package main

import (
	"bytes"
	"fmt"
)

const (
	PLAYER  = 69
	WALL    = 1
	NOTHING = 0
)

type game struct {
}

func (g *game) update() {}

func (g *game) render() {}

func main() {
	width := 80
	height := 20
	level := make([][]byte, height)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			level[h] = make([]byte, width)
		}
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if h == 0 {
				level[h][w] = WALL
			}
			if w == 0 {
				level[h][w] = WALL
			}
			if w == width-1 {
				level[h][w] = WALL
			}
			if h == height-1 {
				level[h][w] = WALL
			}
		}
	}
	buf := new(bytes.Buffer)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if level[h][w] == NOTHING {
				buf.WriteString(" ")
			}
			if level[h][w] == WALL {
				buf.WriteString("▢")
			}
		}
		buf.WriteString("\n")
	}
	fmt.Println(buf.String())
}
