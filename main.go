package main

import (
   "fmt"
   "advent/day1"
   "advent/day2" 
)

func main() {
   path1 := "input.txt"
   path2 := "day2/input_day_2.txt"

   fmt.Println("Day 1:")
   fmt.Printf("Total distance: %d\n", day1.CalculateDistance(path1))
   fmt.Printf("Similarity score: %d\n", day1.CalculateSimilarityScore(path1))

   fmt.Println("\nDay 2:")
   fmt.Printf("Number of safe reports: %d\n", day2.CheckSafeReports(path2))
   fmt.Printf("Number of safe reports with dampener: %d\n", day2.CheckSafeWithDampener(path2))
}
