package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var SPACE_WIDTH, SPACE_HEIGHT = 7, 11

var SPACE_WIDTH, SPACE_HEIGHT = 101, 103

type Position struct {
	y, x int
}

type Robot struct {
	positionX, positionY int
	velocityX, velocityY int
}

func main() {
	robots := getRobots()
	for i := range robots {
		robots[i] = moveRobotBySeconds(robots[i], 100)
	}
	safetyFactor := getSafetyFactor(robots)
	fmt.Println(safetyFactor)

	//part 2
	robots = getRobots()
	for t := 0; t < 10000; t++ {
		for i := range robots {
			robots[i] = moveRobotBySeconds(robots[i], 1)
		}
		printAreaConditionally(robots, t+1)
	}
}

func printAreaConditionally(robots []Robot, secondsPassed int) {
	robotSet := map[Position]bool{}
	for _, robot := range robots {
		robotSet[Position{x: robot.positionX, y: robot.positionY}] = true
	}

	area := []string{}
	shouldPrint := false
	for y := 0; y < SPACE_HEIGHT; y++ {
		line := ""
		for x := 0; x < SPACE_WIDTH; x++ {
			if _, ok := robotSet[Position{x: x, y: y}]; ok {
				line += "■"
			} else {
				line += " "
			}
		}
		area = append(area, line)
		if strings.Contains(line, "■■■■■") { // 5 in a row
			shouldPrint = true
		}
	}

	if shouldPrint {
		fmt.Println(secondsPassed)
		for _, line := range area {
			fmt.Println(line)
		}
	}
}

func getSafetyFactor(robots []Robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, robot := range robots {
		if robot.positionY < SPACE_HEIGHT/2 {
			if robot.positionX < SPACE_WIDTH/2 {
				q1++
			} else if robot.positionX > SPACE_WIDTH/2 {
				q2++
			}
		} else if robot.positionY > SPACE_HEIGHT/2 {
			if robot.positionX < SPACE_WIDTH/2 {
				q3++
			} else if robot.positionX > SPACE_WIDTH/2 {
				q4++
			}
		}
	}

	return q1 * q2 * q3 * q4
}

func moveRobotBySeconds(robot Robot, time int) Robot {
	endStateX := (robot.positionX + robot.velocityX*time) % SPACE_WIDTH
	endStateY := (robot.positionY + robot.velocityY*time) % SPACE_HEIGHT
	if endStateX < 0 {
		endStateX = SPACE_WIDTH + endStateX
	}
	if endStateY < 0 {
		endStateY = SPACE_HEIGHT + endStateY
	}

	robot.positionX = endStateX
	robot.positionY = endStateY
	return robot
}

func getRobots() (robots []Robot) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		positionAndVelocity := strings.Split(line, " ")
		position, velocity := positionAndVelocity[0], positionAndVelocity[1]
		positionXY := strings.Split(strings.Replace(position, "p=", "", 1), ",")
		velocityXY := strings.Split(strings.Replace(velocity, "v=", "", 1), ",")
		positionX, _ := strconv.Atoi(positionXY[0])
		positionY, _ := strconv.Atoi(positionXY[1])
		velocityX, _ := strconv.Atoi(velocityXY[0])
		velocityY, _ := strconv.Atoi(velocityXY[1])
		robots = append(robots, Robot{positionX: positionX, positionY: positionY, velocityX: velocityX, velocityY: velocityY})
	}
	return
}
