package baidu

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	baiduFanyiAPIURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"
)

var (
	strBaiduAppID = os.Getenv("BAIDU_APPID")
	baiduKey      = os.Getenv("BAIDU_KEY")
	baiduAppID    int64
)

func init() {
	rand.Seed(time.Now().UnixNano())
	baiduAppID, _ = strconv.ParseInt(strBaiduAppID, 10, 64)
}

func genSalt() string {
	i := uint64(rand.Int63())
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return fmt.Sprintf("%x", b)
}

func newBaiduTransRequest(q, from, to string) (req *http.Request, err error) {
	salt := rand.Int()

	h := md5.New()
	io.WriteString(h, strBaiduAppID)
	io.WriteString(h, q)
	io.WriteString(h, strconv.Itoa(salt))
	io.WriteString(h, baiduKey)
	sign := fmt.Sprintf("%x", h.Sum(nil))

	body := map[string]interface{}{
		"q":     q,
		"from":  from,
		"to":    to,
		"appid": baiduAppID,
		"salt":  salt,
		"sign":  sign,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	return http.NewRequest("POST", baiduFanyiAPIURL, bytes.NewBuffer(b))
}
