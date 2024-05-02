// Code generated by goa v3.16.1, DO NOT EDIT.
//
// tools service
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package tools

import (
	"context"

	toolsviews "github.com/arduino/arduino-create-agent/gen/tools/views"
	goa "goa.design/goa/v3/pkg"
)

// The tools service manages the available and installed tools
type Service interface {
	// Available implements available.
	Available(context.Context) (res ToolCollection, err error)
	// Installedhead implements installedhead.
	Installedhead(context.Context) (err error)
	// Installed implements installed.
	Installed(context.Context) (res ToolCollection, err error)
	// Install implements install.
	Install(context.Context, *ToolPayload) (res *Operation, err error)
	// Remove implements remove.
	Remove(context.Context, *ToolPayload) (res *Operation, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "arduino-create-agent"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "tools"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"available", "installedhead", "installed", "install", "remove"}

// Operation is the result type of the tools service install method.
type Operation struct {
	// The status of the operation
	Status string
}

// A tool is an executable program that can upload sketches.
type Tool struct {
	// The name of the tool
	Name string
	// The version of the tool
	Version string
	// The packager of the tool
	Packager string
}

// ToolCollection is the result type of the tools service available method.
type ToolCollection []*Tool

// ToolPayload is the payload type of the tools service install method.
type ToolPayload struct {
	// The name of the tool
	Name string
	// The version of the tool
	Version string
	// The packager of the tool
	Packager string
	// The url where the package can be found. Optional.
	// If present checksum must also be present.
	URL *string
	// A checksum of the archive. Mandatory when url is present.
	// This ensures that the package is downloaded correcly.
	Checksum *string
	// The signature used to sign the url. Mandatory when url is present.
	// This ensure the security of the file downloaded
	Signature *string
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_found", false, false, false)
}

// NewToolCollection initializes result type ToolCollection from viewed result
// type ToolCollection.
func NewToolCollection(vres toolsviews.ToolCollection) ToolCollection {
	return newToolCollection(vres.Projected)
}

// NewViewedToolCollection initializes viewed result type ToolCollection from
// result type ToolCollection using the given view.
func NewViewedToolCollection(res ToolCollection, view string) toolsviews.ToolCollection {
	p := newToolCollectionView(res)
	return toolsviews.ToolCollection{Projected: p, View: "default"}
}

// NewOperation initializes result type Operation from viewed result type
// Operation.
func NewOperation(vres *toolsviews.Operation) *Operation {
	return newOperation(vres.Projected)
}

// NewViewedOperation initializes viewed result type Operation from result type
// Operation using the given view.
func NewViewedOperation(res *Operation, view string) *toolsviews.Operation {
	p := newOperationView(res)
	return &toolsviews.Operation{Projected: p, View: "default"}
}

// newToolCollection converts projected type ToolCollection to service type
// ToolCollection.
func newToolCollection(vres toolsviews.ToolCollectionView) ToolCollection {
	res := make(ToolCollection, len(vres))
	for i, n := range vres {
		res[i] = newTool(n)
	}
	return res
}

// newToolCollectionView projects result type ToolCollection to projected type
// ToolCollectionView using the "default" view.
func newToolCollectionView(res ToolCollection) toolsviews.ToolCollectionView {
	vres := make(toolsviews.ToolCollectionView, len(res))
	for i, n := range res {
		vres[i] = newToolView(n)
	}
	return vres
}

// newTool converts projected type Tool to service type Tool.
func newTool(vres *toolsviews.ToolView) *Tool {
	res := &Tool{}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.Packager != nil {
		res.Packager = *vres.Packager
	}
	return res
}

// newToolView projects result type Tool to projected type ToolView using the
// "default" view.
func newToolView(res *Tool) *toolsviews.ToolView {
	vres := &toolsviews.ToolView{
		Name:     &res.Name,
		Version:  &res.Version,
		Packager: &res.Packager,
	}
	return vres
}

// newOperation converts projected type Operation to service type Operation.
func newOperation(vres *toolsviews.OperationView) *Operation {
	res := &Operation{}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	return res
}

// newOperationView projects result type Operation to projected type
// OperationView using the "default" view.
func newOperationView(res *Operation) *toolsviews.OperationView {
	vres := &toolsviews.OperationView{
		Status: &res.Status,
	}
	return vres
}
