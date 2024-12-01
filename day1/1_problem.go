package day1

import (
    "bufio"
    "math"
    "os"
    "sort"
    "strings"
    "strconv"
)

func CalculateDistance(filePath string) int {
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

    sort.Ints(leftList)
    sort.Ints(rightList)

    total := 0
    for i := 0; i < len(leftList); i++ {
        total += int(math.Abs(float64(leftList[i] - rightList[i])))
    }

    return total
}
