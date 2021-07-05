package goResponse

type Cookie struct {
	Name     string
	Value    string
	Duration int
	Path     string
}

func NewCookie() *Cookie {

	return &Cookie{}
}

func (c *Cookie) Create(name string, value string, duration int, path string) {
	c.Name = name
	c.Value = value
	c.Duration = duration
	c.Path = path
}

func (c *Cookie) GetName() string {
	return c.Name
}

func (c *Cookie) GetValue() string {
	return c.Value
}

func (c *Cookie) GetDuration() int {
	return c.Duration
}

func (c *Cookie) GetPath() string {
	return c.Path
}
