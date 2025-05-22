package main

import (
	"math/rand"
	"math"
)


func Walk(board *[][]int, depth int) {
	for range 10 {
		startY := rand.Intn(len(*board))
		startX := rand.Intn(len((*board)[0]))

		walkHelper(board, depth, startX, startY)
	}
}

func walkHelper(board *[][]int, depth, x, y int) {
	directions := []Point{
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}

	dirChoice := rand.Intn(4)

	choice := directions[dirChoice]
	y += choice.Y
	x += choice.X

	if y >= len(*board) || y < 0 {
		return
	}
	if x >= len((*board)[0]) || x < 0{
		return
	}
	if depth <= 0 {
		return
	}
	if (*board)[y][x] == 1 {
		depth += 1
	}


	(*board)[y][x] = 1
	walkHelper(board, depth - 1, x, y)
}

func findStart(board *[][]int) Point {
	y := rand.Intn(len(*board))
	x := rand.Intn(len((*board)[0]))

	for (*board)[y][x] != 1 {
		y = rand.Intn(len(*board))
		x = rand.Intn(len((*board)[0]))
	}

	start := Point{x, y}

	return start
}


func findEnd(board *[][]int, start Point) Point {
	rows := len(*board)
	if rows == 0 {
		return Point{-1, -1}
	}
	cols := len((*board)[0])
	if cols == 0 {
		return Point{-1, -1}
	}

	queue := []Point{start}
	visited := make(map[Point]bool)
	visited[start] = true
	farthestPoint := start
	maxDistance := 0.0

	directions := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentDistance := calculateDistance(start, current)
		if currentDistance > maxDistance {
			maxDistance = currentDistance
			farthestPoint = current
		}

		for _, dir := range directions {
			next := Point{X: current.X + dir.X, Y: current.Y + dir.Y}

			if next.X >= 0 && next.X < cols && next.Y >= 0 && next.Y < rows &&
				(*board)[next.Y][next.X] != 0 && !visited[next] { // Adjust -1 for wall representation
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return farthestPoint
}
func calculateDistance(p1, p2 Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
