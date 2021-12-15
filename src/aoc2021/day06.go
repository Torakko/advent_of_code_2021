package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 10:43:48
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 06);
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
    `3,4,3,1,2`,
};
var part1_test_output = []string{
    `5934`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var fish0 = 0
    var fish1 = 0
    var fish2 = 0
    var fish3 = 0
    var fish4 = 0
    var fish5 = 0
    var fish6 = 0
    var fish7 = 0
    var fish8 = 0
    var vals = strings.Split(inputs[0], ",")
    for _, rawAge := range vals {
        var age, _ = strconv.Atoi(rawAge)
        switch age {
        case 0:
            fish0++
        case 1:
            fish1++
        case 2:
            fish2++
        case 3:
            fish3++
        case 4:
            fish4++
        case 5:
            fish5++
        }
    }
    for day := 1; day<=80; day++ {
        var newFish = fish0
        fish0 = fish1
        fish1 = fish2
        fish2 = fish3
        fish3 = fish4
        fish4 = fish5
        fish5 = fish6
        fish6 = fish7+newFish
        fish7 = fish8
        fish8 = newFish
    }
    var fish = fish0+fish1+fish2+fish3+fish4+fish5+fish6+fish7+fish8
    return strconv.Itoa(fish);
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `26984457539`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var fish0 = 0
    var fish1 = 0
    var fish2 = 0
    var fish3 = 0
    var fish4 = 0
    var fish5 = 0
    var fish6 = 0
    var fish7 = 0
    var fish8 = 0
    var vals = strings.Split(inputs[0], ",")
    for _, rawAge := range vals {
        var age, _ = strconv.Atoi(rawAge)
        switch age {
        case 0:
            fish0++
        case 1:
            fish1++
        case 2:
            fish2++
        case 3:
            fish3++
        case 4:
            fish4++
        case 5:
            fish5++
        }
    }
    for day := 1; day<=256; day++ {
        var newFish = fish0
        fish0 = fish1
        fish1 = fish2
        fish2 = fish3
        fish3 = fish4
        fish4 = fish5
        fish5 = fish6
        fish6 = fish7+newFish
        fish7 = fish8
        fish8 = newFish
    }
    var fish = fish0+fish1+fish2+fish3+fish4+fish5+fish6+fish7+fish8
    return strconv.Itoa(fish);
}
