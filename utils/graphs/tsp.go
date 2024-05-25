package graphs

import (
	"context"
	"golang.org/x/sync/semaphore"
)

const (
	// about half of max safe float
	inf = float64(1 << 52)

	// limit to 32 MB-ish
	maxSolverMemoryUsage = int64(32 * 1024 * 1024)
)

var (
	// semaphore where weight = approx. mem available/used in bytes
	memLock = semaphore.NewWeighted(maxSolverMemoryUsage)
)

func GetTspDistance(
	ctx context.Context,
	startIdx int,
	endIdx int,
	distMatrix [][]float64,
) (res float64, err error) {
	if err = ctx.Err(); err != nil {
		return
	}
	var numVertices = len(distMatrix)

	// reserve (approximately) the size of the tsp dp table
	// (8 bytes per float64) * numVertices * 2^numVertices
	expectedMemoryUsage := 8 * int64(numVertices) * (1 << numVertices)

	// memLock usage lock via semaphore
	err = memLock.Acquire(ctx, expectedMemoryUsage)
	if err != nil {
		return
	}
	defer memLock.Release(expectedMemoryUsage)

	res = tspWithConstraint(numVertices, startIdx, endIdx, distMatrix).solve()
	return
}

type tspConstraint struct {
	numVertices int

	startIdx   int
	endIdx     int
	distMatrix [][]float64

	dpTable []float64
}

func tspWithConstraint(
	numVertices int,
	startIdx int,
	endIdx int,
	distMatrix [][]float64,
) *tspConstraint {
	ret := tspConstraint{
		numVertices: numVertices,
		startIdx:    startIdx,
		endIdx:      endIdx,
		distMatrix:  distMatrix,
	}

	// initialize dp table to required size and fill with -1 (uncomputed)
	tableSize := numVertices * (1 << numVertices)
	ret.dpTable = make([]float64, tableSize)

	// use copy doubling for faster initialization
	ret.dpTable[0] = -1
	for i := 1; i < tableSize; i *= 2 {
		copy(ret.dpTable[i:], ret.dpTable[:i])
	}

	// base terminating case (at endIdx and all vertices visited)
	ret.setDpTableValue(endIdx, 0, 0)
	return &ret
}

func (c *tspConstraint) solve() float64 {
	// all 1 bits (unvisited), equal to 0b111...
	startingMask := (1 << c.numVertices) - 1

	// when we don't need to go back, we need to mark startIdx as visited
	if c.startIdx != c.endIdx {
		// toggle startIdx-th bit to 0 (visited)
		startingMask ^= 1 << c.startIdx
	}

	return c.solveImpl(c.startIdx, startingMask)
}

// typical TSP solver with dynamic programming + bit-masking for state
func (c *tspConstraint) solveImpl(pos, mask int) float64 {
	// already computed
	memoValue := c.getDpTableValue(pos, mask)
	if memoValue >= 0 {
		return memoValue
	}

	minCost := inf
	// try each path through an unvisited vertex
	for i := 0; i < c.numVertices; i++ {
		// i-th vertex not yet visited as i-th bit of mask is still 1
		if (mask>>i)&1 != 0 {
			// compute the cost for taking this path
			takeCost := c.distMatrix[pos][i] + c.solveImpl(i, mask^(1<<i))

			// replace minCost if taking this path means we get a lower cost
			minCost = min(minCost, takeCost)
		}
	}

	// memoize results before returning
	c.setDpTableValue(pos, mask, minCost)
	return minCost
}

func (c *tspConstraint) toDpTableIndex(pos, mask int) int {
	// mask-major indexing on flattened 2D array for speed
	// (large index first dimension ordering heuristic)
	return mask*c.numVertices + pos
}

func (c *tspConstraint) setDpTableValue(pos, mask int, value float64) {
	c.dpTable[c.toDpTableIndex(pos, mask)] = value
}

func (c *tspConstraint) getDpTableValue(pos, mask int) float64 {
	return c.dpTable[c.toDpTableIndex(pos, mask)]
}
