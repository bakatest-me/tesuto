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
	rootCmd := &cobra.Command{}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "tesuto run <file or directory>",
		Run:   CmdRun,
	}
	runCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")
	runCmd.PersistentFlags().BoolP("curl", "c", false, "generate curl command")

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tesuto CLI v0.0.1")
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func CmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Println("Usage: tesuto run <path to JavaScript file or directory>")
		return
	}

	filePath := args[0]
	files, err := util.GetFilePath(filePath)
	if err != nil {
		cmd.Println("Error: ", err)
		return
	}

	isDebug, _ := cmd.Flags().GetBool("debug")
	isCurl, _ := cmd.Flags().GetBool("curl")
	cfg := config.Env{
		Debug:           isDebug,
		GenerateCurlCmd: isCurl,
	}
	tesuto.Run(cfg, files)
}
