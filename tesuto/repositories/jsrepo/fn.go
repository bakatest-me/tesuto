package jsrepo

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
)

func Console() map[string]func(call goja.FunctionCall) goja.Value {
	console := map[string]func(call goja.FunctionCall) goja.Value{
		"log": func(call goja.FunctionCall) goja.Value {
			// Convert arguments to a Go string and print
			args := make([]interface{}, len(call.Arguments))
			for i, arg := range call.Arguments {
				expArg := arg.Export()
				switch expArg.(type) {
				case map[string]interface{}:
					expArgType, ok := expArg.(map[string]interface{})
					if !ok {
						args[i] = arg.Export() // Export converts to Go types (strings, numbers, etc.)
					}
					json, _ := json.MarshalIndent(expArgType, "", "  ")
					args[i] = string(json)
				default:
					args[i] = expArg
				}
			}
			fmt.Println(args...)
			return goja.Undefined()
		},
	}
	return console
}
