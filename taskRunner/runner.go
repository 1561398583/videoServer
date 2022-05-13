package taskRunner

type Runner struct {
	Controller ControlChan
	Error ControlChan
	Data DataChan
	dataSize int
	longLived bool
	Dispatcher Fn
	Executer Fn
}

func NewRunner(dataSize int, longLived bool, dispatcher Fn, executer Fn)  *Runner{
	return &Runner{
		Controller: make(chan string, 1),	//1个size是为了成为非阻塞channel
		Error: make(chan string, 1),
		Data: make(chan interface{}, dataSize),
		dataSize: dataSize,
		longLived: longLived,
		Dispatcher: dispatcher,
		Executer: executer,
	}
}

func (r *Runner) startDispatch()  {
	defer func() {
		if !r.longLived {	//如果不是长久存在的，则关闭资源
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()

	select {
	case c :=<-r.Controller:
		if c == READY_TO_DISPATCH {
			err := r.Dispatcher(r.Data)
			if err != nil {
				r.Error <-CLOSE
			}else {
				r.Controller <-READY_TO_EXECUTE
			}
		}

		if c == READY_TO_EXECUTE {
			err := r.Executer(r.Data)
			if err != nil {
				r.Error <-CLOSE
			}else {
				r.Controller <- READY_TO_DISPATCH
			}
		}
	case e :=<-r.Error:
		if e == CLOSE {
			return
		}
	default:

		
	}
}

func (r *Runner) Start()  {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
