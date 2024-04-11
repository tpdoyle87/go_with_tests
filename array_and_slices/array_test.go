package main

import (
// 	"fmt"
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
}
