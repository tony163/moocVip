package js

import (
	"github.com/go-resty/resty/v2"
	"moocVip/util"
)

func M3u8(e string, t string) string {
	cleint := resty.New()
	i := "False"
	response, err := cleint.R().SetFormData(map[string]string{
		"e": e,
		"t": t,
		"i": i,
	}).Post("http://1.15.249.230:6000/secondarydecrypt")
	if err != nil {
		panic(err)
	}
	m3u8 := response.String()
	m3u8 = util.Bs64Decode(m3u8)
	//fmt.Println(m3u8)
	return m3u8
}

func Key(e string, t string) string {
	cleint := resty.New()
	i := "True"
	response, err := cleint.R().SetFormData(map[string]string{
		"e": e,
		"t": t,
		"i": i,
	}).Post("http://1.15.249.230:6000/secondarydecrypt")
	if err != nil {
		panic(err)
	}
	key := response.String()
	key = util.Bs64Decode(key)
	//fmt.Println(key)
	//uri, _ := url.Parse("https://www.google.com/")
	//cleint.GetClient().Jar.Cookies(uri)
	return key
}
