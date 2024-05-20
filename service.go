package main

import (
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) Service {
	return Service{
		Repository: repository,
	}
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()
	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}
