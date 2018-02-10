package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var prefix string
var suffix string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cubase-rename",
	Short: "Rename wav files generated by \"Channel Batch Export\" function of cubase",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(".")

		for _, file := range files {
			if file.IsDir() || filepath.Ext(file.Name()) != ".wav" {
				continue
			}

			fileNameWithoutExt := strings.Replace(file.Name(), ".wav", "", -1)
			splittedFileName := strings.Split(fileNameWithoutExt, " - ")
			trackName := splittedFileName[len(splittedFileName)-1]
			newFileName := prefix + trackName + suffix + ".wav"
			err = os.Rename(file.Name(), newFileName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%q is renamed to %q\n", file.Name(), newFileName)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "new prefix of file name")
	rootCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "new suffix of file name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
