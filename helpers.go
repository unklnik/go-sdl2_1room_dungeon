package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"

	"math/rand/v2"
)

func orbit(cx, cy, radius, angle float64) (float64, float64) {
	x := cx + radius*math.Cos(angle)
	y := cy + radius*math.Sin(angle)
	return x, y
}
func rotate(cnt Vector2, rad, angl float64) sdl.FPoint { //BETTER VERSION
	radians := angl * math.Pi / 180.0
	newX := cnt.X + rad*math.Cos(float64(radians))
	newY := cnt.Y + rad*math.Sin(float64(radians))
	return sdl.FPoint{float32(newX), float32(newY)}
}
func velXYnewpoint(originalPoint sdl.FPoint, xVelocity, yVelocity, maxSpeed, angle float64) (float32, float32) {

	angleRad := angle * math.Pi / 180.0

	newXVelocity := xVelocity*math.Cos(angleRad) - yVelocity*math.Sin(angleRad)
	newYVelocity := xVelocity*math.Sin(angleRad) + yVelocity*math.Cos(angleRad)

	newSpeed := math.Sqrt(newXVelocity*newXVelocity + newYVelocity*newYVelocity)
	if newSpeed > maxSpeed {
		scaleFactor := maxSpeed / newSpeed
		newXVelocity *= scaleFactor
		newYVelocity *= scaleFactor
	}

	return float32(newXVelocity), float32(newYVelocity)
}

func resizerec(r sdl.Rect, scale int32) sdl.Rect {
	return sdl.Rect{RecCenter(r).X - ((r.W * scale) / 2), RecCenter(r).Y - ((r.H * scale) / 2), r.W * scale, r.H * scale}
}

func ranrec(siz int32) sdl.Rect {

	x := levrecin.X
	y := levrecin.Y
	r := sdl.Rect{}

	canadd := true
	breakcount := 1000

	for {
		x += RI32(0, int(levrecin.W-siz))
		y += RI32(0, int(levrecin.H-siz))
		r = sdl.Rect{x, y, siz, siz}
		canadd = checkaddextra(r)
		breakcount--
		if canadd || breakcount == 0 {
			break
		}
	}

	return r
}

func imfromtile(t TILESHEET, r sdl.Rect) IM {
	i := IM{}
	i.tex = t.tex
	i.r = r
	return i
}
func remMSG(slice []MSG, s int) []MSG {
	return append(slice[:s], slice[s+1:]...)
}
func remPROJ(slice []PROJ, s int) []PROJ {
	return append(slice[:s], slice[s+1:]...)
}
func remANIM(slice []ANIM, s int) []ANIM {
	return append(slice[:s], slice[s+1:]...)
}

type Vector2 struct {
	X, Y float64
}

func sqcntrF(cnt sdl.FPoint, w float32) sdl.FRect {
	return sdl.FRect{cnt.X - w/2, cnt.Y - w/2, w, w}
}

func vel2pointsVECFPOINT(p1 Vector2, p2 sdl.FPoint, maxspeed float32) (xvel float32, yvel float32) {
	xvel = p2.X - float32(p1.X)
	yvel = p2.Y - float32(p1.Y)

	distance := float32(math.Sqrt(float64(xvel*xvel + yvel*yvel)))

	if distance > maxspeed {
		scale := maxspeed / distance
		xvel *= scale
		yvel *= scale
	}

	return xvel, yvel
}

func v2tofpoint(v Vector2) *sdl.FPoint {
	return &sdl.FPoint{float32(v.X), float32(v.Y)}
}
func v2tofpoint2(v Vector2) sdl.FPoint {
	return sdl.FPoint{float32(v.X), float32(v.Y)}
}
func fpoint2v2(p sdl.FPoint) Vector2 {
	return Vector2{float64(p.X), float64(p.Y)}
}
func point2v2(p sdl.Point) Vector2 {
	return Vector2{float64(p.X), float64(p.Y)}
}
func rec2frec(r sdl.Rect) sdl.FRect {
	return sdl.FRect{float32(r.X), float32(r.Y), float32(r.W), float32(r.H)}
}
func frec2rec(r sdl.FRect) sdl.Rect {
	return sdl.Rect{int32(r.X), int32(r.Y), int32(r.W), int32(r.H)}
}
func normalizeSpeed(x, y, maxspd float64) (float64, float64) {
	v := Vector2{x, y}
	v2 := Normalize(v)
	newx := v2.X * maxspd
	newy := v2.Y * maxspd
	return newx, newy
}

func Normalize(v Vector2) Vector2 {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if magnitude == 0 {
		return v // Avoid division by zero
	}
	return Vector2{v.X / magnitude, v.Y / magnitude}
}

func recResizeHeight(w, h, newW int32) int32 {
	return (newW * h) / w
}
func recResizeWidth(w, h, newH int32) int32 {
	return (newH * w) / h
}
func origin(r sdl.Rect) sdl.Point {
	return sdl.Point{r.W / 2, r.H / 2}
}
func origin2(r sdl.Rect) *sdl.Point {
	return &sdl.Point{r.W / 2, r.H / 2}
}
func originF(r sdl.FRect) *sdl.FPoint {
	return &sdl.FPoint{r.W / 2, r.H / 2}
}
func origin2int32toF(r sdl.Rect) *sdl.FPoint {
	return &sdl.FPoint{float32(r.W / 2), float32(r.H / 2)}
}
func checkpointtri(p sdl.Point, tri []sdl.Point) bool {
	as_x := p.X - tri[0].X
	as_y := p.Y - tri[0].Y
	s_ab := (tri[1].X-tri[0].X)*as_y-(tri[1].Y-tri[0].Y)*as_x > 0
	if (tri[2].X-tri[0].X)*as_y-(tri[2].Y-tri[0].Y)*as_x > 0 == s_ab {
		return false
	}
	if (tri[2].X-tri[1].X)*(p.Y-tri[1].Y)-(tri[2].Y-tri[1].Y)*(p.X-tri[1].X) > 0 != s_ab {
		return false
	}
	return true
}

func objSORT(o []OBJ) []OBJ {
	//sort.Slice(o, func(i, j int) bool { return o[i].zi > o[j].zi })
	return o
}
func ChangeAlpha(col []uint8, alpha uint8) []uint8 {
	col[3] = alpha
	return col
}
func RecCenter(r sdl.Rect) sdl.Point {
	return sdl.Point{r.X + r.W/2, r.Y + r.H/2}
}
func RecsFIntersect(r1, r2 sdl.FRect) bool {
	return r1.X < r2.X+r2.W &&
		r1.X+r1.W > r2.X &&
		r1.Y < r2.Y+r2.H &&
		r1.Y+r1.H > r2.Y
}
func RecsIntersect(r1, r2 sdl.Rect) bool {
	return r1.X < r2.X+r2.W &&
		r1.X+r1.W > r2.X &&
		r1.Y < r2.Y+r2.H &&
		r1.Y+r1.H > r2.Y
}
func PointInRecXY(p sdl.Point, x, y, width, height int32) bool {
	return p.X >= x && p.X <= x+width && p.Y >= y && p.Y <= y+height
}
func PointInRec(p sdl.Point, r sdl.Rect) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}
func PointFInRec(p sdl.FPoint, r sdl.Rect) bool {
	return p.X >= float32(r.X) && p.X <= float32(r.X+r.W) && p.Y >= float32(r.Y) && p.Y <= float32(r.Y+r.H)
}
func Point2FPoint(p sdl.Point) sdl.FPoint {
	return sdl.FPoint{float32(p.X), float32(p.Y)}
}
func ColorSliceToSDLColor(color []uint8) sdl.Color {
	return sdl.Color{color[0], color[1], color[2], color[3]}
}
func ColorSDLtoSlice(c sdl.Color) []uint8 {
	return []uint8{c.R, c.G, c.B, c.A}
}
func Abs(n float32) float32 {
	return float32(math.Abs(float64(n)))
}

func pointOnCirc(radius, angl float32, cntr Vector2) Vector2 {
	angleInRadians := float64(angl * math.Pi / 180.0)
	x := float64(radius)*math.Cos(angleInRadians) + float64(cntr.X)
	y := float64(radius)*math.Sin(angleInRadians) + float64(cntr.Y)
	return Vector2{x, y}
}
func angl2points(v1, v2 Vector2) float64 {
	deltaX := v2.X - v1.X
	deltaY := v2.Y - v1.Y
	radians := math.Atan2(deltaY, deltaX)
	degrees := (radians * 180) / math.Pi
	for degrees >= 360 {
		degrees -= 360
	}
	for degrees < 0 {
		degrees += 360
	}
	return degrees
}
func AbsDiff(n, n2 float32) float32 {
	return float32(math.Abs(float64(n) - float64(n2)))
}

func FlipCoin() bool {
	onoff := false
	choose := RINT(0, 100001)
	if choose > 50000 {
		onoff = true
	}
	return onoff
}
func Roll6() int {
	return RINT(1, 7)
}
func Roll12() int {
	return RINT(1, 13)
}
func Roll18() int {
	return RINT(1, 19)
}
func Roll24() int {
	return RINT(1, 25)
}
func Roll30() int {
	return RINT(1, 31)
}
func Roll36() int {
	return RINT(1, 37)
}
func RINT(min, max int) int {
	return min + rand.IntN(max-min)
}
func RUINT8(min2, max2 uint8) uint8 {
	min := int(min2)
	max := int(max2)
	return uint8(min + rand.IntN(max-min))
}
func RI32(min, max int) int32 {
	return int32(min + rand.IntN(max-min))
}
func RF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
func RF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func StF32(t float32) string {
	return fmt.Sprint(t)
}
func StF32NODEC(t float32) string {
	return fmt.Sprintf("%.0f", t)
}
func StINT(t int) string {
	return fmt.Sprint(t)
}
func StINT32(t int32) string {
	return fmt.Sprint(t)
}

// MARK: COLORS
func colorOpacity(c []uint8, opacity uint8) (uint8, uint8, uint8, uint8) {
	return c[0], c[1], c[2], opacity
}

func color2array(c1 uint8, c2 uint8, c3 uint8, c4 uint8) []uint8 {
	var colors []uint8
	colors = append(colors, c1, c2, c3, c4)
	return colors
}

// RANDOM COLORS
func COLORǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255}
}
func DARKGREENǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255}
}
func GREENǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255}
}
func REDǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255}
}
func PINKǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255}
}
func BLUEǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255}
}
func DARKBLUEǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255}
}
func CYANǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255}
}
func YELLOWǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255}
}
func ORANGEǁRANDOMǁ3() sdl.Color {
	return sdl.Color{255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255}
}
func BROWNǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255}
}
func GREYǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255}
}
func DARKGREYǁRANDOMǁ3() sdl.Color {
	return sdl.Color{uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255}
}

func COLORǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255
}
func DARKGREENǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255
}
func GREENǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255
}
func REDǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255
}
func PINKǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255
}
func BLUEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255
}
func DARKBLUEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255
}
func CYANǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255
}
func YELLOWǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255
}
func ORANGEǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return 255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255
}
func BROWNǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255
}
func GREYǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255
}
func DARKGREYǁRANDOMǁ2() (uint8, uint8, uint8, uint8) {
	return uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255
}

func COLORǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255}
}
func DARKGREENǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255}
}
func GREENǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255}
}
func REDǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255}
}
func PINKǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255}
}
func BLUEǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255}
}
func DARKBLUEǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255}
}
func CYANǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255}
}
func YELLOWǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255}
}
func ORANGEǁRANDOM() (rgba []uint8) {
	return []uint8{255, uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255}
}
func BROWNǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255}
}
func GREYǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255}
}
func DARKGREYǁRANDOM() (rgba []uint8) {
	return []uint8{uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255}
}

// SOLID COLORS
func MAROONǁ3() sdl.Color               { return sdl.Color{uint8(128), uint8(0), uint8(0), 255} }
func DARKREDǁ3() sdl.Color              { return sdl.Color{uint8(139), uint8(0), uint8(0), 255} }
func BROWNǁ3() sdl.Color                { return sdl.Color{uint8(165), uint8(42), uint8(42), 255} }
func FIREBRICKǁ3() sdl.Color            { return sdl.Color{uint8(178), uint8(34), uint8(34), 255} }
func CRIMSONǁ3() sdl.Color              { return sdl.Color{uint8(220), uint8(20), uint8(60), 255} }
func REDǁ3() sdl.Color                  { return sdl.Color{uint8(255), uint8(0), uint8(0), 255} }
func TOMATOǁ3() sdl.Color               { return sdl.Color{uint8(255), uint8(99), uint8(71), 255} }
func CORALǁ3() sdl.Color                { return sdl.Color{uint8(255), uint8(127), uint8(80), 255} }
func INDIANREDǁ3() sdl.Color            { return sdl.Color{uint8(205), uint8(92), uint8(92), 255} }
func LIGHTCORALǁ3() sdl.Color           { return sdl.Color{uint8(240), uint8(128), uint8(128), 255} }
func DARKSALMONǁ3() sdl.Color           { return sdl.Color{uint8(233), uint8(150), uint8(122), 255} }
func SALMONǁ3() sdl.Color               { return sdl.Color{uint8(250), uint8(128), uint8(114), 255} }
func LIGHTSALMONǁ3() sdl.Color          { return sdl.Color{uint8(255), uint8(160), uint8(122), 255} }
func ORANGEREDǁ3() sdl.Color            { return sdl.Color{uint8(255), uint8(69), uint8(0), 255} }
func DARKORANGEǁ3() sdl.Color           { return sdl.Color{uint8(255), uint8(140), uint8(0), 255} }
func ORANGEǁ3() sdl.Color               { return sdl.Color{uint8(255), uint8(165), uint8(0), 255} }
func GOLDǁ3() sdl.Color                 { return sdl.Color{uint8(255), uint8(215), uint8(0), 255} }
func DARKGOLDENRODǁ3() sdl.Color        { return sdl.Color{uint8(184), uint8(134), uint8(11), 255} }
func GOLDENRODǁ3() sdl.Color            { return sdl.Color{uint8(218), uint8(165), uint8(32), 255} }
func PALEGOLDENRODǁ3() sdl.Color        { return sdl.Color{uint8(238), uint8(232), uint8(170), 255} }
func DARKKHAKIǁ3() sdl.Color            { return sdl.Color{uint8(189), uint8(183), uint8(107), 255} }
func KHAKIǁ3() sdl.Color                { return sdl.Color{uint8(240), uint8(230), uint8(140), 255} }
func OLIVEǁ3() sdl.Color                { return sdl.Color{uint8(128), uint8(128), uint8(0), 255} }
func YELLOWǁ3() sdl.Color               { return sdl.Color{uint8(255), uint8(255), uint8(0), 255} }
func YELLOWGREENǁ3() sdl.Color          { return sdl.Color{uint8(154), uint8(205), uint8(50), 255} }
func DARKOLIVEGREENǁ3() sdl.Color       { return sdl.Color{uint8(85), uint8(107), uint8(47), 255} }
func OLIVEDRABǁ3() sdl.Color            { return sdl.Color{uint8(107), uint8(142), uint8(35), 255} }
func LAWNGREENǁ3() sdl.Color            { return sdl.Color{uint8(124), uint8(252), uint8(0), 255} }
func CHARTREUSEǁ3() sdl.Color           { return sdl.Color{uint8(127), uint8(255), uint8(0), 255} }
func GREENYELLOWǁ3() sdl.Color          { return sdl.Color{uint8(173), uint8(255), uint8(47), 255} }
func DARKGREENǁ3() sdl.Color            { return sdl.Color{uint8(0), uint8(100), uint8(0), 255} }
func GREENǁ3() sdl.Color                { return sdl.Color{uint8(0), uint8(128), uint8(0), 255} }
func FORESTGREENǁ3() sdl.Color          { return sdl.Color{uint8(34), uint8(139), uint8(34), 255} }
func LIMEǁ3() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(0), 255} }
func LIMEGREENǁ3() sdl.Color            { return sdl.Color{uint8(50), uint8(205), uint8(50), 255} }
func LIGHTGREENǁ3() sdl.Color           { return sdl.Color{uint8(144), uint8(238), uint8(144), 255} }
func PALEGREENǁ3() sdl.Color            { return sdl.Color{uint8(152), uint8(251), uint8(152), 255} }
func DARKSEAGREENǁ3() sdl.Color         { return sdl.Color{uint8(143), uint8(188), uint8(143), 255} }
func MEDIUMSPRINGGREENǁ3() sdl.Color    { return sdl.Color{uint8(0), uint8(250), uint8(154), 255} }
func SPRINGGREENǁ3() sdl.Color          { return sdl.Color{uint8(0), uint8(255), uint8(127), 255} }
func SEAGREENǁ3() sdl.Color             { return sdl.Color{uint8(46), uint8(139), uint8(87), 255} }
func MEDIUMAQUAMARINEǁ3() sdl.Color     { return sdl.Color{uint8(102), uint8(205), uint8(170), 255} }
func MEDIUMSEAGREENǁ3() sdl.Color       { return sdl.Color{uint8(60), uint8(179), uint8(113), 255} }
func LIGHTSEAGREENǁ3() sdl.Color        { return sdl.Color{uint8(32), uint8(178), uint8(170), 255} }
func DARKSLATEGRAYǁ3() sdl.Color        { return sdl.Color{uint8(47), uint8(79), uint8(79), 255} }
func TEALǁ3() sdl.Color                 { return sdl.Color{uint8(0), uint8(128), uint8(128), 255} }
func DARKCYANǁ3() sdl.Color             { return sdl.Color{uint8(0), uint8(139), uint8(139), 255} }
func AQUAǁ3() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(255), 255} }
func CYANǁ3() sdl.Color                 { return sdl.Color{uint8(0), uint8(255), uint8(255), 255} }
func LIGHTCYANǁ3() sdl.Color            { return sdl.Color{uint8(224), uint8(255), uint8(255), 255} }
func DARKTURQUOISEǁ3() sdl.Color        { return sdl.Color{uint8(0), uint8(206), uint8(209), 255} }
func TURQUOISEǁ3() sdl.Color            { return sdl.Color{uint8(64), uint8(224), uint8(208), 255} }
func MEDIUMTURQUOISEǁ3() sdl.Color      { return sdl.Color{uint8(72), uint8(209), uint8(204), 255} }
func PALETURQUOISEǁ3() sdl.Color        { return sdl.Color{uint8(175), uint8(238), uint8(238), 255} }
func AQUAMARINEǁ3() sdl.Color           { return sdl.Color{uint8(127), uint8(255), uint8(212), 255} }
func POWDERBLUEǁ3() sdl.Color           { return sdl.Color{uint8(176), uint8(224), uint8(230), 255} }
func CADETBLUEǁ3() sdl.Color            { return sdl.Color{uint8(95), uint8(158), uint8(160), 255} }
func STEELBLUEǁ3() sdl.Color            { return sdl.Color{uint8(70), uint8(130), uint8(180), 255} }
func CORNFLOWERBLUEǁ3() sdl.Color       { return sdl.Color{uint8(100), uint8(149), uint8(237), 255} }
func DEEPSKYBLUEǁ3() sdl.Color          { return sdl.Color{uint8(0), uint8(191), uint8(255), 255} }
func DODGERBLUEǁ3() sdl.Color           { return sdl.Color{uint8(30), uint8(144), uint8(255), 255} }
func LIGHTBLUEǁ3() sdl.Color            { return sdl.Color{uint8(173), uint8(216), uint8(230), 255} }
func SKYBLUEǁ3() sdl.Color              { return sdl.Color{uint8(135), uint8(206), uint8(235), 255} }
func LIGHTSKYBLUEǁ3() sdl.Color         { return sdl.Color{uint8(135), uint8(206), uint8(250), 255} }
func MIDNIGHTBLUEǁ3() sdl.Color         { return sdl.Color{uint8(25), uint8(25), uint8(112), 255} }
func NAVYǁ3() sdl.Color                 { return sdl.Color{uint8(0), uint8(0), uint8(128), 255} }
func DARKBLUEǁ3() sdl.Color             { return sdl.Color{uint8(0), uint8(0), uint8(139), 255} }
func MEDIUMBLUEǁ3() sdl.Color           { return sdl.Color{uint8(0), uint8(0), uint8(205), 255} }
func BLUEǁ3() sdl.Color                 { return sdl.Color{uint8(0), uint8(0), uint8(255), 255} }
func ROYALBLUEǁ3() sdl.Color            { return sdl.Color{uint8(65), uint8(105), uint8(225), 255} }
func BLUEVIOLETǁ3() sdl.Color           { return sdl.Color{uint8(138), uint8(43), uint8(226), 255} }
func INDIGOǁ3() sdl.Color               { return sdl.Color{uint8(75), uint8(0), uint8(130), 255} }
func DARKSLATEBLUEǁ3() sdl.Color        { return sdl.Color{uint8(72), uint8(61), uint8(139), 255} }
func SLATEBLUEǁ3() sdl.Color            { return sdl.Color{uint8(106), uint8(90), uint8(205), 255} }
func MEDIUMSLATEBLUEǁ3() sdl.Color      { return sdl.Color{uint8(123), uint8(104), uint8(238), 255} }
func MEDIUMPURPLEǁ3() sdl.Color         { return sdl.Color{uint8(147), uint8(112), uint8(219), 255} }
func DARKMAGENTAǁ3() sdl.Color          { return sdl.Color{uint8(139), uint8(0), uint8(139), 255} }
func DARKVIOLETǁ3() sdl.Color           { return sdl.Color{uint8(148), uint8(0), uint8(211), 255} }
func DARKORCHIDǁ3() sdl.Color           { return sdl.Color{uint8(153), uint8(50), uint8(204), 255} }
func MEDIUMORCHIDǁ3() sdl.Color         { return sdl.Color{uint8(186), uint8(85), uint8(211), 255} }
func PURPLEǁ3() sdl.Color               { return sdl.Color{uint8(128), uint8(0), uint8(128), 255} }
func THISTLEǁ3() sdl.Color              { return sdl.Color{uint8(216), uint8(191), uint8(216), 255} }
func PLUMǁ3() sdl.Color                 { return sdl.Color{uint8(221), uint8(160), uint8(221), 255} }
func VIOLETǁ3() sdl.Color               { return sdl.Color{uint8(238), uint8(130), uint8(238), 255} }
func MAGENTAǁ3() sdl.Color              { return sdl.Color{uint8(255), uint8(0), uint8(255), 255} }
func ORCHIDǁ3() sdl.Color               { return sdl.Color{uint8(218), uint8(112), uint8(214), 255} }
func MEDIUMVIOLETREDǁ3() sdl.Color      { return sdl.Color{uint8(199), uint8(21), uint8(133), 255} }
func PALEVIOLETREDǁ3() sdl.Color        { return sdl.Color{uint8(219), uint8(112), uint8(147), 255} }
func DEEPPINKǁ3() sdl.Color             { return sdl.Color{uint8(255), uint8(20), uint8(147), 255} }
func HOTPINKǁ3() sdl.Color              { return sdl.Color{uint8(255), uint8(105), uint8(180), 255} }
func LIGHTPINKǁ3() sdl.Color            { return sdl.Color{uint8(255), uint8(182), uint8(193), 255} }
func PINKǁ3() sdl.Color                 { return sdl.Color{uint8(255), uint8(192), uint8(203), 255} }
func ANTIQUEWHITEǁ3() sdl.Color         { return sdl.Color{uint8(250), uint8(235), uint8(215), 255} }
func BEIGEǁ3() sdl.Color                { return sdl.Color{uint8(245), uint8(245), uint8(220), 255} }
func BISQUEǁ3() sdl.Color               { return sdl.Color{uint8(255), uint8(228), uint8(196), 255} }
func BLANCHEDALMONDǁ3() sdl.Color       { return sdl.Color{uint8(255), uint8(235), uint8(205), 255} }
func WHEATǁ3() sdl.Color                { return sdl.Color{uint8(245), uint8(222), uint8(179), 255} }
func CORNSILKǁ3() sdl.Color             { return sdl.Color{uint8(255), uint8(248), uint8(220), 255} }
func LEMONCHIFFONǁ3() sdl.Color         { return sdl.Color{uint8(255), uint8(250), uint8(205), 255} }
func LIGHTGOLDENRODYELLOWǁ3() sdl.Color { return sdl.Color{uint8(250), uint8(250), uint8(210), 255} }
func LIGHTYELLOWǁ3() sdl.Color          { return sdl.Color{uint8(255), uint8(255), uint8(224), 255} }
func SADDLEBROWNǁ3() sdl.Color          { return sdl.Color{uint8(139), uint8(69), uint8(19), 255} }
func SIENNAǁ3() sdl.Color               { return sdl.Color{uint8(160), uint8(82), uint8(45), 255} }
func CHOCOLATEǁ3() sdl.Color            { return sdl.Color{uint8(210), uint8(105), uint8(30), 255} }
func PERUǁ3() sdl.Color                 { return sdl.Color{uint8(205), uint8(133), uint8(63), 255} }
func SANDYBROWNǁ3() sdl.Color           { return sdl.Color{uint8(244), uint8(164), uint8(96), 255} }
func BURLYWOODǁ3() sdl.Color            { return sdl.Color{uint8(222), uint8(184), uint8(135), 255} }
func TANǁ3() sdl.Color                  { return sdl.Color{uint8(210), uint8(180), uint8(140), 255} }
func ROSYBROWNǁ3() sdl.Color            { return sdl.Color{uint8(188), uint8(143), uint8(143), 255} }
func MOCCASINǁ3() sdl.Color             { return sdl.Color{uint8(255), uint8(228), uint8(181), 255} }
func NAVAJOWHITEǁ3() sdl.Color          { return sdl.Color{uint8(255), uint8(222), uint8(173), 255} }
func PEACHPUFFǁ3() sdl.Color            { return sdl.Color{uint8(255), uint8(218), uint8(185), 255} }
func MISTYROSEǁ3() sdl.Color            { return sdl.Color{uint8(255), uint8(228), uint8(225), 255} }
func LAVENDERBLUSHǁ3() sdl.Color        { return sdl.Color{uint8(255), uint8(240), uint8(245), 255} }
func LINENǁ3() sdl.Color                { return sdl.Color{uint8(250), uint8(240), uint8(230), 255} }
func OLDLACEǁ3() sdl.Color              { return sdl.Color{uint8(253), uint8(245), uint8(230), 255} }
func PAPAYAWHIPǁ3() sdl.Color           { return sdl.Color{uint8(255), uint8(239), uint8(213), 255} }
func SEASHELLǁ3() sdl.Color             { return sdl.Color{uint8(255), uint8(245), uint8(238), 255} }
func MINTCREAMǁ3() sdl.Color            { return sdl.Color{uint8(245), uint8(255), uint8(250), 255} }
func SLATEGRAYǁ3() sdl.Color            { return sdl.Color{uint8(112), uint8(128), uint8(144), 255} }
func LIGHTSLATEGRAYǁ3() sdl.Color       { return sdl.Color{uint8(119), uint8(136), uint8(153), 255} }
func LIGHTSTEELBLUEǁ3() sdl.Color       { return sdl.Color{uint8(176), uint8(196), uint8(222), 255} }
func LAVENDERǁ3() sdl.Color             { return sdl.Color{uint8(230), uint8(230), uint8(250), 255} }
func FLORALWHITEǁ3() sdl.Color          { return sdl.Color{uint8(255), uint8(250), uint8(240), 255} }
func ALICEBLUEǁ3() sdl.Color            { return sdl.Color{uint8(240), uint8(248), uint8(255), 255} }
func GHOSTWHITEǁ3() sdl.Color           { return sdl.Color{uint8(248), uint8(248), uint8(255), 255} }
func HONEYDEWǁ3() sdl.Color             { return sdl.Color{uint8(240), uint8(255), uint8(240), 255} }
func IVORYǁ3() sdl.Color                { return sdl.Color{uint8(255), uint8(255), uint8(240), 255} }
func AZUREǁ3() sdl.Color                { return sdl.Color{uint8(240), uint8(255), uint8(255), 255} }
func SNOWǁ3() sdl.Color                 { return sdl.Color{uint8(255), uint8(250), uint8(250), 255} }
func BLACKǁ3() sdl.Color                { return sdl.Color{uint8(0), uint8(0), uint8(0), 255} }
func DIMGREYǁ3() sdl.Color              { return sdl.Color{uint8(105), uint8(105), uint8(105), 255} }
func GREYǁ3() sdl.Color                 { return sdl.Color{uint8(128), uint8(128), uint8(128), 255} }
func DARKGREYǁ3() sdl.Color             { return sdl.Color{uint8(169), uint8(169), uint8(169), 255} }
func SILVERǁ3() sdl.Color               { return sdl.Color{uint8(192), uint8(192), uint8(192), 255} }
func LIGHTGREYǁ3() sdl.Color            { return sdl.Color{uint8(211), uint8(211), uint8(211), 255} }
func GAINSBOROǁ3() sdl.Color            { return sdl.Color{uint8(220), uint8(220), uint8(220), 255} }
func WHITESMOKEǁ3() sdl.Color           { return sdl.Color{uint8(245), uint8(245), uint8(245), 255} }
func WHITEǁ3() sdl.Color                { return sdl.Color{uint8(255), uint8(255), uint8(255), 255} }

func MAROONǁ2() (uint8, uint8, uint8, uint8)               { return 128, 0, 0, 255 }
func DARKREDǁ2() (uint8, uint8, uint8, uint8)              { return 139, 0, 0, 255 }
func BROWNǁ2() (uint8, uint8, uint8, uint8)                { return 165, 42, 42, 255 }
func FIREBRICKǁ2() (uint8, uint8, uint8, uint8)            { return 178, 34, 34, 255 }
func CRIMSONǁ2() (uint8, uint8, uint8, uint8)              { return 220, 20, 60, 255 }
func REDǁ2() (uint8, uint8, uint8, uint8)                  { return 255, 0, 0, 255 }
func TOMATOǁ2() (uint8, uint8, uint8, uint8)               { return 255, 99, 71, 255 }
func CORALǁ2() (uint8, uint8, uint8, uint8)                { return 255, 127, 80, 255 }
func INDIANREDǁ2() (uint8, uint8, uint8, uint8)            { return 205, 92, 92, 255 }
func LIGHTCORALǁ2() (uint8, uint8, uint8, uint8)           { return 240, 128, 128, 255 }
func DARKSALMONǁ2() (uint8, uint8, uint8, uint8)           { return 233, 150, 122, 255 }
func SALMONǁ2() (uint8, uint8, uint8, uint8)               { return 250, 128, 114, 255 }
func LIGHTSALMONǁ2() (uint8, uint8, uint8, uint8)          { return 255, 160, 122, 255 }
func ORANGEREDǁ2() (uint8, uint8, uint8, uint8)            { return 255, 69, 0, 255 }
func DARKORANGEǁ2() (uint8, uint8, uint8, uint8)           { return 255, 140, 0, 255 }
func ORANGEǁ2() (uint8, uint8, uint8, uint8)               { return 255, 165, 0, 255 }
func GOLDǁ2() (uint8, uint8, uint8, uint8)                 { return 255, 215, 0, 255 }
func DARKGOLDENRODǁ2() (uint8, uint8, uint8, uint8)        { return 184, 134, 11, 255 }
func GOLDENRODǁ2() (uint8, uint8, uint8, uint8)            { return 218, 165, 32, 255 }
func PALEGOLDENRODǁ2() (uint8, uint8, uint8, uint8)        { return 238, 232, 170, 255 }
func DARKKHAKIǁ2() (uint8, uint8, uint8, uint8)            { return 189, 183, 107, 255 }
func KHAKIǁ2() (uint8, uint8, uint8, uint8)                { return 240, 230, 140, 255 }
func OLIVEǁ2() (uint8, uint8, uint8, uint8)                { return 128, 128, 0, 255 }
func YELLOWǁ2() (uint8, uint8, uint8, uint8)               { return 255, 255, 0, 255 }
func YELLOWGREENǁ2() (uint8, uint8, uint8, uint8)          { return 154, 205, 50, 255 }
func DARKOLIVEGREENǁ2() (uint8, uint8, uint8, uint8)       { return 85, 107, 47, 255 }
func OLIVEDRABǁ2() (uint8, uint8, uint8, uint8)            { return 107, 142, 35, 255 }
func LAWNGREENǁ2() (uint8, uint8, uint8, uint8)            { return 124, 252, 0, 255 }
func CHARTREUSEǁ2() (uint8, uint8, uint8, uint8)           { return 127, 255, 0, 255 }
func GREENYELLOWǁ2() (uint8, uint8, uint8, uint8)          { return 173, 255, 47, 255 }
func DARKGREENǁ2() (uint8, uint8, uint8, uint8)            { return 0, 100, 0, 255 }
func GREENǁ2() (uint8, uint8, uint8, uint8)                { return 0, 128, 0, 255 }
func FORESTGREENǁ2() (uint8, uint8, uint8, uint8)          { return 34, 139, 34, 255 }
func LIMEǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 255, 0, 255 }
func LIMEGREENǁ2() (uint8, uint8, uint8, uint8)            { return 50, 205, 50, 255 }
func LIGHTGREENǁ2() (uint8, uint8, uint8, uint8)           { return 144, 238, 144, 255 }
func PALEGREENǁ2() (uint8, uint8, uint8, uint8)            { return 152, 251, 152, 255 }
func DARKSEAGREENǁ2() (uint8, uint8, uint8, uint8)         { return 143, 188, 143, 255 }
func MEDIUMSPRINGGREENǁ2() (uint8, uint8, uint8, uint8)    { return 0, 250, 154, 255 }
func SPRINGGREENǁ2() (uint8, uint8, uint8, uint8)          { return 0, 255, 127, 255 }
func SEAGREENǁ2() (uint8, uint8, uint8, uint8)             { return 46, 139, 87, 255 }
func MEDIUMAQUAMARINEǁ2() (uint8, uint8, uint8, uint8)     { return 102, 205, 170, 255 }
func MEDIUMSEAGREENǁ2() (uint8, uint8, uint8, uint8)       { return 60, 179, 113, 255 }
func LIGHTSEAGREENǁ2() (uint8, uint8, uint8, uint8)        { return 32, 178, 170, 255 }
func DARKSLATEGRAYǁ2() (uint8, uint8, uint8, uint8)        { return 47, 79, 79, 255 }
func TEALǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 128, 128, 255 }
func DARKCYANǁ2() (uint8, uint8, uint8, uint8)             { return 0, 139, 139, 255 }
func AQUAǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 255, 255, 255 }
func CYANǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 255, 255, 255 }
func LIGHTCYANǁ2() (uint8, uint8, uint8, uint8)            { return 224, 255, 255, 255 }
func DARKTURQUOISEǁ2() (uint8, uint8, uint8, uint8)        { return 0, 206, 209, 255 }
func TURQUOISEǁ2() (uint8, uint8, uint8, uint8)            { return 64, 224, 208, 255 }
func MEDIUMTURQUOISEǁ2() (uint8, uint8, uint8, uint8)      { return 72, 209, 204, 255 }
func PALETURQUOISEǁ2() (uint8, uint8, uint8, uint8)        { return 175, 238, 238, 255 }
func AQUAMARINEǁ2() (uint8, uint8, uint8, uint8)           { return 127, 255, 212, 255 }
func POWDERBLUEǁ2() (uint8, uint8, uint8, uint8)           { return 176, 224, 230, 255 }
func CADETBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 95, 158, 160, 255 }
func STEELBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 70, 130, 180, 255 }
func CORNFLOWERBLUEǁ2() (uint8, uint8, uint8, uint8)       { return 100, 149, 237, 255 }
func DEEPSKYBLUEǁ2() (uint8, uint8, uint8, uint8)          { return 0, 191, 255, 255 }
func DODGERBLUEǁ2() (uint8, uint8, uint8, uint8)           { return 30, 144, 255, 255 }
func LIGHTBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 173, 216, 230, 255 }
func SKYBLUEǁ2() (uint8, uint8, uint8, uint8)              { return 135, 206, 235, 255 }
func LIGHTSKYBLUEǁ2() (uint8, uint8, uint8, uint8)         { return 135, 206, 250, 255 }
func MIDNIGHTBLUEǁ2() (uint8, uint8, uint8, uint8)         { return 25, 25, 112, 255 }
func NAVYǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 0, 128, 255 }
func DARKBLUEǁ2() (uint8, uint8, uint8, uint8)             { return 0, 0, 139, 255 }
func MEDIUMBLUEǁ2() (uint8, uint8, uint8, uint8)           { return 0, 0, 205, 255 }
func BLUEǁ2() (uint8, uint8, uint8, uint8)                 { return 0, 0, 255, 255 }
func ROYALBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 65, 105, 225, 255 }
func BLUEVIOLETǁ2() (uint8, uint8, uint8, uint8)           { return 138, 43, 226, 255 }
func INDIGOǁ2() (uint8, uint8, uint8, uint8)               { return 75, 0, 130, 255 }
func DARKSLATEBLUEǁ2() (uint8, uint8, uint8, uint8)        { return 72, 61, 139, 255 }
func SLATEBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 106, 90, 205, 255 }
func MEDIUMSLATEBLUEǁ2() (uint8, uint8, uint8, uint8)      { return 123, 104, 238, 255 }
func MEDIUMPURPLEǁ2() (uint8, uint8, uint8, uint8)         { return 147, 112, 219, 255 }
func DARKMAGENTAǁ2() (uint8, uint8, uint8, uint8)          { return 139, 0, 139, 255 }
func DARKVIOLETǁ2() (uint8, uint8, uint8, uint8)           { return 148, 0, 211, 255 }
func DARKORCHIDǁ2() (uint8, uint8, uint8, uint8)           { return 153, 50, 204, 255 }
func MEDIUMORCHIDǁ2() (uint8, uint8, uint8, uint8)         { return 186, 85, 211, 255 }
func PURPLEǁ2() (uint8, uint8, uint8, uint8)               { return 128, 0, 128, 255 }
func THISTLEǁ2() (uint8, uint8, uint8, uint8)              { return 216, 191, 216, 255 }
func PLUMǁ2() (uint8, uint8, uint8, uint8)                 { return 221, 160, 221, 255 }
func VIOLETǁ2() (uint8, uint8, uint8, uint8)               { return 238, 130, 238, 255 }
func MAGENTAǁ2() (uint8, uint8, uint8, uint8)              { return 255, 0, 255, 255 }
func ORCHIDǁ2() (uint8, uint8, uint8, uint8)               { return 218, 112, 214, 255 }
func MEDIUMVIOLETREDǁ2() (uint8, uint8, uint8, uint8)      { return 199, 21, 133, 255 }
func PALEVIOLETREDǁ2() (uint8, uint8, uint8, uint8)        { return 219, 112, 147, 255 }
func DEEPPINKǁ2() (uint8, uint8, uint8, uint8)             { return 255, 20, 147, 255 }
func HOTPINKǁ2() (uint8, uint8, uint8, uint8)              { return 255, 105, 180, 255 }
func LIGHTPINKǁ2() (uint8, uint8, uint8, uint8)            { return 255, 182, 193, 255 }
func PINKǁ2() (uint8, uint8, uint8, uint8)                 { return 255, 192, 203, 255 }
func ANTIQUEWHITEǁ2() (uint8, uint8, uint8, uint8)         { return 250, 235, 215, 255 }
func BEIGEǁ2() (uint8, uint8, uint8, uint8)                { return 245, 245, 220, 255 }
func BISQUEǁ2() (uint8, uint8, uint8, uint8)               { return 255, 228, 196, 255 }
func BLANCHEDALMONDǁ2() (uint8, uint8, uint8, uint8)       { return 255, 235, 205, 255 }
func WHEATǁ2() (uint8, uint8, uint8, uint8)                { return 245, 222, 179, 255 }
func CORNSILKǁ2() (uint8, uint8, uint8, uint8)             { return 255, 248, 220, 255 }
func LEMONCHIFFONǁ2() (uint8, uint8, uint8, uint8)         { return 255, 250, 205, 255 }
func LIGHTGOLDENRODYELLOWǁ2() (uint8, uint8, uint8, uint8) { return 250, 250, 210, 255 }
func LIGHTYELLOWǁ2() (uint8, uint8, uint8, uint8)          { return 255, 255, 224, 255 }
func SADDLEBROWNǁ2() (uint8, uint8, uint8, uint8)          { return 139, 69, 19, 255 }
func SIENNAǁ2() (uint8, uint8, uint8, uint8)               { return 160, 82, 45, 255 }
func CHOCOLATEǁ2() (uint8, uint8, uint8, uint8)            { return 210, 105, 30, 255 }
func PERUǁ2() (uint8, uint8, uint8, uint8)                 { return 205, 133, 63, 255 }
func SANDYBROWNǁ2() (uint8, uint8, uint8, uint8)           { return 244, 164, 96, 255 }
func BURLYWOODǁ2() (uint8, uint8, uint8, uint8)            { return 222, 184, 135, 255 }
func TANǁ2() (uint8, uint8, uint8, uint8)                  { return 210, 180, 140, 255 }
func ROSYBROWNǁ2() (uint8, uint8, uint8, uint8)            { return 188, 143, 143, 255 }
func MOCCASINǁ2() (uint8, uint8, uint8, uint8)             { return 255, 228, 181, 255 }
func NAVAJOWHITEǁ2() (uint8, uint8, uint8, uint8)          { return 255, 222, 173, 255 }
func PEACHPUFFǁ2() (uint8, uint8, uint8, uint8)            { return 255, 218, 185, 255 }
func MISTYROSEǁ2() (uint8, uint8, uint8, uint8)            { return 255, 228, 225, 255 }
func LAVENDERBLUSHǁ2() (uint8, uint8, uint8, uint8)        { return 255, 240, 245, 255 }
func LINENǁ2() (uint8, uint8, uint8, uint8)                { return 250, 240, 230, 255 }
func OLDLACEǁ2() (uint8, uint8, uint8, uint8)              { return 253, 245, 230, 255 }
func PAPAYAWHIPǁ2() (uint8, uint8, uint8, uint8)           { return 255, 239, 213, 255 }
func SEASHELLǁ2() (uint8, uint8, uint8, uint8)             { return 255, 245, 238, 255 }
func MINTCREAMǁ2() (uint8, uint8, uint8, uint8)            { return 245, 255, 250, 255 }
func SLATEGRAYǁ2() (uint8, uint8, uint8, uint8)            { return 112, 128, 144, 255 }
func LIGHTSLATEGRAYǁ2() (uint8, uint8, uint8, uint8)       { return 119, 136, 153, 255 }
func LIGHTSTEELBLUEǁ2() (uint8, uint8, uint8, uint8)       { return 176, 196, 222, 255 }
func LAVENDERǁ2() (uint8, uint8, uint8, uint8)             { return 230, 230, 250, 255 }
func FLORALWHITEǁ2() (uint8, uint8, uint8, uint8)          { return 255, 250, 240, 255 }
func ALICEBLUEǁ2() (uint8, uint8, uint8, uint8)            { return 240, 248, 255, 255 }
func GHOSTWHITEǁ2() (uint8, uint8, uint8, uint8)           { return 248, 248, 255, 255 }
func HONEYDEWǁ2() (uint8, uint8, uint8, uint8)             { return 240, 255, 240, 255 }
func IVORYǁ2() (uint8, uint8, uint8, uint8)                { return 255, 255, 240, 255 }
func AZUREǁ2() (uint8, uint8, uint8, uint8)                { return 240, 255, 255, 255 }
func SNOWǁ2() (uint8, uint8, uint8, uint8)                 { return 255, 250, 250, 255 }
func BLACKǁ2() (uint8, uint8, uint8, uint8)                { return 0, 0, 0, 255 }
func DIMGREYǁ2() (uint8, uint8, uint8, uint8)              { return 105, 105, 105, 255 }
func GREYǁ2() (uint8, uint8, uint8, uint8)                 { return 128, 128, 128, 255 }
func DARKGREYǁ2() (uint8, uint8, uint8, uint8)             { return 169, 169, 169, 255 }
func SILVERǁ2() (uint8, uint8, uint8, uint8)               { return 192, 192, 192, 255 }
func LIGHTGREYǁ2() (uint8, uint8, uint8, uint8)            { return 211, 211, 211, 255 }
func GAINSBOROǁ2() (uint8, uint8, uint8, uint8)            { return 220, 220, 220, 255 }
func WHITESMOKEǁ2() (uint8, uint8, uint8, uint8)           { return 245, 245, 245, 255 }
func WHITEǁ2() (uint8, uint8, uint8, uint8)                { return 255, 255, 255, 255 }

func MAROON() (rgba []uint8)         { return []uint8{128, 0, 0, 255} }
func DARKRED() (rgba []uint8)        { return []uint8{139, 0, 0, 255} }
func BROWN() (rgba []uint8)          { return []uint8{165, 42, 42, 255} }
func FIREBRICK() (rgba []uint8)      { return []uint8{178, 34, 34, 255} }
func CRIMSON() (rgba []uint8)        { return []uint8{220, 20, 60, 255} }
func RED() (rgba []uint8)            { return []uint8{255, 0, 0, 255} }
func TOMATO() (rgba []uint8)         { return []uint8{255, 99, 71, 255} }
func CORAL() (rgba []uint8)          { return []uint8{255, 127, 80, 255} }
func INDIANRED() (rgba []uint8)      { return []uint8{205, 92, 92, 255} }
func LIGHTCORAL() (rgba []uint8)     { return []uint8{240, 128, 128, 255} }
func DARKSALMON() (rgba []uint8)     { return []uint8{233, 150, 122, 255} }
func SALMON() (rgba []uint8)         { return []uint8{250, 128, 114, 255} }
func LIGHTSALMON() (rgba []uint8)    { return []uint8{255, 160, 122, 255} }
func ORANGERED() (rgba []uint8)      { return []uint8{255, 69, 0, 255} }
func DARKORANGE() (rgba []uint8)     { return []uint8{255, 140, 0, 255} }
func ORANGE() (rgba []uint8)         { return []uint8{255, 165, 0, 255} }
func GOLD() (rgba []uint8)           { return []uint8{255, 215, 0, 255} }
func DARKGOLDENROD() (rgba []uint8)  { return []uint8{184, 134, 11, 255} }
func GOLDENROD() (rgba []uint8)      { return []uint8{218, 165, 32, 255} }
func PALEGOLDENROD() (rgba []uint8)  { return []uint8{238, 232, 170, 255} }
func DARKKHAKI() (rgba []uint8)      { return []uint8{189, 183, 107, 255} }
func KHAKI() (rgba []uint8)          { return []uint8{240, 230, 140, 255} }
func OLIVE() (rgba []uint8)          { return []uint8{128, 128, 0, 255} }
func YELLOW() (rgba []uint8)         { return []uint8{255, 255, 0, 255} }
func YELLOWGREEN() (rgba []uint8)    { return []uint8{154, 205, 50, 255} }
func DARKOLIVEGREEN() (rgba []uint8) { return []uint8{85, 107, 47, 255} }
func OLIVEDRAB() (rgba []uint8)      { return []uint8{107, 142, 35, 255} }
func LAWNGREEN() (rgba []uint8)      { return []uint8{124, 252, 0, 255} }
func CHARTREUSE() (rgba []uint8)     { return []uint8{127, 255, 0, 255} }
func GREENYELLOW() (rgba []uint8)    { return []uint8{173, 255, 47, 255} }
func DARKGREEN() (rgba []uint8)      { return []uint8{0, 100, 0, 255} }
func GREEN() (rgba []uint8)          { return []uint8{0, 128, 0, 255} }
func FORESTGREEN() (rgba []uint8)    { return []uint8{34, 139, 34, 255} }
func LIME() (rgba []uint8)           { return []uint8{0, 255, 0, 255} }
func LIMEGREEN() (rgba []uint8)      { return []uint8{50, 205, 50, 255} }
func LIGHTGREEN() (rgba []uint8)     { return []uint8{144, 238, 144, 255} }
func PALEGREEN() (rgba []uint8)      { return []uint8{152, 251, 152, 255} }
func DARKSEAGREEN() (rgba []uint8)   { return []uint8{143, 188, 143, 255} }
func MEDIUMSPRINGGREEN() (rgba []uint8) {
	return []uint8{0, 250, 154, 255}
}
func SPRINGGREEN() (rgba []uint8) { return []uint8{0, 255, 127, 255} }
func SEAGREEN() (rgba []uint8)    { return []uint8{46, 139, 87, 255} }
func MEDIUMAQUAMARINE() (rgba []uint8) {
	return []uint8{102, 205, 170, 255}
}
func MEDIUMSEAGREEN() (rgba []uint8)  { return []uint8{60, 179, 113, 255} }
func LIGHTSEAGREEN() (rgba []uint8)   { return []uint8{32, 178, 170, 255} }
func DARKSLATEGRAY() (rgba []uint8)   { return []uint8{47, 79, 79, 255} }
func TEAL() (rgba []uint8)            { return []uint8{0, 128, 128, 255} }
func DARKCYAN() (rgba []uint8)        { return []uint8{0, 139, 139, 255} }
func AQUA() (rgba []uint8)            { return []uint8{0, 255, 255, 255} }
func CYAN() (rgba []uint8)            { return []uint8{0, 255, 255, 255} }
func LIGHTCYAN() (rgba []uint8)       { return []uint8{224, 255, 255, 255} }
func DARKTURQUOISE() (rgba []uint8)   { return []uint8{0, 206, 209, 255} }
func TURQUOISE() (rgba []uint8)       { return []uint8{64, 224, 208, 255} }
func MEDIUMTURQUOISE() (rgba []uint8) { return []uint8{72, 209, 204, 255} }
func PALETURQUOISE() (rgba []uint8)   { return []uint8{175, 238, 238, 255} }
func AQUAMARINE() (rgba []uint8)      { return []uint8{127, 255, 212, 255} }
func POWDERBLUE() (rgba []uint8)      { return []uint8{176, 224, 230, 255} }
func CADETBLUE() (rgba []uint8)       { return []uint8{95, 158, 160, 255} }
func STEELBLUE() (rgba []uint8)       { return []uint8{70, 130, 180, 255} }
func CORNFLOWERBLUE() (rgba []uint8)  { return []uint8{100, 149, 237, 255} }
func DEEPSKYBLUE() (rgba []uint8)     { return []uint8{0, 191, 255, 255} }
func DODGERBLUE() (rgba []uint8)      { return []uint8{30, 144, 255, 255} }
func LIGHTBLUE() (rgba []uint8)       { return []uint8{173, 216, 230, 255} }
func SKYBLUE() (rgba []uint8)         { return []uint8{135, 206, 235, 255} }
func LIGHTSKYBLUE() (rgba []uint8)    { return []uint8{135, 206, 250, 255} }
func MIDNIGHTBLUE() (rgba []uint8)    { return []uint8{25, 25, 112, 255} }
func NAVY() (rgba []uint8)            { return []uint8{0, 0, 128, 255} }
func DARKBLUE() (rgba []uint8)        { return []uint8{0, 0, 139, 255} }
func MEDIUMBLUE() (rgba []uint8)      { return []uint8{0, 0, 205, 255} }
func BLUE() (rgba []uint8)            { return []uint8{0, 0, 255, 255} }
func ROYALBLUE() (rgba []uint8)       { return []uint8{65, 105, 225, 255} }
func BLUEVIOLET() (rgba []uint8)      { return []uint8{138, 43, 226, 255} }
func INDIGO() (rgba []uint8)          { return []uint8{75, 0, 130, 255} }
func DARKSLATEBLUE() (rgba []uint8)   { return []uint8{72, 61, 139, 255} }
func SLATEBLUE() (rgba []uint8)       { return []uint8{106, 90, 205, 255} }
func MEDIUMSLATEBLUE() (rgba []uint8) {
	return []uint8{123, 104, 238, 255}
}
func MEDIUMPURPLE() (rgba []uint8)    { return []uint8{147, 112, 219, 255} }
func DARKMAGENTA() (rgba []uint8)     { return []uint8{139, 0, 139, 255} }
func DARKVIOLET() (rgba []uint8)      { return []uint8{148, 0, 211, 255} }
func DARKORCHID() (rgba []uint8)      { return []uint8{153, 50, 204, 255} }
func MEDIUMORCHID() (rgba []uint8)    { return []uint8{186, 85, 211, 255} }
func PURPLE() (rgba []uint8)          { return []uint8{128, 0, 128, 255} }
func THISTLE() (rgba []uint8)         { return []uint8{216, 191, 216, 255} }
func PLUM() (rgba []uint8)            { return []uint8{221, 160, 221, 255} }
func VIOLET() (rgba []uint8)          { return []uint8{238, 130, 238, 255} }
func MAGENTA() (rgba []uint8)         { return []uint8{255, 0, 255, 255} }
func ORCHID() (rgba []uint8)          { return []uint8{218, 112, 214, 255} }
func MEDIUMVIOLETRED() (rgba []uint8) { return []uint8{199, 21, 133, 255} }
func PALEVIOLETRED() (rgba []uint8)   { return []uint8{219, 112, 147, 255} }
func DEEPPINK() (rgba []uint8)        { return []uint8{255, 20, 147, 255} }
func HOTPINK() (rgba []uint8)         { return []uint8{255, 105, 180, 255} }
func LIGHTPINK() (rgba []uint8)       { return []uint8{255, 182, 193, 255} }
func PINK() (rgba []uint8)            { return []uint8{255, 192, 203, 255} }
func ANTIQUEWHITE() (rgba []uint8)    { return []uint8{250, 235, 215, 255} }
func BEIGE() (rgba []uint8)           { return []uint8{245, 245, 220, 255} }
func BISQUE() (rgba []uint8)          { return []uint8{255, 228, 196, 255} }
func BLANCHEDALMOND() (rgba []uint8)  { return []uint8{255, 235, 205, 255} }
func WHEAT() (rgba []uint8)           { return []uint8{245, 222, 179, 255} }
func CORNSILK() (rgba []uint8)        { return []uint8{255, 248, 220, 255} }
func LEMONCHIFFON() (rgba []uint8)    { return []uint8{255, 250, 205, 255} }
func LIGHTGOLDENRODYELLOW() (rgba []uint8) {
	return []uint8{250, 250, 210, 255}
}
func LIGHTYELLOW() (rgba []uint8)    { return []uint8{255, 255, 224, 255} }
func SADDLEBROWN() (rgba []uint8)    { return []uint8{139, 69, 19, 255} }
func SIENNA() (rgba []uint8)         { return []uint8{160, 82, 45, 255} }
func CHOCOLATE() (rgba []uint8)      { return []uint8{210, 105, 30, 255} }
func PERU() (rgba []uint8)           { return []uint8{205, 133, 63, 255} }
func SANDYBROWN() (rgba []uint8)     { return []uint8{244, 164, 96, 255} }
func BURLYWOOD() (rgba []uint8)      { return []uint8{222, 184, 135, 255} }
func TAN() (rgba []uint8)            { return []uint8{210, 180, 140, 255} }
func ROSYBROWN() (rgba []uint8)      { return []uint8{188, 143, 143, 255} }
func MOCCASIN() (rgba []uint8)       { return []uint8{255, 228, 181, 255} }
func NAVAJOWHITE() (rgba []uint8)    { return []uint8{255, 222, 173, 255} }
func PEACHPUFF() (rgba []uint8)      { return []uint8{255, 218, 185, 255} }
func MISTYROSE() (rgba []uint8)      { return []uint8{255, 228, 225, 255} }
func LAVENDERBLUSH() (rgba []uint8)  { return []uint8{255, 240, 245, 255} }
func LINEN() (rgba []uint8)          { return []uint8{250, 240, 230, 255} }
func OLDLACE() (rgba []uint8)        { return []uint8{253, 245, 230, 255} }
func PAPAYAWHIP() (rgba []uint8)     { return []uint8{255, 239, 213, 255} }
func SEASHELL() (rgba []uint8)       { return []uint8{255, 245, 238, 255} }
func MINTCREAM() (rgba []uint8)      { return []uint8{245, 255, 250, 255} }
func SLATEGRAY() (rgba []uint8)      { return []uint8{112, 128, 144, 255} }
func LIGHTSLATEGRAY() (rgba []uint8) { return []uint8{119, 136, 153, 255} }
func LIGHTSTEELBLUE() (rgba []uint8) { return []uint8{176, 196, 222, 255} }
func LAVENDER() (rgba []uint8)       { return []uint8{230, 230, 250, 255} }
func FLORALWHITE() (rgba []uint8)    { return []uint8{255, 250, 240, 255} }
func ALICEBLUE() (rgba []uint8)      { return []uint8{240, 248, 255, 255} }
func GHOSTWHITE() (rgba []uint8)     { return []uint8{248, 248, 255, 255} }
func HONEYDEW() (rgba []uint8)       { return []uint8{240, 255, 240, 255} }
func IVORY() (rgba []uint8)          { return []uint8{255, 255, 240, 255} }
func AZURE() (rgba []uint8)          { return []uint8{240, 255, 255, 255} }
func SNOW() (rgba []uint8)           { return []uint8{255, 250, 250, 255} }
func BLACK() (rgba []uint8)          { return []uint8{0, 0, 0, 255} }
func DIMGREY() (rgba []uint8)        { return []uint8{105, 105, 105, 255} }
func GREY() (rgba []uint8)           { return []uint8{128, 128, 128, 255} }
func DARKGREY() (rgba []uint8)       { return []uint8{169, 169, 169, 255} }
func SILVER() (rgba []uint8)         { return []uint8{192, 192, 192, 255} }
func LIGHTGREY() (rgba []uint8)      { return []uint8{211, 211, 211, 255} }
func GAINSBORO() (rgba []uint8)      { return []uint8{220, 220, 220, 255} }
func WHITESMOKE() (rgba []uint8)     { return []uint8{245, 245, 245, 255} }
func WHITE() (rgba []uint8)          { return []uint8{255, 255, 255, 255} }
