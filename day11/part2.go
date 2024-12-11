package day11

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type MemoKey struct {
	stone    int64
	blinkCnt int
}

func SolveTaskPart2() int64 {
	stones := ParseInputContent()
	fmt.Printf("stones: %v\n", stones)
	var ans atomic.Int64
	ans.Store(0)
	var wg sync.WaitGroup
	wg.Add(len(stones))
	for _, stone := range stones {
		go func() {
			defer wg.Done()
			memo := make(map[MemoKey]int64)
			ans.Add(splitStonesTill75(stone, 0, &memo))
		}()
	}
	wg.Wait()
	return ans.Load()
}

func splitStonesTill75(stone int64, blinkCnt int, memo *map[MemoKey]int64) int64 {
	key := MemoKey{stone: stone, blinkCnt: blinkCnt}
	if _, ok := (*memo)[key]; ok {
		return (*memo)[key]
	}
	if blinkCnt == 75 {
		return 1
	}
	if stone == 0 {
		(*memo)[key] = splitStonesTill75(1, blinkCnt+1, memo)
	} else if DigitCnt(stone)%2 == 0 {
		part1, part2 := SplitEvenNumber(stone)
		(*memo)[key] = splitStonesTill75(part1, blinkCnt+1, memo) + splitStonesTill75(part2, blinkCnt+1, memo)
	} else {
		(*memo)[key] = splitStonesTill75(stone*2024, blinkCnt+1, memo)
	}
	return (*memo)[key]
}
