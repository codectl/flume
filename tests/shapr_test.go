// Package tests is the global check every Terraform module repository runs
package tests

import (
	"os"
	"testing"

	"github.com/codectl/shapr"
)

func TestModule(t *testing.T) {
	opts := []shapr.Option{
		shapr.WithModule(os.Getenv("MODULE_PATH")),
		shapr.WithSections("Goals", "Features", "Contributors", "Testing", "Notes", "References"),
		shapr.WithFiles("CONTRIBUTING.md", "TESTING.md", "CODE_OF_CONDUCT.md", "SECURITY.md", "GOALS.md"),
	}
	if os.Getenv("SCHEMA") != "false" {
		opts = append(opts, shapr.WithSchema())
	}

	v, err := shapr.New(opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := v.Validate(t.Context()); err != nil {
		t.Fatal(err)
	}
}
