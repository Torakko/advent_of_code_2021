package main;

import (
    "fmt"
    "strings"
    "aoc/libs/utils"
    "strconv"
);

/**
  * Start - 23:09:48
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 01);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        if (part1(part1_test_input[i]) != part1_test_output[i]) {
            success = false;
            break;
        }
    }

    fmt.Printf("part1 minitest success: %t! \n", success);
    p1 := part1(input);
    fmt.Printf("part1: %s\n\n", p1);
    
    success = true;
    for i := range part2_test_input {
        if (part2(part2_test_input[i]) != part2_test_output[i]) {
            success = false;
            break;
        }
    }
    fmt.Printf("part2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("part2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `199
200
208
210
200
207
240
269
260
263`,
};
var part1_test_output = []string{
    `7`,
};
func part1(input string) string {
    var inputs = strings.Split(strings.Trim(input, separator), separator)
    var inc = 0;
    var prev = 9999999999999
    for i := range inputs {
        var cur, _ = strconv.Atoi(inputs[i])
        //fmt.Printf("%d\n", prev)
        //fmt.Printf("%d\n", cur)
        if (cur > prev) {
            inc++
        }
        prev = cur
    }
    //fmt.Printf("%d\n", inc)
    return strconv.Itoa(inc)
}

var part2_test_input = part1_test_input

var part2_test_output = []string{
    `5`,
};
func part2(input string) string {
    var inputs = strings.Split(strings.Trim(input, separator), separator);
    inputs = append([]string{"999999999","99999999"}, inputs...)
    var inc = 0;
    var prev = 9999999999999
    for i := 0; i < len(inputs) - 2; i++  {
        var a, _ = strconv.Atoi(inputs[i])
        var b, _ = strconv.Atoi(inputs[i+1])
        var c, _ = strconv.Atoi(inputs[i+2])
        var cur = a+b+c
        //fmt.Printf("%d\n", prev)
        //fmt.Printf("%d\n", cur)
        if (cur > prev) {
            inc++
        }
        prev = cur
    }
    //fmt.Printf("%d\n", inc)
    return strconv.Itoa(inc)
}
