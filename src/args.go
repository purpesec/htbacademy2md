package main

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	moduleUrl   string
	cookies     string
	lang        string
	localImages bool
}

func getArguments() Args {
	var mFlag = flag.String("m", "", "(REQUIRED) Academy Module URL to the first page.")
	var cFlag = flag.String("c", "", "(REQUIRED) Academy Cookies for authorization.")
	var langFlag = flag.String("lang", "", "Language code for translated content (e.g. es, fr, de, ja). Empty = original English content.")
	var imgFlag = flag.Bool("local_images", false, "Save images locally rather than referencing the URL location.")
	flag.Parse()
	arg := Args{
		moduleUrl:   *mFlag,
		cookies:     *cFlag,
		lang:        *langFlag,
		localImages: *imgFlag,
	}

	if arg.moduleUrl == "" || arg.cookies == "" {
		fmt.Println("Missing required arguments for module URL and HTB Academy Cookies. Please use the -h option for help.")
		os.Exit(1)
	}

	return arg
}
