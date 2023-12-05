package vt10x

// ANSI color values
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGrey
	DarkGrey
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

// Default colors are potentially distinct to allow for special behavior.
// For example, a transparent background. Otherwise, the simple case is to
// map default colors to another color.
const (
	DefaultFG Color = 1<<24 + iota
	DefaultBG
	DefaultCursor
)

var colorCache = [256]Color{}

// Color maps to the ANSI colors [0, 16) and the xterm colors [16, 256).
type Color uint32

// ANSI returns true if Color is within [0, 16).
func (c Color) ANSI() bool {
	return (c < 16)
}

// byte2color maps an 8 bit color to a 24-bit RGB Color.
func byte2color(i int) Color {
	if colorCache[0] == 0 {
		loadColorCache()
	}
	return colorCache[i]
}

func loadColorCache() {

	predefinedColors := []Color{
		0x2e3436, 0xcc0000, 0x4e9a06, 0xc4a000, 0x3465a4, 0x75507b, 0x06989a,
		0xd3d7cf, 0x555753, 0xef2929, 0x8ae234, 0xfce94f, 0x729fcf, 0xad7fa8,
		0x34e2e2, 0xeeeeec}

	for i, color := range predefinedColors {
		colorCache[i] = color
	}

	// Generate colors (16-231)
	v := []uint32{0x00, 0x5f, 0x87, 0xaf, 0xd7, 0xff}
	for i := 0; i < 216; i++ {
		r := v[(i/36)%6]
		g := v[(i/6)%6]
		b := v[i%6]
		rgb := (r << 16) | (g << 8) | b
		colorCache[16+i] = Color(rgb)
	}

	// Generate greys (232-255)
	for i := 0; i < 24; i++ {
		c := uint32(8 + i*10)
		rgb := (c << 16) | (c << 8) | c
		colorCache[232+i] = Color(rgb)
	}
}
