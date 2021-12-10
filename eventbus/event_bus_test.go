package eventbus

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	bus := New()
	if bus == nil {
		t.Log("New EventBus not created!")
		t.Fail()
	}
}

func TestHasCallback(t *testing.T) {
	bus := New()
	bus.Subscribe("topic", func() {})
	if bus.HasCallback("topic_topic") {
		t.Fail()
	}
	if !bus.HasCallback("topic") {
		t.Fail()
	}
}

func TestSubscribe(t *testing.T) {
	bus := New()
	if bus.Subscribe("topic", func() {}) != nil {
		t.Fail()
	}
	if bus.Subscribe("topic", "String") == nil {
		t.Fail()
	}
}

func TestSubscribeOnceAndManySubscribe(t *testing.T) {
	bus := New()
	event := "topic"
	flag := 0
	fn := func() { flag += 1 }
	bus.Subscribe(event, fn)
	bus.Subscribe(event, fn)
	bus.Publish(event)

	if flag != 2 {
		t.Fail()
	}
}

func TestUnsubscribe(t *testing.T) {
	bus := New()
	handler := func() {}
	bus.Subscribe("topic", handler)
	if bus.Unsubscribe("topic", handler) != nil {
		t.Fail()
	}
	if bus.Unsubscribe("topic", handler) == nil {
		t.Fail()
	}
}

func TestPublish(t *testing.T) {
	bus := New()
	bus.Subscribe("topic", func(a int, b int) {
		if a != b {
			t.Fail()
		}
	})
	bus.Publish("topic", 10, 10)
}

func TestSubscribeAsync(t *testing.T) {
	results := make(chan int)

	bus := New()
	bus.SubscribeAsync("topic", func(a int, out chan<- int) {
		out <- a
	})

	bus.Publish("topic", 1, results)
	bus.Publish("topic", 2, results)

	numResults := 0
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for _ = range results {
			numResults++
			wg.Done()
		}
	}()
	wg.Wait()

	if numResults != 2 {
		t.Fail()
	}
}
