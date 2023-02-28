package main

import (
	"fmt"

	"github.com/Porcari-Karim/chip8/internal/chip8"
)

func main() {
	fmt.Println("Simple Chip8 Emulator made by Porcari Karim")
	chip8.CreateAndRunEmulator()
}
