package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kshvmdn/gfp"
)

var options struct {
	version bool
	log     bool
	workers int
}

func main() {
	flag.Usage = func() {
		fmt.Printf("usage: %s [options] source target\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&options.log, "show-log", false, "show log output")
	flag.BoolVar(&options.version, "version", false, "print version and exit")
	flag.IntVar(&options.workers, "workers", 6, "number of workers")
	flag.Parse()

	if options.version {
		fmt.Printf("gfp v%v\n", gfp.Version)
		os.Exit(0)
	}

	if !options.log {
		devNull, err := os.Open(os.DevNull)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(devNull)
	}

	accessToken := os.Getenv(gfp.TokenName)
	if accessToken == "" {
		fmt.Printf("Expected access token to be exported as %s.\n", gfp.TokenName)
		os.Exit(1)
	}

	// TODO: Check that credentials are valid.
	client := gfp.GetClient(accessToken)

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	origin, target := flag.Args()[0], flag.Args()[1]
	if origin == target {
		fmt.Println("Expected unique usernames.")
		os.Exit(1)
	}

	user := gfp.Run(origin, target, options.workers, client)
	fmt.Println(user.String())
}
