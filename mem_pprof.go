package warpprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func PprofMemory(frequent int) {
	for i := 0; ; i++ {
		fm, err := os.OpenFile(fmt.Sprintf("pid%v_mem_%v", os.Getpid(), i), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		pprof.WriteHeapProfile(fm)
		time.Sleep(time.Second * time.Duration(frequent))
	}
}

func ListenAndServe(addr string, handler http.Handler) {
	http.ListenAndServe(addr, handler)
}
