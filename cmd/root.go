package cmd

import (
	// "bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	server "github.com/dhth/dstll/server"
	"github.com/dhth/dstll/tsutils"
	"github.com/dhth/dstll/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	configFileName = "dstll/dstll.toml"
	envPrefix      = "DSTLL"
)

var (
	errCouldntGetHomeDir      = errors.New("couldn't get home directory")
	errCouldntGetConfigDir    = errors.New("couldn't get config directory")
	errConfigFileExtIncorrect = errors.New("config file must be a TOML file")
	errConfigFileDoesntExist  = errors.New("config file does not exist")
	errCouldntCreateDirectory = errors.New("could not create directory")
	errNoPathsProvided        = errors.New("no file paths provided")
)

func Execute() error {
	rootCmd, err := NewRootCommand()
	if err != nil {
		return err
	}

	err = rootCmd.Execute()
	return err
}

func NewRootCommand() (*cobra.Command, error) {
	var (
		trimPrefix     string
		fPaths         []string
		userHomeDir    string
		configFilePath string
		configPathFull string

		// root
		plainOutput bool

		// write
		writeOutputDir string
		writeQuiet     bool

		// tui
		tuiViewFileCmdInput []string
		tuiViewFileCmd      []string
	)

	rootCmd := &cobra.Command{
		Use:   "dstll [PATH ...]",
		Short: "dstll gives you a high level overview of various constructs in your code",
		Long: `dstll gives you a high level overview of various constructs in your code.

Its findings can be printed to stdout, written to files, or presented via a TUI/web interface.
`,
		SilenceUsage: true,
		Args:         cobra.MinimumNArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			configPathFull = expandTilde(configFilePath, userHomeDir)

			if filepath.Ext(configPathFull) != ".toml" {
				return errConfigFileExtIncorrect
			}
			_, err := os.Stat(configPathFull)

			fl := cmd.Flags()
			if fl != nil {
				cf := fl.Lookup("config-path")
				if cf != nil && cf.Changed && errors.Is(err, fs.ErrNotExist) {
					return errConfigFileDoesntExist
				}
			}

			var v *viper.Viper
			v, err = initializeConfig(cmd, configPathFull)
			if err != nil {
				return err
			}

			calledAs := cmd.CalledAs()
			if calledAs == "tui" {
				// pretty ugly hack to get around the fact that
				// v.GetStringSlice("view_file_command") always seems to prioritize the config file
				if len(tuiViewFileCmdInput) > 0 && len(tuiViewFileCmdInput[0]) > 0 && !strings.HasPrefix(tuiViewFileCmdInput[0], "[") {
					tuiViewFileCmd = tuiViewFileCmdInput
				} else {
					tuiViewFileCmd = v.GetStringSlice("view-file-command")
				}
				return nil
			}

			fPaths = args
			if len(fPaths) == 0 {
				return errNoPathsProvided
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			results := tsutils.GetResults(fPaths)
			if len(results) == 0 {
				return
			}

			ui.ShowResults(results, trimPrefix, plainOutput)
		},
	}

	writeCmd := &cobra.Command{
		Use:   "write [PATH ...]",
		Short: "Write findings to files",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			results := tsutils.GetResults(fPaths)
			if len(results) == 0 {
				return nil
			}

			err := createDir(writeOutputDir)
			if err != nil {
				return fmt.Errorf("%w: %s", errCouldntCreateDirectory, err.Error())
			}

			ui.WriteResults(results, writeOutputDir, writeQuiet)
			return nil
		},
	}

	tuiCmd := &cobra.Command{
		Use:   "tui",
		Short: "Open dstll TUI",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ui.Config{
				ViewFileCmd: tuiViewFileCmd,
			}
			return ui.RenderUI(config)
		},
	}

	serveCmd := &cobra.Command{
		Use:   "serve [PATH ...]",
		Short: "Serve findings via a web server",
		Args:  cobra.MinimumNArgs(1),
		Run: func(_ *cobra.Command, _ []string) {
			server.Start(fPaths)
		},
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errCouldntGetHomeDir, err.Error())
	}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errCouldntGetConfigDir, err.Error())
	}

	var emptySlice []string
	defaultConfigPath := filepath.Join(userConfigDir, configFileName)

	// rootCmd.Flags().StringSliceVarP(&fPaths, "files", "f", emptySlice, "paths of files to run dstll on")
	rootCmd.Flags().StringVarP(&configFilePath, "config-path", "c", defaultConfigPath, "location of dstll's config file")
	rootCmd.Flags().StringVarP(&trimPrefix, "trim-prefix", "t", "", "prefix to trim from the file path")
	rootCmd.Flags().BoolVarP(&plainOutput, "plain", "p", false, "output plain text")

	writeCmd.Flags().StringVarP(&writeOutputDir, "output-dir", "o", "dstll-output", "directory to write findings in")
	writeCmd.Flags().BoolVarP(&writeQuiet, "quiet", "q", false, "suppress output")

	tuiCmd.Flags().StringSliceVar(&tuiViewFileCmdInput, "view-file-command", emptySlice, "command to use to view, eg. --view-file-command='bat,--style,plain,--paging,always'")

	rootCmd.AddCommand(writeCmd)
	rootCmd.AddCommand(tuiCmd)
	rootCmd.AddCommand(serveCmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd, nil
}

func initializeConfig(cmd *cobra.Command, configFile string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filepath.Base(configFile))
	v.SetConfigType("toml")
	v.AddConfigPath(filepath.Dir(configFile))

	err := v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return v, err
	}

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	err = bindFlags(cmd, v)
	if err != nil {
		return v, err
	}

	return v, nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) error {
	var err error
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := strings.ReplaceAll(f.Name, "-", "_")

		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			fErr := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if fErr != nil {
				err = fErr
				return
			}
		}
	})
	return err
}
