package models

type GitQuery struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
}

type GitResponse struct {
	Data struct {
		User struct {
			Login                   string `json:"login"`
			Typename                string `json:"__typename"`
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
			Followers struct {
				TotalCount int `json:"totalCount"`
			} `json:"followers"`
			Repositories struct {
				Nodes []struct {
					Name           string `json:"name"`
					StargazerCount int    `json:"stargazerCount"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type Profile struct {
	userID          uint
	jobID           uint
	username        string
	followers       uint
	contributions   uint
	mostPopularRepo string
	repoStars       uint
}
