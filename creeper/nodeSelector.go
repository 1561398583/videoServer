package creeper

import (
	"errors"
	"golang.org/x/net/html"
	"strings"
)

type NodeSelector struct {
	node *html.Node
}

func NewNodeSelector(node *html.Node) *NodeSelector {
	return &NodeSelector{node:node}
}


func (ns *NodeSelector) getAttr(attrName string) (string, error){
	for _, attr := range ns.node.Attr {
		if attr.Key == attrName {
			return attr.Val, nil
		}
	}
	return "", errors.New("not found")
}

func (ns *NodeSelector) getText() string{
	result := ""
	for child := ns.node.FirstChild; child != nil; child = child.NextSibling  {
		if child.Type == html.TextNode {
			result += child.Data
		}
	}
	return result
}

func (ns *NodeSelector) getChildrenByTagName(tagName string) []*NodeSelector {
	children := make([]*html.Node, 0)
	kv := KV{key:"!", value:tagName}
	condition := NewTagCondition(kv)
	findChildrenByCondition(ns.node, &children, condition)
	if len(children) == 0 {
		return nil
	}
	result := make([]*NodeSelector, len(children))
	for i := 0; i < len(children); i++  {
		result[i] = NewNodeSelector(children[i])
	}
	return result
}

func (ns *NodeSelector) getChildrenByAttr(key, value string) []*NodeSelector {
	children := make([]*html.Node, 0)
	kv := KV{key:key, value:value}
	condition := NewAttrCondition(kv)
	findChildrenByCondition(ns.node, &children, condition)
	if len(children) == 0 {
		return nil
	}
	result := make([]*NodeSelector, len(children))
	for i := 0; i < len(children); i++  {
		result[i] = NewNodeSelector(children[i])
	}
	return result
}



/**
每个条件以","分隔
！表示tag
其他用key=value表示
例如： "!=p,id=id1,class=class1 class2,name=yx"
*/
func (ns *NodeSelector) getChildrenByCondition(condition string) ([]*NodeSelector, error){
	//分解为key,value数组的形式
	conditions := strings.Split(condition, ",")
	if len(conditions) == 0 {
		return nil, errors.New("have no condition")
	}
	kvs := make([]KV, len(conditions))
	for i := 0; i < len(conditions); i++ {
		kv := strings.Split(conditions[i], "=")
		if len(kv) != 2 {
			return nil, errors.New(conditions[i] + " is illegal syntax")
		}
		kvs[i].key = kv[0]
		kvs[i].value = kv[1]
	}

	//根据第一个条件开始查找
	var result []*html.Node
	if kvs[0].key == "!" {		//如果是tag
		condition := NewTagCondition(kvs[0])
		findChildrenByCondition(ns.node, &result, condition)
		if len(result) == 0 {
			return nil, errors.New("no one")
		}
	}else {		//是Attr
		condition := NewAttrCondition(kvs[0])
		findChildrenByCondition(ns.node, &result, condition)
		if len(result) == 0 {
			return nil, errors.New("no one")
		}
	}

	//接着从第二个条件开始过滤
	for i := 1; i < len(kvs); i++ {
		if kvs[i].key == "!" {		//如果是tag
			result = filterByTagName(kvs[i].value, result)
		}else {		//是Attr
			result = filterByAttr(kvs[i], result)
		}
	}

	//包装结果
	nodeSelectors := make([]*NodeSelector, len(result))
	for i := 0; i < len(result); i++ {
		nodeSelectors[i] = NewNodeSelector(result[i])
	}

	return nodeSelectors, nil
}

type KV struct {
	key, value string
}

func findChildrenByCondition(node *html.Node, result *[]*html.Node, condition Condition) {
	if node.FirstChild == nil {
		return
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling  {
		if condition.Satisfy(child) {
			*result = append(*result, child)
		}

		if child.FirstChild != nil {
			findChildrenByCondition(child, result, condition)
		}
	}
}

type Condition interface {
	Satisfy(node *html.Node) bool
}

type TagCondition struct {
	TagName string
}

func NewTagCondition(kv KV) *TagCondition {
	if kv.key != "!" {
		return nil
	}

	return &TagCondition{TagName:kv.value}
}

func (tc *TagCondition) Satisfy(node *html.Node) bool {
	if node.Type == html.ElementNode && node.Data == tc.TagName {
		return true
	}
	return false
}


type AttrCondition struct {
	key, value string
}

func NewAttrCondition(kv KV)  *AttrCondition{
	return &AttrCondition{key:kv.key, value:kv.value}
}

func (ac *AttrCondition) Satisfy(node *html.Node) bool {
	if len(node.Attr) > 0 {
		for _, attr := range node.Attr {
			if attr.Key == ac.key && attr.Val == ac.value {
				return true
			}
		}
	}

	return false
}

func filterByTagName(tagName string, nodes []*html.Node)  []*html.Node{
	result := make([]*html.Node, 0)
	for _, node := range nodes {
		if node.Type == html.ElementNode && node.Data == tagName {
			result = append(result, node)
		}
	}

	return result
}

func filterByAttr(kv KV, nodes []*html.Node)  []*html.Node{
	result := make([]*html.Node, 0)
	key := kv.key
	value := kv.value
	for _, node := range nodes {
		if len(node.Attr) == 0 {
			continue
		}

		for _, attr := range node.Attr {
			if attr.Key == key && attr.Val == value {
				result = append(result, node)
			}
		}
	}

	return result
}


/**
func (ns *NodeSelector) getChildrenByTagName(tagName string) []*NodeSelector{
	children := make([]*html.Node, 0)
	findChildrenByTagName(ns.node, tagName, &children)
	result := make([]*NodeSelector, len(children))
	for i := 0; i < len(children); i++  {
		result[i] = NewNodeSelector(children[i])
	}
	return result
}

func findChildrenByTagName(node *html.Node,tagName string, result *[]*html.Node) {
	if node.FirstChild == nil {
		return
	}
	for child := node.FirstChild;child != nil ; child = child.NextSibling  {
		if child.Type == html.ElementNode && child.Data == tagName {
			*result = append(*result, child)
		}
		if child.FirstChild != nil {
			findChildrenByTagName(child, tagName, result)
		}
	}

}

func (ns *NodeSelector) getChildrenByAttr(key, value string) []*NodeSelector {
	children := make([]*html.Node, 0)
	findChildrenByAttr(ns.node, key, value, &children)
	result := make([]*NodeSelector, len(children))
	for i := 0; i < len(children); i++  {
		result[i] = NewNodeSelector(children[i])
	}
	return result
}

func findChildrenByAttr(node *html.Node,key string, value string, result *[]*html.Node) {
	if node.FirstChild == nil {
		return
	}
	for child := node.FirstChild;child != nil ; child = child.NextSibling  {
		if len(child.Attr) > 0 {
			for _, kv := range child.Attr {
				if kv.Key == key && kv.Val == value {
					*result = append(*result, child)
				}
			}
		}
		if child.FirstChild != nil {
			findChildrenByAttr(child, key, value, result)
		}
	}
}



*/


