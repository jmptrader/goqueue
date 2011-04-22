package queue

import (
	"testing"
	"log"
)


func TestNewTopic(t *testing.T) {
	var tpc *Topic = NewTopic("topic")
	t.Log(tpc)
}

func TestAddSubscriber(t *testing.T) {
	var tpc *Topic = NewTopic("topic")
	t.Log(tpc)
	log.Println(tpc)
	var s *Subscriber = &Subscriber{"name", "endpoint"}
	tpc.AddSubscriber(s)
	t.Log(tpc)
	if len(tpc.Subscribers) != 1 {
		t.Error("Failed")
	}
	if tpc.Subscribers[0].name != "name" {
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
	var s1 *Subscriber = &Subscriber{"s1", "endpoint"}
	var s2 *Subscriber = &Subscriber{"s2", "endpoint"}
	var s3 *Subscriber = &Subscriber{"s3", "endpoint"}
	var s4 *Subscriber = &Subscriber{"s1", "endpoint"}
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
	if deleted.name != "s1" {
		t.Error("Failed")
	}
	log.Println(tpc)
}

func TestGetSubscriber(t *testing.T) {
	var tpc *Topic = NewTopic("topic")
	var s1 *Subscriber = &Subscriber{"s1", "endpoint"}
	var s2 *Subscriber = &Subscriber{"s2", "endpoint"}
	var s3 *Subscriber = &Subscriber{"s3", "endpoint"}
	var s4 *Subscriber = &Subscriber{"s1", "endpoint"}
	tpc.AddSubscriber(s1)
	tpc.AddSubscriber(s2)
	tpc.AddSubscriber(s3)
	tpc.AddSubscriber(s4)
	log.Println(tpc)
	s := tpc.RemoveSubscriber("s1")
	if s.name != "s1" {
		t.Error("Failed")
	}
	log.Println(tpc)
}


////

