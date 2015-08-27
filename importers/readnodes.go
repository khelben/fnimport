package importers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dbfn "github.com/khelben/fnimport/db"
	"github.com/khelben/fnimport/dto"

	"github.com/jinzhu/gorm"
)

func ReadFietsNetPois() (response *dto.PoisResponse, err error) {
	resp, err := http.Get("http://api.fietsnet.be/v2/m1034/pois/0")
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
	return
}

func ImportFietsNetNodes(PoiList []dto.Poi) (err error) {
	db, err := gorm.Open("postgres", "user=cvp dbname=gis sslmode=disable")
	db.LogMode(true)
	defer db.Close()
	db.Exec("create extension postgis;")

	db.DropTable(&dbfn.FietsNetNode{})
	db.CreateTable(&dbfn.FietsNetNode{})
	for _, poi := range PoiList {
		node := dbfn.FietsNetNode{
			Name:  poi.Label,
			Point: poi.Coordinate.ToGeoPoint(),
		}
		db.Create(&node)
	}
	return
}
