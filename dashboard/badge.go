package dashboard

import (
	"github.com/pnctech/gobadge/svg"
	geo "github.com/pnctech/gobadge/svg/geometry"
)

type Badge struct {
	Label string `json:"label"`
	Tag   string `json:"tag"`
	Color string `json:"color"`
}

func (badge Badge) Draw(canvas *svg.Canvas) {
	canvas.Open(float64(badgeW), float64(badgeH))
	origin := geo.Coordinate{X: 0, Y: 0}
	badge.draw(canvas, origin)
	canvas.Close()
}
