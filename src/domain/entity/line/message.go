package line

type NotifyUserRequest struct {
	UserId   string `json:"user_id"`
	Message     string `json:"message"`
}