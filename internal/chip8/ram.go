package chip8

type Ram struct {
	data []byte
	size int
}

func NewRam() *Ram {
	return &Ram{make([]byte, 4096, 4096), 0}
}
