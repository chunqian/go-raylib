package main

import (
	rl "goray/raylib"

	"runtime"
	"unsafe"
)

var (
	CUBEMAP_SIZE     = 512
	IRRADIANCE_SIZE  = 32
	PREFILTERED_SIZE = 256
	BRDF_SIZE        = 512
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))
	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - pbr material")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(4.0, 4.0, 4.0),
		rl.NewVector3(0, 0.5, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	model := rl.LoadModel("../models/resources/pbr/trooper.obj")
	defer rl.UnloadModel(model)

	rl.MeshTangents(model.Meshes)

	rl.UnloadMaterial(*model.Materials)
	*model.Materialser(0) = LoadMaterialPBR(rl.NewColor(255, 255, 255, 255), 1.0, 1.0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		cameraPos := [3]float32{
			camera.Position.X,
			camera.Position.Y,
			camera.Position.Z,
		}

		rl.SetShaderValue(
			model.Materialser(0).Shader,
			*model.Materialser(0).Shader.Locser(int32(rl.LOC_VECTOR_VIEW)),
			unsafe.Pointer(&cameraPos),
			int32(rl.UNIFORM_VEC3),
		)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, rl.Vector3Zero(), 1.0, rl.White)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

func LoadMaterialPBR(albedo rl.Color, metalness float32, roughness float32) rl.Material {

	mat := rl.LoadMaterialDefault()

	mat.Shader = rl.LoadShader("../models/resources/shaders/glsl330/pbr.vs", "../models/resources/shaders/glsl330/pbr.fs")

	*mat.Shader.Locser(int32(rl.LOC_MAP_ALBEDO)) = rl.GetShaderLocation(mat.Shader, "albedo.sampler")
	*mat.Shader.Locser(int32(rl.LOC_MAP_METALNESS)) = rl.GetShaderLocation(mat.Shader, "metalness.sampler")
	*mat.Shader.Locser(int32(rl.LOC_MAP_NORMAL)) = rl.GetShaderLocation(mat.Shader, "normals.sampler")
	*mat.Shader.Locser(int32(rl.LOC_MAP_ROUGHNESS)) = rl.GetShaderLocation(mat.Shader, "roughness.sampler")
	*mat.Shader.Locser(int32(rl.LOC_MAP_OCCLUSION)) = rl.GetShaderLocation(mat.Shader, "occlusion.sampler")

	*mat.Shader.Locser(int32(rl.LOC_MAP_IRRADIANCE)) = rl.GetShaderLocation(mat.Shader, "irradianceMap")
	*mat.Shader.Locser(int32(rl.LOC_MAP_PREFILTER)) = rl.GetShaderLocation(mat.Shader, "prefilterMap")
	*mat.Shader.Locser(int32(rl.LOC_MAP_BRDF)) = rl.GetShaderLocation(mat.Shader, "brdfLUT")
	*mat.Shader.Locser(int32(rl.LOC_MATRIX_MODEL)) = rl.GetShaderLocation(mat.Shader, "matModel")
	*mat.Shader.Locser(int32(rl.LOC_VECTOR_VIEW)) = rl.GetShaderLocation(mat.Shader, "viewPos")

	mat.Mapser(int32(rl.MAP_ALBEDO)).Texture = rl.LoadTexture("../models/resources/pbr/trooper_albedo.png")
	mat.Mapser(int32(rl.MAP_NORMAL)).Texture = rl.LoadTexture("../models/resources/pbr/trooper_normals.png")
	mat.Mapser(int32(rl.MAP_METALNESS)).Texture = rl.LoadTexture("../models/resources/pbr/trooper_metalness.png")
	mat.Mapser(int32(rl.MAP_ROUGHNESS)).Texture = rl.LoadTexture("../models/resources/pbr/trooper_roughness.png")
	mat.Mapser(int32(rl.MAP_OCCLUSION)).Texture = rl.LoadTexture("../models/resources/pbr/trooper_ao.png")

	shdrCubemap := rl.LoadShader("../models/resources/shaders/glsl330/cubemap.vs", "../models/resources/shaders/glsl330/cubemap.fs")

	shdrIrradiance := rl.LoadShader("../models/resources/shaders/glsl330/skybox.vs", "../models/resources/shaders/glsl330/irradiance.fs")

	shdrPrefilter := rl.LoadShader("../models/resources/shaders/glsl330/skybox.vs", "../models/resources/shaders/glsl330/prefilter.fs")

	shdrBRDF := rl.LoadShader("../models/resources/shaders/glsl330/brdf.vs", "../models/resources/shaders/glsl330/brdf.fs")

	i := []int32{1}
	rl.SetShaderValue(shdrCubemap, rl.GetShaderLocation(shdrCubemap, "equirectangularMap"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))
	rl.SetShaderValue(shdrIrradiance, rl.GetShaderLocation(shdrIrradiance, "environmentMap"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))
	rl.SetShaderValue(shdrPrefilter, rl.GetShaderLocation(shdrPrefilter, "environmentMap"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))

	texHDR := rl.LoadTexture("../models/resources/dresden_square.hdr")
	cubemap := rl.GenTextureCubemap(shdrCubemap, texHDR, int32(CUBEMAP_SIZE))

	mat.Mapser(int32(rl.MAP_IRRADIANCE)).Texture = rl.GenTextureIrradiance(shdrIrradiance, cubemap, int32(IRRADIANCE_SIZE))
	mat.Mapser(int32(rl.MAP_PREFILTER)).Texture = rl.GenTexturePrefilter(shdrPrefilter, cubemap, int32(PREFILTERED_SIZE))
	mat.Mapser(int32(rl.MAP_BRDF)).Texture = rl.GenTextureBRDF(shdrBRDF, int32(BRDF_SIZE))

	rl.UnloadTexture(cubemap)
	rl.UnloadTexture(texHDR)

	rl.UnloadShader(shdrCubemap)
	rl.UnloadShader(shdrIrradiance)
	rl.UnloadShader(shdrPrefilter)
	rl.UnloadShader(shdrBRDF)

	rl.SetTextureFilter(mat.Mapser(int32(rl.MAP_ALBEDO)).Texture, int32(rl.FILTER_BILINEAR))
	rl.SetTextureFilter(mat.Mapser(int32(rl.MAP_NORMAL)).Texture, int32(rl.FILTER_BILINEAR))
	rl.SetTextureFilter(mat.Mapser(int32(rl.MAP_METALNESS)).Texture, int32(rl.FILTER_BILINEAR))
	rl.SetTextureFilter(mat.Mapser(int32(rl.MAP_ROUGHNESS)).Texture, int32(rl.FILTER_BILINEAR))
	rl.SetTextureFilter(mat.Mapser(int32(rl.MAP_OCCLUSION)).Texture, int32(rl.FILTER_BILINEAR))

	rl.SetShaderValue(mat.Shader, rl.GetShaderLocation(mat.Shader, "albedo.useSampler"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))
	rl.SetShaderValue(mat.Shader, rl.GetShaderLocation(mat.Shader, "normals.useSampler"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))
	rl.SetShaderValue(mat.Shader, rl.GetShaderLocation(mat.Shader, "metalness.useSampler"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))
	rl.SetShaderValue(mat.Shader, rl.GetShaderLocation(mat.Shader, "roughness.useSampler"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))
	rl.SetShaderValue(mat.Shader, rl.GetShaderLocation(mat.Shader, "occlusion.useSampler"), unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))

	renderModeLoc := rl.GetShaderLocation(mat.Shader, "renderMode")
	rl.SetShaderValue(mat.Shader, renderModeLoc, unsafe.Pointer(&i[0]), int32(rl.UNIFORM_INT))

	mat.Mapser(int32(rl.MAP_ALBEDO)).Color = albedo
	mat.Mapser(int32(rl.MAP_NORMAL)).Color = rl.NewColor(128, 128, 255, 255)
	mat.Mapser(int32(rl.MAP_METALNESS)).Value = metalness
	mat.Mapser(int32(rl.MAP_ROUGHNESS)).Value = roughness
	mat.Mapser(int32(rl.MAP_OCCLUSION)).Value = 1.0
	mat.Mapser(int32(rl.MAP_EMISSION)).Value = 0.5
	mat.Mapser(int32(rl.MAP_HEIGHT)).Value = 0.5

	return mat
}
