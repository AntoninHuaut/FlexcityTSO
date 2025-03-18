package controller

type TsoController interface {
}

func NewTsoController() TsoController {
	return tsoController{}
}

type tsoController struct {
}
