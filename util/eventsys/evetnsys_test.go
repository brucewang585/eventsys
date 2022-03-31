package eventsys

import (
	"context"
	"fmt"
	"testing"
	"time"
)


type mySource struct {

}
func (m *mySource) Collect(ctx context.Context, cap int) ([]interface{}, error) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")," mysource B handle")

	evs := []interface{}{
		"",
		"",
	}
	return evs,nil
}

type myProcessor struct {

}
func (m *myProcessor) Exec(ctx context.Context, ev interface{}) error {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")," MyProcessor  handle")
	time.Sleep(time.Second*40)
	return nil
}

func TestSum(t *testing.T)  {

	fmt.Println("begin b")
	sys := NewEventSys(3,10,&mySource{},&myProcessor{})
	time.Sleep(time.Second*60)
	sys.Stop()
}


