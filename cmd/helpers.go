package cmd

func ifDotExist(args []string) bool {
	for _, v := range args {
		if v == "." || v == "./" {
			return true
		}
	}
	return false
}
