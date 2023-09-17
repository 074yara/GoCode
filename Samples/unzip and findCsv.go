package main

import (
	"archive/zip"
	"io/fs"

	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	//"strings"
)

func main() {

	zipReader, err := zip.OpenReader("task.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer zipReader.Close()
	for _, file := range zipReader.File {
		r, _ := file.Open()
		if rows, _ := csv.NewReader(r).ReadAll(); len(rows) == 10 && len(rows[4]) == 10 {
			fmt.Println(file.Name, rows[4][2])
		}
		r.Close()

	}

}

func walkerFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	file, _ := os.Open(path)
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(data) > 1 {
		fmt.Println(data[4][2])
	}
	return nil
}

func unZip(src, dest string) {
	src = "task.zip"
	dest = "task_unpacked"

	reader, err := zip.OpenReader(src)
	if err != nil {
		fmt.Println("error")
	}
	defer reader.Close()
	for _, file := range reader.File {
		fp := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			os.Mkdir(fp, os.ModePerm)
			continue
		}
		os.MkdirAll(filepath.Dir(fp), os.ModePerm)
		outFile, _ := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, file.Mode())
		inFile, _ := file.Open()

		io.Copy(outFile, inFile)
		csvReader := csv.NewReader(inFile)
		records, _ := csvReader.ReadAll()
		if len(records) > 0 {
			fmt.Println(records)
		}
		outFile.Close()
	}
}
