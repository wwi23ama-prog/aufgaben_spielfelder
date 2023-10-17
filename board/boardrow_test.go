package board

import "fmt"

func ExampleMakeBoardRow() {
	row := MakeBoardRow(3, "A")
	for _, v := range row {
		fmt.Println(v)
	}
	// Output:
	// A
	// A
	// A
}

func ExampleMakeEmptyBoardRow() {
	row := MakeEmptyBoardRow(3)
	for _, v := range row {
		fmt.Println(v)
	}
	// Output:
	//
	//
	//
}

func ExampleRowContainsOnly() {
	row1 := MakeBoardRow(3, "A")
	fmt.Println(RowContainsOnly(row1, "A"))
	fmt.Println(RowContainsOnly(row1, "B"))

	row2 := MakeEmptyBoardRow(3)
	fmt.Println(RowContainsOnly(row2, " "))
	row2[1] = "A"
	fmt.Println(RowContainsOnly(row2, " "))

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIsEmpty() {
	row1 := MakeBoardRow(3, "A")
	fmt.Println(IsEmpty(row1))

	row2 := MakeEmptyBoardRow(3)
	fmt.Println(IsEmpty(row2))

	// Output:
	// false
	// true
}
