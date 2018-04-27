package client

import (
	"fmt"

	"github.com/weibocom/steem-rpc/steem/translit"
	"github.com/weibocom/steem-rpc/steem/types"
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

	_, err := c.SendTrx(op)

	return err == nil, err

}
