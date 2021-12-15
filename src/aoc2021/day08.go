package main;

import (
    "aoc/libs/utils"
    "fmt"
    "strings"
    "strconv"
);

/**
  * Start - 13:06:17
  * p1 done - {p1Done}
  * p2 done - {p2Done}
  */

func main() {
    var input, _ = utils.Get_input(2021, 8);
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
    `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
    edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
    fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
    fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
    aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
    fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
    dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
    bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
    egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
    gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`,
}
var part1_test_output = []string{
    `26`,
};
func part1(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var count = 0
    for _, line := range(inputs) {
        var split = strings.Split(line, " | ")
        var nums = strings.Split(split[1], " ")
        for _, num := range(nums) {
            switch len(num) {
            case 2: //1
                count++
            case 3: //7
                count++
            case 4: //4
                count++
            case 7: //8
                count++
            }
        }
    }
    return strconv.Itoa(count);
}

var part2_test_input = part1_test_input//[]string{
//    ``,
//};
var part2_test_output = []string{
    `61229`,
};
func part2(input string) string {
    var inputs = utils.Trim_array(strings.Split(strings.Trim(input, separator), separator));
    var total = 0
    for _, line := range(inputs) {
        var split = strings.Split(line, " | ")
        var inputs = strings.Split(split[0], " ")
        var numbers = make([]string, 10)
        var numLen5 []string
        var numLen7 []string
        for _, num := range(inputs) {
            switch len(num) {
            case 2: //1
                numbers[1] = utils.SortStringByCharacter(num)
            case 3: //7
                numbers[7] = utils.SortStringByCharacter(num)
            case 4: //4
                numbers[4] = utils.SortStringByCharacter(num)
            case 5: //2,3,5
                numLen5 = append(numLen5, num)
            case 6: //0,6,9
                numLen7 = append(numLen7, num)
            case 7: //8
                numbers[8] = utils.SortStringByCharacter(num)
            }
        }
        numbers = fixLen7(numLen7, numbers)
        numbers = fixLen5(numLen5, numbers)

        fmt.Printf("Numbers %v\n", numbers)

        var res = 0
        var outputs = strings.Split(split[1], " ")
        fmt.Printf("Outputs %v\n", outputs)
        for i := range(outputs) {
            switch utils.SortStringByCharacter(outputs[3-i]) {
            case numbers[1]:
                res += utils.Pow(10,i)
            case numbers[2]:
                res += utils.Pow(10,i)*2
            case numbers[3]:
                res += utils.Pow(10,i)*3
            case numbers[4]:
                res += utils.Pow(10,i)*4
            case numbers[5]:
                res += utils.Pow(10,i)*5
            case numbers[6]:
                res += utils.Pow(10,i)*6
            case numbers[7]:
                res += utils.Pow(10,i)*7
            case numbers[8]:
                res += utils.Pow(10,i)*8
            case numbers[9]:
                res += utils.Pow(10,i)*9
            }
        }
        total += res
    }
    return strconv.Itoa(total);
}

func fixLen7(numLen7 []string, numbers []string) []string {
    for _, num := range(numLen7) {
        // 1 is not in 6 but not in 0 and 9
        var is6 = false
        for _, char := range(numbers[1]) {
            if !strings.Contains(num, string(char)) {
                is6 = true
                break
            }
        }
        // 4 is in 9 but not in 0 or 6
        var is9 = true
        for _, char := range(numbers[4]) {
            if !strings.Contains(num, string(char)) {
                is9 = false
                break
            }
        }
        if is9 {
            numbers[9] = utils.SortStringByCharacter(num)
        } else if is6 {
            numbers[6] = utils.SortStringByCharacter(num)
        } else {
            numbers[0] = utils.SortStringByCharacter(num)
        }
    }
    return numbers
}

func fixLen5(numLen5 []string, numbers []string) []string {
    for _, num := range(numLen5) {
        // 5 is in 6 but not in 2 or 3
        var is5 = true
        for _, char := range(num) {
            if !strings.Contains(numbers[6], string(char)) {
                is5 = false
                break
            }
        }
        // 1 is in 3 but not in 2 or 5
        var is3 = true
        for _, char := range(numbers[1]) {
            if !strings.Contains(num, string(char)) {
                is3 = false
                break
            }
        }
        if is3 {
            numbers[3] = utils.SortStringByCharacter(num)
        } else if is5 {
            numbers[5] = utils.SortStringByCharacter(num)
        } else {
            numbers[2] = utils.SortStringByCharacter(num)
        }
    }
    return numbers
}
