package main

import (
  "fmt"
)

func swap(values []int, from int, to int) {
  temp := values[to]
  values[to] = values[from]
  values[from] = temp
}

func partition(values []int, left int, right int) int {
  pivot := values[left]

  i := left

  for j := (left + 1); j <= right; j++ {
    if values[j] < pivot {
      i++
      swap(values, j, i)
    }
  }

  swap(values, left, i)

  return i
}

func quicksort(values []int, left int, right int) {
  if left < right {
    pivotIndex := partition(values, left, right)
    quicksort(values, left, pivotIndex - 1)
    quicksort(values, pivotIndex + 1, right)
  }
}

func main() {
  values := []int{7,1,4,9,33,6,2,45,56,0,78,32,21}
  fmt.Println("unsorted: ", values)
  quicksort(values, 0, len(values) - 1)
  fmt.Println("sorted: ", values)
}
