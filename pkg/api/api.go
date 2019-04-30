package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/browny/dpipe/internal/domain/dcard"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
)

const (
	domain = "www.dcard.tw"

	comments = "https://www.dcard.tw/_api/posts/%d/comments"
	posts    = "https://www.dcard.tw/_api/forums/%s/posts"
	post     = "https://www.dcard.tw/_api/posts/%d"
)

func GetComments(postID int64) ([]dcard.Comment, error) {
	url := fmt.Sprintf(post, postID)

	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	respJson := resp.String()
	count := gjson.Get(respJson, "commentCount").Int()

	url = fmt.Sprintf(comments, postID)
	ro := &grequests.RequestOptions{}
	result := []dcard.Comment{}
	for {
		resp, err := grequests.Get(url, ro)
		if err != nil {
			return nil, err
		}
		respJson := resp.String()

		tmpResult := []dcard.Comment{}
		err = json.Unmarshal([]byte(respJson), &tmpResult)
		if err != nil {
			return nil, err
		}
		result = append(result, tmpResult...)

		if result[len(result)-1].Floor >= int(count) {
			break
		}

		// update the `after` parameter for next query
		ro = &grequests.RequestOptions{
			Params: map[string]string{"after": strconv.Itoa(result[len(result)-1].Floor)},
		}
	}
	return result, nil
}

func GetPostsUntil(forum string, postID int64) ([]dcard.Post, error) {
	url := fmt.Sprintf(posts, forum)
	ro := &grequests.RequestOptions{}
	result := []dcard.Post{}

	for {
		resp, err := grequests.Get(url, ro)
		if err != nil {
			return nil, err
		}
		respJson := resp.String()

		contained, idx := contains(respJson, postID)

		tmpResult := []dcard.Post{}
		err = json.Unmarshal([]byte(respJson), &tmpResult)
		if err != nil {
			return nil, err
		}
		result = append(result, tmpResult[0:idx+1]...)

		if contained {
			break
		}

		// update the `before` parameter for next query
		lastId := gjson.Get(respJson, fmt.Sprintf("%d.id", idx)).String()
		ro = &grequests.RequestOptions{
			Params: map[string]string{"before": lastId},
		}
	}

	return result, nil
}

func contains(json string, targetId int64) (bool, int) {
	ids := []int64{}
	for _, v := range gjson.Get(json, "#.id").Array() {
		ids = append(ids, v.Int())
	}
	for idx, val := range ids {
		if targetId == val {
			return true, idx
		}
	}
	return false, len(ids) - 1
}
