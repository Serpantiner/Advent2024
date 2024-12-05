package main

import (
    "fmt"
    "advent/day1"
    "advent/day2"
    "advent/day3"
)

func main() {
    path1 := "input.txt"
    path2 := "day2/input_day_2.txt"
    path3 := "day3/input_day_3.txt"

    fmt.Println("Day 1:")
    fmt.Printf("Total distance: %d\n", day1.CalculateDistance(path1))
    fmt.Printf("Similarity score: %d\n", day1.CalculateSimilarityScore(path1))

    fmt.Println("\nDay 2:")
    fmt.Printf("Number of safe reports: %d\n", day2.CheckSafeReports(path2))
    fmt.Printf("Number of safe reports with dampener: %d\n", day2.CheckSafeWithDampener(path2))

    fmt.Println("\nDay 3:")
    fmt.Printf("Sum of multiplications: %d\n", day3.CalcMul(path3))
    fmt.Printf("Sum of enabled multiplications: %d\n", day3.CalcMulWithControl(path3))
}
