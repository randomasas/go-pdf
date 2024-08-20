package gopdffork

type ICacheColorText interface {
	ICacheContent
	equal(obj ICacheColorText) bool
}
