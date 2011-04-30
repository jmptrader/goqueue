package subscriber


import (
	"http"
	"message"
	"strings"
	"io"
	"os"
)

type HttpPostClient interface {
	Post(url string, bodyType string, body io.Reader) (r *http.Response, err os.Error)
}

type Subscriber struct {
	Name     string
	Endpoint *http.URL
	client   HttpPostClient
}

var defaultClient = new(http.Client)

func NewSubscriber(name string, endpoint string) *Subscriber {
	s := new(Subscriber)
	s.Name = name
	s.Endpoint, _ = http.ParseURL(endpoint)
	return s
}

func (s *Subscriber) Notify(m message.Message) {
	body := strings.NewReader(m.MessageBody)
	if s.client != nil {
		s.client.Post(s.Endpoint.String(), m.MessageBody, body)
	} else {
		defaultClient.Post(s.Endpoint.String(), m.MessageBody, body)
	}
}

