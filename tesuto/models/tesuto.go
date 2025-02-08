package models

import "time"

type Request struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    interface{}         `json:"body,omitempty"`
}

type Response struct {
	Status int           `json:"status"`
	Body   interface{}   `json:"body"`
	Time   time.Duration `json:"time"`
}

type TestCaseResult struct {
	Name   string   `json:"name"`
	IsPass bool     `json:"isPass"`
	Error  error    `json:"error"`
	Req    Request  `json:"req"`
	Resp   Response `json:"resp"`
}

type TestCaseInfo struct {
	Setup   Setting          `json:"setup"`
	Results []TestCaseResult `json:"results"`
}

func ToExpectedParam(statusCode int, result interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": statusCode,
		"body":   result,
	}
}
