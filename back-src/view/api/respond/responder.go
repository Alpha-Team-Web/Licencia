package respond

import (
	"back-src/view/notifications"
	"fmt"
)

func Respond(notification notifications.Notification) {
	ctx := notification.Context
	fmt.Println(notification.Data)
	ctx.Header("Token", notification.Token)
	ctx.JSON(notification.StatusCode, notifications.Response{Message: notification.Message, Data: notification.Data})
}
