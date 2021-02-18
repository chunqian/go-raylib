package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"unsafe"
)

const (
	MAX_LIGHTS     = 4
	LIGHT_DISTANCE = 3.5
	LIGHT_HEIGHT   = 1.0
)

type LightType int32

const (
	LIGHT_DIRECTIONAL LightType = iota
	LIGHT_POINT
)

type Light struct {
	enabled    bool
	ltype      LightType
	position   rl.Vector3
	target     rl.Vector3
	color      rl.Color
	enabledLoc int32
	typeLoc    int32
	posLoc     int32
	targetLoc  int32
	colorLoc   int32
}

var (
	lights      = [MAX_LIGHTS]Light{}
	lightsCount = 0
)

func CreateLight(ltype LightType, pos rl.Vector3, targ rl.Vector3, color rl.Color, shader rl.Shader) {

	light := Light{}

	if lightsCount < MAX_LIGHTS {

		light.enabled = true
		light.ltype = ltype
		light.position = pos
		light.target = targ
		light.color = color

		enabledName := fmt.Sprintf("lights[%d].enabled", lightsCount)
		typeName := fmt.Sprintf("lights[%d].type", lightsCount)
		posName := fmt.Sprintf("lights[%d].position", lightsCount)
		targetName := fmt.Sprintf("lights[%d].target", lightsCount)
		colorName := fmt.Sprintf("lights[%d].color", lightsCount)

		light.enabledLoc = rl.GetShaderLocation(shader, enabledName)
		light.typeLoc = rl.GetShaderLocation(shader, typeName)
		light.posLoc = rl.GetShaderLocation(shader, posName)
		light.targetLoc = rl.GetShaderLocation(shader, targetName)
		light.colorLoc = rl.GetShaderLocation(shader, colorName)

		UpdateLightValues(shader, light)

		lights[lightsCount] = light
		lightsCount++
	}
}

func UpdateLightValues(shader rl.Shader, light Light) {

	rl.SetShaderValue(shader, light.enabledLoc, unsafe.Pointer(&light.enabled), int32(rl.UNIFORM_INT))
	rl.SetShaderValue(shader, light.typeLoc, unsafe.Pointer(&light.ltype), int32(rl.UNIFORM_INT))

	position := [3]float32{light.position.X, light.position.Y, light.position.Z}
	rl.SetShaderValue(shader, light.posLoc, unsafe.Pointer(&position), int32(rl.UNIFORM_VEC3))

	target := [3]float32{light.target.X, light.target.Y, light.target.Z}
	rl.SetShaderValue(shader, light.targetLoc, unsafe.Pointer(&target), int32(rl.UNIFORM_VEC3))

	diff := [4]float32{
		float32(light.color.R) / float32(255),
		float32(light.color.G) / float32(255),
		float32(light.color.B) / float32(255),
		float32(light.color.A) / float32(255),
	}
	rl.SetShaderValue(shader, light.colorLoc, unsafe.Pointer(&diff), int32(rl.UNIFORM_VEC4))
}
