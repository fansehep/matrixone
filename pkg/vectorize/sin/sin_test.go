package sin

import (
	"math"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSinUint8(t *testing.T) {
	lvs := []uint8{1, 2, 10, 12, 21, 212, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinUint8(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}
func TestSinUint16(t *testing.T) {
	lvs := []uint16{1, 2, 10, 12, 21, 212, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinUint16(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}

func TestSinUint32(t *testing.T) {
	lvs := []uint32{1, 2, 10, 12, 21, 212, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinUint32(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}
func TestSinUint64(t *testing.T) {
	lvs := []uint64{1, 2, 10, 12, 21, 212, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinUint64(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}
func TestSinInt8(t *testing.T) {
	lvs := []int8{1, 2, 10, 12, 21, 22, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinInt8(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}

func TestSinInt16(t *testing.T) {
	lvs := []int16{1, 32, 10, 12, 21, 22, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinInt16(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}
func TestSinInt32(t *testing.T) {
	lvs := []int32{1, 32, 10, 12, 21, 22, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinInt32(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}

func TestSinInt64(t *testing.T) {
	lvs := []int64{1, 32, 10, 12, 21, 22, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinInt64(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}
func TestSinFloat32(t *testing.T) {
	lvs := []float32{1, 32, 1.0, 12.13, 21, 22, 31, 0.01}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(float64(lvs[i])))
	}
	rfs = SinFloat32(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}
}
func TestSinFloat64(t *testing.T) {
	lvs := []float64{1, 32.23, 10, 12, 21, 22, 23, 31}
	rfs := []float64{0, 2, 3, 4, 5, 6, 7, 8};
	var result []float64
	for i := range(lvs) {
		result = append(result, math.Sin(lvs[i]))
	}
	rfs = SinFloat64(lvs, rfs)
	for i := range(result) {
		require.Equal(t, rfs[i], result[i])
	}	
}
