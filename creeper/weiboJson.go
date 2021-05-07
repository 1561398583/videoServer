package creeper

type MainPage struct {
	OK int `json:"ok"`
	Data struct{
		TabsInfo struct{
			Tabs []*struct{
				Containerid string `json:"containerid"`
				TabType string `json:"tab_type"`
			}	`json:"tabs"`
		}	`json:"tabsInfo"`
	} `json:"data"`
}


type OnePageIndex struct {
	OK int `json:"ok"`

	Data struct{

		CardlistInfo struct{
			Containerid string	`json:"containerid"`
			Show_style int	`json:"show_style"`
			Since_id int64	`json:"since_id"`
			Total int	`json:"total"`
		}	`json:"cardlistInfo"`

		Cards []Card	`json:"cards"`

		Scheme string	`json:"scheme"`

	} `json:"data"`
}

type UserAllVidoes struct {
	CardlistInfo struct{
		Uid string	`json:"uid"`
		Containerid string	`json:"containerid"`
		Show_style int	`json:"show_style"`
		Since_id int64	`json:"since_id"`
		Total int	`json:"total"`
		V_P string	`json:"v_p"`
	}	`json:"cardlistInfo"`

	Cards []Card	`json:"cards"`
}

type Card struct {
	Card_type int	`json:"card_type"`
	Itemid string	`json:"itemid"`
	Show_type int	`json:"show_type"`
	Title_style int	`json:"title_style"`
	Mblog struct{
		Bid string	`json:"bid"`
		CreatedAt string	`json:"created_at"`
		Fid int64	`json:"fid"`
		Id string	`json:"id"`		//这条微博的id
		Mid string	`json:"mid"`
		CommentsCount int `json:"comments_count"` //评论数
		AttitudesCount int `json:"attitudes_count"` //点赞数
		PageInfo struct{
			Title string	`json:"title"`	//video标题
			MediaInfo struct{	//video info
				Duration float32	`json:"duration"`
				StreamUrl string	`json:"stream_url"`
				StreamUrlHd string	`json:"stream_url_hd"`
			}	`json:"media_info"`
			ObjectId string	`json:"object_id"`
			ObjectType int	`json:"object_type"`
			PagePic struct{
				Height int	`json:"height"`
				Width int	`json:"width"`
				Url string	`json:"url"`
			}	`json:"page_pic"`
			PageUrl string	`json:"page_url"`	//pageUrl
			Type string	`json:"type"`	//video
			UrlOri string	`json:"url_ori"`
			Urls struct{
				Hevc_mp4_hd string	`json:"hevc_mp4_hd"`
				Mp4_720p_mp4 string	`json:"mp4_720p_mp4"`
				Mp4_hd_mp4 string	`json:"mp4_hd_mp4"`
				Mp4_ld_mp4 string	`json:"mp4_ld_mp4"`
			}	`json:"urls"`
		}	`json:"page_info"`
		RewardScheme string	`json:"reward_scheme"`
		Rid string	`json:"rid"`
	}	`json:"mblog"`
}

type CommentList struct {
	Ok int	`json:"ok"`
	Data struct{
		MaxId int64	`json:"max_id"`
		MaxIdType int	`json:"max_id_type"`
		TotalNumber int	`json:"total_number"`
		Data []*Comment	`json:"data"`
	}	`json:"data"`
}

type ChildCommentList struct {
	Ok int	`json:"ok"`
	Data []*Comment	`json:"data"`
	MaxId int64	`json:"max_id"`
	MaxIdType int	`json:"max_id_type"`
}

type Comment struct {
	CreatedAt string	`json:"created_at"`
	Id string	`json:"id"`
	LikeCount int	`json:"like_count"`
	RootId string	`json:"rootid"` //如果Rootid和Id相等，说明不是子评论
	TotalNumber int	`json:"total_number"`	//子评论数
	Text string	`json:"text"`
	User User	`json:"user"`
}

type Level1Comment struct {
	CreatedAt string	`json:"created_at"`
	Id string	`json:"id"`
	LikeCount int	`json:"like_count"`
	Rootid string	`json:"rootid"`
	Chilren []*Level2Comment
	TotalNumber int	`json:"total_number"`	//子评论数
	Text string	`json:"text"`
	User User	`json:"user"`
}

type Level2Comment struct {
	CreatedAt string	`json:"created_at"`
	Id string	`json:"id"`
	LikeCount int	`json:"like_count"`
	Rootid string	`json:"rootid"`
	Text string	`json:"text"`
	User User	`json:"user"`
}


type User struct {
	Id int64	`json:"id"`
	ScreenName string	`json:"screen_name"`
	Avatar_hd string	`json:"avatar_hd"`
}

