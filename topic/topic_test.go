package topic

import (
	"testing"
	"log"
	"subscriber"
)


func TestNewTopic(t *testing.T) {
	var tpc *Topic = NewTopic("topic")
	t.Log(tpc)
}

func TestAddSubscriber(t *testing.T) {
	var tpc *Topic = NewTopic("topic")
	t.Log(tpc)
	log.Println(tpc)
	s := &subscriber.Subscriber{"name", "endpoint"}
	tpc.AddSubscriber(s)
	t.Log(tpc)
	if len(tpc.Subscribers) != 1 {
		t.Error("Failed")
	}
	if tpc.Subscribers[0].Name != "name" {
		t.Error("Failed")
	}
	tpc.AddSubscriber(s)
	tpc.AddSubscriber(s)
	tpc.AddSubscriber(s)
	tpc.AddSubscriber(s)
	log.Println(tpc)
}

func TestRemoveSubscriber(t *testing.T) {
	var tpc *Topic = NewTopic("topic")
	s1 := &subscriber.Subscriber{"s1", "endpoint"}
	s2 := &subscriber.Subscriber{"s2", "endpoint"}
	s3 := &subscriber.Subscriber{"s3", "endpoint"}
	s4 := &subscriber.Subscriber{"s1", "endpoint"}
	tpc.AddSubscriber(s1)
	tpc.AddSubscriber(s2)
	tpc.AddSubscriber(s3)
	tpc.AddSubscriber(s4)
	log.Println(tpc)
	if len(tpc.Subscribers) != 4 {
		t.Error("Failed")
	}
	deleted := tpc.RemoveSubscriber("s1")
	if len(tpc.Subscribers) != 3 {
		t.Error("Failed")
	}
	if deleted.Name != "s1" {
		t.Error("Failed")
	}
	log.Println(tpc)
}

func TestGetSubscriber(t *testing.T) {
	var tpc *Topic = NewTopic("topic")
	s1 := &subscriber.Subscriber{"s1", "endpoint"}
	s2 := &subscriber.Subscriber{"s2", "endpoint"}
	s3 := &subscriber.Subscriber{"s3", "endpoint"}
	s4 := &subscriber.Subscriber{"s1", "endpoint"}
	tpc.AddSubscriber(s1)
	tpc.AddSubscriber(s2)
	tpc.AddSubscriber(s3)
	tpc.AddSubscriber(s4)
	log.Println(tpc)
	s := tpc.RemoveSubscriber("s1")
	if s.Name != "s1" {
		t.Error("Failed")
	}
	log.Println(tpc)
}


////

