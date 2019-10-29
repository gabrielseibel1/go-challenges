//Calculates average housing prices for publishers in regions, from a JSON array of houses

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	House struct {
		Price      float64 `json:"price"`
		PropertyID int     `json:"property_id"`
		Info
	}
	Info struct {
		Region    string `json:"region"`
		Publisher string `json:"publisher"`
	}
)

func (i Info) String() string {
	return fmt.Sprintf("%s %s", i.Region, i.Publisher)
}

func (h House) String() string {
	return fmt.Sprintf("%f %d", h.Price, h.PropertyID) + h.Info.String()
}

func main() {
	jsonPayload := `
	[
		{
		   "price": 514885,
		   "property_id": 0,
		   "publisher": "L0ft",
		   "region": "Porto Alegre"
		},
		{
		   "price": 514887,
		   "property_id": 2,
		   "publisher": "L0ft",
		   "region": "Itaquera"
		},
		{
		   "price": 514885,
		   "property_id": 3,
		   "publisher": "L0ft",
		   "region": "Itaquera"
		},
		{
		   "price": 436372,
		   "property_id": 4,
		   "publisher": "YellowPages",
		   "region": "Itaquera"
		},
		{
		   "price": 378819,
		   "property_id": 5,
		   "publisher": "iHouse",
		   "region": "Itaquera"
		},
		{
		   "price": 446508,
		   "property_id": 6,
		   "publisher": "NextRoof",
		   "region": "Itaquera"
		}
	]`
	avgs := calculateAverages(jsonPayload)
	for _, s := range avgs {
		fmt.Println(s)
	}
}

func calculateAverages(payload string) []string {
	//read
	var houses []House
	if err := json.NewDecoder(strings.NewReader(payload)).Decode(&houses); err != nil {
		panic(err)
	}
	//collect prices for each region-publisher
	infoPrices := make(map[Info][]float64)
	for _, h := range houses {
		infoPrices[h.Info] = append(infoPrices[h.Info], h.Price)
	}
	//calculate mean prices
	var s []string
	for info := range infoPrices {
		sum := 0.0
		for _, p := range infoPrices[info] {
			sum += p
		}
		mean := sum / float64(len(infoPrices[info]))
		s = append(s, info.String()+fmt.Sprintf(" %d", int(mean)))
	}
	return s
}
