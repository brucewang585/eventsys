#需求
1. 开发者简单，只要实现2个对象(source-collect,processor-exec)
2. 支持能力控制，当有空闲processor时，才调用source进行collect
3. 未来可以支持分布式

#角色
- source    收集(任务，命令，消息)
- processor 执行(任务，命令，消息)
- eventsys  调度(source,processor)  

#sample
```
type MySource struct{}

func (m *MySource) Collect(ctx context.Context, cap int) ([]interface{}, error) {	
	return []interface{}{"1","2"}, nil
}

type MyProcessor struct{}

func (m *MyProcessor) Exec(ctx context.Context, ev interface{}) error {
    return nil
}

func main() {
    sys := NewEventSys(3, 10, &MySource{}, &MyProcessor{})
    time.Sleep(time.Second * 60)
    sys.Stop()
}
```