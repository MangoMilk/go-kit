package generator

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSnowflake(t *testing.T) {
	machine := "1"
	newErr := SetupGenerator(SnowflakeGenerator, machine)
	if newErr != nil {
		t.Error(newErr)
	}

	concurrency := 5
	executeNum := 20
	var w sync.WaitGroup
	start := time.Now().UnixNano()
	for k := 0; k < concurrency; k++ {
		w.Add(1)
		go func() {
			for l := 0; l < executeNum; l++ {
				id, _ := Snowflake.GenID()
				fmt.Println(id)
			}
			w.Done()
		}()
	}

	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println(float64(end-start) / 1e6)
}

func TestSnoyflake(t *testing.T) {
	machine := "192.168.1.51"
	newErr := SetupGenerator(SnoyflakeGenerator, machine)
	if newErr != nil {
		t.Error(newErr)
	}

	concurrency := 5
	executeNum := 20
	var w sync.WaitGroup
	start := time.Now().UnixNano()
	for k := 0; k < concurrency; k++ {
		w.Add(1)
		go func() {
			for l := 0; l < executeNum; l++ {
				id, _ := Snoyflake.GenID()
				fmt.Println(id)
			}
			w.Done()
		}()
	}

	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println(float64(end-start) / 1e6)
}
