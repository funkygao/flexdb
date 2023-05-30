package mysql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/funkygao/log4go"
	ormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func newLogger(config ormlog.Config) ormlog.Interface {
	return &logger{
		Config:       config,
		traceStr:     "%s [%.3fms] [rows:%v] %s",
		traceWarnStr: "%s %s [%.3fms] [rows:%v] %s",
		traceErrStr:  "%s %s [%.3fms] [rows:%v] %s",
	}
}

type logger struct {
	ormlog.Config
	traceStr, traceErrStr, traceWarnStr string
}

func (l *logger) LogMode(level ormlog.LogLevel) ormlog.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l logger) Info(ctx context.Context, format string, data ...interface{}) {
	log4go.Info(strings.TrimSuffix(format, "\n"), data...)
}

func (l logger) Warn(ctx context.Context, format string, data ...interface{}) {
	log4go.Warn(strings.TrimSuffix(format, "\n"), data...)
}

func (l logger) Error(ctx context.Context, format string, data ...interface{}) {
	log4go.Error(strings.TrimSuffix(format, "\n"), data...)
}

func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil:
			log4go.Error(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)

		case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			log4go.Warn(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)

		default:
			log4go.Info(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
