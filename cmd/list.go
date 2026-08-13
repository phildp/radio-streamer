package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available stations",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		for key, station := range cfg.Stations {
			fmt.Printf("%-20v%-10v\n", key, station.Title)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
