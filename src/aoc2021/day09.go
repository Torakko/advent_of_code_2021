package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "sort"
);

/**
  * Start - 16:00:30
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 9);
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
    `2199943210
    3987894921
    9856789892
    8767896789
    9899965678`,
};
var part1_test_output = []string{
    `15`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var matrix [][]int
    var nineLine []int
    for i := 0; i<len(inputs[0])+2; i++ {
        nineLine = append(nineLine, 9)
    }
    matrix = append(matrix, nineLine)
    for _, line := range(inputs) {
        var newLine []int
        newLine = append(newLine, 9)
        for _, rawVal := range(line) {
            var val, _ = strconv.Atoi(string(rawVal))
            newLine = append(newLine, val)
        }
        newLine = append(newLine, 9)
        matrix = append(matrix, newLine)
    }
    matrix = append(matrix, nineLine)
    //utils.PrintMatrix(matrix)

    var lowpoints []int
    for i := 1; i < len(matrix)-1; i++ {
        for j := 1; j < len(matrix[i])-1; j++ {
            if matrix[i][j] >= matrix[i+1][j] {
                continue
            } else if matrix[i][j] >= matrix[i-1][j] {
                continue
            } else if matrix[i][j] >= matrix[i][j+1] {
                continue
            } else if matrix[i][j] >= matrix[i][j-1] {
                continue
            } else {
                lowpoints = append(lowpoints, matrix[i][j])
            }
        }
    }
    //fmt.Printf("Lowpoints %v\n", lowpoints)
    return strconv.Itoa(utils.SumArray(lowpoints)+len(lowpoints));
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `1134`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var matrix [][]int
    var nineLine []int
    for i := 0; i<len(inputs[0])+2; i++ {
        nineLine = append(nineLine, 9)
    }
    matrix = append(matrix, nineLine)
    for _, line := range(inputs) {
        var newLine []int
        newLine = append(newLine, 9)
        for _, rawVal := range(line) {
            var val, _ = strconv.Atoi(string(rawVal))
            newLine = append(newLine, val)
        }
        newLine = append(newLine, 9)
        matrix = append(matrix, newLine)
    }
    matrix = append(matrix, nineLine)
    //utils.PrintMatrix(matrix)

    var lowpoints []Point
    for i := 1; i < len(matrix)-1; i++ {
        for j := 1; j < len(matrix[i])-1; j++ {
            if matrix[i][j] >= matrix[i+1][j] {
                continue
            } else if matrix[i][j] >= matrix[i-1][j] {
                continue
            } else if matrix[i][j] >= matrix[i][j+1] {
                continue
            } else if matrix[i][j] >= matrix[i][j-1] {
                continue
            } else {
                lowpoints = append(lowpoints, Point{x:i, y:j})
            }
        }
    }
    fmt.Printf("Lowpoints %v\n", lowpoints)
    
    var basins []int
    for _, p := range(lowpoints) {
        basins = append(basins, p.Basin(matrix))
    }

    sort.Slice(basins, func(i, j int) bool {
        return basins[i] > basins[j]
    })
    fmt.Printf("Basins %v\n", basins)
    return strconv.Itoa(basins[0]*basins[1]*basins[2]);
}

func contains(s []Point, e Point) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

type Point struct {
    x int
    y int
}

func (p Point) Basin(matrix [][]int) int {
    var points = []Point{p}
    for {
        var newPoints = 0
        for _, point := range(points) {
            for _, neighbour := range(point.Neighbours(matrix)) {
                if !contains(points, neighbour) {
                    points = append(points, neighbour)
                    newPoints++
                }
            }
        }
        if newPoints == 0 {
            break
        }
    }
    //fmt.Printf("points: %v\n", points)
    return len(points)
}

func (p Point) Neighbours(matrix [][]int) []Point {
    var ret []Point
    //fmt.Printf("Nieighbours for (%v,%v) %v\n", p.x, p.y, matrix[p.x][p.y])
    if matrix[p.x-1][p.y] != 9 {
        //fmt.Printf("Not 9: (%v,%v) %v\n", p.x-1, p.y, matrix[p.x-1][p.y])
        ret = append(ret, Point{x:p.x-1, y:p.y})
    }
    if matrix[p.x+1][p.y] != 9 {
        //fmt.Printf("Not 9: (%v,%v) %v\n", p.x+1, p.y, matrix[p.x+1][p.y])
        ret = append(ret, Point{x:p.x+1, y:p.y})
    }
    if matrix[p.x][p.y-1] != 9 {
        //fmt.Printf("Not 9: (%v,%v) %v\n", p.x, p.y-1, matrix[p.x][p.y-1])
        ret = append(ret, Point{x:p.x, y:p.y-1})
    }
    if matrix[p.x][p.y+1] != 9 {
        //fmt.Printf("Not 9: (%v,%v) %v\n", p.x, p.y+1, matrix[p.x][p.y+1])
        ret = append(ret, Point{x:p.x, y:p.y+1})
    }
    return ret
}


