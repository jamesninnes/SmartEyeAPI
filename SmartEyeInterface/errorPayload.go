package SmartEyeInterface

import "time"

type errorPayload struct {
	id        int       `json:"id"`
	message   string    `json:"name"`
	created   time.Time `json:"created"`
}

func  throwError(id int, message string) errorPayload {
	var currErrorPayload errorPayload

	currErrorPayload.id = id
	currErrorPayload.message = message
	currErrorPayload.created = time.Now()

	return currErrorPayload
}

func  GetMessage(e errorPayload) string {
	return e.message
}

func  GetTime(e errorPayload) time.Time{
	return e.created

}
