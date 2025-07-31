//go:build js && wasm

package signal

type Signal[T any] struct {
	value    T
	c        int
	effects  map[int]func()
	equalsFn func(a, b T) bool
}

// Creates a new signal and returns a cleanup function to remove all effects from this signal.
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

func (s *Signal[T]) Get() T {
	return s.value
}

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
func (s *Signal[T]) Effect(effect func()) func() {
	s.c++
	id := s.c
	s.effects[id] = effect

	return func() {
		delete(s.effects, id)
	}
}
