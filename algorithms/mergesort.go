package main

import (
  "fmt"
)

func merge(values []int, left int, middle int, right int) {
  //copy the values from left to right to a temp array
  temp := make([]int, len(values))
  for index, _ := range values {
    temp[index] = values[index]
  }

  i := left
  j := middle + 1
  k := left

  for i <= middle && j <= right {
    if temp[i] <= temp[j] {
      values[k] = temp[i]
      i++
    } else {
      values[k] = temp[j]
      j++
    }
    k++
  }

  for i <= middle {
    values[k] = temp[i]
    i++
    k++
  }

  for j <= right {
    values[k] = temp[j]
    j++
    k++
  }
}

func mergesort(values []int, left int, right int) { 
  if left < right {
    middle := (right + left) / 2
    mergesort(values, left, middle)
    mergesort(values, middle + 1, right)
    merge(values, left, middle, right)
  }
}

func main() {
  values := []int{9, 1, 3, 2, 5, 4, 7, 6, 88, 32, 54, 65, 21, 98, 90, 67, 0, 19, 8}
  fmt.Println("unsorted: ", values)
  mergesort(values, 0, len(values) - 1);
  fmt.Println("sorted: ", values)
}
