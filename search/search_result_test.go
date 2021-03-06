package search

import (
	"github.com/Mintegral-official/juno/check"
	"github.com/Mintegral-official/juno/document"
	"github.com/Mintegral-official/juno/index"
	"github.com/Mintegral-official/juno/operation"
	"github.com/Mintegral-official/juno/query"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestSearcher(t *testing.T) {
	var doc1 = &document.DocInfo{
		Id: 10,
		Fields: []*document.Field{
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "1",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldNeme",
				IndexType: 0,
				Value:     "2",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "3",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "4",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "5",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "6",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 1,
				Value:     3,
				ValueType: document.StringFieldType,
			},
		},
	}
	var doc2 = &document.DocInfo{
		Id: 30,
		Fields: []*document.Field{
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "1",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldNeme",
				IndexType: 0,
				Value:     "2",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "3",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "5",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 1,
				Value:     3,
				ValueType: document.StringFieldType,
			},
		},
	}
	var doc3 = &document.DocInfo{
		Id: 40,
		Fields: []*document.Field{
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "1",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldNeme",
				IndexType: 0,
				Value:     "2",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "3",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "4",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "5",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "6",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 1,
				Value:     3,
				ValueType: document.StringFieldType,
			},
		},
	}
	var doc4 = &document.DocInfo{
		Id: 60,
		Fields: []*document.Field{
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "1",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldNeme",
				IndexType: 0,
				Value:     "2",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "3",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "4",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "6",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 1,
				Value:     3,
				ValueType: document.StringFieldType,
			},
		},
	}
	var doc5 = &document.DocInfo{
		Id: 100,
		Fields: []*document.Field{
			{
				Name:      "fieldName",
				IndexType: 0,
				Value:     "1",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "fieldName",
				IndexType: 1,
				Value:     3,
				ValueType: document.StringFieldType,
			},
		},
	}
	ss := index.NewIndex("")
	Convey("Search Test", t, func() {
		_ = ss.Add(doc1)
		_ = ss.Add(doc2)
		_ = ss.Add(doc3)
		_ = ss.Add(doc4)
		_ = ss.Add(doc5)
		s1 := ss.GetInvertedIndex()
		s2 := ss.GetStorageIndex()
		q := query.NewAndQuery([]query.Query{
			query.NewTermQuery(s1.Iterator("fieldName", "1")),
			query.NewTermQuery(s1.Iterator("fieldNeme", "2")),
			query.NewTermQuery(s1.Iterator("fieldName", "3")),
			query.NewAndQuery([]query.Query{
				query.NewTermQuery(s1.Iterator("fieldName", "1")),
				query.NewTermQuery(s1.Iterator("fieldNeme", "2")),
				query.NewTermQuery(s1.Iterator("fieldName", "3")),
				query.NewOrQuery([]query.Query{
					query.NewTermQuery(s1.Iterator("fieldName", "1")),
					query.NewTermQuery(s1.Iterator("fieldNeme", "2")),
					query.NewTermQuery(s1.Iterator("fieldName", "3")),
					query.NewTermQuery(s1.Iterator("fieldName", "4")),
				}, []check.Checker{
					check.NewChecker(s2.Iterator("fieldName"), 2, operation.NE, nil, false),
					check.NewAndChecker([]check.Checker{
						check.NewChecker(s2.Iterator("fieldName"), 2, operation.EQ, nil, false),
						check.NewChecker(s2.Iterator("fieldName"), 3, operation.EQ, nil, false),
					}),
				}),
			}, []check.Checker{
				check.NewInChecker(s2.Iterator("fieldName"), []int{2, 3, 4, 5}, nil, false),
				check.NewOrChecker([]check.Checker{
					check.NewChecker(s2.Iterator("fieldName"), 2, operation.NE, nil, false),
					check.NewChecker(s2.Iterator("fieldName"), 3, operation.EQ, nil, false),
				}),
			}),
		}, []check.Checker{
			check.NewInChecker(s2.Iterator("fieldName"), []int{2, 3, 4}, nil, false),
			check.NewOrChecker([]check.Checker{
				check.NewChecker(s2.Iterator("fieldName"), 2, operation.GT, nil, false),
				check.NewChecker(s2.Iterator("fieldName"), 3, operation.EQ, nil, false),
			}),
		})

		se := NewSearcher()
		se.Search(ss, q)
		testCase := []document.DocId{10, 30, 40, 60}
		for i, expect := range testCase {
			So(se.Docs[i], ShouldEqual, expect)
		}
		v, e := q.Current()
		q.Next()
		So(v, ShouldEqual, 0)
		So(e, ShouldNotBeNil)
	})

}

func TestNewSearcher_Inc_Index(t *testing.T) {

	var doc4 = &document.DocInfo{
		Id: 0,
		Fields: []*document.Field{
			{
				Name:      "field1",
				IndexType: 1,
				Value:     1,
				ValueType: document.IntFieldType,
			},
			{
				Name:      "field2",
				IndexType: 0,
				Value:     "2",
				ValueType: document.StringFieldType,
			},
		},
	}

	var doc5 = &document.DocInfo{
		Id: 0,
		Fields: []*document.Field{
			{
				Name:      "field1",
				IndexType: 1,
				Value:     10,
				ValueType: document.IntFieldType,
			},
			{
				Name:      "field2",
				IndexType: 0,
				Value:     "20",
				ValueType: document.StringFieldType,
			},
			{
				Name:      "field2",
				IndexType: 0,
				Value:     "200",
				ValueType: document.StringFieldType,
			},
		},
	}
	Convey("search inc index", t, func() {
		idx := index.NewIndex("")
		_ = idx.Add(doc4)
		q := query.NewTermQuery(idx.GetInvertedIndex().Iterator("field2", "2"))
		expectMap := [2]map[string][]string{
			{
				"field2": []string{"2"},
			},
			{
				"field1": []string{"1"},
			},
		}
		realMap := idx.GetValueById(0)
		So(reflect.DeepEqual(realMap, expectMap), ShouldBeTrue)

		s1 := NewSearcher()
		s1.Search(idx, q)
		So(s1.Docs[0], ShouldEqual, 0)

		idx.Del(doc5)
		_ = idx.Add(doc5)
		expectMap = [2]map[string][]string{
			{
				"field2": []string{"20", "200"},
			},
			{
				"field1": []string{"10"},
			},
		}
		realMap = idx.GetValueById(0)
		So(reflect.DeepEqual(realMap, expectMap), ShouldBeTrue)

		q = query.NewTermQuery(idx.GetInvertedIndex().Iterator("field2", "20"))
		s1 = NewSearcher()
		s1.Search(idx, q)
		So(s1.Docs[0], ShouldEqual, 0)
	})

}
