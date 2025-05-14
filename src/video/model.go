package video

import "time"

type VideoCall struct {
	CallerID  string    `json:"callerId"`
	CalleeID  string    `json:"calleeId"`
	RoomID    string    `json:"roomId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

type JoinCallRequest struct {
	RoomID string `json:"roomId" binding:"required"`
	UserID string `json:"userId" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

type JoinCallResponse struct {
	Message string `json:"message"`
	RoomID  string `json:"roomId"`
	UserID  string `json:"userId"`
}
