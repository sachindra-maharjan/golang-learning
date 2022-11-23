package main

import "fmt"

type House struct {
	window string
	door   string
	floor  string
}

type IBuilder interface {
	setWindow()
	setDoor()
	setFloor()
	getHouse() House
}

type XBuilder struct {
	House
}

type YBuilder struct {
	House
}

func (h *House) String() {
	fmt.Printf("window=%s, door=%s, floor=%s\n", h.window, h.door, h.floor)
}

func NewXBuilder() *XBuilder {
	return &XBuilder{}
}

func NewYBuilder() *YBuilder {
	return &YBuilder{}
}

func (x *XBuilder) setWindow() {
	x.window = "xWindow"
}
func (x *XBuilder) setDoor() {
	x.door = "xDoor"
}
func (x *XBuilder) setFloor() {
	x.floor = "xFloor"
}
func (x *XBuilder) getHouse() House {
	return House{
		window: x.window,
		door:   x.door,
		floor:  x.floor,
	}
}

func (y *YBuilder) setWindow() {
	y.window = "yWindow"
}
func (y *YBuilder) setDoor() {
	y.door = "yDoor"
}
func (y *YBuilder) setFloor() {
	y.floor = "yFloor"
}
func (y *YBuilder) getHouse() House {
	return House{
		window: y.window,
		door:   y.door,
		floor:  y.floor,
	}
}

func GetBuilder(builder string) IBuilder {
	switch builder {
	case "X":
		return NewXBuilder()
	case "Y":
		return NewYBuilder()
	default:
		return nil
	}
}

type Director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{builder}
}

func (d *Director) SetDirector(builder IBuilder) {
	d.builder = builder
}

func (d *Director) BuildHouse() House {
	d.builder.setDoor()
	d.builder.setFloor()
	d.builder.setWindow()
	return d.builder.getHouse()
}

