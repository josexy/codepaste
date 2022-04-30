package serializer

import "github.com/josexy/gocodepaste/model"

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	CreatedAt int64  `json:"created_at"`
}

func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
