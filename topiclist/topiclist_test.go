package topiclist

import (
	"testing"
)


func TestNew(t *testing.T) {
	tpc := New()
	if tpc == nil {
		t.Error("Failed")
	}
	t.Log(tpc)
}

func TestAddTopic(t *testing.T) {
	tpcl := New()
	tpcl.AddTopic("test")
	if tpcl.GetTopic("test") == nil {
		t.Error("Failed")
	}
	return
}
func TestSubscribe(t *testing.T) {
	tpcl := New()
	tpcl.AddTopic("test")
	tpcl.Subscribe("test", "s", "http://s.com/")
	if tpcl.GetTopic("test") == nil {
		t.Error("Failed")
	}
	if tpcl.GetTopic("test").GetSubscriber("s") == nil {
		t.Error("Failed")
	}
}

