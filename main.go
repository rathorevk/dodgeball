package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Direction struct {
	DX, DY int
	Name   string
}

var directions = []Direction{
	{0, 1, "N"}, {1, 1, "NE"}, {1, 0, "E"}, {1, -1, "SE"},
	{0, -1, "S"}, {-1, -1, "SW"}, {-1, 0, "W"}, {-1, 1, "NW"},
}

func distance(p1, p2 Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func isInDirection(current, target Point, dir Direction) bool {
	if current.X == target.X && current.Y == target.Y {
		return false
	}

	dx := target.X - current.X
	dy := target.Y - current.Y

	if dir.DX == 0 {
		if dx != 0 {
			return false
		}
		if dir.DY > 0 {
			return dy > 0
		} else {
			return dy < 0
		}
	}

	if dir.DY == 0 {
		if dy != 0 {
			return false
		}
		if dir.DX > 0 {
			return dx > 0
		} else {
			return dx < 0
		}
	}

	if dir.DX*dx <= 0 || dir.DY*dy <= 0 {
		return false
	}

	return dx*dir.DY == dy*dir.DX
}

func findNearestInDirection(
	current Point,
	players []Point,
	startDir Direction,
	visited map[int]bool,
) (*Point, string, bool) {
	var nearest *Point
	minDistance := math.Inf(1)
	foundDirection := ""
	found := false

	clockwiseOrder := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}

	startIndex := -1
	for i, dirName := range clockwiseOrder {
		if dirName == startDir.Name {
			startIndex = i
			break
		}
	}

	if startIndex == -1 {
		return nil, "", false
	}

	for offset := 0; offset < 8; offset++ {
		dirIndex := (startIndex + offset) % 8
		dirName := clockwiseOrder[dirIndex]

		dir := getDirectionByName(dirName)
		if dir == nil {
			continue
		}

		for i, player := range players {
			if visited[i] {
				continue
			}
			if isInDirection(current, player, *dir) {
				dist := distance(current, player)
				if dist < minDistance {
					minDistance = dist
					nearest = &player
					foundDirection = dirName
					found = true
				}
			}
		}

		if found {
			break
		}
	}

	return nearest, foundDirection, found
}

func findPlayerIndex(players []Point, target Point) int {
	for i, player := range players {
		if player.X == target.X && player.Y == target.Y {
			return i
		}
	}
	return -1
}

func getDirectionByName(dirName string) *Direction {
	for _, dir := range directions {
		if dir.Name == dirName {
			return &dir
		}
	}
	return nil
}

func getOppositeDirection(direction string) string {
	opposites := map[string]string{
		"N": "S", "S": "N", "E": "W", "W": "E",
		"NE": "SW", "SW": "NE", "NW": "SE", "SE": "NW",
	}
	return opposites[direction]
}

func rotateClockwise(currentDirection string) string {
	clockwiseOrder := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}

	for i, dir := range clockwiseOrder {
		if dir == currentDirection {
			return clockwiseOrder[(i+1)%8]
		}
	}
	return currentDirection
}

func simulateGame(players []Point, startDirection string, startPlayer int) (int, int) {
	currentPlayer := startPlayer
	incomingDirection := startDirection
	throws := 0
	maxThrows := len(players)
	visited := make(map[int]bool)

	visited[currentPlayer] = true

	for throws < maxThrows {
		outgoingDirection := rotateClockwise(incomingDirection)

		dir := getDirectionByName(outgoingDirection)
		if dir == nil {
			break
		}

		current := players[currentPlayer]
		nearest, outgoingDirection, _ := findNearestInDirection(current, players, *dir, visited)

		if nearest == nil {
			break
		}

		targetPlayer := findPlayerIndex(players, *nearest)
		if targetPlayer == -1 {
			break
		}

		throws++
		visited[targetPlayer] = true

		currentPlayer = targetPlayer
		incomingDirection = getOppositeDirection(outgoingDirection)
	}

	return throws, currentPlayer
}

func main() {
	var testCases int
	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var numPlayers int
		fmt.Scan(&numPlayers)

		players := make([]Point, numPlayers)
		for i := 0; i < numPlayers; i++ {
			fmt.Scan(&players[i].X, &players[i].Y)
		}

		var startDirection string
		var startPlayer int
		fmt.Scan(&startDirection)
		fmt.Scan(&startPlayer)

		startPlayer--

		throws, endPlayer := simulateGame(players, startDirection, startPlayer)
		fmt.Printf("%d %d\n", throws, endPlayer+1)
	}
}
