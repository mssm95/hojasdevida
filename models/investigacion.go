package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Investigacion struct {
	Id                  int                `orm:"column(id);pk"`
	PersonaId           int                `orm:"column(persona_id)"`
	NombreInvestigacion string             `orm:"column(nombre_investigacion)"`
	FechaInicio         time.Time          `orm:"column(fecha_inicio);type(date)"`
	FechaFinalizacion   time.Time          `orm:"column(fecha_finalizacion);type(date)"`
	TipoInvestigacion   string             `orm:"column(tipo_investigacion);null"`
	InstitucionId       *Institucion       `orm:"column(institucion_id);rel(fk)"`
	TipoInvestigacionId *TipoInvestigacion `orm:"column(tipo_investigacion_id);rel(fk)"`
	FechaDato           time.Time          `orm:"column(fecha_dato);type(date);null"`
	Vigente             bool               `orm:"column(vigente);null"`
}

func (t *Investigacion) TableName() string {
	return "investigacion"
}

func init() {
	orm.RegisterModel(new(Investigacion))
}

// AddInvestigacion insert a new Investigacion into database and returns
// last inserted Id on success.
func AddInvestigacion(m *Investigacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInvestigacionById retrieves Investigacion by Id. Returns error if
// Id doesn't exist
func GetInvestigacionById(id int) (v *Investigacion, err error) {
	o := orm.NewOrm()
	v = &Investigacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInvestigacion retrieves all Investigacion matches certain condition. Returns empty list if
// no records exist
func GetAllInvestigacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Investigacion))
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

	var l []Investigacion
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

// UpdateInvestigacion updates Investigacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateInvestigacionById(m *Investigacion) (err error) {
	o := orm.NewOrm()
	v := Investigacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInvestigacion deletes Investigacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInvestigacion(id int) (err error) {
	o := orm.NewOrm()
	v := Investigacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Investigacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
