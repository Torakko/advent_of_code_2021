package main;

import (
    "aoc/libs/utils"
    "fmt"
    "sort"
    "strings"
    "strconv"
);

/**
  * Start - 21:29:49
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 10);
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
    `[({(<(())[]>[[{[]{<()<>>
    [(()[<>])]({[<{<<[]>>(
    {([(<{}[<>[]}>{[]{[(<()>
    (((({<>}<{<{<>}{[]{[]{}
    [[<[([]))<([[{}[[()]]]
    [{[{({}]{}}([{[{{{}}([]
    {<[[]]>}<{[{[{[]{()[[[]
    [<(<(<(<{}))><([]([]()
    <{([([[(<>()){}]>(<<{{
    <{([{{}}[<[[[<>{}]]]>[]]`,
};
var part1_test_output = []string{
    `26397`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var score = 0
    for i, line := range(inputs) {
        var lineScore = checkLineCorrupt(line)
        fmt.Printf("Line %v: %v\n", i, lineScore)
        score += lineScore
    }
    return strconv.Itoa(score);
}

func checkLineCorrupt(line string) int{
    var stack []rune
    for _, char := range(line) {
        switch char {
        case '(':
            stack = append(stack, ')')
        case '[':
            stack = append(stack, ']')
        case '{':
            stack = append(stack, '}')
        case '<':
            stack = append(stack, '>')
        default:
            var top = len(stack)-1
            if char == stack[top] {
                //pop
                stack = stack[:top]
            } else {
                // corrupt line
                switch char {
                case ')':
                    return 3
                case ']':
                    return 57
                case '}':
                    return 1197
                case '>':
                    return 25137
                }
            }
        }
    }
    // incomplete line
    return 0
}

func checkLineIncomplete(line string) int{
    var stack []rune
    for _, char := range(line) {
        switch char {
        case '(':
            stack = append(stack, ')')
        case '[':
            stack = append(stack, ']')
        case '{':
            stack = append(stack, '}')
        case '<':
            stack = append(stack, '>')
        default:
            var top = len(stack)-1
            if char == stack[top] {
                //pop
                stack = stack[:top]
            } else {
                // corrupt line
                return 0
            }
        }
    }
    // incomplete line
    var score = 0
    for i := len(stack)-1; i >= 0; i-- {
        score *= 5
        switch stack[i] {
        case ')':
            score += 1
        case ']':
            score += 2
        case '}':
            score += 3
        case '>':
            score += 4
        }
        stack = stack[:i]
    }
    return score
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `288957`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var score []int
    for i, line := range(inputs) {
        var lineScore = checkLineIncomplete(line)
        fmt.Printf("Line %v: %v\n", i, lineScore)
        if lineScore > 0 {
            score = append(score, lineScore)
        }
    }
    sort.Ints(score)
    fmt.Printf("Score: %v\n", score)
    var mid = len(score)/2
    fmt.Printf("Mid: %v\n", score[mid])
    return strconv.Itoa(score[mid]);
}
