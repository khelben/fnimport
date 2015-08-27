package importers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/khelben/fnimport/dto"
)

func ReadFietsNetRtrees() (response *dto.RtreesResponse, err error) {
	resp, err := http.Get("http://api.fietsnet.be/v2/m1034/rtrees/0")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseStream, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(responseStream, &response)
	if err != nil {
		return
	}
	// fmt.Printf("json: %#v", response)
	return
}
