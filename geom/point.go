package geom

import (
	"github.com/vdobler/plot"
	"image/color"
)

type Point struct {
	X, Y float64
	Size float64
	Color color.Color
	Fill color.Color
	Shape int
}

func (p Point) Render(data plot.DataFrame, aes plot.AesMapping, p plot.Plot) {
	am := p.Aes.Merge(aes, plot.DefaultTheme.PointAes)
	for i := 0; i<plot.Length(data); i++ {
		x := plot.Field(data,aes.X)
		y := plot.Field(data,aes.Y)
		size := plot.Field(data,aes.Size)
	}
}
