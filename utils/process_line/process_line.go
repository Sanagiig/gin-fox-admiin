package process_line

import "gin-one/message"

type PLThenFn = func() (msgCode int, err error)
type PLCatchFn = func(err error) (isCatchContinue bool)
type PLRecoverFn = func(err error)
type PLFinallyFn = func()

type ProcessLine struct {
	isCatchContinue bool // catch 后是否继续执行后续的 catch
	msgCode         int
	err             error
}

func (pl *ProcessLine) Then(fn PLThenFn) *ProcessLine {
	if pl.err == nil {
		pl.msgCode, pl.err = fn()
		// 默认错误信息
		if pl.err != nil && pl.msgCode == 0 {
			pl.msgCode = message.OPER_ERR
		}
	}
	return pl
}

func (pl *ProcessLine) Catch(fn PLCatchFn) *ProcessLine {
	if pl.err != nil {
		fn(pl.err)
	}
	return pl
}

func (pl *ProcessLine) Recover(fn PLRecoverFn) *ProcessLine {
	if pl.err != nil {
		fn(pl.err)
		pl.err = nil
	}
	return pl
}

func (pl *ProcessLine) Finally(fn PLFinallyFn) *ProcessLine {
	fn()
	return pl
}

func (pl *ProcessLine) Result() (msgCode int, err error) {
	return pl.msgCode, pl.err
}

func New() *ProcessLine {
	pl := ProcessLine{}
	return &pl
}
