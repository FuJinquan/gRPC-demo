package main

import (
	"context"
	"server/ErrorCode"
	"server/mysql"
	"server/note"
)

type Server struct {
	note.UnimplementedNoteOptionServer
}

//NoteOptionServer
//添加记录
func (s *Server) AddRecord(ctx context.Context, in *note.AddRequest) (*note.Response, error) {
	db := mysql.GetDB()
	note1 := mysql.Note{
		Title: in.Title,
		Text:  in.Text,
	}
	if err := db.Create(&note1).Error; err != nil {
		return &note.Response{
			Code:    ErrorCode.ErrCreateRecord.ErrorCode,
			Data:    &note.AddRequest{Title: "", Text: ""},
			Message: ErrorCode.ErrCreateRecord.ErrorMsg,
		}, err
	}
	return &note.Response{
		Code:    2,
		Data:    &note.AddRequest{Title: in.Title, Text: in.Text},
		Message: "添加新记录成功！",
	}, nil
}

//更新记录
func (s *Server) UpdateRecord(ctx context.Context, in *note.UpdateRequest) (*note.Response, error) {
	db := mysql.GetDB()
	var note1 mysql.Note
	note1.ID = uint(in.Id)
	if err := db.Model(&note1).Updates(mysql.Note{Title: in.Title, Text: in.Text}).Error; err != nil {
		return &note.Response{
			Code:    ErrorCode.ErrUpdate.ErrorCode,
			Data:    &note.AddRequest{Title: "", Text: ""},
			Message: ErrorCode.ErrUpdate.ErrorMsg,
		}, err
	}
	return &note.Response{
		Code:    2,
		Data:    &note.AddRequest{Title: in.Title, Text: in.Text},
		Message: "更新记录成功！",
	}, nil

}

//查找记录
func (s *Server) QueryRecord(ctx context.Context, in *note.Request) (*note.Response, error) {
	var note1 mysql.Note
	db := mysql.GetDB()
	if err := db.Find(&note1, in.Id).Error; err != nil {
		return &note.Response{
			Code:    ErrorCode.ErrSearch.ErrorCode,
			Data:    &note.AddRequest{Title: "", Text: ""},
			Message: ErrorCode.ErrSearch.ErrorMsg,
		}, err
	}
	return &note.Response{
		Code:    2,
		Data:    &note.AddRequest{Title: note1.Title, Text: note1.Text},
		Message: "查找记录成功！",
	}, nil

}

//删除记录
func (s *Server) DeleteRecord(ctx context.Context, in *note.Request) (*note.Response, error) {
	db := mysql.GetDB()
	var note1 mysql.Note
	note1.ID = uint(in.Id)
	if err := db.Unscoped().Delete(&note1).Error; err != nil {
		return &note.Response{
			Code:    ErrorCode.ErrDelete.ErrorCode,
			Data:    &note.AddRequest{Title: "", Text: ""},
			Message: ErrorCode.ErrDelete.ErrorMsg,
		}, err
	}
	return &note.Response{
		Code:    2,
		Data:    &note.AddRequest{Title: note1.Title, Text: note1.Text},
		Message: "记录删除成功！",
	}, nil
}
