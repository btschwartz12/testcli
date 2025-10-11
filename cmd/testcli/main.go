package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/btschwartz12/testcli/pkg/lib"
	"github.com/btschwartz12/testcli/pkg/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "testcli",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			// print all the args
			for _, arg := range args {
				fmt.Println(arg)
			}
		}
		fmt.Println(lib.SpecialString())
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		info := version.BuildInfo()
		output, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
