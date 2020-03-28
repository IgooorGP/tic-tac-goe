package tictacgoe

import (
	"reflect"
	"testing"
)

// Asserts the creation of an initialized game board
func TestCreateGameBoardShouldCreateInitializeBoardFields3x3(t *testing.T) {
	boardSize := 4
	expectedBoard := [][]string{
		{"01", "02", "03", "04"},
		{"05", "06", "07", "08"},
		{"09", "10", "11", "12"},
		{"13", "14", "15", "16"},
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

func TestPrivateGetRowStringBoard5x5(t *testing.T) {
	gameBoard := CreateGameBoard(4)
	expectedFirstRow := " 01 | 02 | 03 | 04 "
	expectedLastRow := " 13 | 14 | 15 | 16 "

	firstRow := getRowString(gameBoard, 0)
	lastRow := getRowString(gameBoard, len(gameBoard.fields)-1)

	if firstRow != expectedFirstRow {
		t.Errorf("First row does not match. Actual: %s, expected: %s", firstRow, expectedFirstRow)
	}

	if lastRow != expectedLastRow {
		t.Errorf("Last row does not match. Actual: %s, expected: %s", lastRow, expectedLastRow)
	}
}