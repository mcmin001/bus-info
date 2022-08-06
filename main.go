package main

import (
	http_util "bus-info/util/http_util"
	"fmt"
)

func main() {
	var url = "http://real-time-bus.whgjzt.com:9087/website//web/420100/line/027-903-0.do?Type=LineDetail"
	result := http_util.HttpGet(url, buildHeader())
	fmt.Println("response :" + result)
}

func buildHeader() (header []http_util.Header) {
	var headerSlice []http_util.Header
	header_User_Agent := http_util.Header{Key: "User-Agent", Value: "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36"}
	header_Accept := http_util.Header{Key: "Accept", Value: "application/json, text/javascript, */*; q=0.01"}
	header_Host := http_util.Header{Key: "Host", Value: "real-time-bus.whgjzt.com:9087"}
	header_Origin := http_util.Header{
		Key:   "Origin",
		Value: "http://wbus.wuhancloud.cn",
	}
	header_Referer := http_util.Header{
		Key:   "Referer",
		Value: "http://wbus.wuhancloud.cn/",
	}

	headerSlice = append(headerSlice, header_User_Agent)
	headerSlice = append(headerSlice, header_Accept)
	headerSlice = append(headerSlice, header_Host)
	headerSlice = append(headerSlice, header_Origin)
	headerSlice = append(headerSlice, header_Referer)

	return headerSlice
}
