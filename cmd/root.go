package cmd

import (
	"bufio"
	"fmt"
	"os"

	"flag"

	"github.com/dhth/dstll/ui"
)

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

var (
	mode       = flag.String("mode", "cli", "mode to operate in; possible values: cli/tui/server")
	trimPrefix = flag.String("trim-prefix", "", "prefix to trim from the file path")
	plain      = flag.Bool("plain", false, "whether to output plain text")
)

func Execute() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\nFlags:\n", helpText)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *mode == "" {
		die("mode cannot be empty")
	}

	if *trimPrefix != "" && *mode != "cli" {
		die("trim-prefix can only be used in CLI mode")
	}

	if *plain == true && *mode != "cli" {
		die("plain can be true only in CLI mode")
	}

	var fPaths []string

	switch *mode {
	case "cli", "server":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fPaths = append(fPaths, scanner.Text())
		}
	}

	switch *mode {
	case "cli":
		ui.ShowResults(fPaths, *trimPrefix, *plain)
	case "server":
		startServer(fPaths)
	case "tui":
		ui.RenderUI()
	}
}
