package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	series := [4][11][2]float64{
		{
			{10, 8.04}, {8, 6.95}, {13, 7.58}, {9, 8.81}, {11, 8.33}, {14, 9.96}, {6, 7.24}, {4, 4.26}, {12, 10.84}, {7, 4.82}, {5, 5.68},
		},
		{
			{10, 9.14}, {8, 8.14}, {13, 8.74}, {9, 8.77}, {11, 9.26}, {14, 8.10}, {6, 6.13}, {4, 3.10}, {12, 9.13}, {7, 7.26}, {5, 4.74},
		},
		{
			{10, 7.46}, {8, 6.77}, {13, 12.74}, {9, 7.11}, {11, 7.81}, {14, 8.84}, {6, 6.08}, {4, 5.39}, {12, 8.15}, {7, 6.42}, {5, 5.73},
		},
		{
			{8, 6.58}, {8, 5.76}, {8, 7.71}, {8, 8.84}, {8, 8.47}, {8, 7.04}, {8, 5.25}, {19, 12.50}, {8, 5.56}, {8, 7.91}, {8, 6.89},
		},
	}

	iterations := 100
	totalExecutionTime := time.Duration(0)

	outputFile, err := os.Create("aiGeneratedCodeOutput.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	for i := 0; i < iterations; i++ {
		start := time.Now()

		for j, data := range series {
			x := make([]float64, len(data))
			y := make([]float64, len(data))

			for k, point := range data {
				x[k] = point[0]
				y[k] = point[1]
			}

			slope, intercept := linearRegression(x, y)
			output := fmt.Sprintf("Series %d: y = %.2f*x + %.2f\n", j+1, slope, intercept)
			outputFile.WriteString(output)
		}

		elapsed := time.Since(start)
		totalExecutionTime += elapsed
	}

	averageExecutionTime := totalExecutionTime / time.Duration(iterations)
	totalExecutionTimeString := strconv.Itoa(int(totalExecutionTime.Microseconds()))
	averageExecutionTimeString := strconv.Itoa(int(averageExecutionTime.Microseconds()))

	outputFile.WriteString("\nTotal execution time: " + totalExecutionTimeString + " microseconds\n")
	outputFile.WriteString("Average iteration execution time: " + averageExecutionTimeString + " microseconds\n")
}

func linearRegression(x, y []float64) (float64, float64) {
	n := float64(len(x))
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	return slope, intercept
}
