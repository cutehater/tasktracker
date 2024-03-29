package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	db "taskservice/db"
	pb "taskservice/protos"
)

type TaskServiceServer struct {
	curTaskId int64
	pb.UnimplementedTaskServiceServer
}

func (s *TaskServiceServer) getTaskByIdAndOwner(id int64, owner string) (*pb.Task, error) {
	var existingTask *pb.Task
	if result := db.DB.First(existingTask, id); result.Error != nil {
		return nil, result.Error
	}

	if owner != existingTask.Owner {
		return nil, errors.New(fmt.Sprintf("User %v is not owner of task %v", owner, id))
	}

	return existingTask, nil
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *pb.Task) (*pb.TaskCreds, error) {
	s.curTaskId++

	task := &pb.Task{
		Id:     s.curTaskId,
		Owner:  req.Owner,
		Body:   req.Body,
		Status: req.Status,
	}

	if result := db.DB.Create(task); result.Error != nil {
		return nil, result.Error
	}

	return &pb.TaskCreds{Id: s.curTaskId}, nil
}

func (s *TaskServiceServer) UpdateTask(ctx context.Context, req *pb.Task) (*empty.Empty, error) {
	if existingTask, err := s.getTaskByIdAndOwner(req.Id, req.Owner); err != nil {
		return nil, err
	} else {
		if req.Body != "" {
			existingTask.Body = req.Body
		}
		if req.Status != pb.Status_DEFAULT {
			existingTask.Status = req.Status
		}

		if result := db.DB.Save(existingTask); result.Error != nil {
			return nil, result.Error
		}

		return nil, nil
	}
}

func (s *TaskServiceServer) DeleteTask(ctx context.Context, req *pb.TaskCreds) (*empty.Empty, error) {
	if existingTask, err := s.getTaskByIdAndOwner(req.Id, req.Owner); err != nil {
		return nil, err
	} else {
		if result := db.DB.Delete(existingTask); result.Error != nil {
			return nil, result.Error
		}

		return nil, nil
	}
}

func (s *TaskServiceServer) GetTask(ctx context.Context, req *pb.TaskCreds) (*pb.Task, error) {
	if existingTask, err := s.getTaskByIdAndOwner(req.Id, req.Owner); err != nil {
		return nil, err
	} else {
		return existingTask, nil
	}
}

func (s *TaskServiceServer) GetTasksByPage(ctx context.Context, req *pb.PageRequest) (*pb.PageResponse, error) {
	offset := (req.Number - 1) * req.Size

	var tasks []*pb.Task
	if result := db.DB.Order("id DESC").Limit(int(req.Size)).Offset(int(offset)).Find(&tasks); result.Error != nil {
		return nil, result.Error
	}

	return &pb.PageResponse{Tasks: tasks}, nil
}
