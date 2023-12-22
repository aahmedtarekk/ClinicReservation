package controllers

import (
	"crypto/rand"
	"fmt"
	"toolsProject/initializers"
	"toolsProject/models"
)

func GenerateUuid(id uint, _type string) string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return ""
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	returnUuid := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	initializers.ActivePool[returnUuid] = models.UserPair{ID: id, Type: _type}
	return returnUuid
}
