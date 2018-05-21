package store

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/weibocom/ipc/model"
)

type DBStore struct {
	db *gorm.DB
}

var _ Store = &DBStore{}

func NewMySQLStore(conn string) *DBStore {
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	s := &DBStore{db: db}

	s.db.AutoMigrate(&model.Account{})
	s.db.AutoMigrate(&model.Member{})
	s.db.AutoMigrate(&model.Post{})

	return s
}
func (s *DBStore) SaveAccount(a *model.Account) error {
	company := getCompany(a.Name)
	a.Company = company
	a.CreatedAt = time.Now()
	return s.db.Save(a).Error
}

func (s *DBStore) LoadAccount(name string) (*model.Account, error) {
	a := &model.Account{Name: name}
	db := s.db.Model(&model.Account{}).Where(a).First(a)

	if db.RecordNotFound() {
		return nil, ErrNonExist
	}

	if db.Error != nil {
		return nil, db.Error
	}
	return a, nil
}

func (s *DBStore) GetAccounts(company string, offset int, limit int) ([]*model.Account, error) {
	var accounts []*model.Account
	db := s.db.Model(&model.Account{}).Where(&model.Account{Company: company}).Offset(offset).Limit(limit).Find(&accounts)

	return accounts, db.Error
}

func (s *DBStore) ExistAccount(name string) (bool, error) {
	a, err := s.LoadAccount(name)
	if err == ErrNonExist {
		return false, nil
	}
	return a != nil, err
}

func (s *DBStore) SaveMember(m *model.Member) error {
	return s.db.Save(m).Error
}

func (s *DBStore) LoadMember(name string) (*model.Member, error) {
	a := &model.Member{Name: name}
	db := s.db.Model(&model.Member{}).Where(a).First(a)
	if db.RecordNotFound() {
		return nil, ErrNonExist
	}
	if db.Error != nil {
		return nil, db.Error
	}
	return a, nil
}

func (s *DBStore) ExistMember(name string) (bool, error) {
	m, err := s.LoadMember(name)
	if err == ErrNonExist {
		return false, nil
	}
	return m != nil, err
}

func (s *DBStore) GetPostCount() (int, error) {
	var count int
	db := s.db.Model(&model.Post{}).Count(&count)

	return count, db.Error
}

func (s *DBStore) SavePost(p *model.Post) error {
	return s.db.Save(p).Error
}

func (s *DBStore) LoadPost(dna model.DNA) (*model.Post, error) {
	a := &model.Post{DNA: dna.ID()}
	db := s.db.Model(&model.Post{}).Where(a).First(a)
	if db.RecordNotFound() {
		return nil, ErrNonExist
	}
	if db.Error != nil {
		return nil, db.Error
	}
	return a, nil
}

func (s *DBStore) ExistPost(dna model.DNA) (bool, error) {
	p, err := s.LoadPost(dna)
	if err == ErrNonExist {
		return false, nil
	}
	return p != nil, err
}

func (s *DBStore) GetLatestPost() (*model.Post, error) {
	a := &model.Post{}
	err := s.db.Model(&model.Post{}).Order("created_at desc").Last(a).Error
	return a, err
}

func (s *DBStore) GetPostByMsgID(author string, mid int64) (*model.Post, error) {
	a := &model.Post{}
	err := s.db.Model(&model.Post{}).Where("author = ? AND mid = ?", author, mid).First(&a).Error
	return a, err
}

func (s *DBStore) GetPostByAuther(author string, offset int, limit int) ([]*model.Post, error) {
	var a []*model.Post
	err := s.db.Model(&model.Post{}).Where("author = ?", author).Order("created_at desc").Offset(offset).Limit(limit).Find(&a).Error
	return a, err
}

func (s *DBStore) GetPostByDNA(dna model.DNA) (*model.Post, error) {
	a := &model.Post{}
	err := s.db.Model(&model.Post{}).Where("dna = ?", dna.ID()).First(&a).Error
	return a, err
}

func (s *DBStore) Close() error {
	return s.db.Close()
}
