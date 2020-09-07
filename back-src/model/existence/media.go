package existence

type Follow struct {
	FollowerUsername    string `json:"follower-username" binding:"min=0,max=25" sql:",notnull"`
	FollowerFreelancer  bool   `json:"follower-freelancer"`
	FollowingUsername   string `json:"following-username" binding:"min=0,max=25" sql:",notnull"`
	FollowingFreelancer bool   `json:"following-freelancer"`
}
