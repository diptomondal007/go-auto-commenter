package cmd

import (
	"github.com/spf13/cobra"
	"go-auto-commenter/pkg"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:	"auto-commenter",
	Short:	"auto commenter",
	Long:	"auto commenter is a tool to automatically comment out all the exported functions",
	Run: func(cmd *cobra.Command, args []string) {
		if ifDotExist(args) {
			pkg.AutoCommentDir(".")
		} else {

		}
	},
}

type A struct {
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
