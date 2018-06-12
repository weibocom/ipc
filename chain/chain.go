package chain

type Chain interface {
	Post(dna string) error
	Verify(dna string) error
	Close() error
}
