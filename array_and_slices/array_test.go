package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
    t.Run("Running over an array of 3 ints", func(t *testing.T) {
        numbers := []int{2, 4, 8}
        total := Sum(numbers)
        expected := 14

        if total != expected {
            t.Errorf("expected %d but got %d given %v", expected, total, numbers)
        }
    })

    t.Run("total two arrays together at respective index", func(t *testing.T) {
        summed := SumAll([]int{1,2}, []int{0,9})
        expect := []int{3,9}

        if !slices.Equal(summed, expect) {
            t.Errorf("Expected %v but got %v", expect, summed)
        }
    })

    t.Run("Running the SumAllTails function", func(t *testing.T) {
        _, tails := SumAllTails([]int{1,2,3}, []int{1,2,3})
        expect := []int{5, 5}

        if !slices.Equal(tails, expect) {
            t.Errorf("Expected %v but got %v", expect, tails)
        }
    })

    t.Run("catching nil slices", func(t *testing.T) {
        err, _ := SumAllTails([]int{}, []int{1,2,4})
        expect := "slice can't be empty"

        if err != expect {
            t.Errorf("Expected empty slice to be caught but it was't")
        }
    })
}
