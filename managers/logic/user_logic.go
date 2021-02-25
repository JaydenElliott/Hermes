package logic

type User struct {
	userId   *string
	username *string
	channels map[*Channel]bool
	threads  map[*Thread]bool
}
