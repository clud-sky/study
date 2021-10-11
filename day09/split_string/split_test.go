package split_string

import (
	"reflect"
	"testing"
)

// 切割字符串
func TestSplit(t *testing.T) {
	ret := Split("abcbdf", "b")
	want := []string{"a", "c", "df"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败了
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("a:c:df", ":")
	want := []string{"a", "c", "df"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败了
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

func Test3Split(t *testing.T) {
	ret := Split("ab1cb1df", "b1")
	want := []string{"a", "c", "df"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败了
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}
