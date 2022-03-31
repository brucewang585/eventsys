package main

import (
	"context"
	"fmt"
	. "github.com/brucewang585/eventsys/util/eventsys"
	"time"
)

type MySource struct{}

func (m *MySource) Collect(ctx context.Context, cap int) ([]interface{}, error) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " mysource B handle")
	time.Sleep(time.Second*5)

	evs := []interface{}{
		"1",
		"2",
	}
	return evs, nil
}

type MyProcessor struct{}

func (m *MyProcessor) Exec(ctx context.Context, ev interface{}) error {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " MyProcessor  handle,",ev)
	time.Sleep(time.Second * 40)
	return nil
}

func main() {
	fmt.Println("begin b")
	sys := NewEventSys(3, 10, &MySource{}, &MyProcessor{})
	time.Sleep(time.Second * 60)
	sys.Stop()
}
