package queue


import (
	"os"
)

type Message struct {
	topicName   string
	messageBody string
}

type Subscriber struct {
	name     string
	endpoint string
}

func NewSubscriber(name string, endpoint string) *Subscriber {
	return &Subscriber{name, endpoint}
}


// A Topic is a point of interest that subscribers whish to get updates on
type Topic struct {
	Name        string
	Subscribers []*Subscriber
}


func NewTopic(name string) *Topic {
	return &Topic{name, [...]*Subscriber{}[:]}
}

func (t *Topic) AddSubscriber(s *Subscriber) {
	if t.Subscribers == nil {
		t.Subscribers = [...]*Subscriber{}[:]
	}
	t.Subscribers = append(t.Subscribers, s)
}

func (t *Topic) RemoveSubscriber(name string) (s *Subscriber) {
	for i, v := range t.Subscribers {
		if v.name == name {
			s = v
			t.Subscribers = append(t.Subscribers[:i], t.Subscribers[i+1:]...)
			break
		}
	}
	return
}

func (t *Topic) GetSubscriber(name string) (s *Subscriber) {
	for _, v := range t.Subscribers {
		if v.name == name {
			s = v
			break
		}
	}
	return
}

type TopicList struct {
	Topics []*Topic
}

func (t *TopicList) New() *TopicList {
	return &TopicList{[...]*Topic{}[:]}
}

func (t *TopicList) AddTopic(name string) (e os.Error) {
	for _, v := range t.Topics {
		if v.Name == name {
			e = os.NewError("Topic already exists")
		}
	}
	t.Topics = append(t.Topics, NewTopic(name))
	return
}


///

