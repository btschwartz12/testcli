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
	Use: "testy",
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

var completionCmd = &cobra.Command{
	Use:       "completion [bash|zsh]",
	ValidArgs: []string{"bash", "zsh"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]
		switch shell {
		case "bash":
			err := rootCmd.GenBashCompletion(os.Stdout)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating bash completion: %v\n", err)
				os.Exit(1)
			}
		case "zsh":
			err := rootCmd.GenZshCompletion(os.Stdout)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating zsh completion: %v\n", err)
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "Unsupported shell: %s. Use 'bash' or 'zsh'\n", shell)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(completionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
