package main

import (
  "fmt"
  "sort"
)

func medianPos(values []int, left int, right int) int {
  medianIndex := (left + right) / 2
  sortedNum := []int{values[left], values[right], values[medianIndex]}
  sort.Ints(sortedNum);
  if sortedNum[1] == values[left]{
    return left
  }else if sortedNum[1] == values[right]{
    return right
  }else{
    return medianIndex
  }
}

func swap(values []int, from int, to int) {
  temp := values[to]
  values[to] = values[from]
  values[from] = temp
}

func partition(values []int, left int, right int) int {
  //Get median index
  medianIndex := medianPos(values, left, right)
  pivotValue := values[medianIndex];

  //Control arrange index
  i := left

  //swap pivot to the beginning of the range
  swap(values, medianIndex, left)

  for index := left + 1; index <= right; index++ {
    if values[index] <= pivotValue {
      i++
      values[i] = values[index]
    }
  }

  swap(values, left, i)

  return i

}

func quickSort(values []int, left int, right int) {
  if left < right {
    //Arrange pivot in array and get its index
    pivotIndex := partition(values, left, right)
    //Handle the left partition
    quickSort(values, left, pivotIndex)
    //Handle the right partition
    quickSort(values, pivotIndex + 1, right)
  }
}

func main() {
  values := []int{9,7,5,3,1,0,8,6,4,2}
  fmt.Println("unsorted: ", values)
  quickSort(values, 0, len(values) -1);
  fmt.Println("sorted: ", values)
}
