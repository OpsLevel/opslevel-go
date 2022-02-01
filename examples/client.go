/*
Wrapping the client can be useful when you want to override default behavior, such as always setting
context or disallowing (to the best of Go's ability) access to specific receiver functions on
`opslevel.Client`. Here is an example of wrapping the client to always require context to be passed
while also maintaining the ability to pass in default options, with the ability to extend your own
set of options if need-be:
*/


type Client struct{
	*opslevel.Client

	do opsLevelDefaultOptions
	customOption string
}

func NewClient(ctx context.Context, apiToken string, options ...Option) *Client {
	var c Client

	for i := range options{
		options[i](&c)
	}
	c.Client = opslevel.NewClient(apiToken, c.do.opsLevelOptions(ctx)...)

	return &c, nil
}

type opsLevelDefaultOptions struct{
	url *string
	pageSize *int
}

func (o *opsLevelDefaultOptions) opsLevelOptions(ctx context.Context) []opslevel.Option {
	opts := []opslevel.Option{opslevel.SetContext(ctx)} // Always set the context.

	if o.url != nil {
		ops = append(opts, opslevel.SetURL(*o.url))
	}
	if o.pageSize != nil {
		ops = append(opts, opslevel.SetPageSize(*o.pageSize))
	}
}

// Option is our own functional option type for our own custom client.
type Option func(*Client)

func WithURL(url string) Option {
	return func(c *Client) {
		c.do.url = &url
	}
}

func WithPageSize(pageSize int) Option {
	return func(c *Client) {
		c.do.pageSize = &pageSize
	}
}

func WithCustomOption(custom string) Option {
	return func(c *Client) {
		c.customOption = custom
	}
}