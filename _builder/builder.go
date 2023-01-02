package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagUseragent = flag.Bool("useragent", false, "build user-agent parsing code")
	flagNormalise = flag.Bool("normalise", false, "build user-agent normalising code")
)

func main() {
	var err error
	flag.Parse()

	if *flagUseragent {
		err = useragent()
	} else if *flagNormalise {
		err = normalise()
	} else {
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not build code: %v\n", err)
		os.Exit(1)
	}
}
