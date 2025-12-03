package extract

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func setToSortedSlice(m map[string]struct{}) []string {
	out := make([]string, 0, len(m))
	for v := range m {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}


// Write lines to files in specified directoy
func writeLines(path string, lines []string) error {
	if len(lines) == 0 {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %q, %w", path, err)
	}

	defer f.Close()

	// create a writer
	writer := bufio.NewWriter(f)

	// for each line write it to the desired path
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("write file %q: %w", path, err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush file %q: %w", path, err)
	}
	return nil

}

// Create directory for each output
func createOutDirectory(cfg Config, outDir string, res Result) (Result, string, error) {

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return Result{}, "", fmt.Errorf("create output dir: %w", err)
	}

	if cfg.WantDomains {
		if err := writeLines(filepath.Join(outDir, "domains.txt"), res.Domains); err != nil {
			return Result{}, "", err
		}
	}
	if cfg.WantSubdomains {
		if err := writeLines(filepath.Join(outDir, "subdomains.txt"), res.SubDomains); err != nil {
			return Result{}, "", err
		}
	}
	if cfg.WantIPs {
		if err := writeLines(filepath.Join(outDir, "ips.txt"), res.IPs); err != nil {
			return Result{}, "", err
		}
	}
	if cfg.WantEmails {
		if err := writeLines(filepath.Join(outDir, "emails.txt"), res.Emails); err != nil {
			return Result{}, "", err
		}
	}
	if cfg.WantPhones {
		if err := writeLines(filepath.Join(outDir, "phones.txt"), res.Phones); err != nil {
			return Result{}, "", err
		}
	}

	return res, outDir, nil
}
