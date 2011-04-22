package queuelist


import (
	"os"
	"topic"
)

type Message struct {
	topicName   string
	messageBody string
}


type TopicList struct {
	Topics []*topic.Topic
}

func (t *TopicList) New() *TopicList {
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

///

