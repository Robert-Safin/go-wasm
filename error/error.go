//go:build js && wasm

package error

import "syscall/js"

type Error struct {
	Value js.Value
}

func ToError(v js.Value) *Error {
	return &Error{v}
}

func (e *Error) Message() string {
	if e.Value.Type() == js.TypeObject && e.Value.Get("message").Type() == js.TypeString {
		return e.Value.Get("message").String()
	}
	return e.Value.String()
}

func (e *Error) Name() string {
	if e.Value.Type() == js.TypeObject && e.Value.Get("name").Type() == js.TypeString {
		return e.Value.Get("name").String()
	}
	return ""
}

func (e *Error) Code() string {
	if e.Value.Type() == js.TypeObject && e.Value.Get("code").Type() == js.TypeString {
		return e.Value.Get("code").String()
	}
	return ""
}
