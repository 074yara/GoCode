package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {

	var minutes, sec int64
	_, err := fmt.Scanf("%d мин. %d сек.", &minutes, &sec)
	if err != nil {
		return
	}

	fmt.Println(time.Unix(now+minutes*60+sec, 0).UTC().Format(time.UnixDate))

}
