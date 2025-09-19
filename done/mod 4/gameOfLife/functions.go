package main

func PlayGameOfLife(initialBoard GameBoard, numGens int) []GameBoard {
	boards := make([]GameBoard, numGens+1)

	boards[0] = initialBoard

	for i := 1; i <= numGens; i++{
		boards[i] = UpdateBoard(boards[i-1])
	}

	return boards
}

func UpdateBoard(currentBoard GameBoard) GameBoard {
	numRows:= CountRows(currentBoard)
	numCols:= CountCols(currentBoard)

	newBoard := initializeBoard(numRows,numCols)

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			newBoard[r][c] = UpdateCell(currentBoard, r, c)
		}
	}
	return newBoard
}

func initializeBoard(numRows, numCols int) GameBoard {
	board := make(GameBoard, numRows)

	for r:=range board {
		board[r] = make([]bool, numCols)
	}

	return board
}

func UpdateCell(currentBoard GameBoard, r, c int) bool {
	numNeighbors := CountLiveNeighbors(currentBoard, r, c)

	if currentBoard[r][c] {
		if numNeighbors == 2 || numNeighbors == 3 {
			return true
		} else {
			return false
		}
	} else {
		if numNeighbors == 3 {
			return true
		} else {
			return false
		}
	}
}

func CountLiveNeighbors(currentboard GameBoard, r, c int) int {
	count := 0

	for i:= r-1; i<=r+1; i++ {
		for j := c-1; j <= c+1; j++ {
			if (i!=r || j!=c) && InField(currentboard, i,j) && currentboard[i][j]{
				count++
			}
		}
	}
	return count
}

func InField(currentboard GameBoard, i, j int) bool {
	numRows:= CountRows(currentboard)
	numCols:= CountCols(currentboard)

	if i<0 || j<0  || i>numRows-1 || j>numCols-1{
		return false
	}
	return true
}
