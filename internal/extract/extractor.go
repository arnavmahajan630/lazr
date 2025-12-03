package extract

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	InputPath      string
	WantDomains    bool
	WantSubdomains bool
	WantIPs        bool
	WantEmails     bool
	WantPhones     bool
	BaseOutDir     string
}

func (c Config) AnyRequested() bool {
	return c.WantDomains || c.WantSubdomains || c.WantIPs || c.WantEmails || c.WantPhones
}

type Result struct {
	Domains    []string
	SubDomains []string
	IPs        []string
	Emails     []string
	Phones     []string
}

func Run(cfg Config) (Result, string, error) {
	// return error if the input file is not present
	f, err := os.Open(cfg.InputPath)
	if err != nil {
		return Result{}, "", fmt.Errorf("open input file: %w", err)
	}

	defer f.Close()

	// set's for unique vlaue of result
	domainsSet := map[string]struct{}{}
	subdomainsSet := map[string]struct{}{}
	ipsSet := map[string]struct{}{}
	emailsSet := map[string]struct{}{}
	phonesSet := map[string]struct{}{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if cfg.WantDomains {
			for _, d := range findDomains(line){
				domainsSet[d] = struct{}{}
			}
		}
		if cfg.WantSubdomains {
			for _, d := range findSubdomains(line){
				subdomainsSet[d] = struct{}{}
			}
		}
		if cfg.WantEmails {
			for _, d := range findEmail(line){
				emailsSet[d] = struct{}{}
			}
		}
		if cfg.WantIPs {
			for _, d := range findIps(line){
				ipsSet[d] = struct{}{}
			}
		}
		if cfg.WantPhones {
			for _, d := range findPhone(line){
				phonesSet[d] = struct{}{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return Result{}, "", fmt.Errorf("scan input: %w", err)
	}

	// fill the result with slices

	res := Result{
		Domains:    setToSortedSlice(domainsSet),
		SubDomains: setToSortedSlice(subdomainsSet),
		IPs:        setToSortedSlice(ipsSet),
		Emails:     setToSortedSlice(emailsSet),
		Phones:     setToSortedSlice(phonesSet),
	}

	// Generate directory path
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	outDir := filepath.Join(cfg.BaseOutDir, "extract_"+timestamp)

	// Create Directories and Write them
	resul, path , err := createOutDirectory(cfg, outDir, res)
	return resul, path , err
}
