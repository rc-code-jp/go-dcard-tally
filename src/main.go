package main

import (
	"fmt"
	"go-dcard-tally/src/lib"
	"go-dcard-tally/src/model"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Start...")

	// 全てのファイルを取得
	files := lib.FilesInDir("./src/assets/csv")

	// CSVファイルのみを抽出
	var csvFiles []string
	for _, file := range files {
		if filepath.Ext(file) == ".csv" {
			csvFiles = append(csvFiles, file)
		}
	}

	// 集計用の配列
	var items []model.Item

	// CSVファイルの中身を取得
	for _, csvFile := range csvFiles {
		rows := lib.CsvScan(csvFile)

		// 仕様の異なるファイルを含めるため先頭の行で分岐
		if len(rows[0]) == 1 {
			// 確定前のデータ
			// 先頭行を削除（タイトル行とヘッダー行）
			rows = append(rows[:0], rows[2:]...)

			// データを取得
			for _, row := range rows {
				if len(row) < 7 {
					continue
				}
				amount, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(row[6]))
				item := model.Item{
					Date:   row[3],
					Amount: amount,
					Place:  row[4],
				}
				items = append(items, item)
			}
		} else {
			// 確定済みのデータ
			// 先頭行を削除
			rows = append(rows[:0], rows[1:]...)

			// データを取得
			for _, row := range rows {
				// 正規表現を使用して文字列から数字を抽出する
				amount, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(row[3]))
				item := model.Item{
					Date:   row[0],
					Amount: amount,
					Place:  row[1],
				}
				items = append(items, item)
			}
		}
	}

	// 集計
	var total int
	for _, item := range items {
		total += item.Amount
	}

	fmt.Println("Total:", total)

	fmt.Println("End...")
}
