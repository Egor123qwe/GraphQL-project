package storage

import (
	"github.com/Egor123qwe/GraphQL-project/internal/storage/repsInterfaces"
)

type Storage interface {
	User() repsInterfaces.User
}
