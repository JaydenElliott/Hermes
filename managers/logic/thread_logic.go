package logic

type Thread struct {
	threadID *string
	users    map[*User]bool
	channel  *Channel
}
