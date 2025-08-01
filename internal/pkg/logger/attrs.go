package logger

import "log/slog"

func (l Logger) Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func (l Logger) Uint64(key string, val uint64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(val),
	}
}
