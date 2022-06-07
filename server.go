package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Rezervation struct {
	Tren           Tren `json:"Tren"`
	RezYapanSayisi int  `json:"RezervasyonYapilacakKisiSayisi"`
	FarkliVagon    bool `json:"KisilerFarkliVagonlaraYerlestirilebilir"`
}

type Tren struct {
	Ad       string     `json:"Ad"`
	Vagonlar []Vagonlar `json:"Vagonlar"`
}

type Vagonlar struct {
	Ad             string `json:"Ad"`
	Kapasite       int    `json:"Kapasite"`
	DoluKoltukAdet int    `json:"DoluKoltukAdet"`
}

type Vagon struct {
	Ad             string `json:"Ad"`
	Kapasite       int    `json:"Kapasite"`
	DoluKoltukAdet int    `json:"DoluKoltukAdet"`
}

func GetInfo(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var rezervation Rezervation
	json.Unmarshal([]byte(value), &rezervation)
	l := len(rezervation.Tren.Vagonlar)

	var rez Train
	person := rezervation.RezYapanSayisi

	for i := 0; i < l; i++ {
		if person <= 0 {
			break
		}
		a := person
		rez = GetRez(rezervation.Tren.Vagonlar[i].Ad, rezervation.Tren.Vagonlar[i].DoluKoltukAdet, rezervation.Tren.Vagonlar[i].Kapasite, rezervation.FarkliVagon)
		person, err = rez.RezervationFunc(person)
		if err != nil {
			c.JSON(200, map[string]any{
				"error": err,
			})
			return
		}
		a = a - person

	}

	c.JSON(200, gin.H{
		"RezervasyonYapilabilir": true,
		"YerlesimAyrinti":        rezervation.Tren.Vagonlar[0].Ad,
	})

}

func GetRez(carName string, capacity, fullness int, difCar bool) Train {
	return Train{
		CarName:  carName,
		Fullness: fullness,
		Capacity: capacity,
		difCar:   difCar,
	}
}
