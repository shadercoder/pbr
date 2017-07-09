package main

import (
	"math"

	"github.com/hunterloftis/pbr/pbr"
)

func main() {
	scene := pbr.EmptyScene()
	camera := pbr.Camera35mm(640, 360, 0.050)
	renderer := pbr.CamRenderer(camera, 1)
	cli := pbr.CliRunner(scene, camera, renderer)

	light := pbr.Light(1500, 1500, 1500)
	redPlastic := pbr.Plastic(1, 0, 0, 1)
	whitePlastic := pbr.Plastic(1, 1, 1, 0.8)
	// grayLambert := pbr.Lambert(0.2, 0.2, 0.2, 0.1)
	bluePlastic := pbr.Plastic(0, 0, 1, 1)
	silver := pbr.Metal(0.95, 0.93, 0.88, 1)
	gold := pbr.Metal(1.022, 0.782, 0.344, 0.9)
	// glass := pbr.Glass(1, 1, 1, 1)
	greenGlass := pbr.Glass(0.2, 1, 0.1, 0.95)

	scene.SetSky(pbr.Vector3{40, 50, 60}, pbr.Vector3{})

	scene.Add(
		pbr.UnitCube(pbr.Ident().Rot(pbr.Vector3{0, -0.25 * math.Pi, 0}).Scale(0.1, 0.1, 0.1), silver, nil),
		pbr.UnitCube(pbr.Ident().Trans(0, 0, -0.4).Rot(pbr.Vector3{0, 0.1 * math.Pi, 0}).Scale(0.1, 0.1, 0.1), gold, nil),
		pbr.UnitCube(pbr.Ident().Trans(-0.4, 0, 0.2).Rot(pbr.Vector3{0, -0.1 * math.Pi, 0}).Scale(0.1, 0.1, 0.1), bluePlastic, nil),
		pbr.UnitCube(pbr.Ident().Trans(0.175, 0.05, 0.175).Rot(pbr.Vector3{0, 0.55 * math.Pi, 0}).Scale(0.01, 0.2, 0.1), greenGlass, nil),
		pbr.UnitCube(pbr.Ident().Trans(0, -0.55, 0).Scale(1000, 1, 1000), whitePlastic, bluePlastic),
		pbr.UnitSphere(pbr.Ident().Trans(-0.2, 0.001, 0.1).Scale(0.1, 0.1, 0.1), greenGlass),
		pbr.UnitSphere(pbr.Ident().Trans(0.3, 0.05, 0).Scale(0.2, 0.2, 0.2), greenGlass),
		pbr.UnitSphere(pbr.Ident().Trans(7, 30, 6).Scale(30, 30, 30), light),
		pbr.UnitSphere(pbr.Ident().Trans(0, -0.025, 0.2).Scale(0.1, 0.05, 0.1), redPlastic),
		pbr.UnitSphere(pbr.Ident().Trans(0.45, 0.05, -0.4).Scale(0.2, 0.2, 0.2), silver),
	)

	camera.MoveTo(-0.7, 0.3, 0.9)
	camera.LookAt(-0.1, 0.001, 0.1)
	camera.Focus(-0.1, 0.001, 0.1, 4)

	cli.Render()
}
