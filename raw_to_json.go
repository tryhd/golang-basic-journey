package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Vehicle struct {
	Name string `json:"name"`
	Data []struct {
		Name string `json:"name"`
		Data []struct {
			Name string `json:"name"`
		} `json:"data"`
	} `json:"data"`
}

func main() {
	raw := []string{"mobil/toyota/land-cruiser", "motor/suzuki/satria-r", "motor/honda/vario-125", "mobil/toyota/kijang-innova", "mobil/daihatsu/xenia", "motor/yamaha/mio", "mobil/daihatsu/sigra", "mobil/nissan/juke", "motor/honda/beat", "motor/yamaha/r15", "mobil/daihatsu/ayla", "motor/honda/cbr-150r", "motor/suzuki/spin", "mobil/nissan/march", "mobil/nissan/livina", "motor/yamaha/r25", "mobil/toyota/avanza", "motor/suzuki/satria-f150"}
	var d = make(map[string]bool)
	var vehicle = []Vehicle{}
	for _, v := range raw {
		tags := strings.Split(v, "/")
		if len(tags) == 3 {
			if _, ok := d[tags[0]]; !ok {
				d[tags[0]] = true
				vehicle = append(vehicle, Vehicle{Name: tags[0]})
			}
		}
	}
	sort.Strings(raw)
	for _, s := range raw {
		tags := strings.Split(s, "/")
		if len(tags) == 3 {
			if tags[0] == "mobil" {
				if _, ok := d[tags[1]]; !ok {
					d[tags[1]] = true
					vehicle[len(vehicle)-2].Data = append(vehicle[len(vehicle)-2].Data, struct {
						Name string `json:"name"`
						Data []struct {
							Name string `json:"name"`
						} `json:"data"`
					}{Name: tags[1]})
				}
				vehicle[len(vehicle)-2].Data[len(vehicle[len(vehicle)-2].Data)-1].Data = append(vehicle[len(vehicle)-2].Data[len(vehicle[len(vehicle)-2].Data)-1].Data, struct {
					Name string `json:"name"`
				}{Name: tags[2]})
			} else {
				if _, ok := d[tags[1]]; !ok {
					d[tags[1]] = true
					vehicle[len(vehicle)-1].Data = append(vehicle[len(vehicle)-1].Data, struct {
						Name string `json:"name"`
						Data []struct {
							Name string `json:"name"`
						} `json:"data"`
					}{Name: tags[1]})
				}
				vehicle[len(vehicle)-1].Data[len(vehicle[len(vehicle)-1].Data)-1].Data = append(vehicle[len(vehicle)-1].Data[len(vehicle[len(vehicle)-1].Data)-1].Data, struct {
					Name string `json:"name"`
				}{Name: tags[2]})
			}
		}
	}
	marshal, _ := json.Marshal(vehicle)
	res := []byte(`{"data":` + string(marshal) + `}`)
	fmt.Println("types:", string(res))
}
