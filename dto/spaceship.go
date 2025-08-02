package dto

type ArmamentDTO struct {
	Title string `json:"title" binding:"required"`
	Qty   int    `json:"qty" binding:"required"`
}

type GetSpaceshipFilterDTO struct {
	Name         string   `form:"name"`
	Class        string   `form:"class"`
	Crew         *int     `form:"crew"`
	Status       string   `form:"status"`
	Value        *float64 `form:"value"`
	ValueMin     *float64 `form:"value_min"`
	ValueMax     *float64 `form:"value_max"`
	CreatedAt    string   `form:"created_at"`
	CreatedAtMin string   `form:"created_at_min"`
	CreatedAtMax string   `form:"created_at_max"`
	UpdatedAt    string   `form:"updated_at"`
	UpdatedAtMin string   `form:"updated_at_min"`
	UpdatedAtMax string   `form:"updated_at_max"`
	Limit        int      `form:"limit"`
	Page         int      `form:"page"`
}

type PostSpaceshipDTO struct {
	Name     string        `json:"name" binding:"required"`
	Class    string        `json:"class" binding:"required"`
	Crew     int           `json:"crew" binding:"required"`
	Image    string        `json:"image"`
	Value    float64       `json:"value" binding:"required"`
	Status   string        `json:"status" binding:"required"`
	Armament []ArmamentDTO `json:"armament" binding:"required"`
}

type PutSpaceshipDTO struct {
	Name     string        `json:"name" binding:"required"`
	Class    string        `json:"class" binding:"required"`
	Crew     int           `json:"crew" binding:"required"`
	Image    string        `json:"image"`
	Value    float64       `json:"value" binding:"required"`
	Status   string        `json:"status" binding:"required"`
	Armament []ArmamentDTO `json:"armament" binding:"required"`
}
