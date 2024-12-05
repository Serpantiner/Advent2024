package day3

import (
    "bufio"
    "os"
    "regexp"
    "strconv"
)

func CalcMul(path string) int {
    f, _ := os.Open(path)
    defer f.Close()

    total := 0
    scan := bufio.NewScanner(f)
    pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

    for scan.Scan() {
        line := scan.Text()
        nums := pattern.FindAllStringSubmatch(line, -1)
    
        for _, n := range nums {
            x, _ := strconv.Atoi(n[1])
            y, _ := strconv.Atoi(n[2])
            total += x * y
        }
    }

    return total
}
