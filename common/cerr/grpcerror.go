package cerr

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var grpcToHttpStatus = map[codes.Code]int{
	codes.InvalidArgument:  http.StatusBadRequest,
	codes.Unauthenticated:  http.StatusUnauthorized,
	codes.PermissionDenied: http.StatusForbidden,
	codes.NotFound:         http.StatusNotFound,
	codes.AlreadyExists:    http.StatusConflict,
	codes.Unavailable:      http.StatusServiceUnavailable,
}

func HandleGrpcError(ctx *gin.Context, err error) {
	grpcErr, ok := status.FromError(err)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	httpStatus, ok := grpcToHttpStatus[grpcErr.Code()]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	ctx.AbortWithStatusJSON(httpStatus, gin.H{
		"error": grpcErr.Message(),
	})
}
