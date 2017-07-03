package models

import "time"

type HistorialVinculacion struct {
	Id_RENAME         int              `orm:"column(id)"`
	VinculaciónId     float64          `orm:"column(vinculación_id)"`
	TipoVinculacionId *TipoVinculacion `orm:"column(tipo_vinculacion_id);rel(fk)"`
	FechaInicio       time.Time        `orm:"column(fecha_inicio);type(date)"`
	FechaFinal        time.Time        `orm:"column(fecha_final);type(date);null"`
}
