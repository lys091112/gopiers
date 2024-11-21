package common

// 根据您选择的AK已为您生成调用代码
// 检测您当前的AK设置了sn检验，本示例中已为您生成sn计算代码

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func calculateSN(queryStr string, sk string) string {
	str := queryStr + sk
	key := md5.Sum([]byte(str))
	sn := fmt.Sprintf("%x", key)
	return sn
}

func BaiduSearch() {
	// 此处填写您在控制台-应用管理-创建应用后获取的AK
	ak := "t7HgxEz7LrwS9qQO2j8Gx1yQWRqHecwe"

	// 此处填写您在控制台-应用管理-创建应用时，校验方式选择sn校验后生成的SK
	sk := "YY1yOEaHWheZ9OyVcL3cx8Ad0hbrJ2sW"

	// 服务地址
	host := "https://api.map.baidu.com"

	// 接口地址
	uri := "/place/v2/search"

	// 设置请求参数
	params := [][]string{
		{"query", "上海交通大学"},
		{"tag", "政府机构"},
		{"region", "全国"},
		{"output", "json"},
		{"ak", ak},
	}

	paramsArr := make([]string, 0)
	for _, v := range params {
		kv := v[0] + "=" + (v[1])
		paramsArr = append(paramsArr, kv)
	}
	paramsStr := strings.Join(paramsArr, "&")

	// 计算sn
	queryStr := url.QueryEscape(uri + "?" + paramsStr)
	sn := calculateSN(queryStr, sk)

	// 发起请求
	request, err := url.Parse(host + uri + "?" + paramsStr + "&sn=" + sn)
	if nil != err {
		fmt.Printf("host error: %v", err)
		return
	}

	resp, err1 := http.Get(request.String())
	// 请注意，此处打印的url为非urlencode后的请求串
	// 如果将该请求串直接粘贴到浏览器中发起请求，由于浏览器会自动进行urlencode，会导致返回sn校验失败
	fmt.Printf("url: %s\n", request.String())
	defer resp.Body.Close()
	if err1 != nil {
		fmt.Printf("request error: %v", err1)
		return
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Printf("response error: %v", err2)
	}

	fmt.Println(string(body))
}
