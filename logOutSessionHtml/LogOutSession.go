package forum

import (
	t "forum/messages"
)

func LogOutSession() {
	//log out of current session
	t.MESSAGES.SessionUser = ""
}
