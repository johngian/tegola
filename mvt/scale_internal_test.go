package mvt

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-spatial/tegola"
	"github.com/go-spatial/tegola/basic"
)

func TestScaleLinestring(t *testing.T) {
	tile := tegola.NewTile(20, 0, 0)

	newLine := func(ptpairs ...float64) (ln basic.Line) {
		for i, j := 0, 1; j < len(ptpairs); i, j = i+2, j+2 {
			pt, err := tile.FromPixel(tegola.WebMercator, [2]float64{ptpairs[i], ptpairs[j]})
			if err != nil {
				panic(fmt.Sprintf("error trying to convert %v,%v to WebMercator. %v", ptpairs[i], ptpairs[j], err))
			}

			ln = append(ln, basic.Point(pt))
		}

		return ln
	}

	type tcase struct {
		g tegola.LineString
		e basic.Line
	}

	fn := func(tc tcase) func(t *testing.T) {
		return func(t *testing.T) {
			got := scalelinestr(tc.g, tile)

			if !reflect.DeepEqual(tc.e, got) {
				t.Errorf("expected %v got %v", tc.e, got)
			}
		}
	}

	tests := map[string]tcase{
		"duplicate pt simple line": {
			g: basic.NewLine(9.0, 9.0, 9.0, 9.0),
		},
		"simple line": {
			g: newLine(10.0, 10.0, 11.0, 11.0),
			e: basic.NewLine(9.0, 9.0, 11.0, 11.0),
		},
		"simple line 3pt": {
			g: newLine(10.0, 10.0, 11.0, 10.0, 11.0, 15.0),
			e: basic.NewLine(9.0, 9.0, 11.0, 9.0, 11.0, 14.0),
		},
	}

	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}
