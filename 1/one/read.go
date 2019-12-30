package main

import (
	"bufio"
	"os"
	"strconv"
)

// Read reads all the lines from an input file.
func Read(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b = append(b, scanner.Text())
	}

	return byteToFloatArray(b), nil
}

func byteToFloatArray(b []string) (f []float64) {
	for _, v := range b {
		float, _ := strconv.ParseFloat(v, 64)
		f = append(f, float64(float))
	}
	return f
}
