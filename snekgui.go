package main

import (
	"os"

	"github.com/cgliu-create/GoSnek/snekdata"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var nh, nv, w, h = 20, 20, 400, 400

var (
	dimensions    = snekdata.NewGrid(nh, nv)
	controls      = snekdata.NewSnekControl()
	startingSnake = snekdata.NewSnake(1, 1)
	startingFood  = snekdata.NewFood(5, 5)
	game          = snekdata.NewSnekGame(dimensions, controls, startingSnake, startingFood)
)

var colors = map[int]*gui.QColor{
	0: gui.NewQColor3(0, 0, 0, 255),
	1: gui.NewQColor3(0, 255, 0, 255),
	2: gui.NewQColor3(255, 0, 0, 255)}

var (
	view  *widgets.QGraphicsView
	scene *widgets.QGraphicsScene
	item  *widgets.QGraphicsPixmapItem
	timer *core.QTimer
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Main Window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Snek")
	window.SetMinimumSize2(w, h)

	scene = widgets.NewQGraphicsScene(nil)
	view = widgets.NewQGraphicsView(nil)
	scene.ConnectKeyPressEvent(keyPressEvent)
	view.ConnectResizeEvent(resizeEvent)

	img := gui.NewQImage3(nh, nv, gui.QImage__Format_ARGB32)
	for i := 0; i < nh; i++ {
		for j := 0; j < nv; j++ {
			img.SetPixelColor2(i, j, colors[0])
		}
	}
	item = widgets.NewQGraphicsPixmapItem2(gui.NewQPixmap().FromImage(img, 0), nil)

	scene.AddItem(item)
	view.SetScene(scene)
	view.Show()
	window.SetCentralWidget(view)

	widgets.QApplication_SetStyle2("fusion")
	window.Show()

	timer = core.NewQTimer(nil)
	timer.SetInterval(500)
	timer.Start2()
	timer.ConnectTimerEvent(timerEvent)

	app.Exec()

}

func timerEvent(e *core.QTimerEvent) {
	gamefunction()
}

//this works
func clear() {
	img := item.Pixmap().ToImage()
	for i := 0; i < nh; i++ {
		for j := 0; j < nv; j++ {
			img.SetPixelColor2(i, j, colors[0])
		}
	}
	item.SetPixmap(gui.NewQPixmap().FromImage(img, 0))
}

// this works
func drawpixel(x, y, c int) {
	img := item.Pixmap().ToImage()
	img.SetPixelColor2(x, y, colors[c])
	item.SetPixmap(gui.NewQPixmap().FromImage(img, 0))
}

// this works
func drawBlocks(blocks []snekdata.Block) {
	for _, block := range blocks {
		c := block.Color
		x, y := block.Coord.X, block.Coord.Y
		drawpixel(x, y, c)
	}
}
func gamefunction() {
	game.MoveOrDie()
	clear()
	drawBlocks(game.Snake)
	drawBlocks(game.Food)
}
func keyPressEvent(e *gui.QKeyEvent) {
	switch int32(e.Key()) {
	case int32(core.Qt__Key_Left):
		game.Player.TurnLeft()
	case int32(core.Qt__Key_Right):
		game.Player.TurnRight()
	case int32(core.Qt__Key_Up):
		game.Player.TurnUp()
	case int32(core.Qt__Key_Down):
		game.Player.TurnDown()
	}
}

func resizeEvent(e *gui.QResizeEvent) {
	view.FitInView(scene.ItemsBoundingRect(), core.Qt__KeepAspectRatio)
}
