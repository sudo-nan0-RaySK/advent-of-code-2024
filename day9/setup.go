package day9

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"slices"
	"strconv"
)

//go:embed static/input.txt
var inputContent string

type Block struct {
	startOffset int
	size        int
	blockId     int
}

func (b *Block) getStartOffset() int {
	return b.startOffset
}

func (b *Block) setStartOffset(offset int) {
	b.startOffset = offset
}

func (b *Block) getSize() int {
	return b.size
}

func (b *Block) setSize(size int) {
	b.size = size
}

func (b *Block) getBlockId() int {
	return b.blockId
}

type DiskLayout struct {
	freeBlocksList     []Block
	occupiedBlocksList []Block
	totalSlots         int
}

func (d *DiskLayout) getFreeBlocks() []Block {
	return d.freeBlocksList
}

func (d *DiskLayout) getOccupiedBlocks() []Block {
	return d.occupiedBlocksList
}

func (d *DiskLayout) GetTotalSlots() int {
	return d.totalSlots
}

func ParseInputContent() DiskLayout {
	freeBlocksList := make([]Block, 0)
	occupiedBlocksList := make([]Block, 0)
	totalSlots, blockId := 0, 0
	for offset, szRaw := range inputContent {
		sz := int(utils.Must(strconv.ParseInt(string(szRaw), 10, 64)))
		if offset%2 == 0 {
			occupiedBlocksList = append(occupiedBlocksList, Block{totalSlots, sz, blockId})
			blockId = blockId + 1
		} else {
			freeBlocksList = append(freeBlocksList, Block{totalSlots, sz, -1})
		}
		totalSlots += sz
	}
	slices.Reverse(freeBlocksList)
	return DiskLayout{freeBlocksList, occupiedBlocksList, totalSlots}
}

func CalculateChecksum(buffer []int) int64 {
	checksum := int64(0)
	for idx := 0; idx < len(buffer); idx++ {
		if buffer[idx] == -1 {
			continue
		}
		checksum += int64(buffer[idx]) * int64(idx)
	}
	return checksum
}
