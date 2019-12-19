package index

import (
	"github.com/Mintegral-official/juno/datastruct"
	"github.com/Mintegral-official/juno/document"
	"github.com/Mintegral-official/juno/helpers"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInvertedIndexer_Add(t *testing.T) {
	s := NewInvertedIndexer()
	Convey("Add", t, func() {
		So(s.Add("fileName", 1), ShouldBeNil)
	})
}

func TestInvertedIndexer_Del(t *testing.T) {
	s := NewInvertedIndexer()
	Convey("Del", t, func() {
		So(s.Del("filename", 1), ShouldBeFalse)
	})
}

func TestInvertedIndexer_Iterator(t *testing.T) {
	s := NewInvertedIndexer()
	Convey("Iterator", t, func() {
		So(s.Iterator("filename"), ShouldNotBeNil)
	})
}

func TestInvertedIndexer(t *testing.T) {
	s := NewInvertedIndexer()
	sl1, _ := datastruct.NewSkipList(datastruct.DefaultMaxLevel, helpers.DocIdFunc)
	s.data.Store("fieldName1", sl1)
	s.data.Store("fieldName2", nil)
	sl2, _ := datastruct.NewSkipList(datastruct.DefaultMaxLevel, helpers.DocIdFunc)
	s.data.Store("fieldName4", sl2)
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
			// fmt.Println(a.Current())
			a.Next()
			c++
		}
		So(c, ShouldEqual, 3)
		So(s.Count(), ShouldEqual, 3)
	})

	Convey("Add", t, func() {
		So(s.Add("fieldName2", 11), ShouldEqual, helpers.ParseError)
	})

	Convey("Add", t, func() {
		So(s.Add("fieldName", 111), ShouldBeNil)
		if v, ok := s.data.Load("fieldName"); ok {
			if v1, ok := v.(*datastruct.SkipList); ok {
				So(v1.Len(), ShouldEqual, 1)
			}
		}
	})
}
