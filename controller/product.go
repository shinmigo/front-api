package controller

import (
	"goshop/front-api/filter"
)

var productFilter *filter.Product

type Product struct {
	Base
}

func (m *Product) Initialise() {
	productFilter = filter.NewProduct(m.Context)
}

func (m *Product) Index() {
	str, err := productFilter.Index()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Product) Detail() {
	str, err := productFilter.Detail()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

// 获取商品标签列表
func (m *Product) Tag() {
	str, err := productFilter.Tag()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}
