package weiboapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var (
	t = ""
)

// TODO: demo 使用，将来移除

func init() {
	t, _ = os.LookupEnv("wbat")

}

type User struct {
	AllowAllActMsg   bool   `json:"allow_all_act_msg"`
	AllowAllComment  bool   `json:"allow_all_comment"`
	AvatarHd         string `json:"avatar_hd"`
	AvatarLarge      string `json:"avatar_large"`
	BiFollowersCount int64  `json:"bi_followers_count"`
	BlockApp         int64  `json:"block_app"`
	BlockWord        int64  `json:"block_word"`
	Cardid           string `json:"cardid"`
	City             string `json:"city"`
	Class            int64  `json:"class"`
	CoverImage       string `json:"cover_image"`
	CoverImagePhone  string `json:"cover_image_phone"`
	CreatedAt        string `json:"created_at"`
	CreditScore      int64  `json:"credit_score"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	FavouritesCount  int64  `json:"favourites_count"`
	FollowMe         bool   `json:"follow_me"`
	FollowersCount   int64  `json:"followers_count"`
	Following        bool   `json:"following"`
	FriendsCount     int64  `json:"friends_count"`
	Gender           string `json:"gender"`
	GeoEnabled       bool   `json:"geo_enabled"`
	ID               int64  `json:"id"`
	Idstr            string `json:"idstr"`
	Insecurity       struct {
		SexualContent bool `json:"sexual_content"`
	} `json:"insecurity"`
	Lang             string `json:"lang"`
	Like             bool   `json:"like"`
	LikeMe           bool   `json:"like_me"`
	Location         string `json:"location"`
	Mbrank           int64  `json:"mbrank"`
	Mbtype           int64  `json:"mbtype"`
	Name             string `json:"name"`
	OnlineStatus     int64  `json:"online_status"`
	PagefriendsCount int64  `json:"pagefriends_count"`
	ProfileImageURL  string `json:"profile_image_url"`
	ProfileURL       string `json:"profile_url"`
	Province         string `json:"province"`
	Ptype            int64  `json:"ptype"`
	Remark           string `json:"remark"`
	ScreenName       string `json:"screen_name"`
	Star             int64  `json:"star"`
	Status           struct {
		AttitudesCount    int64  `json:"attitudes_count"`
		BizFeature        int64  `json:"biz_feature"`
		CanEdit           bool   `json:"can_edit"`
		Cardid            string `json:"cardid"`
		CommentManageInfo struct {
			ApprovalCommentType   int64 `json:"approval_comment_type"`
			CommentPermissionType int64 `json:"comment_permission_type"`
		} `json:"comment_manage_info"`
		CommentsCount        int64         `json:"comments_count"`
		ContentAuth          int64         `json:"content_auth"`
		CreatedAt            string        `json:"created_at"`
		DarwinTags           []interface{} `json:"darwin_tags"`
		Favorited            bool          `json:"favorited"`
		Geo                  interface{}   `json:"geo"`
		GifIds               string        `json:"gif_ids"`
		HasActionTypeCard    int64         `json:"hasActionTypeCard"`
		HotWeiboTags         []interface{} `json:"hot_weibo_tags"`
		ID                   int64         `json:"id"`
		Idstr                string        `json:"idstr"`
		InReplyToScreenName  string        `json:"in_reply_to_screen_name"`
		InReplyToStatusID    string        `json:"in_reply_to_status_id"`
		InReplyToUserID      string        `json:"in_reply_to_user_id"`
		IsLongText           bool          `json:"isLongText"`
		IsPaid               bool          `json:"is_paid"`
		IsShowBulletin       int64         `json:"is_show_bulletin"`
		MblogVipType         int64         `json:"mblog_vip_type"`
		Mid                  string        `json:"mid"`
		Mlevel               int64         `json:"mlevel"`
		MoreInfoType         int64         `json:"more_info_type"`
		PendingApprovalCount int64         `json:"pending_approval_count"`
		PicUrls              []interface{} `json:"pic_urls"`
		PositiveRecomFlag    int64         `json:"positive_recom_flag"`
		RepostsCount         int64         `json:"reposts_count"`
		Source               string        `json:"source"`
		SourceAllowclick     int64         `json:"source_allowclick"`
		SourceType           int64         `json:"source_type"`
		Text                 string        `json:"text"`
		TextTagTips          []interface{} `json:"text_tag_tips"`
		Truncated            bool          `json:"truncated"`
		UserType             int64         `json:"userType"`
		Visible              struct {
			ListID int64 `json:"list_id"`
			Type   int64 `json:"type"`
		} `json:"visible"`
	} `json:"status"`
	StatusesCount     int64  `json:"statuses_count"`
	StoryReadState    int64  `json:"story_read_state"`
	Urank             int64  `json:"urank"`
	URL               string `json:"url"`
	UserAbility       int64  `json:"user_ability"`
	VclubMember       int64  `json:"vclub_member"`
	Verified          bool   `json:"verified"`
	VerifiedReason    string `json:"verified_reason"`
	VerifiedReasonURL string `json:"verified_reason_url"`
	VerifiedSource    string `json:"verified_source"`
	VerifiedSourceURL string `json:"verified_source_url"`
	VerifiedTrade     string `json:"verified_trade"`
	VerifiedType      int64  `json:"verified_type"`
	Weihao            string `json:"weihao"`
}

func GetIDByName(name string) (int64, error) {
	if t == "" {
		return -1, nil
	}
	name = url.QueryEscape(name)
	resp, err := http.Get("https://api.weibo.com/2/users/show.json?screen_name=" + name + "&access_token=" + t)
	if err != nil {
		return -1, nil
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil
	}

	user := User{}
	json.Unmarshal(buf, &user)
	if err != nil {
		return -1, nil
	}

	return user.ID, nil
}
