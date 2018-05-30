package store

// import (
// 	"time"

// 	shell "github.com/ipfs/go-ipfs-api"
// 	"github.com/weibocom/ipc/model"
// )

// type IPFSStore struct {
// 	url    string
// 	client *shell.Shell
// }

// func NewIPFSStore(u string) *IPFSStore {
// 	s := &IPFSStore{
// 		url: u,
// 	}

// 	client := shell.NewShell(u)
// 	s.client = client

// 	return s
// }

// func (s *IPFSStore) SaveAccount(a *model.Account) error {
// 	company := getCompany(a.Name)
// 	a.Company = company
// 	a.CreatedAt = time.Now()
// 	return s.db.Save(a).Error
// }

// func (s *IPFSStore) LoadAccount(name string) (*model.Account, error) {
// 	a := &model.Account{Name: name}
// 	db := s.db.Model(&model.Account{}).Where(a).First(a)

// 	if db.RecordNotFound() {
// 		return nil, ErrNonExist
// 	}

// 	if db.Error != nil {
// 		return nil, db.Error
// 	}
// 	return a, nil
// }

// func (s *IPFSStore) GetAccounts(company string, offset int, limit int) ([]*model.Account, error) {
// 	var accounts []*model.Account
// 	db := s.db.Model(&model.Account{}).Where(&model.Account{Company: company}).Order("created_at desc").Offset(offset).Limit(limit).Find(&accounts)

// 	return accounts, db.Error
// }

// func (s *IPFSStore) ExistAccount(name string) (bool, error) {
// 	a, err := s.LoadAccount(name)
// 	if err == ErrNonExist {
// 		return false, nil
// 	}
// 	return a != nil, err
// }

// func (s *IPFSStore) GetAccountCount() (int, error) {
// 	var count int
// 	db := s.db.Model(&model.Account{}).Count(&count)

// 	return count, db.Error
// }

// func (s *IPFSStore) SaveMember(m *model.Member) error {
// 	return s.db.Save(m).Error
// }

// func (s *IPFSStore) LoadMember(name string) (*model.Member, error) {
// 	a := &model.Member{Name: name}
// 	db := s.db.Model(&model.Member{}).Where(a).First(a)
// 	if db.RecordNotFound() {
// 		return nil, ErrNonExist
// 	}
// 	if db.Error != nil {
// 		return nil, db.Error
// 	}
// 	return a, nil
// }

// func (s *IPFSStore) ExistMember(name string) (bool, error) {
// 	m, err := s.LoadMember(name)
// 	if err == ErrNonExist {
// 		return false, nil
// 	}
// 	return m != nil, err
// }

// func (s *IPFSStore) GetPostCount() (int, error) {
// 	var count int
// 	db := s.db.Model(&model.Post{}).Count(&count)

// 	return count, db.Error
// }

// func (s *IPFSStore) SavePost(p *model.Post) error {
// 	return s.db.Save(p).Error
// }

// func (s *IPFSStore) LoadPost(dna model.DNA) (*model.Post, error) {
// 	a := &model.Post{DNA: dna.ID()}
// 	db := s.db.Model(&model.Post{}).Where(a).First(a)
// 	if db.RecordNotFound() {
// 		return nil, ErrNonExist
// 	}
// 	if db.Error != nil {
// 		return nil, db.Error
// 	}
// 	return a, nil
// }

// func (s *IPFSStore) ExistPost(dna model.DNA) (bool, error) {
// 	p, err := s.LoadPost(dna)
// 	if err == ErrNonExist {
// 		return false, nil
// 	}
// 	return p != nil, err
// }

// func (s *IPFSStore) GetLatestPost() (*model.Post, error) {
// 	a := &model.Post{}
// 	err := s.db.Model(&model.Post{}).Order("created_at desc").Last(a).Error
// 	return a, err
// }

// func (s *IPFSStore) GetPostByMsgID(author string, mid int64) (*model.Post, error) {
// 	a := &model.Post{}
// 	err := s.db.Model(&model.Post{}).Where("author = ? AND mid = ?", author, mid).First(&a).Error
// 	return a, err
// }

// func (s *IPFSStore) GetPostByAuther(author string, offset int, limit int) ([]*model.Post, error) {
// 	var a []*model.Post
// 	err := s.db.Model(&model.Post{}).Where("author = ?", author).Order("created_at desc").Offset(offset).Limit(limit).Find(&a).Error
// 	return a, err
// }

// func (s *IPFSStore) GetPostByDNA(dna model.DNA) (*model.Post, error) {
// 	s.ls
// 	return a, err
// }

// func (s *IPFSStore) Close() error {
// 	return nil
// }
