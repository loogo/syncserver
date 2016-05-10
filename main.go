package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/loogo/syncserver/jsonrpc"
)

func main() {
	db, err := sqlx.Open("mysql", "wcl:1@tcp(192.168.31.127:3306)/gzl")
	if err != nil {
		fmt.Println("Connect Error: ", err)
	}
	defer db.Close()

	config := loadconfig()
	jsonObj := loadmetadate()

	for _, table := range jsonObj.Tables {
		data := query(db, table, "")
		// out, _ := json.Marshal(data)
		// fmt.Println(string(out))
		params := map[string]interface{}{
			"service": "object",
			"method":  "execute",
			"args": []interface{}{
				config.DB, config.User, config.Password, "product.template", "product_import", data,
			},
		}
		url := config.URL
		result, err := jsonrpc.Call(url, params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		// for _, vals := range data {
		// 	for key, val := range vals {
		// 		fmt.Printf("%s: %v\n", key, val)
		// 	}
		// 	fmt.Println("****************************************")
		// }
	}
	// rows, err := db.Queryx("select *,opt.title as opt_title from yuanshop_shop_goods as good join yuanshop_shop_goods_option as opt on good.id = opt.goodsid")

}
