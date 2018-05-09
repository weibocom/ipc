package store

import (
	"github.com/jinzhu/gorm"
	"github.com/weibocom/ipc/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBStore struct {
	db *gorm.DB
}

var _ Store = &DBStore{}

func NewMySQLStore(conn string) *DBStore {
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
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

func (s *DBStore) GetPosts(author string, afterDNA model.DNA, limit int) ([]*model.Post, error) {
	a := &model.Post{DNA: afterDNA.ID()}
	var afterID int64

	if s.db.Model(&model.Post{}).First(a).Error == nil {
		afterID = a.ID
	}

	var result []*model.Post
	db := s.db.Model(&model.Post{}).Order("id desc").Where("id > ?", afterID).Limit(limit).Find(&result)
	if db.RecordNotFound() {
		return nil, ErrNonExist
	}
	return result, db.Error
}

func (s *DBStore) GetLatestPost() (*model.Post, error) {
	a := &model.Post{}
	err := s.db.Model(&model.Post{}).Last(a).Error
	return a, err
}

func (s *DBStore) GetPostByURI(author string, uri string) (*model.Post, error) {
	a := &model.Post{}
	err := s.db.Model(&model.Post{}).Where("author = ? AND uri = ?", author, uri).First(&a).Error
	return a, err
}

func (s *DBStore) Close() error {
	return s.db.Close()
}
