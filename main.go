package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"

	"github.com/ghodss/yaml"
	"k8s.io/helm/pkg/repo"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func noop(obj interface{}) {

}

func sleep(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}

func main() {
	Unmarshal()
}

func Unmarshal() {
	data, err := ioutil.ReadFile("/helm-index.yaml")
	check(err)
	noop(data)
	PrintMemUsage()

	sleep(2)

	index := &repo.IndexFile{}
	fmt.Printf("unmarshal\n")
	err = yaml.Unmarshal(data, index)
	check(err)
	PrintMemUsage()

	sleep(10)

	versions := 0
	for _, entry := range index.Entries {
		versions += len(entry)
	}
	fmt.Printf("total entries: %d, total versions: %d\n", len(index.Entries), versions)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tHeapAlloc = %v MiB", bToMb(m.HeapAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
