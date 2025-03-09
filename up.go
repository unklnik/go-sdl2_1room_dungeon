package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	//MOUSE
	mouseL, mouseR, mouseM bool
	mouse                  sdl.Point
)

func UPDATE() {
	MOUSE()
	uPL()
	if len(projpl) > 0 {
		clear := upPROJPL()
		if clear {
			for i := 0; i < len(projpl); i++ {
				if projpl[i].onoff {
					projpl = remPROJ(projpl, i)
				}
			}
		}
	}

	if len(runeslev) > 0 {
		uRUNES()
	}

	//TIMERS
	if enmT > 0 {
		enmT--
		if enmT == 0 {
			nxENEMIES()
		}
	}
	if frames%15 == 0 {
		t15 = !t15
	}
	if clickpauseL > 0 {
		clickpauseL--
	}
	if clickpauseR > 0 {
		clickpauseR--
	}
	if pl.projT > 0 {
		pl.projT--
	}
}

func collectrune(num int) { //MARK: COLLECT RUNE

	found := false
	if len(pl.runes) > 0 {
		for i := range pl.runes {
			if runeslev[num].nm == pl.runes[i].nm {
				switch runeslev[num].nm {

				/*

					ADD BELOW

				*/
				case "splinter":
					if pl.splinter < 9 {
						pl.splinter++
						pl.runes[i].invnum++
					} else {
						addmsg("splinter rune max")
					}
					found = true
				case "orbital":
					if pl.orbital < 9 {
						pl.orbital++
						pl.runes[i].invnum++
					} else {
						addmsg("orbital rune max")
					}
					found = true
				case "zigzag":
					if pl.zigzag < 9 {
						pl.zigzag++
						pl.runes[i].invnum++
					} else {
						addmsg("zigzag rune max")
					}
					found = true
				case "bounce":
					if pl.bounce < 9 {
						pl.bounce++
						pl.runes[i].invnum++
					} else {
						addmsg("bounce rune max")
					}
					found = true
				case "split":
					if pl.split < 9 {
						pl.split++
						pl.runes[i].invnum++
					} else {
						addmsg("split rune max")
					}
					found = true
				case "rear":
					if pl.rear < 9 {
						pl.rear++
						pl.runes[i].invnum++
					} else {
						addmsg("rear rune max")
					}
					found = true
				case "faster":
					if pl.faster < 9 {
						pl.faster++
						pl.runes[i].invnum++
					} else {
						addmsg("faster rune max")
					}
					found = true
				case "grow":
					if pl.grow < 9 {
						pl.grow++
						pl.runes[i].invnum++
					} else {
						addmsg("grow rune max")
					}
					found = true

					/*

						ADD BELOW

					*/
				}
			}
		}
	}
	if !found { //ADD HERE
		pl.runes = append(pl.runes, runeslev[num])

		switch runeslev[num].nm {
		case "splinter":
			pl.splinter++
		case "orbital":
			pl.orbital++
		case "zigzag":
			pl.zigzag++
		case "bounce":
			pl.bounce++
		case "split":
			pl.split++
		case "rear":
			pl.rear++
		case "faster":
			pl.faster++
		case "grow":
			pl.grow++
		}
	}
}
func upPROJPL() bool { //MARK: PL PROJ
	clear := false
	for i := range projpl {
		if !projpl[i].onoff {

			//GROW
			if projpl[i].grow > 0 {
				multi := float32(projpl[i].grow) * 7
				w := multi * projpl[i].oW
				if projpl[i].r.W < w {
					projpl[i].r.W += 4
				}

			}

			//ZIGZAG
			if projpl[i].zigzag > 0 {
				if Roll12() == 12 {
					var angl float64
					for {
						angl = RF64(-60, 61)
						if Abs(float32(angl)) > 10 {
							break
						}
					}
					projpl[i].velx, projpl[i].vely = velXYnewpoint(projpl[i].cnt, float64(projpl[i].velx), float64(projpl[i].vely), float64(projpl[i].spd), angl)
					projpl[i].zigzag--
				}

			}

			projpl[i].cnt.X += projpl[i].velx + Delta
			projpl[i].cnt.Y += projpl[i].vely + Delta
			projpl[i].r = sqcntrF(projpl[i].cnt, projpl[i].r.W)
			projpl[i].cr = sqcntrF(projpl[i].cnt, projpl[i].cr.W)

			//SPLIT
			if projpl[i].split > 0 {
				if Roll18()+projpl[i].split > 5 {
					projpl[i].split = 0
					p2 := projpl[i]
					for {
						p2.velx += RF32(-float32(tileSize/5), float32(tileSize/5))
						if Abs(p2.velx) > p2.spd/2 {
							break
						}
					}
					for {
						p2.vely += RF32(-float32(tileSize/5), float32(tileSize/5))
						projpl = append(projpl, p2)
						if Abs(p2.vely) > p2.spd/2 {
							break
						}
					}
				}
			}

			if !checkinnwallFrec(projpl[i].cr) {
				if projpl[i].bounce > 0 { //BOUNCE
					projpl[i].velx *= -1
					projpl[i].vely *= -1
					projpl[i].bounce--
				} else {
					if projpl[i].splinter > 0 {
						splinterproj(i)
					}
					projpl[i].onoff = true
					anm := endprojanm
					anm.once = true
					anm.spd = 2
					anm.dr = resizerec(frec2rec(projpl[i].r), 1)
					fxanm = append(fxanm, anm)
				}
			} else if !PointFInRec(projpl[i].cnt, levrecfloor) {
				if projpl[i].bounce > 0 { //BOUNCE
					projpl[i].velx *= -1
					projpl[i].vely *= -1
					projpl[i].bounce--
				} else {
					if projpl[i].splinter > 0 {
						splinterproj(i)
					}
					projpl[i].onoff = true
					anm := endprojanm
					anm.once = true
					anm.spd = 2
					anm.dr = resizerec(frec2rec(projpl[i].r), 1)
					fxanm = append(fxanm, anm)
				}
			}
		} else {
			clear = true
		}
	}

	return clear
}

func splinterproj(num int) {

	if projpl[num].splinter >= 1 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = 0
		p2.vely = p2.spd
		projpl = append(projpl, p2)
		p2.vely = -p2.spd
		projpl = append(projpl, p2)
		p2.velx = p2.spd
		p2.vely = 0
		projpl = append(projpl, p2)
		p2.velx = -p2.spd
		p2.vely = 0
		projpl = append(projpl, p2)
	}
	if projpl[num].splinter >= 2 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = p2.spd / 2
		p2.vely = p2.spd / 2
		projpl = append(projpl, p2)
		p2.velx = -p2.spd / 2
		p2.vely = -p2.spd / 2
		projpl = append(projpl, p2)
	}
	if projpl[num].splinter >= 3 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = -p2.spd / 2
		p2.vely = p2.spd / 2
		projpl = append(projpl, p2)
		p2.velx = p2.spd / 2
		p2.vely = -p2.spd / 2
		projpl = append(projpl, p2)
	}
	if projpl[num].splinter >= 4 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = -((p2.spd / 4) * 3)
		p2.vely = p2.spd / 4
		projpl = append(projpl, p2)
		p2.velx = p2.spd / 4
		p2.vely = -p2.spd / 4
		projpl = append(projpl, p2)
	}
	if projpl[num].splinter >= 5 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = ((p2.spd / 4) * 3)
		p2.vely = p2.spd / 4
		projpl = append(projpl, p2)
		p2.velx = p2.spd / 4
		p2.vely = ((p2.spd / 4) * 3)
		projpl = append(projpl, p2)
	}
	if projpl[num].splinter >= 6 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = -((p2.spd / 8) * 5)
		p2.vely = ((p2.spd / 8) * 3)
		projpl = append(projpl, p2)
		p2.velx = ((p2.spd / 8) * 5)
		p2.vely = -((p2.spd / 8) * 5)
		projpl = append(projpl, p2)
	}
	if projpl[num].splinter >= 7 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = -((p2.spd / 12) * 5)
		p2.vely = ((p2.spd / 12) * 7)
		projpl = append(projpl, p2)
		p2.velx = ((p2.spd / 12) * 5)
		p2.vely = -((p2.spd / 12) * 7)
		projpl = append(projpl, p2)
	}
	if projpl[num].splinter >= 8 {
		p2 := projpl[num]
		p2.splinter = 0
		p2.velx = -((p2.spd / 12) * 5)
		p2.vely = -((p2.spd / 12) * 7)
		projpl = append(projpl, p2)
		p2.velx = ((p2.spd / 12) * 5)
		p2.vely = ((p2.spd / 12) * 7)
		projpl = append(projpl, p2)
	}

	if projpl[num].splinter > 0 {
		projpl[num].splinter--
	}
}

func uRUNES() {
	clear := false
	for i := range runeslev {
		if !runeslev[i].collected {
			if RecsIntersect(runeslev[i].r, frec2rec(pl.r)) {
				runeslev[i].collected = true
				clear = true
				collectrune(i)
			}
		}
	}
	if clear {
		for i := 0; i < len(runeslev); i++ {
			if runeslev[i].collected {
				runeslev = remRUNE(runeslev, i)
			}
		}
	}
}

func checkinnwallFrec(r sdl.FRect) bool {
	canadd := true
	if len(iw) > 0 {
		for i := range iw {
			if RecsIntersect(iw[i].r, frec2rec(r)) {
				canadd = false
				break
			}
		}
	}
	return canadd
}

func uPL() { //MARK: PLAYER

	if kR {
		if pl.velx < 0 {
			pl.velx = 0
		}
		if pl.velx < pl.spd {
			pl.velx += pl.acc
		} else if pl.velx > pl.spd {
			pl.velx = pl.spd
		}
	}
	if kL {
		if pl.velx > 0 {
			pl.velx = 0
		}
		if pl.velx > -pl.spd {
			pl.velx -= pl.acc
		} else if pl.velx < -pl.spd {
			pl.velx = -pl.spd
		}
	}
	if kU {
		if pl.vely > 0 {
			pl.vely = 0
		}
		if pl.vely > -pl.spd {
			pl.vely -= pl.acc
		} else if pl.vely < -pl.spd {
			pl.vely = -pl.spd
		}
	}
	if kD {
		if pl.vely < 0 {
			pl.vely = 0
		}
		if pl.vely < pl.spd {
			pl.vely += pl.acc
		} else if pl.vely > pl.spd {
			pl.vely = pl.spd
		}
	}

	//IDLE SLOWDOWN
	if !kR && !kL && !kU && !kD {

		pl.state = 0
		pl.velx = 0
		pl.vely = 0
	}

	//MOVE
	if pl.velx != 0 && pl.vely != 0 {
		pl.velx, pl.vely = normalizeSpeed(pl.velx, pl.vely, pl.spd)
	}
	if pl.velx != 0 || pl.vely != 0 {
		if checkPLmoveX() {
			pl.state = 1
			pl.cnt.X += pl.velx + float64(Delta)
		} else {
			pl.velx = 0
		}
		if checkPLmoveY() {
			pl.state = 1
			pl.cnt.Y += pl.vely + float64(Delta)
		} else {
			pl.vely = 0
		}
		if pl.velx != 0 || pl.vely != 0 {
			pl.r = sdl.FRect{float32(pl.cnt.X) - pl.r.W/2, float32(pl.cnt.Y) - pl.r.H/2, pl.r.W, pl.r.H}
			pl.collisr = sdl.FRect{float32(pl.cnt.X) - pl.collisr.W/2, float32(pl.cnt.Y) - pl.collisr.H/2, pl.collisr.W, pl.collisr.H}
			pl.collisr.Y = pl.r.Y + pl.r.H - pl.collisr.H
			pl.wr = pl.r
			pl.wr.Y -= pl.r.W / 5
			if pl.lr {
				pl.wr.X -= pl.r.W / 5
			} else {
				pl.wr.X += pl.r.W / 5
			}

		}
	}

	pl.targ = v2tofpoint2(pointOnCirc(pl.radtarg, float32(angl2points(pl.cnt, point2v2(mouse))), pl.cnt))

}
func checkPLmoveX() bool {
	canmove := true

	r := pl.collisr
	r.X += float32(pl.velx)

	for i := range len(iw) {
		if RecsFIntersect(r, rec2frec(iw[i].r)) {
			canmove = false
			break
		}
	}
	if canmove {
		for i := range len(floor) {
			if floor[i].solid {
				if RecsFIntersect(r, rec2frec(floor[i].r)) {
					canmove = false
					break
				}
			}
		}
	}
	return canmove
}
func checkPLmoveY() bool {
	canmove := true

	r := pl.collisr
	r.Y += float32(pl.vely)

	for i := range len(iw) {
		if RecsFIntersect(r, rec2frec(iw[i].r)) {
			canmove = false
			break
		}
	}
	if canmove {
		for i := range len(floor) {
			if floor[i].solid {
				if RecsFIntersect(r, rec2frec(floor[i].r)) {
					canmove = false
					break
				}
			}
		}
	}
	return canmove
}
