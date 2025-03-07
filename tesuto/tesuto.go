package tesuto

import (
	"strings"
	"tesuto/config"
	"tesuto/tesuto/consts"
	"tesuto/tesuto/models"
	"time"

	"github.com/dop251/goja"
	"resty.dev/v3"
)

type JsRepository interface {
	GetSetup() (models.Setting, error)
	GetTestCases() ([]models.TestCase, error)
	ExpectedFn(expFn goja.Value, resp map[string]interface{}) (bool, error)
}

type executor struct {
	engine JsRepository
	cfg    config.Env
}

func NewExecutor(cfg config.Env, engine JsRepository) *executor {
	return &executor{
		cfg:    cfg,
		engine: engine,
	}
}

func (m executor) Run(rawJS string) (models.TestCaseInfo, error) {
	engine := m.engine
	setup, err := engine.GetSetup()
	if err != nil {
		return models.TestCaseInfo{}, err
	}

	testCases, err := engine.GetTestCases()
	if err != nil {
		return models.TestCaseInfo{}, err
	}

	r := resty.New()
	testCaseResults := []models.TestCaseResult{}
	for _, v := range testCases {
		tcResult := models.TestCaseResult{
			Name: v.Name,
		}

		var result interface{}
		req := r.R().
			EnableTrace().
			SetPathParams(v.Params).
			SetQueryParams(v.Query).
			SetHeaders(setup.Headers).
			SetBody(v.Body).
			SetResult(&result).
			EnableGenerateCurlCmd()

		if m.cfg.GenerateCurlCmd {
			req = req.SetTimeout(1 * time.Nanosecond)
		}

		resp, err := req.Execute(setup.Method, setup.URL)

		if m.cfg.GenerateCurlCmd {
			tcResult.CurlCmd = strings.ReplaceAll(req.CurlCmd(), consts.HeaderAcceptEncoding, "")
			testCaseResults = append(testCaseResults, tcResult)
			continue
		}

		if err != nil {
			tcResult.Error = err
			testCaseResults = append(testCaseResults, tcResult)
			continue
		}

		statusCode := resp.StatusCode()
		param := models.ToExpectedParam(statusCode, result)
		isPass, err := engine.ExpectedFn(v.Expected, param)
		if err != nil {
			tcResult.Error = err
		}

		tcResult.Req = models.Request{
			Method:  resp.Request.Method,
			URL:     resp.Request.URL,
			Headers: resp.Request.Header,
		}
		if v.Body != nil {
			tcResult.Req.Body = resp.Request.Body
		}
		tcResult.Resp = models.Response{
			Status: statusCode,
			Body:   result,
			Time:   resp.Request.TraceInfo().ResponseTime,
		}
		tcResult.IsPass = isPass
		testCaseResults = append(testCaseResults, tcResult)
	}

	return models.TestCaseInfo{
		Setup:   setup,
		Results: testCaseResults,
	}, nil
}
