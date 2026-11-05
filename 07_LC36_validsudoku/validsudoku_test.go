package validsudoku

import (
	"slices"
	"testing"
)

// newBoard builds a board from 9 strings of 9 characters each, where '.'
// is an empty cell. Each character becomes a byte cell (e.g. '5' or '.').
// It makes the test cases readable as a grid.
func newBoard(t *testing.T, rows [9]string) [][]byte {
	t.Helper()

	board := make([][]byte, 9)
	for i, row := range rows {
		if len(row) != 9 {
			t.Fatalf("test setup error: row %d has length %d, want 9", i, len(row))
		}
		board[i] = []byte(row)
	}

	return board
}

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		name  string
		board [9]string
		want  bool
	}{
		{
			"leetcode example 1",
			[9]string{
				"53..7....",
				"6..195...",
				".98....6.",
				"8...6...3",
				"4..8.3..1",
				"7...2...6",
				".6....28.",
				"...419..5",
				"....8..79",
			},
			true,
		},
		{
			"leetcode example 2: repeated 8 in the top left box",
			[9]string{
				"83..7....",
				"6..195...",
				".98....6.",
				"8...6...3",
				"4..8.3..1",
				"7...2...6",
				".6....28.",
				"...419..5",
				"....8..79",
			},
			false,
		},
		{
			"empty board",
			[9]string{
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			},
			true,
		},
		{
			"completely filled valid board",
			[9]string{
				"534678912",
				"672195348",
				"198342567",
				"859761423",
				"426853791",
				"713924856",
				"961537284",
				"287419635",
				"345286179",
			},
			true,
		},
		{
			"completely filled board with one wrong cell",
			[9]string{
				"534678913",
				"672195348",
				"198342567",
				"859761423",
				"426853791",
				"713924856",
				"961537284",
				"287419635",
				"345286179",
			},
			false,
		},
		{
			"repeated digit only in a row (different columns and boxes)",
			[9]string{
				"5...5....",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			},
			false,
		},
		{
			"repeated digit only in a column (different rows and boxes)",
			[9]string{
				"5........",
				".........",
				".........",
				"5........",
				".........",
				".........",
				".........",
				".........",
				".........",
			},
			false,
		},
		{
			"repeated digit only in a box (different rows and columns)",
			[9]string{
				"5........",
				".5.......",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			},
			false,
		},
		{
			"repeated digit in the last row",
			[9]string{
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				"1.......1",
			},
			false,
		},
		{
			"repeated digit in the last column",
			[9]string{
				"........7",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				"........7",
			},
			false,
		},
		{
			"repeated digit in the bottom right box",
			[9]string{
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				"......9..",
				".........",
				"........9",
			},
			false,
		},
		{
			"same digit in different rows, columns and boxes",
			[9]string{
				"5........",
				"...5.....",
				"......5..",
				".5.......",
				"....5....",
				".......5.",
				"..5......",
				".....5...",
				"........5",
			},
			true,
		},
		{
			"same digit in diagonal boxes only",
			[9]string{
				"..3......",
				".........",
				".........",
				"...3.....",
				".........",
				".........",
				".........",
				".........",
				".........",
			},
			true,
		},
		{
			"valid but not solvable (cell (0,8) cannot take any digit)",
			[9]string{
				"12345678.",
				"........9",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := newBoard(t, tt.board)
			original := make([][]byte, len(board))
			for i := range board {
				original[i] = slices.Clone(board[i])
			}

			got := IsValidSudoku(board)

			if got != tt.want {
				t.Errorf("IsValidSudoku(%v) = %v, want %v", tt.board, got, tt.want)
			}
			for i := range board {
				if !slices.Equal(board[i], original[i]) {
					t.Errorf("IsValidSudoku modified the board: row %d is %q, want %q", i, board[i], original[i])
				}
			}
		})
	}
}
