package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	type Item struct {
		Course string `json:"course"`
		Work   string `json:"work"`
	}
	num := 0
	var items []Item
	url := "https://passport2-api.chaoxing.com/v11/loginregister?code=" + "你的学习通密码" + "&cx_xxt_passport=json&uname=" + "你的学习通手机号" + "&loginType=1&roleSelect=true"
	method := "GET"
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	res, err := client.Do(req)
	cookie := res.Cookies()
	cookies := ""
	for _, cook := range cookie {
		cookies += cook.Raw + ";"
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	print(string(body))
	url = "http://mooc1-api.chaoxing.com/work/stu-work"
	req, _ = http.NewRequest(method, url, nil)
	req.Header.Add("Cookie", cookies)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")
	res, _ = client.Do(req)
	if res != nil {
		defer res.Body.Close()
		//body, _ = ioutil.ReadAll(res.Body)
	}
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find("#content").Each(func(i int, selection1 *goquery.Selection) {
		selection1.Find(".redPoint+div").Each(func(i int, selection2 *goquery.Selection) {
			if selection2.Find(".fr").Text() != "" {
				course := selection2.Find(".status+span").Text()
				work := selection2.Find("p").Text()
				item := Item{
					Course: course,
					Work:   work,
				}
				items = append(items, item)
				num += 1
			}
		})
	})
	if len(items) != 0 {
		url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=微信公众appid&secret=微信公众secret"
		req, _ = http.NewRequest("GET", url, nil)
		res1, _ := client.Do(req)
		defer res1.Body.Close()
		body, _ = ioutil.ReadAll(res1.Body)
		access_token := gjson.Get(string(body), "access_token")
		url = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token.String()
		type ValueColor struct {
			Value string `json:"value"`
			Color string `json:"color"`
		}
		type WxData struct {
			Course ValueColor `json:"course"`
			Num    ValueColor `json:"num"`
			Work   ValueColor `json:"work"`
		}
		type WxPayload struct {
			Touser     string `json:"touser"`
			TemplateId string `json:"template_id"`
			Url        string `json:"url"`
			Data       WxData `json:"data"`
		}
		var courses string
		var works string
		for _, v := range items {
			courses += v.Course
			works += v.Work
		}
		payload := WxPayload{
			Touser:     "你的微信openid",
			TemplateId: "你的模板消息id",
			Data: WxData{
				Course: ValueColor{
					Value: courses,
					Color: "#f1939c",
				},
				Work: ValueColor{
					Value: works,
					Color: "#61ac85",
				},
				Num: ValueColor{
					Value: strconv.Itoa(num),
					Color: "#2775b6",
				},
			},
		}
		payloaddata, _ := json.Marshal(payload)
		data := bytes.NewReader(payloaddata)
		req, _ = http.NewRequest("POST", url, data)
		res2, _ := client.Do(req)
		defer res2.Body.Close()
		body, _ = ioutil.ReadAll(res2.Body)
		fmt.Printf("string(body): %v\n", string(body))
	}
}
