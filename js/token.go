package js

import (
	"github.com/go-resty/resty/v2"
	"moocVip/util"
)

func Token(e string) string {
	client := resty.New()
	client.SetHeader("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
	response, err := client.R().SetFormData(map[string]string{
		"e":            e,
		"qualityValue": "1",
	}).Post("http://1.15.249.230:6000/token")
	if err != nil {
		panic(err)
	}
	videoToke := response.String()
	//fmt.Println(videoToke)
	videoToke = util.Bs64Decode(videoToke)
	//fmt.Println(videoToke)
	return videoToke
}
