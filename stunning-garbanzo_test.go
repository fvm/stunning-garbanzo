package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

// TODO: Define standard output settings
// TODO: Benchmark an entire series of exports on a varied range of inputs

func BenchmarkBicycle(b *testing.B) {
	base := "./bench"

	s := make([]setting, 6)
	for i := range s {
		s[i] = setting{
			path:   fmt.Sprintf("%s/%d", base, i),
			height: rand.Intn(2000),
			width:  rand.Intn(2000),
		}
	}
	image, err := ioutil.ReadFile("bike.jpg")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("BenchmarkConvertParallel", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertParallel(image, s)
		}

	})
	b.Run("BenchmarkConvertSequential", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertSequential(image, s)
		}

	})

}
func BenchmarkChess(b *testing.B) {
	base := "./bench"

	s := make([]setting, 6)
	for i := range s {
		s[i] = setting{
			path:   fmt.Sprintf("%s/%d", base, i),
			height: rand.Intn(2000),
			width:  rand.Intn(2000),
		}
	}
	image, err := ioutil.ReadFile("chess.jpg")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("BenchmarkConvertParallel", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertParallel(image, s)
		}

	})
	b.Run("BenchmarkConvertSequential", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertSequential(image, s)
		}

	})

}
func BenchmarkNoise(b *testing.B) {
	base := "./bench"

	s := make([]setting, 6)
	for i := range s {
		s[i] = setting{
			path:   fmt.Sprintf("%s/%d", base, i),
			height: rand.Intn(2000),
			width:  rand.Intn(2000),
		}
	}
	image, err := ioutil.ReadFile("noise.jpg")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("BenchmarkConvertParallel", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertParallel(image, s)
		}

	})
	b.Run("BenchmarkConvertSequential", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertSequential(image, s)
		}

	})

}
func BenchmarkTiff(b *testing.B) {
	base := "./bench"

	s := make([]setting, 6)
	for i := range s {
		s[i] = setting{
			path:   fmt.Sprintf("%s/%d", base, i),
			height: rand.Intn(2000),
			width:  rand.Intn(2000),
		}
	}
	image, err := ioutil.ReadFile("tiff.tif")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("BenchmarkConvertParallel", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertParallel(image, s)
		}

	})
	b.Run("BenchmarkConvertSequential", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			convertSequential(image, s)
		}

	})

}

func TestConvertParallel(t *testing.T) {
	base := "./test"
	os.MkdirAll(base, 0777)
	defer os.RemoveAll("./test")

	s := []setting{
		{
			path:   fmt.Sprintf("%s/%d", base, 1),
			height: 100,
			width:  100,
		},
		{
			path:   fmt.Sprintf("%s/%d", base, 2),
			height: 200,
			width:  200,
		},
	}

	b, err := ioutil.ReadFile("noise.jpg")
	if err != nil {
		t.Fatal(err)
	}

	if err := convertParallel(b, s); err != nil {
		t.Error(err)
	}
}

func TestConvertSequential(t *testing.T) {
	base := "./test"
	os.MkdirAll(base, 0777)
	defer os.RemoveAll("./test")

	s := []setting{
		{
			path:   fmt.Sprintf("%s/%d", base, 1),
			height: 100,
			width:  100,
		},
		{
			path:   fmt.Sprintf("%s/%d", base, 2),
			height: 200,
			width:  200,
		},
	}
	b, err := ioutil.ReadFile("noise.jpg")
	if err != nil {
		t.Fatal(err)
	}

	if err := convertParallel(b, s); err != nil {
		t.Error(err)
	}
}
