package types

type GetConfigRequest struct {
	Path string `json:"path"`
}

type GetConfigResponse struct {
	Path string `json:"path"`
	Body string `json:"body"`
}

type PutConfigRequest struct {
	Path string `json:"path"`
	Body string `json:"body"`
}

type PutConfigResponse struct {
	Path string `json:"path"`
	Body string `json:"body"`
}
