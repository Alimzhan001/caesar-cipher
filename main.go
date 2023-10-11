package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	kz = "АӘБВГҒДЕЁЖЗИЙКҚЛМНҢОӨПРСТУҰҮФХҺЦЧШЩЪЫІЬЭЮЯаәбвгғдеёжзийкқлмнңоөпрстуұүфхһцчшщъыіьэюя"
	en = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ru = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя"
)

func DoCipher(text **string, encStr *string, alphabet, method string, keyInt int) {
	var words, upper, lower int

	switch alphabet {
	case en:
		words, upper, lower = 26, 25, 51
	case kz:
		words, upper, lower = 42, 41, 83
	case ru:
		words, upper, lower = 33, 32, 64
	}

	if method == "1" {
		for _, char := range **text {
			if strings.ContainsRune(alphabet, char) {
				index := strings.IndexRune(alphabet, char)
				if alphabet != en {
					index = index / 2
				}
				if index <= upper {
					if index+keyInt > upper {
						*encStr += string([]rune(alphabet)[index-words+keyInt])
					} else {
						*encStr += string([]rune(alphabet)[index+keyInt])
					}
				} else if index > upper {
					if index+keyInt > lower {
						*encStr += string([]rune(alphabet)[index-words+keyInt])
					} else {
						*encStr += string([]rune(alphabet)[index+keyInt])
					}
				}

			} else {
				*encStr += string(char)
			}
		}
	} else {
		for _, char := range **text {
			if strings.ContainsRune(alphabet, char) {
				index := strings.IndexRune(alphabet, char)
				if alphabet != en {
					index = index / 2
				}
				if index <= upper {
					if index-keyInt < 0 {
						*encStr += string([]rune(alphabet)[words+index-keyInt])
					} else {
						*encStr += string([]rune(alphabet)[index-keyInt])
					}
				} else if index > upper {
					if index-keyInt <= upper {
						*encStr += string([]rune(alphabet)[words+index-keyInt])
					} else {
						*encStr += string([]rune(alphabet)[index-keyInt])
					}
				}

			} else {
				*encStr += string(char)
			}
		}
	}

}

func Cipher(text *string, method string) string {
	if method == "1" {
		fmt.Println("to encrypt using English alphabet type number '1'\nto encrypt using Kazakh alphabet type number '2' \nto encrypt using Russian alphabet type number '3'")
	} else if method == "2" {
		fmt.Println("to decrypt using English alphabet type number '1'\nto decrypt using Kazakh alphabet type number '2' \nto decrypt using Russian alphabet type number '3'")
	}

	reader := bufio.NewReader(os.Stdin)
	alphabetChoice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	alphabetChoice = strings.TrimSpace(alphabetChoice)

	fmt.Println("type a key")
	reader = bufio.NewReader(os.Stdin)
	key, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	key = strings.TrimSpace(key)
	keyInt, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
	}

	var alphabet string
	switch alphabetChoice {
	case "1":
		alphabet = en
	case "2":
		alphabet = kz
	case "3":
		alphabet = ru
	default:
		fmt.Println("Error: bad argument")
		return ""
	}

	var encStr string

	if method == "1" {
		DoCipher(&text, &encStr, alphabet, method, keyInt)
	} else if method == "2" {
		DoCipher(&text, &encStr, alphabet, method, keyInt)
	} else {
		fmt.Println("Error: bad argument")
	}

	return encStr
}

func main() {

	fmt.Println("Type number '1' to encrypt, type number '2' to decrypt:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSpace(line)

	file, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	var finalStr string

	finalStr = Cipher(&text, line)

	err = ioutil.WriteFile("Result.txt", []byte(finalStr), 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Encrypted successfully: check Result.txt file")
	}

}
