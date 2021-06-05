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
		dirs := make([]string, 0)
		files := make([]string, 0)
		if ifDotExist(args) {
			log.Println("Auto commenting the current directory and ignoring other arguments")
			pkg.AutoCommentDir(".")
		} else {
			for index := range args {
				if isDir(args[index]) {
					dirs = append(dirs, args[index])
				} else if isFileExist(args[index]) {
					files = append(files, args[index])
				}
			}
		}

		for index := range dirs {
			pkg.AutoCommentDir(dirs[index])
		}

		pkg.AutoCommentFiles(files...)
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
