package day11

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

func SolveTaskPart1() int64 {
	stones := ParseInputContent()
	fmt.Printf("stones: %v\n", stones)
	var ans atomic.Int64
	ans.Store(0)
	var wg sync.WaitGroup
	wg.Add(len(stones))
	for _, stone := range stones {
		go func() {
			defer wg.Done()
			ans.Add(splitStonesTill25(stone, 0))
		}()
	}
	wg.Wait()
	return ans.Load()
}

func splitStonesTill25(stone int64, blinkCnt int) int64 {
	if blinkCnt == 25 {
		return 1
	}
	if stone == 0 {
		return splitStonesTill25(1, blinkCnt+1)
	} else if DigitCnt(stone)%2 == 0 {
		part1, part2 := SplitEvenNumber(stone)
		return splitStonesTill25(part1, blinkCnt+1) + splitStonesTill25(part2, blinkCnt+1)
	} else {
		return splitStonesTill25(stone*2024, blinkCnt+1)
	}
}

func SplitEvenNumber(stone int64) (int64, int64) {
	stoneStr := strconv.FormatInt(stone, 10)
	ln := len(stoneStr)
	return utils.Must(strconv.ParseInt(stoneStr[:ln/2], 10, 64)), utils.Must(strconv.ParseInt(stoneStr[ln/2:], 10, 64))
}

func DigitCnt(stone int64) int {
	return len(strconv.FormatInt(stone, 10))
}
