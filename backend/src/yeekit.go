package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	yeekitFanyiApiUrl = "http://api.yeekit.com/dotranslate.php"
)

func newYeekitTransRequest(q, from, to string) (req *http.Request, err error) {
	body := map[string]interface{}{
		"text":    q,
		"from":    from,
		"to":      to,
		"app_id":  yeekitAppId,
		"app_key": yeekitKey,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return
	}
	fmt.Println(string(b))

	return http.NewRequest("POST", yeekitFanyiApiUrl, bytes.NewBuffer(b))
}
