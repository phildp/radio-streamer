package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/phildp/radio-streamer/internal/player"
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

		fmt.Println("Starting radio service...")
		fmt.Printf("Listening to %s (volume: %d%%)\n", rs.Title, int(volume*100))
		fmt.Println("Press Ctrl+C to stop.")

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		if err := player.Play(ctx, rs.Filename, volume); err != nil && err != context.Canceled {
			return fmt.Errorf("play stream: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&station, "station", "s", "", "Set radio station")
	startCmd.Flags().Float64VarP(&volume, "volume", "v", 0.5, "Set the volume (float [0,1])")
}
