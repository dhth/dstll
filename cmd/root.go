package cmd

import (
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
	mode       = flag.String("mode", "cli", "mode to operate in; possible values: cli/tui")
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

	switch *mode {
	case "cli":
		ui.ShowResults(*trimPrefix, *plain)
	case "tui":
		ui.RenderUI()
	}
}
