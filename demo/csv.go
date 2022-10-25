package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVFile struct {
	filename string
	fp       *os.File
}

func Open(filename string) (*CSVFile, error) {
	fp, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	fp.WriteString("\xEF\xBB\xBF") ////写入UTF-8 BOM,此处如果不写入就会导致写入的汉字乱码 不写好像也没问题
	return &CSVFile{filename: filename, fp: fp}, nil
}
func (f *CSVFile) SaveCsv(data [][]string) error {
	w := csv.NewWriter(f.fp)
	err := w.WriteAll(data) // calls Flush internally
	return err
}
func (f *CSVFile) Close() {
	f.fp.Close()
	f.filename = ""
}

func main() {
	file, err := Open("test.csv")
	if err != nil {
		fmt.Printf("Open error: err=%+v\n", err)
		return
	}
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "陈", "rob"},
		{"Ken", "中李老师街坊邻居阿拉法", "ken"},
		{"Robert", "将上课附件阿九", "gri"},
	}
	err = file.SaveCsv(records)
	err = file.SaveCsv(records)
	if err != nil {
		fmt.Printf("SaveCsv error: err=%+v\n", err)
		return
	}
	file.Close()

}
