package bilibili

import (
	"encoding/json"
	"io"

	"github.com/bitly/go-simplejson"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type SearchDynamicAtParam struct {
	Uid     int    `json:"uid"`     // 自己的uid
	Keyword string `json:"keyword"` // 搜索关键字
}

type DynamicGroupItem struct {
	Uid                int    `json:"uid"`                  // 用户id
	Uname              string `json:"uname"`                // 用户昵称
	Face               string `json:"face"`                 // 用户头像url
	Fans               int    `json:"fans"`                 // 用户粉丝数
	OfficialVerifyType int    `json:"official_verify_type"` // 认证信息?
}

type DynamicGroup struct {
	GroupType int                `json:"group_type"` // 2:我的关注。4:其他
	GroupName string             `json:"group_name"` // 分组名字
	Items     []DynamicGroupItem `json:"items"`      // 用户信息
}

type SearchDynamicAtResult struct {
	Groups []DynamicGroup `json:"groups"` // 内容分组
	Gt     int            `json:"_gt_"`   // 固定值0
}

// SearchDynamicAt 根据关键字搜索用户(at别人时的填充列表)
func (c *Client) SearchDynamicAt(param SearchDynamicAtParam) (*SearchDynamicAtResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_mix/v1/dynamic_mix/at_search"
	)
	return execute[*SearchDynamicAtResult](c, method, url, param)
}

type GetDynamicRepostDetailParam struct {
	DynamicId int `json:"dynamic_id"`                                 // 动态id
	Offset    int `json:"offset,omitempty" request:"query,omitempty"` // 偏移量
}

type DynamicRepostDetail struct {
	HasMore int `json:"has_more"` // 是否还有下一页
	Total   int `json:"total"`    // 总计包含
	Items   []struct {
		Desc struct {
			Uid         int   `json:"uid"`
			Type        int   `json:"type"`
			Rid         int64 `json:"rid"`
			Acl         int   `json:"acl"`
			View        int   `json:"view"`
			Repost      int   `json:"repost"`
			Like        int   `json:"like"`
			IsLiked     int   `json:"is_liked"`
			DynamicId   int64 `json:"dynamic_id"`
			Timestamp   int   `json:"timestamp"`
			PreDyId     int64 `json:"pre_dy_id"`
			OrigDyId    int64 `json:"orig_dy_id"`
			OrigType    int   `json:"orig_type"`
			UserProfile struct {
				Info struct {
					Uid     int    `json:"uid"`
					Uname   string `json:"uname"`
					Face    string `json:"face"`
					FaceNft int    `json:"face_nft"`
				} `json:"info"`
				Card struct {
					OfficialVerify struct {
						Type int    `json:"type"`
						Desc string `json:"desc"`
					} `json:"official_verify"`
				} `json:"card"`
				Vip struct {
					VipType    int   `json:"vipType"`
					VipDueDate int64 `json:"vipDueDate"`
					VipStatus  int   `json:"vipStatus"`
					ThemeType  int   `json:"themeType"`
					Label      struct {
						Path        string `json:"path"`
						Text        string `json:"text"`
						LabelTheme  string `json:"label_theme"`
						TextColor   string `json:"text_color"`
						BgStyle     int    `json:"bg_style"`
						BgColor     string `json:"bg_color"`
						BorderColor string `json:"border_color"`
					} `json:"label"`
					AvatarSubscript    int    `json:"avatar_subscript"`
					NicknameColor      string `json:"nickname_color"`
					Role               int    `json:"role"`
					AvatarSubscriptUrl string `json:"avatar_subscript_url"`
				} `json:"vip"`
				Pendant struct {
					Pid               int    `json:"pid"`
					Name              string `json:"name"`
					Image             string `json:"image"`
					Expire            int    `json:"expire"`
					ImageEnhance      string `json:"image_enhance"`
					ImageEnhanceFrame string `json:"image_enhance_frame"`
				} `json:"pendant"`
				Rank      string `json:"rank"`
				Sign      string `json:"sign"`
				LevelInfo struct {
					CurrentLevel int `json:"current_level"`
				} `json:"level_info"`
			} `json:"user_profile"`
			UidType      int    `json:"uid_type"`
			Stype        int    `json:"stype"`
			RType        int    `json:"r_type"`
			InnerId      int    `json:"inner_id"`
			Status       int    `json:"status"`
			DynamicIdStr string `json:"dynamic_id_str"`
			PreDyIdStr   string `json:"pre_dy_id_str"`
			OrigDyIdStr  string `json:"orig_dy_id_str"`
			RidStr       string `json:"rid_str"`
			Origin       struct {
				Uid          int    `json:"uid"`
				Type         int    `json:"type"`
				Rid          int    `json:"rid"`
				Acl          int    `json:"acl"`
				View         int    `json:"view"`
				Repost       int    `json:"repost"`
				Like         int    `json:"like"`
				DynamicId    int64  `json:"dynamic_id"`
				Timestamp    int    `json:"timestamp"`
				PreDyId      int    `json:"pre_dy_id"`
				OrigDyId     int    `json:"orig_dy_id"`
				UidType      int    `json:"uid_type"`
				Stype        int    `json:"stype"`
				RType        int    `json:"r_type"`
				InnerId      int    `json:"inner_id"`
				Status       int    `json:"status"`
				DynamicIdStr string `json:"dynamic_id_str"`
				PreDyIdStr   string `json:"pre_dy_id_str"`
				OrigDyIdStr  string `json:"orig_dy_id_str"`
				RidStr       string `json:"rid_str"`
			} `json:"origin"`
			Previous struct {
				Uid          int    `json:"uid"`
				Type         int    `json:"type"`
				Rid          int64  `json:"rid"`
				Acl          int    `json:"acl"`
				View         int    `json:"view"`
				Repost       int    `json:"repost"`
				Like         int    `json:"like"`
				DynamicId    int64  `json:"dynamic_id"`
				Timestamp    int    `json:"timestamp"`
				PreDyId      int64  `json:"pre_dy_id"`
				OrigDyId     int64  `json:"orig_dy_id"`
				UidType      int    `json:"uid_type"`
				Stype        int    `json:"stype"`
				RType        int    `json:"r_type"`
				InnerId      int    `json:"inner_id"`
				Status       int    `json:"status"`
				DynamicIdStr string `json:"dynamic_id_str"`
				PreDyIdStr   string `json:"pre_dy_id_str"`
				OrigDyIdStr  string `json:"orig_dy_id_str"`
				RidStr       string `json:"rid_str"`
			} `json:"previous"`
		} `json:"desc"`
		Card       string `json:"card"`
		ExtendJson string `json:"extend_json"`
		Display    struct {
			Origin struct {
				EmojiInfo struct {
					EmojiDetails []struct {
						EmojiName string `json:"emoji_name"`
						Id        int    `json:"id"`
						PackageId int    `json:"package_id"`
						State     int    `json:"state"`
						Type      int    `json:"type"`
						Attr      int    `json:"attr"`
						Text      string `json:"text"`
						Url       string `json:"url"`
						Meta      struct {
							Size int `json:"size"`
						} `json:"meta"`
						Mtime int `json:"mtime"`
					} `json:"emoji_details"`
				} `json:"emoji_info"`
				Relation struct {
					Status     int `json:"status"`
					IsFollow   int `json:"is_follow"`
					IsFollowed int `json:"is_followed"`
				} `json:"relation"`
			} `json:"origin"`
			Relation struct {
				Status     int `json:"status"`
				IsFollow   int `json:"is_follow"`
				IsFollowed int `json:"is_followed"`
			} `json:"relation"`
		} `json:"display"`
	} `json:"items"`
	Gt int `json:"_gt_"` // 固定值0
}

// GetDynamicRepostDetail 获取动态转发列表
//
// 见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/basicInfo.md
func (c *Client) GetDynamicRepostDetail(param GetDynamicRepostDetailParam) (*DynamicRepostDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_repost/v1/dynamic_repost/repost_detail"
	)
	return execute[*DynamicRepostDetail](c, method, url, param)
}

type GetDynamicLikeListParam struct {
	DynamicId int64 `json:"dynamic_id"`                             // 动态id
	Pn        int64 `json:"pn,omitempty" request:"query,omitempty"` // 页码
	Ps        int64 `json:"ps,omitempty" request:"query,omitempty"` // 每页数量。该值不得大于20
}

type DynamicLikeList struct {
	ItemLikes []struct { // 点赞信息列表主体
		Uid      int    `json:"uid"`
		Time     int    `json:"time"`
		FaceUrl  string `json:"face_url"`
		Uname    string `json:"uname"`
		UserInfo struct {
			Uid            int    `json:"uid"`
			Uname          string `json:"uname"`
			Face           string `json:"face"`
			Rank           string `json:"rank"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Vip struct {
				VipType    int   `json:"vipType"`
				VipDueDate int64 `json:"vipDueDate"`
				VipStatus  int   `json:"vipStatus"`
				ThemeType  int   `json:"themeType"`
				Label      struct {
					Path        string `json:"path"`
					Text        string `json:"text"`
					LabelTheme  string `json:"label_theme"`
					TextColor   string `json:"text_color"`
					BgStyle     int    `json:"bg_style"`
					BgColor     string `json:"bg_color"`
					BorderColor string `json:"border_color"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			} `json:"vip"`
			Pendant struct {
				Pid               int    `json:"pid"`
				Name              string `json:"name"`
				Image             string `json:"image"`
				Expire            int    `json:"expire"`
				ImageEnhance      string `json:"image_enhance"`
				ImageEnhanceFrame string `json:"image_enhance_frame"`
			} `json:"pendant"`
			Sign      string `json:"sign"`
			LevelInfo struct {
				CurrentLevel int `json:"current_level"`
			} `json:"level_info"`
		} `json:"user_info"`
		Attend int `json:"attend"`
	} `json:"item_likes"`
	HasMore    int `json:"has_more"`    // 是否还有下一页
	TotalCount int `json:"total_count"` // 总计点赞数
	Gt         int `json:"_gt_"`        // 固定值0
}

// GetDynamicLikeList 获取动态点赞列表
func (c *Client) GetDynamicLikeList(param GetDynamicLikeListParam) (*DynamicLikeList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_like/v1/dynamic_like/spec_item_likes"
	)
	return execute[*DynamicLikeList](c, method, url, param)
}

type GetDynamicLiveUserListParam struct {
	Size int `json:"size,omitempty" request:"query,omitempty"` // 每页显示数。默认为10
}

type DynamicLiveUserList struct {
	Count int        `json:"count"` // 直播者数量
	Group string     `json:"group"` // 固定值"default"，作用尚不明确
	Items []struct { // 直播者列表
		Uid   int    `json:"uid"`   // 直播者id
		Uname string `json:"uname"` // 直播者昵称
		Face  string `json:"face"`  // 直播者头像
		Link  string `json:"link"`  // 直播链接
		Title string `json:"title"` // 直播标题
	} `json:"items"`
	Gt int `json:"_gt_"` // 固定值0，作用尚不明确
}

// GetDynamicLiveUserList 获取正在直播的已关注者
func (c *Client) GetDynamicLiveUserList(param GetDynamicLiveUserListParam) (*DynamicLiveUserList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/w_live_users"
	)
	return execute[*DynamicLiveUserList](c, method, url, param)
}

type DynamicUpList struct {
	ButtonStatement string     `json:"button_statement"` // 固定值空，作用尚不明确
	Items           []struct { // 更新者列表
		UserProfile struct {
			Info struct {
				Uid   int    `json:"uid"`
				Uname string `json:"uname"`
				Face  string `json:"face"`
			} `json:"info"`
			Card struct {
				OfficialVerify struct {
					Type int    `json:"type"`
					Desc string `json:"desc"`
				} `json:"official_verify"`
			} `json:"card"`
			Vip struct {
				VipType       int    `json:"vipType"`
				VipDueDate    int64  `json:"vipDueDate"`
				DueRemark     string `json:"dueRemark"`
				AccessStatus  int    `json:"accessStatus"`
				VipStatus     int    `json:"vipStatus"`
				VipStatusWarn string `json:"vipStatusWarn"`
				ThemeType     int    `json:"themeType"`
				Label         struct {
					Path string `json:"path"`
				} `json:"label"`
			} `json:"vip"`
			Pendant struct {
				Pid          int    `json:"pid"`
				Name         string `json:"name"`
				Image        string `json:"image"`
				Expire       int    `json:"expire"`
				ImageEnhance string `json:"image_enhance"`
			} `json:"pendant"`
			Rank      string `json:"rank"`
			Sign      string `json:"sign"`
			LevelInfo struct {
				CurrentLevel int    `json:"current_level"`
				CurrentMin   int    `json:"current_min"`
				CurrentExp   int    `json:"current_exp"`
				NextExp      string `json:"next_exp"`
			} `json:"level_info"`
		} `json:"user_profile"`
		HasUpdate int `json:"has_update"`
	} `json:"items"`
	Gt int `json:"_gt_"` // 固定值0，作用尚不明确
}

type GetDynamicUpListParam struct {
	TeenagersMode int `json:"teenagers_mode,omitempty" request:"query,omitempty"` // 是否开启青少年模式。否：0。是：1
}

// GetDynamicUpList 获取发布新动态的已关注者
func (c *Client) GetDynamicUpList(param GetDynamicUpListParam) (*DynamicUpList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/w_dyn_uplist"
	)
	return execute[*DynamicUpList](c, method, url, param)
}

type RemoveDynamicParam struct {
	DynamicId int `json:"dynamic_id"` // 动态id
}

// RemoveDynamic 删除动态
func (c *Client) RemoveDynamic(param RemoveDynamicParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/rm_dynamic"
	)
	_, err := execute[any](c, method, url, param, fillCsrf(c))
	return err
}

type GetDynamicDetailParam struct {
	DynamicId int `json:"dynamic_id"` // 动态id
}

// DynamicCard 动态卡片内容。因为 ActivityInfos 、 Desc 、 Display 等字段会随着此动态类型不同发生一定的变化，无法统一，因此都转换成了 map[string]any ，请自行解析
type DynamicCard struct {
	ActivityInfos map[string]any `json:"activity_infos"` // 该条动态参与的活动
	Card          string         `json:"card"`           // 动态详细信息
	Desc          map[string]any `json:"desc"`           // 动态相关信息
	Display       map[string]any `json:"display"`        // 动态部分的可操作项
	ExtendJson    string         `json:"extend_json"`    // 动态扩展项
}

type DynamicDetail struct {
	Card   *DynamicCard `json:"card"` // 动态卡片内容
	Result int          `json:"result"`
	Gt     int          `json:"_gt_"`
}

// GetDynamicDetail 获取特定动态卡片信息
func (c *Client) GetDynamicDetail(param GetDynamicDetailParam) (*DynamicDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail"
	)
	return execute[*DynamicDetail](c, method, url, param)
}

type DynamicPortal struct {
	MyInfo struct { // 个人关注的一些信息
		Dyns      int      `json:"dyns"`      // 个人动态
		Face      string   `json:"face"`      // 头像url
		FaceNft   int      `json:"face_nft"`  // 含义尚不明确
		Follower  int      `json:"follower"`  // 粉丝数量
		Following int      `json:"following"` // 我的关注
		LevelInfo struct { // 本人等级内容
			CurrentExp   int   `json:"current_exp"`
			CurrentLevel int   `json:"current_level"` // 当前等级，0-6级
			CurrentMin   int   `json:"current_min"`
			LevelUp      int64 `json:"level_up"`
			NextExp      int   `json:"next_exp"`
		} `json:"level_info"`
		Mid      int      `json:"mid"`  // 账户mid
		Name     string   `json:"name"` // 账户名称
		Official struct { // 认证信息
			Desc  string `json:"desc"`  // 认证备注
			Role  int    `json:"role"`  // 认证类型，0：无，1 2 7：个人认证，3 4 5 6：机构认证
			Title string `json:"title"` // 认证信息
			Type  int    `json:"type"`  // 是否认证，-1：无，0：认证
		} `json:"official"`
		SpaceBg string   `json:"space_bg"` // 账户个人中心的背景横幅url
		Vip     struct { // vip信息
			AvatarSubscript    int      `json:"avatar_subscript"`     // 是否显示会员图标，0：不显示，1：显示
			AvatarSubscriptUrl string   `json:"avatar_subscript_url"` // 大会员角标地址
			DueDate            int64    `json:"due_date"`             // 会员过期时间，Unix时间戳（毫秒）
			Label              struct { // 会员标签
				BgColor               string `json:"bg_color"`                  // 会员标签背景颜色，颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
				BgStyle               int    `json:"bg_style"`                  // 固定值1，作用尚不明确
				BorderColor           string `json:"border_color"`              // 会员标签边框颜色，未使用
				ImgLabelUriHans       string `json:"img_label_uri_hans"`        // 固定值空
				ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"` // 大会员牌子图片，简体版
				ImgLabelUriHant       string `json:"img_label_uri_hant"`        // 固定值空
				ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"` // 大会员牌子图片，繁体版
				LabelTheme            string `json:"label_theme"`               // 会员标签，vip，annual_vip，ten_annual_vip，hundred_annual_vip，fools_day_hundred_annual_vip
				Path                  string `json:"path"`                      // 固定值空，作用尚不明确
				Text                  string `json:"text"`                      // 会员类型文案，大会员，年度大会员，十年大会员，百年大会员，最强绿鲤鱼
				TextColor             string `json:"text_color"`                // 会员标签文字颜色
				UseImgLabel           bool   `json:"use_img_label"`             // 固定值true
			} `json:"label"`
			NicknameColor string `json:"nickname_color"`  // 会员昵称颜色，颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
			Role          int    `json:"role"`            // 大会员类型，1：月度大会员，3：年度大会员，7：十年大会员，15：百年大会员
			Status        int    `json:"status"`          // 会员状态，0：无，1：有
			ThemeType     int    `json:"theme_type"`      // 固定值0，作用尚不明确
			TvVipPayType  int    `json:"tv_vip_pay_type"` // 电视大会员支付类型
			TvVipStatus   int    `json:"tv_vip_status"`   // 电视大会员状态，0：未开通
			Type          int    `json:"type"`            // 会员类型，0：无，1：月大会员，2：年度及以上大会员
			VipPayType    int    `json:"vip_pay_type"`    // 支付类型，0：未支付，1：已支付
		} `json:"vip"`
	} `json:"my_info"`
	UpList []struct { // 最近更新的up主列表
		Face            string `json:"face"`       // UP主头像
		HasUpdate       bool   `json:"has_update"` // 最近是否有更新
		IsReserveRecall bool   `json:"is_reserve_recall"`
		Mid             int    `json:"mid"`   // UP主mid
		Uname           string `json:"uname"` // UP主昵称
	} `json:"up_list"`
}

// GetDynamicPortal 获取最近更新UP主列表（其实就是获取自己的动态门户）
func (c *Client) GetDynamicPortal() (*DynamicPortal, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/portal"
	)
	return execute[*DynamicPortal](c, method, url, nil)
}

// UploadDynamicBfs 为图片动态上传图片
func (c *Client) UploadDynamicBfs(fileName string, file io.Reader, category string) (url string, size Size, err error) {
	biliJct := c.getCookie("bili_jct")
	if len(biliJct) == 0 {
		return "", Size{}, errors.New("B站登录过期")
	}
	resp, err := c.resty.R().
		SetFileReader("file_up", fileName, file).SetQueryParams(map[string]string{
		"category": category,
		"csrf":     biliJct,
	}).Post("https://api.bilibili.com/x/dynamic/feed/draw/upload_bfs")
	if err != nil {
		return "", Size{}, errors.WithStack(err)
	}
	if resp.StatusCode() != 200 {
		return "", Size{}, errors.Errorf("status code: %d", resp.StatusCode())
	}
	var response commonResp[struct {
		ImageUrl    string `json:"image_url"`
		ImageWidth  int    `json:"image_width"`
		ImageHeight int    `json:"image_height"`
	}]
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return "", Size{}, errors.WithStack(err)
	}
	if response.Code != 0 {
		return "", Size{}, errors.Errorf("错误码: %d, 错误信息: %s", response.Code, response.Message)
	}
	data := response.Data
	return data.ImageUrl, Size{Width: data.ImageWidth, Height: data.ImageHeight}, errors.WithStack(err)
}

type CreateDynamicParam struct {
	DynamicId       int          `json:"dynamic_id"`                                            // 0
	Type            int          `json:"type"`                                                  // 4
	Rid             int          `json:"rid"`                                                   // 0
	Content         string       `json:"content"`                                               // 动态内容
	UpChooseComment int          `json:"up_choose_comment,omitempty" request:"query,omitempty"` // 0
	UpCloseComment  int          `json:"up_close_comment,omitempty" request:"query,omitempty"`  // 0
	Extension       string       `json:"extension,omitempty" request:"query,omitempty"`         // 位置信息，参考 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/dynamic/publish.md
	AtUids          string       `json:"at_uids,omitempty" request:"query,omitempty"`           // 动态中 at 到的用户的 uid。使用逗号,分隔
	Ctrl            []FormatCtrl `json:"ctrl,omitempty" request:"query,omitempty"`              // 特殊格式控制 (如 at 别人时的蓝字体和链接)
}

type CreateDynamicResult struct {
	Result       int    `json:"result"`         // 0
	Errmsg       string `json:"errmsg"`         // 像是服务器日志一样的东西
	DynamicId    int    `json:"dynamic_id"`     // 动态 id
	CreateResult int    `json:"create_result"`  // 1
	DynamicIdStr string `json:"dynamic_id_str"` // 动态 id。字符串格式
	Gt           int    `json:"_gt_"`           // 0
}

// CreateDynamic 发表纯文本动态
func (c *Client) CreateDynamic(param CreateDynamicParam) (*CreateDynamicResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/create"
	)
	return execute[*CreateDynamicResult](c, method, url, param, fillCsrf(c))
}

// DynamicList 包含置顶及热门的动态列表
//
// TODO: 因为不清楚 attentions 字段（关注列表）的格式，暂未对此字段进行解析
type DynamicList struct {
	Cards         *DynamicCard `json:"cards"` // 动态列表
	FounderUid    int          `json:"founder_uid,omitempty"`
	HasMore       int          `json:"has_more"` // 当前话题是否有额外的动态，0：无额外动态，1：有额外动态
	IsDrawerTopic int          `json:"is_drawer_topic,omitempty"`
	Offset        string       `json:"offset"` // 接下来获取列表时的偏移值，一般为当前获取的话题列表下最后一个动态id
	Gt            int          `json:"_gt_"`   // 固定值0，作用尚不明确
}

type DynamicItem struct {
	Basic struct { // 见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/all.md#data%E5%AF%B9%E8%B1%A1---items%E6%95%B0%E7%BB%84%E4%B8%AD%E7%9A%84%E5%AF%B9%E8%B1%A1---basic%E5%AF%B9%E8%B1%A1
		CommentIdStr string `json:"comment_id_str"`
		CommentType  int    `json:"comment_type"`
		LikeIcon     struct {
			ActionUrl string `json:"action_url"`
			EndUrl    string `json:"end_url"`
			Id        int    `json:"id"`
			StartUrl  string `json:"start_url"`
		} `json:"like_icon"`
		RidStr string `json:"rid_str"`
	} `json:"basic"`
	IdStr   any `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string，不知道为什么。这里用any会带来一个问题，number会解析成为float64，有可能存在丢失精度问题。请谨慎使用
	Modules struct {
		ModuleAuthor struct {
			Avatar struct {
				ContainerSize struct {
					Height float64 `json:"height"`
					Width  float64 `json:"width"`
				} `json:"container_size"`
				FallbackLayers struct {
					IsCriticalGroup bool `json:"is_critical_group"`
					Layers          []struct {
						GeneralSpec struct {
							PosSpec struct {
								AxisX         float64 `json:"axis_x"`
								AxisY         float64 `json:"axis_y"`
								CoordinatePos int     `json:"coordinate_pos"`
							} `json:"pos_spec"`
							RenderSpec struct {
								Opacity int `json:"opacity"`
							} `json:"render_spec"`
							SizeSpec struct {
								Height float64 `json:"height"`
								Width  float64 `json:"width"`
							} `json:"size_spec"`
						} `json:"general_spec"`
						LayerConfig struct {
							IsCritical bool `json:"is_critical,omitempty"`
							Tags       struct {
								AvatarLayer struct {
								} `json:"AVATAR_LAYER,omitempty"`
								GeneralCfg struct {
									ConfigType    int `json:"config_type"`
									GeneralConfig struct {
										WebCssStyle struct {
											BorderRadius    string `json:"borderRadius"`
											BackgroundColor string `json:"background-color,omitempty"`
											Border          string `json:"border,omitempty"`
											BoxSizing       string `json:"boxSizing,omitempty"`
										} `json:"web_css_style"`
									} `json:"general_config"`
								} `json:"GENERAL_CFG"`
								IconLayer struct{} `json:"ICON_LAYER,omitempty"`
							} `json:"tags"`
						} `json:"layer_config"`
						Resource struct {
							ResAnimation struct {
								WebpSrc struct {
									Placeholder int `json:"placeholder"`
									Remote      struct {
										BfsStyle string `json:"bfs_style"`
										Url      string `json:"url"`
									} `json:"remote"`
									SrcType int `json:"src_type"`
								} `json:"webp_src"`
							} `json:"res_animation,omitempty"`
							ResType  int `json:"res_type"`
							ResImage struct {
								ImageSrc struct {
									Local   int `json:"local"`
									SrcType int `json:"src_type"`
								} `json:"image_src"`
							} `json:"res_image,omitempty"`
						} `json:"resource"`
						Visible bool `json:"visible"`
					} `json:"layers"`
				} `json:"fallback_layers"`
				Mid string `json:"mid"`
			} `json:"avatar"`
			Face           string `json:"face"`
			FaceNft        bool   `json:"face_nft"`
			Following      any    `json:"following"`
			JumpUrl        string `json:"jump_url"`
			Label          string `json:"label"`
			Mid            int    `json:"mid"`
			Name           string `json:"name"`
			OfficialVerify struct {
				Desc string `json:"desc"`
				Type int    `json:"type"`
			} `json:"official_verify"`
			Pendant struct {
				Expire            int    `json:"expire"`
				Image             string `json:"image"`
				ImageEnhance      string `json:"image_enhance"`
				ImageEnhanceFrame string `json:"image_enhance_frame"`
				NPid              int    `json:"n_pid"`
				Name              string `json:"name"`
				Pid               int    `json:"pid"`
			} `json:"pendant"`
			PubAction       string `json:"pub_action"`
			PubLocationText string `json:"pub_location_text"`
			PubTime         string `json:"pub_time"`
			PubTs           int    `json:"pub_ts"`
			Type            string `json:"type"`
			Vip             struct {
				AvatarSubscript    int    `json:"avatar_subscript"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
				DueDate            int64  `json:"due_date"`
				Label              struct {
					BgColor               string `json:"bg_color"`
					BgStyle               int    `json:"bg_style"`
					BorderColor           string `json:"border_color"`
					ImgLabelUriHans       string `json:"img_label_uri_hans"`
					ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelUriHant       string `json:"img_label_uri_hant"`
					ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
					LabelTheme            string `json:"label_theme"`
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					TextColor             string `json:"text_color"`
					UseImgLabel           bool   `json:"use_img_label"`
				} `json:"label"`
				NicknameColor string `json:"nickname_color"`
				Status        int    `json:"status"`
				ThemeType     int    `json:"theme_type"`
				Type          int    `json:"type"`
			} `json:"vip"`
		} `json:"module_author"`
		ModuleDynamic struct {
			Additional any `json:"additional"`
			Desc       *struct {
				RichTextNodes []struct {
					OrigText string `json:"orig_text"`
					Text     string `json:"text"`
					Type     string `json:"type"`
					JumpUrl  string `json:"jump_url,omitempty"`
					Style    any    `json:"style"`
					Emoji    struct {
						IconUrl string `json:"icon_url"`
						Size    int    `json:"size"`
						Text    string `json:"text"`
						Type    int    `json:"type"`
					} `json:"emoji,omitempty"`
					Rid string `json:"rid,omitempty"`
				} `json:"rich_text_nodes"`
				Text string `json:"text"`
			} `json:"desc"`
			Major *struct {
				Draw struct {
					Id    int `json:"id"`
					Items []struct {
						Height int     `json:"height"`
						Size   float64 `json:"size"`
						Src    string  `json:"src"`
						Tags   []any   `json:"tags"`
						Width  int     `json:"width"`
					} `json:"items"`
				} `json:"draw,omitempty"`
				Type    string `json:"type"`
				Archive struct {
					Aid   string `json:"aid"`
					Badge struct {
						BgColor string `json:"bg_color"`
						Color   string `json:"color"`
						IconUrl any    `json:"icon_url"`
						Text    string `json:"text"`
					} `json:"badge"`
					Bvid           string `json:"bvid"`
					Cover          string `json:"cover"`
					Desc           string `json:"desc"`
					DisablePreview int    `json:"disable_preview"`
					DurationText   string `json:"duration_text"`
					JumpUrl        string `json:"jump_url"`
					Stat           struct {
						Danmaku string `json:"danmaku"`
						Play    string `json:"play"`
					} `json:"stat"`
					Title string `json:"title"`
					Type  int    `json:"type"`
				} `json:"archive,omitempty"`
				Opus struct {
					Summary struct {
						RichTextNodes []struct {
							OrigText string `json:"orig_text,omitempty"`
							Text     string `json:"text,omitempty"`
							Type     string `json:"type,omitempty"`
						} `json:"rich_text_nodes,omitempty"`
						Text string `json:"text,omitempty"`
					} `json:"summary,omitempty"`
					Title string `json:"title,omitempty"`
				} `json:"opus,omitempty"`
			} `json:"major"`
			Topic any `json:"topic"`
		} `json:"module_dynamic"`
		ModuleMore struct {
			ThreePointItems []struct {
				Label string `json:"label"`
				Type  string `json:"type"`
			} `json:"three_point_items"`
		} `json:"module_more"`
		ModuleStat struct {
			Comment struct {
				Count     int  `json:"count"`
				Forbidden bool `json:"forbidden"`
			} `json:"comment"`
			Forward struct {
				Count     int  `json:"count"`
				Forbidden bool `json:"forbidden"`
			} `json:"forward"`
			Like struct {
				Count     int  `json:"count"`
				Forbidden bool `json:"forbidden"`
				Status    bool `json:"status"`
			} `json:"like"`
		} `json:"module_stat"`
	} `json:"modules"`
	Orig struct {
		Basic struct {
			CommentIdStr string `json:"comment_id_str"`
			CommentType  int    `json:"comment_type"`
			LikeIcon     struct {
				ActionUrl string `json:"action_url"`
				EndUrl    string `json:"end_url"`
				Id        int    `json:"id"`
				StartUrl  string `json:"start_url"`
			} `json:"like_icon"`
			RidStr string `json:"rid_str"`
		} `json:"basic"`
		IdStr   any `json:"id_str"`
		Modules struct {
			ModuleAuthor struct {
				Avatar struct {
					ContainerSize struct {
						Height float64 `json:"height"`
						Width  float64 `json:"width"`
					} `json:"container_size"`
					FallbackLayers struct {
						IsCriticalGroup bool `json:"is_critical_group"`
						Layers          []struct {
							GeneralSpec struct {
								PosSpec struct {
									AxisX         float64 `json:"axis_x"`
									AxisY         float64 `json:"axis_y"`
									CoordinatePos int     `json:"coordinate_pos"`
								} `json:"pos_spec"`
								RenderSpec struct {
									Opacity int `json:"opacity"`
								} `json:"render_spec"`
								SizeSpec struct {
									Height float64 `json:"height"`
									Width  float64 `json:"width"`
								} `json:"size_spec"`
							} `json:"general_spec"`
							LayerConfig struct {
								IsCritical bool `json:"is_critical,omitempty"`
								Tags       struct {
									AVATARLAYER struct {
									} `json:"AVATAR_LAYER,omitempty"`
									GENERALCFG struct {
										ConfigType    int `json:"config_type"`
										GeneralConfig struct {
											WebCssStyle struct {
												BorderRadius    string `json:"borderRadius"`
												BackgroundColor string `json:"background-color,omitempty"`
												Border          string `json:"border,omitempty"`
												BoxSizing       string `json:"boxSizing,omitempty"`
											} `json:"web_css_style"`
										} `json:"general_config"`
									} `json:"GENERAL_CFG,omitempty"`
									PENDENTLAYER struct {
									} `json:"PENDENT_LAYER,omitempty"`
									ICONLAYER struct {
									} `json:"ICON_LAYER,omitempty"`
								} `json:"tags"`
							} `json:"layer_config"`
							Resource struct {
								ResImage struct {
									ImageSrc struct {
										Placeholder int `json:"placeholder,omitempty"`
										Remote      struct {
											BfsStyle string `json:"bfs_style"`
											Url      string `json:"url"`
										} `json:"remote,omitempty"`
										SrcType int `json:"src_type"`
										Local   int `json:"local,omitempty"`
									} `json:"image_src"`
								} `json:"res_image"`
								ResType int `json:"res_type"`
							} `json:"resource"`
							Visible bool `json:"visible"`
						} `json:"layers"`
					} `json:"fallback_layers"`
					Mid string `json:"mid"`
				} `json:"avatar"`
				Decorate struct {
					CardUrl string `json:"card_url"`
					Fan     struct {
						Color  string `json:"color"`
						IsFan  bool   `json:"is_fan"`
						NumStr string `json:"num_str"`
						Number int    `json:"number"`
					} `json:"fan"`
					Id      int    `json:"id"`
					JumpUrl string `json:"jump_url"`
					Name    string `json:"name"`
					Type    int    `json:"type"`
				} `json:"decorate,omitempty"`
				Face           string `json:"face"`
				FaceNft        bool   `json:"face_nft"`
				Following      any    `json:"following"`
				JumpUrl        string `json:"jump_url"`
				Label          string `json:"label"`
				Mid            int    `json:"mid"`
				Name           string `json:"name"`
				OfficialVerify struct {
					Desc string `json:"desc"`
					Type int    `json:"type"`
				} `json:"official_verify"`
				Pendant struct {
					Expire            int    `json:"expire"`
					Image             string `json:"image"`
					ImageEnhance      string `json:"image_enhance"`
					ImageEnhanceFrame string `json:"image_enhance_frame"`
					NPid              int    `json:"n_pid"`
					Name              string `json:"name"`
					Pid               int    `json:"pid"`
				} `json:"pendant"`
				PubAction string `json:"pub_action"`
				PubTime   string `json:"pub_time"`
				PubTs     int    `json:"pub_ts"`
				Type      string `json:"type"`
				Vip       struct {
					AvatarSubscript    int    `json:"avatar_subscript"`
					AvatarSubscriptUrl string `json:"avatar_subscript_url"`
					DueDate            int64  `json:"due_date"`
					Label              struct {
						BgColor               string `json:"bg_color"`
						BgStyle               int    `json:"bg_style"`
						BorderColor           string `json:"border_color"`
						ImgLabelUriHans       string `json:"img_label_uri_hans"`
						ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
						ImgLabelUriHant       string `json:"img_label_uri_hant"`
						ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
						LabelTheme            string `json:"label_theme"`
						Path                  string `json:"path"`
						Text                  string `json:"text"`
						TextColor             string `json:"text_color"`
						UseImgLabel           bool   `json:"use_img_label"`
					} `json:"label"`
					NicknameColor string `json:"nickname_color"`
					Status        int    `json:"status"`
					ThemeType     int    `json:"theme_type"`
					Type          int    `json:"type"`
				} `json:"vip"`
			} `json:"module_author"`
			ModuleDynamic struct {
				Additional any `json:"additional"`
				Desc       *struct {
					RichTextNodes []struct {
						JumpUrl  string `json:"jump_url,omitempty"`
						OrigText string `json:"orig_text"`
						Text     string `json:"text"`
						Type     string `json:"type"`
						Emoji    struct {
							IconUrl string `json:"icon_url"`
							Size    int    `json:"size"`
							Text    string `json:"text"`
							Type    int    `json:"type"`
						} `json:"emoji,omitempty"`
					} `json:"rich_text_nodes"`
					Text string `json:"text"`
				} `json:"desc"`
				Major struct {
					Archive struct {
						Aid   string `json:"aid"`
						Badge struct {
							BgColor string `json:"bg_color"`
							Color   string `json:"color"`
							IconUrl any    `json:"icon_url"`
							Text    string `json:"text"`
						} `json:"badge"`
						Bvid           string `json:"bvid"`
						Cover          string `json:"cover"`
						Desc           string `json:"desc"`
						DisablePreview int    `json:"disable_preview"`
						DurationText   string `json:"duration_text"`
						JumpUrl        string `json:"jump_url"`
						Stat           struct {
							Danmaku string `json:"danmaku"`
							Play    string `json:"play"`
						} `json:"stat"`
						Title string `json:"title"`
						Type  int    `json:"type"`
					} `json:"archive,omitempty"`
					Type string `json:"type"`
					Draw struct {
						Id    int `json:"id"`
						Items []struct {
							Height int     `json:"height"`
							Size   float64 `json:"size"`
							Src    string  `json:"src"`
							Tags   []any   `json:"tags"`
							Width  int     `json:"width"`
						} `json:"items"`
					} `json:"draw,omitempty"`
				} `json:"major"`
				Topic any `json:"topic"`
			} `json:"module_dynamic"`
		} `json:"modules"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"orig,omitempty"`
	Type    string `json:"type"`
	Visible bool   `json:"visible"`
}

type DynamicInfo struct {
	HasMore        bool          `json:"has_more"`        // 是否有更多数据
	Items          []DynamicItem `json:"items"`           // 数据数组
	Offset         string        `json:"offset"`          // 偏移量，等于items中最后一条记录的id，获取下一页时使用
	UpdateBaseline string        `json:"update_baseline"` // 更新基线，等于items中第一条记录的id
	UpdateNum      int           `json:"update_num"`      // 本次获取获取到了多少条新动态，在更新基线以上的动态条数
}

type GetUserSpaceDynamicParam struct {
	Offset         string `json:"offset"`          // 分页偏移量
	HostMid        string `json:"host_mid"`        // 用户UID
	TimezoneOffset int    `json:"timezone_offset"` // -480
	Features       string `json:"features"`        // itemOpusStyle
}

// GetUserSpaceDynamic 获取用户空间动态，mid就是用户UID，无需登录。
//
// 返回结构较为繁琐，见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/space.md
func (c *Client) GetUserSpaceDynamic(param GetUserSpaceDynamicParam) (*DynamicInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space"
	)
	return execute[*DynamicInfo](c, method, url, param)
}

func (c *Client) RawGetUserSpaceDynamic(param map[string]string) (*simplejson.Json, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space"
	)
	return RawExecute(c, method, url, ContentTypeUrl, param, nil)
}
