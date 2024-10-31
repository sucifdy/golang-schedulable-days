package main

import (
    "fmt"
)

func SchedulableDays(date1 []int, date2 []int) []int {
    // Menggunakan map untuk menyimpan tanggal kosong dari date1
    dateMap := make(map[int]bool)
    result := []int{}

    // Menyimpan tanggal kosong date1 ke dalam map
    for _, date := range date1 {
        dateMap[date] = true
    }

    // Mencari tanggal kosong yang ada di date2 dan juga ada di map
    for _, date := range date2 {
        if dateMap[date] {
            result = append(result, date) // Tambahkan tanggal ke hasil
        }
    }

    return result
}

func main() {
    // Contoh test case
    fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5})) // Output: [3, 4]
    fmt.Println(SchedulableDays([]int{11, 12, 13, 14, 15}, []int{5, 10, 12, 13, 20, 21})) // Output: [12, 13]
    fmt.Println(SchedulableDays([]int{2, 7, 12, 20, 21, 22}, []int{1, 3, 6, 10})) // Output: []
    fmt.Println(SchedulableDays([]int{}, []int{})) // Output: []
}
