package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	msgs   []MSG
	runes  []RUNE
	wands  []WAND
	fxanm  []ANIM
	chests []CHEST

	wandnm1 = []string{"baton", "staff", "pole", "rod", "sceptre", "twig", "sprig", "wand", "mace"}
)

type WAND struct {
	nm        string
	atk, gems int
	ir        sdl.Rect
}
type RUNE struct {
	nm        string
	im        IM
	r         sdl.Rect
	collected bool
	ro        float64

	invnum int
}
type CHEST struct {
	im   IM
	r    sdl.Rect
	num  int
	open bool
}

type MSG struct {
	t     string
	timer int
	onoff bool
}

func mRUNES() {

	for i := range runestiles.r {
		s := RUNE{}
		s.im = imfromtile(runestiles, runestiles.r[i])
		s.invnum = 1
		switch i {
		case 0: // BOUNCE
			s.nm = "bounce"
		case 1: // SPLIT
			s.nm = "split"
		case 2: // REAR
			s.nm = "rear"
		case 3: // ZIGZAG
			s.nm = "zigzag"
		case 4: // FASTER
			s.nm = "faster"
		case 5: // GROW
			s.nm = "grow"
		case 6: // ORBITAL
			s.nm = "orbital"
		case 7: // SPLINTER
			s.nm = "splinter"
		default:
			s.nm = fmt.Sprint(i)
		}

		runes = append(runes, s)
	}

}
func addmsg(t string) {
	m := MSG{}
	m.t = t
	m.timer = int(FPS) * 3
	msgs = append(msgs, m)
}

func mCHEST() {
	siz := tileSize
	c := CHEST{}
	c.im = extrasIM[RI32(28, 31)]
	c.r = sdl.Rect{CNTR.X - siz/2, CNTR.Y - siz/2, siz, siz}
	c.r.X += tileSize * 2
	chests = append(chests, c)
}

func mWANDS() {

	for i := range wandtiles.r {
		w := WAND{}
		w.ir = wandtiles.r[i]
		w.nm = mWANDNAME()
		w.atk = RINT(1, 3)
		w.gems = RINT(0, 3)
		wands = append(wands, w)
	}

}
func mWANDNAME() string {
	txt := " of "
	txt = wandnm1[RINT(0, len(wandnm1))] + txt
	return txt
}

func remRUNE(slice []RUNE, s int) []RUNE {
	return append(slice[:s], slice[s+1:]...)
}
