package day17

func SolveTaskPart1() string {
	program := ParseInputContent()
	regState := program.RegisterState
	codeSegment := program.Instructions
	for int(regState.IP) < len(codeSegment) {
		codeSegment[int(regState.IP)].Execute()
	}
	outCode := regState.PinOut
	return outCode[:len(outCode)-1]
}
