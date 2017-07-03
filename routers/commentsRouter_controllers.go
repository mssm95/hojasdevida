package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ContactoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EscalafonEquivalenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:HistoricoEscalafonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ModalidadXInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelEscalafonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelIdiomaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:PersonaIdiomaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProcesoDisciplinarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoEstadoAdministrativoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:VinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
