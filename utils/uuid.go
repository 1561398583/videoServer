package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUID()  (string, error){
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	idStr := id.String()
	ss := strings.Split(idStr, "-")
	r := ""
	for i := 0; i < len(ss); i++ {
		r += ss[i]
	}
	return r, nil
}
