//go:build debug && !linux

package dunmain

func rusageMaxRSS() float64 {
	return -1
}
