package store

// import (
// 	"testing"
// 	"time"

// 	"github.com/weibocom/ipc/model"
// )

// func TestMySQL(t *testing.T) {
// 	s := NewMySQLStore("root@/ipc?charset=utf8&parseTime=True&loc=Local")

// 	a := &model.Account{
// 		Name: "wbuser-1",
// 		WIF:  "abcdefghijklmndfddsfdfdsfwfe3ewfd",
// 	}
// 	err := s.SaveAccount(a)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ok, err := s.ExistAccount(a.Name)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !ok {
// 		t.Errorf("account %s doesn't be saved into db", a.Name)
// 	}

// 	m := &model.Member{
// 		Name:       "wbuser-1",
// 		ID:         12345678,
// 		SigningKey: "STMabcddsadasdasdsadasdasdasdasd",
// 		CreatedAt:  time.Now(),
// 	}

// 	err = s.SaveMember(m)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ok, err = s.ExistMember(m.Name)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !ok {
// 		t.Errorf("member %s doesn't be saved into db", m.Name)
// 	}

// 	p := &model.Post{
// 		DNA:     "abcdefg",
// 		Author:  "wbuser-1",
// 		Title:   "this is a title",
// 		Content: "Thisis the body",
// 		URI:     "http://weibo.com/abcdefg",
// 		Digest:  "asfwefwavdgreiwfwfwefwe",
// 	}

// 	err = s.SavePost(p)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ok, err = s.ExistPost(model.DNA(p.DNA))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !ok {
// 		t.Errorf("post %s doesn't be saved into db", m.Name)
// 	}
// }
