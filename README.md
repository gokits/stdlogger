# stdlogger
Logger interface and wrappers for common implementations

# What this package provides
## Minimal leveled logger interface
A minimal leveled logger interface is provided. We can use this interface in other packages so the package would not be tied together with any specific logger implementation. 
```go
type LeveledLogger interface {
	Debugf(format string, v ...interface{})
	Debug(v ...interface{})
	Infof(format string, v ...interface{})
	Info(v ...interface{})
	Warnf(format string, v ...interface{})
	Warn(v ...interface{})
	Errorf(format string, v ...interface{})
	Error(v ...interface{})
}
```
## Wrappers of common logger implementations
This package alreay include some wrappers of logger implementations. And many other logger implementation already implemented `LeveledLogger` interface(`zap.Sugar` for example). If you want, you can implementation your own wrappers very easy.

example for [logrus](https://github.com/sirupsen/logrus)
```go
import "github.com/gokits/stdlogger/logrus"

//...

var logglobal LeveledLogger = logrus.FromGlobal()
var logentry LeveledLogger = logrus.FromEntry(...)
```

example for [zap](https://github.com/uber-go/zap)

```go
import "go.uber.org/zap"

//...

logger, _ := zap.NewProduction()
defer logger.Sync() // flushes buffer, if any
var sugar LeveledLogger = logger.Sugar()
```
