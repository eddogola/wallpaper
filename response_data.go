package wallpapr

// User defines public fields Unsplash provides on a user
type User struct {
	ID                string `json:"id"`
	UpdatedAt         string `json:"updated_at"`
	Username          string `json:"username"`
	Name              string `json:"name"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	InstagramUsername string `json:"instagram_username"`
	TwitterUsername   string `json:"twitter_username"`
	PortfolioURL      string `json:"portfolio_url"`
	Bio               string `json:"bio"`
	Location          string `json:"location"`
	TotalLikes        int    `json:"total_likes"`
	TotalPhotos       int    `json:"total_photos"`
	TotalCollections  int    `json:"total_collections"`
	FolloweddByUser   bool   `json:"followed_by_user"`
	FollowersCount    int    `json:"followers_count"`
	FollowingCount    int    `json:"following_count"`
	Downloads         int    `json:"downloads"`
	AcceptedTos       bool   `json:"accepted_tos"`
	ProfileImage      struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"profile_image"`
	Badge struct {
		Title   string `json:"title"`
		Primary bool   `json:"primary"`
		Slug    string `json:"slug"`
		Link    string `json:"link"`
	} `json:"badge"`
	Links struct {
		Self      string `json:"self"`
		HTML      string `json:"html"`
		Photos    string `json:"photos"`
		Likes     string `json:"likes"`
		Portfolio string `json:"portfolio"`
		Following string `json:"following"`
		Followers string `json:"followers"`
	} `json:"links"`
}

// Photo defines fields in a photo resource
type Photo struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	Downloads      int    `json:"downloads"`
	BlurHash       string `json:"blur_hash"`
	Likes          int    `json:"likes"`
	LikedByUser    bool   `json:"liked_by_user"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Exif           struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposure_time"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focal_length"`
		ISO          int    `json:"iso"`
	} `json:"exif"`
	Location struct {
		City     string `json:"city"`
		Country  string `json:"country"`
		Position struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
	} `json:"location"`
	Tags                   []Tag        `json:"tags"`
	CurrentUserCollections []Collection `json:"current_user_collections"`
	URLs                   struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	User User `json:"user"`
}

// Collection defines fields in a collection resource
type Collection struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublishedAt     string `json:"published_at"`
	LastCollectedAt string `json:"last_collected_at"`
	UpdatedAt       string `json:"updated_at"`
	Featured        bool   `json:"featured"`
	TotalPhotos     int    `json:"total_photos"`
	Private         bool   `json:"private"`
	ShareKey        string `json:"share_key"`
	CoverPhoto      Photo  `json:"cover_photo"`
	User            User   `json:"user"`
	Links           struct {
		Self   string `json:"self"`
		HTML   string `json:"html"`
		Photos string `json:"photos"`
	} `json:"links"`
}

// Tag defines fields in a photo's tag
type Tag struct {
	Title string `json:"title"`
}

// Topic defines fields in an Unsplash topic
type Topic struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	UpdatedAt   string `json:"updated_at"`
	StartsAt    string `json:"starts_at"`
	EndsAt      string `json:"ends_at"`
	Featured    bool   `json:"featured"`
	TotalPhotos int    `json:"total_photos"`
	Links       struct {
		Self   string `json:"self"`
		HTML   string `json:"html"`
		Photos string `json:"photos"`
	} `json:"links"`
	Status                      string  `json:"status"`
	Owners                      []User  `json:"owners"`
	TopContributors             []User  `json:"top_contributors"`
	CurrentUserContributions    []Photo `json:"current_user_contributions"`
	TotalCurrentUserSubmissions int     `json:"total_current_user_submissions"`
	CoverPhoto                  Photo   `json:"cover_photo"`
	PreviewPhotos               []Photo `json:"preview_photos"`
}

// SearchResult defines the structure of the response gotten
// after searching for a picture
type SearchResult struct {
	Total string `json:"total"`
	TotalPages int `json:"total_pages"`
	Results []Photo `json:"results"`
}
