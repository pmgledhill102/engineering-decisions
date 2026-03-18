package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func readTestdata(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name)) //nolint:gosec // test helper reads fixtures
	if err != nil {
		t.Fatalf("reading testdata/%s: %v", name, err)
	}
	return string(data)
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		wantErrors []string // substrings that must appear in errors
		wantClean  bool     // if true, expect zero errors
	}{
		{
			name:      "valid draft RFC",
			filename:  "valid.md",
			wantClean: true,
		},
		{
			name:      "valid accepted RFC with decision date",
			filename:  "valid-accepted.md",
			wantClean: true,
		},
		{
			name:     "template file passes validation",
			filename: "_template.md",
			// template is loaded from rfcs/_template.md but we simulate the filename
			wantClean: true,
		},
		{
			name:     "no frontmatter",
			filename: "no-frontmatter.md",
			wantErrors: []string{
				"does not start with YAML frontmatter delimiter",
			},
		},
		{
			name:     "missing required sections",
			filename: "missing-sections.md",
			wantErrors: []string{
				`missing required section: "Background"`,
				`missing required section: "Business Justification"`,
				`missing required section: "Proposal"`,
				`missing required section: "Alternatives Considered"`,
				`missing required section: "Migration and Rollout"`,
				`missing required section: "Risks and Mitigations"`,
			},
		},
		{
			name:     "empty sections (comments only)",
			filename: "empty-sections.md",
			wantErrors: []string{
				`section "Summary" is empty`,
				`section "Background" is empty`,
				`section "Business Justification" is empty`,
				`section "Proposal" is empty`,
				`section "Alternatives Considered" is empty`,
				`section "Migration and Rollout" is empty`,
				`section "Risks and Mitigations" is empty`,
			},
		},
		{
			name:     "invalid status",
			filename: "bad-status.md",
			wantErrors: []string{
				`invalid status "pending"`,
			},
		},
		{
			name:     "accepted without decision date",
			filename: "accepted-no-decision-date.md",
			wantErrors: []string{
				`decision-date is required when status is "accepted"`,
			},
		},
		{
			name:     "title mismatch with H1",
			filename: "0007-title-mismatch.md",
			wantErrors: []string{
				`frontmatter title "Frontmatter Title" does not match H1 heading "Different H1 Title"`,
			},
		},
		{
			name:     "RFC number does not match filename",
			filename: "0008-number-mismatch.md",
			wantErrors: []string{
				"rfc number 99 does not match filename number 8",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var content string
			if tt.filename == "_template.md" {
				// Read from the actual template in the repo root.
				data, err := os.ReadFile(filepath.Join("..", "..", "rfcs", "_template.md"))
				if err != nil {
					t.Fatalf("reading template: %v", err)
				}
				content = string(data)
			} else {
				content = readTestdata(t, tt.filename)
			}

			// Use the testdata filename for validation (so filename-based checks work).
			errs := Validate(tt.filename, content)

			if tt.wantClean {
				if len(errs) != 0 {
					t.Errorf("expected no errors, got %d:\n  %s", len(errs), strings.Join(errs, "\n  "))
				}
				return
			}

			if len(errs) == 0 {
				t.Fatal("expected errors but got none")
			}

			errStr := strings.Join(errs, "\n")
			for _, want := range tt.wantErrors {
				if !strings.Contains(errStr, want) {
					t.Errorf("expected error containing %q, got:\n  %s", want, strings.Join(errs, "\n  "))
				}
			}
		})
	}
}

func TestRFCNumberFromFilename(t *testing.T) {
	tests := []struct {
		filename string
		want     int
	}{
		{"0001-my-rfc.md", 1},
		{"0042-another-rfc.md", 42},
		{"9999-max.md", 9999},
		{"_template.md", -1},
		{"no-number.md", -1},
		{"rfcs/0001-nested.md", 1},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got := rfcNumberFromFilename(tt.filename)
			if got != tt.want {
				t.Errorf("rfcNumberFromFilename(%q) = %d, want %d", tt.filename, got, tt.want)
			}
		})
	}
}

func TestSectionHasContent(t *testing.T) {
	tests := []struct {
		name string
		body string
		want bool
	}{
		{"real content", "This is content.\n", true},
		{"empty", "", false},
		{"whitespace only", "   \n\n  \n", false},
		{"comment only", "<!-- Fill this in -->\n", false},
		{"comment and content", "<!-- Note -->\nReal content here.\n", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sectionHasContent(tt.body)
			if got != tt.want {
				t.Errorf("sectionHasContent(%q) = %v, want %v", tt.body, got, tt.want)
			}
		})
	}
}
