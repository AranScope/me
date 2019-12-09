package types

type PatchPlugBody struct {
	State string `json:"state"`
}

type GetPlugResponse struct {
	State string `json:"state"`
	Model string `json:"model"`
}
