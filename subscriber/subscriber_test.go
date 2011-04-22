package subscriber

import (
	"testing"
)


func TestNewSubscriber(t *testing.T) {
	var tpc *Subscriber = NewSubscriber("topic", "endpoint")
	if tpc == nil {
		t.Error("Failed")
	}
	t.Log(tpc)
}

