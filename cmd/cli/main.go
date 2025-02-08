package main

import (
	"fmt"
	"os"
	"tesuto/config"
	"tesuto/pkg/util"
	"tesuto/tesuto"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "run",
		Run: CmdRun,
	}
	rootCmd.Flags().BoolP("debug", "d", false, "debug mode")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func CmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		cmd.Println("Usage: tesuto run <path to JavaScript file or directory>")
		return
	}

	filePath := args[1]
	files, err := util.GetFilePath(filePath)
	if err != nil {
		cmd.Println("Error: ", err)
		return
	}

	isDebug, _ := cmd.Flags().GetBool("debug")
	cfg := config.Env{
		Debug: isDebug,
	}
	tesuto.Run(cfg, files)
}
