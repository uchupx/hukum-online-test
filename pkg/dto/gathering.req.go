package dto

type GatheringPutV1 struct {
	Creator     *int64  `json:"creator"`
	Location    *string `json:"location"`
	ScheduledAt *string `json:"scheduled_at"`
	Name        *string `json:"name"`
	Type        *string `json:"type"`
}

type GatheringPostV1 struct {
	Creator     int64  `json:"creator" binding:"required"`
	Location    string `json:"location" binding:"required"`
	ScheduledAt string `json:"scheduled_at" binding:"required" time_format:"2006-01-02 15:04:05"`
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
}
