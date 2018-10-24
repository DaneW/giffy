package giphy

// Rank represents the seralized verion of database Rank without userIDs
type Rank struct {
	UpVotes   int `json:"up_votes"`
	DownVotes int `json:"down_votes"`
}

// Search represents a search response from the Giphy API
type Search struct {
	Data       []Data     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GIF represents a ID response from the Giphy API
type GIF struct {
	Data Data
}

// Data contains all the fields in a data response from the Giphy API
type Data struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	EmbedURL         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Caption          string `json:"caption"`
	ContentURL       string `json:"content_url"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`
	Images           Images `json:"images"`
	Rank             Rank   `json:"rank"`
}

// Images represents all the different types of images
type Images struct {
	Downsized      Image `json:"downsized"`
	DownsizedStill Image `json:"downsized_still"`
	Original       Image `json:"original"`
	OriginalStill  Image `json:"original_still"`
}

// Image represents an image
type Image struct {
	URL    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Size   string `json:"size,omitempty"`
	Frames string `json:"frames,omitempty"`
}

// Pagination represents the pagination section in a Giphy API response
type Pagination struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Offset     int `json:"offset"`
}
