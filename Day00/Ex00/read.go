package main

import (
	"bufio"
	"errors"
	"log"
	"strconv"
)

var (
	EmptyInput       = errors.New("empty input")
	WrongInput       = errors.New("wrong input")
	NumberIsNotRange = errors.New("number is not range")
	NoDataProvided   = errors.New("no data provided")
)

func ReadInput(scanner *bufio.Scanner) ([]float64, error) {

	var in []float64

	for scanner.Scan() {

		text := scanner.Text()
		if text == "" {
			log.Println(EmptyInput)
			continue
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			log.Println(WrongInput)
		} else if num < -100000 || num > 100000 {
			log.Println(NumberIsNotRange)
		} else {
			in = append(in, float64(num))
		}
	}

	if len(in) == 0 {
		return nil, NoDataProvided
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}
