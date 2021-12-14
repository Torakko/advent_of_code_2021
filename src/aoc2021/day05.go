package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 20:09:30
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 05);
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
    `0,9 -> 5,9
    8,0 -> 0,8
    9,4 -> 3,4
    2,2 -> 2,1
    7,0 -> 7,4
    6,4 -> 2,0
    0,9 -> 2,9
    3,4 -> 1,4
    0,0 -> 8,8
    5,5 -> 8,2`,
};
var part1_test_output = []string{
    `5`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var size = 1000
    var matrix = make([][]int, size)
    for i, _ := range matrix {
        matrix[i] = make([]int, size)
    }
    for _, line := range inputs {
        var split = strings.Split(line, " -> ")
        //fmt.Printf("A %s B %s\n", split[0], split[1])
        var a = strings.Split(split[0], ",")
        var b = strings.Split(split[1], ",")
        var ay, _ = strconv.Atoi(a[0])
        var ax, _ = strconv.Atoi(a[1])
        var by, _ = strconv.Atoi(b[0])
        var bx, _ = strconv.Atoi(b[1])
        if ax == bx {
            if ay < by {
                filly(matrix, ax, ay, by)
            } else {
                filly(matrix, ax, by, ay)
            }
        } else if ay == by {
            if ax < bx {
                fillx(matrix, ay, ax, bx)
            } else {
                fillx(matrix, ay, bx, ax)
            }
        } else {
            continue
        }
        //utils.PrintMatrix(matrix)
    }
    var ret = count(matrix)
    return strconv.Itoa(ret);
}

func filly(matrix [][]int, x int, from int, to int) {
    for i := from; i <= to; i++ {
        matrix[x][i]++
    }
}

func fillx(matrix [][]int, y int, from int, to int) {
    for i := from; i <= to; i++ {
        matrix[i][y]++
    }
}

func count(matrix [][]int) int {
    var ret = 0
    for _, line := range matrix {
        for _, val := range line {
            if val > 1 {
                ret++
            }
        }
    }
    return ret
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `12`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var size = 1000
    var matrix = make([][]int, size)
    for i, _ := range matrix {
        matrix[i] = make([]int, size)
    }
    for _, line := range inputs {
        var split = strings.Split(line, " -> ")
        //fmt.Printf("A %s B %s\n", split[0], split[1])
        var a = strings.Split(split[0], ",")
        var b = strings.Split(split[1], ",")
        var ay, _ = strconv.Atoi(a[0])
        var ax, _ = strconv.Atoi(a[1])
        var by, _ = strconv.Atoi(b[0])
        var bx, _ = strconv.Atoi(b[1])
        if ax == bx {
            if ay < by {
                filly(matrix, ax, ay, by)
            } else {
                filly(matrix, ax, by, ay)
            }
        } else if ay == by {
            if ax < bx {
                fillx(matrix, ay, ax, bx)
            } else {
                fillx(matrix, ay, bx, ax)
            }
        } else {
            if ax < bx {
                filld(matrix, ax, ay, bx, by)
            } else {
                filld(matrix, bx, by, ax, ay)
            }
        }
        //utils.PrintMatrix(matrix)
    }
    var ret = count(matrix)
    return strconv.Itoa(ret);
}

func filld(matrix [][]int, ax int, ay int, bx int, by int) {
        for i := 0; i <= bx-ax; i++ {
            if ay <= by {
                matrix[ax+i][ay+i]++
            } else {
                matrix[ax+i][ay-i]++
            }
        }
}

