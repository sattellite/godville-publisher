package publisher

type Publisher interface {
	SendMessage(string) error
	UpdateStatus(string) error
}
