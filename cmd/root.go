package cmd

import (
	"fmt"

	"github.com/phildp/radio-streamer/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var configPath string

var rootCmd = &cobra.Command{
	Use:   "radio",
	Short: "RadioStreamer CLI",
	Long:  `Listen to your favourite online through your command line.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/radio/stations.yml)")
}

func initConfig() {
	if cfgFile != "" {
		configPath = cfgFile
		return
	}

	path, err := config.DefaultPath()
	if err != nil {
		cobra.CheckErr(fmt.Errorf("resolve config path: %w", err))
	}
	configPath = path
}

func loadConfig() (*config.Config, error) {
	return config.Load(configPath)
}
