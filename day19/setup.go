package day19

import (
	_ "embed"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type CombinationInfo struct {
	AvailableDesigns []string
	DesignsToCheck   []string
}

func ParseInputContent() CombinationInfo {
	parts := strings.Split(inputContent, "\n\n")
	availableDesigns := strings.Split(parts[0], ",")
	for idx, availableDesign := range availableDesigns {
		availableDesigns[idx] = strings.TrimSpace(availableDesign)
	}
	designsToCheck := strings.Split(parts[1], "\n")
	return CombinationInfo{AvailableDesigns: availableDesigns, DesignsToCheck: designsToCheck}
}
