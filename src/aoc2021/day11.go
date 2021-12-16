package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 19:58:24
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 11);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        var res1 = part1(part1_test_input[i])
        if (res1 != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    res1,
                    part1_test_output[i]);
            break;
        }
    }

    fmt.Printf("part1 minitest success: %t! \n", success);
    if success == false {
        return
    }
    p1 := part1(input);
    fmt.Printf("part1: %s\n\n", p1);
    
    success = true;
    for i := range part2_test_input {
        var res2 = part2(part2_test_input[i])
        if (res2 != part2_test_output[i]) {
            success = false;
            fmt.Printf("part2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    res2,
                    part2_test_output[i]);
            break;
        }
    }
    fmt.Printf("part2 minitest success: %t! \n", success);
    if success == false {
        return
    }
    p2 := part2(input);
    fmt.Printf("part2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `5483143223
    2745854711
    5264556173
    6141336146
    6357385478
    4167524645
    2176841721
    6882881134
    4846848554
    5283751526`,
};
var part1_test_output = []string{
    `1656`,
};
func part1(input string) string {
    // Step 0 create matrix
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var size = len(inputs)
    var matrix [][]int
    for _, line := range(inputs) {
        var newLine []int
        for _, rawVal := range(line) {
            var val, _ = strconv.Atoi(string(rawVal))
            newLine = append(newLine, val)
        }
        matrix = append(matrix, newLine)
    }
    fmt.Printf("Step: %v\n", 0);
    utils.PrintMatrix(matrix)
    var flashCount = 0
    for step := 1; step <= 100; step++ {
        fmt.Printf("Step: %v\n", step);
        // Step 1 increase 1
        for i, line := range(matrix) {
            for j := range(line) {
                matrix[i][j]++
            }
        }
        //fmt.Printf("Increase done:\n");
        //utils.PrintMatrix(matrix)

        // Step 2 chain flashing
        var flashMatrix = getFlashMatrix(size, size)
        for {
            var anyFlashed = false
            for i, line := range(matrix) {
                for j := range(line) {
                    if matrix[i][j] > 9 && !flashMatrix[i][j] {
                        //fmt.Printf("Flashed in: i %v j %v\n", i, j);
                        flashCount++
                        anyFlashed = true
                        flashMatrix[i][j] = true
                        if !(i-1<0) && !(j-1<0) {
                            //fmt.Printf("NV\n");
                            matrix[i-1][j-1]++
                        }
                        if !(i-1<0) {
                            //fmt.Printf("N\n");
                            matrix[i-1][j]++
                        }
                        if !(i-1<0) && !(j+1>size-1) {
                            //fmt.Printf("NE\n");
                            matrix[i-1][j+1]++
                        }
                        if !(j-1<0) {
                            //fmt.Printf("V\n");
                            matrix[i][j-1]++
                        }
                        if !(j+1>size-1) {
                            //fmt.Printf("E\n");
                            matrix[i][j+1]++
                        }
                        if !(i+1>size-1) && !(j-1<0) {
                            //fmt.Printf("SV\n");
                            matrix[i+1][j-1]++
                        }
                        if !(i+1>size-1) {
                            //fmt.Printf("S\n");
                            matrix[i+1][j]++
                        }
                        if !(i+1>size-1) && !(j+1>size-1) {
                            //fmt.Printf("SE\n");
                            matrix[i+1][j+1]++
                        }
                        //utils.PrintMatrix(matrix)
                    }
                }
            }
            if !anyFlashed {
                break
            }
        }

        // Step 3 Flased set to 0
        for i, line := range(matrix) {
            for j := range(line) {
                if flashMatrix[i][j] {
                    matrix[i][j] = 0
                }
            }
        }
        //fmt.Printf("Cleaned:\n");
        utils.PrintMatrix(matrix)

    }

    return strconv.Itoa(flashCount);
}

func getFlashMatrix(x int, y int) [][]bool {
    var matrix [][]bool
    for i := 0; i<x; i++ {
        var newLine = make([]bool, y)
        matrix = append(matrix, newLine)
    }
    return matrix
}


type Point struct {
    x int
    y int
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `195`,
};
func part2(input string) string {
    // Step 0 create matrix
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var size = len(inputs)
    var matrix [][]int
    for _, line := range(inputs) {
        var newLine []int
        for _, rawVal := range(line) {
            var val, _ = strconv.Atoi(string(rawVal))
            newLine = append(newLine, val)
        }
        matrix = append(matrix, newLine)
    }
    fmt.Printf("Step: %v\n", 0);
    utils.PrintMatrix(matrix)
    var step = 0
    for {
        step++
        fmt.Printf("Step: %v\n", step);
        // Step 1 increase 1
        for i, line := range(matrix) {
            for j := range(line) {
                matrix[i][j]++
            }
        }
        //fmt.Printf("Increase done:\n");
        //utils.PrintMatrix(matrix)

        // Step 2 chain flashing
        var flashMatrix = getFlashMatrix(size, size)
        for {
            var anyFlashed = false
            for i, line := range(matrix) {
                for j := range(line) {
                    if matrix[i][j] > 9 && !flashMatrix[i][j] {
                        //fmt.Printf("Flashed in: i %v j %v\n", i, j);
                        anyFlashed = true
                        flashMatrix[i][j] = true
                        if !(i-1<0) && !(j-1<0) {
                            //fmt.Printf("NV\n");
                            matrix[i-1][j-1]++
                        }
                        if !(i-1<0) {
                            //fmt.Printf("N\n");
                            matrix[i-1][j]++
                        }
                        if !(i-1<0) && !(j+1>size-1) {
                            //fmt.Printf("NE\n");
                            matrix[i-1][j+1]++
                        }
                        if !(j-1<0) {
                            //fmt.Printf("V\n");
                            matrix[i][j-1]++
                        }
                        if !(j+1>size-1) {
                            //fmt.Printf("E\n");
                            matrix[i][j+1]++
                        }
                        if !(i+1>size-1) && !(j-1<0) {
                            //fmt.Printf("SV\n");
                            matrix[i+1][j-1]++
                        }
                        if !(i+1>size-1) {
                            //fmt.Printf("S\n");
                            matrix[i+1][j]++
                        }
                        if !(i+1>size-1) && !(j+1>size-1) {
                            //fmt.Printf("SE\n");
                            matrix[i+1][j+1]++
                        }
                        //utils.PrintMatrix(matrix)
                    }
                }
            }
            if !anyFlashed {
                break
            }
        }

        // Step 3 Flased set to 0
        for i, line := range(matrix) {
            for j := range(line) {
                if flashMatrix[i][j] {
                    matrix[i][j] = 0
                }
            }
        }
        //fmt.Printf("Cleaned:\n");
        utils.PrintMatrix(matrix)

        // Step 4 check if all flashed
        if allFlashed(flashMatrix) {
            return strconv.Itoa(step);
        }
    }
    return ""
}

func allFlashed(matrix [][]bool) bool {
    for i, line := range(matrix) {
        for j := range(line) {
            if !matrix[i][j] {
                return false
            }
        }
    }
    return true
}
