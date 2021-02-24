package logic

type Channel struct {
	channelID   *string
	channelName *string
	users       []*User
	threads     []*Thread
}
