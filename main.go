package main

type MobilePhone struct {
	Camara Camera
}

func (r *MobilePhone) AbrirCamara() {
	r.Camara.AccederCamara("Default")
}

func (r *MobilePhone) CapturaPhoto() Photo {
	return r.Camara.Capturar()
}

type Camera struct {
	modo string
}

func (r *Camera) AccederCamara(modo string) {
	r.modo = modo
}

func (r *Camera) Capturar() Photo {
	var foto Photo
	switch r.modo {
	case "Default":
		foto.PixelesAncho = 1920
		foto.PixelesAlto = 1080
	}
	return foto
}

type Photo struct {
	PixelesAncho int
	PixelesAlto  int
}
