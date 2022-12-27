package domain

type FinetuneType string

func (f FinetuneType) ToString() string {
	return string(f)
}

var (
	Caption FinetuneType = "opt-caption"
)

type FinetuneTaskType string

func (f FinetuneTaskType) ToString() string {
	return string(f)
}

var (
	FinetuneTask FinetuneTaskType = "finetune"
)

type FinetuneParameter string

func (f FinetuneParameter) ToString() string {
	return string(f)
}

var (
	// CaptionEpochs ,value >= 1
	CaptionEpochs FinetuneParameter = "epochs"

	// CaptionStartLearningRate ,value is float
	CaptionStartLearningRate FinetuneParameter = "start_learning_rate"

	// CaptionEndLearningRate ,value is float
	CaptionEndLearningRate FinetuneParameter = "end_learning_rate"
)
