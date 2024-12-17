package day17

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type OperandType int

const (
	COMBO OperandType = iota
	LITERAL
)

const ModMask = 0b111

type RegisterState struct {
	A, B, C int64
	PinOut  string
	IP      int64
}

func (r *RegisterState) Combo(x int64) int64 {
	switch x {
	case 0:
		return x
	case 1:
		return x
	case 2:
		return x
	case 3:
		return x
	case 4:
		return r.A
	case 5:
		return r.B
	case 6:
		return r.C
	default:
		panic("Invalid Combo Value")
	}
}

type Instruction interface {
	Code() int
	Name() string
	String() string
	Execute()
	Operand() int64
}

func MakeInstruction(insCode, operand int64, regState *RegisterState) interface{} {
	switch insCode {
	case 0:
		return &Adv{operand: operand, registerState: regState}
	case 1:
		return &Bxl{operand: operand, registerState: regState}
	case 2:
		return &Bst{operand: operand, registerState: regState}
	case 3:
		return &Jnz{operand: operand, registerState: regState}
	case 4:
		return &Bxc{operand: operand, registerState: regState}
	case 5:
		return &Out{operand: operand, registerState: regState}
	case 6:
		return &Bdv{operand: operand, registerState: regState}
	case 7:
		return &Cdv{operand: operand, registerState: regState}
	}
	panic("unreachable")
}

type Adv struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (a *Adv) Code() int {
	return 0
}

func (a *Adv) Name() string {
	return "adv"
}

func (a *Adv) String() string {
	return fmt.Sprintf("adv %d ; A <- A // (1 << Combo(%d))", a.Operand(), a.Operand())
}

func (a *Adv) Execute() {
	a.registerState.A = a.registerState.A / (1 << a.registerState.Combo(a.Operand()))
	a.registerState.IP += 1
}

func (a *Adv) Operand() int64 {
	return a.operand
}

type Bxl struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (b *Bxl) Code() int {
	return 1
}

func (b *Bxl) Name() string {
	return "bxl"
}

func (b *Bxl) String() string {
	return fmt.Sprintf("bxl %d ; B <- B XOR %d", b.Operand(), b.Operand())
}

func (b *Bxl) Execute() {
	b.registerState.B = b.registerState.B ^ b.Operand()
	b.registerState.IP += 1
}

func (b *Bxl) Operand() int64 {
	return b.operand
}

type Bst struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (b *Bst) Code() int {
	return 2
}

func (b *Bst) Name() string {
	return "bst"
}

func (b *Bst) String() string {
	return fmt.Sprintf("bst %d ; B <- Combo(%d) %% 8", b.Operand(), b.Operand())
}

func (b *Bst) Execute() {
	b.registerState.B = b.registerState.Combo(b.Operand()) & ModMask
	b.registerState.IP += 1
}

func (b *Bst) Operand() int64 {
	return b.operand
}

type Jnz struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (j *Jnz) Code() int {
	return 3
}

func (j *Jnz) Name() string {
	return "jnz"
}

func (j *Jnz) String() string {
	return fmt.Sprintf("jnz #%d ; jump to #%d if A != 0 else nop", j.Operand(), j.Operand())
}

func (j *Jnz) Execute() {
	if j.registerState.A != 0 {
		j.registerState.IP = j.Operand()
	} else {
		j.registerState.IP += 1
	}
}

func (j *Jnz) Operand() int64 {
	return j.operand
}

type Bxc struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (b *Bxc) Code() int {
	return 4
}

func (b *Bxc) Name() string {
	return "bxc"
}

func (b *Bxc) String() string {
	return fmt.Sprintf("bxc %d ; B <- B XOR C", b.Operand())
}

func (b *Bxc) Execute() {
	b.registerState.B = b.registerState.B ^ b.registerState.C
	b.registerState.IP += 1
}

func (b *Bxc) Operand() int64 {
	return b.operand
}

type Out struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (o *Out) Code() int {
	return 5
}

func (o *Out) Name() string {
	return "out"
}

func (o *Out) String() string {
	return fmt.Sprintf("out %d ; PinOut <- PinOut + (Combo(%d) %% 8)", o.Operand(), o.Operand())
}

func (o *Out) Execute() {
	o.registerState.PinOut += strconv.Itoa(int(o.registerState.Combo(o.Operand())&ModMask)) + ","
	o.registerState.IP += 1
}

func (o *Out) Operand() int64 {
	return o.operand
}

type Bdv struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (b *Bdv) Code() int {
	return 6
}

func (b *Bdv) Name() string {
	return "bdv"
}

func (b *Bdv) String() string {
	return fmt.Sprintf("bdv %d ; B <- A // (1 << Combo(%d))", b.Operand(), b.Operand())
}

func (b *Bdv) Execute() {
	b.registerState.B = b.registerState.A / (1 << b.registerState.Combo(b.Operand()))
	b.registerState.IP += 1
}

func (b *Bdv) Operand() int64 {
	return b.operand
}

type Cdv struct {
	operand       int64
	registerState *RegisterState
	Instruction
}

func (c *Cdv) Code() int {
	return 7
}

func (c *Cdv) Name() string {
	return "cdv"
}

func (c *Cdv) String() string {
	return fmt.Sprintf("cdv %d ; C <- A // (1 << Combo(%d))", c.Operand(), c.Operand())
}

func (c *Cdv) Execute() {
	c.registerState.C = c.registerState.A / (1 << c.registerState.Combo(c.Operand()))
	c.registerState.IP += 1
}

func (c *Cdv) Operand() int64 {
	return c.operand
}

type Program struct {
	RegisterState *RegisterState
	Instructions  []Instruction
}

func ParseInputContent() Program {
	parsedInput := strings.Split(inputContent, "\n\n")
	rA, rB, rC := parseRegisterValues(parsedInput[0])
	registerState := RegisterState{A: rA, B: rB, C: rC, PinOut: "", IP: 0}
	instructionsRaw := parseInstructionsArrayStr(parsedInput)
	codeSegment := make([]Instruction, len(instructionsRaw)/2)
	insIdx := 0
	for idx := 0; idx < len(instructionsRaw); idx += 2 {
		ins := MakeInstruction(utils.Must(strconv.ParseInt(instructionsRaw[idx], 10, 64)), utils.Must(strconv.ParseInt(instructionsRaw[idx+1], 10, 64)), &registerState)
		codeSegment[insIdx] = ins.(Instruction)
		insIdx += 1
	}
	return Program{&registerState, codeSegment}
}

func parseInstructionsArrayStr(parsedInput []string) []string {
	return strings.Split(strings.Trim(strings.Split(parsedInput[1], ":")[1], " "), ",")
}

func parseRegisterValues(s string) (int64, int64, int64) {
	ans := make([]int64, 3)
	splitted := strings.Split(s, "\n")
	for idx, line := range splitted {
		val := strings.Split(line, ":")[1]
		ans[idx] = utils.Must(strconv.ParseInt(strings.Trim(val, " "), 10, 64))
	}
	return ans[0], ans[1], ans[2]
}
