package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Vec2i struct {
	X int
	Y int
}

func NewVec2i(x, y int) Vec2i {
	return Vec2i{X: x, Y: y}
}
func (v Vec2i) ToVector2() rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}
