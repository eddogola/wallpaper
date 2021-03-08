package wallpapr

// Logger provides functions implemented by any logging library
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
