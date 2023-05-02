package load

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
)

func makeRequests(wg *sync.WaitGroup, url string, nRequests uint32) {
	for i := 0; i < int(nRequests); i++ {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("\nReceived Status: %v", resp.StatusCode)
	}
	defer wg.Done()
}

func requestPerConnection(connections uint32, requests uint32) []uint32 {
	var rpc []uint32
	reqs_per_channel := math.Floor(float64(requests / connections))
	for x := 0; x < int(connections-1); x++ {
		rpc = append(rpc, uint32(reqs_per_channel))
	}

	remaing_request := requests - (uint32(reqs_per_channel) * (connections - 1))
	rpc = append(rpc, remaing_request)

	return rpc
}

func Hammer(url string, connections uint32, requests uint32) {
	var wg sync.WaitGroup

	rpc := requestPerConnection(connections, requests)
	fmt.Println(rpc)

	fmt.Printf("Running Load on: %v", url)
	for x := 0; x < int(connections); x++ {
		wg.Add(1)
		go makeRequests(&wg, url, rpc[x])
	}

	wg.Wait()
	fmt.Printf("\nAll finished")
}
