package main

import (
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (

	//TEXT
	standardCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;<=>?!#$%&'()*+,-./@[]^_`{|}~'\"' "

	font                    *ttf.Font
	tx1                     = 18
	fontchars               []FONTCH
	txSurf                  *sdl.Surface
	LineHeight, LetterSpace int32 = int32(tx1 + 4), 1

	//DEBUG
	DEBUG bool

	//TIMERS
	t15                      bool
	clickpauseL, clickpauseR int

	//FPS
	nowDelta, lastDelta uint64
	Delta               float32
	targetFPS           float64
	frameDelay          = float64(1000 / targetFPS) // Delay in milliseconds
	frameStart          time.Time
	frameTime           float64
	FPS                 float64
	frameCount, frames  int
	startTime           = time.Now()
)

type FONTCH struct {
	character string
	tex       *sdl.Texture
	rec       sdl.Rect
}

func RESTART() {
	floor = nil
	iw = nil
	extras = nil
	mLEV()
	//pl = PLAYER{}
	//mPL()
	chests = nil
	mCHEST()
	fxanm = nil

}

func B4() {
	frameStart = time.Now()
	EVENTS()
	RNDR.SetRenderTarget(TEX)
	RNDR.SetDrawColor(0, 0, 0, 0)
	RNDR.Clear()
}
func AFTER() {

	RNDR.SetRenderTarget(nil)

	RNDR.SetDrawColor(0, 0, 0, 255)
	RNDR.Clear()

	RNDR.Copy(TEX, nil, &sdl.Rect{0, 0, winW, winH})

	RNDR.Present()

	//FPS
	frames++
	frameCount++

	frameTime = float64(time.Since(frameStart).Milliseconds())
	if frameTime < frameDelay {
		sdl.Delay(uint32(frameDelay - frameTime)) // Wait to achieve target FPS
	}
	if time.Since(startTime).Seconds() >= 1 {
		FPS = float64(frameCount) / time.Since(startTime).Seconds()
		frameCount = 0
		startTime = time.Now()
	}

}

func EXIT() {
	//surface.Free()
	//texture.Destroy()
	loadsurf.Free()
	loadtex.Destroy()
	RNDR.Destroy()
	WIN.Destroy()
	OFF = true
}

func getDelta() {
	tickT := sdl.GetTicks64()
	Delta = float32(tickT-lastDelta) * 0.001
	lastDelta = tickT
}

// TEXT
func txtXY(txt string, x, y int32) {
	t := strings.Split(txt, "")
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontchars); j++ {
			if t[i] == fontchars[j].character {
				RNDR.Copy(fontchars[j].tex, &fontchars[j].rec, &sdl.Rect{x, y, fontchars[j].rec.W, fontchars[j].rec.H})
				x += fontchars[j].rec.W + LetterSpace
				break
			}
		}
	}
}
func txtXYCOLOR(txt string, x, y int32, c sdl.Color) {
	t := strings.Split(txt, "")
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontchars); j++ {
			if t[i] == fontchars[j].character {
				fontchars[j].tex.SetColorMod(c.R, c.G, c.B)
				RNDR.Copy(fontchars[j].tex, &fontchars[j].rec, &sdl.Rect{x, y, fontchars[j].rec.W, fontchars[j].rec.H})
				fontchars[j].tex.SetColorMod(255, 255, 255)
				x += fontchars[j].rec.W + LetterSpace
				break
			}
		}
	}
}

func fontTexRevert(f []FONTCH) []FONTCH {
	for i := 0; i < len(f); i++ {
		fontchars[i].tex.SetColorMod(255, 255, 255)
	}
	return f
}
func fontTextureSheet() {
	t := strings.Split(standardCharacters, "")
	for i := 0; i < len(t); i++ {
		fontchars = append(fontchars, fontCharCreate(t[i], 0))
	}
}

func fontCharCreate(singleCharacter string, fontNum int) FONTCH {

	var w, h int

	switch fontNum {
	case 0: //FONT DEFAULT
		txSurf, _ = font.RenderUTF8Blended(singleCharacter, WHITEǁ3())
		defer txSurf.Free()
		w, h, _ = font.SizeUTF8(singleCharacter)
	}

	f := FONTCH{}
	f.tex, _ = RNDR.CreateTextureFromSurface(txSurf)
	f.rec = sdl.Rect{0, 0, int32(w), int32(h)}
	f.character = singleCharacter
	return f
}
