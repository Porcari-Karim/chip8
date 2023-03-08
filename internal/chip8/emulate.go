package chip8

import (
	"fmt"
	"math/rand"
)

func (e *Emulator) FetchOpCode() OpCode {
	pc := e.registers.Pc
	opcode := OpCode(e.ram.data[pc])
	opcode = opcode<<8 | OpCode(e.ram.data[pc+1])
	return OpCode(opcode)
}

func (e *Emulator) DecodeOpCode(opcode OpCode) {
	I := opcode & 0xF000
	x := byte((opcode & 0x0F00) >> 2)
	y := byte((opcode & 0x00F0) >> 1)
	n := byte(opcode & 0x000F)
	//nnn := opcode & 0x0FFF
	//fmt.Println(x, y, n)

	switch I {
	case 0x0000:
		{
			break
		}

	case 0x1000:
		{
			break
		}

	case 0x2000:
		{
			break
		}

	case 0x3000:
		{
			break
		}

	case 0x4000:
		{
			break
		}

	case 0x5000:
		{
			break
		}

	case 0x6000:
		{
			break
		}

	case 0x7000:
		{
			break
		}
	case 0x8000:
		{
			break
		}
	case 0x9000:
		{
			break
		}
	case 0xA000:
		{
			break
		}
	case 0xD000:
		{
			fmt.Printf("0x%04x \n", opcode)
			e.DRWD(x, y, n)
			break
		}
	}
}

func (e *Emulator) ClearScreen() {
	e.display.pixels = make([]byte, (32*64)/8)
}

func (e *Emulator) GoTo(adress uint16) {
	e.registers.Pc = adress
}

func (e *Emulator) Return() {
	e.registers.Pc = e.stack.Pop().(uint16)
}

func (e *Emulator) Call(adress uint16) {
	e.stack.Push(e.registers.Pc)
	e.registers.Pc = adress
}

func (e *Emulator) SE3(x byte, kk byte) {
	if e.registers.Vx[x] == kk {
		e.registers.Pc += 2
	}
}
func (e *Emulator) SE4(x byte, kk byte) { // Aka SNE
	if e.registers.Vx[x] != kk {
		e.registers.Pc += 2
	}
}
func (e *Emulator) SE5(x byte, y byte) {
	if e.registers.Vx[x] == e.registers.Vx[y] {
		e.registers.Pc += 2
	}
}
func (e *Emulator) LD6(x byte, kk byte) {
	e.registers.Vx[x] = kk
}
func (e *Emulator) ADD7(x byte, kk byte) {
	e.registers.Vx[x] += kk
}
func (e *Emulator) LD8(x byte, y byte) {
	e.registers.Vx[x] = e.registers.Vx[y]
}
func (e *Emulator) OR8(x byte, y byte) {
	e.registers.Vx[x] = e.registers.Vx[x] | e.registers.Vx[y]
}

func (e *Emulator) AND8(x byte, y byte) {
	e.registers.Vx[x] = e.registers.Vx[x] & e.registers.Vx[y]
}
func (e *Emulator) XOR8(x byte, y byte) {
	e.registers.Vx[x] = e.registers.Vx[x] ^ e.registers.Vx[y]
}
func (e *Emulator) ADD8(x byte, y byte) {
	e.registers.Vx[x] += e.registers.Vx[y]
	if e.registers.Vx[x] > 255 {
		e.registers.Vx[0xF] = 1
		return
	}
	e.registers.Vx[0xF] = 0
}
func (e *Emulator) SUB8(x byte, y byte) {
	e.registers.Vx[0xF] = 0
	if e.registers.Vx[x] > e.registers.Vx[y] {
		e.registers.Vx[0xF] = 1
	}
	e.registers.Vx[x] -= e.registers.Vx[y]
}
func (e *Emulator) SHR8(x byte) {
	last_significant_bit := x & byte(0x01)
	e.registers.Vx[0xF] = last_significant_bit
	e.registers.Vx[x] = e.registers.Vx[x] >> 1
}
func (e *Emulator) SUBN8(x byte, y byte) {
	e.registers.Vx[0xF] = 0
	if e.registers.Vx[y] > e.registers.Vx[x] {
		e.registers.Vx[0xF] = 1
	}
	e.registers.Vx[y] -= e.registers.Vx[x]
}
func (e *Emulator) SHL8(x byte) {
	most_significant_bit := x & byte(0xF0)
	e.registers.Vx[0xF] = most_significant_bit >> 3
	e.registers.Vx[x] = e.registers.Vx[x] << 1
}
func (e *Emulator) SNE9(x byte, y byte) {
	if e.registers.Vx[x] != e.registers.Vx[y] {
		e.registers.Pc += 2
	}
}
func (e *Emulator) LDA(adress uint16) {
	e.registers.I = adress
}
func (e *Emulator) JPB(adress uint16) {
	e.registers.Pc = uint16(e.registers.Vx[0]) + adress
}
func (e *Emulator) RNDC(x byte) {
	random_byte := byte(rand.Intn(255))
	e.registers.Vx[x] = e.registers.Vx[x] & random_byte
}
func (e *Emulator) DRWD(x byte, y byte, n uint8) {
	I := e.registers.I
	coord_x := e.registers.Vx[x]
	coord_y := e.registers.Vx[y]
	fmt.Println("Drawing: ", "x = ", x, "y = ", y, "n = ", n, " I = ", I, "Vx = ", coord_x, "Vy = ", coord_y)
	row := byte(0)
	cast_n := uint16(n)
	coord_x *= 8
	iter := I
	for iter < cast_n+I {
		fmt.Println("Ram Data: ", e.ram.data[I])
		e.display.pixels[coord_x+row*8+coord_y] ^= e.ram.data[iter]
		coord_x += 8
		iter++
	}

}

func (e *Emulator) SKPE(x byte) {

}
