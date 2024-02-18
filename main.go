package main

import (
	"encoding/binary"
	"image/color"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sstallion/go-hid"
)

var wheelRotation uint16
var throttle uint8
var breakPedal uint8
var handbrake bool

func main() {
	initHID()

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(385, 160, "go-obs-joystick-overlay")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	graphPoints := 60 * 3
	throttleGraph := NewGraph(graphPoints)
	breakGraph := NewGraph(graphPoints)

	font := rl.LoadFont("Roboto-Bold.ttf")
	rl.SetTextureFilter(font.Texture, rl.FilterTrilinear)
	defer rl.UnloadFont(font)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		if handbrake {
			rl.DrawTextEx(font, "Handbrake", rl.NewVector2(45, 74), 14, 2, rl.Red)
		} else {
			rl.DrawTextEx(font, "Handbrake", rl.NewVector2(45, 74), 14, 2, rl.Gray)
		}

		drawWheel(80, 80)
		drawVerticalBar(150, 20, 15, 120, rl.Remap(float32(throttle), 255, 0, 0, 1), rl.Green)
		drawVerticalBar(170, 20, 15, 120, rl.Remap(float32(breakPedal), 255, 0, 0, 1), rl.Red)

		breakGraph.AddPoint(breakPedal)
		drawGraph(breakGraph, 190, 20, 120, rl.Red, true)
		throttleGraph.AddPoint(throttle)
		drawGraph(throttleGraph, 190, 20, 120, rl.Green, false)

		rl.EndDrawing()
	}
}

func initHID() {
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	d, err := hid.OpenFirst(0x46d, 0xc262)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 15)
	go func() {
		for {
			d.Read(b)
			wheelRotation = binary.LittleEndian.Uint16(b[4:6])
			throttle = b[6]
			breakPedal = b[7]
			// clutch = b[8]
			handbrake = b[1]&0b100000 > 0
		}
	}()
}

// value must be in the range 0.0 - 1.0.
// func drawHorizontalBar(x, y, width, height int32, value float32, color color.RGBA) {
// 	rl.DrawRectangle(x, y, int32(value*float32(width)), height, color)
// 	rl.DrawRectangleLines(x, y, width, height, rl.White)
// }

// value must be in the range 0.0 - 1.0.
func drawVerticalBar(x, y, width, height int32, value float32, color color.RGBA) {
	barHeight := int32(value * float32(height))
	rl.DrawRectangle(x, y+(height-barHeight), width, barHeight, color)
	rl.DrawRectangleLines(x, y, width, height, rl.White)
}

func drawWheel(x, y int) {
	direction := rl.Remap(float32(wheelRotation), 0x0000, 0xFFFF, 0, 900)
	rl.DrawRing(rl.NewVector2(float32(x), float32(y)), 50, 60, 0, 360, 64, rl.White)
	rl.DrawRing(rl.NewVector2(float32(x), float32(y)), 49, 61, -180+direction-10, -180+direction+10, 64, rl.Red)
}

type Graph struct {
	maxPoints    int
	points       []uint8
	currentIndex int
}

func NewGraph(maxPoints int) Graph {
	return Graph{
		maxPoints:    maxPoints,
		points:       make([]uint8, maxPoints),
		currentIndex: 0,
	}
}

func (g *Graph) AddPoint(value uint8) {
	g.points[g.currentIndex] = value

	if g.currentIndex+1 < g.maxPoints {
		g.currentIndex++
	} else {
		g.currentIndex = 0
	}
}

func drawGraph(graph Graph, x, y, height int, color color.RGBA, drawBorder bool) {
	rl.SetLineWidth(2)
	if drawBorder {
		rl.DrawRectangleLines(int32(x), int32(y), int32(graph.maxPoints), int32(height), rl.White)
	}

	for i := 1; i < graph.maxPoints; i++ {
		graphFrame := (graph.currentIndex + i) % graph.maxPoints

		var last uint8
		if graphFrame-1 >= 0 {
			last = graph.points[graphFrame-1]
		} else {
			last = graph.points[graph.maxPoints-1]
		}

		cur := graph.points[graphFrame]

		v := rl.Remap(float32(cur), 0, 255, 0, float32(height))

		vlast := rl.Remap(float32(last), 0, 255, 0, float32(height))
		rl.EnableSmoothLines()
		rl.DrawLine(
			int32(x+i-1),
			int32(y)+int32(vlast),
			int32(x+i),
			int32(y)+int32(v), color)
	}
}
