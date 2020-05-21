package shapes

import (
	"fmt"
	"strings"

	geo "github.com/pnctech/gobadge/svg/geometry"
)

type Text struct {
	ID     string
	Value  string
	Font   Font
	Origin geo.Coordinate
}

type Font struct {
	Family string
	Size   int
	Weight string
}

func (text Text) Vectorize() string {

	elements := strings.Join([]string{
		fmt.Sprintf(`id="%s"`, text.ID),
		fmt.Sprintf(`font-family="%s"`, text.Font.Family),
		fmt.Sprintf(`font-size="%d"`, text.Font.Size),
		fmt.Sprintf(`font-weight="%s"`, text.Font.Weight),
	}, " ")

	contents := fmt.Sprintf(`<tspan %s>%s</tspan>`, text.Origin.Vectorize(), text.Value)
	template := `<text %s>%s</text>`

	return fmt.Sprintf(template, elements, contents)
}
