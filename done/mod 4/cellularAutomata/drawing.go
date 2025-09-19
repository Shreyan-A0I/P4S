package main

import (
	"image"
	"canvas"
	"strconv"
)

// DrawGameBoards takes as input a slice of (rectangular) GameBoards and an integer cellWidth.
// It returns a slice of Image objects corresponding to each of these GameBoards,
// where each cell in the image has width and height equal to cellWidth.
func DrawGameBoards(boards []GameBoard, cellWidth int) []image.Image {
    numGenerations := len(boards)
    imageList := make([]image.Image, numGenerations)
    for i := range boards {
        imageList[i] = DrawGameBoard(boards[i], cellWidth)
    }
    return imageList
}

// DrawGameBoard takes as input a (rectangular) GameBoard and an integer cellWidth.
// It returns an Image object corresponding to the GameBoard, where each cell in the image has width and height equal to cellWidth.

// DrawGameBoard takes as input a (rectangular) GameBoard and an integer cellWidth.
// It returns an Image object corresponding to the GameBoard,
// where each cell in the image has width and height equal to cellWidth.
func DrawGameBoard(board GameBoard, cellWidth int) image.Image {
    numRows := CountRows(board)
    numColumns := CountColumns(board)

    // create the canvas
    height := numRows * cellWidth
    width := numColumns * cellWidth
    c := canvas.CreateNewCanvas(width, height)

    // declare colors
    darkGray := canvas.MakeColor(60, 60, 60)
    white := canvas.MakeColor(255, 255, 255)
    red := canvas.MakeColor(239, 71, 111)
    green := canvas.MakeColor(6, 214, 160)
    yellow := canvas.MakeColor(255, 255, 0)
    orange := canvas.MakeColor(255, 165, 0)
    purple := canvas.MakeColor(160, 32, 240)
    blue := canvas.MakeColor(17, 138, 178)
    black := canvas.MakeColor(0, 0, 0)

    //set the entire board as dark gray
    c.SetFillColor(darkGray)
    c.Clear()

    // fill in colored squares
    for i := range board {
        for j := range board[i] {
            switch board[i][j] {
            case 0:
                c.SetFillColor(darkGray)
            case 1:
                c.SetFillColor(white)
            case 2:
                c.SetFillColor(red)
            case 3:
                c.SetFillColor(green)
            case 4:
                c.SetFillColor(yellow)
            case 5:
                c.SetFillColor(orange)
            case 6:
                c.SetFillColor(purple)
            case 7:
                c.SetFillColor(blue)
            case 8:
                c.SetFillColor(black)
            default:
                panic("Error: Out of range value " + strconv.Itoa(board[i][j]) + " in board when drawing board.")
            }

            //only need to draw circle if it's not dead since it would equal background color
            if board[i][j] != 0 {
                // set circle's central coordinates
                x := float64(j*cellWidth) + float64(cellWidth)/2
                y := float64(i*cellWidth) + float64(cellWidth)/2

                scalingFactor := 0.8 // to make circle smaller

                c.Circle(float64(x), float64(y), scalingFactor*float64(cellWidth)/2)
                c.Fill()
            }
        }
    }

    return c.GetImage()
}