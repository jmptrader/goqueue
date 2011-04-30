package topiclist


import (
	"os"
	"topic"
	"subscriber"
	"message"
)


type TopicList struct {
	Topics []*topic.Topic
}

func New() *TopicList {
	return &TopicList{[...]*topic.Topic{}[:]}
}

func (t *TopicList) AddTopic(name string) (e os.Error) {
	for _, v := range t.Topics {
		if v.Name == name {
			e = os.NewError("Topic already exists")
		}
	}
	t.Topics = append(t.Topics, topic.NewTopic(name))
	return
}

func (t *TopicList) GetTopic(name string) (tpc *topic.Topic) {
	for _, v := range t.Topics {
		if v.Name == name {
			tpc = v
			break
		}
	}
	return
}

func (t *TopicList) Subscribe(topicName string, subscriberName string, subscriberEndPoint string) {
	tpc := t.GetTopic(topicName)
	if tpc != nil {
		tpc.AddSubscriber(subscriber.NewSubscriber(subscriberName, subscriberName))
	}
}

func (t *TopicList) Publish(topicName string, m message.Message) {
	tpc := t.GetTopic(topicName)
	if tpc != nil {
		tpc.Publish(m)
	}
}
///

