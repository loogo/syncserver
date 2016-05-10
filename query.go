package main

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func query(db *sqlx.DB, table TableType, where string) []map[string]interface{} {
	cols := strings.Join(table.Columns.Keys(), ",")
	sql := fmt.Sprintf("select %s from %s", cols, table.Name)
	if len(where) > 0 {
		sql += " where " + where
	}
	fmt.Println(sql)
	rows, err := db.Queryx(sql)
	if err != nil {
		fmt.Println("Sql Error: ", err)
	}
	var data []map[string]interface{}
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		if err != nil {
			fmt.Println("Scan Error: ", err)
		}

		for key, value := range results {
			if value != nil {
				switch v := value.(type) {
				case []byte:
					results[key] = string(value.([]byte))
					// fmt.Println(key, ": ", string(value.([]byte)))
				default:
					fmt.Println("unknown", v)
				}
			}
		}
		if table.Children != nil {
			for _, chi := range table.Children {
				chidata := query(db, chi.TableType, chi.RelCol+"="+results["id"].(string))
				results[chi.Name] = chidata
			}
		}
		data = append(data, results)
	}

	return data
}
