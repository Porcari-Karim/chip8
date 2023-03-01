package chip8

import (
	"fmt"
)

func (e *Emulator) FetchOpCode() OpCode {
	pc := e.registers.Pc
	opcode := OpCode(e.ram.data[pc])
	opcode = opcode << 8 | OpCode(e.ram.data[pc+1])
	return OpCode(opcode)
}

func (e *Emulator) DecodeOpCode(opcode OpCode) {
	I := opcode & 0xF000
	fmt.Printf("0x%04x \n", opcode)
	switch(I){
	case 0x0000 :
		{
			break
		}
		
	case 0x1000 :
		{
			break
		}
		
	case 0x2000 :
		{
			break
		}
		
	case 0x3000 :
		{
			break
		}
		
	case 0x4000 :
		{
			break
		}
		
	case 0x5000 :
		{
			break
		}
		
	case 0x6000 :
		{
			break
		}
		
	case 0x7000 :
		{
			break
		}
		case 0x8000 :
		{
			break
		}
	case 0x9000 :
		{
			break
		}
	case 0xA000 :
		{
			break
		}
	}
}
 
func (e *Emulator) ClearScreen(){
	e.display.pixels = [32 * 64]byte{}
}

func (e * Emulator) GoTo(adress uint16){
	e.registers.Pc = adress
}

func (e *Emulator) Return(){
	e.registers.Pc = e.stack.Pop().(uint16)
}

func (e * Emulator) Call(adress uint16) {
	e.stack.Push(e.registers.Pc)
	e.registers.Pc = adress
}

func (e* Emulator) SE3(x byte, kk byte)  {
	if e.registers.Vx[x] == kk {
		e.registers.Pc += 2
	}
}
func (e* Emulator) SE4(x byte, kk byte)  { // Aka SNE
	if e.registers.Vx[x] != kk {
		e.registers.Pc += 2
	}
}
func (e* Emulator) SE5(x byte, y byte)  {
	if e.registers.Vx[x] == e.registers.Vx[y] {
		e.registers.Pc += 2
	}
}
func (e* Emulator) LD6(x byte, kk byte)  {
	e.registers.Vx[x] = kk
}
func (e* Emulator) ADD7(x byte, kk byte)  {
	e.registers.Vx[x] += kk
}
func (e* Emulator) LD8(x byte, y byte)  {
	e.registers.Vx[x] = e.registers.Vx[y]
}
func (e* Emulator) OR8(x byte, y byte)  {
	e.registers.Vx[x] = e.registers.Vx[x] | e.registers.Vx[y]
}
