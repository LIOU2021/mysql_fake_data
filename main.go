package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit"
	_ "github.com/go-sql-driver/mysql"
)

// 批次插入的大小
var batchSize = 1000

// 要插入的記錄數量
var numRecords = 10000 * 100

func main() {
	// 資料庫連線設定
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	process(db)
}

// count 為要執行幾次
func process(db *sql.DB) {

	queryString := ""
	for i := 0; i < numRecords; i++ {
		name := gofakeit.Name()
		address := gofakeit.Address().Address
		status := fmt.Sprintf("%d", i%10) // 將數字轉為字串

		if i == (numRecords - 1) {
			queryString += fmt.Sprintf("('%s', '%s', '%s')", name, address, status)
			break
		}
		queryString += fmt.Sprintf("('%s', '%s', '%s'), ", name, address, status)

		if i != 0 && (i+1)%batchSize == 0 {
			queryString = queryString[0 : len(queryString)-2]
			var sql = "INSERT INTO users (name, address, status) VALUES " + queryString
			_, err := db.Exec(sql)
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(sql)
			queryString = ""
		}
	}
	fmt.Printf("插入 %d 筆資料完成\n", numRecords)
}
