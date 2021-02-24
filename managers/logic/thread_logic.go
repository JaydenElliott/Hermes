package logic

type Thread struct {
	threadID *string
	users    []*User
	channel  *Channel
}
