package http

import "luwu/libs/error"

type Response struct {
	error.Error
	Data interface{} `json:"data"`
}

func NewResponse(data interface{}, err *error.Error) *Response {
	e := error.Error{}
	if err != nil {
		e = *err
	}
	return &Response{
		Error: e,
		Data:  data,
	}
}
