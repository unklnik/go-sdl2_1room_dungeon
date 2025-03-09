package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	pZERO    = sdl.Point{0, 0}
	pZEROF32 = sdl.FPoint{0, 0}

	debugNUM int
)

// DRAW TO SCREEN
func DRAW() {

	dLEV()
	if len(runeslev) > 0 {
		dRUNES()
	}
	if len(chests) > 0 {
		dCHESTS()
	}
	if len(enm) > 0 {
		dENM()
	}
	if len(fxanm) > 0 {
		dFXANM()
	}

	dPL()
	dUI()

	if len(msgs) > 0 {
		dMSG()
	}

	if DEBUG {
		dRECFLINESIZE(pl.r, 2, ORANGEǁ3())
		dRECFLINESIZE(pl.collisr, 2, ORANGEǁ3())

		for i := range floor {
			if floor[i].solid {
				dRECFILLALPHA(floor[i].r, MAGENTAǁ3(), 50)
			}
			dRECLINE(floor[i].r, MAGENTAǁ3())
			txtXY(fmt.Sprint(i), floor[i].r.X+4, floor[i].r.Y+4)

		}

		dRECLINESIZECOLOR(levrecfloor, 4, DARKORANGEǁ3())
		dRECLINESIZECOLOR(levrecin, 4, DARKORANGEǁ3())
		dRECLINESIZE(levrec, 4)

		dSQCNTRF(enmspawnpoints[0], float32(tileSize/4))
		dSQCNTRF(enmspawnpoints[1], float32(tileSize/4))
		dSQCNTRF(enmspawnpoints[2], float32(tileSize/4))
		dSQCNTRF(enmspawnpoints[3], float32(tileSize/4))

		debug()

	}

}

// ENEMIES
func dENM() {

	for i := range enm {
		dFRECFILL(enm[i].r, MAGENTAǁ3())
	}

}

// MSG
func dMSG() {
	counting := false
	clear := false
	for i := range msgs {
		if !msgs[i].onoff && !counting {
			counting = true
			txtXY(msgs[i].t, levrec.X, levrec.Y+levrec.H)
			msgs[i].timer--
			if msgs[i].timer == 0 {
				msgs[i].onoff = true
				clear = true
			}
		}
	}

	if clear {
		for i := 0; i < len(msgs); i++ {
			msgs = remMSG(msgs, i)
		}
	}
}

// CHESTS
func dCHESTS() {
	for i := range chests {
		if !chests[i].open {
			dIMSHADOW(chests[i].im, chests[i].r, 4, true)
			if RecsIntersect(frec2rec(pl.r), chests[i].r) {
				chests[i].open = true
				fxanm = append(fxanm, mANIMONCE(chestopen, 2, resizerec(chests[i].r, 4), 1))
			}
		}
	}
}

// FX ANIM
func dFXANM() {
	for i := range fxanm {
		if !fxanm[i].onoff {
			if fxanm[i].once {
				fxanm[i], fxanm[i].onoff = dANIMONCE(fxanm[i], fxanm[i].dr, fxanm[i].spd, fxanm[i].onoff)
			} else {
				fxanm[i] = dANIMI32(fxanm[i], fxanm[i].dr, fxanm[i].spd)
			}
		}
	}
	//CLEAR
	for i := 0; i < len(fxanm); i++ {
		if fxanm[i].onoff {
			uFXANIMNUMOFF(i)
			fxanm = remANIM(fxanm, i)
		}
	}
}
func uFXANIMNUMOFF(num int) {
	switch fxanm[num].num {
	case 1: //CHEST OPEN

		runeslev = append(runeslev, runes[RINT(0, 8)])

		w := recResizeWidth(runestiles.r[0].W, runestiles.r[0].H, tileSize/2)
		runeslev[0].r = sdl.Rect{RecCenter(fxanm[num].dr).X, RecCenter(fxanm[num].dr).Y, w, tileSize / 2}

	}
}

// SCROLLS
func dRUNES() {
	for i := range len(runeslev) {
		runeslev[i].ro = dIMSHADOWRO(runeslev[i].im, runeslev[i].r, runeslev[i].ro, 8)
	}
}

// UI
func dUI() {

	x := floor[0].r.X
	x -= tileSize
	y := floor[0].ir.Y + tileSize + tileSize/2
	spc := int32(tileSize / 7)
	inset := int32(tileSize / 3)
	siz := tileSize

	for i := 0; i < pl.slots; i++ {
		r := sdl.Rect{x, y, siz, siz}
		dIMBLUR(etcIM[11], r, 4)

		r2 := r
		r2.X += inset / 2
		r2.Y += inset / 2
		r2.W -= inset
		r2.H -= inset
		if i < len(pl.runes) {
			dIM(pl.runes[i].im, r2)
		}

		r3 := r
		r3.W = r.W / 3
		r3.H = r.H / 3
		r3.X = r.X + r.W - r3.W
		r3.Y = r.Y + r.H - r3.H
		dIM(etcIM[12], r3)
		if i < len(pl.runes) {
			txtXYCOLOR(fmt.Sprint(pl.runes[i].invnum), r3.X+7, r3.Y+1, ORANGEǁ3())

		}
		y += r.H + spc
	}

}

// PLAYER
func dPL() { //MARK: PLAYER
	//WAND
	dWAND()
	//PL SHADOW
	sr := pl.r
	sr.Y += pl.r.H / 2
	sr.W -= pl.r.W / 3
	sr.X += pl.r.W / 6
	sr.H -= pl.r.W / 3
	sr.Y += pl.r.W / 6
	dIMF(etcIM[0], sr)
	//PL IM
	anmspd := 10
	switch pl.state {
	case 1: //MOVE
		anmspd = 3
		if pl.lr {
			plR = dANIMMIRROR(plR, pl.r, anmspd)
		} else {
			plR = dANIM(plR, pl.r, anmspd)
		}
	case 0: //IDLE
		anmspd = 15
		if pl.lr {
			plI = dANIMMIRROR(plI, pl.r, anmspd)
		} else {
			plI = dANIM(plI, pl.r, anmspd)
		}
	}
	//TARGET
	dTARG()

	if pl.orbital > 0 {
		dORBITAL()
	}
}

func dORBITAL() {

	siz := float32(tileSize / 2)

	if pl.orbital >= 1 {
		pl.oangl += 3 + float64(Delta)
		pl.o2angl += 12
		pl.orb1 = rotate(pl.cnt, float64(tileSize*3), pl.oangl)
		r := sdl.FRect{pl.orb1.X, pl.orb1.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 2 {
		pl.oangl2 -= 3 + float64(Delta)
		pl.o2angl2 += 12
		pl.orb2 = rotate(pl.cnt, float64(tileSize*3), pl.oangl2)
		r := sdl.FRect{pl.orb2.X, pl.orb2.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl2, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 3 {
		pl.oangl3 += 3 + float64(Delta)
		pl.o2angl3 += 12
		pl.orb3 = rotate(pl.cnt, float64(tileSize*2), pl.oangl3)
		r := sdl.FRect{pl.orb3.X, pl.orb3.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl3, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 4 {
		pl.oangl4 -= 3 + float64(Delta)
		pl.o2angl4 += 12
		pl.orb4 = rotate(pl.cnt, float64(tileSize*2), pl.oangl4)
		r := sdl.FRect{pl.orb4.X, pl.orb4.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl4, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 5 {
		pl.oangl5 += 3 + float64(Delta)
		pl.o2angl5 += 12
		pl.orb5 = rotate(pl.cnt, float64(tileSize+tileSize/2), pl.oangl5)
		r := sdl.FRect{pl.orb5.X, pl.orb5.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl5, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 6 {
		pl.oangl6 -= 3 + float64(Delta)
		pl.o2angl6 += 12
		pl.orb6 = rotate(pl.cnt, float64(tileSize+tileSize/2), pl.oangl6)
		r := sdl.FRect{pl.orb6.X, pl.orb6.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl6, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 7 {
		pl.oangl7 += 3 + float64(Delta)
		pl.o2angl7 += 12
		pl.orb7 = rotate(pl.cnt, float64(tileSize*4), pl.oangl7)
		r := sdl.FRect{pl.orb7.X, pl.orb7.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl7, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 8 {
		pl.oangl8 -= 3 + float64(Delta)
		pl.o2angl8 += 12
		pl.orb8 = rotate(pl.cnt, float64(tileSize*4), pl.oangl8)
		r := sdl.FRect{pl.orb8.X, pl.orb8.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl8, originF(r), sdl.FLIP_NONE)
	}
	if pl.orbital >= 9 {
		pl.oangl9 += 3 + float64(Delta)
		pl.o2angl9 += 12
		pl.orb9 = rotate(pl.cnt, float64(tileSize*4), pl.oangl9)
		r := sdl.FRect{pl.orb9.X, pl.orb9.Y, siz, siz}
		RNDR.CopyExF(etcIM[3].tex, &etcIM[3].r, &r, pl.o2angl9, originF(r), sdl.FLIP_NONE)
	}
}

func dTARG() {
	siz := float32(tileSize / 2)
	r := sdl.FRect{pl.targ.X - siz/2, pl.targ.Y - siz/2, siz, siz}
	RNDR.CopyExF(etcIM[1].tex, &etcIM[1].r, &r, targangl, originF(r), sdl.FLIP_NONE)
	targangl++

}
func dWAND() {
	pl.wandangl = 60
	if pl.lr {
		pl.wandangl -= 30
	}

	r2 := pl.wr
	if pl.state == 0 {
		if t15 {
			r2.Y -= 2
		}
	}
	RNDR.CopyExF(wandtiles.tex, &pl.wand.ir, &r2, pl.wandangl, originF(pl.wr), sdl.FLIP_NONE)
}

// LEVEL
func dLEV() {
	//FLOOR
	for i := range floor {
		if floor[i].outer {
			dIM(floor[i].im, floor[i].ir)
		} else {
			dIM(floor[i].im, floor[i].r)
		}
	}
	//EXTRAS
	dEXTRAS()
	//DOORS
	dDOORS()
	//INNER WALLS
	for i := range iw {
		dIMRECOUTLINESHADOW(iw[i].im, iw[i].r, 4)
	}
	//PROJ
	dPLPROJ()

}
func dPLPROJ() { //MARK: PL PROJ

	if len(projpl) > 0 {

		for i := range projpl {
			if !projpl[i].onoff {
				projpl[i].anm = dANIM(projpl[i].anm, projpl[i].r, 1)

				//dIMSHADOWONLY(projpl[i].an)

				if DEBUG {
					dRECFLINESIZE(projpl[i].r, 4, MAGENTAǁ3())
					dRECFLINESIZE(projpl[i].cr, 4, ORANGEǁ3())
				}
			}
		}

	}

}
func dDOORS() {
	for i := range doorIM {
		dIM(doorIM[i], doorIM[i].dr)
	}
}
func dEXTRAS() {
	for i := range extras {
		if extras[i].ro != 0 {
			dIMRO(extras[i].im, extras[i].r, extras[i].ro)
		} else {
			//dIMSHADOW(extras[i].im, extras[i].r, 4)
			dIMSHADOW(extras[i].im, extras[i].r, 5, false)
			dIMBLUR(extras[i].im, extras[i].r, 4)
		}
	}
}

// FX
func dIMBLUR(i IM, r sdl.Rect, offset float32) {
	RNDR.Copy(i.tex, &i.r, &r)
	r2 := rec2frec(r)
	r2.X += RF32(-offset, offset)
	r2.Y += RF32(-offset, offset)
	if frames%3 == 0 {
		i.tex.SetAlphaMod(50)
		RNDR.CopyF(i.tex, &i.r, &r2)
		i.tex.SetAlphaMod(255)
	}
}
func dIMROBLUR(i IM, r sdl.Rect, angl float64, offset float32) {
	RNDR.CopyEx(i.tex, &i.r, &r, angl, origin2(r), sdl.FLIP_NONE)

	r2 := rec2frec(r)
	r2.X += RF32(-offset, offset)
	r2.Y += RF32(-offset, offset)
	if frames%3 == 0 {
		i.tex.SetAlphaMod(50)
		RNDR.CopyExF(i.tex, &i.r, &r2, angl, origin2int32toF(r), sdl.FLIP_NONE)
		i.tex.SetAlphaMod(255)
	}
}

// ANIM
func dANIMI32(a ANIM, r sdl.Rect, spd int) ANIM {
	a = pANIM(a, spd)
	RNDR.Copy(a.tex, &a.r, &r)
	return a
}
func dANIM(a ANIM, r sdl.FRect, spd int) ANIM {
	a = pANIM(a, spd)
	RNDR.CopyF(a.tex, &a.r, &r)
	return a
}
func dANIMONCEF(a ANIM, r sdl.FRect, spd int, complete bool) (ANIM, bool) {
	a, complete = pANIMONCE(a, spd, complete)
	RNDR.CopyF(a.tex, &a.r, &r)
	return a, complete
}
func dANIMONCE(a ANIM, r sdl.Rect, spd int, complete bool) (ANIM, bool) {
	a, complete = pANIMONCE(a, spd, complete)
	RNDR.Copy(a.tex, &a.r, &r)
	return a, complete
}
func dANIMMIRROR(a ANIM, r sdl.FRect, spd int) ANIM {
	a = pANIM(a, spd)
	RNDR.CopyExF(a.tex, &a.r, &r, 0, &pZEROF32, sdl.FLIP_HORIZONTAL)
	return a
}
func pANIM(a ANIM, spd int) ANIM {
	if frames%spd == 0 {
		a.r.X += a.r.W
		a.framecount++
		if a.framecount == a.frames {
			a.r.X = a.xl
			a.framecount = 0
		}
	}
	return a
}
func pANIMONCE(a ANIM, spd int, complete bool) (ANIM, bool) {
	if !complete {
		if frames%spd == 0 {
			a.r.X += a.r.W
			a.framecount++
			if a.framecount == a.frames {
				a.r.X = a.xl
				a.framecount = 0
				complete = true
			}
		}
	}
	return a, complete
}

// IMAGES
func dIMSHADOWONLY(i IM, r sdl.Rect, offset int32, lr bool) {
	i.tex.SetAlphaMod(100)
	i.tex.SetColorMod(0, 0, 0)
	r2 := r
	if lr {
		r2.X += offset
	} else {
		r2.X -= offset
	}
	r2.Y += offset
	RNDR.Copy(i.tex, &i.r, &r2)
}
func dIMSHADOWRO(i IM, r sdl.Rect, ro float64, offset int32) float64 {
	ro++
	i.tex.SetAlphaMod(100)
	i.tex.SetColorMod(0, 0, 0)
	r2 := r
	r2.X -= offset
	r2.Y += offset
	RNDR.CopyEx(i.tex, &i.r, &r2, ro, origin2(r), sdl.FLIP_NONE)
	i.tex.SetAlphaMod(255)
	i.tex.SetColorMod(255, 255, 255)
	RNDR.CopyEx(i.tex, &i.r, &r, ro, origin2(r), sdl.FLIP_NONE)
	return ro
}
func dIM(i IM, r sdl.Rect) {
	RNDR.Copy(i.tex, &i.r, &r)
}
func dIMRO(i IM, r sdl.Rect, angl float64) {
	RNDR.CopyEx(i.tex, &i.r, &r, angl, origin2(r), sdl.FLIP_NONE)
}
func dIMF(i IM, r sdl.FRect) {
	RNDR.CopyF(i.tex, &i.r, &r)
}
func dIMSHADOW(i IM, r sdl.Rect, offset int32, lr bool) {
	i.tex.SetAlphaMod(100)
	i.tex.SetColorMod(0, 0, 0)
	r2 := r
	if lr {
		r2.X += offset
	} else {
		r2.X -= offset
	}
	r2.Y += offset
	RNDR.Copy(i.tex, &i.r, &r2)
	i.tex.SetAlphaMod(255)
	i.tex.SetColorMod(255, 255, 255)
	RNDR.Copy(i.tex, &i.r, &r)
}
func dIMRECSHADOW(i IM, r sdl.Rect, offset int32) {
	RNDR.SetDrawColor(0, 0, 0, 170)
	r2 := r
	r2.X -= offset
	r2.Y += offset
	RNDR.FillRect(&r2)
	RNDR.Copy(i.tex, &i.r, &r)
}
func dIMRECOUTLINESHADOW(i IM, r sdl.Rect, offset int32) {
	RNDR.SetDrawColor(0, 0, 0, 155)
	r2 := r
	r2.X -= offset
	r2.Y += offset
	RNDR.FillRect(&r2)
	RNDR.Copy(i.tex, &i.r, &r)
	RNDR.SetDrawColor(0, 0, 0, 255)
	RNDR.DrawRect(&r)
}

// ISO RECS
func dISOREC(p []sdl.Point, c []uint8) {
	RNDR.SetDrawColor(c[0], c[1], c[2], c[3])
	dLINE(p[0], p[1])
	dLINE(p[1], p[2])
	dLINE(p[2], p[3])
	dLINE(p[0], p[3])
}

// TILESHEET
func dTILESHEET(t TILESHEET, x, y, space, zoom int32) {
	r := sdl.Rect{x, y, t.r[0].W * zoom, t.r[0].H * zoom}
	for i := range t.r {
		RNDR.Copy(t.tex, &t.r[i], &r)
		txtXY(fmt.Sprint(i), r.X+2, r.Y+2)
		r.X += r.W + space
	}
}

// LINES
func dLINE(p1, p2 sdl.Point) {
	RNDR.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
}

// RECS
func dRECFILLALPHA(r sdl.Rect, c sdl.Color, a uint8) {
	c.A = a
	RNDR.SetDrawColor(c.R, c.G, c.B, c.A)
	RNDR.FillRect(&r)

}
func dRECLINESIZE(r sdl.Rect, lineW int32) {
	RNDR.SetDrawColor(MAGENTAǁ2())
	RNDR.DrawRect(&r)
	for lineW > 0 {
		r.X--
		r.Y--
		r.W += 2
		r.H += 2
		RNDR.DrawRect(&r)
		lineW--
	}
}
func dRECLINESIZECOLOR(r sdl.Rect, lineW int32, c sdl.Color) {
	RNDR.SetDrawColor(c.R, c.G, c.B, c.A)
	RNDR.DrawRect(&r)
	for lineW > 0 {
		r.X--
		r.Y--
		r.W += 2
		r.H += 2
		RNDR.DrawRect(&r)
		lineW--
	}
}
func dRECFLINESIZE(r sdl.FRect, lineW float32, c sdl.Color) {
	RNDR.SetDrawColor(c.R, c.G, c.B, c.A)
	RNDR.DrawRectF(&r)
	for lineW > 0 {
		r.X--
		r.Y--
		r.W += 2
		r.H += 2
		RNDR.DrawRectF(&r)
		lineW--
	}
}
func dRECLINE(r sdl.Rect, c sdl.Color) {
	RNDR.SetDrawColor(c.R, c.G, c.B, c.A)
	RNDR.DrawRect(&r)
}
func dSQCNTRF(cnt sdl.FPoint, w float32) {
	RNDR.SetDrawColor(MAGENTAǁ2())
	r := sdl.FRect{cnt.X - w/2, cnt.Y - w/2, w, w}
	RNDR.FillRectF(&r)
}
func dSQCNTR(cnt sdl.Point, w int32) {
	RNDR.SetDrawColor(MAGENTAǁ2())
	r := sdl.Rect{cnt.X - w/2, cnt.Y - w/2, w, w}
	RNDR.FillRect(&r)
}
func dFRECFILL(r sdl.FRect, c sdl.Color) {
	RNDR.SetDrawColor(c.R, c.G, c.B, c.A)
	RNDR.FillRectF(&r)
}
