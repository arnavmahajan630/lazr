package extract

import "regexp"

var (
	regDomain   = regexp.MustCompile(`\b(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}\b`)
	reSubdomain = regexp.MustCompile(`\b(?:[a-zA-Z0-9-]+\.){2,}[a-zA-Z]{2,}\b`)
	reIP        = regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	reEmail     = regexp.MustCompile(`\b[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}\b`)
	rePhone     = regexp.MustCompile(`\b\+?\d[\d\s\-]{7,}\d\b`)
)

func findDomains(s string) []string {
	return regDomain.FindAllString(s, -1)
}
func findSubdomains(s string) []string {
	return reSubdomain.FindAllString(s, -1)
}
func findPhone(s string) []string {
	return rePhone.FindAllString(s, -1)
}
func findEmail(s string) []string {
	return reEmail.FindAllString(s, -1)
}
func findIps(s string) []string {
	return reIP.FindAllString(s, -1)
}
