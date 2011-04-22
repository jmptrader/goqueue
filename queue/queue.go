package queue


import (
	"os"
	"subscriber"
)

type Message struct {
	topicName   string
	messageBody string
}


// A Topic is a point of interest that subscribers whish to get updates on
type Topic struct {
	Name        string
	Subscribers []*subscriber.Subscriber
}


func NewTopic(name string) *Topic {
	return &Topic{name, [...]*subscriber.Subscriber{}[:]}
}

func (t *Topic) AddSubscriber(s *subscriber.Subscriber) {
	if t.Subscribers == nil {
		t.Subscribers = [...]*subscriber.Subscriber{}[:]
	}
	t.Subscribers = append(t.Subscribers, s)
}

func (t *Topic) RemoveSubscriber(name string) (s *subscriber.Subscriber) {
	for i, v := range t.Subscribers {
		if v.Name == name {
			s = v
			t.Subscribers = append(t.Subscribers[:i], t.Subscribers[i+1:]...)
			break
		}
	}
	return
}

func (t *Topic) GetSubscriber(name string) (s *subscriber.Subscriber) {
	for _, v := range t.Subscribers {
		if v.Name == name {
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

func (t *TopicList) GetTopic(name string) (tpc *Topic) {
	for _, v := range t.Topics {
		if v.Name == name {
			tpc = v
			break
		}
	}
	return
}

///

