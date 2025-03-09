package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	loadtex                 *sdl.Texture
	loadsurf                *sdl.Surface
	etcIM, doorIM, extrasIM []IM
	targangl                float64

	floortiles, walltiles, wandtiles, doortiles, inwalltiles1, inwalltiles2, inwalltiles3, rugtiles, runestiles TILESHEET

	enmspawn, chestopen, fireball, endprojanm, plA, plD, plI, plI2, plInvis, plJ, plR ANIM
)

type TILESHEET struct {
	tex *sdl.Texture
	r   []sdl.Rect
}
type ANIM struct {
	r, dr                  sdl.Rect
	frames, framecount, xl int32
	tex                    *sdl.Texture
	onoff, once            bool
	spd, num               int
}

type IM struct {
	tex   *sdl.Texture
	r, dr sdl.Rect
}

func mIMAGES() {

	loadsurf, _ = img.Load("im/floor.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	floortiles = mTILES(loadtex, 40, 40, 4)
	loadsurf, _ = img.Load("im/outwall.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	walltiles.tex = loadtex
	walltiles.r = append(walltiles.r, sdl.Rect{40, 41, 80, 120})  //0 CORNER UL
	walltiles.r = append(walltiles.r, sdl.Rect{160, 41, 40, 120}) //1 WALL U
	walltiles.r = append(walltiles.r, sdl.Rect{240, 41, 40, 120}) //2 WALL U2
	walltiles.r = append(walltiles.r, sdl.Rect{320, 41, 80, 120}) //3 CORNER UR
	walltiles.r = append(walltiles.r, sdl.Rect{40, 361, 80, 80})  //4 CORNER BL
	walltiles.r = append(walltiles.r, sdl.Rect{160, 401, 40, 40}) //5 WALL B
	walltiles.r = append(walltiles.r, sdl.Rect{240, 401, 40, 40}) //6 WALL B2
	walltiles.r = append(walltiles.r, sdl.Rect{320, 361, 80, 80}) //7 CORNER BR
	walltiles.r = append(walltiles.r, sdl.Rect{40, 201, 40, 40})  //8 WALL L
	walltiles.r = append(walltiles.r, sdl.Rect{40, 281, 40, 40})  //9 WALL L2
	walltiles.r = append(walltiles.r, sdl.Rect{360, 201, 40, 40}) //10 WALL R
	walltiles.r = append(walltiles.r, sdl.Rect{360, 281, 40, 40}) //11 WALL R2
	walltiles.r = append(walltiles.r, sdl.Rect{160, 401, 40, 40}) //12 WALL B
	walltiles.r = append(walltiles.r, sdl.Rect{240, 401, 40, 40}) //13 WALL B2

	loadsurf, _ = img.Load("im/player.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	plJ = mANIM(loadtex, 0, 160, 32, 32, 8)
	plR = mANIM(loadtex, 0, 96, 32, 32, 8)
	plD = mANIM(loadtex, 0, 224, 32, 32, 8)
	plA = mANIM(loadtex, 0, 256, 32, 32, 8)
	plI = mANIM(loadtex, 0, 32, 32, 32, 2)
	plI2 = mANIM(loadtex, 0, 0, 32, 32, 2)
	plInvis = mANIM(loadtex, 0, 192, 32, 32, 4)

	i := IM{}
	i.r = sdl.Rect{218, 6, 24, 24}
	i.tex = loadtex
	etcIM = append(etcIM, i) //0 PL SHADOW

	loadsurf, _ = img.Load("im/wands.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	wandtiles = mTILES(loadtex, 128, 128, 10)
	mWANDS()

	loadsurf, _ = img.Load("im/etc.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	i = IM{}
	i.r = sdl.Rect{0, 0, 16, 16}
	i.tex = loadtex
	etcIM = append(etcIM, i) //1 TARGET
	i.r = sdl.Rect{18, 1, 18, 18}
	etcIM = append(etcIM, i) //2 FLAME
	i.r = sdl.Rect{41, 3, 14, 14}
	etcIM = append(etcIM, i) //3 SHURIKEN
	i.r = sdl.Rect{59, 2, 18, 18}
	etcIM = append(etcIM, i) //4 LIGHTNING
	i.r = sdl.Rect{79, 3, 16, 16}
	etcIM = append(etcIM, i) //5 ARROW
	i.r = sdl.Rect{98, 3, 13, 13}
	etcIM = append(etcIM, i) //6 SPIKE
	i.r = sdl.Rect{113, 3, 15, 15}
	etcIM = append(etcIM, i) //7 ROCKET
	i.r = sdl.Rect{132, 2, 16, 16}
	etcIM = append(etcIM, i) //8 SHOTGUN
	i.r = sdl.Rect{150, 4, 14, 14}
	etcIM = append(etcIM, i) //9 CARROT
	i.r = sdl.Rect{167, 4, 16, 16}
	etcIM = append(etcIM, i) //10 SHIELD
	i.r = sdl.Rect{3, 24, 62, 62}
	etcIM = append(etcIM, i) //11 BORDER
	i.r = sdl.Rect{68, 24, 62, 62}
	etcIM = append(etcIM, i) //12 BORDER2

	loadsurf, _ = img.Load("im/doors.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	doortiles = mTILES(loadtex, 256, 256, 12)

	loadsurf, _ = img.Load("im/runes.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	runestiles = mTILES(loadtex, 54, 60, 35)

	loadsurf, _ = img.Load("im/inwalls1.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	inwalltiles1 = mTILES(loadtex, 16, 16, 48)

	loadsurf, _ = img.Load("im/inwalls2.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	inwalltiles2 = mTILES(loadtex, 16, 16, 48)

	loadsurf, _ = img.Load("im/inwalls3.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	inwalltiles3 = mTILES(loadtex, 16, 16, 48)

	loadsurf, _ = img.Load("im/rugs.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	rugtiles = mTILES(loadtex, 128, 128, 11)

	loadsurf, _ = img.Load("im/proj.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	fireball = mANIM(loadtex, 0, 0, 64, 64, 12)
	endprojanm = mANIM(loadtex, 0, 64, 64, 64, 18)
	chestopen = mANIM(loadtex, 0, 128, 64, 64, 14)
	enmspawn = mANIM(loadtex, 0, 192, 64, 64, 7)

	loadsurf, _ = img.Load("im/interior.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	i = IM{}
	i.r = sdl.Rect{0, 0, 96, 96}
	i.tex = loadtex
	extrasIM = append(extrasIM, i) //0 RUG
	i.r = sdl.Rect{96, 0, 96, 96}
	extrasIM = append(extrasIM, i) //1 RUG2
	i.r = sdl.Rect{192, 0, 96, 96}
	extrasIM = append(extrasIM, i) //2 RUG3
	i.r = sdl.Rect{292, 0, 42, 66}
	extrasIM = append(extrasIM, i) //3 PLANT
	i.r = sdl.Rect{334, 0, 42, 66}
	extrasIM = append(extrasIM, i) //4 PLANT2
	i.r = sdl.Rect{376, 0, 47, 75}
	extrasIM = append(extrasIM, i) //5 PLANT3
	i.r = sdl.Rect{423, 0, 47, 75}
	extrasIM = append(extrasIM, i) //6 PLANT4
	i.r = sdl.Rect{470, 0, 47, 75}
	extrasIM = append(extrasIM, i) //7 PLANT5
	i.r = sdl.Rect{517, 0, 47, 75}
	extrasIM = append(extrasIM, i) //8 PLANT6
	i.r = sdl.Rect{564, 0, 30, 36}
	extrasIM = append(extrasIM, i) //9 PLANT7
	i.r = sdl.Rect{594, 0, 30, 39}
	extrasIM = append(extrasIM, i) //10 PLANT8
	i.r = sdl.Rect{564, 36, 30, 36}
	extrasIM = append(extrasIM, i) //11 PLANT9
	i.r = sdl.Rect{594, 39, 36, 42}
	extrasIM = append(extrasIM, i) //12 POT
	i.r = sdl.Rect{630, 1, 16, 14}
	extrasIM = append(extrasIM, i) //13 ROCK
	i.r = sdl.Rect{630, 16, 18, 15}
	extrasIM = append(extrasIM, i) //14 ROCK2
	i.r = sdl.Rect{632, 32, 28, 21}
	extrasIM = append(extrasIM, i) //15 ROCK3
	i.r = sdl.Rect{634, 56, 34, 35}
	extrasIM = append(extrasIM, i) //16 ROCK4
	i.r = sdl.Rect{654, 0, 15, 32}
	extrasIM = append(extrasIM, i) //17 ROCK5
	i.r = sdl.Rect{673, 63, 33, 25}
	extrasIM = append(extrasIM, i) //18 ROCK6
	i.r = sdl.Rect{669, 1, 19, 24}
	extrasIM = append(extrasIM, i) //19 DUMMY
	i.r = sdl.Rect{690, 1, 18, 24}
	extrasIM = append(extrasIM, i) //20 DUMMY
	i.r = sdl.Rect{708, 1, 16, 24}
	extrasIM = append(extrasIM, i) //21 POLE
	i.r = sdl.Rect{726, 1, 12, 24}
	extrasIM = append(extrasIM, i) //22 CLOCK
	i.r = sdl.Rect{664, 32, 24, 22}
	extrasIM = append(extrasIM, i) //23 FLAG
	i.r = sdl.Rect{690, 31, 24, 22}
	extrasIM = append(extrasIM, i) //24 FLAG2
	i.r = sdl.Rect{743, 2, 32, 32}
	extrasIM = append(extrasIM, i) //25 BARREL
	i.r = sdl.Rect{778, 2, 32, 32}
	extrasIM = append(extrasIM, i) //26 CRATE
	i.r = sdl.Rect{715, 39, 69, 37}
	extrasIM = append(extrasIM, i) //27 STONE SIGN
	i.r = sdl.Rect{787, 43, 48, 39}
	extrasIM = append(extrasIM, i) //28 CHEST
	i.r = sdl.Rect{839, 47, 61, 48}
	extrasIM = append(extrasIM, i) //29 CHEST2
	i.r = sdl.Rect{836, 4, 48, 39}
	extrasIM = append(extrasIM, i) //30 CHEST3
	i.r = sdl.Rect{897, 5, 54, 48}
	extrasIM = append(extrasIM, i) //31 CHEST4
	i.r = sdl.Rect{959, 4, 42, 65}
	extrasIM = append(extrasIM, i) //32 ALTAR
	i.r = sdl.Rect{1003, 4, 48, 67}
	extrasIM = append(extrasIM, i) //33 ALTAR2
	i.r = sdl.Rect{1055, 2, 45, 78}
	extrasIM = append(extrasIM, i) //34 ALTAR3
	i.r = sdl.Rect{1101, 2, 50, 75}
	extrasIM = append(extrasIM, i) //35 ALTAR4
	i.r = sdl.Rect{1152, 2, 46, 57}
	extrasIM = append(extrasIM, i) //36 ALTAR5
}
func mANIMONCE(a ANIM, spd int, r sdl.Rect, num int) ANIM {
	a.spd = spd
	a.dr = r
	a.once = true
	a.num = num
	return a
}
func mANIM(t *sdl.Texture, x, y, w, h, frames int32) ANIM {
	a := ANIM{}
	a.tex = t
	a.r = sdl.Rect{x, y, w, h}
	a.xl = x
	a.frames = frames

	return a
}
func mTILES(t *sdl.Texture, w, h, num int32) TILESHEET {

	ts := TILESHEET{}
	ts.tex = t
	var x, y int32 = 0, 0

	for num > 0 {
		r := sdl.Rect{x, y, w, h}
		ts.r = append(ts.r, r)
		x += w
		num--
	}

	return ts
}
func mTILESXY(t *sdl.Texture, x, y, w, h, num int32) TILESHEET {

	ts := TILESHEET{}
	ts.tex = t

	for num > 0 {
		r := sdl.Rect{x, y, w, h}
		ts.r = append(ts.r, r)
		x += w
		num--
	}

	return ts
}
