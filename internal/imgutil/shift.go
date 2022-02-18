package imgutil

import (
	"image"

	"github.com/ejuju/go-algo-art/internal/random"
)

const minShiftZones = 3  // min number of random shift zones for pixel shifting algorithm
const maxShiftZones = 20 // max number of random shift zones for pixel shifting algorithm

//
type PixelShiftZone struct {
	DirX bool
	DirY bool
	X    int
	Y    int
}

//
func ApplyPixelShift() TransformationFunc {
	return func(img image.Image) image.Image {
		newImg := image.NewRGBA(image.Rect(0, 0, Width(img), Height(img)))

		minX := Width(img) / 4
		maxX := 3 * Width(img) / 4
		minY := Height(img) / 4
		maxY := 3 * Height(img) / 4

		// get random zones within limits
		var zones [maxShiftZones]PixelShiftZone
		nbZones := random.IntMinMax(minShiftZones, maxShiftZones)
		maxZoneXAmpl := (maxX - minX) / 10 // max amplitude
		maxZoneYAmpl := (maxY - minY) / 10

		for i := 0; i < nbZones; i++ {
			zones[i] = PixelShiftZone{
				DirX: random.Bool(),
				DirY: random.Bool(),
				X:    random.IntMinMax(minX, maxX),
				Y:    random.IntMinMax(minY, maxY),
			}
		}

		_ = ForEachPixel(img, func(x, y int) error {
			xShift := 0
			yShift := 0

			ampl := 10

			for _, zone := range zones {
				zoneMinX := zone.X - maxZoneXAmpl
				zoneMinY := zone.Y - maxZoneYAmpl
				zoneMaxX := zone.X + maxZoneXAmpl
				zoneMaxY := zone.Y + maxZoneYAmpl
				if x > zoneMinX && x < zoneMaxX && y > zoneMinY && y < zoneMaxY {
					if zone.DirX == true {
						xShift = random.IntMinMax(0, ampl)
					} else if zone.DirX == false {
						xShift = random.IntMinMax(-ampl, 0)
					} else if zone.DirY == true {
						yShift = random.IntMinMax(0, ampl)
					} else if zone.DirY == false {
						yShift = random.IntMinMax(-ampl, 0)
					}
				}
			}

			newImg.Set(x, y, img.At(x+xShift, y+yShift))
			return nil
		})

		return newImg
	}

}
