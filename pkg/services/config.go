package services


type Config	struct {
	ZincsearchUrl	string
}

func (c	*Config) SetConfig(zincsearchUrl string) {
	c.ZincsearchUrl = zincsearchUrl
}