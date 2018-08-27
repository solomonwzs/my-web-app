package yeekit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	yeekitFanyiAPIURL = "http://api.yeekit.com/dotranslate.php"
)

var (
	yeekitAppID = os.Getenv("YEEKIT_APPID")
	yeekitKey   = os.Getenv("YEEKIT_KEY")
)

func newYeekitTransRequest(q, from, to string) (req *http.Request, err error) {
	body := map[string]interface{}{
		"text":    q,
		"from":    from,
		"to":      to,
		"app_id":  yeekitAppID,
		"app_key": yeekitKey,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return
	}
	fmt.Println(string(b))

	return http.NewRequest("POST", yeekitFanyiAPIURL, bytes.NewBuffer(b))
}
