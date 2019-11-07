package index

import (
	"fmt"
	"github.com/Mintegral-official/juno/document"
	"github.com/Mintegral-official/juno/helpers"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSimpleInvertedIndex_Add(t *testing.T) {
	s := NewSimpleInvertedIndex()
	Convey("Add", t, func() {
        So(s.Add("fileName", 1), ShouldBeNil)
	})
}

func TestSimpleInvertedIndex_Del(t *testing.T) {
	s := NewSimpleInvertedIndex()
	Convey("Del", t, func() {
		So(s.Del("filename", 1), ShouldBeFalse)
	})
}

func TestSimpleInvertedIndex_Iterator(t *testing.T) {
	s := NewSimpleInvertedIndex()
	Convey("Iterator", t, func() {
		So(s.Iterator("filename"), ShouldBeNil)
	})
}

func TestSimpleInvertedIndex(t *testing.T) {
	s := NewSimpleInvertedIndex()
	s.data.Store("fieldName1", NewSkipListIterator(DEFAULT_MAX_LEVEL, helpers.DocIdFunc))
	s.data.Store("fieldName2", nil)
	s.data.Store("fieldName4", NewSkipListIterator(DEFAULT_MAX_LEVEL, helpers.DocIdFunc))
	Convey("Add", t, func() {
		So(s.Add("fieldName1", document.DocId(1)), ShouldBeNil)
		So(s.Add("fieldName1", document.DocId(5)), ShouldBeNil)
		So(s.Add("fieldName1", document.DocId(6)), ShouldBeNil)
		So(s.Add("fieldName1", document.DocId(7)), ShouldBeNil)
		So(s.Add("fieldName4", document.DocId(2)), ShouldBeNil)
		So(s.Del("fieldName", 1), ShouldBeFalse)
		So(s.Del("fieldName1", document.DocId(1)), ShouldBeTrue)
		a := s.Iterator("fieldName1")
		So(s.Iterator("fieldName1"), ShouldNotBeNil)
		c := 0
		for a.HasNext() {
			fmt.Println(a.Next().key)
			c++
		}
		So(c, ShouldEqual, 3)
	})

	Convey("Add", t, func() {
		So(s.Add("fieldName2", 11), ShouldEqual, helpers.PARSE_ERROR)
	})

	Convey("Add", t, func() {
		So(s.Add("fieldName", 111), ShouldBeNil)
		if v, ok := s.data.Load("fieldName"); ok {
			if v1, ok := v.(*SkipList); ok {
				So(v1.length, ShouldEqual, 1)
			}
		}
	})
}