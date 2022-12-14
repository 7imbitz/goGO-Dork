package main

import (
	"github.com/7imbitz/goGO-Dork/cmd/pkg/args"
	"github.com/7imbitz/goGO-Dork/cmd/pkg/core"

	"os"
	"os/signal"

	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
)

var options = &args.Options{}

func main() {
	gracefulExit()
	core.ShowBanner()
	readArgs()
	core.ParseOptions(options)
}

func readArgs() {
	set := goflags.NewFlagSet()
	set.SetDescription("Simple Google Dork Search")
	set.StringVarP(&options.Domain, "domain", "d", "", "Domain to scan")
	set.IntVarP(&options.Results, "result", "r", 10, "Number of results per search")

	if err := set.Parse(); err != nil {
		gologger.Fatal().Msgf("Could not parse flags: %s\n", err)
	}

	_ = set.Parse()
}

func gracefulExit() {
	//setting up graceful exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			gologger.Fatal().Msgf("CTRL+C pressed: Exiting\n")
			os.Exit(1)
		}
	}()
}
