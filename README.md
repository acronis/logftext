# logf Appender and Encoder for colored text logs

[![Go Reference](https://pkg.go.dev/badge/github.com/acronis/logftext.svg)](https://pkg.go.dev/github.com/acronis/logftext)
[![Build Status](https://github.com/acronis/logftext/actions/workflows/go.yml/badge.svg)](https://github.com/acronis/logftext/actions/workflows/go.yml)

Package `logftext` provides `logf` Appender and Encoder for colored text logs.

## Example

The following example creates the new `logf` logger with `logftext` Encoder.

```go
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
```

The output example is:
![Alt text](https://user-images.githubusercontent.com/1574981/48273352-bed64f00-e451-11e8-9f5d-7ef1d76d222d.png)

## TODOs

* Handle terminals with a light backgrounds