package chip8

type Registers struct {
	Vx [16]byte

	Delay byte
	Sound byte

	I byte
	Sp byte
	Pc uint16
}

func NewRegisters() *Registers {
	return &Registers{
		Pc: 0x200,
	}
}
