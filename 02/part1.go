package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SubmarinePt1 struct {
	HorizontalPosition int
	Depth              int
}

func (c *SubmarinePt1) moveForward(x int) {
	c.HorizontalPosition += x
}

func (c *SubmarinePt1) adjustDepth(x int) {
	c.Depth += x
}

func (c *SubmarinePt1) executeCommand(command string) {
	command_array := strings.Split(command, " ")
	instruction := command_array[0]
	length, _ := strconv.Atoi(command_array[1])

	switch instruction {
	case "forward":
		c.moveForward(length)
	case "down":
		c.adjustDepth(length)
	case "up":
		c.adjustDepth(length * -1)
	}
}

func part1(movements []string) {
	submarinePt1 := SubmarinePt1{0, 0}

	for _, movement := range movements {
		submarinePt1.executeCommand(movement)
	}
	fmt.Println(submarinePt1.Depth * submarinePt1.HorizontalPosition)
}
