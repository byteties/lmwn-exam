package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
		"lmwn-exam/configs"
		"lmwn-exam/apps"
		"lmwn-exam/models"
		"lmwn-exam/responses"
		"io/ioutil"
		"crypto/tls"
		"encoding/json"
)


func Summary() gin.HandlerFunc {
	return func(c *gin.Context) {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, err := http.Get(configs.EnvCovidURI())
		if err != nil {
			var msg = "Get " + configs.EnvCovidURI() + " Error"
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": msg,
			})
		}
		body, _ := ioutil.ReadAll(resp.Body)
		data := string(body)

		var result models.CovidResponseData
		jsonErr := json.Unmarshal([]byte(data), &result)

		if jsonErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": jsonErr.Error(),
			})
		}
		ageGroup,provinceGroup := apps.CountingCovid(result.Data)
		c.JSON(http.StatusOK, responses.CovidResponse{AgeGroup: ageGroup,Province: provinceGroup})
	}
}