package main

import (
	"testing"
)

func Test_sacarFotoTelefono(t *testing.T) {
	var telefono MobilePhone
	telefono.AbrirCamara()
	foto := telefono.CapturaPhoto()

	width := foto.PixelesAncho == 1920
	height := foto.PixelesAlto == 1080

	if !(width && height) {
		t.Errorf("el tamaño no es el correcto")
	}
}
