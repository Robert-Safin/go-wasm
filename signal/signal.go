//go:build js && wasm

package signal

import (
	"syscall/js"
)

type EffectFunc func()

type Signal[T any] struct {
	value       T
	subscribers []EffectFunc
}

func NewSignal[T any](initial T) *Signal[T] {
	return &Signal[T]{
		value:       initial,
		subscribers: []EffectFunc{},
	}
}

func (s *Signal[T]) Get() T {
	if currentEffect != nil {
		s.subscribers = append(s.subscribers, currentEffect)
	}
	return s.value
}

func (s *Signal[T]) Set(v T) {
	s.value = v
	subs := append([]EffectFunc{}, s.subscribers...)

	for _, sub := range subs {
		sub()
	}
}

var currentEffect EffectFunc

func Effect(f EffectFunc) {
	var wrapped EffectFunc
	wrapped = func() {
		prev := currentEffect
		currentEffect = wrapped
		defer func() { currentEffect = prev }()
		f()
	}
	wrapped()
}

func BindText(el js.Value, sig *Signal[string]) {
	Effect(func() {
		el.Set("textContent", sig.Get())
	})
}
