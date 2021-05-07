package creeper

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetFileNameFromUrl(t *testing.T) {
	var tests = []struct{
		input string
		want string
	}{
		{"https://f.video.weibocdn.com/lAVblciAlx07L20y92Ok01041200cf3R0E010.mp4?label=mp4_hd&template=540x960.24.0&trans_finger=7c347e6ee1691b93dc7e5726f4ef34b3&ori=0&ps=1BThIDDSKizpZI&Expires=1617549087&ssig=LH7pc7gTuV&KID=unistore,video",
			"lAVblciAlx07L20y92Ok01041200cf3R0E010.mp4",
		},
		{"https://f.video.weibocdn.com/lKLsObu5lx07L0xFuIpG010412008RrO0E010.mp4",
			"lKLsObu5lx07L0xFuIpG010412008RrO0E010.mp4",
		},
	}

	for _, test := range tests{
		if s := GetFileNameFromUrl(test.input); s != test.want {
			errorStr := fmt.Sprintf("GetFileNameFromUrl => want : %s, but : %s", test.want, s)
			t.Errorf(errorStr)
		}
	}
}

func TestUrlAddParams(t *testing.T) {
	tests := [2]struct{
		url string
		params map[string]string
		want string
	}{}

	tests[0].url = "https://f.video.weibocdn.com/lAVblciAlx07L20y92Ok01041200cf3R0E010.mp4"
	tests[0].want = "https://f.video.weibocdn.com/lAVblciAlx07L20y92Ok01041200cf3R0E010.mp4?label=mp4_hd&template=540x960.24.0&trans_finger=7c347e6ee1691b93dc7e5726f4ef34b3"
	params0 := make(map[string]string)
	params0["label"] = "mp4_hd"
	params0["template"] = "540x960.24.0"
	params0["trans_finger"] = "7c347e6ee1691b93dc7e5726f4ef34b3"
	tests[0].params = params0

	tests[1].url = "https://f.video.weibocdn.com/lKLsObu5lx07L0xFuIpG010412008RrO0E010.mp4"
	tests[1].want = "https://f.video.weibocdn.com/lKLsObu5lx07L0xFuIpG010412008RrO0E010.mp4"
	params1 := make(map[string]string)
	tests[1].params = params1


	for _, test := range tests{
		if s := UrlAddParams(test.url, test.params); s != test.want {
			errorStr := fmt.Sprintf("want : %s, but : %s", test.want, s)
			t.Errorf(errorStr)
		}
	}
}

func TestSetCookie(t *testing.T) {
	tests := []struct{
		SetCookie string
		Name string
		Value string
	}{
		{
			"_T_WM=10058446894; expires=Wed, 07-Apr-2021 16:00:00 GMT; Max-Age=83286; path=/; domain=.weibo.cn",
			"_T_WM",
			"10058446894",
		},
		{
			"XSRF-TOKEN=deleted; expires=Thu, 01-Jan-1970 00:00:01 GMT; Max-Age=0; path=/; domain=.weibo.cn; HttpOnly",
			"XSRF-TOKEN",
			"deleted",
		},
		{
			"XSRF-TOKEN=5cb711; expires=Tue, 06-Apr-2021 17:11:54 GMT; Max-Age=1200; path=/; domain=m.weibo.cn",
			"XSRF-TOKEN",
			"5cb711",
		},
		{
			"M_WEIBOCN_PARAMS=fid%3D1005052640113513%26uicode%3D10000011; expires=Tue, 06-Apr-2021 17:01:54 GMT; Max-Age=600; path=/; domain=.weibo.cn; HttpOnly",
			"M_WEIBOCN_PARAMS",
			"fid%3D1005052640113513%26uicode%3D10000011",
		},
	}

	cookies := make([]string, 0)
	for _, to := range tests {
		cookies = append(cookies, to.SetCookie)
	}
	cookieMap := make(map[string]string)
	SetCookie(cookies)
	if cookieMap["_T_WM"] != "10058446894" {
		t.Errorf("want : %s, but : %s", "10058446894", cookieMap["_T_WM"])
	}
	if cookieMap["XSRF-TOKEN"] != "5cb711" {
		t.Errorf("want : %s, but : %s", "5cb711", cookieMap["XSRF-TOKEN"])
	}
	if cookieMap["M_WEIBOCN_PARAMS"] != "fid%3D1005052640113513%26uicode%3D10000011" {
		t.Errorf("want : %s, but : %s", "fid%3D1005052640113513%26uicode%3D10000011", cookieMap["M_WEIBOCN_PARAMS"])
	}
}

func TestUrlEncode(t *testing.T)  {
	s, err := url.QueryUnescape("oid%3D4617137337402598%26luicode%3D10000011%26lfid%3D2315672640113513%26fid%3D2315672640113513%26uicode%3D10000011")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s)
}


