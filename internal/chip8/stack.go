package chip8

type Stack []uint16

func NewStack() Stack {
	return make([]uint16, 16, 16)
}
