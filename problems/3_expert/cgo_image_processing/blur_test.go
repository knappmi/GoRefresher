package cgoimage

import "testing"

func TestBlur(t *testing.T) {
	w,h := 3,3
	// 3x3 gradient
	px := []uint8{
		10,20,30,
		40,50,60,
		70,80,90,
	}
	Blur(px,w,h)
	center := px[4]
	// Compute expected manual average of 3x3 neighborhood = (10+20+30+40+50+60+70+80+90)/9
	want := uint8((10+20+30+40+50+60+70+80+90)/9)
	if center != want { t.Fatalf("want %d got %d", want, center) }
}
