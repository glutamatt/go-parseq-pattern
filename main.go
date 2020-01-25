package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Input int
type Output string

const concurrency = 4

func main() {

	inputs := make([]Input, 50)
	for i := range inputs {
		inputs[i] = Input(i)
	}

	process := func(in Input, out chan<- Output) {
		time.Sleep(time.Duration(rand.Float32() * float32(time.Second)))
		out <- Output(fmt.Sprintf("result of %d", in))
	}

	results := make(chan chan Output, concurrency)

	go func() {
		for _, input := range inputs {
			output := make(chan Output)
			results <- output
			go process(input, output)
		}
		close(results)
	}()

	for res := range results {
		println(<-res)
	}
}
