package desktop

import "github.com/gabrielluizsf/antui/canvas"

// Palette holds the colour roles of the web stylesheet, so both versions of
// the app look like the same product.
type Palette struct {
	Primary, PrimaryHover, OnPrimary         canvas.Color
	PrimaryContainer, OnPrimaryContainer     canvas.Color
	SecondaryContainer, OnSecondaryContainer canvas.Color
	ErrorContainer, OnErrorContainer         canvas.Color
	Error, Warning                           canvas.Color
	Surface, SurfaceVariant, Background      canvas.Color
	OnSurface, OnSurfaceVariant, Outline     canvas.Color
	Scrim, InverseSurface, InverseOnSurface  canvas.Color
	InversePrimary                           canvas.Color
}

// LightPalette is the web version's :root palette.
func LightPalette() Palette {
	return Palette{
		Primary:              0xFF4F46E5,
		PrimaryHover:         0xFF4338CA,
		OnPrimary:            0xFFFFFFFF,
		PrimaryContainer:     0xFFE0E7FF,
		OnPrimaryContainer:   0xFF1E1B4B,
		SecondaryContainer:   0xFFEEEBFF,
		OnSecondaryContainer: 0xFF322E67,
		ErrorContainer:       0xFFFFE8E6,
		OnErrorContainer:     0xFF7A1E1B,
		Error:                0xFFE53935,
		Warning:              0xFFD97706,
		Surface:              0xFFFFFFFF,
		SurfaceVariant:       0xFFF5F6FB,
		Background:           0xFFE9EBF4,
		OnSurface:            0xFF1A1D29,
		OnSurfaceVariant:     0xFF5B606E,
		Outline:              0xFFC9CCD9,
		Scrim:                0x66101828,
		InverseSurface:       0xFF2B2E3B,
		InverseOnSurface:     0xFFF1F2F8,
		InversePrimary:       0xFFB4B0FF,
	}
}

// DarkPalette is the same roles for a system that paints itself dark.
func DarkPalette() Palette {
	return Palette{
		Primary:              0xFF6D66F0,
		PrimaryHover:         0xFF7F79F3,
		OnPrimary:            0xFFFFFFFF,
		PrimaryContainer:     0xFF2E2B66,
		OnPrimaryContainer:   0xFFE0E7FF,
		SecondaryContainer:   0xFF2C2A4A,
		OnSecondaryContainer: 0xFFDCD8FF,
		ErrorContainer:       0xFF4A1F1E,
		OnErrorContainer:     0xFFFFDAD6,
		Error:                0xFFEF5350,
		Warning:              0xFFF59E0B,
		Surface:              0xFF1C1E27,
		SurfaceVariant:       0xFF252835,
		Background:           0xFF12131A,
		OnSurface:            0xFFE6E7EE,
		OnSurfaceVariant:     0xFFA3A7B7,
		Outline:              0xFF3A3E4E,
		Scrim:                0x99000000,
		InverseSurface:       0xFFE6E7EE,
		InverseOnSurface:     0xFF1A1D29,
		InversePrimary:       0xFF4F46E5,
	}
}

// RingColor turns the ring amber and then red as the time runs out, like
// the web version.
func (p Palette) RingColor(fraction float64) canvas.Color {
	switch {
	case fraction <= 0.1:
		return p.Error
	case fraction <= 0.35:
		return p.Warning
	default:
		return p.Primary
	}
}
