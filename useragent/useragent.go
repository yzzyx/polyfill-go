package useragent

import "fmt"

// Useragent describes a browser useragent
type Useragent struct {
	Family string
	Major  int
	Minor  int
	Patch  int
}

// String returns a string representation of a user agent
func (ua Useragent) String() string {
	return fmt.Sprintf("%s/%d.%d.%d", ua.Family, ua.Major, ua.Minor, ua.Patch)
}
