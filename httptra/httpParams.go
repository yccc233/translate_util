package httptra

import (
	"TranslateUtil/constent"
	"TranslateUtil/stru"
	"TranslateUtil/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)


/*
返回翻译url 懒人做法
 */
func GetTansURL(q,f,t string)  string{
	return GetWholeUrl(constent.URL_BAIDU_PRE,q,f,t)
}

/*
返回带域名参数的url
 */
func GetWholeUrl(pre, q,f,t string) string{
	p := pieceParams(q,f,t)
	return pre+"?"+getParamsUrl(p)
}

/*
生成参数url
 */
func getParamsUrl(p stru.Params) string {
	return fmt.Sprintf(constent.URL_FORMAT,
		p.Query,
		p.From,
		p.To,
		p.AppId,
		p.Salt,
		p.Sign,
	)
}

/*
拼接参数，将参数补全
 */
func pieceParams(q,f,t string) stru.Params {
	var p = stru.Params{Query: q, From: f, To: t}
	p.AppId = constent.APPID
	p.Salt = utils.GetRand(20)
	p.Sign = makeSignMd5(p)
	return p
}

/*
生成标签，要求参数p中APPID，query，随机数	，密钥 有值
 */
func makeSignMd5(p stru.Params) string {
	var sign = p.AppId + p.Query + p.Salt + constent.SECRETID
	m := md5.New()
	m.Write([]byte(sign))
	return hex.EncodeToString(m.Sum(nil))
}