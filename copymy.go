package main

import (
	"flag"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/pelletier/go-toml"
)

func main() {
	flag.Parse()
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		key := flag.Arg(0)
		value := config.Get(key).(string)
		if err := clipboard.WriteAll(value); err != nil {
			panic(err)
		} else {
			fmt.Printf("Copied %s to keyboard\n", value)
		}
	}
}
