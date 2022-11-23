package main

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

func (x *XBuilder) setWindow() string {
	return "xWindow"
}
func (x *XBuilder) setDoor() string {
	return "xDoor"
}
func (x *XBuilder) setFloor() string {
	return "xFloor"
}
func (x *XBuilder) getHouse() House {
	return House{
		window: x.window,
		door:   x.door,
		floor:  x.floor,
	}
}

func (y *YBuilder) setWindow() string {
	return "xWindow"
}
func (y *YBuilder) setDoor() string {
	return "xDoor"
}
func (y *YBuilder) setFloor() string {
	return "xFloor"
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
		return &XBuilder{}
	case "Y":
		return &YBuilder{}
	default:
		return nil
	}
}
