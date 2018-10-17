package cameras

import (
	"github.com/tokkenno/seed/core"
	"github.com/tokkenno/seed/core/math"
)

type Camera struct {
	core.Object3

	matrixWorldInverse      *math.Matrix4
	projectionMatrix        *math.Matrix4
	projectionMatrixInverse *math.Matrix4
}

func (camera *Camera) GetMatrixWorldInverse() *math.Matrix4 {
	return camera.matrixWorldInverse
}

func (camera *Camera) GetProjectionMatrix() *math.Matrix4 {
	return camera.projectionMatrix
}

func (camera *Camera) Copy(source *Camera, recursive bool) {
	camera.Object3.Copy(&source.Object3, recursive)

	camera.matrixWorldInverse.Copy(source.matrixWorldInverse)
	camera.projectionMatrix.Copy(source.projectionMatrix)
	camera.projectionMatrixInverse.Copy(source.projectionMatrixInverse)
}

func (camera *Camera) Clone() *Camera {
	newCamera := new(Camera)
	newCamera.Copy(camera, true)
	return newCamera
}

func (camera *Camera) UpdateMatrixWorld(force bool) {
	camera.Object3.UpdateMatrixWorld(force)
	camera.matrixWorldInverse.Inverse(camera.GetMatrixWorld(), false)
}

func (camera *Camera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	camera.UpdateMatrixWorld(true)
	e := camera.GetMatrixWorld().GetElements()

	target.Set(-e[8], -e[9], -e[10])
	target.Normalize()
	return target
}
