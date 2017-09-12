package tree

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

//Tree 树
type Tree struct {
	Type      string
	TypeOrder int
	Title     string
	Key       string
	Order     int
	ParentKey string
	Parent    *Tree `json:"-"`
	Child     []Tree
	Value     Node
}

type treejson struct {
	Type      string
	TypeOrder int
	Title     string
	Key       string
	Order     int
	ParentKey string
	Parent    *treejson `json:"-"`
	Child     []treejson
}

//Node 树节点接口
type Node interface {
	TypeName() string
	TypeOrder() int
	Title() string
	Key() string
	OrderNumber() int
	ParentKey() string
	GetParent() (Node, error)
}

//NewTree 创建一个新树
func NewTree(node Node) *Tree {
	t := Tree{
		Type:      node.TypeName(),
		TypeOrder: node.TypeOrder(),
		Title:     node.Title(),
		Key:       node.Key(),
		Order:     node.OrderNumber(),
		ParentKey: node.ParentKey(),
		Child:     make([]Tree, 0),
		Value:     node,
	}
	return &t
}

//SortChild 对子节点进行排序
func (t *Tree) SortChild() {
	if t.Child == nil || len(t.Child) == 0 {
		return
	}
	sort.Slice(t.Child, func(i, j int) bool {
		if t.Child[i].TypeOrder == t.Child[j].TypeOrder {
			return t.Child[i].Order < t.Child[j].Order
		}
		return t.Child[i].TypeOrder < t.Child[j].TypeOrder
	})

}

//ToJSON tree转json
func (t *Tree) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

//ToJSONwithoutValue tree转json
func (t Tree) ToJSONwithoutValue() ([]byte, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	jsont := treejson{}
	err = json.Unmarshal(data, &jsont)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsont)

}

//WalkFunc 遍历函数
type WalkFunc func(node Tree)

//Walk 遍历
func (t *Tree) Walk(fn WalkFunc) {
	var innerfn WalkFunc
	innerfn = func(node Tree) {
		fn(node)
		if node.Child == nil {
			return
		}

		for i := 0; i < len(node.Child); i++ {
			innerfn(node.Child[i])
		}
	}
	innerfn(*t)
}

//NodeSliceToTree 由TreeNode切片或数组生成Tree
func NodeSliceToTree(slice interface{}) ([]Tree, error) {

	format := "需要的是TreeNode切片或数组，不能传入：%v\r\n"
	if slice == nil {
		return nil, fmt.Errorf(format, "nil")
	}

	v := reflect.Indirect(reflect.ValueOf(slice))
	t := v.Type()

	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {
			return nil, fmt.Errorf(format, "空数组或切片")
		}
		var trees []Tree
		trees = make([]Tree, 0)
		tempMap := make(map[string][]Tree)
		tempKeyList := make(map[string]bool)
		for i := 0; i < v.Len(); i++ {
			node, ok := v.Index(i).Interface().(Node)
			if !ok {
				return nil, fmt.Errorf(format, "未实现TreeNode接口的切片")
			}

			tnode := NewTree(node)
			tempMap[tnode.ParentKey] = append(tempMap[tnode.ParentKey], *tnode)
			tempKeyList[tnode.Key] = true
			trees = append(trees, *tnode)
		}

		//按照key进行字符排序
		sort.SliceStable(trees, func(i, j int) bool {
			ti, tj := trees[i], trees[j]
			return ti.Key < tj.Key
		})
		root := &Tree{}
		root.Child = make([]Tree, 0)

		//建立父子关系，找不到父的，是顶级
		for i := 0; i < len(trees); i++ {
			if !tempKeyList[trees[i].ParentKey] {
				root.Child = append(root.Child, trees[i])
			}
		}

		//建立树
		root.SortChild()
		var fn func(r *Tree)

		fn = func(r *Tree) {
			child, has := tempMap[r.Key]
			if has {
				r.Child = child
				delete(tempMap, r.Key)
				if r.Child != nil && len(r.Child) > 0 {
					r.SortChild()
					for i := 0; i < len(r.Child); i++ {
						r.Child[i].Parent = r
						fn(&r.Child[i])
					}
				}
			}
		}
		for i := 0; i < len(root.Child); i++ {
			fn(&root.Child[i])
		}
		return root.Child, nil
	default:
		return nil, fmt.Errorf(format, t)
	}

}
