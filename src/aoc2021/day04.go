package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 20:48:29
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 04);
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
    `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

    22 13 17 11  0
     8  2 23  4 24
    21  9 14 16  7
     6 10  3 18  5
     1 12 20 15 19

     3 15  0  2 22
     9 18 13 17  5
    19  8  7 25 23
    20 11 10 24  4
    14 21 16 12  6

    14 21 17 24  4
    10 16 15  9 19
    18  8 23 26 20
    22 11 13  6  5
     2  0 12  3  7`,
};
var part1_test_output = []string{
    `4512`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var drawn = strings.Split(inputs[0], ",")
    //fmt.Printf("Drawn: %s\n", drawn)
    inputs = inputs[1:]
    var nums = make([][]int, len(inputs))
    for i := range inputs {
        var line = strings.Split(strings.ReplaceAll(inputs[i], "  ", " "), " ")
        var new_line = make([]int, len(line))
        for j, number := range line {
            new_line[j], _ = strconv.Atoi(number)
        }
        nums[i] = new_line
    }

    //fmt.Printf("nums: %d\n", nums)
    var bingos = make([]Bingo, 0)
    for line := 0; line < len(nums); line += 6 {
        var duttMatrix = make([][]bool, 5)
        for i, _ := range duttMatrix {
            duttMatrix[i] = make([]bool, 5)
        }
        var bingo = Bingo{input:nums[line+1:line+6],
                          dutt:duttMatrix}
        bingos = append(bingos, bingo)
    }
    //fmt.Printf("Bingos: %d\n", len(bingos))
    for _, bingo := range bingos {
        bingo.Print()
    }

    for _, num := range drawn {
        for _, bingo := range bingos {
            var num2, _ = strconv.Atoi(num)
            //fmt.Printf("Draw: %d\n", num2)
            bingo.Dutt(num2)
            bingo.Print()
            if bingo.Victory() {
                return strconv.Itoa(bingo.Sum() * num2)
            }
        }
    }
    return "";
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `1924`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var drawn = strings.Split(inputs[0], ",")
    //fmt.Printf("Drawn: %s\n", drawn)
    inputs = inputs[1:]
    var nums = make([][]int, len(inputs))
    for i := range inputs {
        var line = strings.Split(strings.ReplaceAll(inputs[i], "  ", " "), " ")
        var new_line = make([]int, len(line))
        for j, number := range line {
            new_line[j], _ = strconv.Atoi(number)
        }
        nums[i] = new_line
    }

    //fmt.Printf("nums: %d\n", nums)
    var bingos = make([]Bingo, 0)
    for line := 0; line < len(nums); line += 6 {
        var duttMatrix = make([][]bool, 5)
        for i, _ := range duttMatrix {
            duttMatrix[i] = make([]bool, 5)
        }
        var bingo = Bingo{input:nums[line+1:line+6],
                          dutt:duttMatrix}
        bingos = append(bingos, bingo)
    }
    var bingosRes = make([]bool, len(bingos))
    var lastBingo = 0
    var lastBingoNum = 0
    for _, num := range drawn {
        fmt.Printf("Bingos: %d\n", len(bingos))
        var num2, _ = strconv.Atoi(num)
        fmt.Printf("Draw: %d\n", num2)
        for id, bingo := range bingos {
            if !bingosRes[id] {
                bingo.Dutt(num2)
                //bingo.Print()
                if bingo.Victory() {
                    lastBingo = id
                    lastBingoNum = num2
                    bingosRes[id] = true
                }
            }

        }
    }
    return strconv.Itoa(bingos[lastBingo].Sum() * lastBingoNum)
}

func RemoveFromSlice(s []Bingo, i int) []Bingo {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}
type Bingo struct {
    input [][]int
    dutt [][]bool
}

func (b Bingo) Dutt(num int) {
    for i, line := range b.input {
        for j, val := range line {
            if val == num {
                b.dutt[i][j] = true
                //fmt.Printf("Dutting in %d %d\n", i, j)
            }
        }
    }
}


func (b Bingo) Victory() bool {
    for i, line := range b.dutt {
        if line[0] && line[1] && line[2] && line[3] && line[4] {
            return true
        }
        if b.dutt[0][i] && b.dutt[1][i] && b.dutt[2][i] && b.dutt[3][i] && b.dutt[4][i] {
            return true
        }
    }
    return false
}

func (b Bingo) Sum() int{
    var sum = 0
    for i, line := range b.dutt {
        for j, val := range line {
            if !val {
                sum += b.input[i][j]
            }
        }
    }
    fmt.Printf("Sum is %d\n", sum)
    return sum
}

func (b Bingo) Print() {
    fmt.Printf("%d\n", b.input)
    fmt.Printf("%v\n", b.dutt)
}





