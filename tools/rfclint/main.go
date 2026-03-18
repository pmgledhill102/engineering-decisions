// Package main implements rfclint, a validation tool for RFC documents.
//
// Usage:
//
//	rfclint <file.md> [file2.md ...]
//
// rfclint checks that RFC files conform to the required format:
//   - Valid YAML frontmatter with required fields
//   - RFC number matches the filename
//   - All required sections are present and non-empty
//   - Status is a valid lifecycle value
//   - Decision date is present when status requires it
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Frontmatter represents the YAML frontmatter of an RFC document.
type Frontmatter struct {
	RFC          any      `yaml:"rfc"`
	Title        string   `yaml:"title"`
	Status       string   `yaml:"status"`
	Authors      []string `yaml:"authors"`
	Created      string   `yaml:"created"`
	Updated      string   `yaml:"updated"`
	DecisionDate string   `yaml:"decision-date"`
	Supersedes   any      `yaml:"supersedes"`
	SupersededBy any      `yaml:"superseded-by"`
}

// rfcNumber returns the RFC number as an int, or 0 if the value is not numeric.
func (f Frontmatter) rfcNumber() int {
	switch v := f.RFC.(type) {
	case int:
		return v
	case float64:
		return int(v)
	default:
		return 0
	}
}

var validStatuses = map[string]bool{
	"draft":      true,
	"proposed":   true,
	"accepted":   true,
	"rejected":   true,
	"withdrawn":  true,
	"superseded": true,
}

var requiredSections = []string{
	"Summary",
	"Background",
	"Business Justification",
	"Proposal",
	"Alternatives Considered",
	"Migration and Rollout",
	"Risks and Mitigations",
}

// rfcNumberFromFilename extracts the RFC number from a filename like "0001-short-slug.md".
// Returns -1 if the filename does not match the expected pattern.
func rfcNumberFromFilename(filename string) int {
	base := filepath.Base(filename)
	if base == "_template.md" {
		return -1
	}
	re := regexp.MustCompile(`^(\d{4})-`)
	m := re.FindStringSubmatch(base)
	if m == nil {
		return -1
	}
	n, err := strconv.Atoi(m[1])
	if err != nil {
		return -1
	}
	return n
}

// parseFrontmatter extracts and parses YAML frontmatter from the raw file content.
// It returns the parsed frontmatter and any remaining body content after the closing delimiter.
func parseFrontmatter(content string) (Frontmatter, string, error) {
	var fm Frontmatter

	if !strings.HasPrefix(content, "---\n") {
		return fm, "", fmt.Errorf("file does not start with YAML frontmatter delimiter (---)")
	}

	end := strings.Index(content[4:], "\n---")
	if end == -1 {
		return fm, "", fmt.Errorf("no closing YAML frontmatter delimiter (---) found")
	}

	yamlContent := content[4 : 4+end]
	body := content[4+end+4:]

	if err := yaml.Unmarshal([]byte(yamlContent), &fm); err != nil {
		return fm, "", fmt.Errorf("invalid YAML in frontmatter: %w", err)
	}

	return fm, body, nil
}

// isValidDate checks whether a string is a valid YYYY-MM-DD date.
func isValidDate(s string) bool {
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}

// htmlCommentRe matches HTML comments like <!-- ... -->.
var htmlCommentRe = regexp.MustCompile(`<!--[\s\S]*?-->`)

// sectionHasContent checks whether a section body contains meaningful content
// beyond HTML template comments and whitespace.
func sectionHasContent(body string) bool {
	stripped := htmlCommentRe.ReplaceAllString(body, "")
	return strings.TrimSpace(stripped) != ""
}

// extractSections parses the markdown body into a map of H2 heading -> section body.
func extractSections(body string) map[string]string {
	sections := make(map[string]string)
	var currentHeading string
	var currentBody strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "## ") {
			if currentHeading != "" {
				sections[currentHeading] = currentBody.String()
			}
			currentHeading = strings.TrimPrefix(line, "## ")
			currentBody.Reset()
		} else if currentHeading != "" {
			currentBody.WriteString(line)
			currentBody.WriteString("\n")
		}
	}
	if currentHeading != "" {
		sections[currentHeading] = currentBody.String()
	}

	return sections
}

// extractH1 returns the first H1 heading found in the body, or empty string if none.
func extractH1(body string) string {
	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") && !strings.HasPrefix(line, "## ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return ""
}

// Validate checks an RFC file and returns a list of errors found.
func Validate(filename, content string) []string {
	var errs []string

	fm, body, err := parseFrontmatter(content)
	if err != nil {
		return []string{err.Error()}
	}

	isTemplate := filepath.Base(filename) == "_template.md"

	// Validate RFC number.
	if !isTemplate {
		rfcNum := fm.rfcNumber()
		if rfcNum <= 0 {
			errs = append(errs, "rfc number must be a positive integer")
		}
		fileNum := rfcNumberFromFilename(filename)
		if fileNum >= 0 && rfcNum != fileNum {
			errs = append(errs, fmt.Sprintf("rfc number %d does not match filename number %d", rfcNum, fileNum))
		}
	}

	// Validate title.
	if strings.TrimSpace(fm.Title) == "" {
		errs = append(errs, "title is required")
	}

	// Validate title matches H1.
	if !isTemplate && strings.TrimSpace(fm.Title) != "" {
		h1 := extractH1(body)
		if h1 != "" && h1 != fm.Title {
			errs = append(errs, fmt.Sprintf("frontmatter title %q does not match H1 heading %q", fm.Title, h1))
		}
	}

	// Validate status.
	if !validStatuses[fm.Status] {
		errs = append(errs, fmt.Sprintf("invalid status %q; must be one of: draft, proposed, accepted, rejected, withdrawn, superseded", fm.Status))
	}

	// Validate authors.
	if len(fm.Authors) == 0 {
		errs = append(errs, "at least one author is required")
	}

	// Validate created date.
	if !isTemplate {
		if fm.Created == "" {
			errs = append(errs, "created date is required")
		} else if !isValidDate(fm.Created) {
			errs = append(errs, fmt.Sprintf("created date %q is not valid YYYY-MM-DD", fm.Created))
		}
	}

	// Validate decision-date is present for terminal statuses.
	if fm.Status == "accepted" || fm.Status == "rejected" {
		if fm.DecisionDate == "" {
			errs = append(errs, fmt.Sprintf("decision-date is required when status is %q", fm.Status))
		} else if !isValidDate(fm.DecisionDate) {
			errs = append(errs, fmt.Sprintf("decision-date %q is not valid YYYY-MM-DD", fm.DecisionDate))
		}
	}

	// Validate required sections.
	sections := extractSections(body)
	for _, req := range requiredSections {
		sectionBody, found := sections[req]
		if !found {
			errs = append(errs, fmt.Sprintf("missing required section: %q", req))
		} else if !isTemplate && !sectionHasContent(sectionBody) {
			errs = append(errs, fmt.Sprintf("section %q is empty (contains only comments or whitespace)", req))
		}
	}

	return errs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: rfclint <file.md> [file2.md ...]\n")
		os.Exit(2)
	}

	exitCode := 0
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename) //nolint:gosec // CLI tool intentionally reads user-specified files
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", filename, err)
			exitCode = 1
			continue
		}

		errs := Validate(filename, string(data))
		if len(errs) > 0 {
			for _, e := range errs {
				fmt.Fprintf(os.Stderr, "%s: %s\n", filename, e)
			}
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
