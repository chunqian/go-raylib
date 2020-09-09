// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 09 Sep 2020 15:11:08 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package physac

/*
#include "../lib/raylib/src/physac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// Vector2 as declared in src/physac.h:130
type gVector2 struct {
	gX             float32
	gY             float32
	ref29ca61a5    *C.Vector2
	allocs29ca61a5 interface{}
}
type Vector2 struct {
	X float32
	Y float32
}

// Matrix2x2 as declared in src/physac.h:142
type gMatrix2x2 struct {
	gM00           float32
	gM01           float32
	gM10           float32
	gM11           float32
	refb92f06e6    *C.Matrix2x2
	allocsb92f06e6 interface{}
}
type Matrix2x2 struct {
	M00 float32
	M01 float32
	M10 float32
	M11 float32
}

// PolygonData as declared in src/physac.h:148
type gPolygonData struct {
	gVertexCount   uint32
	gPositions     [24]gVector2
	gNormals       [24]gVector2
	ref87a94eb7    *C.PolygonData
	allocs87a94eb7 interface{}
}
type PolygonData struct {
	VertexCount uint32
	Positions   [24]Vector2
	Normals     [24]Vector2
}

// PhysicsShape as declared in src/physac.h:157
type gPhysicsShape struct {
	gType          PhysicsShapeType
	gBody          []gPhysicsBodyData
	gRadius        float32
	gTransform     gMatrix2x2
	gVertexData    gPolygonData
	ref4c540f1a    *C.PhysicsShape
	allocs4c540f1a interface{}
}
type PhysicsShape struct {
	Type       PhysicsShapeType
	Body       *PhysicsBodyData
	Radius     float32
	Transform  Matrix2x2
	VertexData PolygonData
}

// PhysicsBodyData as declared in src/physac.h:179
type gPhysicsBodyData struct {
	gId              uint32
	gEnabled         bool
	gPosition        gVector2
	gVelocity        gVector2
	gForce           gVector2
	gAngularVelocity float32
	gTorque          float32
	gOrient          float32
	gInertia         float32
	gInverseInertia  float32
	gMass            float32
	gInverseMass     float32
	gStaticFriction  float32
	gDynamicFriction float32
	gRestitution     float32
	gUseGravity      bool
	gIsGrounded      bool
	gFreezeOrient    bool
	gShape           gPhysicsShape
	refd780e53       *C.PhysicsBodyData
	allocsd780e53    interface{}
}
type PhysicsBodyData struct {
	Id              uint32
	Enabled         bool
	Position        Vector2
	Velocity        Vector2
	Force           Vector2
	AngularVelocity float32
	Torque          float32
	Orient          float32
	Inertia         float32
	InverseInertia  float32
	Mass            float32
	InverseMass     float32
	StaticFriction  float32
	DynamicFriction float32
	Restitution     float32
	UseGravity      bool
	IsGrounded      bool
	FreezeOrient    bool
	Shape           PhysicsShape
}

// PhysicsManifoldData as declared in src/physac.h:194
type gPhysicsManifoldData struct {
	gId              uint32
	gBodyA           []gPhysicsBodyData
	gBodyB           []gPhysicsBodyData
	gPenetration     float32
	gNormal          gVector2
	gContacts        [2]gVector2
	gContactsCount   uint32
	gRestitution     float32
	gDynamicFriction float32
	gStaticFriction  float32
	ref10b92967      *C.PhysicsManifoldData
	allocs10b92967   interface{}
}
type PhysicsManifoldData struct {
	Id              uint32
	BodyA           *PhysicsBodyData
	BodyB           *PhysicsBodyData
	Penetration     float32
	Normal          Vector2
	Contacts        [2]Vector2
	ContactsCount   uint32
	Restitution     float32
	DynamicFriction float32
	StaticFriction  float32
}
