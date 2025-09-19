package main

//place your functions from the assignment here.

// func DiffuseBoardOneParticle_i(currentBoard [][]float64, kernel [3][3]float64) [][]float64 {
//     rows := len(currentBoard)
//     cols := len(currentBoard[0])

//     newBoard := make([][]float64, rows)
//     for i := range newBoard {
//         newBoard[i] = make([]float64, cols)
//     }

//     // update each cell with diffusion effect
//     for r := 0; r < rows; r++ {
//         for c := 0; c < cols; c++ {
//             newBoard[r][c] = currentBoard[r][c] + ApplyKernel(currentBoard, kernel, r, c)
//         }
//     }

//     return newBoard
// }


// func InField(board [][]float64, row, col int) bool {
//     return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
// }

// func ApplyKernel(board [][]float64, kernel [3][3]float64, row, col int) float64 {
//     sum := 0.0
//     for kr := -1; kr <= 1; kr++ {
//         for kc := -1; kc <= 1; kc++ {
//             rr := row + kr
//             cc := col + kc

//             if InField(board, rr, cc) {
//                 sum += kernel[kr+1][kc+1]*board[rr][cc]
//             }
//         }
//     }
//     return sum
// }

// DiffuseBoardOneParticle applies one diffusion step using kernel * diffusionRate
// func DiffuseBoardOneParticle(currentBoard [][]float64, kernel [3][3]float64) [][]float64 {
//     rows := len(currentBoard)
//     cols := len(currentBoard[0])

//     newBoard := make([][]float64, rows)
//     for i := range newBoard {
//         newBoard[i] = make([]float64, cols)
//     }

//     // update each cell with diffusion effect
//     for r := 0; r < rows; r++ {
//         for c := 0; c < cols; c++ {
//             newBoard[r][c] = currentBoard[r][c] + ApplyKernel(currentBoard, kernel, r, c)
//         }
//     }

//     return newBoard
// }


func SumCells(cells ...Cell) Cell {
    var s Cell
	for _, c := range cells {
        s[0] += c[0]
        s[1] += c[1]
	}
	return s
}

func ChangeDueToReactions(currentCell Cell, feedRate, killRate float64) Cell {
	a := currentCell[0]
	b := currentCell[1]
    
    var c Cell
	// Gray-Scott equations
	c[0] = -a*b*b + feedRate*(1-a)
	c[1] = a*b*b - (killRate)*b
    
    
    return c
}

func ChangeDueToDiffusion(currentBoard Board, row, col int, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {
    var c Cell
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if InField(currentBoard, row+i, col+j) {
                c[0] += kernel[i+1][j+1]*currentBoard[i+row][j+col][0]*preyDiffusionRate
                c[1] += kernel[i+1][j+1]*currentBoard[i+row][j+col][1]*predatorDiffusionRate
            }
        }
    }
    return c
}

func InField(board Board, row, col int) bool {
    return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}

func UpdateBoard(currentBoard Board, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Board {
    numRows := len(currentBoard)
    numCols := len(currentBoard[0])

    newBoard := make(Board, numRows)
    for row := range newBoard {
        newBoard[row] = make([]Cell, numCols)
    }

    for row := 0; row < numRows; row++ {
        for col :=0; col < numCols; col++ {
            newBoard[row][col] = UpdateCell(currentBoard, row, col, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
        }
    }
    return newBoard
}

func UpdateCell(currentBoard Board, row, col int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {
    currentCell := currentBoard[row][col]
    diffusionValues := ChangeDueToDiffusion(currentBoard, row, col, preyDiffusionRate, predatorDiffusionRate, kernel)
    reactionValues := ChangeDueToReactions(currentCell, feedRate, killRate)
    return SumCells(currentCell, diffusionValues, reactionValues)
}

func SimulateGrayScott(initialBoard Board, numGens int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) []Board {
    boards := make([]Board, numGens+1)
    boards[0] = initialBoard

    for i := 1; i <= numGens;  i++ {
        boards[i] = UpdateBoard(boards[i-1], feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
    }
    return boards
}

func InitializeBoard(numRows, numCols int) Board {
	b := make(Board, numRows)
	for i := range b {
		b[i] = make([]Cell, numCols)
	}

	return b
}