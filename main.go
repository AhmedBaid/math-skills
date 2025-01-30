package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	mathFunctions "mathSkills/mathFunctions"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Error while parsing arguments")
		return
	}
	Filepath := arguments[1]
	File, err := os.Open(Filepath)
	if err != nil {
		fmt.Println("Error while opening file", err)
	}
	defer File.Close()
	scanner := bufio.NewScanner(File)
	var numbers []float64
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number", err)
		}
		numbers = append(numbers, float64(num))
	}
	fmt.Println("Average:", fmt.Sprintf("%.0f", math.Round(mathFunctions.Average(numbers))))
	fmt.Println("Median:", fmt.Sprintf("%.0f", math.Round(mathFunctions.Median(numbers))))
	fmt.Println("Variance:", fmt.Sprintf("%.0f", math.Round(mathFunctions.Variance(numbers))))
	fmt.Println("Standard Deviation:", fmt.Sprintf("%.0f", math.Round(mathFunctions.StandardDeviation(numbers))))
}
