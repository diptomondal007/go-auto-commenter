package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:	"auto-commenter",
	Short:	"auto commenter",
	Long:	"auto commenter is a tool to automatically comment out all the exported functions",
	Run: func(cmd *cobra.Command, args []string) {
		if ifDotExist(args) {
			autoCommentDir(".")
		}
	},
}

type A struct {
}

// TODO: document exported function
func NewA() {

}

// TODO: document exported function
func (a *A) New() {

}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
