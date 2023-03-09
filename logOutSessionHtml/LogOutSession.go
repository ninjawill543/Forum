package forum

import (
	t "forum/messages"
)

func LogOutSession() {
	t.Messages.SessionUser = ""
}
