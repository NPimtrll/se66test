package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `valid:"required~กรุณาระบุชื่อจริง"`
	LastName  string `valid:"required~กรุณาระบุนามสกุล"`
	StudentID string `valid:"required~กรุณาระบุรหัสนักศึกษา, matches(^[BDM]\\d{7}$)~กรุณาใส่รูปแบบให้ถูกต้อง"`
	Phone     string `valid:"stringlength(10|10)~กรุณาใส่ให้ครบ10หลัก"`
	Age       int    `valid:"range(18|50)~อายุของคุณไม่อยู่ในช่วงที่กำหนด"`
	Email     string `valid:"email~กรุณาระบุอีเมลให้ถูกต้อง"`
	Birthday  time.Time 

	
}
