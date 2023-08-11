package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 資料庫連線設定
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 要插入的記錄數量
	numRecords := 1000000 * 1

	for i := 0; i < numRecords; i++ {
		name := gofakeit.Name()
		address := gofakeit.Address().Address
		status := "Active"
		if rand.Intn(2) == 0 {
			status = "Inactive"
		}

		// 插入資料
		_, err := db.Exec("INSERT INTO users (name, address, status) VALUES (?, ?, ?)", name, address, status)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("插入 %d 筆資料完成\n", numRecords)
}
