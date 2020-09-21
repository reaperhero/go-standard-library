package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Idiom struct {
	Title string
}

func main() {
	showapi_appid := "372021"
	showapi_sign := "ffb2608e59364988922fc2985227eefd"
	keyword := "肉"
	url := "http://route.showapi.com/1196-1?showapi_appid=" + showapi_appid + "&keyword=" + keyword + "&showapi_sign=" + showapi_sign


	resp, _ := http.Get(url)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	resultJson := string(body)
	fmt.Println(resultJson)
	//{
	//	"showapi_res_error": "",
	//	"showapi_res_id": "5f6772648d57babe4e13e98e",
	//	"showapi_res_code": 0,
	//	"showapi_res_body": {"ret_message":"Success","ret_code":0,"last_page":11,"total":105,"data":[{"title":"皮开肉绽"},{"title":"白骨再肉"},{"title":"髀里肉生"},{"title":"髀肉复生"},{"title":"不吃羊肉空惹一身膻"},{"臭肉来蝇"},{"title":"肥鱼大肉"},{"title":"凡夫肉眼"},{"title":"凡胎肉眼"}]}
	//}


	tempMap := make(map[string]interface{})
	json.Unmarshal(body, &tempMap)
	fmt.Println(tempMap)
	//map[showapi_res_body:map[data:[map[title:皮开肉绽] map[title:白骨再肉] map[title:髀里肉生] map[title:髀肉复生] map[title:不吃羊肉空惹一身膻] map[title:不知肉味] map[title:臭肉来蝇] map[title:肥鱼大肉] map[title:凡夫肉眼] map[tissage:Success total:105] showapi_res_code:0 showapi_res_error: showapi_res_id:5f6772648d57babe4e13e98e]

	dataSilce := tempMap["showapi_res_body"].(map[string]interface{})["data"].([]interface{})
	idiomMap := make(map[string]Idiom)
	for _, v := range dataSilce {
		title := v.(map[string]interface{})["title"].(string)
		idiom := Idiom{Title: title}
		idiomMap[title] = idiom
	}
	fmt.Println(idiomMap)
	// map[不吃羊肉空惹一身膻:{不吃羊肉空惹一身膻} 不知肉味:{不知肉味} 凡夫肉眼:{凡夫肉眼} 凡胎肉眼:{凡胎肉眼} 白骨再肉:{白骨再肉} 皮开肉绽:{皮开肉绽} 肥鱼大肉:{肥鱼大肉} 臭肉来蝇:{臭肉来蝇} 髀肉复生:{髀肉复生} 髀里肉生:{髀里肉生}]
}
