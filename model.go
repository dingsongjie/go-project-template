package main

import "time"

type requestModel struct {
	Paths []string
}

type ResponseModel struct {
	Paths []string
}

type FileUserInformation struct {
	OriginPath           string
	NewPath              string
	CreatedUserCode      string
	CreatedTime          time.Time
	LastModifiedUserCode string
	LastModifiedTime     time.Time
}
