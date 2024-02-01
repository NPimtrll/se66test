package unit

import (
	"testing"
	"time"

	"github.com/NPimtrll/se66test/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestFirstname(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`firstname is required`, func(t *testing.T) {
		user := entity.User{
			FirstName: "",
			LastName:  "Srimartpirom",
			StudentID: "B6401887",
			Phone:     "0811387170",
			Age:       20,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("กรุณาระบุชื่อจริง"))
	})

	t.Run(`firstName is valid`, func(t *testing.T) {
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "Srimartpirom",
			StudentID: "B6401887",
			Phone:     "0811387170",
			Age:       20,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestSurName(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`surname is Required`, func(*testing.T) {
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "",
			StudentID: "B6401887",
			Phone:     "0811387170",
			Age:       20,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(`กรุณาระบุนามสกุล`))

	})

	t.Run(`surname is valid`, func(*testing.T){
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "Srimartpirom",
			StudentID: "B6401887",
			Phone:     "0811387170",
			Age:       20,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

}


func TestStudentID(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`student_id is required`,func(*testing.T){
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "Srimartpirom",
			StudentID: "",
			Phone:     "0811387170",
			Age:       20,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(`กรุณาระบุรหัสนักศึกษา`))
	})

	t.Run(`student_id pattern is not true`, func(*testing.T){
		user := entity.User{
		FirstName: "Nannapat",
		LastName:  "Srimartpirom",
		StudentID: "K1234567",
		Phone:     "0811387170",
		Age:       20,
		Email:     "ploy54218740@gmail.com",
		Birthday:  time.Now(),
		}

		ok,err := govalidator.ValidateStruct(user)
	
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
	
		g.Expect(err.Error()).To(Equal("กรุณาใส่รูปแบบให้ถูกต้อง"))
	})

	t.Run(`StudentID is valid`,func(*testing.T){
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "Srimartpirom",
			StudentID: "B6401887",
			Phone:     "0811387170",
			Age:       20,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})


}

func TestPhone(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`phone_number check 10 digit` , func(*testing.T){
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "Srimartpirom",
			StudentID: "B6401887",
			Phone:     "341",
			Age:       20,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}
	
		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("กรุณาใส่ให้ครบ10หลัก"))
	})
}

func TestAge(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Age not in range` , func(*testing.T){
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "Srimartpirom",
			StudentID: "B6401887",
			Phone:     "0811387170",
			Age:       10,
			Email:     "ploy54218740@gmail.com",
			Birthday:  time.Now(),
		}

	ok,err := govalidator.ValidateStruct(user)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("อายุของคุณไม่อยู่ในช่วงที่กำหนด"))
	})
}

func TestEmail(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`email pattern not true` , func(*testing.T){
		user := entity.User{
			FirstName: "Nannapat",
			LastName:  "Srimartpirom",
			StudentID: "B6401887",
			Phone:     "0811387170",
			Age:       19,
			Email:     "ploy54218740gmailcom",
			Birthday:  time.Now(),
		}

	ok,err := govalidator.ValidateStruct(user)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("กรุณาระบุอีเมลให้ถูกต้อง"))
	})
}