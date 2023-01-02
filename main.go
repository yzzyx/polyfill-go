package main

import (
	"fmt"
	"os"

	"github.com/yzzyx/polyfill-go/useragent"
)

func main() {
	ua := useragent.Normalise("Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko")

	pf, err := New("/home/elias/devel/polyfill-library/polyfills/__dist")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not build polyfiller: %v\n", err)
		os.Exit(1)
	}

	r, err := pf.Generate([]string{"fetch"}, ua, Options{Raw: false})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not generate polyfill: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf(r)
}
