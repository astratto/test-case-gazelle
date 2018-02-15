package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// TestConf example conf
type TestConf struct {
	Count int `toml:"count,omitempty"`
}

func main() {
	var testSimple = "count = 100"

	var conf TestConf
	if _, err := toml.Decode(testSimple, &conf); err != nil {
		log.Fatal(err)
	}

	log.Println("Count: ", conf.Count)
}
