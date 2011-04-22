package topic


import (
	"subscriber"
)


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

