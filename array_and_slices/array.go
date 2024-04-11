package main

import "fmt"

func Sum(i []int) (total int) {
	for _, n := range i {
		total += n
	}
	return
}

func SumAll(numsToSum ...[]int) []int {
    var sums []int
    for _, nums := range numsToSum {
        sums = append(sums, Sum(nums))
    }
    return sums
}

func SumAllTails(tailsToSum ...[]int) (string, []int) {
    var sums []int
    for _, nums := range tailsToSum {
        if len(nums) == 0 {
            return errorNilSlice(), nil
        }
        sums = append(sums, Sum(nums[1:]))
    }
    return "", sums
}

func errorNilSlice() string {
    return fmt.Sprintf("slice can't be empty")
}