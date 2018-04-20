package client

import (
	"fmt"
	"log"
	"os"

	"github.com/weibocom/steem-rpc/apis/networkbroadcast"
	"github.com/weibocom/steem-rpc/steem"
	"github.com/weibocom/steem-rpc/transactions"
	"github.com/weibocom/steem-rpc/translit"
	"github.com/weibocom/steem-rpc/types"
)

// Post add a post.
func (c *Client) Post(authorname, title, body, permlink, ptag, postImage string, tags []string) (bool, error) {
	if permlink == "" {
		permlink = translit.EncodeTitle(title)
	} else {
		permlink = translit.EncodeTitle(permlink)
	}
	tag := translit.EncodeTags(tags)
	if ptag == "" {
		ptag = translit.EncodeTag(tags[0])
	} else {
		ptag = translit.EncodeTag(ptag)
	}

	jsonMeta := fmt.Sprintf(`{
		"tags":  [%s],
		"image": [%s],
		"lib":   "go-steem-rpc"
	}`, tag, postImage)

	op := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         authorname,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JsonMetadata:   jsonMeta,
	}

	resp, err := c.SendTrx(op)

	if err != nil {
		return false, err
	}

	log.Printf("response: %+v", resp)

	return true, nil

}

// SendTrx signs and sends transactions.
func (c *Client) SendTrx(operations ...types.Operation) (resp *networkbroadcast.BroadcastResponse, err error) {
	props, cfgErr := c.Database.GetDynamicGlobalProperties()
	if cfgErr != nil {
		fmt.Println("failed to get config:%s", cfgErr.Error())
		os.Exit(-3)
	}

	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		fmt.Printf("failed to parse ref block prefix:%v\n", err.Error())
		return
	}

	tx := &types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	}

	stx := transactions.NewSignedTransaction(tx)

	for _, op := range operations {
		stx.PushOperation(op)
	}

	if err := stx.Sign(steem.GetPrivateKeys(), steem.SteemChain); err != nil {
		log.Printf("transaction sig err:%v\n", err.Error())
		return nil, err
	}

	resp, err = c.NetworkBroadcast.BroadcastTransactionSynchronous(stx.Transaction)
	if err != nil {
		log.Printf("broadcast failed:%v\n", err.Error())
		return
	}

	return resp, err
}
