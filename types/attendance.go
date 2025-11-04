package types

type AttendanceCreate struct {
	Side  string `json:"side"`
	Name  string `json:"name"`
	Meal  string `json:"meal"`
	Count int    `json:"count"`
}

type AttendanceItem struct {
    Id        int   `json:"id"`
    Side      string `json:"side"`
    Name      string `json:"name"`
    Meal      string `json:"meal"`
    Count     int    `json:"count"`
    Timestamp int64  `json:"timestamp"`
}

type AttendanceListResponse struct {
    Items []AttendanceItem `json:"items"`
    Total int              `json:"total"`
}
