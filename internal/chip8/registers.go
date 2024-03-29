package chip8

type Registers struct {
	Vx [16]byte

	Delay byte
	Sound byte

	I  uint16
	Sp byte
	Pc uint16
}

func NewRegisters() *Registers {
	return &Registers{
		Pc: 0x200,
	}
}

func (e *Emulator) GetRegisters() *Registers {
	return e.registers
}
