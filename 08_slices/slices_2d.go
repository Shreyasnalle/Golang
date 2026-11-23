package main

import (
	"fmt"
	"strings"
)

// slice_2d_literal -> a slice of slices, every row is its own []string
// this is the tic tac toe board example, the 2d part i had skipped in slices.go
func slice_2d_literal() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	for _, row := range board {
		fmt.Println(strings.Join(row, " ")) // Join glues the row elements with a space
	}
}

// slice_2d_make -> building a rows x cols grid with make, first the outer slice, then every inner row
// rows can even have different lengths (jagged), something a 2d array can never do
func slice_2d_make(rows, cols int) [][]int {
	grid := make([][]int, rows) // outer slice with `rows` nil rows
	for i := range grid {
		grid[i] = make([]int, cols) // give every row its own slice
		for j := range grid[i] {
			grid[i][j] = i * j
		}
	}
	return grid
}

// slice_3d -> just one more level, a slice of slices of slices
func slice_3d() {
	cube := [][][]int{
		{{1, 2}, {3, 4}},
		{{5, 6}, {7, 8}},
	}
	fmt.Println("3d:", cube, "| cube[1][0][1] =", cube[1][0][1]) // 6
}

// slice_of_struct_2d -> using a struct inside the 2d slice, so each cell can hold different data types
func slice_of_struct_2d() {
	type cell struct {
		value int
		label string
	}
	table := [][]cell{
		{{1, "a"}, {2, "b"}},
		{{3, "c"}, {4, "d"}},
	}
	fmt.Println("2d slice of struct:", table)
}
