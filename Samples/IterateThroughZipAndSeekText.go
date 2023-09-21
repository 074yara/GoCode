package main

import (
	"archive/zip"
	"fmt"
	"os"
	"strings"
)

func main() {

	zipReader, err := zip.OpenReader("Shakespeare.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer func(zipReader *zip.ReadCloser) {
		err := zipReader.Close()
		if err != nil {

		}
	}(zipReader)
	for _, file := range zipReader.File {
		_, err := file.Open()
		if err != nil {
			return
		}
		f, _ := os.ReadFile(file.Name)
		rawTextSlice := strings.Split(string(f), "\n")
		for _, v := range rawTextSlice {
			if strings.Contains(v, "love") {
				newFile, err := os.OpenFile("newFile.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
				_, err = newFile.WriteString(v + "\n")
				if err != nil {
					fmt.Println(err)
				}
				err = newFile.Close()
				if err != nil {
					return
				}
			}
		}

	}

}
