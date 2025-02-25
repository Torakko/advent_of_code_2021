package utils;

import (
  "fmt"
  "os"
  "errors"
  "sort"
  "strings"
  "strconv"
  "net/http"
  "io/ioutil"
  "math"
);

const url_format string = "https://adventofcode.com/%d/day/%d/input";
const path_session_token = "cookie_session";

func get_session() (string, error) {
    content, err := ioutil.ReadFile(path_session_token)
    if err != nil {
        fmt.Fprintln(os.Stderr, "ERROR (utils.get_session): Failed to read ./"+path_session_token);
        return "", err
    }

    return strings.Trim(string(content), " \n"), nil;
}

func Get_input(year int, day int) (string, error) {
    var session, err = get_session();

    // Sanity and errors
    if err != nil {
        return "", err;
    }
    if session == "" {
        const msg string = "ERROR (utils.Get_input): Session-token is an empty string. Make sure you edited ./cookie_session";
        fmt.Fprintln(os.Stderr, msg);
        return "", errors.New(msg);
    }

    var url = fmt.Sprintf(url_format, year, day);

    // Declare http client
    var client = &http.Client{};

    // Declare HTTP Method and Url
    req, err := http.NewRequest("GET", url, nil);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }   
    
    // Set cookie
    req.Header.Set("Cookie", fmt.Sprintf("session=%s; count=x", session));

    resp, err := client.Do(req);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }
    // Read response
    data, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    return string(data), nil;
}

func Trim_array(strs []string) []string {
    for i, str := range strs {
        strs[i] = strings.Trim(str, " ");
    }
    return strs;
}

func StrToInt_array(strs []string) ([]int, error) {
    var ints = make([]int, len(strs));
    for i, str := range strs {
        var d, err = strconv.Atoi(str);
        if err != nil {
            return nil, err;
        }
        ints[i] = d;
    }
    return ints, nil;
}

func StrToFloat_array(strs []string) ([]float64, error) {
    var floats = make([]float64,len(strs));
    for i, str := range strs {
        var f, err = strconv.ParseFloat(str, 64);
        if err != nil {
            return nil, err;
        }
        floats[i] = f;
    }
    return floats, nil;
}

func StrByteToInt(input string) int {
    var res = 0
    for pos := 0; pos < len(input); pos++ {
        if (input[len(input)-pos-1] == '1') {
            res += int(math.Pow(2,float64(pos)))
        }
    }
    return res
}

func RemoveFromSlice(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func PrintMatrix(matrix [][]int) {
    for _, line := range matrix {
        fmt.Printf("%v\n", line)
    }
}

func Pow(a int, exp int) int {
    return int(math.Pow(float64(a), float64(exp)))
}

func StringToRuneSlice(s string) []rune {
    var r []rune
    for _, runeValue := range s {
        r = append(r, runeValue)
    }
    return r
}

func SortStringByCharacter(s string) string {
    r := StringToRuneSlice(s)
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
    return string(r)
}

func Max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func Min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func MinArray(a []int) int {
    var res = math.MaxInt32
    for _, x := range(a) {
        res = Min(res, x)
    }
    return res
}

func Abs(a int) int {
    if a > 0 {
        return a
    } else {
        return -a
    }
}

func SumArray(a []int) int {
    var res = 0
    for _, val := range(a) {
        res += val
    }
    return res
}
