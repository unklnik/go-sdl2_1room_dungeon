package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	WIN  *sdl.Window
	RNDR *sdl.Renderer

	winW, winH int32 = 1920, 1080

	displayMode sdl.DisplayMode

	SURF *sdl.Surface
	TEX  *sdl.Texture

	OFF bool

	setFPS = 60

	CNTR sdl.Point
)

func debug() { //MARK: DEBUG

	h := int32(font.Height())
	var x, y int32 = 4, 4
	txtXY("FPS "+fmt.Sprintf("%.0f", FPS)+" "+"tileSize "+fmt.Sprint(tileSize), x, y)
	y += h
	txtXY("mouse X "+fmt.Sprint(mouse.X)+" "+"mouse Y "+fmt.Sprint(mouse.Y), x, y)
	y += h
	txtXY("frames "+fmt.Sprint(frames)+" "+"debugNUM "+fmt.Sprint(debugNUM), x, y)
	y += h
	txtXY("kU "+fmt.Sprint(kU)+" "+"kD "+fmt.Sprint(kD), x, y)
	y += h
	txtXY("kR "+fmt.Sprint(kR)+" "+"kL "+fmt.Sprint(kL), x, y)
	y += h
	txtXY("pl.velx "+fmt.Sprintf("%.0f", pl.velx)+" "+"pl.vely "+fmt.Sprintf("%.0f", pl.vely), x, y)
	y += h
	txtXY("levW "+fmt.Sprint(levW)+" "+"levH "+fmt.Sprint(levH), x, y)
	y += h
	txtXY("mouseL "+fmt.Sprint(mouseL)+" "+"mouseR "+fmt.Sprint(mouseR), x, y)
	y += h
	txtXY("len(blokscntr) "+fmt.Sprint(len(blokscntr))+" "+"len(projpl) "+fmt.Sprint(len(projpl)), x, y)
	y += h
	txtXY("len(runeslev) "+fmt.Sprint(len(runeslev))+" "+"len(runes) "+fmt.Sprint(len(runes)), x, y)
	y += h
	txtXY("pl.bounce "+fmt.Sprint(pl.bounce)+" "+"pl.split "+fmt.Sprint(pl.split), x, y)
	y += h
	txtXY("pl.projT "+fmt.Sprint(pl.projT)+" "+"pl.grow "+fmt.Sprint(pl.grow), x, y)
	y += h
	txtXY("pl.faster "+fmt.Sprint(pl.faster)+" "+"pl.zigzag "+fmt.Sprint(pl.zigzag), x, y)
	y += h
	txtXY("Delta "+fmt.Sprintf("%.4f", Delta)+" "+"enmT "+fmt.Sprint(enmT), x, y)
	y += h

}

func PLAY() {

	INITIAL()

	for !OFF {

		B4()
		getDelta()
		DRAW()
		AFTER()
		UPDATE()
	}

}
func INITIAL() {

	sdl.SetHint(sdl.HINT_MOUSE_RELATIVE_SPEED_SCALE, "5")
	sdl.SetRelativeMouseMode(true)
	sdl.ShowCursor(sdl.DISABLE)
	ttf.Init()
	font, _ = ttf.OpenFont("Rubik-Medium.ttf", tx1)
	fontTextureSheet()
	mIMAGES()
	mLEV()
	mPL()
	mRUNES()

	mCHEST()
	/*
		runeslev = append(runeslev, runes[0])
		w := recResizeWidth(runestiles.r[0].W, runestiles.r[0].H, tileSize)
		runeslev[0].r = sdl.Rect{CNTR.X, CNTR.Y, w, tileSize}
		runeslev[0].r.X += tileSize
		runeslev[0].r.W = runeslev[0].r.W / 2
		runeslev[0].r.H = runeslev[0].r.H / 2
		runeslev[0].bounce = 1
	*/
}
func main() {

	//GET SCREEN SIZE
	WIN, _ = sdl.CreateWindow("", 0, 0, 0, 0, sdl.WINDOW_HIDDEN)
	defer WIN.Destroy()
	displayMode, _ = sdl.GetCurrentDisplayMode(0)
	if displayMode.W > 0 {
		winW = displayMode.W
	}
	if displayMode.H > 0 {
		winH = displayMode.H
	}
	WIN.Destroy()

	//winW, winH = 1920, 1080

	//CREATE WINDOW
	WIN, _ = sdl.CreateWindow("SDL2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winW, winH, sdl.WINDOW_ALLOW_HIGHDPI|sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)
	RNDR, _ = sdl.CreateRenderer(WIN, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_TARGETTEXTURE)

	RNDR.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	TEX, _ = RNDR.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, winW, winH)

	defer WIN.Destroy()
	defer RNDR.Destroy()
	defer TEX.Destroy()

	CNTR = sdl.Point{winW / 2, winH / 2}

	tileSize = winH / scale

	//FPS
	targetFPS = float64(setFPS)
	frameDelay = 1000 / targetFPS // Delay in milliseconds

	PLAY()
}
