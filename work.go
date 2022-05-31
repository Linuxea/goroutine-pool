package goroutinepool

type WorkFun func()

func (w WorkFun) work() {
	w()
}

type Worker interface {
	work()
}
