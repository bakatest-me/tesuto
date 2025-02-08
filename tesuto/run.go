package tesuto

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"tesuto/config"
	"tesuto/tesuto/models"
	"tesuto/tesuto/repositories/jsrepo"

	"github.com/dop251/goja"
	"github.com/fatih/color"
)

func Run(cfg config.Env, filePath []string) {
	for _, path := range filePath {
		fmt.Println("File: ", path)
		rawJS, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		vm := goja.New()
		engine, err := jsrepo.NewEngine(vm, string(rawJS))
		if err != nil {
			log.Fatalf("Error creating engine: %v", err)
		}

		t := NewExecutor(cfg, engine)
		info, err := t.Run(string(rawJS))
		if err != nil {
			log.Fatalf("Error running tesuto: %v", err)
		}

		PrintResult(cfg, info)
	}
}

func PrintResult(cfg config.Env, info models.TestCaseInfo) {
	color.White("%s %s\n", color.CyanString(info.Setup.Method), color.YellowString(info.Setup.URL))
	for k, v := range info.Setup.Headers {
		color.White("%s: %s\n", color.CyanString(k), color.YellowString(v))
	}
	for _, r := range info.Results {

		respTimeFmt := color.WhiteString("[%v]", r.Resp.Time)
		if r.Error != nil {
			color.Red("✘ %s %v", r.Name, respTimeFmt)
			color.White("    " + r.Error.Error())
			continue
		}

		if r.IsPass {
			color.Green("✔ %s %v", r.Name, respTimeFmt)
		} else {
			color.Red("✘ " + r.Name)
		}

		if !cfg.Debug {
			continue
		}

		color.White("-----------------------------")
		color.White("Request:")
		color.White("-----------------------------")
		color.White("Method: %s", r.Req.Method)
		color.White("URL: %s", r.Req.URL)
		color.White("Headers:")
		for k, v := range r.Req.Headers {
			color.White("    %s: %s", k, v)
		}

		b, _ := json.MarshalIndent(r.Resp.Body, "", "  ")
		color.White("-----------------------------")
		color.White("Response:")
		color.White("-----------------------------")
		color.White("Status: %d", r.Resp.Status)
		color.White("Time: %v", r.Resp.Time)
		body := string(b)
		if body == "null" {
			body = ""
		}
		color.White("Body: \n%s", body)
		color.White("-----------------------------")
		fmt.Println()
	}

	fmt.Println()
}
