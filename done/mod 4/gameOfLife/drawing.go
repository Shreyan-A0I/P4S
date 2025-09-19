package main

import (
	"canvas"
	"image"	
)

func DrawGameBoard(board GameBoard, cellwidth int) image.Image {
	width := CountCols(board) * cellwidth
	height := CountRows(board) * cellwidth
	c := canvas.CreateNewCanvas(width,height)

	darkgrey := canvas.MakeColor(60,60,60)
	white := canvas.MakeColor(255,255,255)

	c.SetFillColor(darkgrey)
	c.Clear()

	for i := range board {
		for j:= range board[i] {
			if board[i][j] {
				c.SetFillColor(white)

				minX := j*cellwidth
				minY := j*cellwidth
				maxX := minX + cellwidth
				maxY := minY + cellwidth
				c.ClearRect(minX,minY,maxX,maxY)
			}
		}
	}
	return c.GetImage()
}

func CountCols(board GameBoard) int {
	AssertRectangular(board)

	if len(board) == 0 {
		panic("error")
	}
	return len(board[0])
	
}

func CountRows(board GameBoard) int {
	return len(board)
}

func AssertRectangular(board GameBoard) {
	if len(board) == 0 {
		panic("error")
	}

	//every row has same length
	firstrowlength := len(board[0])

	for row:=1; row<len(board); row++ {
		if len(board[row]) != firstrowlength{
			panic("error")
		}
	}
}