package creeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"yx.com/videos/db"
)

func GetVideoComments(vidoeId string) (int,error) {
	cs, err := fetchComments(vidoeId)
	if err != nil {
		return 0,errors.New("GetVideoComments error : " + err.Error())
	}
	//获取所有用户
	users := getCommentsUsers(cs)
	for _, user := range users {
		err := createUser(user)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = commnets2DbComments(cs, vidoeId)
	if err != nil {
		return 0,errors.New("GetVideoComments error : " + err.Error())
	}
	return len(cs), nil
}


func commnets2DbComments(cs []*Comment, videoId string)  error{
	//comments 写入数据库
	dbcs := make([]*db.Comment, 0)
	for _, c := range cs {
		createAt, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", c.CreatedAt)
		if err != nil {
			panic(err)
		}
		createT := createAt.Format("2006-01-02 15:04:05")
		//评论等级1是video评论；2是1的子评论
		level := 1
		if c.Id != c.RootId {
			level = 2
		}
		uid := strconv.FormatInt(c.User.Id, 10)
		dbc := db.Comment{
			ID:              c.Id,
			CreatedAt:       createT,
			VideoId:         videoId,
			Comment:         c.Text,
			UserId:          uid,
			ToUserId:        "",
			ChildNum:        c.TotalNumber,
			LikeNum:         c.LikeCount,
			Level: level,
			RootId: c.RootId,
		}
		dbcs = append(dbcs, &dbc)
	}
	err := db.AddComments(dbcs)
	if err != nil {
		return errors.New("comments2DbComments error : " + err.Error())
	}
	return nil
}

//获取所有评论的所有user（不重复）
func getCommentsUsers(cs []*Comment) []*User {
	userMap := make(map[int64]*User)
	for _, c := range cs {
		userMap[c.User.Id] = &c.User
	}
	users := make([]*User, 0)
	for _, u := range userMap {
		users = append(users, u)
	}
	return users
}



//抓取评论
//wbId video的微博的id
func fetchComments(wbId string) ([]*Comment,error){
	allComments := make([]*Comment, 0)

	id := wbId
	mid := wbId
	maxId := ""
	maxIdType := "0"

	for  {
		param := CommentParam{
			Id:        id,
			Mid:       mid,
			MaxId:     maxId,
			MaxIdType: maxIdType,
		}
		pageComments, err := getOnePageComments(&param)
		if err != nil {	//出错就不用再执行下面的了
			return nil, errors.New("GetComments error : " + err.Error())
		}
		//获取子评论
		for _, c := range pageComments.Data.Data {
			allComments = append(allComments, c)
			if !(c.TotalNumber > 0) {
				continue
			}
			childComments, err := getAllChildComment(c.Id)
			if err != nil {
				return nil, errors.New("GetComments error : " + err.Error())
			}
			allComments = append(allComments, childComments...)
		}

		if pageComments.Data.MaxId == 0 {
			break
		}
		maxId = strconv.FormatInt(pageComments.Data.MaxId, 10)
		time.Sleep(5 * time.Second)
	}

	return allComments, nil
}

func getOnePageComments(param *CommentParam)  (*CommentList, error){
	ReqUrl := "https://m.weibo.cn/comments/hotflow"
	urlParams := make(map[string]string)
	urlParams["id"] = param.Id
	urlParams["mid"] = param.Mid
	if param.MaxId != "" {
		urlParams["max_id"] = param.MaxId
	}
	urlParams["max_id_type"] = param.MaxIdType

	headers := make(map[string]string)
	headers["accept"] = "application/json, text/plain, */*"
	headers["accept-encoding"] = "gzip, deflate, br"
	headers["accept-language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["mweibo-pwa"] = "1"
	refererUrl := "https://m.weibo.cn/detail/" + param.Id
	headers["referer"] = refererUrl
	headers["x-requested-with"] = "XMLHttpRequest"
	headers["x-xsrf-token"] = cookieMap["xsrf-token"]

	resp, err := GetResponse(ReqUrl, urlParams, headers)
	if err != nil {
		return nil, errors.New("getOnePageComments error : " + err.Error())
	}
	//ShowResponseHeader(resp)
	/*
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}
	*/
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(bs))

	//json转obj
	commentList := CommentList{}
	err = json.Unmarshal(bs, &commentList)
	if err != nil {
		panic(err)
	}

	return &commentList, nil
}


func getAllChildComment(fatherId string) ([]*Comment,error) {
	cs := make([]*Comment, 0)
	cid := fatherId
	max_id := "0"
	max_id_type := "0"

	for  {
		params := CommentParam{
			Id:        cid,
			Mid:       "",
			MaxId:     max_id,
			MaxIdType: max_id_type,
		}
		ccl, err := getOnePageChildComments(&params)
		if err != nil {
			return nil, errors.New("GetAllComments error : " + err.Error())
		}
		cs = append(cs, ccl.Data...)
		if ccl.MaxId == 0 {
			break
		}
		max_id = strconv.FormatInt(ccl.MaxId, 10)
		time.Sleep(5 * time.Second)
	}
	return cs, nil
}

func getOnePageChildComments(param *CommentParam)  (*ChildCommentList, error){
	ReqUrl := "https://m.weibo.cn/comments/hotFlowChild"
	urlParams := make(map[string]string)
	urlParams["cid"] = param.Id
	urlParams["max_id"] = param.MaxId
	urlParams["max_id_type"] = param.MaxIdType

	headers := make(map[string]string)
	headers["accept"] = "application/json, text/plain, */*"
	headers["accept-encoding"] = "gzip, deflate, br"
	headers["accept-language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["mweibo-pwa"] = "1"
	//https://m.weibo.cn/detail/4624180705758429?cid=4624181452604570
	refererUrl := "https://m.weibo.cn/detail/" + getMWeiboCnParams("oid") + "?cid=" + param.Id
	headers["referer"] = refererUrl
	headers["x-requested-with"] = "XMLHttpRequest"
	headers["x-xsrf-token"] = cookieMap["xsrf-token"]

	resp,err := GetResponse(ReqUrl, urlParams, headers)
	if err != nil {
		return nil, errors.New("getOnePageChildComments error : " + err.Error())
	}
	//ShowResponseHeader(resp)
	/*
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}

	 */
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(bs))

	//json转obj
	ccl := ChildCommentList{}
	err = json.Unmarshal(bs, &ccl)
	if err != nil {
		panic(err)
	}

	return &ccl, nil
}

type CommentParam struct {
	Id string
	Mid string
	MaxId string
	MaxIdType string
}


