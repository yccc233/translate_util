package constent

import "time"

const (
	URL_BAIDU_PRE = "http://api.fanyi.baidu.com/api/trans/vip/translate"

	URL_FORMAT = "q=%s&from=%s&to=%s&appid=%s&salt=%s&sign=%s"

	APPID = "20210324000740657"

	SECRETID = "13pXNyQraMMH2I2H_bTM"

	//语种，目前仅中文 英文
	LANG_ZH = "zh"
	LANG_EN = "en"

	LOCKTIME = time.Second
)
