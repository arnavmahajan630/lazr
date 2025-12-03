/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/arnavmahajan630/lazr/internal/extract"
	"github.com/spf13/cobra"
)

var (
	exFlagDomains    bool
	exFlagSubdomains bool
	exFlagIPs        bool
	exFlagEmails     bool
	exFlagPhones     bool
	exFlagAll        bool
	exFlagOutDir     string
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract <input-file>",
	Short: "Extract Desired data from file",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := args[0]

		// validating the path through status
		if _, err := os.Stat(inputPath); err != nil {
			return fmt.Errorf("cannot access input file %q: %w", inputPath, err)
		}

		outdir := exFlagOutDir
		// set the defualt output
		if outdir == "" {
			outdir = "./lazr_output"
		}
		cfg := extract.Config{
			InputPath:      inputPath,
			WantDomains:    exFlagAll || exFlagDomains,
			WantSubdomains: exFlagAll || exFlagSubdomains,
			WantIPs:        exFlagAll || exFlagIPs,
			WantEmails:     exFlagAll || exFlagEmails,
			WantPhones:     exFlagAll || exFlagPhones,
			BaseOutDir: 	exFlagOutDir,
		}

		// if no flag was set return with error
		if !cfg.AnyRequested() {
			return fmt.Errorf("no extract type specified")
		}	

		// start execution
		_, outDir, err := extract.Run(cfg)
		if err != nil {
			return err
		}

		// executed successfully !
		fmt.Printf("✔ Extraction completed\n📁 Output: %s\n", outDir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// set flags
	extractCmd.Flags().BoolVar(&exFlagDomains, "domains", false, "Extract domains")
	extractCmd.Flags().BoolVar(&exFlagSubdomains, "subdomains", false, "Extract Sub-domains")
	extractCmd.Flags().BoolVar(&exFlagIPs, "ips", false, "Extract IP addresses")
	extractCmd.Flags().BoolVar(&exFlagEmails, "emails", false, "Extract email addresses")
	extractCmd.Flags().BoolVar(&exFlagPhones, "phones", false, "Extract phone numbers")
	extractCmd.Flags().BoolVar(&exFlagAll, "all", false, "Extract all supported artifacts")
	extractCmd.Flags().StringVar(&exFlagOutDir, "outdir", "", "Base output directory")
}
