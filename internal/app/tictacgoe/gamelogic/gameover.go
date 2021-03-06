// Functions that scan the board to check if the game's over!
package gamelogic

import (
	"fmt"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"
)

func IsSliceFilledWithSinglePlayerMark(line []string) bool {
	symbolsMap := make(map[string]bool) // map that works as a symbol count!

	// traverses whole slice appending new values to the symbols set
	for _, value := range line {
		symbolsMap[value] = true // if a new symbol is found, then a NEW key is added, otherwise remains as is!

		// slice is NOT homogeneous OR still has digits (not completely filled)
		if len(symbolsMap) > 1 || HasDigits(value) {
			return false
		}
	}

	// slice is homogeneous (symbols set remained as one)
	return true
}

// Evals the board to check for draws and wins/losses -- Naive implementation
func IsGameOverDueToAVictory(validatedFieldPosition int, board *Board) bool {
	row, col := GetBoardRowAndCol(validatedFieldPosition, board)
	relatedRowSlice := GetBoardRowSlice(row, board)
	relatedColumnSlice := GetBoardColumnSlice(col, board)

	// always check for a row or column-based victory
	rowVictory := IsSliceFilledWithSinglePlayerMark(relatedRowSlice)
	columnVictory := IsSliceFilledWithSinglePlayerMark(relatedColumnSlice)

	// not every field position needs to trigger a diagonal/reverse diagonal scan
	diagonalVictory := false
	reverseDiagonalVictory := false

	// if the row and col is on the board's diagonal, check for diagonal victory
	if IsFieldPositionOnBoardDiagonal(row, col) {
		relatedDiagonalSlice := GetBoardDiagonalSlice(board)
		diagonalVictory = IsSliceFilledWithSinglePlayerMark(relatedDiagonalSlice)
	}

	// if the row and col is on the board's REVERSE diagonal, check for REVERSE diagonal victory
	if IsFieldPositionOnBoardReverseDiagonal(row, col, board) {
		relatedReverseDiagonalSlice := GetBoardReverseDiagonalSlice(board)
		reverseDiagonalVictory = IsSliceFilledWithSinglePlayerMark(relatedReverseDiagonalSlice)
	}

	// one of them is enough for a victory!
	return rowVictory || columnVictory || diagonalVictory || reverseDiagonalVictory
}

// Evals board to see if the free Fields are gone, i.e., the game's a tie
func isGameOverDueToDraw(board *Board) bool {
	return len(board.freeFields) == 0
}

// Applies Game Over logic for wins and draws!
func IsGameOver(playerMark string, validatedFieldPosition int, board *Board) bool {
	isGameOver, gameOverMsg := false, ""

	// victories :)!
	if IsGameOverDueToAVictory(validatedFieldPosition, board) {
		winner := ""

		// improve this later!
		if playerMark == settings.Player1Mark {
			winner = "HUMAN"
		} else {
			winner = "COMPUTER"
		}

		isGameOver = true
		gameOverMsg = fmt.Sprintf(settings.VictoryMsg, winner)

		DisplayBoardWithSpaces(board)
		DisplayScreenMessage(gameOverMsg, true)

		return isGameOver
	}

	// ties :/
	if isGameOverDueToDraw(board) {
		isGameOver = true
		gameOverMsg = settings.TieMatchMsg

		DisplayBoardWithSpaces(board)
		DisplayScreenMessage(gameOverMsg, true)

		return isGameOver
	}

	// game's still up!
	return isGameOver
}
