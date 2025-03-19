package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

// Функция для выполнения всех шагов
func processVariables(variables []interface{}) (string, string) {

	var combinedString string
	for _, v := range variables {
		fmt.Printf("Тип переменной %v: %s\n", v, reflect.TypeOf(v))
		combinedString += fmt.Sprintf("%v ", v)
	}
	// Преобразуем строку в срез рун
	runes := []rune(combinedString)
	// Хэшируем срез рун SHA256 с добавлением соли
	salt := "go-2024"
	saltedRunes := append(runes[:len(runes)/2], []rune(salt)...)
	saltedRunes = append(saltedRunes, runes[len(runes)/2:]...)
	hash := sha256.Sum256([]byte(string(saltedRunes)))
	return combinedString, hex.EncodeToString(hash[:])
}
func main() {
	// Создаем переменные различных типов
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i
	variables := []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	combinedString, hash := processVariables(variables)
	fmt.Printf("Объединенная строка: %s\n", combinedString)
	fmt.Printf("SHA256 Hash: %s\n", hash)
}
