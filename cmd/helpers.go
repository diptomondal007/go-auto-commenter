package cmd

import "os"

func ifDotExist(args []string) bool {
	for _, v := range args {
		if v == "." || v == "./" {
			return true
		}
	}
	return false
}

func isDir(file string) bool {
	fi, err := os.Stat(file)
	return err == nil && fi.IsDir()
}

func isFileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
