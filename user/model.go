package user

type User struct {
	Id          string              `json:"id"`
	Phase       phase               `json:"phase"`
	HashTag     *string             `json:"hash_tag"`
	Completed   map[quest]bool      `json:"completed"`
	TestPhase   int                 `json:"test_phase"`
	Answers     map[AnswerValue]int `json:"answers"`
	ProjectText string              `json:"project_text"`
}

func NewUser(id string, phase phase, hashTag *string) *User {
	return &User{
		Id:        id,
		Phase:     phase,
		HashTag:   hashTag,
		Completed: make(map[quest]bool),
		TestPhase: 0,
		Answers:   make(map[AnswerValue]int),
	}
}

type phase string

const (
	LoggingPhase phase = "LOGGING_PHASE"
	MainPhase    phase = "MAIN_PHASE"
	TestPhase    phase = "TEST_PHASE"
	ProjectPhase phase = "PROJECT_PHASE"
)

type quest string

const (
	Project quest = "project"
	Test    quest = "test"
)

type AnswerValue string

const (
	SMM     AnswerValue = "smm"
	Back    AnswerValue = "back"
	Mobile  AnswerValue = "mobile"
	Design  AnswerValue = "design"
	Manager AnswerValue = "manager"
	Front   AnswerValue = "front"
)
