package stru


type Params struct {
	Query string
	From  string
	To    string
	AppId string
	Salt  string
	Sign  string
}




/*
返回体的结构体
*/
type Resp struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct{
		Src 	string `json:"src"`
		Dst 	string `json:"dst"`
	} `json:"trans_result"`
}

type Req struct {
	Query string `json:"query"`
	From        string `json:"from"`
	To          string `json:"to"`
}
