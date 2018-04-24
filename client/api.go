package client

type IClient interface {
	CreateAccount(name string, meta interface{}) (*Account, error)
	Post(author string, title string, content []byte, abstracts []string, url string) (*DNA, error)
	LookupContent(dna DNA) (*Content, error)
	Verify(author string, dna DNA) (error, bool)
	CheckSimilar(a, b DNA) (int, error)
	Members() (error, []Member)
	AddMember(m Member) error
	RemoveMember(m Member) error
}
