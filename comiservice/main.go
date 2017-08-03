package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/judwhite/go-svc/svc"
	options "github.com/mreiferson/go-options"
)

type program struct {
	comiservice *ComiService
}

func comiFlagSet(opts *Options) *flag.FlagSet {
	flagSet := flag.NewFlagSet("comiservice", flag.ExitOnError)
	flagSet.String("config", "", "path to config file")
	flagSet.Bool("version", false, "print version string")
	flagSet.String("log-level", "info", "set log verbosity: debug, info, warn, error, or fatal")
	flagSet.String("log-prefix", "[comiservice] ", "log message prefix")
	flagSet.Bool("verbose", false, "deprecated in favor of log-level")
	flagSet.String("http-address", opts.HTTPAddress, "<addr>:<port> to listen on for HTTP clients")
	return flagSet
}

func main() {
	runtime.GOMAXPROCS(2)
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *program) Start() error {
	opts := NewOptions()
	flagSet := comiFlagSet(opts)
	flagSet.Parse(os.Args[1:])

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println("nsqlookupd")
		os.Exit(0)
	}

	var cfg map[string]interface{}
	configFile := flagSet.Lookup("config").Value.String()
	if configFile != "" {
		_, err := toml.DecodeFile(configFile, &cfg)
		if err != nil {
			log.Fatalf("ERROR: failed to load config file %s - %s", configFile, err.Error())
		}
	}

	options.Resolve(opts, flagSet, cfg)
	service := New(opts)

	service.Main()
	p.comiservice = service
	return nil
}

func (p *program) Stop() error {
	if p.comiservice != nil {
		p.comiservice.Exit()
	}
	return nil
}
