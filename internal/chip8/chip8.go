package chip8

import (
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/golang-collections/collections/stack"
)

type Emulator struct {
	ram       *Ram
	registers *Registers
	display   *Display
	stack     *stack.Stack
	key       KeyState
	rom_size int
}

func NewInstance() *Emulator {
	return &Emulator{  
		NewRam(),
		NewRegisters(),
		NewDisplay(),
		stack.New(),
		KeyState{},
		0,
	}
}

func (e *Emulator) EmulateCycle() {
	if int(e.registers.Pc) < (e.rom_size + 0x200) {
		opcode := e.FetchOpCode()
		e.DecodeOpCode(opcode)
		e.registers.Pc +=2
	}

}

func (e *Emulator) LoadGame() (int64, error) {
	rom, err := os.ReadFile("assets/rom/rom.rom")
	if err != nil {
		return 0, err
	}
	fi , err := os.Stat("assets/rom/rom.rom")
	if err != nil {
		return 0, err
	}
	copy(e.ram.data[0x200:], rom)
	return fi.Size(), nil
}

func run() {
	emulator := NewInstance()
	cfg := pixelgl.WindowConfig{
		Title:  "Emulation Test",
		Bounds: pixel.R(0, 0, 640, 320),
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rom_size , err := emulator.LoadGame()
	if err != nil {
		panic(err)
	}
	//emulator.display.pixels[100] = 0b11111111
	emulator.rom_size = int(rom_size)
	for !window.Closed() {
		emulator.EmulateCycle()
		emulator.Draw(window)
		window.Update()
	}
}

func CreateAndRunEmulator() {
	pixelgl.Run(run)
}
