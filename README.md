# カード明細CSVの集計

## init

```bash
make init
```

## file

1. csvファイルをsrc/assets/csvに置く（ファイル名は自由、csvの形式は変更不可）
  - ご利用内訳明細_キャッシングご返済明細_YYYYMMDD.csv
  - ご利用明細照会_YYYY年M月お支払い分
1. runを実行する

## run

```bash
make app
go build src/main.go && go run src/main.go 
```
