package bitbucket

type Addons struct {
	c *Client
}

func (addons *Addons) Delete() error {
	urlStr := addons.c.requestUrl("/addon")
	_, err := addons.c.executePaginated("DELETE", urlStr, "")

	return err
}