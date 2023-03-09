package forum

type User struct {
	Username        string
	Email           string
	CreationDate    string
	BirthDate       string
	Uuid            string
	Admin           int
	TopicsCreated   []string
	MessagesSend    []MessageSend `MessageSend`
	UuidOfTopics    []string
	PrivateMessages []PrivateMessage `PrivateMessage`
}

type PrivateMessage struct {
	PrivateMessage        string
	PrivateMessage2ndUser string
}

type MessageSend struct {
	MessageSendByUser string
	TopicSentInName   string
}
