package state

const (
	Port string = ":8019"
	Path string = "/state"
)

type Message struct {
	Head string `json:"head"`
	Body string `json:"body"`
}

func (self *Message) String() string {
	return self.Head + " with " + self.Body
}

var maxId int = 0
var channelBufSize int = 19
