package main

import (
	"flag"

	"github.com/alochym01/hardware-exporter/router"
)

func main() {
	// Command-line Flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	env := flag.String("env", "debug", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line
	// This reads in the command-line flag value and assigns it to the addr
	flag.Parse()

	r := router.Router(*env)

	r.Run(*addr)
}
