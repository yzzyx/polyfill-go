// Code generated by "builder -normalise"; DO NOT EDIT.
package useragent

import (
	"regexp"
	"strconv"
	"strings"
)

var nm_re_isNormalized = regexp.MustCompile(`(?i)^(\w+)\/(\d+)(?:\.(\d+)(?:\.(\d+))?)?$`)
var nm_re_0 = regexp.MustCompile(`(?i) GSA\/[\d.]+`)
var nm_re_1 = regexp.MustCompile(`(?i) Instagram [\d.]+`)
var nm_re_2 = regexp.MustCompile(`(?i) PTST\/[\d.]+`)
var nm_re_3 = regexp.MustCompile(`(?i) Waterfox\/[\d.]+`)
var nm_re_4 = regexp.MustCompile(`(?i) Goanna\/[\d.]+`)
var nm_re_5 = regexp.MustCompile(`(?i) PaleMoon\/[\d.]+`)
var nm_re_6 = regexp.MustCompile(`(YaBrowser)\/(\d+\.)+\d+ `)
var nm_re_7 = regexp.MustCompile(`(?i) (Crosswalk)\/(\d+)\.(\d+)\.(\d+)\.(\d+)`)
var nm_re_8 = regexp.MustCompile(`((CriOS|OPiOS)\/(\d+)\.(\d+)\.(\d+)\.(\d+)|(FxiOS\/(\d+)\.(\d+)))`)
var nm_re_9 = regexp.MustCompile(`(?i) vivaldi\/[\d.]+\d+`)
var nm_re_10 = regexp.MustCompile(`(?i) \[(FB_IAB|FBAN|FBIOS|FB4A)\/[^\]]+\]`)
var nm_re_11 = regexp.MustCompile(`(?i) Electron\/[\d.]+\d+`)
var nm_re_12 = regexp.MustCompile(`(?i) Edg\/[\d.]+\d+`)
var nm_re_13 = regexp.MustCompile(`(?i)Safari.* Googlebot\/2\.1; \+http:\/\/www\.google\.com\/bot\.html\)`)
var nm_re_14 = regexp.MustCompile(`(?i) Googlebot\/2\.1; \+http:\/\/www\.google\.com\/bot\.html\) `)

// Normalise parses a user-agent string and normalises it according to the same rules as https://github.com/Financial-Times/polyfill-useragent-normaliser
func Normalise(ua string) Useragent {
	if match := nm_re_isNormalized.FindStringSubmatch(ua); match != nil {
		major, _ := strconv.Atoi(match[2])
		minor, _ := strconv.Atoi(match[3])
		return Useragent{
			Family: match[0],
			Major:  major,
			Minor:  minor,
			Patch:  0,
		}
	}
	if match := nm_re_0.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_1.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_2.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_3.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_4.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_5.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_6.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_7.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_8.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_9.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_10.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_11.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_12.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_13.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	if match := nm_re_14.FindStringSubmatchIndex(ua); match != nil {
		ua = ua[0:match[0]] + ua[match[1]:]
	}
	client := Parse(ua)

	// For improved CDN cache performance, remove the patch version.
	//  There are few cases in which a patch release drops the requirement for a polyfill, but if so, the polyfill
	//  can simply be served unnecessarily to the patch versions that contain the fix, and we can stop targeting
	//  at the next minor release.
	client.Patch = 0
	client.Family = strings.ToLower(client.Family)
	// Apply aliases to useragent family and version numbers
	if client.Family == `blackberry webkit` {
		client.Family = `bb`
	}
	if client.Family == `blackberry` {
		client.Family = `bb`
	}
	if client.Family == `pale moon (firefox variant)` {
		client.Family = `firefox`
	}
	if client.Family == `pale moon` {
		client.Family = `firefox`
	}
	if client.Family == `firefox mobile` {
		client.Family = `firefox_mob`
	}
	if client.Family == `firefox namoroka` {
		client.Family = `firefox`
	}
	if client.Family == `firefox shiretoko` {
		client.Family = `firefox`
	}
	if client.Family == `firefox minefield` {
		client.Family = `firefox`
	}
	if client.Family == `firefox alpha` {
		client.Family = `firefox`
	}
	if client.Family == `firefox beta` {
		client.Family = `firefox`
	}
	if client.Family == `microb` {
		client.Family = `firefox`
	}
	if client.Family == `mozilladeveloperpreview` {
		client.Family = `firefox`
	}
	if client.Family == `iceweasel` {
		client.Family = `firefox`
	}
	if client.Family == `opera tablet` {
		client.Family = `opera`
	}
	if client.Family == `opera mobile` {
		client.Family = `op_mob`
	}
	if client.Family == `opera mini` {
		client.Family = `op_mini`
	}
	if client.Family == `chrome mobile webview` {
		client.Family = `chrome`
	}
	if client.Family == `chrome mobile` {
		client.Family = `chrome`
	}
	if client.Family == `chrome frame` {
		client.Family = `chrome`
	}
	if client.Family == `chromium` {
		client.Family = `chrome`
	}
	if client.Family == `headlesschrome` {
		client.Family = `chrome`
	}
	if client.Family == `ie mobile` {
		client.Family = `ie_mob`
	}
	if client.Family == `ie large screen` {
		client.Family = `ie`
	}
	if client.Family == `internet explorer` {
		client.Family = `ie`
	}
	if client.Family == `edge mobile` {
		client.Family = `edge_mob`
	}
	if client.Family == `uc browser` {
		if client.Major == 9 && client.Minor == 9 {
			client.Family = `ie`
			client.Major = 10
			client.Minor = 0
		}
	}
	if client.Family == `chrome mobile ios` {
		client.Family = `ios_chr`
	}
	if client.Family == `mobile safari` {
		client.Family = `ios_saf`
	}
	if client.Family == `iphone` {
		client.Family = `ios_saf`
	}
	if client.Family == `iphone simulator` {
		client.Family = `ios_saf`
	}
	if client.Family == `mobile safari uiwebview` {
		client.Family = `ios_saf`
	}
	if client.Family == `mobile safari ui/wkwebview` {
		client.Family = `ios_saf`
	}
	if client.Family == `mobile safari/wkwebview` {
		client.Family = `ios_saf`
	}
	if client.Family == `samsung internet` {
		client.Family = `samsung_mob`
	}
	if client.Family == `phantomjs` {
		client.Family = `safari`
		client.Major = 5
		client.Minor = 0
	}
	if client.Family == `opera` {
		if client.Major == 20 {
			client.Family = `chrome`
			client.Major = 33
			client.Minor = 0
		}
		if client.Major == 24 {
			client.Family = `chrome`
			client.Major = 37
			client.Minor = 0
		}
		if client.Major == 39 {
			client.Family = `chrome`
			client.Major = 52
			client.Minor = 0
		}
		if client.Major == 42 {
			client.Family = `chrome`
			client.Major = 55
			client.Minor = 0
		}
		if client.Major == 57 {
			client.Family = `chrome`
			client.Major = 70
			client.Minor = 0
		}
		if client.Major == 62 {
			client.Family = `chrome`
			client.Major = 75
			client.Minor = 0
		}
		if client.Major == 65 {
			client.Family = `chrome`
			client.Major = 78
			client.Minor = 0
		}
		if client.Major == 28 {
			client.Family = `chrome`
			client.Major = 41
			client.Minor = 0
		}
		if client.Major == 29 {
			client.Family = `chrome`
			client.Major = 42
			client.Minor = 0
		}
		if client.Major == 51 {
			client.Family = `chrome`
			client.Major = 64
			client.Minor = 0
		}
		if client.Major == 53 {
			client.Family = `chrome`
			client.Major = 66
			client.Minor = 0
		}
		if client.Major == 66 {
			client.Family = `chrome`
			client.Major = 79
			client.Minor = 0
		}
		if client.Major == 44 {
			client.Family = `chrome`
			client.Major = 57
			client.Minor = 0
		}
		if client.Major == 49 {
			client.Family = `chrome`
			client.Major = 62
			client.Minor = 0
		}
		if client.Major == 64 {
			client.Family = `chrome`
			client.Major = 77
			client.Minor = 0
		}
		if client.Major == 21 {
			client.Family = `chrome`
			client.Major = 34
			client.Minor = 0
		}
		if client.Major == 33 {
			client.Family = `chrome`
			client.Major = 46
			client.Minor = 0
		}
		if client.Major == 34 {
			client.Family = `chrome`
			client.Major = 47
			client.Minor = 0
		}
		if client.Major == 35 {
			client.Family = `chrome`
			client.Major = 48
			client.Minor = 0
		}
		if client.Major == 41 {
			client.Family = `chrome`
			client.Major = 54
			client.Minor = 0
		}
		if client.Major == 45 {
			client.Family = `chrome`
			client.Major = 58
			client.Minor = 0
		}
		if client.Major == 54 {
			client.Family = `chrome`
			client.Major = 67
			client.Minor = 0
		}
		if client.Major == 26 {
			client.Family = `chrome`
			client.Major = 39
			client.Minor = 0
		}
		if client.Major == 30 {
			client.Family = `chrome`
			client.Major = 43
			client.Minor = 0
		}
		if client.Major == 32 {
			client.Family = `chrome`
			client.Major = 45
			client.Minor = 0
		}
		if client.Major == 40 {
			client.Family = `chrome`
			client.Major = 53
			client.Minor = 0
		}
		if client.Major == 43 {
			client.Family = `chrome`
			client.Major = 56
			client.Minor = 0
		}
		if client.Major == 47 {
			client.Family = `chrome`
			client.Major = 60
			client.Minor = 0
		}
		if client.Major == 25 {
			client.Family = `chrome`
			client.Major = 38
			client.Minor = 0
		}
		if client.Major == 31 {
			client.Family = `chrome`
			client.Major = 44
			client.Minor = 0
		}
		if client.Major == 48 {
			client.Family = `chrome`
			client.Major = 61
			client.Minor = 0
		}
		if client.Major == 52 {
			client.Family = `chrome`
			client.Major = 65
			client.Minor = 0
		}
		if client.Major == 56 {
			client.Family = `chrome`
			client.Major = 69
			client.Minor = 0
		}
		if client.Major == 22 {
			client.Family = `chrome`
			client.Major = 35
			client.Minor = 0
		}
		if client.Major == 23 {
			client.Family = `chrome`
			client.Major = 36
			client.Minor = 0
		}
		if client.Major == 36 {
			client.Family = `chrome`
			client.Major = 49
			client.Minor = 0
		}
		if client.Major == 37 {
			client.Family = `chrome`
			client.Major = 50
			client.Minor = 0
		}
		if client.Major == 46 {
			client.Family = `chrome`
			client.Major = 59
			client.Minor = 0
		}
		if client.Major == 50 {
			client.Family = `chrome`
			client.Major = 63
			client.Minor = 0
		}
		if client.Major == 55 {
			client.Family = `chrome`
			client.Major = 68
			client.Minor = 0
		}
		if client.Major == 60 {
			client.Family = `chrome`
			client.Major = 73
			client.Minor = 0
		}
		if client.Major == 61 {
			client.Family = `chrome`
			client.Major = 74
			client.Minor = 0
		}
		if client.Major == 27 {
			client.Family = `chrome`
			client.Major = 40
			client.Minor = 0
		}
		if client.Major == 38 {
			client.Family = `chrome`
			client.Major = 51
			client.Minor = 0
		}
		if client.Major == 58 {
			client.Family = `chrome`
			client.Major = 71
			client.Minor = 0
		}
		if client.Major == 59 {
			client.Family = `chrome`
			client.Major = 72
			client.Minor = 0
		}
		if client.Major == 63 {
			client.Family = `chrome`
			client.Major = 76
			client.Minor = 0
		}
		if client.Major == 67 {
			client.Family = `chrome`
			client.Major = 80
			client.Minor = 0
		}
	}
	if client.Family == `googlebot` {
		if client.Major == 2 && client.Minor == 1 {
			client.Family = `chrome`
			client.Major = 41
			client.Minor = 0
		}
	}
	// Compare result against list of polyfill baseline
	if (client.Family == `safari` && client.Major >= 9) ||
		(client.Family == `samsung_mob` && client.Major >= 4) ||
		(client.Family == `firefox` && client.Major >= 38) ||
		(client.Family == `firefox_mob` && client.Major >= 38) ||
		(client.Family == `opera` && client.Major >= 33) ||
		(client.Family == `op_mob` && client.Major >= 10) ||
		(client.Family == `edge_mob`) ||
		(client.Family == `chrome` && client.Major >= 29) ||
		(client.Family == `ios_saf` && client.Major >= 9) ||
		(client.Family == `ios_chr` && client.Major >= 9) ||
		(client.Family == `op_mini` && client.Major >= 5) ||
		(client.Family == `bb` && client.Major >= 6) ||
		(client.Family == `edge`) ||
		(client.Family == `ie` && client.Major >= 9) ||
		(client.Family == `ie_mob` && client.Major >= 11) ||
		(client.Family == `android` && client.Major > 4 || (client.Major == 4 && client.Minor >= 3)) { /* empty */
	} else {
		return Useragent{Family: `other`}
	}

	return client
}
