package models

import "github.com/dop251/goja"

type TestCase struct {
	Name     string            `json:"name"`
	Body     goja.Value        `json:"body"`
	Params   map[string]string `json:"params"`
	Query    map[string]string `json:"query"`
	Expected goja.Value        `json:"expected"`
}

type Setting struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}
