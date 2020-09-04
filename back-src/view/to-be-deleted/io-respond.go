package to_be_deleted

import (
	"back-src/view/notifications"
)

func RespondIO(notification notifications.Notification) {
	notification.Context.Header("Token", notification.Token)
	notification.Context.JSON(notification.StatusCode, notification.Response{Message: notification.Message})
}
