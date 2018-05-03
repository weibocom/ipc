package client

import (
	"fmt"

	"github.com/weibocom/ipc/steem/translit"
	"github.com/weibocom/ipc/steem/types"
)

// Post add a post.
func (c *Client) Post(authorname, title, body, permlink, ptag, postImage string, tags []string) (bool, error) {
	op := CreateCommentOperation(authorname, title, body, permlink, ptag, postImage, tags)
	_, err := c.SendTrx(op)

	return err == nil, err

}

// CreateCommentOperation creates a CommentOeration.
func CreateCommentOperation(authorname, title, body, permlink, ptag, postImage string, tags []string) *types.CommentOperation {
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

	return &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         authorname,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JsonMetadata:   jsonMeta,
	}
}

func (c *Client) BatchPost(ops []types.Operation) (bool, error) {
	_, err := c.SendTrx(ops...)
	return err == nil, err
}