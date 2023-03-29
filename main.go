package main

import "fmt"

func main() {
	input := "selamat malam"

	// looping dengan variable yang berisi string suatu kalimat
	// dan pecahlah kalimat tersebut menjadi 1 per 1 kata
	for _, char := range input {
		fmt.Println(string(char))
	}

	// lakukan perhitungan munculnya kata dari variable tersebut
	// dengan cara mapping golang
	charCount := make(map[string]int)

	for _, char := range input {
		charCount[string(char)] += 1
	}

	fmt.Println(charCount)
}
