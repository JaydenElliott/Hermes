package logic

type User struct {
	userId   *string
	email    *string
	username *string
	channels []*Channel // chat channels this user is apart of  ** REVIEW (would a map be better)
}
