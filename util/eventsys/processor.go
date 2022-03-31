package eventsys

import "context"

type Processor interface {
	Exec(ctx context.Context, msg interface{}) error
}
