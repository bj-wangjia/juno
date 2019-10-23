package index

type InvertedIndex interface {
	Add(id DocId)
	Del(id DocId)
	HasNext() bool
	Next() DocId
	GetGE(id DocId) DocId
}
