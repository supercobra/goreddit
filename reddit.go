package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Stringer interface {
	String() string
}

// Item describes a Reddit item
type Item struct {
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}

type response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
		com = " (0 comments)"
	case 1:
		com = " (1 comment)"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

// Get fetches the most recent Items posted to the specifed subreddit
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "your bot 0.1")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}

	return items, nil
}

func foo() {
	items, err := Get("golang")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}

}
