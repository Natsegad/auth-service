package user

import "github.com/google/uuid"

func GenUserId() uint64 {
	return uint64(uuid.New().ID())
}
