package main

import (
	"log"
	"os"

	_ "github.com/fendijatmiko/rss-go/matchers"
	"github.com/fendijatmiko/rss-go/search"
)

// init is called prior to main
func init() {
	// change the device to logging to stdout
	log.SetOutput(os.Stdout)

}

func main() {
	// perform the search for the specific search
	search.Run("president")
}
