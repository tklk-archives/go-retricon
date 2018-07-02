package retricon

import (
	"crypto/sha512"
	"encoding/binary"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

func New(seed []byte) image.Image {
	buf := sha512.Sum512(seed)
	seed_int := binary.BigEndian.Uint64(buf[56:])
	rand.Seed(int64(seed_int))

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)

	c := color.RGBA{uint8(r), uint8(g), uint8(b), 255}

	for j := 0; j < 5; j++ {
		for k := 0; k < 3; k++ {
			if rand.Intn(2) != 0 {
				draw.Draw(img, image.Rect(k*20, j*20, k*20+20, j*20+20), &image.Uniform{c}, image.ZP, draw.Over)
				if k == 1 {
					draw.Draw(img, image.Rect(k*20+40, j*20, k*20+20+40, j*20+20), &image.Uniform{c}, image.ZP, draw.Over)
				} else {
					draw.Draw(img, image.Rect(k*20+80, j*20, k*20+20+80, j*20+20), &image.Uniform{c}, image.ZP, draw.Over)
				}
			}
		}
	}
	return img
}