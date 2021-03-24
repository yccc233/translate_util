package translate

import (
	"TranslateUtil/constent"
	"TranslateUtil/httptra"
	"TranslateUtil/lock"
	"TranslateUtil/stru"
	"encoding/json"
	"time"
)

/*
中 -> 英

返回结构体 以及 原串
*/
func getTransCtEStruct(trans string) (stru.Resp,string) {
	url := httptra.GetTansURL(trans,constent.LANG_ZH, constent.LANG_EN)
	str := httptra.GetHttp(url)

	//str := "{\"from\":\"en\",\"to\":\"zh\",\"trans_result\":[{\"src\":\"lie\",\"dst\":\"\\u8eba\"}]}\n"

	var res = stru.Resp{}

	err := json.Unmarshal([]byte(str),&res)
	if err != nil {
		panic(err)
	}
	return res,str
}


/*
中 -> 英

返回第一个翻译
*/
func GetTransCtESMean(trans string) string{
	time.Sleep(time.Second * time.Duration(lock.LockSeconds()))


	res,_ := getTransCtEStruct(trans)

	return res.TransResult[0].Dst
}