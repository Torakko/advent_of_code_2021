package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 19:33:32
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 02);
    // fmt.Printf("Input: %s \n", input);

    var success = true;
    for i := range part1_test_input {
        if (part1(part1_test_input[i]) != part1_test_output[i]) {
            success = false;
            fmt.Printf("part1 failed with input %s: result %s != expected %s \n",
                    part1_test_input[i],
                    part1(part1_test_input[i]),
                    part1_test_output[i]);
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
            fmt.Printf("part2 failed with input %s: result %s != expected %s \n",
                    part2_test_input[i],
                    part2(part2_test_input[i]),
                    part2_test_output[i]);
            break;
        }
    }
    fmt.Printf("part2 minitest success: %t! \n", success);
    p2 := part2(input);
    fmt.Printf("part2: %s\n", p2);
}

const separator string = "\n";

var part1_test_input = []string{
    `forward 5
down 5
forward 8
up 3
down 8
forward 2`,
};
var part1_test_output = []string{
    `150`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator))
    var depth = 0
    var horizontal = 0
    for i := 0; i < len(inputs); i++  {
        var splitted = strings.Split(inputs[i], " ")
        var direction = splitted[0]
        var distance, _ = strconv.Atoi(splitted[1])
        //fmt.Printf("direction: %s\n", direction)
        //fmt.Printf("distance: %d\n", distance)
        switch direction {
        case "forward":
            horizontal += distance
        case "up":
            depth -= distance
        case "down":
            depth += distance
        }
    }
    //fmt.Printf("depth: %d\n", depth)
    //fmt.Printf("horizontal: %d\n", horizontal)
    return strconv.Itoa(depth*horizontal)
}

var part2_test_input = part1_test_input
var part2_test_output = []string{
    `900`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator))
    var depth = 0
    var horizontal = 0
    var aim = 0
    for i := 0; i < len(inputs); i++  {
        var splitted = strings.Split(inputs[i], " ")
        var direction = splitted[0]
        var distance, _ = strconv.Atoi(splitted[1])
        //fmt.Printf("direction: %s\n", direction)
        //fmt.Printf("distance: %d\n", distance)
        switch direction {
        case "forward":
            depth += aim*distance
            horizontal += distance
        case "up":
            aim -= distance
        case "down":
            aim += distance
        }
        //fmt.Printf("aim: %d\n", aim)
        //fmt.Printf("depth: %d\n", depth)
        //fmt.Printf("horizontal: %d\n", horizontal)
        //fmt.Printf("-----------------------------\n")
    }
    return strconv.Itoa(depth*horizontal)
}
