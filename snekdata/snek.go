package snekdata

import (
	"math/rand"

	"github.com/cgliu-create/GoSnek/duplicateobj"
)

type snekControl struct {
	isGoingLeft, isGoingRight, isGoingUp, isGoingDown bool
}

func newSnekControl() *snekControl {
	var newSC snekControl
	newSC.turnRight()
	return &newSC
}
func (c *snekControl) turnLeft() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = true, false, false, false
}
func (c *snekControl) turnRight() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = false, true, false, false
}
func (c *snekControl) turnUp() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = false, false, true, false
}
func (c *snekControl) turnDown() {
	c.isGoingLeft, c.isGoingRight, c.isGoingUp, c.isGoingDown = false, false, false, true
}

var colors = map[int]string{0: "black", 1: "green", 2: "red"}

type grid struct {
	blockWidth, blockheight, numHoriz, numVert int
}

// Point coordinates
type Point struct {
	X, Y int
}
type block struct {
	coord Point
	color int
}
type snekGame struct {
	gridData    grid
	player      snekControl
	snake, food []block
}

func newSnekGame(igrid grid, icontrol snekControl, isnake, ifood []block) *snekGame {
	newGame := snekGame{gridData: igrid, player: icontrol, snake: isnake, food: ifood}
	return &newGame
}
func (sg *snekGame) generateFood() {
	rx := rand.Intn(sg.gridData.numHoriz)
	ry := rand.Intn(sg.gridData.numVert)
	sg.food = append(sg.food, block{coord: Point{X: rx, Y: ry}, color: 2})
}
func (sg *snekGame) checkFoodEaten() bool {
	if len(sg.food) == 0 {
		sg.generateFood()
		return false
	}
	head := sg.snake[len(sg.snake)-1]
	ifood := sg.food[0]
	if head.coord.X == ifood.coord.X && head.coord.Y == ifood.coord.Y {
		sg.food = []block{}
		return true
	}
	return false
}
func (sg *snekGame) growSnake(nx, ny int) {
	sg.snake = append(sg.snake, block{coord: Point{X: nx, Y: ny}, color: 1})
	if !sg.checkFoodEaten() {
		sg.snake = sg.snake[1:]
	}
}
func (sg *snekGame) checkSelfCollision() bool {
	var Pointlist []interface{}
	fields := []string{"X", "Y"}
	for _, b := range sg.snake {
		Pointlist = append(Pointlist, b.coord)
	}
	repeats := duplicateobj.FindDuplicateObj(fields, Pointlist)
	return len(repeats) == 0
}
func (sg *snekGame) checkOutOfBounds(x, y int) bool {
	return y < 0 || x < 0 || y >= sg.gridData.numVert || x >= sg.gridData.numHoriz
}
func (sg *snekGame) moveOrDie() bool {
	head := sg.snake[len(sg.snake)-1]
	nx, ny := head.coord.X, head.coord.Y
	if sg.player.isGoingLeft {
		nx--
	}
	if sg.player.isGoingRight {
		nx++
	}
	if sg.player.isGoingUp {
		ny--
	}
	if sg.player.isGoingDown {
		ny++
	}
	if sg.checkOutOfBounds(nx, ny) || sg.checkSelfCollision() {
		return false
	}
	sg.growSnake(nx, ny)
	return true
}
