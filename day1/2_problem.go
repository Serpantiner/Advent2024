package day1

import (
    "bufio"
    "os"
    "strings"
    "strconv"
)

func CalculateSimilarityScore(filePath string) int {
    file, _ := os.Open(filePath)
    defer file.Close()

    leftList := []int{}
    rightList := []int{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.Fields(scanner.Text())
        left, _ := strconv.Atoi(line[0])
        right, _ := strconv.Atoi(line[1])
        leftList = append(leftList, left)
        rightList = append(rightList, right)
    }

    total := 0
    for _, left := range leftList {
        count := 0
        for _, right := range rightList {
            if left == right {
                count++
            }
        }
        total += left * count
    }

    return total
}