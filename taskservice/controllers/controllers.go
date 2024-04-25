package controllers

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"taskservice/db"
	"taskservice/protos"
)

type TaskServiceServer struct {
	curTaskId int64
	protos.UnimplementedTaskServiceServer
}

func (s *TaskServiceServer) getTaskByIdAndOwner(id int64, owner string) (*db.Task, error) {
	var existingTask *db.Task
	if result := db.DB.First(&existingTask, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		} else {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
	}

	if owner != existingTask.Owner {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("User %v is not owner of task %v", owner, id))
	}

	return existingTask, nil
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *protos.Task) (*protos.TaskCreds, error) {
	s.curTaskId++

	task := &db.Task{Task: protos.Task{
		Id:     s.curTaskId,
		Owner:  req.Owner,
		Body:   req.Body,
		Status: req.Status,
	}}

	if result := db.DB.Create(task); result.Error != nil {
		return nil, result.Error
	}

	return &protos.TaskCreds{Id: s.curTaskId}, nil
}

func (s *TaskServiceServer) UpdateTask(ctx context.Context, req *protos.Task) (*empty.Empty, error) {
	if existingTask, err := s.getTaskByIdAndOwner(req.Id, req.Owner); err != nil {
		return nil, err
	} else {
		if req.Body != "" {
			existingTask.Body = req.Body
		}
		if req.Status != protos.Status_DEFAULT {
			existingTask.Status = req.Status
		}

		if result := db.DB.Save(existingTask); result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		return nil, nil
	}
}

func (s *TaskServiceServer) DeleteTask(ctx context.Context, req *protos.TaskCreds) (*empty.Empty, error) {
	if existingTask, err := s.getTaskByIdAndOwner(req.Id, req.Owner); err != nil {
		return nil, err
	} else {
		if result := db.DB.Delete(existingTask); result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}

		return nil, nil
	}
}

func (s *TaskServiceServer) GetTask(ctx context.Context, req *protos.TaskCreds) (*protos.Task, error) {
	if existingTask, err := s.getTaskByIdAndOwner(req.Id, req.Owner); err != nil {
		return nil, err
	} else {
		return &existingTask.Task, nil
	}
}

func (s *TaskServiceServer) GetTasksByPage(ctx context.Context, req *protos.PageRequest) (*protos.PageResponse, error) {
	offset := (req.Number - 1) * req.Size

	var tasks []*db.Task
	if result := db.DB.Where("owner = ?", req.Owner).Order("id DESC").Limit(int(req.Size)).Offset(int(offset)).Find(&tasks); result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var protoTasks []*protos.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, &task.Task)
	}

	return &protos.PageResponse{Tasks: protoTasks}, nil
}
