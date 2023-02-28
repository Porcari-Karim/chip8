package chip8


func (e *Emulator) FetchOpCode() OpCode {
	pc := e.registers.Pc
	opcode := e.ram.data[pc]<<8 | e.ram.data[pc+1]
	return OpCode(opcode)
}

func (e *Emulator) DecodeOpCode() {

}
