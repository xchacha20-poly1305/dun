package dunmain

import (
	"github.com/xchacha20-poly1305/dun/dunbox"
)

func DisableColor() {
	disableColor = true
	dunbox.DisableColor()
}
