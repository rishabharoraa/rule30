package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const DEPTH int = 2048

func createFirstLine(depth int) []int {
	var firstLine []int

	for i := 1; i <= (2*depth)-1; i++ {

		if i == ((2*depth)-1)/2 {
			firstLine = append(firstLine, 1)
		}

		firstLine = append(firstLine, 0)
	}

	return firstLine
}

func iterate(length int, previousLine []int) []int {
	var line []int
	for i := 0; i < length; i++ {

		if i == 0 {
			line = append(line, previousLine[i]|previousLine[i+1])
			continue
		}

		if i == length-1 {
			line = append(line, previousLine[i-1]^previousLine[i])
			continue
		}

		line = append(line, previousLine[i-1]^(previousLine[i]|previousLine[i+1]))
	}
	return line
}

func Paint(graph [][]int) {

	depth := len(graph)
	length := len(graph[0])

	img := image.NewGray(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{length, depth},
		},
	)

	for y := 0; y < depth-1; y++ {
		for x := 0; x < length-1; x++ {
			var col uint8
			if graph[y][x] == 1 {
				col = 255
			} else {
				col = 0
			}
			img.Set(x, y, color.Gray{col})
		}
	}

	file, _ := os.Create("rule30.png")
	png.Encode(file, img)
}

func Plot(depth int, length int) {

	var graph [][]int

	graph = append(graph, createFirstLine(depth))

	for i := 1; i < depth; i++ {
		graph = append(graph, iterate(length, graph[i-1]))
	}

	Paint(graph)

}

func main() {
	Plot(DEPTH, DEPTH*2-1)
}
