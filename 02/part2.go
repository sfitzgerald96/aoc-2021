package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SubmarinePt2 struct {
	HorizontalPosition int
	Depth              int
	Aim                int
}

func (c *SubmarinePt2) moveForward(x int) {
	c.HorizontalPosition += x
	c.Depth += c.Aim * x
}

func (c *SubmarinePt2) adjustDepth(x int) {
	c.Aim += x
}

func (c *SubmarinePt2) executeCommand(command string) {
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

func part2(movements []string) {
	submarinePt2 := SubmarinePt2{0, 0, 0}

	for _, movement := range movements {
		submarinePt2.executeCommand(movement)
	}

	fmt.Println(submarinePt2.Depth * submarinePt2.HorizontalPosition)
}
