//go:build js && wasm

package signal

type Signal[T any] struct {
	// Underlying value stored.
	value T
	// Internal counter, used as ID for effects.
	c int
	// Map of effects on a signal with c as key.
	// Effects are function to execute whenever value changes.
	effects map[int]func()
	// [Optional] Used to compare previous value to new value. Pass nil to opt out.
	// If equality is found effects will not re-run.
	// If equality is not found or equalsFn is nil, effects will always re-run on value change.
	equalsFn func(a, b T) bool
}

// Creates a new signal and returns a cleanup function that removes all effects from this signal.
// Running cleanup function prevents memory leaks.
func NewSignal[T any](initialState T, equalsFn func(a, b T) bool) (*Signal[T], func()) {
	signal := &Signal[T]{
		value:    initialState,
		c:        0,
		effects:  make(map[int]func()),
		equalsFn: equalsFn,
	}

	cleanup := func() {
		for id := range signal.effects {
			delete(signal.effects, id)
		}
	}

	return signal, cleanup
}

// Get underlying value.
func (s *Signal[T]) Get() T {
	return s.value
}

// First: if equalsFn is present and equality is found, signal update is aborted.
// Second: sets underlying value to new value passed in.
// Last: runs all effects on the signal.
func (s *Signal[T]) Set(new T) {
	if s.equalsFn != nil && s.equalsFn(s.value, new) {
		return
	}
	s.value = new
	for _, effect := range s.effects {
		effect()
	}
}

// Register an effect and return a cleanup function to remove just that effect.
// Cleanup function prevents memory leaks.
func (s *Signal[T]) Effect(effect func()) func() {
	s.c++
	id := s.c
	s.effects[id] = effect

	return func() {
		delete(s.effects, id)
	}
}
