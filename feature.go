package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Feature describes a single package in the polyfill library (uses meta.json file)
type Feature struct {
	Aliases      []string          `json:"aliases"`
	Dependencies []string          `json:"dependencies"`
	License      string            `json:"license"`
	Browsers     map[string]string `json:"browsers"`
	DetectSource string            `json:"detectSource"`

	// Generated fields
	Name string
	Path string

	/* The following are currently not used, so we ignore them to save some memory
	Notes        []string          `json:"notes"`
	Spec         string            `json:"spec"`
	Repo         string            `json:"repo"`
	Docs         string            `json:"docs"`
	BaseDir      string            `json:"baseDir"`
	HasTests     bool              `json:"hasTests"`
	IsTestable   bool              `json:"isTestable"`
	IsPublic     bool              `json:"isPublic"`
	Size         int               `json:"size"`
	MetaInstall  struct {
		Module string   `json:"module"`
		Paths  []string `json:"paths"`
	} `json:"install"`
	*/
}

type Options struct {
	Gated bool
	Raw   bool
}

func (p *Feature) BuildPolyfill(opts Options) (string, error) {
	s := &strings.Builder{}

	filename := "min.js"
	if opts.Raw {
		filename = "raw.js"
	}

	if opts.Gated && p.DetectSource != "" {
		fmt.Fprint(s, "if (!(", p.DetectSource, ")) {\n")
	}

	contents, err := os.ReadFile(filepath.Join(p.Path, filename))
	if err != nil {
		return "", err
	}
	s.Write(contents)

	if opts.Gated && p.DetectSource != "" {
		fmt.Fprint(s, "\n}\n")
	}
	fmt.Fprint(s, "\n")
	return s.String(), nil
}
