package domain

type Finetune interface {
	Token() (string, error)

	CreateFinetune(*CreateFinetuneOptions) (string, error)
	GetFinetune(jobId string) (FinetuneData, error)
	DeleteFinetune(jobId string) (err error)
	TerminateFinetune(jobId string) (err error)
	FinetuneLog(jobId string) (content string, err error)
}

// TokenInfo 获取token
type TokenInfo struct {
	Duration int64  `json:"duration"`
	Token    string `json:"token"`
	Msg      string `json:"msg"`
	Status   string `json:"status"`
}

type CreateFinetuneOptions struct {
	User            string           `json:"user"`
	TaskName        string           `json:"task_name"`
	FoundationModel FinetuneType     `json:"foundation_model"`
	TaskType        FinetuneTaskType `json:"task_type"`
	Parameters      []Parameter      `json:"parameters,omitempty"`
}

type Parameter struct {
	Name FinetuneParameter `json:"name"`

	Value string `json:"value"`
}

// CreateFinetuneInfo 创建微调
type CreateFinetuneInfo struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	JobId  string `json:"job_id"`
}

// GetFinetuneInfo 获取微调
type GetFinetuneInfo struct {
	Status int          `json:"status"`
	Msg    string       `json:"msg"`
	Data   FinetuneData `json:"data"`
}

type FinetuneData struct {
	TaskName   string `json:"task_name"`
	Framework  string `json:"framework"`
	Phase      string `json:"phase"`
	TaskType   string `json:"task_type"`
	Runtime    int    `json:"runtime"`
	CreatedAt  string `json:"created_at"`
	EngineName string `json:"engine_name"`
}

// FinetuneLogInfo 获取日志
type FinetuneLogInfo struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	ObsUrl string `json:"obs_url"`
}
