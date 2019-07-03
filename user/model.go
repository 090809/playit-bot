package user

type User struct {
	Id 			string 				`json:"id"`
	Phase 		Phase				`json:"phase"`
	HashTag 	*string  			`json:"hash_tag"`
	Completed 	*map[string]Quest  	`json:"completed"`
}

func NewUser(id string, phase Phase, hashTag *string) *User {
	return &User{Id: id, Phase: phase, HashTag: hashTag}
}

type Phase string
const (
	LoggingPhase Phase = "LOGGING_PHASE"
	MainPhase Phase = "MAIN_PHASE"
	TestPhase Phase = "TEST_PHASE"
)

type Quest string
const (
	Quest1 Quest = "quest1"
	Quest2 Quest = "quest2"
	Quest3 Quest = "quest3"
	Quest4 Quest = "quest4"
	Quest5 Quest = "quest5"
)

