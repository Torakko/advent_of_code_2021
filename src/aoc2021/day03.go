package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
    "math"
);

/**
  * Start - 22:20:28
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 03);

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
    `00100
    11110
    10110
    10111
    10101
    01111
    00111
    11100
    10000
    11001
    00010
    01010`,
};
var part1_test_output = []string{
    `198`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    //var nums, _ = utils.StrToInt_array(inputs);
    var gamma = 0
    var epsilon = 0
    for pos := 0; pos < len(inputs[0]); pos++ {
        var one = 0
        var zero = 0
        for line := 0; line < len(inputs); line++  {
            if (inputs[line][len(inputs[line])-pos-1] == '1') {
                one++
            } else {
                zero++
            }
        }
        //fmt.Printf("pos %d: 1:%d 0:%d\n", pos, one, zero)
        if (one > zero) {
            gamma += int(math.Pow(2,float64(pos)))
            //fmt.Printf("gamma pos %d: %d\n", pos, int(math.Pow(2,float64(pos))))
        } else {
            epsilon += int(math.Pow(2,float64(pos)))
        }
    }

    fmt.Printf("gamma: %b %d\n", gamma, gamma)
    fmt.Printf("epsilon: %b %d\n", epsilon, epsilon)

    return strconv.Itoa(gamma*epsilon);
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `230`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var inputsCopy = inputs
    var oxygenGeneratorRating = 0
    var co2ScrupperRating = 0
    for pos := 0; pos < len(inputs[0]); pos++ {
        inputs = return_common(inputs, pos)
        if len(inputs) == 1 {
            oxygenGeneratorRating = utils.StrByteToInt(inputs[0])
            break
        }
    }
    for pos := 0; pos < len(inputsCopy[0]); pos++ {
        inputsCopy = return_uncommon(inputsCopy, pos)
        if len(inputsCopy) == 1 {
            co2ScrupperRating = utils.StrByteToInt(inputsCopy[0])
            break
        }
    }

    fmt.Printf("oxygenGeneratorRating: %b %d\n", oxygenGeneratorRating, oxygenGeneratorRating)
    fmt.Printf("co2ScrupperRating: %b %d\n", co2ScrupperRating, co2ScrupperRating)

    return strconv.Itoa(oxygenGeneratorRating*co2ScrupperRating);
}

func return_common(input []string, pos int) []string {
    //fmt.Printf("len input: %d\n", len(input))
    //fmt.Printf("pos: %d\n", pos)
    var one = make([]string, 0)
    var zero = make([]string, 0)
    for line := 0; line < len(input); line++  {
        //fmt.Printf("line: %s\n", input[line])
        if (input[line][pos] == '1') {
            //fmt.Printf("ONE!\n")
            one = append(one, input[line])
        } else {
            //fmt.Printf("ZERO!\n")
            zero = append(zero, input[line])
        }
    }
    //fmt.Printf("ONE: %s\n", one)
    //fmt.Printf("ZERO: %s\n", zero)
    if len(one) >= len(zero) {
        //fmt.Printf("out one len: %d\n", len(one))
        return one
    } else {
        //fmt.Printf("out zero len: %d\n", len(zero))
        return zero
    }
}

func return_uncommon(input []string, pos int) []string {
    //fmt.Printf("len input: %d\n", len(input))
    //fmt.Printf("pos: %d\n", pos)
    var one = make([]string, 0)
    var zero = make([]string, 0)
    for line := 0; line < len(input); line++  {
        //fmt.Printf("line: %s\n", input[line])
        if (input[line][pos] == '1') {
            //fmt.Printf("ONE!\n")
            one = append(one, input[line])
        } else {
            //fmt.Printf("ZERO!\n")
            zero = append(zero, input[line])
        }
    }
    //fmt.Printf("ONE: %s\n", one)
    //fmt.Printf("ZERO: %s\n", zero)
    if len(zero) <= len(one) {
        //fmt.Printf("out zero len: %d\n", len(zero))
        return zero
    } else {
        //fmt.Printf("out one len: %d\n", len(one))
        return one
    }
}
