package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	// "github.com/jinzhu/gorm"
	"github.com/khelben/fnimport/dto"
	"github.com/khelben/fnimport/importers"
)

func main() {
	_, err := importers.ReadFietsNetRtrees()
	if err != nil {
		panic(err)
	}
}

func loadTracks(trackname string, m *sync.WaitGroup) {
	fmt.Printf(">>>> %s\n", trackname)
	resp, _ := http.Get(fmt.Sprintf("http://api.fietsnet.be/v2/m1034/tracks/%s", trackname))
	defer resp.Body.Close()
	responseStream, _ := ioutil.ReadAll(resp.Body)

	var response dto.TrackResponse
	json.Unmarshal(responseStream, &response)
	fmt.Printf("<<<< %s\n", trackname)
}
