//go:build js && wasm

// [EXPERIMENTAL]
package react

type State[T any] struct {
	value    T
	c        int
	effects  map[int]func()
	equalsFn func(a, b T) bool
}

type StateGetter[T any] func() T
type StateSetter[T any] func(T)

func (s *State[T]) Get() T {
	return s.value
}

func (s *State[T]) Set(new T) {
	if s.equalsFn != nil && s.equalsFn(s.value, new) {
		return
	}
	s.value = new
	for _, effect := range s.effects {
		effect()
	}
}

func (s *State[T]) RegisterEffect(effect func()) {
	s.c++
	s.effects[s.c] = effect
}

func UseState[T any](initialState T, equalsFn func(a, b T) bool) (*State[T], StateGetter[T], StateSetter[T]) {
	state := &State[T]{
		value:    initialState,
		effects:  make(map[int]func()),
		equalsFn: equalsFn,
	}

	getter := func() T {
		return state.value
	}

	setter := func(new T) {
		state.Set(new)
	}

	return state, getter, setter
}

func UseEffect[T any](sideEffect func(), dependencies ...*State[T]) {
	for _, dep := range dependencies {
		dep.RegisterEffect(sideEffect)
	}
}
