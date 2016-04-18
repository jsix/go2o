/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-08 10:53
 * description :
 * history :
 */

package sale

import (
	"bytes"
	"go2o/src/core/domain/interface/sale"
	"strconv"
	"strings"
	"go2o/src/core/infrastructure/domain"
	"fmt"
)

var _ sale.ICategory = new(Category)

type Category struct {
	_value           *sale.ValueCategory
	_rep             sale.ISaleRep
	_parentIdChanged bool
	_childIdArr      []int
	_opt domain.IOptionStore
}

func newCategory(saleRep sale.ISaleRep, v *sale.ValueCategory) sale.ICategory {
	return &Category{
		_value: v,
		_rep:   saleRep,
	}
}

func (this *Category) GetDomainId() int {
	return this._value.Id
}

func (this *Category) GetValue() *sale.ValueCategory {
	return this._value
}

func (this *Category)GetOption()domain.IOptionStore {
	if this._opt == nil {
		opt := newCategoryOption(this)
		if err := opt.Stat();err != nil {
			opt.Set(sale.C_OptionViewName, &domain.Option{
				Key:sale.C_OptionViewName,
				Type:domain.OptionTypeString,
				Must:false,
				Title:"显示页面",
				Value:"list.html",
			})
			opt.Set(sale.C_OptionDescribe, &domain.Option{
				Key:sale.C_OptionDescribe,
				Type:domain.OptionTypeString,
				Must:false,
				Title:"描述",
				Value:"",
			})
			opt.Flush()
		}
		this._opt = opt
	}
	return this._opt
}

func (this *Category) SetValue(v *sale.ValueCategory) error {
	val := this._value
	if val.Id == v.Id {
		val.Description = v.Description
		val.Enabled = v.Enabled
		val.Name = v.Name
		val.OrderIndex = v.OrderIndex
		if val.ParentId != v.ParentId {
			this._parentIdChanged = true
			val.ParentId = v.ParentId
		} else {
			this._parentIdChanged = false
		}
	}
	return nil
}

// 获取子栏目的编号
func (this *Category) GetChildId() []int {
	if this._childIdArr == nil {
		childCats := this._rep.GetChildCategories(this._value.PartnerId, this.GetDomainId())
		this._childIdArr = make([]int, len(childCats))
		for i, v := range childCats {
			this._childIdArr[i] = v.Id
		}
	}
	return this._childIdArr
}

func (this *Category) Save() (int, error) {
	id, err := this._rep.SaveCategory(this._value)
	if err == nil {
		this._value.Id = id
		if len(this._value.Url) == 0 || (this._parentIdChanged &&
			strings.HasPrefix(this._value.Url, "/c-")) {
			this._value.Url = this.getAutomaticUrl(this._value.PartnerId, id)
			this._parentIdChanged = false
			return this.Save()
		}
	}
	return id, err
}

func (this *Category) getAutomaticUrl(partnerId, id int) string {
	var relCategories []*sale.ValueCategory = this._rep.GetRelationCategories(partnerId, id)
	var buf *bytes.Buffer = bytes.NewBufferString("/c")
	var l int = len(relCategories)
	for i := l; i > 0; i-- {
		buf.WriteString("-" + strconv.Itoa(relCategories[i-1].Id))
	}
	buf.WriteString(".htm")
	return buf.String()
}

var _ domain.IOptionStore = new(categoryOption)
type categoryOption struct{
	domain.IOptionStore
	_partnerId int
	_c *Category
}

func newCategoryOption(c *Category)domain.IOptionStore {
	i := fmt.Sprintf("conf/%d/option/c/%d",c.GetValue().PartnerId, c.GetDomainId())
	return &categoryOption{
		_partnerId:c.GetValue().ParentId,
		_c : c,
		IOptionStore:domain.NewOptionStoreWrapper(i),
	}
}