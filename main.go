package main

import (
	"github.com/nuclio/nuclio-sdk-go"
)

func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	context.Logger.Info("This is an unstrucured %s", "log")

	evt := &nuclio.MemoryEvent{Body: []byte("body")}
	resp, err := context.Platform.CallFunction("test_nuclio2", evt)

	if err != nil {
		return 0, err
	}

	return nuclio.Response{
		StatusCode:  200,
		ContentType: "application/text",
		Body:        resp.Body,
	}, nil
}
