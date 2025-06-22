package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input provided. Run './program help' to see what inputs are required.")
		return
	}
	mode := os.Args[1]

	if mode == "reverse" {
		if len(os.Args) != 4 {
			fmt.Println("Incorrect inputs. Run './program help' to see what inputs are required.")
			return
		}
		inputPath := os.Args[2]
		inputContent, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		reversedBytes := reverseBytes(inputContent)

		outputPath := os.Args[3]
		err = os.WriteFile(outputPath, reversedBytes, 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println("Operation done!")
	} else if mode == "copy" {
		if len(os.Args) != 4 {
			fmt.Println("Incorrect inputs. Run './program help' to see what inputs are required.")
			return
		}
		inputPath := os.Args[2]
		inputContent, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		outputPath := os.Args[3]
		err = os.WriteFile(outputPath, inputContent, 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println("Operation done!")
	} else if mode == "duplicate-contents" {
		if len(os.Args) != 4 {
			fmt.Println("Incorrect inputs. Run './program help' to see what inputs are required.")
			return
		}
		inputPath := os.Args[2]
		inputContent, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		n, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if n < 0 {
			fmt.Println("Please provide positive number.")
			return
		}

		duplicatedBytes := duplicateContents(inputContent, n)

		err = os.WriteFile(inputPath, duplicatedBytes, 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println("Operation done!")
	} else if mode == "replace-string" {
		if len(os.Args) != 5 {
			fmt.Println("Incorrect inputs. Run './program help' to see what inputs are required.")
			return
		}
		inputPath := os.Args[2]
		inputContent, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		oldStr := os.Args[3]
		newStr := os.Args[4]
		updatedStr := strings.ReplaceAll(string(inputContent), oldStr, newStr)
		err = os.WriteFile(inputPath, []byte(updatedStr), 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println("Operation done!")
	} else if mode == "help" {
		printHelp()
	} else {
		fmt.Printf("No option called '%s' available\n", mode)
		printHelp()
	}
}

func reverseBytes(inputBytes []byte) []byte {
	inputRunes := []rune(string(inputBytes))

	for i, j := 0, len(inputRunes)-1; i < j; i, j = i+1, j-1 {
		inputRunes[i], inputRunes[j] = inputRunes[j], inputRunes[i]
	}

	outputBytes := []byte(string(inputRunes))
	return outputBytes
}

func duplicateContents(inputBytes []byte, n int) []byte {
	n++
	duplicatedBytes := make([]byte, 0, len(inputBytes)*n)
	for i := 0; i < n; i++ {
		duplicatedBytes = append(duplicatedBytes, inputBytes...)
	}
	return duplicatedBytes
}

func printHelp() {
	fmt.Println("Options:")
	fmt.Println("reverse: ./program reverse <input file (string)> <output file (string)>")
	fmt.Println("copy: ./program copy <input file (string)> <output file (string)>")
	fmt.Println("duplicate-contents: ./program duplicate-contents <input file (string)> <how many times to duplicate (int)>")
	fmt.Println("replace-string: ./program replace-string <input file (string)> <old string (string)> <new string (string)>")
}
