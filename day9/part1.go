package day9

func SolveTaskPart1() int64 {
	diskMeta := ParseInputContent()
	diskBuffer := make([]int, diskMeta.GetTotalSlots())
	for idx := 0; idx < diskMeta.GetTotalSlots(); idx++ {
		diskBuffer[idx] = -1
	}
	freeList, occupiedList := diskMeta.getFreeBlocks(), diskMeta.getOccupiedBlocks()
	for len(freeList) > 0 && len(occupiedList) > 0 {
		nextFreeBlock := &freeList[len(freeList)-1]
		nextOccupiedBlock := &occupiedList[len(occupiedList)-1]
		if nextFreeBlock.getStartOffset() >= nextOccupiedBlock.getStartOffset() {
			break
		}
		if nextFreeBlock.getSize() >= nextOccupiedBlock.getSize() {
			for idx := nextFreeBlock.getStartOffset(); idx < nextFreeBlock.getStartOffset()+nextOccupiedBlock.getSize(); idx++ {
				diskBuffer[idx] = nextOccupiedBlock.getBlockId()
			}
			nextFreeBlock.setSize(nextFreeBlock.getSize() - nextOccupiedBlock.getSize())
			nextFreeBlock.setStartOffset(nextFreeBlock.getStartOffset() + nextOccupiedBlock.getSize())
			occupiedList = occupiedList[:len(occupiedList)-1]
			if nextFreeBlock.getSize() == 0 {
				freeList = freeList[:len(freeList)-1]
			}
		} else {
			for idx := nextFreeBlock.getStartOffset(); idx < nextFreeBlock.getStartOffset()+nextFreeBlock.getSize(); idx++ {
				diskBuffer[idx] = nextOccupiedBlock.getBlockId()
			}
			nextOccupiedBlock.setSize(nextOccupiedBlock.getSize() - nextFreeBlock.getSize())
			freeList = freeList[:len(freeList)-1]
		}
	}
	for _, occupiedBlock := range occupiedList {
		for idx := occupiedBlock.getStartOffset(); idx < occupiedBlock.getStartOffset()+occupiedBlock.getSize(); idx++ {
			diskBuffer[idx] = occupiedBlock.getBlockId()
		}
	}
	return CalculateChecksum(diskBuffer)
}
