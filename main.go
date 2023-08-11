package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
)

// generate 100W fake data in mysql
func main() {
	// 資料庫連線設定
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	process(2, db)
}

// count 為要執行幾次
func process(count int, db *sql.DB) {
	var wg sync.WaitGroup
	for c := 0; c < count; c++ {
		wg.Add(1)
		go func() {
			// 批次插入的大小
			batchSize := 1000
			// 要插入的記錄數量
			numRecords := 10000 * 50
			// numRecords := 1500

			// 建立一個批次切片
			var batchValues []interface{}

			for i := 0; i < numRecords; i++ {
				name := gofakeit.Name()
				address := gofakeit.Address().Address
				status := fmt.Sprintf("%d", i%10) // 將數字轉為字串

				// 將資料添加到批次切片
				batchValues = append(batchValues, name, address, status)

				// 每當達到批次大小，執行批次插入
				if len(batchValues) == batchSize || i == numRecords-1 {
					sql := "INSERT INTO users (name, address, status) VALUES "
					valueStrings := make([]string, 0, len(batchValues)/3)
					for j := 0; j < len(batchValues); j += 3 {
						valueStrings = append(valueStrings, fmt.Sprintf("('%s', '%s', '%s')", name, address, status))
					}
					sql += strings.Join(valueStrings, ", ")

					_, err := db.Exec(sql)
					if err != nil {
						log.Fatal(err)
					}
					// fmt.Println(sql)
					// 清空批次切片
					batchValues = nil
				}
			}

			fmt.Printf("插入 %d 筆資料完成\n", numRecords)
			wg.Done()
		}()
	}
	wg.Wait()
}
