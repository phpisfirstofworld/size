package size

import (
	"fmt"
	"strconv"
)

type Size struct {
	byte int64
}

func NewSize(byte int64) *Size {

	return &Size{byte: byte}
}

func (s *Size) Bytes() int64 {

	return s.byte
}

func (s *Size) Kb() float64 {

	value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", float64(s.byte)/1024.0), 64)

	return value
}

func (s Size) Mb() float64 {

	value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", float64(s.byte)/1024.0/1024.0), 64)

	return value
}

func (s Size) Gb() float64 {

	value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", float64(s.byte)/1024.0/1024.0/1024.0), 64)

	return value
}
