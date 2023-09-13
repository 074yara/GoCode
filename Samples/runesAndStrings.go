package main

import "fmt"

func main() {
	var txt string = "Слово"
	txtRunes := []rune(txt)
	fmt.Printf("\nСрез рун: %v", txtRunes)
	fmt.Printf("\nИз рун получаем снова текст: %v", string(txtRunes))
	lenTxtRunes := len(txtRunes)
	fmt.Printf("\nДлина текста (среза): %v", lenTxtRunes)

	for i, v := range txtRunes {
		fmt.Printf("\nСимвол %d: %s", i, string(v))
	}

	txtRunes[lenTxtRunes-1] = rune('а')
	fmt.Printf("\nЗаменили последнюю букву на \"а\": %v", string(txtRunes))

	txtRunes = append(txtRunes[:4])
	fmt.Printf("\nОтрезали последнюю букву: %v", string(txtRunes))

	txtRunes = append(txtRunes, 'а', 'р', 'ь')
	fmt.Printf("\nДобавили три буквы в конец: %v", string(txtRunes))
}
