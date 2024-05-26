package graphs_test

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/utils/graphs"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

/*
             m
             a
             a                           h
             s       h   s               u
             t   a   e   i   g           l
             r   a   e   t   e           s
             i   c   r   t   l   e   b   b   a       e
             c   h   l   a   e   c   o   e   n   o   p
             h   e   e   r   e   h   n   r   n   h   e
             t   n   n   d   n   t   n   g   e   e   n
-------------------------------------------------------
maastricht | 0   29  20  21  16  31  100 12  4   31  18
    aachen | 29  0   15  29  28  40  72  21  29  41  12
   heerlen | 20  15  0   15  14  25  81  9   23  27  13
   sittard | 21  29  15  0   4   12  92  12  25  13  25
    geleen | 16  28  14  4   0   16  94  9   20  16  22
      echt | 31  40  25  12  16  0   95  24  36  3   37
      bonn | 100 72  81  92  94  95  0   90  101 99  84
  hulsberg | 12  21  9   12  9   24  90  0   15  25  13
     kanne | 4   29  23  25  20  36  101 15  0   35  18
       ohe | 31  41  27  13  16  3   99  25  35  0   38
      epen | 18  12  13  25  22  37  84  13  18  38  0

Optimal (by program): cities 0-7-4-3-9-5-2-6-1-10-8-0 = 253km
maastricht -> hulsberg -> geleen -> sittard -> ohe -> kanne -> echt
-> heerlen -> bonn -> aachen -> epen -> kanne -> maastricht

source: https://stackoverflow.com/questions/11007355/data-for-simple-tsp
*/

var distMatrix = [][]float64{
	{0, 29, 20, 21, 16, 31, 100, 12, 4, 31, 18},
	{29, 0, 15, 29, 28, 40, 72, 21, 29, 41, 12},
	{20, 15, 0, 15, 14, 25, 81, 9, 23, 27, 13},
	{21, 29, 15, 0, 4, 12, 92, 12, 25, 13, 25},
	{16, 28, 14, 4, 0, 16, 94, 9, 20, 16, 22},
	{31, 40, 25, 12, 16, 0, 95, 24, 36, 3, 37},
	{100, 72, 81, 92, 94, 95, 0, 90, 101, 99, 84},
	{12, 21, 9, 12, 9, 24, 90, 0, 15, 25, 13},
	{4, 29, 23, 25, 20, 36, 101, 15, 0, 35, 18},
	{31, 41, 27, 13, 16, 3, 99, 25, 35, 0, 38},
	{18, 12, 13, 25, 22, 37, 84, 13, 18, 38, 0},
}

func TestGetTspDistance(t *testing.T) {
	Convey("When getting the TSP of a graph", t, func() {
		Convey("Should return the optimal distance", func() {
			res, err := graphs.GetTspDistance(
				context.TODO(),
				0,
				0,
				distMatrix,
			)
			So(err, ShouldBeNil)
			So(helper.IsEqualFloat(res, 253.0), ShouldBeTrue)
		})
	})
	Convey("When getting the TSP of just 3 points in a line", t, func() {
		Convey("Should return the optimal distance", func() {
			small := [][]float64{
				{0, 2, 100},
				{100, 0, 3},
				{100, 100, 0},
			}
			res, err := graphs.GetTspDistance(
				context.TODO(),
				0,
				2,
				small,
			)
			So(err, ShouldBeNil)
			So(helper.IsEqualFloat(res, 5.0), ShouldBeTrue)
		})
	})
}

func benchmarkGetTspDistance(b *testing.B, size int) {
	subMatrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		subMatrix[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			subMatrix[i][j] = distMatrix[i][j]
		}
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < b.N; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			defer wg.Done()

			_, err := graphs.GetTspDistance(
				context.TODO(),
				0,
				0,
				subMatrix,
			)
			if err != nil {
				panic(err)
			}
		}(&wg)
	}
}

// test machine:
// - Ryzen 9 5900 HS
// - linux 6.6.31
//
// versions of stuff:
// - go 1.22.2
// - gcc 14.1.1
// - glibc 2.39

// ~150k calls per second for 11 vertices
func BenchmarkGetTspDistance11(b *testing.B) {
	benchmarkGetTspDistance(b, 11)
}

// ~700k calls per second for 9 vertices (our supposed max), fast enough lol
func BenchmarkGetTspDistance9(b *testing.B) {
	benchmarkGetTspDistance(b, 9)
}

// this should take basically no time at all (~5 million calls per second)
func BenchmarkGetTspDistance7(b *testing.B) {
	benchmarkGetTspDistance(b, 7)
}
