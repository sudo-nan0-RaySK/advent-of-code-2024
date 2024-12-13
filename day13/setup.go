package day13

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type Point struct {
	x, y int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

type MachineInfo struct {
	btnADelta     Point
	btnBDelta     Point
	prizeLocation Point
}

func (m *MachineInfo) ButtonADelta() Point {
	return m.btnADelta
}

func (m *MachineInfo) ButtonBDelta() Point {
	return m.btnBDelta
}

func (m *MachineInfo) PrizeLocation() Point {
	return m.prizeLocation
}

func ParseInputContent() []MachineInfo {
	lines := strings.Split(inputContent, "\n\n")
	machineInfos := make([]MachineInfo, len(lines))
	for idx, machineInfoRaw := range lines {
		machineInfo := strings.Split(machineInfoRaw, "\n")
		btnADelta := parsePointInfo(machineInfo[0][10:])
		btnBDelta := parsePointInfo(machineInfo[1][10:])
		prizeLocation := parsePointInfo(machineInfo[2][7:])
		machineInfos[idx] = MachineInfo{btnADelta: btnADelta, btnBDelta: btnBDelta, prizeLocation: prizeLocation}
	}
	return machineInfos
}

func parsePointInfo(btnInfoRaw string) Point {
	pts := strings.Split(btnInfoRaw, ",")
	ptX, ptY := strings.Trim(pts[0], " "), strings.Trim(pts[1], " ")
	ptX = strings.Replace(ptX, "=", "", -1)
	ptY = strings.Replace(ptY, "=", "", -1)
	x := int(utils.Must(strconv.ParseInt(ptX[1:], 10, 64)))
	y := int(utils.Must(strconv.ParseInt(ptY[1:], 10, 64)))
	return Point{x, y}
}
