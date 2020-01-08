package uuid

import (
	"encoding/hex"

	"github.com/gofrs/uuid"
)

//GetUUID generate and retrn uinique id
func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(uuid.Bytes())
}
