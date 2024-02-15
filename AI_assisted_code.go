package main

// This program is used to calculate the runtime of the Linear Regression function
// The Linear Regression function is executed 100 times for each of the four Anscombe Quartet data series
// The runtime for each of the 400 experimental trials is appended to a slice called runtime_slice
// The average runtime for each of the 400 experimental trials is calculated and printed to the console

// import the packages that will be needed for this project
import (
	"fmt"
	"os"
	"time"
)

// Series is a container for a series of data
type Series []Coordinate

// Coordinate holds the data in a series
type Coordinate struct {
	X, Y float64
}

// Define the four Anscombe Quartet data series
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
	{14, 8.1},
	{6, 6.13},
	{4, 3.1},
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
	{19, 12.5},
	{8, 5.56},
	{8, 7.91},
	{8, 6.89},
}

func main() {
	// Execute 100 iterations of the experimental trials
	// Append the runtime for each experimental trial to a slice called runtime_slice

	var runtimeSlice []int64

	outputFile, err := os.Create("aiAssistedCodeOutput.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	loopCounter := 0
	for i := 0; i < 100; i++ {
		// Execute the Linear Regression function for each of the four Anscombe Quartet data series
		// Print the slope and intercept of the regression line for each of the four data series
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

		loopCounter += i
	}
	printSlice(runtimeSlice)

	// Calculate and print the average runtime across each of the experimental trials
	var runtimeSum int64 = 0

	for i := 0; i < 100; i++ {
		runtimeSum += runtimeSlice[i]
	}
	//avgRuntime := (float64(runtimeSum)) / (float64(100))
	avgRuntime := runtimeSum / 100

	fmt.Fprintf(outputFile, "Total runtime: %d\n", runtimeSum)
	fmt.Fprintf(outputFile, "Average runtime: %d\n", avgRuntime)

}

// Define a function that will print out the slice of runtimes from each experimental trial
func printSlice(s []int64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// LinearRegression finds the least squares linear regression on data series
// Then it prints the slope and intercept of the regression line
func LinearRegression(s Series) (regressions Series, err error) {

	// If the data series is empty, return an error
	if len(s) == 0 {
		return nil, fmt.Errorf("Empty data series")
	}

	// Placeholder for the math to be done
	var sum [5]float64

	// Loop over data keeping index in place
	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X
		sum[3] += s[i].X * s[i].Y
		sum[4] += s[i].Y * s[i].Y
	}

	// Find gradient and intercept
	f := float64(i)
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (gradient * sum[0] / f)
	fmt.Println("Regression Slope = ", gradient, "\nRegression Intercept = ", intercept)

	return regressions, nil
}
