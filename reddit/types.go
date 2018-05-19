package reddit

type AuthScope string

const (
	ScopeAccount          AuthScope = "account"          // Update preferences and related account information. Will not have access to your email or password.
	ScopeCreddits                   = "creddits"         // Spend my reddit gold creddits on giving gold to other users.
	ScopeEdit                       = "edit"             // Edit and delete my comments and submissions.
	ScopeFlair                      = "flair"            // Select my subreddit flair. Change link flair on my submissions.
	ScopeHistory                    = "history"          // Access my voting history and comments or submissions I've saved or hidden.
	ScopeIdentity                   = "identity"         // Access my reddit username and signup date.
	ScopeLivemanage                 = "livemanage"       // Manage settings and contributors of live threads I contribute to.
	ScopeModconfig                  = "modconfig"        // Manage the configuration, sidebar, and CSS of subreddits I moderate.
	ScopeModcontributors            = "modcontributors"  // Add/remove users to approved submitter lists and ban/unban or mute/unmute users from subreddits I moderate.
	ScopeModflair                   = "modflair"         // Manage and assign flair in subreddits I moderate.
	ScopeModlog                     = "modlog"           // Access the moderation log in subreddits I moderate.
	ScopeModmail                    = "modmail"          // Access and manage modmail via mod.reddit.com.
	ScopeModothers                  = "modothers"        // Invite or remove other moderators from subreddits I moderate.
	ScopeModposts                   = "modposts"         // Approve, remove, mark nsfw, and distinguish content in subreddits I moderate.
	ScopeModself                    = "modself"          // Accept invitations to moderate a subreddit. Remove myself as a moderator or contributor of subreddits I moderate or contribute to.
	ScopeModtraffic                 = "modtraffic"       // Access traffic stats in subreddits I moderate.
	ScopeModwiki                    = "modwiki"          // Change editors and visibility of wiki pages in subreddits I moderate.
	ScopeMysubreddits               = "mysubreddits"     // Access the list of subreddits I moderate, contribute to, and subscribe to.
	ScopePrivatemessages            = "privatemessages"  // Access my inbox and send private messages to other users.
	ScopeRead                       = "read"             // Access posts and comments through my account.
	ScopeReport                     = "report"           // Report content for rules violations. Hide & show individual submissions.
	ScopeSave                       = "save"             // Save and unsave comments and submissions.
	ScopeStructuredstyles           = "structuredstyles" // Edit structured styles for a subreddit I moderate.
	ScopeSubmit                     = "submit"           // Submit links and comments from my account.
	ScopeSubscribe                  = "subscribe"        // Manage my subreddit subscriptions. Manage \"friends\" - users whose content I follow.
	ScopeVote                       = "vote"             // Submit and change my votes on comments and submissions.
	ScopeWikiedit                   = "wikiedit"         // Edit wiki pages on my behalf.
	ScopeWikiread                   = "wikiread"         // Read wiki pages through my account.
)

type ListingSort string

const (
	SortDefault       ListingSort = ""
	SortHot                       = "hot"
	SortNew                       = "new"
	SortRising                    = "rising"
	SortTop                       = "top"
	SortControversial             = "controversial"
)

type ListingTime string

const (
	TimeDefault ListingTime = ""
	TimeHour                = "hour"
	TimeDay                 = "day"
	TimeWeek                = "week"
	TimeMonth               = "month"
	TimeYear                = "year"
	TimeAllTime             = "all"
)

type ListingLocation string

const (
	GLOBAL ListingLocation = "GLOBAL"
	US                     = "US"
	AR                     = "AR"
	AU                     = "AU"
	BG                     = "BG"
	CA                     = "CA"
	CL                     = "CL"
	CO                     = "CO"
	HR                     = "HR"
	CZ                     = "CZ"
	FI                     = "FI"
	GR                     = "GR"
	HU                     = "HU"
	IS                     = "IS"
	IN                     = "IN"
	IE                     = "IE"
	JP                     = "JP"
	MY                     = "MY"
	MX                     = "MX"
	NZ                     = "NZ"
	PH                     = "PH"
	PL                     = "PL"
	PT                     = "PT"
	PR                     = "PR"
	RO                     = "RO"
	RS                     = "RS"
	SG                     = "SG"
	SE                     = "SE"
	TW                     = "TW"
	TH                     = "TH"
	TR                     = "TR"
	GB                     = "GB"
	USWA                   = "US_WA"
	USDE                   = "US_DE"
	USDC                   = "US_DC"
	USWI                   = "US_WI"
	USWV                   = "US_WV"
	USHI                   = "US_HI"
	USFL                   = "US_FL"
	USWY                   = "US_WY"
	USNH                   = "US_NH"
	USNJ                   = "US_NJ"
	USNM                   = "US_NM"
	USTX                   = "US_TX"
	USLA                   = "US_LA"
	USNC                   = "US_NC"
	USND                   = "US_ND"
	USNE                   = "US_NE"
	USTN                   = "US_TN"
	USNY                   = "US_NY"
	USPA                   = "US_PA"
	USCA                   = "US_CA"
	USNV                   = "US_NV"
	USVA                   = "US_VA"
	USCO                   = "US_CO"
	USAK                   = "US_AK"
	USAL                   = "US_AL"
	USAR                   = "US_AR"
	USVT                   = "US_VT"
	USIL                   = "US_IL"
	USGA                   = "US_GA"
	USIN                   = "US_IN"
	USIA                   = "US_IA"
	USOK                   = "US_OK"
	USAZ                   = "US_AZ"
	USID                   = "US_ID"
	USCT                   = "US_CT"
	USME                   = "US_ME"
	USMD                   = "US_MD"
	USMA                   = "US_MA"
	USOH                   = "US_OH"
	USUT                   = "US_UT"
	USMO                   = "US_MO"
	USMN                   = "US_MN"
	USMI                   = "US_MI"
	USRI                   = "US_RI"
	USKS                   = "US_KS"
	USMT                   = "US_MT"
	USMS                   = "US_MS"
	USSC                   = "US_SC"
	USKY                   = "US_KY"
	USOR                   = "US_OR"
	USSD                   = "US_SD"
)
