package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 12:28:46
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 07);
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
    `16,1,2,0,4,2,7,1,2,14`,
};
var part1_test_output = []string{
    `37`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var vals = strings.Split(inputs[0], ",")
    var crabs []int
    var max = 0
    for _, rawAge := range vals {
        var pos, _ = strconv.Atoi(rawAge)
        crabs = append(crabs, pos)
        max = utils.Max(max, pos)
    }
    var fuel = make([]int, max+1)
    for pos := 0; pos<max+1; pos++ {
        var sum = 0
        for _, crab := range crabs {
            sum += utils.Abs(pos-crab)
        }
        fuel[pos] = sum
    }
    //fmt.Printf("Fuel %v\n", fuel)
    var result = utils.MinArray(fuel)
    return strconv.Itoa(result);
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `168`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var vals = strings.Split(inputs[0], ",")
    var crabs []int
    var max = 0
    for _, rawAge := range vals {
        var pos, _ = strconv.Atoi(rawAge)
        crabs = append(crabs, pos)
        max = utils.Max(max, pos)
    }

    var cost = make([]int, max+1)
    cost[0] = 0
    cost[1] = 1
    for i := 2; i<max+1; i++ {
        cost[i] = cost[i-1] + i
    }
    //fmt.Printf("Cost %v\n", cost)

    var fuel = make([]int, max+1)
    for pos := 0; pos<max+1; pos++ {
        var sum = 0
        for _, crab := range crabs {
            sum += cost[utils.Abs(pos-crab)]
        }
        fuel[pos] = sum
    }
    //fmt.Printf("Fuel %v\n", fuel)
    var result = utils.MinArray(fuel)
    return strconv.Itoa(result);
}
