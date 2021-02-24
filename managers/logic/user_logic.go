package logic

type User struct {
	userId   *string
	username *string
	channels []*Channel
	threads  []*Thread // Change later to thread type
}
