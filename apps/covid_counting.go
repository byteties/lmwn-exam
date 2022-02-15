package apps

import (
	"lmwn-exam/models"
)

type Provinces map[string] int
type Ages map[string] int

func GetCountAges(ages Ages,age int) {
	if age < 0 {
		ages["N/A"] += 1
	}	else if age < 31 {
		ages["0-30"] += 1
	} else if age < 61 {
		ages["31-60"] += 1
	} else {
		ages["61+"] += 1
	}
}

func GetCountProvince(province Provinces,key string) {
	if key == "" {
		key = "N/A"
	}
	v,ok := province[key]
	if !ok {
		province[key] = 1
	}else {
		province[key] = v + 1
	}
}

func CountingCovid(input []models.CovidData) (Ages,Provinces){
	ages := Ages{}
	province := Provinces{}
	for i := 0; i < len(input); i++ {
		data := input[i]
		GetCountAges(ages,data.Age)
		GetCountProvince(province,data.Province)
	}
	return ages,province
}