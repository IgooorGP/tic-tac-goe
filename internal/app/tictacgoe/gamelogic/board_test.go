package gamelogic

import (
	"reflect"
	"testing"
)

// Asserts the creation of an initialized game board
func TestCreateGameBoardShouldCreateInitializeBoardFields3x3(t *testing.T) {
	boardSize := 4
	expectedBoard := [][]string{
		{"00", "01", "02", "03"},
		{"04", "05", "06", "07"},
		{"08", "09", "10", "11"},
		{"12", "13", "14", "15"},
	}

	// function invocation
	boardFields := CreateGameBoard(boardSize).fields

	// board assertions for each row
	for i := range expectedBoard {
		// compares each board row with the expected board
		if !reflect.DeepEqual(boardFields[i], expectedBoard[i]) {
			t.Errorf("Board row %d has failed assertion! board: %s, expected: %s", i, boardFields[i], expectedBoard[i])
		}
	}
}

// Asserts appropriate row printing
func TestPrivateGetRowStringBoard5x5(t *testing.T) {
	gameBoard := CreateGameBoard(4)
	expectedFirstRow := " 00 | 01 | 02 | 03 "
	expectedLastRow := " 12 | 13 | 14 | 15 "

	firstRow := getRowString(gameBoard, 0)
	lastRow := getRowString(gameBoard, len(gameBoard.fields)-1)

	if firstRow != expectedFirstRow {
		t.Errorf("First row does not match. Actual: %s, expected: %s", firstRow, expectedFirstRow)
	}

	if lastRow != expectedLastRow {
		t.Errorf("Last row does not match. Actual: %s, expected: %s", lastRow, expectedLastRow)
	}
}

// Asserts proper conversion of a field number to a row, col of the board matrix
// TODO: Improve this test with more clarity
func TestGetBoardRowAndCol5x5(t *testing.T) {
	board := CreateGameBoard(5)

	row, col := GetBoardRowAndCol(0, board)
	expectedRow, expectedCol := 0, 0

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(13, board)
	expectedRow, expectedCol = 2, 3

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(15, board)
	expectedRow, expectedCol = 3, 0

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(6, board)
	expectedRow, expectedCol = 1, 1

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(9, board)
	expectedRow, expectedCol = 1, 4

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(24, board)
	expectedRow, expectedCol = 4, 4

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

}

// Asserts GetFieldValue returns the appropriate value from the board
func TestGetFieldValue(t *testing.T) {
	gameBoard := CreateGameBoard(3)

	gameBoard.fields[0][1] = "IG"

	value1 := GetFieldValue(0, gameBoard)
	value2 := GetFieldValue(1, gameBoard)

	if value1 != "00" {
		t.Errorf("Unexpected value returned from board. actual: %s, expected: %s", value1, "00")
	}

	if value2 != "IG" {
		t.Errorf("Unexpected value returned from board. actual: %s, expected: %s", value1, "IG")
	}

}

// Asserts only the desired board row is returned based on a row index
func TestGetBoardRowSlice(t *testing.T) {
	gameBoard := CreateGameBoard(5)
	expectedRowSlice := []string{"20", "21", "22", "23", "24"}

	rowSlice := GetBoardRowSlice(4, gameBoard)

	for i, val := range rowSlice {
		if val != expectedRowSlice[i] {
			t.Errorf("Unexpected element at row slice. actual: %s, expected: %s, index: %d", val, expectedRowSlice[i], i)
		}
	}
}

// Asserts only the desired board column is returned based on a col index
func TestGetBoardColumnSlice(t *testing.T) {
	gameBoard := CreateGameBoard(5)
	expectedColumnSlice := []string{"04", "09", "14", "19", "24"} // last column
	expectedColumnSlice2 := []string{"02", "07", "12", "17", "22"}

	columnSlice := GetBoardColumnSlice(4, gameBoard)

	for i, val := range columnSlice {
		if val != expectedColumnSlice[i] {
			t.Errorf("Unexpected element at column slice. actual: %s, expected: %s, index: %d", val, expectedColumnSlice[i], i)
		}
	}

	columnSlice2 := GetBoardColumnSlice(2, gameBoard)

	for i, val := range columnSlice2 {
		if val != expectedColumnSlice2[i] {
			t.Errorf("Unexpected element at column slice. actual: %s, expected: %s, index: %d", val, expectedColumnSlice2[i], i)
		}
	}
}

// Asserts only the desired board diagonal elements are returned
func TestGetBoardDiagonalSlice(t *testing.T) {
	gameBoard := CreateGameBoard(4)
	gameBoard.fields[1][1] = "IG" // has a player mark

	expectedDiagonalSlice := []string{"00", "IG", "10", "15"}

	diagonalSlice := GetBoardDiagonalSlice(gameBoard)

	for i, val := range diagonalSlice {
		if val != expectedDiagonalSlice[i] {
			t.Errorf("Unexpected element at diagonal slice. actual: %s, expected: %s, index: %d", val, expectedDiagonalSlice[i], i)
		}
	}
}

// Asserts only the desired board reverse diagonal elements are returned on a 3x3 board
func TestGetBoardReverseDiagonalSlice3x3(t *testing.T) {
	gameBoard := CreateGameBoard(3)
	gameBoard.fields[1][1] = "IG" // has a player mark

	expectedReverseDiagonalSlice := []string{"02", "IG", "06"}

	reverseDiagonalSlice := GetBoardReverseDiagonalSlice(gameBoard)

	for i, val := range reverseDiagonalSlice {
		if val != expectedReverseDiagonalSlice[i] {
			t.Errorf("Unexpected element at diagonal slice. actual: %s, expected: %s, index: %d", val, expectedReverseDiagonalSlice[i], i)
		}
	}
}

// Asserts only the desired board reverse diagonal elements are returned on a 6x6 board
func TestGetBoardReverseDiagonalSlice5x5(t *testing.T) {
	gameBoard := CreateGameBoard(6)
	gameBoard.fields[2][3] = "IG" // has a player mark

	expectedReverseDiagonalSlice := []string{"05", "10", "IG", "20", "25", "30"} // expected reverse diagonal of 6x6 board

	reverseDiagonalSlice := GetBoardReverseDiagonalSlice(gameBoard)

	for i, val := range reverseDiagonalSlice {
		if val != expectedReverseDiagonalSlice[i] {
			t.Errorf("Unexpected element at diagonal slice. actual: %s, expected: %s, index: %d", val, expectedReverseDiagonalSlice[i], i)
		}
	}
}
