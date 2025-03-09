package main

import "github.com/veandco/go-sdl2/sdl"

var (
	pl                   PLAYER
	projpl               []PROJ
	enmlist, enm, enmall []ENEMY

	enmT, enmnum int

	enmspawnpoints []sdl.FPoint
)

type ENEMY struct {
	r       sdl.FRect
	cnt     sdl.FPoint
	spawned bool
}

type PROJ struct {
	r, cr               sdl.FRect
	spd, velx, vely, oW float32
	cnt                 sdl.FPoint
	anm                 ANIM
	onoff               bool

	bounce, split, zigzag, grow, splinter int
}

type PLAYER struct {
	r, collisr, wr sdl.FRect
	cnt            Vector2
	lr             bool
	wand           WAND
	radtarg        float32
	runes          []RUNE

	targ, orb1, orb2, orb3, orb4, orb5, orb6, orb7, orb8, orb9 sdl.FPoint

	oangl, oangl2, oangl3, oangl4, oangl5, oangl6, oangl7, oangl8, oangl9 float64

	o2angl, o2angl2, o2angl3, o2angl4, o2angl5, o2angl6, o2angl7, o2angl8, o2angl9 float64

	state, projT, projP, slots     int
	spd, velx, vely, acc, wandangl float64

	bounce, split, rear, zigzag, faster, grow, orbital, splinter int
}

func mENEMIES() {

	p1 := sdl.FPoint{float32(floor[25].r.X + floor[25].r.W), float32(floor[25].r.Y + floor[25].r.H)}
	p2 := sdl.FPoint{float32(floor[46].r.X), float32(floor[46].r.Y + floor[46].r.H)}
	p3 := sdl.FPoint{float32(floor[289].r.X + floor[289].r.W), float32(floor[289].r.Y)}
	p4 := sdl.FPoint{float32(floor[310].r.X), float32(floor[310].r.Y)}

	enmspawnpoints = append(enmspawnpoints, p1, p2, p3, p4)

	enmT = int(FPS)
	enmnum = 10

	for i := 0; i < enmnum; i++ {

		siz := float32(tileSize)
		e := ENEMY{}
		choose := RINT(0, 5)
		switch choose {
		case 0:
			e.cnt = enmspawnpoints[0]
		case 1:
			e.cnt = enmspawnpoints[1]
		case 2:
			e.cnt = enmspawnpoints[2]
		case 3:
			e.cnt = enmspawnpoints[3]
		}

		e.r = sqcntrF(e.cnt, siz)

		enmlist = append(enmlist, e)
	}

}

func nxENEMIES() {

	if enmnum > 0 {
		for i := range enmlist {
			if !enmlist[i].spawned {
				enmlist[i].spawned = true
				enm = append(enm, enmlist[i])
				enmT = int(FPS) * RINT(1, 4)
				break
			}
		}
	}

}
func mPLPROJ() {

	siz := float32(tileSize)
	p := PROJ{}
	p.spd = float32(tileSize) / 6
	p.cnt = sdl.FPoint{float32(pl.cnt.X), float32(pl.cnt.Y)}
	p.r = sdl.FRect{float32(pl.cnt.X) - siz/2, float32(pl.cnt.Y) - siz/2, siz, siz}
	p.oW = p.r.W
	p.cr = p.r
	p.cr.W -= p.r.W / 2
	p.cr.H -= p.r.H / 2
	p.cr.X += p.r.W / 4
	p.cr.Y += p.r.H / 4

	p.velx, p.vely = vel2pointsVECFPOINT(pl.cnt, pl.targ, p.spd)

	//FASTER
	if pl.faster > 0 {
		p.velx *= (float32(pl.faster) / 10) + 1.5
		p.vely *= (float32(pl.faster) / 10) + 1.5
	}

	p.anm = fireball

	p.splinter = pl.splinter
	p.bounce = pl.bounce
	p.split = pl.split
	p.zigzag = pl.zigzag
	p.grow = pl.grow

	projpl = append(projpl, p)

	//REAR
	if pl.rear > 0 {
		p2 := p
		p2.velx *= -1
		p2.vely *= -1
		projpl = append(projpl, p2)
		if pl.rear > 1 {
			p2.cnt.X += p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
		if pl.rear > 2 {
			p2.cnt.X += p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
		if pl.rear > 3 {
			p2.cnt.X += p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
		if pl.rear > 4 {
			p2.cnt.X += p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
		if pl.rear > 5 {
			p2 = p
			p2.cnt.X -= p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
		if pl.rear > 6 {
			p2.cnt.X -= p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
		if pl.rear > 7 {
			p2.cnt.X -= p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
		if pl.rear > 8 {
			p2.cnt.X -= p2.r.W / 3
			p2.cnt.Y += p2.r.W / 3
			projpl = append(projpl, p2)
		}
	}

}

func mPL() {

	siz := float32(tileSize)

	pl.cnt = Vector2{float64(CNTR.X), float64(CNTR.Y)}
	pl.r = sdl.FRect{float32(pl.cnt.X) - siz/2, float32(pl.cnt.Y) - siz/2, siz, siz}
	pl.collisr.W = float32(iw[0].r.W) - 12
	pl.collisr.H = float32(iw[0].r.W) - 12
	pl.collisr = sdl.FRect{float32(pl.cnt.X) - pl.collisr.W/2, float32(pl.cnt.Y) - pl.collisr.H/2, pl.collisr.W, pl.collisr.H}
	pl.collisr.Y = pl.r.Y + pl.r.H - pl.collisr.H
	pl.wr = pl.r
	pl.wr.Y -= pl.r.W / 5
	if pl.lr {
		pl.wr.X -= pl.r.W / 4
	} else {
		pl.wr.X += pl.r.W / 4
	}
	pl.spd = float64(tileSize) / 9
	pl.acc = 1
	pl.wand = wands[RINT(0, len(wands))]

	pl.radtarg = float32(tileSize * 4)
	pl.targ = sdl.FPoint{float32(pl.cnt.X) + pl.radtarg, float32(pl.cnt.Y)}

	pl.projP = int(FPS) * 2

	pl.wandangl = 60
	pl.slots = 10

	pl.o2angl2 = 40
	pl.o2angl3 = 80
	pl.o2angl4 = 120
	pl.o2angl5 = 160
	pl.o2angl6 = 200
	pl.o2angl7 = 240
	pl.o2angl8 = 280
	pl.o2angl9 = 320

}
