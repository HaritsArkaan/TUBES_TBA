package main

import (
	"fmt"
)

type Stack [10]string
type state string

func main() {
	var statusAkhir bool
	var kata, status string
	var kalimat [99]string
	var SPOK [10]string
	var idx, i int

	fmt.Println("Input Kalimat (akhiri dengan '.'): ")
	fmt.Scan(&kata)
	for kata != "." {
		kalimat[idx] = kata
		idx++
		fmt.Scan(&kata)
	}

	for i < idx {
		status = token(status, kalimat, i)
		switch status {
		case "subjek":
			SPOK[i] = "S"
		case "predikat":
			SPOK[i] = "P"
		case "objek":
			SPOK[i] = "O"
		case "keterangan":
			SPOK[i] = "K"
		default:
			SPOK[i] = "Z"
		}
		i++
	}
	SPOK[i] = " "
	i++
	PDA(SPOK, i, &statusAkhir)
	if statusAkhir {
		fmt.Println("Kalimat tersebut memenuhi struktur Bahasa Indonesia")
	} else {
		fmt.Println("Kalimat tersebut tidak memenuhi struktur Bahasa Indonesia")
	}

}

func token(status string, kalimat [99]string, i int) string {
	n := kalimat[i]
	if subjekFA(n) {
		return "subjek"
	} else if predikatFA(n) {
		return "predikat"
	} else if objectFA(n) {
		return "objek"
	} else if keteranganFA(n) {
		return "keterangan"
	} else {
		return "notFound"
	}
}

func subjekFA(input string) bool {
	recognized := false
	pattern := [5]string{"harits", "rangga", "zahwa", "mereka", "kami"}
	for i := 0; i < 5; i++ {
		patternMatched := true
		for j := 0; len(input) == len(pattern[i]) && j < len(input) && patternMatched; j++ {
			if pattern[i][j] != input[j] {
				patternMatched = false
			}
		}
		if patternMatched && len(input) == len(pattern[i]) {
			recognized = true
		}
	}
	return recognized
}
func predikatFA(input string) bool {
	recognized := false
	pattern := [5]string{"menyukai", "makan", "minum", "bermain", "belajar"}
	for i := 0; i < 5; i++ {
		patternMatched := true
		for j := 0; len(input) == len(pattern[i]) && j < len(input) && patternMatched; j++ {
			if pattern[i][j] != input[j] {
				patternMatched = false
			}
		}
		if patternMatched && len(input) == len(pattern[i]) {
			recognized = true
		}
	}
	return recognized
}
func objectFA(input string) bool {
	recognized := false
	pattern := [5]string{"bola", "bunga", "kopi", "susu", "kue"}
	for i := 0; i < 5; i++ {
		patternMatched := true
		for j := 0; len(input) == len(pattern[i]) && j < len(input) && patternMatched; j++ {
			if pattern[i][j] != input[j] {
				patternMatched = false
			}
		}
		if patternMatched && len(input) == len(pattern[i]) {
			recognized = true
		}
	}
	return recognized
}
func keteranganFA(input string) bool {
	recognized := false
	pattern := [5]string{"hariini", "kemarin", "dilapangan", "dirumah", "dikampus"}
	for i := 0; i < 5; i++ {
		patternMatched := true
		for j := 0; len(input) == len(pattern[i]) && j < len(input) && patternMatched; j++ {
			if pattern[i][j] != input[j] {
				patternMatched = false
			}
		}
		if patternMatched && len(input) == len(pattern[i]) {
			recognized = true
		}
	}
	return recognized
}

func PDA(parseInput [10]string, n int, status *bool) {
	var S Stack
	var j int = 0
	var current state = "start"
	push(&S, "#", &j)
	current = "q1"
	for i := 0; i <= n && current != "acc"; i++ {
		if current == "q1" {
			if parseInput[i] == "S" {
				push(&S, "x", &j)
				current = "q2"
			} else {
				break
			}
		} else if current == "q2" {
			if parseInput[i] == "P" {
				if top(S, j) == "x" {
					pop(&S, &j)
					current = "q3"
				}
			} else {
				break
			}
		} else if current == "q3" {
			if parseInput[i] == "K" {
				current = "q5"
			} else if parseInput[i] == "O" {
				push(&S, "y", &j)
				current = "q4"
			} else if parseInput[i] == " " {
				if top(S, j) == "#" {
					pop(&S, &j)
					current = "acc"
				}
			} else {
				break
			}
		} else if current == "q4" {
			if parseInput[i] == " " || parseInput[i] == "K" {
				if top(S, j) == "y" {
					pop(&S, &j)
					current = "q5"
				}
			} else {
				break
			}
		} else if current == "q5" {
			if top(S, j) == "#" {
				pop(&S, &j)
				current = "acc"
			} else {
				break
			}
		}
	}
	if current == "acc" {
		*status = true
	} else {
		*status = false
	}
}

func push(S *Stack, x string, n *int) {
	S[*n] = x
	*n++
}

func pop(S *Stack, n *int) string {
	result := S[*n-1]
	*n--
	return result
}

func top(S Stack, n int) string {
	return S[n-1]
}
