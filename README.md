# plog

Deadly simple logger library for Go

## Functionality

`plog` supports Following three log levels.

* Info
  * Call `plog.Infof` function to log message as information
* Debug
  * Call `plog.Debugf` function to log message as debug
  * To enable debug, `plog.SetDebug(true)` must be called in advance.
* Error
  * Call `plog.Errorf` function to log message as error

`plog.SetPrinter(p Printer)` function is for control output format/destination of logging.

## LICENSE

MIT

## Author

Yosuke Akatsuka (@pankona)
* [Twitter](https://twitter.com/pankona)
* [GitHub](https://github.com/pankona)
* [Qiita](https://qiita.com/pankona)

