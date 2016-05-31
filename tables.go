package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// JSONObject Base Type
type jsonobject struct {
	Tables tableTypes
}

// TableType Type
type TableType struct {
	Name     string
	Model    string
	Method   string
	Args     map[string]interface{}
	Seq      string
	Filter   string
	Alias    string
	Columns  mapcolumn
	Children []ChildrenType
}
type tableTypes []TableType

func (a tableTypes) Len() int           { return len(a) }
func (a tableTypes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a tableTypes) Less(i, j int) bool { return a[i].Seq < a[j].Seq }

// ChildrenType Type
type ChildrenType struct {
	TableType
	RelCol string
}

// ColumnType col type
type ColumnType struct {
	Ctype    string `json:"ctype"`
	Alias    string
	Relation string
	Relcol   string
	Select   []string
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
	sort.Sort(tables.Tables)
	// fmt.Println(tables)
	return tables
}
