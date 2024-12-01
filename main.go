package main

import (
    "fmt"
    "advent/day1"
)

func main() {
    path := "input.txt"
    fmt.Println(day1.CalculateDistance(path))
    fmt.Println(day1.CalculateSimilarityScore(path))
}