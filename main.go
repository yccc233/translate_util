package main

import (
	"TranslateUtil/constent"
	"TranslateUtil/translate"
	"fmt"
	"net/http"
)

//fmt.Println(translate.GetTransEtCSMean("ABO"))
//
//fmt.Println(translate.GetTransCtESMean("很高兴认识你！"))


func main() {

	//提供英-汉译文
	http.HandleFunc("/etc", func(writer http.ResponseWriter, request *http.Request) {

		trans := request.URL.Query().Get("q")
		var bd string

		if trans == "" {
			bd = "请输入正确的参数名称或参数"
		}else {
			bd = translate.GetTransEtCSMean(trans)
		}
		_,err := writer.Write([]byte(bd))
		if err != nil {
			panic("http trans error ! made by this program!")
		}
		fmt.Println("请求成功！")
	})

	//提供汉-英译文
	http.HandleFunc("/cte", func(writer http.ResponseWriter, request *http.Request) {

		trans := request.URL.Query().Get("q")
		var bd string
		if trans == "" {
			bd = "请输入正确的参数名称或参数"
		}else {
			bd = translate.GetTransCtESMean(trans)
		}
		_,err := writer.Write([]byte(bd))
		if err != nil {
			panic("http trans error ! made by this program!")
		}
		fmt.Println("请求成功！")
	})

	//根据自己选择提供翻译方向
	http.HandleFunc("/translate", func(writer http.ResponseWriter, request *http.Request) {

		trans := request.URL.Query().Get("query")
		var bd string

		if trans == "" || request.URL.Query().Get("from") == "" || request.URL.Query().Get("to") == ""{
			bd = "请输入正确的url"
		}else {
			if request.URL.Query().Get("from") == constent.LANG_ZH {
				bd = translate.GetTransCtESMean(trans)
			}else {
				bd = translate.GetTransEtCSMean(trans)
			}


		}
		_,err := writer.Write([]byte(bd))
		if err != nil {
			panic("http trans error ! made by this program!")
		}
		fmt.Println("请求成功！")
	})


	fmt.Println("即将开启8888端口通道，请使用以下形式访问：\n" +
		" /etc?q=hello\n" +
		" /cte?q=你好\n" +
		" /translate?query=你好&from=zh&to=en\n" +
		"...")


	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
