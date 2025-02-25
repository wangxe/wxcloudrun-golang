package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	clientId     = "d6bd0b0f5f465b23b704"
	clientSecret = "80b3f02de8c5b1fda2deaeb4a5a790e4a9550eee"
)

var token map[string]interface{}

// HelloHandler hello接口
func AuthorizeHandler(c *gin.Context) {

	request, err := http.NewRequest(http.MethodGet, "https://github.com/login/oauth/authorize", nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("client_id", clientId)
	params.Add("redirect_uri", "https://golang-3k3w-1824429-1311191960.ap-shanghai.run.tcloudbase.com/code/12345")
	params.Add("login", "login")
	params.Add("scope", "user%2Crepo%2Cread%3Arepo_hook%2Cwrite%3Arepo_hook%2Cadmin%3Arepo_hook%2Cadmin%3Aorg_hook%2Cread%3Aorg")
	request.URL.RawQuery = params.Encode()

	_, err = http.DefaultClient.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func GetCode(c *gin.Context) {
	code := c.Query("code")

	data := make(map[string]string)
	data["client_id"] = clientId
	data["client_secret"] = clientSecret
	data["code"] = code
	data["redirect_uri"] = "https://golang-3k3w-1824429-1311191960.ap-shanghai.run.tcloudbase.com/token/12345"
	marshal, _ := json.Marshal(data)

	body := strings.NewReader(string(marshal))

	resp, err := http.Post("https://github.com/login/oauth/access_token", "application/json", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	token = make(map[string]interface{})
	token["data"] = all

	c.JSON(http.StatusOK, "ok")

}

func GetToken(c *gin.Context) {

	if token != nil {
		c.JSON(http.StatusOK, token)
		return
	}

	c.JSON(http.StatusOK, "token not exist")

}
