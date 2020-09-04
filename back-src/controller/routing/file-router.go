package routing

import (
	"back-src/view/users"
	"github.com/gin-gonic/gin"
)

func (router *router) addFileEndpoints() {
	router.addNewEndpointGroup("/files", "files", "")
	router.addNewEndpointGroup("/profile-pic", "profile-pic", "files")
	router.addNewEndpointGroup("/user", "user-profile-pic", "profile-pic")

	router.addHandlerToPath("/upload", "user-profile-pic", Post, func(context *gin.Context) {
		err := router.handler.UploadUserImage(context)

	})
}
