package domain

import "encoding/base64"

type OEM string
type OEMPageURL string
type OEMPageResultUrl string

func GetOEMURL() string {
	OEMsURLEnc := "aHR0cHM6Ly93d3cuYXV0b3RyYWRlci5jby56YS9zZWFyY2gvbWFrZW1vZGVsc2F1dG9jb21wbGV0ZQ=="
	sDec, _ := base64.StdEncoding.DecodeString(OEMsURLEnc)
	return string(sDec)
}

func GetOEMPagesURL() string {
	OEMsURLEnc := "aHR0cHM6Ly93d3cuYXV0b3RyYWRlci5jby56YS9jYXJzLWZvci1zYWxlLw=="
	sDec, _ := base64.StdEncoding.DecodeString(OEMsURLEnc)
	return string(sDec)
}

func GetBaseURL() string {
	OEMsURLEnc := "aHR0cHM6Ly93d3cuYXV0b3RyYWRlci5jby56YQ=="
	sDec, _ := base64.StdEncoding.DecodeString(OEMsURLEnc)
	return string(sDec)
}
