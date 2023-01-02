package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type normalisation struct {
	Reason string `json:"reason"`
	Regex  string `json:"regex"`
}

type OrderedMap struct {
	Map   map[string]interface{}
	Order []string
}

func (m *OrderedMap) UnmarshalJSON(b []byte) error {
	// First, decode to map as usually
	err := json.Unmarshal(b, &m.Map)
	if err != nil {
		return err
	}

	// Then, we'll decode token-by-token and save the order
	dec := json.NewDecoder(bytes.NewReader(b))
	// First well get the delimiter '{'
	t, err := dec.Token()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		return fmt.Errorf("expected '{', got %T: %v", t, t)
	}

	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch v := t.(type) {
		case string:
			m.Order = append(m.Order, v)
		case json.Delim:
			if v == '}' {
				// This should be the last token before EOF
				continue
			}
		}

		t, err = dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if delim, ok := t.(json.Delim); ok {
			if delim == '{' || delim == '[' {
				err = decodeObject(dec, delim)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("unexpected delimiter %v at this position", delim)
			}
		}
	}

	return nil
}

// decodeObject decodes a json object/array and ignores the result
func decodeObject(dec *json.Decoder, startDelim json.Delim) error {
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if delim, ok := t.(json.Delim); ok {
			if startDelim == '{' && delim == '}' {
				return nil
			} else if startDelim == '[' && delim == ']' {
				return nil
			}
			err = decodeObject(dec, delim)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type normalisationData struct {
	IsNormalized     string            `json:"isNormalized"`
	BaselineVersions map[string]string `json:"baselineVersions"`
	Normalisations   []normalisation   `json:"normalisations"`
	Aliases          OrderedMap        `json:"aliases"`
}

//go:embed polyfill-useragent-normaliser/data.json
var normalisationJSON []byte

func normalise() error {
	nd := normalisationData{}
	err := json.Unmarshal(normalisationJSON, &nd)
	if err != nil {
		return err
	}

	checks := &strings.Builder{}
	globals := &strings.Builder{}

	fmt.Fprintf(globals, "var nm_re_isNormalized = regexp.MustCompile(`%s`)\n", strings.Trim(nd.IsNormalized, "/"))
	fmt.Fprintf(checks, `if match := nm_re_isNormalized.FindStringSubmatch(ua); match != nil {
		major, _ := strconv.Atoi(match[2])
		minor, _ := strconv.Atoi(match[3])
		return Useragent{
			Family: match[0],
			Major:  major,
			Minor:  minor,
			Patch:  0,
		}
}
`)

	for idx, normalisation := range nd.Normalisations {
		fmt.Fprintf(globals, "var nm_re_%d = regexp.MustCompile(`%s`)\n", idx, strings.Trim(normalisation.Regex, "/"))
		fmt.Fprintf(checks, "if match := nm_re_%d.FindStringSubmatchIndex(ua); match != nil {\n"+
			"	ua = ua[0:match[0]] + ua[match[1]:]\n"+
			"}\n", idx)
	}

	fmt.Fprintf(checks, "client := Parse(ua)\n")

	fmt.Fprintf(checks, `
	// For improved CDN cache performance, remove the patch version.
	//  There are few cases in which a patch release drops the requirement for a polyfill, but if so, the polyfill
	//  can simply be served unnecessarily to the patch versions that contain the fix, and we can stop targeting
	//  at the next minor release.
	client.Patch = 0
	client.Family = strings.ToLower(client.Family)
`)

	fmt.Fprintf(checks, "// Apply aliases to useragent family and version numbers\n")

	for _, key := range nd.Aliases.Order {
		value := nd.Aliases.Map[key]

		fmt.Fprintf(checks, "if client.Family == `%s` {", key)
		switch v := value.(type) {
		case string:
			fmt.Fprintf(checks, "client.Family = `%s`\n", v)
		case map[string]interface{}:
			for version, value := range v {
				// Expect value to be an array with two entries - family and major version
				arr, ok := value.([]interface{})
				if !ok {
					continue
				}

				var major, minor int
				parts := strings.Split(version, ".")
				major, _ = strconv.Atoi(parts[0])

				if len(parts) > 1 {
					minor, _ = strconv.Atoi(parts[1])
					fmt.Fprintf(checks, "if client.Major == %d && client.Minor == %d {\n", major, minor)
				} else {
					fmt.Fprintf(checks, "if client.Major == %d {\n", major)
				}

				fmt.Fprintf(checks, "client.Family = `%s`\n", arr[0].(string))
				fmt.Fprintf(checks, "client.Major = %d\n", int(arr[1].(float64)))
				fmt.Fprintf(checks, "client.Minor = 0\n")
				fmt.Fprintf(checks, "}\n")
			}
		case []interface{}:
			fmt.Fprintf(checks, "client.Family = `%s`\n", v[0].(string))
			fmt.Fprintf(checks, "client.Major = %d\n", int(v[1].(float64)))
			fmt.Fprintf(checks, "client.Minor = 0\n")
		}
		fmt.Fprintf(checks, "}\n")
	}

	fmt.Fprintf(checks, "// Compare result against list of polyfill baseline\n")
	// 4. Check if browser and version are in the baseline supported browser versions
	baselines := []string{}
	for family, version := range nd.BaselineVersions {
		bl := fmt.Sprintf("(client.Family == `%s`", family)

		if version == "*" {
			bl += ")"
			baselines = append(baselines, bl)
			continue
		}

		parts := strings.Split(version, ".")
		var major, minor int
		major, _ = strconv.Atoi(parts[0])
		if len(parts) > 1 {
			minor, _ = strconv.Atoi(parts[1])
			bl += fmt.Sprintf(" && client.Major > %d || (client.Major == %d && client.Minor >= %d))", major, major, minor)
			baselines = append(baselines, bl)
		} else {
			bl += fmt.Sprintf(" && client.Major >= %d )", major)
			baselines = append(baselines, bl)
		}
	}
	fmt.Fprintf(checks, "if %s { /* empty */ } else {\n", strings.Join(baselines, " ||\n"))
	fmt.Fprintf(checks, " return Useragent{Family: `other`}\n")
	fmt.Fprintf(checks, "}\n")

	fmt.Printf(`// Code generated by "builder -normalise"; DO NOT EDIT.
package useragent

import (
"regexp"
"strconv"
"strings"
)
`)
	fmt.Printf(globals.String())
	fmt.Printf(`
// Normalise parses a user-agent string and normalises it according to the same rules as https://github.com/Financial-Times/polyfill-useragent-normaliser
func Normalise(ua string) Useragent {
%s
	return client
}`, checks)
	return nil
}
