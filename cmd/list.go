package cmd

import (
  "fmt"
  "os"

  "gopkg.in/yaml.v2"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
)


var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List all the available stations",
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

    for key, station := range cfg.Stations {
      fmt.Printf("%-20v%-10v\n", key, station.Title)
    }
  },
}

func init() {
  rootCmd.AddCommand(listCmd)
}