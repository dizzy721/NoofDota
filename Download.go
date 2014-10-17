package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

const key string = "536F1B65B347004500327794F4CFFC99"
const historyUrl string = "https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/V001/?key=" + key 
const detailsUrl string = "https://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/V001/?key=" + key

func main() {
	log.Println("Beginning Retreival of JSON")
	getLatest25("Noof.Dizzy")
	getMatchDetails("957303201")
	log.Println("JSON Retrieved")
}

func getLatest25(playerName string) []byte{
	var url string = historyUrl + "&player_name=" + playerName
	var ret []byte = getJson(url)
	return ret
} 

func getMatchDetails(matchId string) []byte {
	var url string = detailsUrl + "&match_id="+ matchId
	return getJson(url)
}

func getJson(url string) []byte{
	log.Println("Retrieving from url: " + url)
        res, err := http.Get(url)
        if err != nil {
                log.Fatal(err)
        }
        json,  err := ioutil.ReadAll(res.Body)
        res.Body.Close()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s", json)
        return json
}
