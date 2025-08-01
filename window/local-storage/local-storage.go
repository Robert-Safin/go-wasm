//go:build js && wasm

package localStroage

import "syscall/js"

type LocalStorage struct {
	Value js.Value
}

func Init() LocalStorage {
	return LocalStorage{js.Global().Get("localStorage")}
}

func (l LocalStorage) SetLocalStorageItem(key, value string) {
	l.Value.Call("setItem", key, value)
}

func (l LocalStorage) GetLocalStorageItem(key string) (string, bool) {
	val := l.Value.Call("getItem", key)
	if !val.Truthy() {
		return "", false
	}
	return val.String(), true
}

func (l LocalStorage) RemoveLocalStorageItem(key string) {
	l.Value.Call("removeItem", key)
}

func (l LocalStorage) ClearLocalStorage() {
	l.Value.Call("clear")
}
