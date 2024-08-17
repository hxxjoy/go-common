package chttp

import "fmt"

type Result struct {
	Code int8
	Msg  string
	Data any
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Test() {
	fmt.Println("Test")
}

func (r *Result) Fail(err error, data any, code int8) *Result {
	return &Result{
		Code: code,
		Msg:  err.Error(),
		Data: data,
	}
}

func (r *Result) Success(data any) *Result {
	fmt.Println("=====================", data)
	return &Result{
		Code: 2,
		Msg:  "3333333333",
		Data: data,
	}
}

func (r *Result) Deal(err error, data any) *Result {
	fmt.Println(err, "------------------------------------------------------------")
	if err != nil {
		fmt.Println("+++++++++++++++++++++++++++")
		return r.Fail(err, data, 0)
	}
	return r.Success(data)
}
