package main

import (
	"strconv"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 1000} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLength := size
			if wantLength < 0 {
				wantLength = 0
			}
			if len(data) != wantLength {
				t.Fatalf("length = %d, want %d", len(data), wantLength)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want positive", i, value)
				}
			}
		})
	}
}

func TestMaximumFunctions(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"zero", []int{0}, 0},
		{"first", []int{9, 2, 1}, 9},
		{"middle", []int{1, 9, 2}, 9},
		{"last", []int{1, 2, 9}, 9},
		{"equal", []int{7, 7, 7}, 7},
		{"negative", []int{-9, -2, -3}, -2},
		{"eight", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
		{"negative chunks", []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
	}
	for name, findMaximum := range map[string]func([]int) int{
		"maximum":   maximum,
		"maxChunks": maxChunks,
	} {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					if got := findMaximum(test.data); got != test.want {
						t.Errorf("got %d, want %d", got, test.want)
					}
				})
			}
		})
	}
}
