package reddit

import (
	"strings"
	"net/url"
	"strconv"
)

func (r Reddit) GetListing(options ListingOptions) (posts *ListingResponse, err error) {

	err = options.Validate()
	if err != nil {
		return posts, err
	}

	q := url.Values{}

	if options.After != "" {
		q.Set("after", options.After)
	}
	if options.Before != "" {
		q.Set("before", options.Before)
	}
	if options.Count > 0 {
		q.Set("count", strconv.Itoa(options.Count))
	}
	if options.Limit > 0 {
		q.Set("limit", strconv.Itoa(options.Limit))
	}
	if options.Show {
		q.Set("show", "all")
	}
	if options.Detail {
		q.Set("sr_detail", "")
	}
	if options.Time == SortTop || options.Time == SortControversial {
		q.Set("t", string(options.Time))
	}

	var u string
	if options.Reddit != "" {
		u = u + "r/" + options.Reddit
	}
	if options.Sort != "" {
		u = u + "/" + string(options.Sort)
	}

	posts = new(ListingResponse)
	err = r.fetchGet(u, q, posts)
	if err != nil {
		return posts, err
	}

	return posts, err
}

type ListingOptions struct {
	After    string
	Before   string
	Count    int
	Limit    int
	Show     bool
	Detail   bool
	Location ListingLocation // Hot only
	Time     ListingTime     // Top & Controversial only
	Sort     ListingSort
	Reddit   string
}

func (l ListingOptions) Validate() error {

	return nil
}

type ListingResponse struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string        `json:"modhash"`
		Dist     int           `json:"dist"`
		Children []ListingPost `json:"children"`
		After    string        `json:"after"`
		Before   interface{}   `json:"before"`
	} `json:"data"`
}

type ListingPost struct {
	Kind string          `json:"kind"`
	Data ListingPostData `json:"data"`
}

type ListingPostData struct {
	ApprovedAtUtc  interface{}   `json:"approved_at_utc"`
	Subreddit      string        `json:"subreddit"`
	Selftext       string        `json:"selftext"`
	UserReports    []interface{} `json:"user_reports"`
	IsSaved        bool          `json:"saved"`
	ModReasonTitle interface{}   `json:"mod_reason_title"`
	Gilded         int           `json:"gilded"`
	IsClicked      bool          `json:"clicked"`
	Title          string        `json:"title"`
	LinkFlairRichtext []struct {
		E string `json:"e"`
		T string `json:"t"`
	} `json:"link_flair_richtext"`
	SubredditNamePrefixed      string      `json:"subreddit_name_prefixed"`
	IsHidden                   bool        `json:"hidden"`
	Pwls                       int         `json:"pwls"`
	LinkFlairCSSClass          string      `json:"link_flair_css_class"`
	Downs                      int         `json:"downs"`
	ThumbnailHeight            int         `json:"thumbnail_height"`
	ParentWhitelistStatus      string      `json:"parent_whitelist_status"`
	HideScore                  bool        `json:"hide_score"`
	Name                       string      `json:"name"`
	Quarantine                 bool        `json:"quarantine"`
	LinkFlairTextColor         string      `json:"link_flair_text_color"`
	AuthorFlairBackgroundColor interface{} `json:"author_flair_background_color"`
	SubredditType              string      `json:"subreddit_type"`
	Ups                        int         `json:"ups"`
	Domain                     string      `json:"domain"`
	MediaEmbed struct {
	} `json:"media_embed"`
	ThumbnailWidth        int         `json:"thumbnail_width"`
	AuthorFlairTemplateID interface{} `json:"author_flair_template_id"`
	IsOriginalContent     bool        `json:"is_original_content"`
	SecureMedia           interface{} `json:"secure_media"`
	IsRedditMediaDomain   bool        `json:"is_reddit_media_domain"`
	Category              interface{} `json:"category"`
	SecureMediaEmbed struct {
	} `json:"secure_media_embed"`
	LinkFlairText string      `json:"link_flair_text"`
	CanModPost    bool        `json:"can_mod_post"`
	Score         int         `json:"score"`
	ApprovedBy    interface{} `json:"approved_by"`
	Thumbnail     string      `json:"thumbnail"`
	//Edited              bool          `json:"edited"` // Timestamp or false
	AuthorFlairCSSClass string        `json:"author_flair_css_class"`
	AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
	PostHint            string        `json:"post_hint"`
	IsSelf              bool          `json:"is_self"`
	ModNote             interface{}   `json:"mod_note"`
	Created             float64       `json:"created"`
	LinkFlairType       string        `json:"link_flair_type"`
	Wls                 int           `json:"wls"`
	PostCategories      interface{}   `json:"post_categories"`
	BannedBy            interface{}   `json:"banned_by"`
	AuthorFlairType     string        `json:"author_flair_type"`
	ContestMode         bool          `json:"contest_mode"`
	SelftextHTML        interface{}   `json:"selftext_html"`
	Likes               interface{}   `json:"likes"`
	SuggestedSort       interface{}   `json:"suggested_sort"`
	BannedAtUtc         interface{}   `json:"banned_at_utc"`
	ViewCount           interface{}   `json:"view_count"`
	Archived            bool          `json:"archived"`
	NoFollow            bool          `json:"no_follow"`
	IsCrosspostable     bool          `json:"is_crosspostable"`
	Pinned              bool          `json:"pinned"`
	IsOver18            bool          `json:"over_18"`
	Preview struct {
		Images []struct {
			Source struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"source"`
			Resolutions []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"resolutions"`
			Variants struct {
			} `json:"variants"`
			ID string `json:"id"`
		} `json:"images"`
		Enabled bool `json:"enabled"`
	} `json:"preview"`
	CanGild              bool          `json:"can_gild"`
	IsSpoiler            bool          `json:"spoiler"`
	Locked               bool          `json:"locked"`
	AuthorFlairText      string        `json:"author_flair_text"`
	RteMode              string        `json:"rte_mode"`
	IsVisited            bool          `json:"visited"`
	NumReports           interface{}   `json:"num_reports"`
	Distinguished        interface{}   `json:"distinguished"`
	SubredditID          string        `json:"subreddit_id"`
	ModReasonBy          interface{}   `json:"mod_reason_by"`
	RemovalReason        interface{}   `json:"removal_reason"`
	ID                   string        `json:"id"`
	ReportReasons        interface{}   `json:"report_reasons"`
	Author               string        `json:"author"`
	NumCrossposts        int           `json:"num_crossposts"`
	NumComments          int           `json:"num_comments"`
	SendReplies          bool          `json:"send_replies"`
	ModReports           []interface{} `json:"mod_reports"`
	AuthorFlairTextColor interface{}   `json:"author_flair_text_color"`
	Permalink            string        `json:"permalink"`
	WhitelistStatus      string        `json:"whitelist_status"`
	Stickied             bool          `json:"stickied"`
	URL                  string        `json:"url"`
	SubredditSubscribers int           `json:"subreddit_subscribers"`
	CreatedUtc           float64       `json:"created_utc"`
	Media                interface{}   `json:"media"`
	IsVideo              bool          `json:"is_video"`
}

func (d ListingPostData) IsImage() bool {
	return strings.HasSuffix(d.URL, ".jpg") || strings.HasSuffix(d.URL, ".jpeg") || strings.HasSuffix(d.URL, ".png") || strings.HasSuffix(d.URL, ".gif")
}

func (r Reddit) Save(id string, category string) (err error) {

	q := url.Values{}
	q.Set("id", id)
	q.Set("category", category)

	err = r.fetchPost("api/save", q, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r Reddit) Unsave(id string) (err error) {

	q := url.Values{}
	q.Set("id", id)

	err = r.fetchPost("api/unsave", q, nil)
	if err != nil {
		return err
	}

	return err
}
