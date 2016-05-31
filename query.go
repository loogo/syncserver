package main

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func query(db *sqlx.DB, table *TableType) []map[string]interface{} {
	cols := strings.Join(table.Columns.Keys(), ",")
	sql := fmt.Sprintf("select %s from %s", cols, table.Name)
	filter := table.Filter
	if len(filter) > 0 {
		sql += " where " + filter
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
					val := string(value.([]byte))
					col := table.Columns.getByAlias(key)

					if len(val) > 0 {
						switch col.Ctype {
						case "image":
							val = cfg.ImageRoot + val
						case "relation":
							sql = fmt.Sprintf("select %s from %s where %s=%s", strings.Join(col.Select, ","), col.Relation, col.Relcol, val)

							relVal := make(map[string]interface{})
							db.QueryRowx(sql).MapScan(relVal)
							for rkey, rvalue := range relVal {
								results[rkey] = string(rvalue.([]byte))
							}

							continue
						}
					}
					results[key] = val
					// fmt.Println(key, ": ", string(value.([]byte)))
				default:
					fmt.Printf("%s   %s    %s\n", key, value, v)
				}
			}
		}
		if table.Children != nil {
			for _, chi := range table.Children {
				where := chi.RelCol + "=" + results[table.Columns["id"].Alias].(string)
				if len(chi.Filter) > 0 {
					chi.Filter += " and " + where
				} else {
					chi.Filter = where
				}
				chidata := query(db, &chi.TableType)
				if len(chi.Alias) > 0 {
					results[chi.Alias] = chidata
				} else {
					results[chi.Name] = chidata
				}
			}
		}
		delete(results, "id")
		data = append(data, results)
	}

	return data
}
