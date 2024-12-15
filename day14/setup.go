package day14

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type Point struct {
	x, y int64
}

func (p *Point) X() int64 {
	return p.x
}

func (p *Point) Y() int64 {
	return p.y
}

type RobotInfo struct {
	initial  Point
	velocity Point
}

func (r *RobotInfo) Initial() Point {
	return r.initial
}

func (r *RobotInfo) SetInitial(p Point) {
	r.initial = p
}

func (r *RobotInfo) Velocity() Point {
	return r.velocity
}

func ParseInputContent() []RobotInfo {
	lines := strings.Split(inputContent, "\n")
	botInfos := make([]RobotInfo, len(lines))
	for idx, line := range lines {
		botInfos[idx] = parseBotInfoLine(line)
	}
	return botInfos
}

func parseBotInfoLine(line string) RobotInfo {
	parsed := strings.Split(line, " ")
	initial, velocity := parsed[0], parsed[1]
	return RobotInfo{initial: parsePoint(initial), velocity: parsePoint(velocity)}
}

func parsePoint(rawPart string) Point {
	parsed := strings.Split(rawPart[2:], ",")
	x := utils.Must(strconv.ParseInt(parsed[0], 10, 64))
	y := utils.Must(strconv.ParseInt(parsed[1], 10, 64))
	return Point{x, y}
}

func Mod(a, m int64) int64 {
	return (a%m + m) % m
}
