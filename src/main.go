package main

import (
	"fmt"
	"go-dcard-tally/src/lib"
	"go-dcard-tally/src/model"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("******集計を開始******")

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

	// 月毎に集計するためにソート
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})

	// 集計
	total := 0
	perMonth := make(map[string]int)
	for _, item := range items {
		total += item.Amount
		key := item.Date[:7]
		if _, exists := perMonth[key]; !exists {
			perMonth[key] = 0
		} else {
			perMonth[key] += item.Amount
		}
	}

	// 先頭と末尾の日付を取得
	startDate := items[0].Date
	endDate := items[len(items)-1].Date
	fmt.Println(startDate + "から" + endDate + "まで")

	// 全体
	fmt.Println("合計:", lib.FormatCurrency(total)+"円")
	fmt.Println("行数:", len(items))

	// 月毎
	for key, value := range perMonth {
		fmt.Println(key+":", lib.FormatCurrency(value)+"円")
	}

	fmt.Println("******正常終了******")
}
