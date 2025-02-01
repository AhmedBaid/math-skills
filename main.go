package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	mathFunctions "mathSkills/mathFunctions"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Error while parsing arguments")
		return
	}
	Filepath := arguments[1]
	if !strings.HasSuffix(Filepath, "txt") {
		fmt.Println("the file path should be a txt file")
		return
	}
	File, err := os.Open(Filepath)
	if err != nil {
		fmt.Println("Error while opening file", err)
		return
	}
	defer File.Close()
	scanner := bufio.NewScanner(File)
	var numbers []float64
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number", err)
			return
		}
		numbers = append(numbers, float64(num))
	}
	fmt.Println("Average:", int(math.Round(mathFunctions.Average(numbers))))
	fmt.Println("Median:", int(math.Round(mathFunctions.Median(numbers))))
	fmt.Println("Variance:", int(math.Round(mathFunctions.Variance(numbers))))
	fmt.Println("Standard Deviation:", int(math.Round(mathFunctions.StandardDeviation(numbers))))
}
