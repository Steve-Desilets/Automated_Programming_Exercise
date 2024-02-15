package main

//go:generate go run automated_code_generation.go

import (
	"fmt"
	"log"
	"os"
	"time"
)

// define a type called Series that is a container for a series of data
type Series []Coordinate

// Define a type called Coordinate that holds the x and y values of a point
type Coordinate struct {
	X, Y float64
}

// Define four variables called AnscombeSeries1, AnscombeSeries2, AnscomebeSeries3, and AnscombeSeries4
// Each variable will be of type Series and will be a container for data points stored as Coordinates
// The coordinates will correspond to those defined in Anscombe's Quartet
var AnscombeSeries1 = []Coordinate{
	{10, 8.04},
	{8, 6.95},
	{13, 7.58},
	{9, 8.81},
	{11, 8.33},
	{14, 9.96},
	{6, 7.24},
	{4, 4.26},
	{12, 10.84},
	{7, 4.82},
	{5, 5.68},
}

var AnscombeSeries2 = []Coordinate{
	{10, 9.14},
	{8, 8.14},
	{13, 8.74},
	{9, 8.77},
	{11, 9.26},
	{14, 8.10},
	{6, 6.13},
	{4, 3.10},
	{12, 9.13},
	{7, 7.26},
	{5, 4.74},
}

var AnscombeSeries3 = []Coordinate{
	{10, 7.46},
	{8, 6.77},
	{13, 12.74},
	{9, 7.11},
	{11, 7.81},
	{14, 8.84},
	{6, 6.08},
	{4, 5.39},
	{12, 8.15},
	{7, 6.42},
	{5, 5.73},
}

var AnscombeSeries4 = []Coordinate{
	{8, 6.58},
	{8, 5.76},
	{8, 7.71},
	{8, 8.84},
	{8, 8.47},
	{8, 7.04},
	{8, 5.25},
	{19, 12.50},
	{8, 5.56},
	{8, 7.91},
	{8, 6.89},
}

func main() {
	// Execute 100 iterations of the experimental trials
	// Append the runtime for each experimental trial to a slice called runtimeSlice
	// Each experimental trial should leverage a function called LinearRegression each of the Anscombe series

	var runtimeSlice []int64
	var loopCounter int = 0

	for i := 0; i < 100; i++ {
		// Execute the experimental trial
		// Measure the runtime
		// Append the runtime to the slice
		// runtimeSlice = append(runtimeSlice, runtime)

		startTime := time.Now()

		fmt.Println("Anscombe Quartet 1 - Regression Outputs")
		regression1, _ := LinearRegression(AnscombeSeries1)
		fmt.Println(regression1)

		fmt.Println("Anscombe Quartet 2 - Regression Outputs")
		regression2, _ := LinearRegression(AnscombeSeries2)
		fmt.Println(regression2)

		fmt.Println("Anscombe Quartet 3 - Regression Outputs")
		regression3, _ := LinearRegression(AnscombeSeries3)
		fmt.Println(regression3)

		fmt.Println("Anscombe Quartet 4 - Regression Outputs")
		regression4, _ := LinearRegression(AnscombeSeries4)
		fmt.Println(regression4)

		endTime := time.Now()
		executionTime := endTime.Sub(startTime)
		runtimeSlice = append(runtimeSlice, executionTime.Microseconds())

		loopCounter += 1
	}

	// Print the slice of runtimes
	printSlice(runtimeSlice)

	// Calculate the total runtime and the average runtime
	var runtimeSum int64 = 0
	var averageRuntime int64 = 0

	for i := 0; i < len(runtimeSlice); i++ {
		runtimeSum += runtimeSlice[i]
	}
	averageRuntime = runtimeSum / int64(len(runtimeSlice))

	// Print the total runtime and the average runtime to a file called automatedCodeGenerationOutput.txt
	file, err := os.Create("automatedCodeGenerationOutput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "Total runtime: %d\n", runtimeSum)
	fmt.Fprintf(file, "Average runtime: %d\n", averageRuntime)

}

// Define a function called printSlice that will print out the slice of runtimes from each experimental trial
func printSlice(slice []int64) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

// Define a function called LinearRegression that will calculate the slope and y-intercept of a line of best fit

func LinearRegression(s Series) (regressions Series, err error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("no data points provided")
	}

	var xSum, ySum, xSquaredSum, xySum float64
	for _, c := range s {
		xSum += c.X
		ySum += c.Y
		xSquaredSum += c.X * c.X
		xySum += c.X * c.Y
	}

	n := float64(len(s))
	m := (n*xySum - xSum*ySum) / (n*xSquaredSum - xSum*xSum)
	b := (ySum - m*xSum) / n
	fmt.Println("Regression Slope = ", m, "\nRegression Y-Intercept = ", b)
	return regressions, nil

}
