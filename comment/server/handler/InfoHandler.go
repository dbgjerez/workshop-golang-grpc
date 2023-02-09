package handler

import (
	"context"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"github.com/dbgjerez/workshop-golang-grpc/comment/server/utils"
)

const (
	serviceVersion       = "SERVICE_VERSION"
	serviceName          = "SERVICE_NAME"
	buildTime            = "SERVICE_BUILD_TIME"
	SERVICE_NAME_DEFAULT = "comment-grpc"
	NO_VALUE             = "no default value"
)

type InfoHandler struct {
	c.UnsafeInfoServiceServer
	info *c.Info
}

func NewInfoHandler() *InfoHandler {
	info := c.Info{
		Name:      utils.GetEnv(serviceName, SERVICE_NAME_DEFAULT),
		Version:   utils.GetEnv(serviceVersion, NO_VALUE),
		BuildTime: utils.GetEnv(buildTime, NO_VALUE),
	}
	return &InfoHandler{info: &info}
}

func (ih *InfoHandler) GetInfo(context.Context, *c.RetrieveInfo) (*c.Info, error) {
	// This application has not external dependencies as ddbb so we needn't check anymore
	return ih.info, nil
}
