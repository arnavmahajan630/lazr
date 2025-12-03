/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

var version = "v0.1.0"
var rootCmd = &cobra.Command{
	Use:   "lazr",
	Short: "recon helper for most used recon activites",
	Long: `Lazr is a tool made make your recon journey simpler. It has all the small helper functions
	that you would need for cleaning and sorting for your day to day recon activites`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	 Run: func(cmd *cobra.Command, args []string) {
		printBanner()
	  },
	  Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lazr.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


const banner = ` _       ___   ____________ 
| |     / _ \ |___  /| ___ \
| |    / /_\ \   / / | |_/ /
| |    |  _  |  / /  |    / 
| |____| | | |./ /___| |\ \ 
\_____/\_| |_/\_____/\_| \_|
                            
                            `

func printBanner() {
	fmt.Println(banner)
	fmt.Printf("Version: %s\n", "0.1.0")
	fmt.Printf("Author: Arnav Mahajan\n")
	fmt.Println("You are welcome to contribute. Do Star repo if you like the it")	
	fmt.Println()

	fmt.Println("Check out -h | --help to get started :)")
}


