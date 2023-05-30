package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	configFn string

	options struct {
		logfile  string
		loglevel string

		printVersion    bool
		showAPIsAndExit bool
		migrateDB       bool
	}
)

const (
	usage = `flexdb

Flags:	
`
)

func parseFlags() {
	flag.StringVar(&configFn, "cf", "", "config file name")

	flag.BoolVar(&options.showAPIsAndExit, "apis", false, "show RESTful APIs and exit")
	flag.BoolVar(&options.migrateDB, "migrate", false, "reset database(Dangerous!)")
	flag.BoolVar(&options.printVersion, "ver", false, "show version and exit")
	flag.StringVar(&options.logfile, "logfile", "", "master log file path, default stdout")
	flag.StringVar(&options.loglevel, "loglevel", "trace", "log level")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
	}

	flag.Parse()
}
