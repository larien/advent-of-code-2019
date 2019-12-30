package main

import (
	"io/ioutil"
	"os"
)

func Read(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return byteToFloatArray(b), nil
}

func byteToFloatArray(b []byte) []float64 {
	f := make([]float64, len(b))
	for _, v := range b {
		f = append(f, float64(v))
	}
	return f
}
