package wallpaper

// User defines public fields Unsplash provides on a user
type User struct {
	ID                string
	UpdatedAt         string
	Username          string
	Name              string
	PortfolioURL      string
	Bio               string
	Location          string
	TotalLikes        int
	TotalPhotos       int
	TotalCollections  int
	InstagramUsername string
	TwitterUsername   string
	AcceptedTos       bool
	ProfileImage      struct {
		Small  string
		Medium string
		Large  string
	}
	Links struct {
		Self      string
		HTML      string
		Photos    string
		Likes     string
		Portfolio string
		Following string
		Followers string
	}
}

// Photo defines fields in a photo resource
type Photo struct {
	ID                     string
	CreatedAt              string
	UpdatedAt              string
	Width                  int
	Height                 int
	Color                  string
	BlurHash               string
	Likes                  int
	LikesByUser            bool
	Description            string
	User                   User
	CurrentUserCollections []Collection
	Exif                   struct {
		Make         string
		Model        string
		ExposureTime string
		Aperture     string
		FocalLength  string
		ISO          int
	}
	Location struct {
		City     string
		Country  string
		Position struct {
			Latitude  float64
			Longitude float64
		}
	}
	Tags []Tag
	URLs struct {
		Raw     string
		Full    string
		Regular string
		Small   string
		Thumb   string
	}
	Links struct {
		Self             string
		HTML             string
		Download         string
		DownloadLocation string
	}
}

// Collection defines fields in a collection resource
type Collection struct {
	ID              int
	Title           string
	Description     string
	PublishedAt     string
	LastCollectedAt string
	UpdatedAt       string
	TotalPhotos     int
	Private         bool
	ShareKey        string
	CoverPhoto      Photo
	User            User
}

// Tag defines fields in a photo's tag
type Tag struct {
	Title string
}

// Topic defines fields in an Unsplash topic
type Topic struct {
	ID          string
	Slug        string
	Title       string
	Description string
	PublishedAt string
	UpdatedAt   string
	StartsAt    string
	EndsAt      string
	Featured    bool
	TotalPhotos int
	Links       struct {
		Self   string
		HTML   string
		Photos string
	}
	Status string
	Owners []User
	CurrentUserContributions []Photo
	TotalCurrentUserSubmissions int
	CoverPhoto Photo
	PreviewPhotos []Photo
}
