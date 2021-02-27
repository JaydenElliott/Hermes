package logic

/*
 Parameters for all functions relating to logic package
*/

// Update channel name
type UpdateName_ struct {
	UpdatedName string
}

// Parameters for CreateChannel function
type CreateChannel_ struct {
	ChannelName string
}

// Parameters for getUsers function
type GetUsersParams_ struct {
	ReturnType string // [userId, username
}
