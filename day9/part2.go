package day9

func SolveTaskPart2() int64 {
	diskMeta := ParseInputContent()
	diskBuffer := make([]int, diskMeta.GetTotalSlots())
	for idx := 0; idx < diskMeta.GetTotalSlots(); idx++ {
		diskBuffer[idx] = -1
	}
	freeList, occupiedList := diskMeta.getFreeBlocks(), diskMeta.getOccupiedBlocks()
	occupiedListIdx := len(occupiedList) - 1
	for len(freeList) > 0 && occupiedListIdx > 0 {
		nextOccupiedBlock := &occupiedList[occupiedListIdx]
		var bestBlockIdx int
		var bestBlock *Block
		for freeBlockIdx, freeBlock := range freeList {
			if freeBlock.getSize() >= nextOccupiedBlock.getSize() && freeBlock.getStartOffset() < nextOccupiedBlock.getStartOffset() {
				if bestBlock == nil || freeBlock.getStartOffset() < bestBlock.getStartOffset() {
					bestBlockIdx = freeBlockIdx
					bestBlock = &freeList[freeBlockIdx]
				}
			}
		}
		if bestBlock == nil {
			occupiedListIdx -= 1
			continue
		}

		for idx := bestBlock.getStartOffset(); idx < bestBlock.getStartOffset()+nextOccupiedBlock.getSize(); idx++ {
			diskBuffer[idx] = nextOccupiedBlock.getBlockId()
		}

		bestBlock.setStartOffset(bestBlock.getStartOffset() + nextOccupiedBlock.getSize())
		bestBlock.setSize(bestBlock.getSize() - nextOccupiedBlock.getSize())

		if bestBlock.getSize() == 0 {
			freeList = append(freeList[:bestBlockIdx], freeList[bestBlockIdx+1:]...)
		}

		nextOccupiedBlock.setSize(0)
		occupiedListIdx -= 1
	}

	for _, occupiedBlock := range occupiedList {
		if occupiedBlock.getSize() == 0 {
			continue
		}
		for idx := occupiedBlock.getStartOffset(); idx < occupiedBlock.getStartOffset()+occupiedBlock.getSize(); idx++ {
			diskBuffer[idx] = occupiedBlock.getBlockId()
		}
	}

	return CalculateChecksum(diskBuffer)
}
