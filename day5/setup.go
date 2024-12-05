package day5

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type OrderingAndUpdates struct {
	ordering map[int64]map[int64]bool // map[x] = {a, b, c} means x comes before a, b and c
	updates  [][]int64
}

func (o *OrderingAndUpdates) GetOrdering() map[int64]map[int64]bool {
	return o.ordering
}

func (o *OrderingAndUpdates) GetUpdates() [][]int64 {
	return o.updates
}

func ParseInputContent() *OrderingAndUpdates {
	lines := strings.Split(inputContent, "\n")
	ordering := make(map[int64]map[int64]bool)
	updatesList := make([][]int64, 0)
	parseOrdering := false
	for _, line := range lines {
		if line == "" {
			parseOrdering = true
			continue
		}

		if parseOrdering {
			updatesRaw := strings.Split(line, ",")
			updates := make([]int64, len(updatesRaw))
			for idx, update := range updatesRaw {
				updates[idx] = utils.Must(strconv.ParseInt(update, 10, 64))
			}
			updatesList = append(updatesList, updates)
		} else {
			nodes := strings.Split(line, "|")
			nodeA := utils.Must(strconv.ParseInt(nodes[0], 10, 64))
			nodeB := utils.Must(strconv.ParseInt(nodes[1], 10, 64))
			if set, ok := ordering[nodeA]; ok {
				set[nodeB] = true
			} else {
				ordering[nodeA] = make(map[int64]bool)
				ordering[nodeA][nodeB] = true
			}
		}
	}

	return &OrderingAndUpdates{updates: updatesList, ordering: ordering}
}
