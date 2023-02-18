package adapters

type LogAdapter interface {
	Log(info string, message string)
}
