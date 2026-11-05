// Package validsudoku solves the "Valid sudoku" problem.
package validsudoku

// Signature records which digits (1-9) have already been seen in a row,
// a column or a box. Position i is true if the digit i+1 was already seen.
type Signature [9]bool

// IsValidSudoku reports whether a 9 x 9 Sudoku board is valid, meaning no
// row, column or 3 x 3 box repeats a digit. Empty cells (".") are ignored.
// It does not check that the board is solvable.
//
// It goes through the board row by row, keeping a Signature for each column
// and for each of the three boxes in the current group of three rows. The
// box signatures are reset every three rows.
//
// Time complexity: O(1), since the board is always 9 x 9 (81 cells). For an
// n x n board it would be O(n^2).
// Space complexity: O(1), since only a fixed number of signatures is stored.
//
// NOTE: This is my solution.
// I first considered the signature as a "map[string]int", but using "[9]bool" is
// more optimal (that was a recommendation from AI).
func IsValidSudoku(board [][]byte) bool {
	columnSigns := make([]Signature, 9)
	boxSigns := make([]Signature, 3)
	for i, row := range board {
		// Checking the row
		ok := addToSignatureAndValidate(&Signature{}, row...)
		if !ok {
			return false
		}

		// Checking the columns
		for j, val := range row {
			ok := addToSignatureAndValidate(&columnSigns[j], val)
			if !ok {
				return false
			}
		}

		// Checking the boxes
		if i > 0 && i%3 == 0 {
			boxSigns = make([]Signature, 3)
		}
		for k := range boxSigns {
			ok := addToSignatureAndValidate(&boxSigns[k], row[k*3:(k+1)*3]...)
			if !ok {
				return false
			}
		}
	}

	return true
}

// addToSignatureAndValidate marks each digit in vals as seen in sign, and
// returns false as soon as one of them was already marked (a repeated digit).
// Empty cells (".") are skipped. sign is a pointer so the marks are kept
// between calls.
func addToSignatureAndValidate(sign *Signature, vals ...byte) bool {
	for _, val := range vals {
		if val == '.' {
			continue
		}
		v := val - '0' // Converting the digit from byte to int
		if sign[v-1] {
			return false
		}
		sign[v-1] = true
	}

	return true
}
