//go:build arm64
// +build arm64

package dsp

import "github.com/zaf/g711"

func L16MixSat160(dst, src *int16) {

}

func LinearToUlaw(linear int64) (ulaw int64) {
	return int64(g711.EncodeUlawFrame(int16(linear)))
}

// Turns a μ-Law byte back into an audio sample.
func UlawToLinear(ulaw int64) (linear int64) {
	linear = int64(g711.DecodeUlawFrame(uint8(ulaw)))
	return
}
