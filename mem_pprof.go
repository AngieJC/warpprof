package warpprof

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func pprofMemory(frequent int) {
	for i := 0; ; i++ {
		fm, err := os.OpenFile(fmt.Sprintf("pid%v_mem_%v", os.Getpid(), i), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		pprof.WriteHeapProfile(fm)
		time.Sleep(time.Second * time.Duration(frequent))
	}
}
