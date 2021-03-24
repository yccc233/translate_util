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
英 -> 中

返回结构体 以及 原串
 */
func getTransEtCStruct(trans string) (stru.Resp,string) {

	url := httptra.GetTansURL(trans,constent.LANG_EN, constent.LANG_ZH)
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
英 -> 中

返回第一个翻译
*/
func GetTransEtCSMean(trans string) string{

	time.Sleep(time.Second * time.Duration(lock.LockSeconds()))

	res,_ := getTransEtCStruct(trans)


	return res.TransResult[0].Dst
}