package main

import (
	"github.com/spf13/cobra"
)

// rootCMD is the base command when called without any subcommands.
var rootCMD = &cobra.Command{
	Short:             "This app is a food ordering service.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// main is the main application entry.
func main() {
	cobra.CheckErr(rootCMD.Execute())
}
