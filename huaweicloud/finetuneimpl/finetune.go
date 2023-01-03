package finetuneimpl

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/xihe-finetune/domain"
)

type finetuneImpl struct {
	cfg    *ModelartsConfig
	cli    utils.HttpClient
	token  string
	expire time.Time
}

func NewFinetune(cfg *ModelartsConfig) (domain.Finetune, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	return &finetuneImpl{cfg: cfg, cli: utils.NewHttpClient(3)}, nil
}

func (f *finetuneImpl) baseUrl(s string) string {
	return fmt.Sprintf("%s/%s", f.cfg.Endpoint, s)
}

func (f *finetuneImpl) checkTokenExpire() bool {
	if len(f.token) > 0 && time.Now().Before(f.expire) {
		return true
	}

	return false
}

// Token get Authorization token
func (f *finetuneImpl) Token() (t string, err error) {
	if f.checkTokenExpire() {
		t = f.token

		return
	}
	var req *http.Request
	body := fmt.Sprintf(`{"username":"%s", "password":"%s"}`, f.cfg.Username, f.cfg.Password)

	req, err = http.NewRequest("POST", f.baseUrl(tokenRoutePath), strings.NewReader(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	var res = new(domain.TokenInfo)
	//type TokenInfo struct {
	//	Duration int64  `json:"duration"`
	//	Token    string `json:"token"`
	//	Msg      string `json:"msg"`
	//	Status   string `json:"status"`
	//}
	_, err = f.cli.ForwardTo(req, res)
	if err != nil {
		return "", err
	}

	if res.Status != "200" {
		err = errors.New(res.Msg)
		return
	}

	t = res.Token
	err = nil

	f.token = t
	f.expire = time.Now().Add(time.Duration(res.Duration) * time.Second)

	return
}

// CreateFinetune create finetune job
func (f *finetuneImpl) CreateFinetune(options *domain.CreateFinetuneOptions) (jobId string, err error) {
	var bys []byte
	var req *http.Request
	bys, err = utils.JsonMarshal(options)
	if err != nil {
		return
	}

	var token string
	token, err = f.Token()
	if err != nil {
		return
	}

	req, err = http.NewRequest("POST", f.baseUrl(CreateFinetuneV1), bytes.NewReader(bys))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JWT "+token)

	var res = new(domain.CreateFinetuneInfo)
	//type CreateFinetuneInfo struct {
	//	Status int    `json:"status"`
	//	Msg    string `json:"msg"`
	//	JobId  string `json:"job_id"`
	//}
	_, err = f.cli.ForwardTo(req, res)
	if err != nil {
		return
	}

	if res.Status == -1 {
		err = errors.New(res.Msg)
		return
	}

	jobId = res.JobId

	return
}

// GetFinetune get finetune job
func (f *finetuneImpl) GetFinetune(jobId string) (info domain.FinetuneData, err error) {
	var req *http.Request
	url := strings.ReplaceAll(f.baseUrl(GetFinetune), ":jobId", jobId)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	var token string
	token, err = f.Token()
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "JWT "+token)

	var res = new(domain.GetFinetuneInfo)
	//type GetFinetuneInfo struct {
	//	Status int          `json:"status"`
	//	Msg    string       `json:"msg"`
	//	Data   FinetuneData `json:"data"`
	//}
	//
	//type FinetuneData struct {
	//	TaskName   string `json:"task_name"`
	//	Framework  string `json:"framework"`
	//	Phase      string `json:"phase"`
	//	TaskType   string `json:"task_type"`
	//	Runtime    int    `json:"runtime"`
	//	CreatedAt  string `json:"created_at"`
	//	EngineName string `json:"engine_name"`
	//}

	_, err = f.cli.ForwardTo(req, res)
	if err != nil {
		return
	}

	if res.Status == -1 {
		err = fmt.Errorf(res.Msg)
		return
	}

	info = res.Data

	return
}

// DeleteFinetune delete finetune job
func (f *finetuneImpl) DeleteFinetune(jobId string) (err error) {
	var req *http.Request
	url := strings.ReplaceAll(f.baseUrl(GetFinetune), ":jobId", jobId)

	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}

	var token string
	token, err = f.Token()
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "JWT "+token)

	var res = struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	}{}

	_, err = f.cli.ForwardTo(req, &res)
	if err != nil {
		return
	}

	if res.Status == -1 {
		err = errors.New(res.Msg)
		return
	}

	return
}

func (f *finetuneImpl) TerminateFinetune(jobId string) (err error) {
	var req *http.Request
	url := strings.ReplaceAll(f.baseUrl(GetFinetune), ":jobId", jobId)

	req, err = http.NewRequest("PUT", url, nil)
	if err != nil {
		return
	}

	var token string
	token, err = f.Token()
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "JWT "+token)

	var res = struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	}{}

	_, err = f.cli.ForwardTo(req, &res)
	if err != nil {
		return
	}

	if res.Status == -1 {
		err = errors.New(res.Msg)
		return
	}

	return
}

func (f *finetuneImpl) FinetuneLog(jobId string) (content string, err error) {
	var req *http.Request
	url := strings.ReplaceAll(f.baseUrl(FinetuneJobLog), ":jobId", jobId)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	var token string
	token, err = f.Token()
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "JWT "+token)

	var res = new(domain.FinetuneLogInfo)
	//type FinetuneLogInfo struct {
	//	Status int    `json:"status"`
	//	Msg    string `json:"msg"`
	//	ObsUrl string `json:"obs_url"`
	//}

	_, err = f.cli.ForwardTo(req, res)
	if err != nil {
		return
	}

	if res.Status == -1 {
		err = errors.New(res.Msg)
		return
	}

	content = res.ObsUrl

	return
}
