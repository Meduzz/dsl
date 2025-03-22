package qapi

import (
	"fmt"

	"github.com/Meduzz/dsl/api"
)

const cType = "application/json"

func Quickapi(api *api.Api, prefix, name string, entity any) {
	create := api.POST(createUrl("/%s/", prefix, name))
	create.Description = fmt.Sprintf("Create a new %s entity", name)
	createReq := create.BodyVariable("body", cType)
	createReq.SetType(entity)
	createResp := create.SetResponse(cType)
	createResp.SetType(entity)
	create.SetPermission(fmt.Sprintf("%s.create", name))

	read := api.GET(createUrl("/%s/:id", prefix, name))
	read.Description = fmt.Sprintf("Read an %s entity by id", name)
	read.PathVariable("id")
	readResp := read.SetResponse(cType)
	readResp.SetType(entity)
	read.SetPermission(fmt.Sprintf("%s.read", name))

	update := api.PUT(createUrl("/%s/:id", prefix, name))
	update.Description = fmt.Sprintf("Update an %s entity by id", name)
	updateReq := update.BodyVariable("body", cType)
	updateReq.SetType(entity)
	updateResp := update.SetResponse(cType)
	updateResp.SetType(entity)
	update.SetPermission(fmt.Sprintf("%s.update", name))

	remove := api.DELETE(createUrl("/%s/:id", prefix, name))
	remove.Description = fmt.Sprintf("Delete an %s entity by id", name)
	remove.PathVariable("id")
	remove.SetPermission(fmt.Sprintf("%s.delete", name))

	search := api.GET(createUrl("/%s/", prefix, name))
	search.Description = fmt.Sprintf("Read an %s entity by id", name)
	sSkip := search.QueryVariable("skip")
	sSkip.SetType(0)
	sTake := search.QueryVariable("take")
	sTake.SetType(25)
	sWhere := search.QueryVariable("where")
	sWhere.MapOf("string")
	sSort := search.QueryVariable("sort")
	sSort.MapOf("string")
	sPreload := search.QueryVariable("preload")
	sPreload.MapOf("string")
	searchResp := read.SetResponse(cType)
	searchResp.ArrayOf(entity)
	search.SetPermission(fmt.Sprintf("%s.read", name))

	patch := api.PATCH(createUrl("/%s/:id", prefix, name))
	patch.Description = fmt.Sprintf("Patch individual fields of an %s entity", name)
	patch.PathVariable("id")
	patchReq := patch.BodyVariable("body", cType)
	patchReq.Map = true
	patchReq.Type = "any"
	patchResp := patch.SetResponse(cType)
	patchResp.SetType(entity)
	patch.SetPermission(fmt.Sprintf("%s.update", name))
}

func createUrl(url, prefix, name string) string {
	if prefix != "" {
		return fmt.Sprintf("%s%s", prefix, fmt.Sprintf(url, name))
	}

	return fmt.Sprintf(url, name)
}
