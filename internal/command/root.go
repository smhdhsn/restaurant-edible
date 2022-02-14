package command

import (
	"github.com/spf13/cobra"
)

// rootCMD is the base command when called without any subcommands.
var rootCMD = &cobra.Command{
	Short:             "This app is a food ordering service.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// Execute runs requested commands and handles errors gracefully.
func Execute() {
	cobra.CheckErr(rootCMD.Execute())
}
