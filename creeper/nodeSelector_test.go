package creeper

import (
	"testing"
)

/*
func init()  {
	f, err := os.Open("E:\\go_project\\videoProject\\server\\video\\creeper\\test.html")
	if err != nil {
		panic(err)
	}
	rootNode, err := html.Parse(f)
	if err != nil {
		panic(err)
	}
	rootNodeSelector = NewNodeSelector(rootNode)
	fmt.Println("init nodeSelector test finish")
}

 */

var rootNodeSelector *NodeSelector



func TestGetChildrenByCondition(t *testing.T)  {
	ps, err := rootNodeSelector.getChildrenByCondition("class=h")
	if err != nil {
		t.Error(err)
	}
	if ps[0].getText() != "我是a1" {
		t.Errorf("except 我是a1, but %s", ps[0].getText())
	}
	if ps[1].getText() != "我是b1" {
		t.Errorf("except 我是b1, but %s", ps[1].getText())
	}

	ps, err = rootNodeSelector.getChildrenByCondition("!=b,class=h")
	if err != nil {
		t.Error(err)
	}
	if ps[0].getText() != "我是b1" {
		t.Errorf("except 我是b1, but %s", ps[0].getText())
	}

	ps, err = rootNodeSelector.getChildrenByCondition("!=a,name=a1")
	if err != nil {
		t.Error(err)
	}
	if ps[0].getText() != "我是a1" {
		t.Errorf("except 我是a1, but %s", ps[0].getText())
	}
}


func TestGetChildrenByAttr(t *testing.T)  {
	ps := rootNodeSelector.getChildrenByAttr("name", "b1")
	if ps[0].getText() != "我是b1" {
		t.Errorf("except 我是b1, but %s", ps[0].getText())
	}
}

