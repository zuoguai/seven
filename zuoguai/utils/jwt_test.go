package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"strings"
	"testing"
	"zuoguai/internal/config"
)

func TestParseJwtToken(t *testing.T) {

	r, err := ParseJwtToken("", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

type News struct {
	Code  int      `json:"code"`
	Msg   string   `json:"msg"`
	Data  NewsData `json:"data"`
	Time  int      `json:"time"`
	Usage int      `json:"usage"`
	LogId string   `json:"log_id"`
}

type NewsData struct {
	Data      string   `json:"date"`
	News      []string `json:"news"`
	WeiYu     string   `json:"weiyu"`
	Image     string   `json:"image"`
	HeadImage string   `json:"head_image"`
}

func TestGetNews(t *testing.T) {
	configs := config.GetConfigs("C:\\Users\\作怪\\Desktop\\app.yaml")

	url := "https://v2.alapi.cn/api/zaobao"

	payload := strings.NewReader(fmt.Sprintf("token=%s&format=json", configs.PrivateToken.ApiToken))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println("????", string(body))
	news := News{}
	if err := binding.JSON.BindBody(body, &news); err != nil {
		fmt.Println(err)
	}
	fmt.Println(news)

	fmt.Println(news.Data.Image)
	fmt.Println(news.Data.HeadImage)
	for i := 0; i < len(news.Data.News); i++ {
		fmt.Println(news.Data.News[i])
	}

}
