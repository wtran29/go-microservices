package main

import (
	"context"

	"github.com/wtran29/go-microservices/logger/data"
	"github.com/wtran29/go-microservices/logger/logs"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer             // ensures backwards compatibility
	Models                             data.Models // methods to write to mongo
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	//return response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}
