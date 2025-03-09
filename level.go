package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	blokscntr                     []sdl.Rect
	floor, iw, extras             []OBJ
	scale                               = int32(20)
	tileSize                            = int32(32)
	levWnum, levHnum              int32 = 24, 14
	levW, levH, levWin, levHin    int32
	levrecin, levrec, levrecfloor sdl.Rect

	runeslev []RUNE
)

type OBJ struct {
	r, ir        sdl.Rect
	solid, outer bool
	ro           float64
	im           IM
}

func mEXTRAS() {

	w2, h2 := levrecin.W, levrecin.H
	x, y := levrecin.X, levrecin.Y

	//RUG
	siz := tileSize * 3
	o := OBJ{}
	o.ro = RF64(0, 360)
	x += RI32(0, int(w2-siz))
	y += RI32(0, int(h2-siz))
	o.r = sdl.Rect{x, y, siz, siz}
	o.im.r = rugtiles.r[RINT(0, len(rugtiles.r))]
	o.im.tex = rugtiles.tex
	extras = append(extras, o)

	countbreak := 100
	num := RINT(2, 7)
	for num > 0 && countbreak > 0 {
		x, y = levrecin.X, levrecin.Y
		siz = tileSize / 2
		o = OBJ{}
		x += RI32(0, int(w2-siz))
		y += RI32(0, int(h2-siz))
		choose := RINT(3, 28)
		if choose > 12 && choose != 22 && choose != 27 {
			siz += siz / 4
		}
		if choose == 27 {
			siz += siz
		}
		o.im = extrasIM[choose]
		o.r = sdl.Rect{x, y, siz, recResizeHeight(o.im.r.W, o.im.r.H, siz)}
		if checkaddextra(o.r) {
			extras = append(extras, o)
			num--
		}
		countbreak--
	}

}
func checkaddextra(r sdl.Rect) bool {
	canadd := true
	if len(iw) > 0 {
		for i := range iw {
			if RecsIntersect(iw[i].r, r) {
				canadd = false
				break
			}
		}
	}
	if canadd {
		for i := range blokscntr {
			if RecsIntersect(blokscntr[i], r) {
				canadd = false
				break
			}
		}
	}

	if canadd {
		if len(extras) > 0 {
			for i := range extras {
				if RecsIntersect(extras[i].r, r) {
					canadd = false
					break
				}
			}
		}
	}

	return canadd
}

func mLEV() {

	//MAKE GRID
	a := levWnum * levHnum
	w := levWnum * tileSize
	x := CNTR.X - w/2
	ox := x
	h := levHnum * tileSize
	y := CNTR.Y - h/2

	c := 0
	for a > 0 {

		o := OBJ{}
		o.r = sdl.Rect{x, y, tileSize, tileSize}
		if len(floor) < int(levWnum) {
			o.solid = true
		} else if len(floor)%int(levWnum) == 0 {
			o.solid = true
		} else if len(floor) > (int(levWnum*levHnum - levWnum)) {
			o.solid = true
		} else if (len(floor)+1)%int(levWnum) == 0 {
			o.solid = true
		}
		floor = append(floor, o)

		x += tileSize
		c++
		if c == int(levWnum) {
			c = 0
			x = ox
			y += tileSize
		} //
		a--
	}

	//BOUNDARY
	levW = floor[levWnum-1].r.X + floor[levWnum-1].r.W - floor[0].r.X
	levH = floor[len(floor)-int(levWnum)].r.Y + floor[0].r.H - floor[0].r.Y
	levWin = levW - tileSize*5
	levHin = levH - tileSize*5
	levrecin = sdl.Rect{CNTR.X - levWin/2, CNTR.Y - levHin/2, levWin, levHin}
	levrec = sdl.Rect{CNTR.X - levW/2, CNTR.Y - levH/2, levW, levH}
	levrecfloor = levrec
	levrecfloor.X += tileSize
	levrecfloor.Y += tileSize
	levrecfloor.W -= tileSize * 2
	levrecfloor.H -= tileSize * 2

	blokscntr = append(blokscntr, floor[130].r, floor[131].r, floor[132].r, floor[133].r, floor[154].r, floor[155].r, floor[156].r, floor[157].r, floor[178].r, floor[179].r, floor[180].r, floor[181].r, floor[202].r, floor[203].r, floor[204].r, floor[205].r)

	//INNERWALLS INTERIOR
	mINNWALLS()
	mEXTRAS()

	//FLOOR TILES
	for i := range floor {
		if !floor[i].solid {
			floor[i].im.tex = floortiles.tex
			switch RINT(0, 4) {
			case 0:
				floor[i].im.r = floortiles.r[0]
			case 1:
				floor[i].im.r = floortiles.r[1]

			case 2:
				floor[i].im.r = floortiles.r[2]

			case 3:
				floor[i].im.r = floortiles.r[3]
			}
		}
	}

	//WALL TILES
	//CORNERS TOP
	floor[1].outer = true
	floor[1].im.r = walltiles.r[0]
	floor[1].im.tex = walltiles.tex
	newH := recResizeHeight(floor[1].im.r.W, floor[1].im.r.H, tileSize*2)
	floor[1].ir = sdl.Rect{floor[1].r.X - tileSize, floor[1].r.Y + floor[1].r.H, tileSize * 2, newH}
	floor[1].ir.Y -= floor[1].ir.H

	floor[levWnum-1].outer = true
	floor[levWnum-1].im.r = walltiles.r[3]
	floor[levWnum-1].im.tex = walltiles.tex
	newH = recResizeHeight(floor[levWnum-1].im.r.W, floor[levWnum-1].im.r.H, tileSize*2)
	floor[levWnum-1].ir = sdl.Rect{floor[levWnum-1].r.X - tileSize, floor[levWnum-1].r.Y + floor[levWnum-1].r.H, tileSize * 2, newH}
	floor[levWnum-1].ir.Y -= floor[levWnum-1].ir.H

	//TOP WALL
	for i := 2; i < int(levWnum-1); i++ {
		floor[i].outer = true
		floor[i].im.tex = walltiles.tex
		floor[i].im.r = walltiles.r[2]
		if Roll6() > 4 {
			floor[i].im.r = walltiles.r[1]
		}
		newH = recResizeHeight(floor[i].im.r.W, floor[i].im.r.H, tileSize)
		floor[i].ir = sdl.Rect{floor[i].r.X - tileSize, floor[i].r.Y + floor[i].r.H, tileSize, newH}
		floor[i].ir.Y -= floor[i].ir.H
	}

	//DOORS
	mDOORS(newH)

	//CORNERS BOTTOM
	num := int(levWnum*levHnum - levWnum)
	floor[num].outer = true
	floor[num].im.r = walltiles.r[4]
	floor[num].im.tex = walltiles.tex
	newH = recResizeHeight(floor[num].im.r.W, floor[num].im.r.H, tileSize*2)
	floor[num].ir = sdl.Rect{floor[num].r.X, floor[num].r.Y + floor[num].r.H, tileSize * 2, newH}
	floor[num].ir.Y -= floor[num].ir.H

	num = int(levWnum*levHnum - 1)
	floor[num].outer = true
	floor[num].im.r = walltiles.r[7]
	floor[num].im.tex = walltiles.tex
	newH = recResizeHeight(floor[num].im.r.W, floor[num].im.r.H, tileSize*2)
	floor[num].ir = sdl.Rect{floor[num].r.X - tileSize, floor[num].r.Y + floor[num].r.H, tileSize * 2, newH}
	floor[num].ir.Y -= floor[num].ir.H

	//SIDE WALLS
	for i := int(levWnum); i < int(levWnum*levHnum-(levWnum*2)); i++ {
		if i%int(levWnum) == 0 {
			floor[i].im.tex = walltiles.tex
			floor[i].im.r = walltiles.r[9]
			if Roll6() > 4 {
				floor[i].im.r = walltiles.r[8]
			}
		}
		if (i+1)%int(levWnum) == 0 {
			floor[i].im.tex = walltiles.tex
			floor[i].im.r = walltiles.r[11]
			if Roll6() > 4 {
				floor[i].im.r = walltiles.r[10]
			}
		}
	}

	for i := (levWnum * levHnum) - (levWnum - 2); i < (levWnum*levHnum)-2; i++ {
		floor[i].im.tex = walltiles.tex
		floor[i].im.r = walltiles.r[13]
		if Roll6() > 4 {
			floor[i].im.r = walltiles.r[12]
		}
	}

	mENEMIES()
}
func mINNWALLS() {

	selwall := inwalltiles1
	switch RINT(0, 3) {
	case 1:
		selwall = inwalltiles2
	case 2:
		selwall = inwalltiles3
	}

	w2, h2 := levrecin.W, levrecin.H
	siz := (tileSize / 5) * 3

	num := RINT(3, 9)

	for num > 0 {
		choose := RINT(0, 15)
		//choose = 14
		x, y := levrecin.X, levrecin.Y

		switch choose {
		case 14: //TRIANGLE UP
			x += RI32(0, int(w2-siz*7))
			y += RI32(int(siz*4), int(h2-siz*2))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 13: //TRIANGLE UP
			x += RI32(0, int(w2-siz*7))
			y += RI32(int(siz*4), int(h2-siz*2))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 12: //LR 2 BLOK HORIZ
			x += RI32(0, int(w2-siz*6))
			y += RI32(0, int(h2-siz*6))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x, y, siz, siz}
			o.r.Y += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x, y, siz, siz}
			o.r.Y += siz * 4
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 11: //UP DOWN DIAGONAL X3
			x += RI32(0, int(w2-siz*6))
			y += RI32(0, int(h2-siz*2))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 10: //5 SINGLE BLOKS HORIZ
			x += RI32(0, int(w2-siz*15))
			y += RI32(0, int(h2-siz))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 9: //3 SINGLE BLOKS VERT
			x += RI32(0, int(w2-siz))
			y += RI32(0, int(h2-siz*9))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 8: //3 SINGLE BLOKS HORIZ
			x += RI32(0, int(w2-siz*9))
			y += RI32(0, int(h2-siz))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 3
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 7: //3X3 CROSS X2 VERT
			x += RI32(0, int(w2-siz*3))
			y += RI32(0, int(h2-siz*10))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x + siz, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.r.Y += siz
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			//CROSS 2
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz * 5
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.r.Y += siz * 5
			o.r.Y += siz
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 6: //3X3 CROSS X3 HORIZ
			x += RI32(0, int(w2-siz*15))
			y += RI32(0, int(h2-siz*3))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x + siz, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.r.Y += siz
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			//CROSS 2
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 5
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.r.X += siz * 5
			o.r.Y += siz
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}

			//CROSS 3
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 10
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.r.X += siz * 10
			o.r.Y += siz
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 5: //3X3 CROSS
			x += RI32(0, int(w2-siz*3))
			y += RI32(0, int(h2-siz*3))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x + siz, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r = sdl.Rect{x + siz, y, siz, siz}
			o.r.Y += siz
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 4: //3X3 DIAGONAL CROSS
			x += RI32(0, int(w2-siz*3))
			y += RI32(0, int(h2-siz*3))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.r = sdl.Rect{x, y, siz, siz}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 3: //3X3 BLOK
			x += RI32(0, int(w2-siz*3))
			y += RI32(0, int(h2-siz*3))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 2: //3X2 BLOK
			x += RI32(0, int(w2-siz*3))
			y += RI32(0, int(h2-siz*2))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 1: //4 BLOK
			x += RI32(0, int(w2-siz*2))
			y += RI32(0, int(h2-siz*2))
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.Y += siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		case 0: //CORNERS DIAGONALS
			o := OBJ{}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.im.tex = selwall.tex
			o.r = sdl.Rect{x, y, siz, siz}
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			o.r.Y += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			o.r.Y += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			o.r.Y += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.r = sdl.Rect{x, y, siz, siz}
			o.r.X += levrecin.W - o.r.W
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			o.r.Y += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			o.r.Y += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			o.r.Y += siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.r = sdl.Rect{x, y, siz, siz}
			o.r.Y += levrecin.H - o.r.W
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			o.r.Y -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			o.r.Y -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X += siz * 2
			o.r.Y -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.r = sdl.Rect{x, y, siz, siz}
			o.r.Y += levrecin.H - o.r.W
			o.r.X += levrecin.W - o.r.W
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			o.r.Y -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			o.r.Y -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
			o.im.r = selwall.r[RINT(0, len(inwalltiles1.r))]
			o.r.X -= siz * 2
			o.r.Y -= siz * 2
			if checkaddinnwall(o.r) {
				iw = append(iw, o)
			}
		}
		num--
	}
}

func checkaddinnwall(r sdl.Rect) bool {
	canadd := true
	if len(iw) > 0 {
		for i := range iw {
			if RecsIntersect(iw[i].r, r) {
				canadd = false
				break
			}
		}
	}
	if canadd {
		for i := range blokscntr {
			if RecsIntersect(blokscntr[i], r) {
				canadd = false
				break
			}
		}
	}
	return canadd
}
func mDOORS(h int32) {
	h -= tileSize + tileSize/4
	newW := recResizeWidth(doortiles.r[0].W, doortiles.r[0].H, h)
	y := floor[1].ir.Y + floor[1].ir.H - h
	for i := range 3 {
		d := IM{}
		d.tex = doortiles.tex
		d.r = doortiles.r[RINT(0, len(doortiles.r))]
		switch i {
		case 0:
			d.dr = sdl.Rect{CNTR.X - newW/2, y, newW, h}
		case 1:
			x := CNTR.X - levW/4
			d.dr = sdl.Rect{x - newW/2, y, newW, h}
		case 2:
			x := CNTR.X + levW/4
			d.dr = sdl.Rect{x - newW/2, y, newW, h}
		}
		doorIM = append(doorIM, d)
	}
}
