# Math Statistics Calculator

## Description
This project is a command-line tool that reads numerical data from a file and computes key statistical metrics:
- **Average**
- **Median**
- **Variance**
- **Standard Deviation**

## Features
- Reads numerical values from a file (one number per line)
- Computes statistical metrics
- Uses QuickSort for sorting in median calculation
- Implements mathematical functions in a separate package

## File Structure
```
math-stats/
│── main.go
│── mathFunctions/
│   ├── average.go
│   ├── median.go
│   ├── variance.go
│   ├── standardDeviation.go
│── data.txt (example input file)
│── README.md
```

## Installation
Make sure you have [Go installed](https://go.dev/dl/). Clone the repository and navigate to the project directory.
```sh
git clone <repository-url>
cd math-stats
```

## Usage
Compile and run the program, providing a file containing numerical data as an argument:
```sh
go run main.go data.txt
```

### Example Input (data.txt)
```
189
113
121
114
145
110
```

### Example Output
```
Average: 132
Median: 121
Variance: 1005
Standard Deviation: 31
```

## Implementation Details
- **`main.go`**: Reads input data, processes it, and prints results.
- **`mathFunctions/average.go`**: Calculates the arithmetic mean.
- **`mathFunctions/median.go`**: Uses QuickSort to sort and find the median.
- **`mathFunctions/variance.go`**: Computes variance.
- **`mathFunctions/standardDeviation.go`**: Computes standard deviation using variance.

## License
This project is open-source and available under the MIT License.