package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
	// type = *os.PathError его значение nil
	// data = nil >>> интерфейс не пустой, так как есть информация о типе.
}

func main() {
	err := Foo()
	fmt.Println(err)        // печатает <nil>, так как печать интерфейса показывает его data, которое nil
	fmt.Println(err == nil) // false:  даёт false, потому что сравниваются оба поля интерфейса, так как type = *os.PathError
}

// Отличие пустого интерфейса: интерфейс хранит тип и данные. Разница только в том, что у обычного интерфейса требуется проверка методов, а у пустого — иъ нет (методов).
