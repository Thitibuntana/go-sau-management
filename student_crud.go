package main

import (
	"gorm.io/gorm"
)

func createStudent(db *gorm.DB, student *Student)error{
	return db.Create(student).Error
}

func getAllStudent(db *gorm.DB) ([]Student, error){
	var student []Student
	result := db.Find(&student)
	return student, result.Error
}

func getStudentByID(db *gorm.DB, id int) (*Student, error){
	var student Student

	result := db.First(&student, id)
	if result.Error != nil{
		return nil, result.Error
	}

	return &student, nil
}

func getStudentByName (db *gorm.DB, StuName string)([]Student, error){
	var students []Student

	result := db.Where("stu_name LIKE ?","%" +StuName+"%").Find(&students)

	return students, result.Error
}

func getStudentByLevel (db *gorm.DB, StuLevel int) ([]Student, error){
	var students []Student

	result := db.Where("stu_level = ?",StuLevel).Find(&students)

	return students, result.Error
}

func updateStudent(db *gorm.DB, id int, student *Student) error{
	return db.Model(&Student{}).Where("id = ?", id).Updates(student).Error
}

func deleteStudent(db *gorm.DB, id int) error{
	return db.Delete(&Student{}, id).Error
}