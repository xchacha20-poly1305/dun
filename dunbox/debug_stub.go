//go:build !(linux || darwin)

package dunbox

func rusageMaxRSS() float64 {
	return -1
}
