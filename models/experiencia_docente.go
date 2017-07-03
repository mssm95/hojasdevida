package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ExperienciaDocente struct {
	Id                int             `orm:"column(id);pk"`
	InstitucionId     *Institucion    `orm:"column(institucion_id);rel(fk)"`
	TipoActividad     string          `orm:"column(tipo_actividad)"`
	CampoEnseñanza    string          `orm:"column(campo_enseñanza);null"`
	PersonaId         int             `orm:"column(persona_id)"`
	Validacion        bool            `orm:"column(validacion)"`
	FechaInicio       time.Time       `orm:"column(fecha_inicio);type(date)"`
	FechaFinalizacion time.Time       `orm:"column(fecha_finalizacion);type(date)"`
	TipoDedicacionId  *TipoDedicacion `orm:"column(tipo_dedicacion_id);rel(fk)"`
	FechaDato         time.Time       `orm:"column(fecha_dato);type(date);null"`
	Vigencia          bool            `orm:"column(vigencia);null"`
}

func (t *ExperienciaDocente) TableName() string {
	return "experiencia_docente"
}

func init() {
	orm.RegisterModel(new(ExperienciaDocente))
}

// AddExperienciaDocente insert a new ExperienciaDocente into database and returns
// last inserted Id on success.
func AddExperienciaDocente(m *ExperienciaDocente) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExperienciaDocenteById retrieves ExperienciaDocente by Id. Returns error if
// Id doesn't exist
func GetExperienciaDocenteById(id int) (v *ExperienciaDocente, err error) {
	o := orm.NewOrm()
	v = &ExperienciaDocente{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllExperienciaDocente retrieves all ExperienciaDocente matches certain condition. Returns empty list if
// no records exist
func GetAllExperienciaDocente(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ExperienciaDocente))
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

	var l []ExperienciaDocente
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

// UpdateExperienciaDocente updates ExperienciaDocente by Id and returns error if
// the record to be updated doesn't exist
func UpdateExperienciaDocenteById(m *ExperienciaDocente) (err error) {
	o := orm.NewOrm()
	v := ExperienciaDocente{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteExperienciaDocente deletes ExperienciaDocente by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExperienciaDocente(id int) (err error) {
	o := orm.NewOrm()
	v := ExperienciaDocente{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ExperienciaDocente{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
