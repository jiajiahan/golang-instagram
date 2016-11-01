package instagram

// Instagram User Object. Note that user objects are not always fully returned.
// Be sure to see the descriptions on the instagram documentation for any given endpoint.
type User struct {
	Id             string      `json:"id,omitempty"`
	Username       string      `json:"username,omitempty"`
	FullName       string      `json:"full_name,omitempty"`
	ProfilePicture string      `json:"profile_picture,omitempty"`
	Bio            string      `json:"bio,omitempty"`
	Website        string      `json:"website,omitempty"`
	Likes          int64       `json:"likes"`
	Comments       int64       `json:"comments"`
	FollowedAt     int64       `json:"followed_at,omitempty"`
	FollowingAt    int64       `json:"following_at,omitempty"`
	UnfollowedAt   int64       `json:"unfollowed_at,omitempty"`
	UnfollowingAt  int64       `json:"unfollowing_at,omitempty"`
	Counts         *UserCounts `json:"counts,omitempty"`
	Rank           int64       `json:"rank"`
}

// Instagram User Counts object. Returned on User objects
type UserCounts struct {
	Media      int64 `json:"media,omitempty"`
	Follows    int64 `json:"follows,omitempty"`
	FollowedBy int64 `json:"followed_by,omitempty"`
}

// Instagram Media object
type Media struct {
	Type         string         `json:"type,omitempty"`
	Id           string         `json:"id,omitempty"`
	UsersInPhoto []UserPosition `json:"users_in_photo,omitempty"`
	Filter       string         `json:"filter,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	Comments     *Comments      `json:"comments,omitempty"`
	Caption      *Caption       `json:"caption,omitempty"`
	Likes        *Likes         `json:"likes,omitempty"`
	Link         string         `json:"link,omitempty"`
	User         *User          `json:"user,omitempty"`
	CreatedTime  StringUnixTime `json:"created_time,omitempty"`
	Images       *Images        `json:"images,omitempty"`
	Videos       *Videos        `json:"videos,omitempty"`
	Location     *Location      `json:"location,omitempty"`
	UserHasLiked bool           `json:"user_has_liked,omitempty"`
	Attribution  *Attribution   `json:"attribution,omitempty"`
}

// A pair of user object and position
type UserPosition struct {
	User     *User     `json:"user,omitempty"`
	Position *Position `json:"position,omitempty"`
}

// A position in a media
type Position struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

// Instagram tag
type Tag struct {
	MediaCount int64  `json:"media_count,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Comments struct {
	Count int64     `json:"count,omitempty"`
	Data  []Comment `json:"data,omitempty"`
}

type Comment struct {
	CreatedTime StringUnixTime `json:"created_time,omitempty"`
	Text        string         `json:"text,omitempty"`
	From        *User          `json:"from,omitempty"`
	Id          string         `json:"id,omitempty"`
}

type Caption Comment

type Likes struct {
	Count int64  `json:"count,omitempty"`
	Data  []User `json:"data,omitempty"`
}

type Images struct {
	LowResolution      *Image `json:"low_resolution,omitempty"`
	Thumbnail          *Image `json:"thumbnail,omitempty"`
	StandardResolution *Image `json:"standard_resolution,omitempty"`
}

type Image struct {
	Url    string `json:"url,omitempty"`
	Width  int64  `json:"width,omitempty"`
	Height int64  `json:"height,omitempty"`
}

type Videos struct {
	LowResolution      *Video `json:"low_resolution,omitempty"`
	StandardResolution *Video `json:"standard_resolution,omitempty"`
}

type Video Image

type Location struct {
	Id        LocationId `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Latitude  float64    `json:"latitude,omitempty"`
	Longitude float64    `json:"longitude,omitempty"`
}

type Relationship struct {
	IncomingStatus string `json:"incoming_status,omitempty"`
	OutgoingStatus string `json:"outgoing_status,omitempty"`
}

// If another app uploaded the media, then this is the place it is given. As of 11/2013, Hipstamic is the only allowed app
type Attribution struct {
	Website   string `json:"website,omitempty"`
	ItunesUrl string `json:"itunes_url,omitempty"`
	Name      string `json:"name,omitempty"`
}
