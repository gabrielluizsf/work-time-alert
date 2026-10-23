package desktop

import (
	"math"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func near(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestArcPointsStartAtTwelveAndGoClockwise(t *testing.T) {
	pts := arcPoints(100, 100, 50, 0.25)
	assert.True(t, near(pts[0].X, 100) && near(pts[0].Y, 50))
	last := pts[len(pts)-1]
	assert.True(t, near(last.X, 150) && near(last.Y, 100))
}

func TestArcPointsClampTheFraction(t *testing.T) {
	full := arcPoints(0, 0, 10, 2)
	first, last := full[0], full[len(full)-1]
	assert.True(t, near(first.X, last.X) && near(first.Y, last.Y))
	assert.Equal(t, len(arcPoints(0, 0, 10, -1)), 3)
}
