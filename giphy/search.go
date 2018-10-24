package giphy

import (
	"fmt"
)

// Search returns a search response from the Giphy API
func (c *Client) Search(args string) (Search, error) {
	path := fmt.Sprintf("/gifs/search?limit=%v&q=%s", c.limit, args)
	req, err := c.NewRequest(path)
	if err != nil {
		return Search{}, err
	}

	var search Search
	if _, err = c.Do(req, &search); err != nil {
		return Search{}, err
	}

	return search, nil
}

// SearchOffset returns a search response from the Giphy API with specified offset
func (c *Client) SearchOffset(args string, offset string) (Search, error) {
	path := fmt.Sprintf("/gifs/search?limit=%v&q=%s&offset=%s", c.limit, args, offset)
	req, err := c.NewRequest(path)
	if err != nil {
		return Search{}, err
	}

	var search Search
	if _, err = c.Do(req, &search); err != nil {
		return Search{}, err
	}

	return search, nil
}
