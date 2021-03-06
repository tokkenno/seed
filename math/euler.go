package math

import (
	"github.com/nuberu/engine/event"
)

type EulerOrder int

const (
	EulerOrderXYZ     EulerOrder = 0
	EulerOrderYZX     EulerOrder = 1
	EulerOrderZXY     EulerOrder = 2
	EulerOrderXZY     EulerOrder = 3
	EulerOrderYXZ     EulerOrder = 4
	EulerOrderZYX     EulerOrder = 5
	EulerOrderDefault            = EulerOrderXYZ
)

type Euler struct {
	x     float32
	y     float32
	z     float32
	order EulerOrder

	changeEvent *event.Emitter
}

func (eu *Euler) OnChange() *event.Handler {
	return eu.changeEvent.GetHandler()
}

func NewDefaultEuler() *Euler {
	return NewEuler(0, 0, 0, EulerOrderDefault)
}

func NewEuler(x float32, y float32, z float32, order EulerOrder) *Euler {
	return &Euler{
		x:           x,
		y:           y,
		z:           z,
		order:       order,
		changeEvent: event.NewEvent(),
	}
}

func NewEulerFromArray(arr []float32, offset int) *Euler {
	return NewEuler(
		arr[offset],
		arr[offset+1],
		arr[offset+2],
		EulerOrder(int(arr[offset+3])),
	)
}

func (eu *Euler) GetX() float32 {
	return eu.x
}

func (eu *Euler) SetX(x float32) {
	if eu.x != x {
		eu.x = x
		eu.changeEvent.Emit(eu, nil)
	}
}

func (eu *Euler) GetY() float32 {
	return eu.y
}

func (eu *Euler) SetY(y float32) {
	if eu.y != y {
		eu.y = y
		eu.changeEvent.Emit(eu, nil)
	}
}

func (eu *Euler) GetZ() float32 {
	return eu.z
}

func (eu *Euler) SetZ(z float32) {
	if eu.z != z {
		eu.z = z
		eu.changeEvent.Emit(eu, nil)
	}
}

func (eu *Euler) GetOrder() EulerOrder {
	return eu.order
}

func (eu *Euler) SetOrder(order EulerOrder) {
	if eu.order != order {
		eu.order = order
		eu.changeEvent.Emit(eu, nil)
	}
}

func (eu *Euler) Set(x float32, y float32, z float32, order EulerOrder) {
	eu.x = x
	eu.y = y
	eu.z = z
	eu.order = order
	eu.changeEvent.Emit(eu, nil)
}

func (eu *Euler) SetFromRotationMatrixAndOrder(ma *Matrix4, order EulerOrder, update bool) {
	if order == EulerOrderXYZ {
		eu.y = Asin(Clamp(ma.elements[8 ], - 1, 1))
		if Abs(ma.elements[8 ]) < 0.99999 {
			eu.x = Atan2(- ma.elements[9 ], ma.elements[10])
			eu.z = Atan2(- ma.elements[4 ], ma.elements[0 ])
		} else {
			eu.x = Atan2(ma.elements[6 ], ma.elements[5 ])
			eu.z = 0
		}
	} else if order == EulerOrderYXZ {
		eu.x = Asin(-Clamp(ma.elements[9 ], - 1, 1))
		if Abs(ma.elements[9 ]) < 0.99999 {
			eu.y = Atan2(ma.elements[8 ], ma.elements[10])
			eu.z = Atan2(ma.elements[1 ], ma.elements[5 ])
		} else {
			eu.y = Atan2(- ma.elements[2 ], ma.elements[0 ])
			eu.z = 0
		}
	} else if order == EulerOrderZXY {
		eu.x = Asin(Clamp(ma.elements[6 ], - 1, 1))
		if Abs(ma.elements[6 ]) < 0.99999 {
			eu.y = Atan2(- ma.elements[2 ], ma.elements[10])
			eu.z = Atan2(- ma.elements[4 ], ma.elements[5 ])
		} else {
			eu.y = 0
			eu.z = Atan2(ma.elements[1 ], ma.elements[0 ])
		}
	} else if order == EulerOrderZYX {
		eu.y = Asin(-Clamp(ma.elements[2 ], - 1, 1))
		if Abs(ma.elements[2 ]) < 0.99999 {
			eu.x = Atan2(ma.elements[6 ], ma.elements[10])
			eu.z = Atan2(ma.elements[1 ], ma.elements[0 ])
		} else {
			eu.x = 0
			eu.z = Atan2(- ma.elements[4 ], ma.elements[5 ])
		}
	} else if order == EulerOrderYZX {
		eu.z = Asin(Clamp(ma.elements[1 ], - 1, 1))
		if Abs(ma.elements[1 ]) < 0.99999 {
			eu.x = Atan2(- ma.elements[9 ], ma.elements[5 ])
			eu.y = Atan2(- ma.elements[2 ], ma.elements[0 ])
		} else {
			eu.x = 0
			eu.y = Atan2(ma.elements[8 ], ma.elements[10])
		}
	} else if order == EulerOrderXZY {
		eu.z = Asin(-Clamp(ma.elements[4 ], - 1, 1))
		if Abs(ma.elements[4 ]) < 0.99999 {
			eu.x = Atan2(ma.elements[6 ], ma.elements[5 ])
			eu.y = Atan2(ma.elements[8 ], ma.elements[0 ])
		} else {
			eu.x = Atan2(- ma.elements[9 ], ma.elements[10])
			eu.y = 0
		}
	}

	eu.order = order
	if update != false {
		eu.changeEvent.Emit(eu, nil)
	}
}

func (eu *Euler) SetFromQuaternion(q *Quaternion, order EulerOrder, update bool) {
	tmpMatrix := NewDefaultMatrix4()
	tmpMatrix.MakeRotationFromQuaternion(q)
	eu.SetFromRotationMatrixAndOrder(tmpMatrix, order, update)
}

func (eu *Euler) SetFromVector3(v *Vector3) {
	eu.SetFromVector3AndOrder(v, eu.order)
}

func (eu *Euler) SetFromVector3AndOrder(v *Vector3, order EulerOrder) {
	eu.Set(v.X, v.Y, v.Z, order)
}

func (eu *Euler) Clone() *Euler {
	return &Euler{
		x:           eu.x,
		y:           eu.y,
		z:           eu.z,
		order:       eu.order,
		changeEvent: event.NewEvent(),
	}
}

func (eu *Euler) Copy(source *Euler) {
	eu.x = source.x
	eu.y = source.y
	eu.z = source.z
	eu.order = source.order
	eu.changeEvent.Emit(eu, nil)
}

func (eu *Euler) Reorder(newOrder EulerOrder) {
	tmpQua := NewDefaultQuaternion()
	tmpQua.SetFromEuler(eu, false)
	eu.SetFromQuaternion(tmpQua, newOrder, false)
}

func (eu *Euler) Equals(other *Euler) bool {
	return eu.x == other.x && eu.y == other.y && eu.z == other.z && eu.order == other.order
}

func (eu *Euler) ToArray() [4]float32 {
	return [4]float32{eu.x, eu.y, eu.z, float32(eu.order)}
}

func (eu *Euler) ToVector3() *Vector3 {
	return NewVector3(eu.x, eu.y, eu.z)
}
