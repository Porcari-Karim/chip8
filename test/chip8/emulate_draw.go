package main

import (
	"fmt"

	"github.com/Porcari-Karim/chip8/internal/chip8"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func test_run() {
	emulator := chip8.NewInstance()
	cfg := pixelgl.WindowConfig{
		Title:  "Emulation Test",
		Bounds: pixel.R(0, 0, 640, 320),
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rom_size, err := emulator.LoadGame()
	if err != nil {
		fmt.Println(rom_size)
		panic(err)
	}
	//emulator.display.pixels[100] = 0b11111111
	line := make([]byte, 5)
	line[0] = 0b11111111
	line[1] = 0b10011001
	line[2] = 0b11111111
	line[3] = 0b11111111
	line[4] = 0b11011011
	emulator.Populate(line, 4000)
	fmt.Println(emulator.GetData(4000, 5))
	registers_ptr := emulator.GetRegisters()
	registers_ptr.I = 4000
	registers_ptr.Vx[0] = 0
	registers_ptr.Vx[1] = 0
	emulator.DRWD(0, 1, 5)

	for !window.Closed() {
		//emulator.EmulateCycle()
		emulator.Draw(window)
		window.Update()
	}
}

func main() {
	pixelgl.Run(test_run)
}
