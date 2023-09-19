/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type Config struct {
	Stations map[string]struct {
		Title    string `yaml:"title"`
		Filename string `yaml:"filename"`
	} `yaml:"stations"`
}

var station string
var volume float64

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the stream",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		f, err := os.Open(home + "/.config/radio/stations.yml")
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		defer f.Close()
		var cfg Config
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		rs, ok := cfg.Stations[station]
		if !ok {
			fmt.Println("radio station doesn't exist in stations.yml")
			os.Exit(2)
		}
		filename := rs.Filename

		cmad := exec.Command("vlc", "-I", "dummy", filename, fmt.Sprint("--gain=", volume))
		fmt.Println("Starting radio service...")
		fmt.Printf("Listening to %s (volume: %d%%)\n", rs.Title, int(volume*100))
		err = cmad.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&station, "station", "s", "", "Set radio station")
	startCmd.Flags().Float64VarP(&volume, "volume", "v", 0.5, "Set the volume (float [0,1])")
}
