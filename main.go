package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/arpitbbhayani/arlt/arlt"

	"github.com/paulbellamy/ratecounter"
)

func randomAPI1(basicRlt *arlt.Arlt, servedCounter, discardCounter *ratecounter.RateCounter) int {
	exceeded, err := basicRlt.DidLimitExceed("randomAPI1", arlt.Configuration{
		MaxTicksPerWindow:       100,
		WindowdurationInSeconds: 10,
	})
	if err != nil {
		fmt.Println(err)
	}

	if exceeded {
		discardCounter.Incr(1)
		return 429
	}

	servedCounter.Incr(1)
	return 200
}

func main() {
	t := 0
	wg := &sync.WaitGroup{}

	basicRlt, err := arlt.NewArlt(arlt.DefaultSetting())
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	totalCounter := ratecounter.NewRateCounter(1 * time.Second)
	servedCounter := ratecounter.NewRateCounter(1 * time.Second)
	discardCounter := ratecounter.NewRateCounter(1 * time.Second)
	go func() {
		for {
			fmt.Println("t = ", t, "total = ", totalCounter.Rate(), "served = ", servedCounter.Rate(), "discard = ", discardCounter.Rate())
			time.Sleep(1 * time.Second)
			t++
		}
	}()

	wg.Add(1)
	go func() {
		for {
			totalCounter.Incr(1)
			randomAPI1(basicRlt, servedCounter, discardCounter)
			time.Sleep(time.Duration(rand.Uint32()%100) * time.Millisecond)
		}
	}()

	wg.Wait()
}
