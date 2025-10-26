package cgoimage

// Blur applies a naive box blur to a grayscale pixel slice of length w*h.
func Blur(pixels []uint8, w, h int) {
	if w*h == 0 || len(pixels) < w*h { return }
	out := make([]uint8, w*h)
	for y:=0; y<h; y++ {
		for x:=0; x<w; x++ {
			var sum, count int
			for dy:=-1; dy<=1; dy++ {
				ny := y+dy; if ny<0 || ny>=h { continue }
				for dx:=-1; dx<=1; dx++ {
					nx := x+dx; if nx<0 || nx>=w { continue }
					idx := ny*w + nx
					sum += int(pixels[idx]); count++
				}
			}
			out[y*w+x] = uint8(sum / count)
		}
	}
	copy(pixels, out)
}
