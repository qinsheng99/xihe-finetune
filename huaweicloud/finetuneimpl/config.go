package finetuneimpl

import (
	"fmt"
)

const (
	tokenRoutePath   = "foundation-model/token"
	CreateFinetuneV1 = "v1/foundation-model/finetune"
	GetFinetune      = "v1/foundation-model/finetune/:jobId"
	FinetuneJobLog   = "v1/foundation-model/finetune/:jobId/log"
)

type Config struct {
	Modelarts ModelartsConfig `json:"modelarts"   required:"true"`
}

type ModelartsConfig struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`

	// finetune endpoint
	Endpoint string `json:"endpoint" required:"true"`
}

func (m *ModelartsConfig) Validate() error {
	if len(m.Username) == 0 || len(m.Password) == 0 || len(m.Endpoint) == 0 {
		return fmt.Errorf("parameter is empty")
	}
	return nil
}
