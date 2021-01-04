// Package dummyhandler
package dummyhandler

type DummyHandler struct {
	ret string
}

func (h DummyHandler) Run(interface{}) (string, error) {
	return "dummy response", nil
}
