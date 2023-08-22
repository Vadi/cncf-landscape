package types

// ProccessedLandscapeList is the struct that is used to store the processed landscape list
type ProccessedLandscapeList struct {
	Landscape []struct {
		Category      interface{} `yaml:"category"`
		Name          string      `yaml:"name"`
		Subcategories []struct {
			Subcategory interface{} `yaml:"subcategory"`
			Name        string      `yaml:"name"`
			Items       []struct {
				Item           interface{} `yaml:"item"`
				Name           string      `yaml:"name"`
				HomepageURL    string      `yaml:"homepage_url"`
				RepoURL        string      `yaml:"repo_url,omitempty"`
				Logo           string      `yaml:"logo"`
				Twitter        string      `yaml:"twitter,omitempty"`
				Crunchbase     string      `yaml:"crunchbase"`
				CrunchbaseData struct {
					Name            string        `yaml:"name"`
					Description     string        `yaml:"description"`
					NumEmployeesMin int           `yaml:"num_employees_min"`
					NumEmployeesMax int           `yaml:"num_employees_max"`
					Homepage        string        `yaml:"homepage"`
					City            string        `yaml:"city"`
					Region          string        `yaml:"region"`
					Country         string        `yaml:"country"`
					Twitter         string        `yaml:"twitter"`
					Linkedin        string        `yaml:"linkedin"`
					Acquisitions    []interface{} `yaml:"acquisitions"`
					Parents         []interface{} `yaml:"parents"`
					StockExchange   interface{}   `yaml:"stockExchange"`
					CompanyType     string        `yaml:"company_type"`
					Industries      []interface{} `yaml:"industries"`
				} `yaml:"crunchbase_data,omitempty"`
				GithubData struct {
					Languages []struct {
						Name  string `yaml:"name"`
						Value int    `yaml:"value"`
						Color string `yaml:"color"`
					} `yaml:"languages"`
					Contributions     string `yaml:"contributions"`
					FirstWeek         string `yaml:"firstWeek"`
					Stars             int    `yaml:"stars"`
					License           string `yaml:"license"`
					Description       string `yaml:"description"`
					LatestCommitDate  string `yaml:"latest_commit_date"`
					LatestCommitLink  string `yaml:"latest_commit_link"`
					ReleaseDate       string `yaml:"release_date"`
					ReleaseLink       string `yaml:"release_link"`
					ContributorsCount int    `yaml:"contributors_count"`
					ContributorsLink  string `yaml:"contributors_link"`
				} `yaml:"github_data,omitempty"`
				Repos []struct {
					URL   string `yaml:"url"`
					Stars int    `yaml:"stars"`
				} `yaml:"repos,omitempty"`
				GithubStartCommitData struct {
					StartCommitLink string `yaml:"start_commit_link"`
					StartDate       string `yaml:"start_date"`
				} `yaml:"github_start_commit_data,omitempty"`
				ImageData struct {
					FileName string `yaml:"fileName"`
					Hash     string `yaml:"hash"`
				} `yaml:"image_data"`
				BestPracticeData struct {
					Badge      string      `yaml:"badge"`
					Percentage interface{} `yaml:"percentage"`
				} `yaml:"best_practice_data"`
				TwitterData struct {
					LatestTweetDate string `yaml:"latest_tweet_date"`
				} `yaml:"twitter_data,omitempty"`
				Project string `yaml:"project,omitempty"`
				Extra   struct {
					Accepted         string `yaml:"accepted"`
					AnnualReviewDate string `yaml:"annual_review_date"`
					AnnualReviewURL  string `yaml:"annual_review_url"`
					DevStatsURL      string `yaml:"dev_stats_url"`
					BlogURL          string `yaml:"blog_url"`
					ArtworkURL       string `yaml:"artwork_url"`
					SlackURL         string `yaml:"slack_url"`
					ChatChannel      string `yaml:"chat_channel"`
					ClomonitorName   string `yaml:"clomonitor_name"`
					ClomonitorSvg    string `yaml:"clomonitor_svg"`
				} `yaml:"extra,omitempty"`
			} `yaml:"items"`
		} `yaml:"subcategories"`
	} `yaml:"landscape"`
	TwitterOptions struct {
		Count   int    `yaml:"count"`
		SinceID string `yaml:"since_id"`
	} `yaml:"twitter_options"`
	UpdatedAt string `yaml:"updated_at"`
}
