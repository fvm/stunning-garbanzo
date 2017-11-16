package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	base := "./bench"

	s := make([]Setting, 6)
	for i := range s {
		s[i] = Setting{
			path:   fmt.Sprintf("%s/%d", base, i),
			height: rand.Intn(2000),
			width:  rand.Intn(2000),
		}
	}
	image, err := ioutil.ReadFile("OV-fiets2009.jpg")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("BenchmarkConvertParallel", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ConvertParallel(image, s)
		}

	})
	b.Run("BenchmarkConvertSequential", func(b *testing.B) {

		os.MkdirAll(base, 0777)

		defer os.RemoveAll(base)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ConvertSequential(image, s)
		}

	})

}

func TestConvertParallel(t *testing.T) {
	base := "./test"
	os.MkdirAll(base, 0777)
	defer os.RemoveAll("./test")

	s := []Setting{
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

	if err := ConvertParallel(b, s); err != nil {
		t.Error(err)
	}
}

func TestConvertSequential(t *testing.T) {
	base := "./test"
	os.MkdirAll(base, 0777)
	defer os.RemoveAll("./test")

	s := []Setting{
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

	if err := ConvertParallel(b, s); err != nil {
		t.Error(err)
	}
}
