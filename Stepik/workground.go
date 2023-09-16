package main

import (
	"archive/zip"

	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	//"strings"
)

func main() {
	//src := "task.zip"
	dest := "task_unpacked"
	//unZip(src, dest)
	if err := filepath.Walk(dest, walkerFunc); err != nil {

	}

}

func walkerFunc(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	file, _ := os.Open(filepath.Join())
	//fmt.Println(file)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil

	}
	fmt.Println(records)
	file.Close()

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
