package index

import (
	"fmt"
	"github.com/Mintegral-official/juno/datastruct"
	"github.com/Mintegral-official/juno/document"
	"github.com/Mintegral-official/juno/helpers"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var doc1 = &document.DocInfo{
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
		{
			Name:      "field1",
			IndexType: 2,
			Value:     "1",
			ValueType: document.StringFieldType,
		},
	},
}

var doc2 = &document.DocInfo{
	Id: 1,
	Fields: []*document.Field{
		{
			Name:      "field1",
			IndexType: 0,
			Value:     "1",
			ValueType: document.StringFieldType,
		},
		{
			Name:      "field2",
			IndexType: 1,
			Value:     "2",
			ValueType: document.StringFieldType,
		},
		{
			Name:      "field1",
			IndexType: 0,
			Value:     "1",
			ValueType: document.StringFieldType,
		},
	},
}

var doc3 = &document.DocInfo{
	Id: 2,
	Fields: []*document.Field{
		{
			Name:      "field1",
			IndexType: 0,
			Value:     "1",
			ValueType: document.StringFieldType,
		},
		{
			Name:      "field2",
			IndexType: 0,
			Value:     "2",
			ValueType: document.StringFieldType,
		},
		{
			Name:      "field1",
			IndexType: 1,
			Value:     1,
			ValueType: document.IntFieldType,
		},
	},
}

func TestNewIndex(t *testing.T) {
	Convey("NewIndex", t, func() {
		So(NewIndex("index"), ShouldNotBeNil)
	})

	Convey("Add", t, func() {
		index := NewIndex("index")
		So(index.Add(nil), ShouldEqual, helpers.DocumentError)
		So(index.Add(doc1), ShouldBeNil)
		So(index.Add(doc2), ShouldBeNil)
		So(index.Add(doc3), ShouldBeNil)
		if1 := index.GetInvertedIndex().Iterator("field1", "1")
		c := 0
		for if1.HasNext() {
			if if1.Current() != nil {
				c++
			}
			if1.Next()
		}
		So(c, ShouldEqual, 3)

		if2 := index.invertedIndex.Iterator("field2", "2")
		c = 0
		for if2.HasNext() {
			if if2.Current() != nil {
				c++
			}
			if2.Next()
		}
		So(c, ShouldEqual, 2)
		sf1 := index.GetStorageIndex().Iterator("field1")
		c = 0
		for sf1.HasNext() {
			if sf1.Current() != nil {
				c++
			}
			sf1.Next()
		}
		So(c, ShouldEqual, 2)
		sf2 := index.storageIndex.Iterator("field2")
		c = 0
		for sf2.HasNext() {
			if sf2.Current() != nil {
				c++
			}
			sf2.Next()
		}
		So(c, ShouldEqual, 1)
		So(len(*index.GetBitMap()), ShouldEqual, 32768)
		So(index.GetCampaignMap(), ShouldNotBeNil)
		So(index.GetDataType("field1"), ShouldEqual, 1)
		So(index.GetDataType("field2"), ShouldEqual, 3)
	})

	Convey("Del", t, func() {
		index := NewIndex("index")
		So(index.Add(nil), ShouldEqual, helpers.DocumentError)
		So(index.Add(doc1), ShouldBeNil)
		So(index.Add(doc2), ShouldBeNil)
		So(index.Add(doc3), ShouldBeNil)
		index.Del(doc1)
		if1 := index.GetInvertedIndex().Iterator("field1", "1")
		c := 0
		for if1.HasNext() {
			if if1.Current() != nil {
				c++
			}
			if1.Next()
		}
		So(c, ShouldEqual, 2)

		if2 := index.invertedIndex.Iterator("field2", "2")
		c = 0
		for if2.HasNext() {
			if if2.Current() != nil {
				c++
			}
			if2.Next()
		}
		So(c, ShouldEqual, 1)
		sf1 := index.GetStorageIndex().Iterator("field1")
		c = 0
		for sf1.HasNext() {
			if sf1.Current() != nil {
				c++
			}
			sf1.Next()
		}
		So(c, ShouldEqual, 1)
		sf2 := index.storageIndex.Iterator("field2")
		c = 0
		for sf2.HasNext() {
			if sf2.Current() != nil {
				c++
			}
			sf2.Next()
		}
		So(c, ShouldEqual, 1)
		So(len(*index.GetBitMap()), ShouldEqual, 32768)
		So(index.GetCampaignMap(), ShouldNotBeNil)
		So(index.DebugInfo(), ShouldNotBeNil)
		So(index.GetDataType("field1"), ShouldEqual, 1)
		So(index.GetDataType("field2"), ShouldEqual, 3)
	})
}

func f1(a interface{}) interface{} {
	return a.(bool)
}

func TestInterface(t *testing.T) {
	var a interface{} = true
	fmt.Println(f1(a))
}

func TestNewIndex2(t *testing.T) {

	//Convey("mongoIndex", t, func() {
	//	mon, err := model.NewMongo(cfg)
	//pkgNames := make([]string, 10)
	//osVersionCodeV2 := ""
	//timestamp := -1
	//category := -1
	//networkId := -1
	//directMarket := -1
	//networkType := -1
	//So(mon, ShouldBeNil)
	//So(err, ShouldNotBeNil)
	// fmt.Println(mon)
	//r, e := mon.Find()
	//So(e, ShouldNotBeNil)
	//fmt.Println(e)
	//for i := 0; i < len(r); i++ {
	//	if !helpers.In(r[i].PackageName, pkgNames) {
	//		continue
	//	}
	//	if !r[i].IsSSPlatform() {
	//		continue
	//	}
	//	if int(*r[i].AdvertiserId) == 919 || int(*r[i].AdvertiserId) == 976 {
	//		continue
	//	}
	//	if !helpers.In(osVersionCodeV2, r[i].OsVersionMinV2, r[i].OsVersionMaxV2) {
	//		continue
	//	}
	//	if !helpers.In(timestamp, r[i].StartTime, r[i].EndTime) {
	//		continue
	//	}
	//	if !helpers.Equal(category, "UNKNOWN") || !helpers.Equal(category, int(*r[i].Category)) {
	//		continue
	//	}
	//	if !helpers.Equal(networkId, r[i].Network) {
	//		continue
	//	}
	//	if !helpers.Equal(directMarket, "UNKNOWN") || !helpers.Equal(directMarket, "NO_LIMIT") ||
	//		!(helpers.Equal(directMarket, "NO_APK") && (helpers.Equal(r[i].CampaignType, "GooglePlay") ||
	//			helpers.Equal(r[i].CampaignType, "OTHER"))) || !(helpers.Equal(directMarket, "ONLY_APK") &&
	//		helpers.Equal(r[i].CampaignType, "APK")) || !(helpers.Equal(directMarket, "ONLY_GP") &&
	//		helpers.Equal(r[i].CampaignType, "GooglePlay")) {
	//		continue
	//	}
	//	if !helpers.In("UNKNOWN", r[i].NetWorkType) || !helpers.In(networkType, r[i].NetWorkType) ||
	//		!(helpers.In("RECALL_IF_UNKNOWN", r[i].NetWorkType) && helpers.Equal(networkType, "UNKNOWN")) {
	//		continue
	//	}
	//
	//}
	//})

}

//func TestIndexer_Del(t *testing.T) {
//	var a interface{} = 100
//	var b interface{} = 100
//	fmt.Println(uintptr(unsafe.Pointer(&a)) == uintptr(unsafe.Pointer(&b)))
//	fmt.Println(unsafe.Pointer(&a))
//	fmt.Println(unsafe.Pointer(&b))
//
//}

//func TestNewStorageIndexer2(t *testing.T) {
//	var s = make([]int, 10, 100)
//	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
//	fmt.Println(Len, len(s))
//
//	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
//	fmt.Println(Cap, cap(s))
//
//	fmt.Println(unsafe.Sizeof(s))
//	fmt.Println(unsafe.Alignof(s))
//	fmt.Println(unsafe.Sizeof([]int{}))
//	fmt.Println(unsafe.Alignof([]int{}))
//}

//func TestIndexer_Add(t *testing.T) {
//	type slice struct {
//		array unsafe.Pointer
//		len   int
//		cap   int
//	}
//	s := &slice{
//		array: unsafe.Pointer(&[]int{}),
//		len:   10,
//		cap:   100,
//	}
//	fmt.Println(unsafe.Sizeof(s.array))
//	fmt.Println(unsafe.Sizeof(s.len))
//	fmt.Println(unsafe.Sizeof(s.cap))
//	fmt.Println(unsafe.Sizeof(s))
//	fmt.Println(unsafe.Alignof(s))
//	fmt.Println(unsafe.Alignof(s.array))
//	fmt.Println(unsafe.Alignof(s.cap))
//	fmt.Println(unsafe.Alignof(s.len))
//	fmt.Println(unsafe.Offsetof(s.array))
//	fmt.Println(unsafe.Offsetof(s.len))
//	fmt.Println(unsafe.Offsetof(s.cap))
//
//	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
//	fmt.Println(Len, s.len)
//
//	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
//	fmt.Println(Cap, s.cap)
//
//	fmt.Println(unsafe.Alignof(&[]int{}))
//	fmt.Println(unsafe.Alignof(true))
//}

func TestStorageIndexer_Add(t *testing.T) {
	var a = &document.DocInfo{
		Id: 0,
		Fields: []*document.Field{
			{
				Name:      "f1",
				IndexType: 0,
				Value:     []int64{1, 2, 3},
				ValueType: document.SliceFieldType,
			},
			{
				Name:      "f2",
				IndexType: 1,
				Value:     []float64{1.1, 2.2, 3.4},
				ValueType: document.SliceFieldType,
			},
			{
				Name:      "f1",
				IndexType: 0,
				Value:     []int64{1, 22, 33},
				ValueType: document.SliceFieldType,
			},
		},
	}
	Convey("add", t, func() {
		index := NewIndex("index")
		_ = index.Add(a)
		idx := index.invertedIndex.Iterator("f1", "1")
		So(idx.HasNext(), ShouldBeTrue)
		c := 0
		for idx.HasNext() {
			if idx.Current() != nil {
				c++
			}
			idx.Next()
		}
		So(c, ShouldEqual, 1)
		sto := index.storageIndex.Iterator("f2")
		So(sto.HasNext(), ShouldBeTrue)
		c = 0
		for sto.HasNext() {
			if sto.Current().(*datastruct.Element).Value() != nil {
				c++
			}
			sto.Next()
		}
		So(c, ShouldEqual, 1)
	})

}
