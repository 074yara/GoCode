package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	//timeString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//timeString = strings.Trim(timeString, "\r\n")

	timeNew, _ := time.Parse("2006-01-02 15:04:05", scan.Text())
	if timeNew.Hour() > 13 {
		fmt.Println(timeNew.Add(time.Hour * 24).Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(timeNew.Format("2006-01-02 15:04:05"))
	}

}
