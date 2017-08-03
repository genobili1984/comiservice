package main

import (
	"comiserver/internal"
	"log"
	"os"
)

type Options struct {
	Verbose     bool        `flag:"verbose"`
	LogLevel    string      `flag:"log-level"`
	LogPrefix   string      `flag:"log-prefix"`
	logLevel    lg.LogLevel // private, not really an option
	HTTPAddress string      `flag:"http-address"`
	Logger      Logger
}

func NewOptions() *Options {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hostname)
	return &Options{
		LogPrefix:   "[comiservice]",
		HTTPAddress: "0.0.0.0:8004",
	}
}
