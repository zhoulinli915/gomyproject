package models

import (
	"github.com/astaxie/beego/orm"
	//"github.com/george518/PPGo_Job/libs"
	"strconv"
	"github.com/astaxie/beego"
	"time"
)

type Order struct {
	Id        int64
	Oid       string
	Uid       int64
	Types     int64
	Price     float64
	Amount    float64
	Status    int64
	CoinTypes string
	Created   int64
}

func init() {
	orm.RegisterModel(new(Order))
}

func (u *Order) TableName() string {
	return "order"
}

//添加
func OrderAdd(this *Order ) (int64, error) {
	id, _ := orm.NewOrm().Insert(this)
	//单号生成
	oid := OrderNum(id)
	order := new(Order)
	order.Id = id
	order.Oid = oid
	return orm.NewOrm().Update(order, "oid")
}


func OrderNum(id int64) string {
	//beego.Date(time.Unix(v.Created, 0), "Y-m-d H:i:s")
	//time.Unix(created, 0).Format("2006-01-02 15:04:05")
	s := strconv.FormatInt(id, 10)
	date := beego.Date(time.Now(), "YmdHis")
	str := date + s
	return str
}

//列表
func OrderGetList(page, pageSize int, filters []interface{}, orderMap ... string) ([]*Order, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Order, 0)
	query := orm.NewOrm().QueryTable(new(Order).TableName())
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	for _, v := range orderMap {
		query = query.OrderBy(v)
	}
	query.Limit(pageSize, offset).All(&list)

	return list, total
}
