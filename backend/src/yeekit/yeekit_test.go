package yeekit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestYeekitTrans(t *testing.T) {
	req, err := newYeekitTransRequest("apple", "en", "zh")
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
