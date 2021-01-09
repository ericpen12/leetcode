package _95_max_area_of_island

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_maxAreaOfIsland1(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"1", [][]int{{1, 1}}, 2},
		{"2", [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}, 4},
	}
	for _, tt := range tests {
		Convey(tt.name, t, func() {
			So(maxAreaOfIsland(tt.grid), ShouldEqual, tt.want)
		})
	}
}
