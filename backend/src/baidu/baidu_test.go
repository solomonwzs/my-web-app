package baidu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBaiduTrans(t *testing.T) {
	req, err := newBaiduTransRequest("apple", "auto", "zh")
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
