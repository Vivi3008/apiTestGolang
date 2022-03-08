package logging

import (
	"context"
	"time"

	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/sirupsen/logrus"
)

type LoggerKeyType int

const LoggerKey LoggerKeyType = iota

func NewLog(ctx context.Context, operation string) *logrus.Entry {
	log := FromContext(ctx).WithFields(logrus.Fields{
		"operation": operation,
		"time":      time.Now().UTC().Format(response.DateLayout),
	})
	return log
}

func FromContext(ctx context.Context) *logrus.Entry {
	if entry, ok := ctx.Value(LoggerKey).(*logrus.Entry); ok {
		return entry
	}
	return logrus.NewEntry(logrus.New())
}
