package httptra

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func GetHttp(url string) string {
	client := &http.Client{Timeout: 10*time.Second}

	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	/*
		防止response过长，可以使用计时器，暂时用不到
	 */
	var buffer [2048]byte

	result := bytes.NewBuffer(nil)

	for {
		n, err := response.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

