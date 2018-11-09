package main

import (
	"errors"
	"os"
	"runtime"

	"github.com/acronis/logf"
	"github.com/acronis/logftext"
)

func main() {
	// Create ChannelWriter with text Encoder.
	writer, writerClose := logf.NewChannelWriter(logf.ChannelWriterConfig{
		Appender: logftext.NewAppender(os.Stdout, logftext.EncoderConfig{}),
	})
	defer writerClose()

	// Create Logger with ChannelWriter.
	logger := logf.NewLogger(logf.LevelInfo, writer).WithCaller().WithName("main")

	logger.Info("got cpu info", logf.Int("count", runtime.NumCPU()))
	logger.Error("error example", logf.Error(errors.New("failed to do nothing")))
}
