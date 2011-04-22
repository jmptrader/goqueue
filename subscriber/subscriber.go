package subscriber


import ()


type Subscriber struct {
	Name     string
	Endpoint string
}

func NewSubscriber(name string, endpoint string) *Subscriber {
	return &Subscriber{name, endpoint}
}

