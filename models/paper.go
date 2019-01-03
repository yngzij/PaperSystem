package models

import "time"

type Paper struct {
	Id        int
	Uuid      string
	Account     string
	Title    string
	UserId    int
	Image_1      string
	Image_2      string
	Image_3      string
	Image_4      string
	Body   string
	CreatedAt time.Time
}

type PaperInfo struct {
	
}


func GetAllPaper() (papers []Paper, err error) {
	err= Db.Find(&papers)
	return
}

func (paper *Paper) Create() (err error) {
	_, err = Db.Insert(paper)
	return nil
}


