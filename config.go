package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	DB, URL, Password, ImageRoot, DataSourceName string
	User                                         int
}

// LoadConfig load json config
func loadconfig() (config config) {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Printf("file error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(tables)
	return config
}
