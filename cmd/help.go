package cmd

import "fmt"

var (
	helpText = `dstll gives you a high level overview of various constructs in your code.
It can run in both CLI and TUI mode, dictated by the -mode flag.
Additionally, it can also present its findings via an HTTP server; use -mode=server for that.

Usage: dstll [flags]
`
	configSampleFormat = `
[tui]
view_file_command = ["your", "command"]
# for example, ["bat", "--style", "plain", "--paging", "always"]
# will run 'bat --style plain --paging always <file-path>'
`
)

func cfgErrSuggestion(msg string) string {
	return fmt.Sprintf(`%s

Make sure to structure dstll's toml config file as follows:

------
%s
------

Use "dstll -help" for more information`,
		msg,
		configSampleFormat,
	)
}
