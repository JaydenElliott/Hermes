package channel

type UpdateName_ struct {
	UpdatedName string
}

// Parameters for CreateChannel function
type Create_ struct {
	ChannelName string
}

// Parameters for getUsers function
type GetUsersParams_ struct {
	ReturnType string // [userId, username
}
