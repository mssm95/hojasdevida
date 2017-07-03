package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ModalidadGrado struct {
	Id                      int          `orm:"column(id);pk"`
	NombreModalidad         string       `orm:"column(nombre_modalidad)"`
	InstitucionId           *Institucion `orm:"column(institucion_id);rel(fk)"`
	FechaDato               time.Time    `orm:"column(fecha_dato);type(date);null"`
	Vigente                 bool         `orm:"column(vigente);null"`
	IdModalidadxinstitucion float64      `orm:"column(id_modalidadxinstitucion)"`
}

func (t *ModalidadGrado) TableName() string {
	return "modalidad_grado"
}

func init() {
	orm.RegisterModel(new(ModalidadGrado))
}

// AddModalidadGrado insert a new ModalidadGrado into database and returns
// last inserted Id on success.
func AddModalidadGrado(m *ModalidadGrado) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetModalidadGradoById retrieves ModalidadGrado by Id. Returns error if
// Id doesn't exist
func GetModalidadGradoById(id int) (v *ModalidadGrado, err error) {
	o := orm.NewOrm()
	v = &ModalidadGrado{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllModalidadGrado retrieves all ModalidadGrado matches certain condition. Returns empty list if
// no records exist
func GetAllModalidadGrado(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ModalidadGrado))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ModalidadGrado
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateModalidadGrado updates ModalidadGrado by Id and returns error if
// the record to be updated doesn't exist
func UpdateModalidadGradoById(m *ModalidadGrado) (err error) {
	o := orm.NewOrm()
	v := ModalidadGrado{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteModalidadGrado deletes ModalidadGrado by Id and returns error if
// the record to be deleted doesn't exist
func DeleteModalidadGrado(id int) (err error) {
	o := orm.NewOrm()
	v := ModalidadGrado{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ModalidadGrado{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
