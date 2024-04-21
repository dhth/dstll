package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"flag"

	server "github.com/dhth/dstll/server"
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
	currentUser, err := user.Current()
	var defaultConfigFP string

	if err != nil {
		fmt.Printf("Couldn't determine your home directory, if you want to use the config file, pass the location to it manually (via the -config-file-path flag\n")
	} else {
		defaultConfigFP = fmt.Sprintf("%s/.config/dstll/dstll.toml", currentUser.HomeDir)
	}

	configFilePath := flag.String("config-file-path", defaultConfigFP, "location of dstll's config file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\nFlags:\n", helpText)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n------\n\ndstll's toml config looks like this:\n%s", configSampleFormat)
	}

	flag.Parse()

	if *configFilePath == "" {
		die("config-file-path cannot be empty")
	}

	if *mode == "" {
		die("mode cannot be empty")
	}

	if *trimPrefix != "" && *mode != "cli" {
		die("trim-prefix can only be used in CLI mode")
	}

	if *plain == true && *mode != "cli" {
		die("plain can be true only in CLI mode")
	}
	var defaultConfigPathUsed bool

	if *configFilePath == defaultConfigFP {
		defaultConfigPathUsed = true
	}

	var configFPExpanded string
	if strings.Contains(*configFilePath, "~") {
		configFPExpanded, err = expandTilde(*configFilePath)
		if err != nil {
			die("Something went horribly wrong. Let @dhth know about this error on github: ", err.Error())
		}
	} else {
		configFPExpanded = *configFilePath
	}

	var cfg dstllConfig
	var cfgErr error
	_, err = os.Stat(configFPExpanded)

	if os.IsNotExist(err) {
		if !defaultConfigPathUsed {
			die(cfgErrSuggestion(fmt.Sprintf("Error: file doesn't exist at %q", configFPExpanded)))
		}
	}

	if err == nil {
		cfg, cfgErr = readConfig(configFPExpanded)
		if cfgErr != nil {
			die(cfgErrSuggestion(fmt.Sprintf("Error reading config: %s", cfgErr.Error())))
		}
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
		server.Start(fPaths)
	case "tui":
		config := ui.Config{
			ViewFileCmd: cfg.TUICfg.ViewFileCmd,
		}
		ui.RenderUI(config)
	}
}

func expandTilde(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return strings.Replace(path, "~", usr.HomeDir, 1), nil
	}
	return path, nil
}
