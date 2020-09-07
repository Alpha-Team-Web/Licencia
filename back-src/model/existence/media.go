package existence

type Follow struct {
	FollowerUsername    string `json:"follower-username" binding:"min=0,max=25" sql:",notnull"`
	FollowerFreelancer  bool   `json:"follower-freelancer"`
	FollowingUsername   string `json:"following-username" binding:"min=0,max=25" sql:",notnull"`
	FollowingFreelancer bool   `json:"following-freelancer"`
}

//Freelancer events
const (
	FEUpdateProfile    = "update profile"
	FETakeProject      = "take project"
	FEDeliverProject   = "deliver project"
	FEFollowFreelancer = "follow freelancer"
	FEFollowEmployer   = "follow employer"
)

//Employer events
const (
	EEUpdateProfile    = "update profile"
	EEAddProject       = "add project"
	EEAssignProject    = "assign project"
	EEExtendProject    = "extend project"
	EEFinishProject    = "finish project"
	EEFollowFreelancer = "follow freelancer"
	EEFollowEmployer   = "follow employer"
)

type Event struct {
	Username     string `json:"username" binding:"min=0,max=25" sql:",notnull"`
	IsFreelancer bool   `json:"is-freelancer"`
	EventMessage string `json:"event-message" binding:"min=0,max=25" sql:",notnull"`
	Id           string `json:"id" sql:",nopk"`
}
