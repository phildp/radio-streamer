package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var station string
var volume float64

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the stream",
	RunE: func(cmd *cobra.Command, args []string) error {
		if station == "" {
			return fmt.Errorf("--station is required")
		}

		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		rs, ok := cfg.Stations[station]
		if !ok {
			return fmt.Errorf("radio station %q doesn't exist in config", station)
		}

		vlc := exec.Command("vlc", "-I", "dummy", rs.Filename, fmt.Sprintf("--gain=%g", volume))
		fmt.Println("Starting radio service...")
		fmt.Printf("Listening to %s (volume: %d%%)\n", rs.Title, int(volume*100))
		if err := vlc.Run(); err != nil {
			return fmt.Errorf("run vlc: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&station, "station", "s", "", "Set radio station")
	startCmd.Flags().Float64VarP(&volume, "volume", "v", 0.5, "Set the volume (float [0,1])")
}
