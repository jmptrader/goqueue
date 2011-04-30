package subscriber

import (
	"testing"
	"fmt"
	"http"
	"http/httptest"
	"message"
	"io/ioutil"
	"io"
	"os"
)

func TestNewSubscriber(t *testing.T) {
	var tpc *Subscriber = NewSubscriber("topic", "endpoint")
	if tpc == nil {
		t.Error("Failed")
	}
	t.Log(tpc)
}


type testHandler struct {
	body string
	req  *http.Request
}

func (th *testHandler) handlef() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var b []byte
		b, _ = ioutil.ReadAll(req.Body)
		req.Body.Close()
		th.body = string(b)
		th.req = req
	}
}

var th = new(testHandler)

var thf = http.HandlerFunc(th.handlef())


type testClient struct {
	body string
	url  string
}

func (tc *testClient) Post(url string, bodyType string, body io.Reader) (r *http.Response, err os.Error) {
	var b []byte
	b, _ = ioutil.ReadAll(body)
	tc.body = string(b)
	tc.url = url
	fmt.Println(tc)
	return nil, nil
}
func TestNotify(t *testing.T) {
	ts := httptest.NewServer(thf)
	defer ts.Close()
	testClient := &testClient{}
	url, _ := http.ParseURL(ts.URL)
	var tpc *Subscriber = &Subscriber{"topic", url, testClient}
	tpc.Notify(message.Message{"test message body"})
	if testClient.body != "test message body" {
		t.Error("Failed")
	}
	var sub *Subscriber = NewSubscriber("topic", ts.URL)
	sub.Notify(message.Message{"real client test message body"})
	if th.body != "real client test message body" {
		t.Error("Message not posted by subscriber")
	}
}

