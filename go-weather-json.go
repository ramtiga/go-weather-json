package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

// Weather Hacks API : example city is Tokyo(130010)
var URI string = "http://weather.livedoor.com/forecast/webservice/json/v1?city=130010"

type Forecasts struct {
    Date        string        `json:"date"`
    Datelabel   string        `json:"dateLabel"`
    Telop       string        `json:"telop"`
    Temperature Temperature   `json:"temperature"`
}

type Temperature struct {
    Max Max `json:"max"`
    Min Min `json:"min"`
}

type Max struct {
    Celsius    string `json:"celsius"`
    Fahrenheit string `json:"fahrenheit"`
}

type Min struct {
    Celsius    string `json:"celsius"`
    Fahrenheit string `json:"fahrenheit"`
}

type Location struct {
    Area       string `json:"area"`
    City       string `json:"city"`
    Prefecture string `json:"prefecture"`
}

type JsonData struct {
    Forecasts []Forecasts `json:"forecasts"`
    Location  Location  `json:"location"`
}

func main() {
    jd := new(JsonData)
    err := jd.JsonProc()
    if err != nil {
        log.Fatalf("Log: %v", err)
        return
    }
    for _, v := range jd.Forecasts {
        fmt.Printf("%v %v: %v  最高気温：%v  最低気温：%v\n", v.Date, v.Datelabel, v.Telop, v.Temperature.Max.Celsius, v.Temperature.Min.Celsius)
    }
}

func (p *JsonData) JsonProc() (err error){
    res, err := http.Get(URI)
    if err != nil {
        return err
    }

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal([]byte(b), &p)
    if err != nil {
        return err
    }
    return nil
}

