package csvi

import (
	"container/list"
	"strings"
	"testing"

	"github.com/hymkor/csvi/uncsv"
)

func TestFitAllColumns(t *testing.T) {
	app := &Application{csvLines: list.New()}
	mode := &uncsv.Mode{Comma: ','}
	rows, err := uncsv.ReadAll(strings.NewReader(
		"a,bb,c\naaaaa,b\nx,y,"+strings.Repeat("z", 60)+",w\n"), mode)
	if err != nil {
		t.Fatal(err)
	}
	for i := range rows {
		app.push(&rows[i])
	}
	cw := NewCellWidth()
	cw.Option[1] = 7 // explicit width must survive autofit

	app.fitAllColumns(cw)

	for col, want := range []int{6, 7, 40, 3} {
		if got := cw.Get(col); got != want {
			t.Errorf("column %d: expected width %d, got %d", col, want, got)
		}
	}
}
