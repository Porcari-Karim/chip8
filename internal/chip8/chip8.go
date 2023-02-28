package chip8

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"os"
)

type Emulator struct {
	ram       *Ram
	registers *Registers
	display   *Display
	stack     Stack
	key       KeyState
}

func NewInstance() *Emulator {
	return &Emulator{
		NewRam(),
		NewRegisters(),
		NewDisplay(),
		NewStack(),
		KeyState{},
	}
}

func (e *Emulator) EmulateCycle() {

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
	emulator.registers.Pc = 0x200
	for _, b := range emulator.ram.data[0x200:0x200 + rom_size] {
		fmt.Printf("0x%04x \n", b)
	}
	for !window.Closed() {
		emulator.EmulateCycle()
		window.Update()
	}
}

func CreateAndRunEmulator() {
	pixelgl.Run(run)
}
