package jsons

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type VerifyJSON struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	SiteName   string `json:"sitename"`
	Site       string `json:"site"`
	SiteUser   string `json:"siteuser"`
	About      string `json:"about"`
	VerifyCode string `json:"code"`
}

type DatasJSON struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	SiteName string `json:"sitename"`
	Site     string `json:"site"`
	SiteUser string `json:"siteuser"`
	About    string `json:"about"`
	Active   bool   `json:"active"`
}

type DataJSON struct {
	Datas  []DatasJSON  `json:"data"`
	Verify []VerifyJSON `json:"verify"`
}

func ReadAllJSON() DataJSON {
	jsonFile, err := os.Open("data/all.json")
	if err != nil {
		log.Println("[jsons] Error: falied to open json data file,", err)
		return DataJSON{}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var datas DataJSON

	json.Unmarshal(byteValue, &datas)

	return datas
}

func SetAllJson(datas DataJSON) {
	jsondata, errj := json.Marshal(datas)
	if errj != nil {
		fmt.Println("[jsons] Error: Convert json to string error,", errj)
		return
	}
	jsonFile, _ := os.OpenFile("data/all.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	jsonFile.Write(jsondata)
	jsonFile.Close()
}
