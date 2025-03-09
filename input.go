package main

import "github.com/veandco/go-sdl2/sdl"

var (
	//KEYS
	kEscape, kF1, kF2, kF3, kL, kR, kU, kD bool
)

func EVENTS() {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch k := event.(type) {
		case sdl.QuitEvent:
			kEscape = true

		case sdl.KeyboardEvent:
			if event.GetType() == sdl.KEYDOWN {
				switch k.Keysym.Scancode {
				case sdl.SCANCODE_W, sdl.SCANCODE_UP:
					kU = true
				case sdl.SCANCODE_S, sdl.SCANCODE_DOWN:
					kD = true
				case sdl.SCANCODE_D, sdl.SCANCODE_RIGHT:
					kR = true
					if mouse.X > int32(pl.cnt.X) {
						pl.lr = false
					}
				case sdl.SCANCODE_A, sdl.SCANCODE_LEFT:
					kL = true
					if mouse.X < int32(pl.cnt.X) {
						pl.lr = true
					}
				case sdl.SCANCODE_ESCAPE:
					kEscape = true
				case sdl.SCANCODE_F1:
					kF1 = true
				case sdl.SCANCODE_F2:
					kF2 = true
				case sdl.SCANCODE_F3:
					kF3 = true
				}
			}
			if event.GetType() == sdl.KEYUP {
				switch k.Keysym.Scancode {
				case sdl.SCANCODE_W, sdl.SCANCODE_UP:
					kU = false
				case sdl.SCANCODE_S, sdl.SCANCODE_DOWN:
					kD = false
				case sdl.SCANCODE_D, sdl.SCANCODE_RIGHT:
					kR = false
				case sdl.SCANCODE_A, sdl.SCANCODE_LEFT:
					kL = false
				}
			}
		}
	}

	//KEYS
	if kEscape {
		EXIT()
	}
	if kF3 {
		addmsg("this is a message")
		kF3 = false
	}
	if kF2 {
		RESTART()
		kF2 = false
	}
	if kF1 {
		DEBUG = !DEBUG
		kF1 = false
	}

}

func MOUSE() {
	x, y, click := sdl.GetMouseState()
	mouse.X, mouse.Y = x, y

	if click == sdl.ButtonLMask && clickpauseL == 0 {
		clickpauseL = 4
		mouseL = true
	} else {
		mouseL = false
	}

	if click == sdl.ButtonRMask && clickpauseR == 0 {
		clickpauseR = 4
		mouseR = true
	} else {
		mouseR = false
	}

	if click == sdl.ButtonMMask {
		mouseM = true
	} else {
		mouseM = false
	}

	if mouseL && pl.projT == 0 {
		pl.projT = pl.projP
		mPLPROJ()
	}

	if mouse.X < int32(pl.cnt.X) && !pl.lr {
		pl.lr = true
	}
	if mouse.X > int32(pl.cnt.X) && pl.lr {
		pl.lr = false
	}
}
