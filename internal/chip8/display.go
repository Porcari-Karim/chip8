package chip8

type Display struct {
	pixels [32 * 64]byte //
}

func NewDisplay() *Display {
	return &Display{}
}
