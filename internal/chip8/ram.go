package chip8

type Ram struct {
	data []byte
	size int
}

func NewRam() *Ram {
	return &Ram{make([]byte, 4096, 4096), 0}
}

func (e *Emulator) Populate(data []byte, index int) {
	temp := append(e.ram.data[:index], data...)
	e.ram.data = append(temp, e.ram.data[len(data)+index:]...)

}

func (e *Emulator) GetData(index int, size int) []byte {
	return e.ram.data[index : index+size]
}
