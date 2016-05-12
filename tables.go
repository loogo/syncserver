package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JSONObject Base Type
type jsonobject struct {
	Tables []TableType
}

// TableType Type
type TableType struct {
	Name     string
	Filter   string
	Alias    string
	Columns  mapcolumn
	Children []ChildrenType
}

// ChildrenType Type
type ChildrenType struct {
	TableType
	RelCol string
}

// ColumnType col type
type ColumnType struct {
	Ctype string `json:"ctype"`
	Alias string
}

// LoadConfig load json config
func loadmetadate() (tables jsonobject) {
	file, err := ioutil.ReadFile("./config/data.json")
	if err != nil {
		fmt.Printf("file error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	err = json.Unmarshal(file, &tables)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(tables)
	return tables
}
