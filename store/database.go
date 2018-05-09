package store

import (
	"github.com/jinzhu/gorm"
	"github.com/weibocom/ipc/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBStore struct {
	db *gorm.DB
}

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
	err := s.db.Model(&model.Account{}).First(a).Error
	return a, err
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
	err := s.db.Model(&model.Member{}).First(a).Error
	return a, err
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
	err := s.db.Model(&model.Post{}).First(a).Error
	return a, err
}

func (s *DBStore) ExistPost(dna model.DNA) (bool, error) {
	p, err := s.LoadPost(dna)
	if err == ErrNonExist {
		return false, nil
	}
	return p != nil, err
}

func (s *DBStore) Close() error {
	return s.db.Close()
}
