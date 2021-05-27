package matchstr
import "github.com/google/uuid"

func GenerateUUID() (uuid.UUID) {
	uuid_string:= uuid.New()
	return uuid_string
}
