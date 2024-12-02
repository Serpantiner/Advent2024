package day2

import (
    "bufio"
    "strings"
    "strconv"
    "os"
)

func CheckSafeWithDampener(filePath string) int {
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

        if isSafeWithDampener(nums) {
            safe++
        }
    }

    return safe
}

func isSafeWithDampener(nums []int) bool {
    if isSafe(nums) {
        return true
    }
	
    for i := 0; i < len(nums); i++ {
        newNums := make([]int, len(nums)-1)
        copy(newNums[:i], nums[:i])
        copy(newNums[i:], nums[i+1:])

        if isSafe(newNums) {
            return true
        }
    }

    return false
}