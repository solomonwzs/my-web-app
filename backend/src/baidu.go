package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	baiduFanyiApiUrl = "http://api.fanyi.baidu.com/api/trans/vip/translate"
)

func genSalt() string {
	i := uint64(rand.Int63())
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return fmt.Sprintf("%x", b)
}

func newBaiduTransRequest(q, from, to string) (req *http.Request, err error) {
	salt := rand.Int()

	h := md5.New()
	io.WriteString(h, strBaiduAppId)
	io.WriteString(h, q)
	io.WriteString(h, strconv.Itoa(salt))
	io.WriteString(h, baiduKey)
	sign := fmt.Sprintf("%x", h.Sum(nil))

	body := map[string]interface{}{
		"q":     q,
		"from":  from,
		"to":    to,
		"appid": baiduAppId,
		"salt":  salt,
		"sign":  sign,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	return http.NewRequest("POST", baiduFanyiApiUrl, bytes.NewBuffer(b))
}
