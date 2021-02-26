package logic

/*
 Hosts the parameter structs for functions utilised by 2 or more files in the logic module
*/

// Parameters for getUsers function (see channel_logic.go and thread_logic.go)
type GetUsersParams_ struct {
	ReturnType string // [userId, username
}
