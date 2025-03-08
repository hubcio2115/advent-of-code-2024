package day6

import (
	"slices"
	"strings"
)

const (
	UP       = "^"
	DOWN     = "v"
	LEFT     = "<"
	RIGHT    = ">"
	OBSTACLE = '#'
)

type Coords struct {
	x int
	y int
}

// BUG: Test is passing but it gives a wrong solution
func Part1(input string) int {
	rows := strings.Split(input, "\n")
	boardSize := Coords{len(rows[0]), len(rows)}

	var guard Guard
	for y, row := range rows {
		for x, character := range strings.Split(row, "") {
			if slices.Contains([]string{UP, DOWN, LEFT, RIGHT}, character) {
				guard = Guard{Coords{x, y}, character}
			}
		}
	}

	uniquePositions := map[Coords]bool{}
	for guard.position.x > 0 || guard.position.x < boardSize.x || guard.position.y > 0 || guard.position.y < boardSize.y {
		nextPosition := Coords{guard.position.x, guard.position.y}
		switch guard.direction {
		case UP:
			nextPosition.y -= 1
		case DOWN:
			nextPosition.y += 1
		case LEFT:
			nextPosition.x -= 1
		case RIGHT:
			nextPosition.x += 1
		}

		isOutOfBounds := nextPosition.x == boardSize.x || nextPosition.x < 0 || nextPosition.y == boardSize.y || nextPosition.y < 0
		if isOutOfBounds {
			break
		}

		isObstacle := []rune(rows[nextPosition.y])[nextPosition.x] == OBSTACLE
		if isObstacle {
			guard.rotateClockwise()
		} else {
			guard.position = nextPosition
			uniquePositions[nextPosition] = true
		}
	}

	return len(uniquePositions)
}

type Guard struct {
	position  Coords
	direction string
}

func (g *Guard) rotateClockwise() {
	switch g.direction {
	case UP:
		g.direction = RIGHT
	case DOWN:
		g.direction = LEFT
	case LEFT:
		g.direction = UP
	case RIGHT:
		g.direction = DOWN
	}
}
