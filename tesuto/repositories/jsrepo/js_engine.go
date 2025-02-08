package jsrepo

import (
	"errors"
	"tesuto/tesuto/models"

	"github.com/dop251/goja"
)

type engine struct {
	vm *goja.Runtime
}

func NewEngine(vm *goja.Runtime, rawJS string) (*engine, error) {
	if err := vm.Set("console", Console()); err != nil {
		return nil, err
	}

	_, err := vm.RunString(rawJS)
	if err != nil {
		return nil, err
	}

	return &engine{vm: vm}, nil
}

func (m *engine) SafeGet(key string) SafeGet {
	return NewSafeGet(m.vm, key)
}

func (m *engine) GetSetup() (models.Setting, error) {
	setting := m.SafeGet("setup").ToObject()

	method := setting.SafeGet("method").String()
	url := setting.SafeGet("url").String()

	headers := setting.SafeGet("headers").ToObject()
	headersMap := map[string]string{}
	for _, key := range headers.Keys() {
		headersMap[key] = headers.SafeGet(key).String()
	}

	return models.Setting{
		Method:  method,
		URL:     url,
		Headers: headersMap,
	}, nil
}

func (m *engine) GetTestCases() ([]models.TestCase, error) {
	tc := m.SafeGet("testcase").ToObject()

	listTestCase := []models.TestCase{}
	for _, key := range tc.Keys() {
		testCase := tc.SafeGet(key).ToObject()
		expected := testCase.SafeGet("expected").Raw()

		params := testCase.SafeGet("params").ToObject()
		paramsMap := map[string]string{}
		for _, key := range params.Keys() {
			paramsMap[key] = params.SafeGet(key).String()
		}

		query := testCase.SafeGet("query").ToObject()
		queryMap := map[string]string{}
		for _, key := range query.Keys() {
			queryMap[key] = query.SafeGet(key).String()
		}

		listTestCase = append(listTestCase, models.TestCase{
			Name:     key,
			Body:     testCase.SafeGet("body").Raw(),
			Params:   paramsMap,
			Query:    queryMap,
			Expected: expected,
		})
	}
	return listTestCase, nil
}

func (m *engine) ExpectedFn(expFn goja.Value, resp map[string]interface{}) (bool, error) {
	expectedFn, ok := goja.AssertFunction(expFn)
	if !ok {
		return false, errors.New("error expected function")
	}

	result, err := expectedFn(goja.Undefined(), m.vm.ToValue(resp))
	if err != nil {
		return false, err
	}

	return result.ToBoolean(), nil
}
