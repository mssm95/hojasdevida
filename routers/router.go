// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hojasdevida/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/modalidad_x_institucion",
			beego.NSInclude(
				&controllers.ModalidadXInstitucionController{},
			),
		),

		beego.NSNamespace("/escalafon_equivalente",
			beego.NSInclude(
				&controllers.EscalafonEquivalenteController{},
			),
		),

		beego.NSNamespace("/distincion",
			beego.NSInclude(
				&controllers.DistincionController{},
			),
		),

		beego.NSNamespace("/experiencia_docente",
			beego.NSInclude(
				&controllers.ExperienciaDocenteController{},
			),
		),

		beego.NSNamespace("/estado_administrativo",
			beego.NSInclude(
				&controllers.EstadoAdministrativoController{},
			),
		),

		beego.NSNamespace("/tipo_vinculacion",
			beego.NSInclude(
				&controllers.TipoVinculacionController{},
			),
		),

		beego.NSNamespace("/produccion_academica",
			beego.NSInclude(
				&controllers.ProduccionAcademicaController{},
			),
		),

		beego.NSNamespace("/cursos",
			beego.NSInclude(
				&controllers.CursosController{},
			),
		),

		beego.NSNamespace("/tipo_produccion",
			beego.NSInclude(
				&controllers.TipoProduccionController{},
			),
		),

		beego.NSNamespace("/investigacion",
			beego.NSInclude(
				&controllers.InvestigacionController{},
			),
		),

		beego.NSNamespace("/dato_produccion",
			beego.NSInclude(
				&controllers.DatoProduccionController{},
			),
		),

		beego.NSNamespace("/evaluador",
			beego.NSInclude(
				&controllers.EvaluadorController{},
			),
		),

		beego.NSNamespace("/nivel_escalafon",
			beego.NSInclude(
				&controllers.NivelEscalafonController{},
			),
		),

		beego.NSNamespace("/tipo_investigacion",
			beego.NSInclude(
				&controllers.TipoInvestigacionController{},
			),
		),

		beego.NSNamespace("/modalidad_grado",
			beego.NSInclude(
				&controllers.ModalidadGradoController{},
			),
		),

		beego.NSNamespace("/nivel_formacion",
			beego.NSInclude(
				&controllers.NivelFormacionController{},
			),
		),

		beego.NSNamespace("/tipo_contacto",
			beego.NSInclude(
				&controllers.TipoContactoController{},
			),
		),

		beego.NSNamespace("/tipo_estado_administrativo",
			beego.NSInclude(
				&controllers.TipoEstadoAdministrativoController{},
			),
		),

		beego.NSNamespace("/vinculacion",
			beego.NSInclude(
				&controllers.VinculacionController{},
			),
		),

		beego.NSNamespace("/persona_idioma",
			beego.NSInclude(
				&controllers.PersonaIdiomaController{},
			),
		),

		beego.NSNamespace("/programa",
			beego.NSInclude(
				&controllers.ProgramaController{},
			),
		),

		beego.NSNamespace("/tipo_dedicacion",
			beego.NSInclude(
				&controllers.TipoDedicacionController{},
			),
		),

		beego.NSNamespace("/proceso_disciplinario",
			beego.NSInclude(
				&controllers.ProcesoDisciplinarioController{},
			),
		),

		beego.NSNamespace("/contacto",
			beego.NSInclude(
				&controllers.ContactoController{},
			),
		),

		beego.NSNamespace("/escalafon",
			beego.NSInclude(
				&controllers.EscalafonController{},
			),
		),

		beego.NSNamespace("/criterio_evaluacion",
			beego.NSInclude(
				&controllers.CriterioEvaluacionController{},
			),
		),

		beego.NSNamespace("/subtipo_produccion",
			beego.NSInclude(
				&controllers.SubtipoProduccionController{},
			),
		),

		beego.NSNamespace("/historico_escalafon",
			beego.NSInclude(
				&controllers.HistoricoEscalafonController{},
			),
		),

		beego.NSNamespace("/formacion_academica",
			beego.NSInclude(
				&controllers.FormacionAcademicaController{},
			),
		),

		beego.NSNamespace("/dato_subtipo",
			beego.NSInclude(
				&controllers.DatoSubtipoController{},
			),
		),

		beego.NSNamespace("/experiencia_laboral",
			beego.NSInclude(
				&controllers.ExperienciaLaboralController{},
			),
		),

		beego.NSNamespace("/idioma",
			beego.NSInclude(
				&controllers.IdiomaController{},
			),
		),

		beego.NSNamespace("/dato",
			beego.NSInclude(
				&controllers.DatoController{},
			),
		),

		beego.NSNamespace("/institucion",
			beego.NSInclude(
				&controllers.InstitucionController{},
			),
		),

		beego.NSNamespace("/opcion_dato",
			beego.NSInclude(
				&controllers.OpcionDatoController{},
			),
		),

		beego.NSNamespace("/titulo",
			beego.NSInclude(
				&controllers.TituloController{},
			),
		),

		beego.NSNamespace("/nivel_idioma",
			beego.NSInclude(
				&controllers.NivelIdiomaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
