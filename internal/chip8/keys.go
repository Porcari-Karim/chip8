package chip8

import (
	"github.com/faiface/pixel/pixelgl"
)

type KeyState struct {
	window *pixelgl.Window
}

func (k *KeyState) CheckKeyState(key pixelgl.Button) bool {
	return k.window.JustPressed(key)
}
