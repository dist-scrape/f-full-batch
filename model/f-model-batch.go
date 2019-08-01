package model

type MessageSubModel struct {
	OEM   string `json:"oem"`
	Model string `json:"model"`
}

func CreateMessageSubModelFromMakeModelResponseOEM(i MakeModelResponseOEM) []MessageSubModel {
	results := make([]MessageSubModel, 0)
	for _, m := range i.Models {
		r := MessageSubModel{
			OEM:   i.Title,
			Model: m.Title,
		}
		results = append(results, r)
	}
	return results
}
