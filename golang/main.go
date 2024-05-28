package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func randomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getHeadersAuthen() map[string]string {
	currentTimestamp := time.Now().UnixMilli()
	timeDiff := -randomNumber(99, 199)
	currentTimestamp = currentTimestamp - int64(timeDiff)

	date := time.UnixMilli(currentTimestamp)
	day := fmt.Sprintf("%02d", date.Day())
	month := fmt.Sprintf("%02d", date.Month())
	year := fmt.Sprintf("%d", date.Year())
	hour := fmt.Sprintf("%02d", date.Hour())
	minute := fmt.Sprintf("%02d", date.Minute())
	second := fmt.Sprintf("%02d", date.Second())

	dateValue := year + month + day
	timeValue := hour + minute + second
	md5Value := getMD5Hash(dateValue + timeValue)
	keyValue := md5Value[:3] + md5Value[len(md5Value)-3:]

	keyAccess := "Kh0ngDuLieu" + dateValue + "C0R0i" + timeValue + "Kh0aAnT0an" + keyValue
	headers := map[string]string{
		"X-SFD-Key":    getMD5Hash(keyAccess),
		"X-SFD-Date":   dateValue + timeValue,
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
	}
	return headers
}

func getData(channel string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.thvli.vn/backend/cm/get_detail/%s-hd/?timezone=Asia/Ho_Chi_Minh", channel)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	headers := getHeadersAuthen()
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Structure for the response body
type Response struct {
	PlayInfo struct {
		Data struct {
			LinkPlay string `json:"link_play"`
		} `json:"data"`
	} `json:"play_info"`
}

// Gin handler function
func tvHandler(ctx *gin.Context) {
	channel := ctx.Query("channel")
	resp, err := getData(channel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	linkPlay := result.PlayInfo.Data.LinkPlay
	if linkPlay == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	ctx.Redirect(http.StatusFound, linkPlay)
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/tv", tvHandler)
	r.Run(":43009") // Run the server on port 8080
}
