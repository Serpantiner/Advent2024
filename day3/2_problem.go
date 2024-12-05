package day3

import (
    "bufio"
    "os"
    "regexp"
    "strconv"
)

func CalcMulWithControl(path string) int {
    f, _ := os.Open(path)
    defer f.Close()

    total := 0
    enabled := true
    scan := bufio.NewScanner(f)

    pattern := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)

    for scan.Scan() {
        line := scan.Text()
        matches := pattern.FindAllStringSubmatch(line, -1)
        
        for _, match := range matches {
            instruction := match[1]
            
            switch {
            case instruction == "do()":
                enabled = true
            case instruction == "don't()":
                enabled = false
            case enabled && len(match) > 3:  
                x, _ := strconv.Atoi(match[2])
                y, _ := strconv.Atoi(match[3])
                total += x * y
            }
        }
    }
    
    return total
}
