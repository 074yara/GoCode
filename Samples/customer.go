package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

// Создаём тип Customer.
type Customer struct {
	Name string
	Age  int
}

// Реализуем метод WriteJSON, который берёт io.Writer в виде параметра.
// Он отправляет структуру Сustomer в JSON, и если всё отрабатывает
// успешно, то вызывается соответствующий метод Write() из io.Writer.
func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)
	return err
}

func main3() {
	// Инициализируем структуру Customer.
	c := &Customer{Name: "Alice", Age: 21}

	// Затем с помощью Buffer можем вызвать метод WriteJSON
	var buf bytes.Buffer
	err := c.WriteJSON(buf)
	if err != nil {
		log.Fatal(err)
	}

	// или воспользоваться файлом.
	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
