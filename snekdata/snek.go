package snekdata

import (
	"math/rand"

	"github.com/cgliu-create/GoSnek/duplicateobj"
)

// SnekControl - holds direction booleans
type SnekControl struct {
	isGoingLeft, isGoingRight, isGoingUp, isGoingDown bool
}

// NewSnekControl - creates a default SnekControl
func NewSnekControl() *SnekControl {
	var newSC SnekControl
	newSC.TurnRight()
	return &newSC
}

// TurnLeft - changes only left to true
func (c *SnekControl) TurnLeft() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = true, false, false, false
}

// TurnRight - changes only right to true
func (c *SnekControl) TurnRight() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = false, true, false, false
}

// TurnUp - changes only up to true
func (c *SnekControl) TurnUp() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = false, false, true, false
}

// TurnDown - changes only down to true
func (c *SnekControl) TurnDown() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = false, false, false, true
}

// Point - holds coordinates
type Point struct {
	X, Y int
}

// Block - holds a point and color reference
type Block struct {
	Coord Point
	Color int
}

// Grid - holds game dimensions
type Grid struct {
	NumHoriz, NumVert int
}

// SnekGame - holds relevant game data
type SnekGame struct {
	GridData    Grid
	Player      *SnekControl
	Snake, Food []Block
}

// NewGrid - creates a new grid struct with specified dimensions
func NewGrid(nh, nv int) Grid {
	grid := Grid{NumHoriz: nh, NumVert: nv}
	return grid
}

// NewSnake - creates a snake block list with an inital block
func NewSnake(x, y int) []Block {
	snake := []Block{Block{Coord: Point{X: x, Y: y}, Color: 1}}
	return snake
}

// NewFood - creates a food block list with an inital block
func NewFood(x, y int) []Block {
	food := []Block{Block{Coord: Point{X: x, Y: y}, Color: 2}}
	return food
}

// NewSnekGame - creates a new Snek Game with defined components
func NewSnekGame(igrid Grid, icontrol *SnekControl, isnake, ifood []Block) *SnekGame {
	game := SnekGame{GridData: igrid, Player: icontrol, Snake: isnake, Food: ifood}
	return &game
}
func (sg *SnekGame) generateFood() {
	rx := rand.Intn(sg.GridData.NumHoriz)
	ry := rand.Intn(sg.GridData.NumVert)
	sg.Food = append(sg.Food, Block{Coord: Point{X: rx, Y: ry}, Color: 2})
}
func (sg *SnekGame) checkFoodEaten() bool {
	if len(sg.Food) == 0 {
		sg.generateFood()
		return false
	}
	head := sg.Snake[len(sg.Snake)-1]
	ifood := sg.Food[0]
	if head.Coord.X == ifood.Coord.X && head.Coord.Y == ifood.Coord.Y {
		sg.Food = []Block{}
		return true
	}
	return false
}
func (sg *SnekGame) growSnake(nx, ny int) {
	sg.Snake = append(sg.Snake, Block{Coord: Point{X: nx, Y: ny}, Color: 1})
	if !sg.checkFoodEaten() {
		sg.Snake = sg.Snake[1:]
	}
}
func (sg *SnekGame) checkSelfCollision() bool {
	var Pointlist []interface{}
	fields := []string{"X", "Y"}
	for _, b := range sg.Snake {
		Pointlist = append(Pointlist, b.Coord)
	}
	repeats := duplicateobj.FindDuplicateObj(fields, Pointlist)
	return len(repeats) != 0
}

func (sg *SnekGame) checkOutOfBounds(x, y int) bool {
	return y < 0 || x < 0 || y > sg.GridData.NumVert || x > sg.GridData.NumHoriz
}

// MoveOrDie - calls all the relevant game operations
// returns if snake is alive
func (sg *SnekGame) MoveOrDie() bool {
	head := sg.Snake[len(sg.Snake)-1]
	nx, ny := head.Coord.X, head.Coord.Y
	if sg.Player.isGoingLeft {
		nx--
	}
	if sg.Player.isGoingRight {
		nx++
	}
	if sg.Player.isGoingUp {
		ny--
	}
	if sg.Player.isGoingDown {
		ny++
	}
	sg.growSnake(nx, ny)
	if sg.checkOutOfBounds(nx, ny) || sg.checkSelfCollision() {
		return false
	}
	return true
}
