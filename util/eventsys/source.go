package eventsys

import "context"

type Source interface {
	//cap: 允许最多*msg的数目，返回如果大于cap,会丢弃多余的
	Collect(ctx context.Context, cap int) ([]interface{}, error)
}
