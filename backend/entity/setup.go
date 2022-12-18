package entity

import (
	"github.com/Polalius/sa-65-example/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("SportsCenter.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		//Employee system
		&Login{},
		&Role{},
		&Employee{},
		//WormUp
		&WormUp{},
		&Exercise{},
		&Stretch{},
		&ExerciseProgramList{},
	)

	db = database

	//เราจะทำการสร้าง Admin account ไว้สำหรับการสร้าง Employee และ role ต่างๆ

	password, _ := services.Hash("123456")

	admin := Role{
		Name: "admin",
	}
	db.Model(&Role{}).Create(&admin)

	trainer := Role{
		Name: "trainer",
	}

	db.Model(&Role{}).Create(&trainer)

	//Admin 1
	admin1 := Login{
		User:     "Admin1",
		Password: string(password),
	}
	admin1err := db.Model(&Login{}).Create(&admin1)
	admin1emp := Employee{
		Name:    "Admin1",
		Surname: "Example1",
		Login:   admin1,
		Role:    admin,
	}
	if admin1err.Error == nil {
		db.Model(&Employee{}).Create(&admin1emp)
	}

	//Admin 2
	admin2 := Login{
		User:     "Admin2",
		Password: string(password),
	}
	admin2err := db.Model(&Login{}).Create(&admin2)
	admin2emp := Employee{
		Name:    "Admin1",
		Surname: "Example1",
		Login:   admin2,
		Role:    admin,
	}
	if admin2err.Error == nil {
		db.Model(&Employee{}).Create(&admin2emp)
	}

	//Trainer 1
	Trainer1 := Login{
		User:     "Trainer1",
		Password: string(password),
	}
	train1err := db.Model(&Login{}).Create(&Trainer1)

	train1emp := Employee{
		Name:    "Trainer1",
		Surname: "Example1",
		Login:   Trainer1,
		Role:    trainer,
	}
	if train1err.Error == nil {
		db.Model(&Employee{}).Create(&train1emp)
	}

	//------------------------------ trainer create program --------------------------
	// worm up
	worm1 := WormUp{
		SetName:          "Set_AA",
		ExercisePoseture: "Scapula Retraction, Chin in",
	}
	db.Model(&WormUp{}).Create(&worm1)
	worm2 := WormUp{
		SetName:          "Set_AB",
		ExercisePoseture: "Wall Slide, Cat Camel",
	}
	db.Model(&WormUp{}).Create(&worm2)
	worm3 := WormUp{
		SetName:          "Set_AC",
		ExercisePoseture: "Chin In, Wall Slide",
	}
	db.Model(&WormUp{}).Create(&worm3)
	// exercise
	exer1 := Exercise{
		SetName:          "Set_BA",
		ExercisePoseture: "Cable Row, Dumpbell Press, Lat Pulldown",
	}
	db.Model(&Exercise{}).Create(&exer1)
	exer2 := Exercise{
		SetName:          "Set_BB",
		ExercisePoseture: "Cable Face Pull, DB Reverse Fly, Plank",
	}
	db.Model(&Exercise{}).Create(&exer2)
	exer3 := Exercise{
		SetName:          "Set_BC",
		ExercisePoseture: "Cable Row, DB Reverse Fly, Plank",
	}
	db.Model(&Exercise{}).Create(&exer3)
	// stratch
	stre1 := Stretch{
		SetName:          "Set_CA",
		ExercisePoseture: "Seated Erector Stretch, Neck Stretch",
	}
	db.Model(&Stretch{}).Create(&stre1)
	stre2 := Stretch{
		SetName:          "Set_CB",
		ExercisePoseture: "Front Delt Stretch, Lying Twist Chest Stretch",
	}
	db.Model(&Stretch{}).Create(&stre2)
	stre3 := Stretch{
		SetName:          "Set_CC",
		ExercisePoseture: "Levetor Stretch, Chest Wall Stretch",
	}
	db.Model(&Stretch{}).Create(&stre3)
	// Order
	exeprogram1 := ExerciseProgramList{
		ProgramName: "Program_A",
		Employee:    train1emp,
		WormUp:      worm1,
		Exercise:    exer1,
		Stretch:     stre1,
		Minute:      60,
	}
	db.Model(&ExerciseProgramList{}).Create(&exeprogram1)
	exeprogram2 := ExerciseProgramList{
		ProgramName: "Program_B",
		Employee:    train1emp,
		WormUp:      worm2,
		Exercise:    exer1,
		Stretch:     stre3,
		Minute:      90,
	}
	db.Model(&ExerciseProgramList{}).Create(&exeprogram2)
}
