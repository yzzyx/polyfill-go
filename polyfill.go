package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/yzzyx/polyfill-go/useragent"
)

// Polyfiller creates packages of polyfills based on the supplied user-agent
type Polyfiller struct {
	dataDir    fs.ReadDirFS
	aliases    map[string][]string
	featureMap map[string]*Feature
}

// New creates a new Polyfiller instance that reads polyfills from the supplied dataPath
func New(dataPath string) (*Polyfiller, error) {
	aliases := map[string][]string{}
	featureMap := map[string]*Feature{}

	dir, err := os.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range dir {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		pkgName := entry.Name()

		p := filepath.Join(dataPath, pkgName, "meta.json")
		f, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("could not read meta-file for package %s: %w", pkgName, err)
		}

		meta := Feature{}
		err = json.Unmarshal(f, &meta)
		if err != nil {
			return nil, fmt.Errorf("could not parse meta-file for package %s: %w", pkgName, err)
		}

		meta.Name = pkgName
		meta.Path = filepath.Join(dataPath, pkgName)
		featureMap[pkgName] = &meta
		for _, a := range meta.Aliases {
			aliases[a] = append(aliases[a], pkgName)
		}
	}

	return &Polyfiller{
		aliases:    aliases,
		featureMap: featureMap,
	}, nil
}

func (p *Polyfiller) Generate(features []string, useragent useragent.Useragent, opts Options) (string, error) {
	uaVersion, err := semver.NewVersion(useragent.Version())
	if err != nil {
		uaVersion, _ = semver.NewVersion("0.0.0")
	}

	var featureList []string
	for _, featureName := range features {
		// First, check if it's a package or an alias
		if l, ok := p.aliases[featureName]; ok {
			featureList = append(featureList, l...)
			continue
		}

		// Ignore features we cannot find
		if _, ok := p.featureMap[featureName]; !ok {
			continue
		}
		featureList = append(featureList, featureName)

	}

	var dependencyLists [][]string
	visited := map[string]bool{}

	var idx int
	for {
		lastIdx := len(featureList)
		var added []string
		for ; idx < lastIdx; idx++ {
			pkgName := featureList[idx]

			// Do not add duplicates
			if visited[pkgName] {
				continue
			}
			visited[pkgName] = true

			// Check if this feature is applicable for the current browser
			constraintStr, ok := p.featureMap[pkgName].Browsers[useragent.Family]
			if !ok {
				continue
			}

			comp, err := semver.NewConstraint(constraintStr)
			if err != nil {
				return "", fmt.Errorf("could not compile constraint '%s' (package %s, family %s): %v\n", constraintStr, pkgName, useragent.Family, err)
			}

			if !comp.Check(uaVersion) {
				continue
			}

			added = append(added, pkgName)

			// Iterate through all dependencies also
			for _, dep := range p.featureMap[pkgName].Dependencies {
				featureList = append(featureList, dep)
			}
		}

		// No new entries added to list
		if len(featureList) == lastIdx {
			break
		}
		depSorter(added).Sort()
		dependencyLists = append(dependencyLists, added)
	}

	// Reverse our dependency list, so that we execute the entries most other features
	// depend upon first.
	featureList = []string{}
	for idx := len(dependencyLists) - 1; idx >= 0; idx-- {
		featureList = append(featureList, dependencyLists[idx]...)
	}

	buf := &strings.Builder{}
	contents := &strings.Builder{}
	fmt.Fprintln(buf, "/* polyfill-service")
	// Check which packages apply to this user-agent
	for _, pkgName := range featureList {
		f := p.featureMap[pkgName]
		if opts.Raw {
			license := f.License
			if license == "" {
				license = "CC0"
			}
			_, _ = fmt.Fprintf(buf, " * - %s, License: %s\n", pkgName, license)
		}

		polyfill, err := f.BuildPolyfill(opts)
		if err != nil {
			return "", err
		}
		fmt.Fprint(contents, polyfill)
	}
	fmt.Fprint(buf, " */\n")

	if len(featureList) > 0 {
		fmt.Fprint(buf, "(function(self, undefined) {\n")
		fmt.Fprint(buf, contents.String())
		fmt.Fprint(buf, "})('object' === typeof window && window || 'object' === typeof self && self || 'object' === typeof global && global || {});\n")
	} else if opts.Raw {
		fmt.Fprint(buf, "\n/* No polyfills needed for current settings and browser */\n")
	}
	return buf.String(), nil
}

// depSorter is used to sort dependencies, where we want all deps starting with '_' first.
type depSorter []string

func (x depSorter) Len() int { return len(x) }
func (x depSorter) Less(i, j int) bool {
	ip := strings.HasPrefix(x[i], "_")
	jp := strings.HasPrefix(x[j], "_")
	if ip && !jp {
		return true
	}
	if !ip && jp {
		return false
	}
	return x[i] < x[j]
}
func (x depSorter) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x depSorter) Sort()         { sort.Sort(x) }
