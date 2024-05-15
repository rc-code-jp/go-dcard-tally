package main

import (
	"fmt"
	"go-dcard-tally/src/lib"
	"go-dcard-tally/src/model"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

			// 末尾行を削除（合計行）
			rows = append(rows[:0], rows[:len(rows)-4]...)

			// データを取得
			for _, row := range rows {
				if len(row) < 7 {
					continue
				}
				// 正規表現を使用して文字列から数字を抽出する
				numbers := regexp.MustCompile(`\d+|-`).FindAllString(row[6], -1)
				amount, _ := strconv.Atoi(strings.Join(numbers, ""))
				item := model.Item{
					Date:     row[3],
					Amount:   amount,
					Place:    row[4],
					FileType: "確定前",
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
				numbers := regexp.MustCompile(`\d+|-`).FindAllString(row[3], -1)
				amount, _ := strconv.Atoi(strings.Join(numbers, ""))
				item := model.Item{
					Date:     row[0],
					Amount:   amount,
					Place:    row[1],
					FileType: "確定済",
				}
				items = append(items, item)
			}
		}
	}

	// 集計
	var total int
	for _, item := range items {
		fmt.Println(item)
		total += item.Amount
	}

	fmt.Println("Total:", total)

	fmt.Println("Total items:", len(items))

	fmt.Println("End...")
}
