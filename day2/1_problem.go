package day2

import (
    "bufio"
    "strings"
    "strconv"
    "os"
)

func CheckSafeReports(filePath string) int {

    file, _ := os.Open(filePath)
    defer file.Close()

    safe := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        numbers := strings.Fields(scanner.Text())
        nums := make([]int, len(numbers))
        for i := 0; i < len(numbers); i++ {
            nums[i], _ = strconv.Atoi(numbers[i])
        }
        if isSafe(nums) {
            safe++
        }
    }

    return safe
}

func isSafe(nums []int) bool {

    if len(nums) < 2 {
        return false
    }

    goingUp := nums[1] > nums[0]

    for i := 1; i < len(nums); i++ {
        diff := nums[i] - nums[i-1]

        if diff == 0 || diff > 3 || diff < -3 {
            return false
        }

        if goingUp && diff < 0 {
            return false
        }
        if !goingUp && diff > 0 {
            return false
        }
    }

    return true
}
