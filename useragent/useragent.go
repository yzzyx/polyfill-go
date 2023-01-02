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
	return ua.Family + "/" + ua.Version()
}

// Version returns a string representation of the version
func (ua Useragent) Version() string {
	return fmt.Sprintf("%d.%d.%d", ua.Major, ua.Minor, ua.Patch)
}
