package model

type Terminal struct {
	*TerminalManual `gorm:"embedded"`
	*TerminalBasic  `gorm:"embedded"`
	*DefaultField
}

type TerminalManual struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type DefaultField struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt int  `json:"created_at"` // gorm自动使用当前时间戳的秒数填充
	UpdatedAt int  `json:"updated_at"`
}
