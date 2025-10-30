package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	var intDecimalVar int = 10
	var intOctVar int = 052
	var intHexVar int = 0x2A
	var floatVar float64 = 1.23
	var stringVar string = "task1"
	var boolVar bool = true
	var complexVar complex64 = 1 + 1i

	printType(intDecimalVar, intOctVar, intHexVar, floatVar, stringVar, boolVar, complexVar)

	convStringSlice := convertToString(intDecimalVar, intOctVar, intHexVar, floatVar, stringVar, boolVar, complexVar)
	fmt.Println("Слайс строковых переменных: ", convStringSlice)

	concString := concatenateStrings(convStringSlice)
	fmt.Println("Строка с переменными: ", concString)

	sliceRunes := stringToRunes(concString)
	fmt.Println("Слайс рун: ", sliceRunes)

	hashRunes := hashRunes(sliceRunes)
	fmt.Println("Захешированные руны", hashRunes)

}

func printType(intDecimalVar int, intOctVar int, intHexVar int, floatVar float64, stringVar string, boolVar bool, complexVar complex64) {
	fmt.Printf("type of intDecimalVar is %s \n", reflect.TypeOf(intDecimalVar))
	fmt.Printf("type of intOctVar is %s \n", reflect.TypeOf(intOctVar))
	fmt.Printf("type of intHexVar is %s \n", reflect.TypeOf(intHexVar))
	fmt.Printf("type of floatVar is %s \n", reflect.TypeOf(floatVar))
	fmt.Printf("type of stringVar is %s \n", reflect.TypeOf(stringVar))
	fmt.Printf("type of boolVar is %s \n", reflect.TypeOf(boolVar))
	fmt.Printf("type of complexVar is %s \n", reflect.TypeOf(complexVar))
}

func convertToString(numDecimal, numOctal, numHexadecimal int, pi float64, name string, isActive bool, complexNum complex64) []string {
	var stringsSlice []string

	stringsSlice = append(stringsSlice, strconv.Itoa(numDecimal))
	stringsSlice = append(stringsSlice, strconv.Itoa(numOctal))
	stringsSlice = append(stringsSlice, strconv.Itoa(numHexadecimal))
	stringsSlice = append(stringsSlice, strconv.FormatFloat(pi, 'f', -1, 64))
	stringsSlice = append(stringsSlice, name)
	stringsSlice = append(stringsSlice, strconv.FormatBool(isActive))
	stringsSlice = append(stringsSlice, fmt.Sprint(complexNum))

	return stringsSlice
}

func concatenateStrings(strSlice []string) string {
	return strings.Join(strSlice, " ")
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

func hashRunes(runes []rune) string {
	str := string(runes)

	// Добавляем соль в середину
	mid := len(str) / 2
	strWithSalt := str[:mid] + "go-2024" + str[mid:]

	// Вычисляем хеш
	hasher := sha256.New()
	hasher.Write([]byte(strWithSalt))
	return hex.EncodeToString(hasher.Sum(nil))

}
