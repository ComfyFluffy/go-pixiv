package pixiv

import "time"

/*
  These models are built exactly following Pixiv's AppAPI.
  There may be deprecated and strange fields here.
*/

// Generated by https://quicktype.io

// Profile is embedded in RespUserDetail
type Profile struct {
	Webpage string `json:"webpage"`
	Gender  string `json:"gender"`

	// Format: 1999-04-10
	Birth string `json:"birth"`

	// Deprecated: Use Birth instead.
	// Format: 04-10
	BirthDay string `json:"birth_day"`
	// Deprecated: Use Birth instead.
	BirthYear int `json:"birth_year"`

	Region                     string `json:"region"`
	AddressID                  int    `json:"address_id"`
	CountryCode                string `json:"country_code"`
	Job                        string `json:"job"`
	JobID                      int    `json:"job_id"`
	TotalFollowUsers           int    `json:"total_follow_users"`
	TotalMypixivUsers          int    `json:"total_mypixiv_users"`
	TotalIllusts               int    `json:"total_illusts"`
	TotalManga                 int    `json:"total_manga"`
	TotalNovels                int    `json:"total_novels"`
	TotalIllustBookmarksPublic int    `json:"total_illust_bookmarks_public"`
	TotalIllustSeries          int    `json:"total_illust_series"`
	TotalNovelSeries           int    `json:"total_novel_series"`
	BackgroundImageURL         string `json:"background_image_url"`
	TwitterAccount             string `json:"twitter_account"`
	TwitterURL                 string `json:"twitter_url"`
	PawooURL                   string `json:"pawoo_url"`
	IsPremium                  bool   `json:"is_premium"`
	IsUsingCustomProfileImage  bool   `json:"is_using_custom_profile_image"`
}

// ProfilePublicity is embedded in RespUserDetail
//
// All fields here except Pawoo are all "private" or "public"
type ProfilePublicity struct {
	Gender    string `json:"gender"`
	Region    string `json:"region"`
	BirthDay  string `json:"birth_day"`
	BirthYear string `json:"birth_year"`
	Job       string `json:"job"`
	Pawoo     bool   `json:"pawoo"`
}

// User may be embedded in Illust, Novel, Comment
type User struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Account          string `json:"account"`
	ProfileImageURLs struct {
		Medium string `json:"medium"`
	} `json:"profile_image_urls"`
	Comment    string `json:"comment"`
	IsFollowed bool   `json:"is_followed"`
}

// Illust is embedded in RespIllusts
type Illust struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`

	// Deprecated: Only contains the image URLs of the first page.
	// Use MetaSinglePage or MetaPages instead.
	ImageURLs ImageURLs `json:"image_urls"`

	Caption        string    `json:"caption"`
	Restrict       int       `json:"restrict"`
	User           User      `json:"user"`
	Tags           []Tag     `json:"tags"`
	Tools          []string  `json:"tools"`
	CreateDate     time.Time `json:"create_date"`
	PageCount      int       `json:"page_count"`
	Width          int       `json:"width"`
	Height         int       `json:"height"`
	SanityLevel    int       `json:"sanity_level"`
	XRestrict      int       `json:"x_restrict"`
	Series         Series    `json:"series"`
	MetaSinglePage struct {
		OriginalImageURL string `json:"original_image_url,omitempty"`
	} `json:"meta_single_page"`
	MetaPages []struct {
		ImageURLs ImageURLs `json:"image_urls"`
	} `json:"meta_pages"`
	TotalView      int  `json:"total_view"`
	TotalBookmarks int  `json:"total_bookmarks"`
	IsBookmarked   bool `json:"is_bookmarked"`
	Visible        bool `json:"visible"`
	IsMuted        bool `json:"is_muted"`
}

// ImageURLs is embedded in Illust, MetaPage, Novel
type ImageURLs struct {
	SquareMedium string `json:"square_medium"`
	Medium       string `json:"medium"`
	Large        string `json:"large"`
	Original     string `json:"original,omitempty"`
}

// NovelMarker is embedded in RespNovelText
type NovelMarker struct {
	Page int `json:"page"`
}

// Novel is embedded in RespNovelText, RespNovels
type Novel struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Caption        string    `json:"caption"`
	Restrict       int       `json:"restrict"`
	XRestrict      int       `json:"x_restrict"`
	ImageURLs      ImageURLs `json:"image_urls"`
	CreateDate     time.Time `json:"create_date"`
	Tags           []Tag     `json:"tags"`
	PageCount      int       `json:"page_count"`
	TextLength     int       `json:"text_length"`
	User           User      `json:"user"`
	Series         Series    `json:"series"`
	IsBookmarked   bool      `json:"is_bookmarked"`
	TotalBookmarks int       `json:"total_bookmarks"`
	TotalView      int       `json:"total_view"`
	Visible        bool      `json:"visible"`
	TotalComments  int       `json:"total_comments"`
	IsMuted        bool      `json:"is_muted"`
	IsMypixivOnly  bool      `json:"is_mypixiv_only"`
	IsXRestricted  bool      `json:"is_x_restricted"`
}

// Series is embedded in Illust(where Type="manga"), Novel
type Series struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Tag is embedded in Illust, Novel
type Tag struct {
	Name                string `json:"name"`
	TranslatedName      string `json:"translated_name"`
	AddedByUploadedUser *bool  `json:"added_by_uploaded_user"`
}

// Comment is embedded in RespComments
type Comment struct {
	ID         int       `json:"id"`
	Comment    string    `json:"comment"`
	Date       time.Time `json:"date"`
	User       User      `json:"user"`
	HasReplies bool      `json:"has_replies"`
}

/*
type PrivacyPolicy struct {
	Version string `json:"version"`
	Message string `json:"message"`
	URL     string `json:"url"`
}
*/
