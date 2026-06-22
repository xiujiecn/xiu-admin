package model

type InfoModel struct {
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Version     string                 `json:"version"`
	Build       string                 `json:"build"`
	Commit      string                 `json:"commit"`
	Data        map[string]interface{} `json:"data"`
}

func NewInfoModel() *InfoModel {
	return &InfoModel{
		Data: make(map[string]interface{}),
	}
}
