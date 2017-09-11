package tree

import (
	"fmt"
	"testing"
	
)

var _ Node = testNode{}

type testNode struct {
	PID int
	Name int
	
}

func (t testNode) TypeName() string {
	return ""
}
func (t testNode) TypeOrder() int {
	return 0
}
func (t testNode) Title() string {
	return fmt.Sprintf("%d",t.Name)
}
func (t testNode) Key() string {
	return fmt.Sprintf("%d",t.Name)
}
func (t testNode) OrderNumber() int {
	return 0
}
func (t testNode) ParentKey() string {
	return fmt.Sprintf("%d",t.PID)
}
func (t testNode) Parent() (Node, error) {
	return testNode{Name:t.PID}, nil
}
// func TestTree_ToJSONwithValue(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		t       Tree
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "测试转json带value",
// 			t:    Tree{Title: "aaaa", Value: testNode{Name: "hahahaha"}, Child: []Tree{{Title: "aaaa", Value: testNode{Name: "hahahaha"}}}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.t.ToJSONwithoutValue()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Tree.ToJSONwithoutValue() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			t.Log(string(got))

// 			// if !reflect.DeepEqual(got, tt.want) {
// 			// 	t.Errorf("Tree.ToJSONwithValue() = %v, want %v", got, tt.want)
// 			// }
// 		})
// 	}
// }

func TestAA(t *testing.T){
	nodes:=[]testNode{
		{0,1},
		{1,2},
		{1,3},
		{1,4},
		{2,5},
		{2,6},
		{5,7},
		{5,8},
		{5,9},

	}
 ts,err:=NodeSliceToTree(nodes)
 t.Log(err)
 t.Log(len(ts))

}