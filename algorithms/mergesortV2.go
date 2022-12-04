package main

import (
  "fmt"
)

func merge(values []int, left int, middle int, right int) {
  helper := make([]int, len(values))
  for index, _ := range values {
    helper[index] = values[index]
  }

  i := left
  j := middle + 1
  k := left

  for i <= middle && j <= right {
    if helper[i] <= helper[j]{
      values[k] = helper[i]
      i++ 
    }else {
      values[k] = helper[j]
      j++
    }
    k++
  }

  for i <= middle {
    values[k] = helper[i]
    k++
    i++
  }

  for j <= right {
    values[k] = helper[j]
    k++
    j++
  }
}

func mergesort(values []int, left int, right int) {
  if left < right {
    //Calculate the middle of the array range
    middle := (left + right) / 2
    //Handle the left partition
    mergesort(values, left, middle)
    //Handle the right partition
    mergesort(values, middle + 1, right)
    //Merge partitions
    merge(values, left, middle, right)
  }
}

func main() {
  values := []int{9,7,5,3,1,0,8,6,4,2}
  fmt.Println("unsorted: ", values)
  mergesort(values, 0, len(values) -1)
  fmt.Println("sorted: ", values)
}
