package controllers

import (
	"context"
	"fmt"

	"statisticsservice/db"
	"statisticsservice/protos"
)

type StatisticsServiceServer struct {
	protos.UnimplementedStatisticsServiceServer
}

func (s *StatisticsServiceServer) GetTotalViewsLikesCount(ctx context.Context, req *protos.SpecificTaskRequest) (*protos.SpecificTaskResponse, error) {
	queryViews := fmt.Sprintf(`
        SELECT COUNT(*) AS views_count
        FROM %s
        WHERE event_type = 'View' AND task_id = ?
    `, db.TableName)
	queryLikes := fmt.Sprintf(`
        SELECT COUNT(*) AS likes_count
        FROM %s
        WHERE event_type = 'Like' AND task_id = ?
    `, db.TableName)

	resp := protos.SpecificTaskResponse{TaskID: req.TaskID}

	err := db.DB.QueryRow(queryViews, req.TaskID).Scan(&resp.ViewsCount)
	if err != nil {
		return nil, err
	}
	err = db.DB.QueryRow(queryLikes, req.TaskID).Scan(&resp.LikesCount)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *StatisticsServiceServer) GetTopTasks(ctx context.Context, req *protos.TopRequest) (*protos.TopTasksResponse, error) {
	query := fmt.Sprintf(`
        SELECT task_id, username, COUNT(*) AS count
        FROM %s
        WHERE event_type = '%s'
        GROUP BY task_id
        ORDER BY count DESC
        LIMIT 5
    `, db.TableName, req.Type.String())

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := protos.TopTasksResponse{}
	for rows.Next() {
		task := protos.TaskResponse{StatisticsType: req.Type}
		if err := rows.Scan(&task.TaskID, &task.Username, &task.StatisticsCount); err != nil {
			return nil, err
		}
		resp.Tasks = append(resp.Tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *StatisticsServiceServer) GetTopUsers(ctx context.Context, req *protos.TopRequest) (*protos.TopUsersResponse, error) {
	query := fmt.Sprintf(`
        SELECT username, COUNT(*) AS count
        FROM %s
        WHERE event_type = '%s'
        GROUP BY user_id
        ORDER BY count DESC
        LIMIT 3
    `, db.TableName, req.Type.String())

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := protos.TopUsersResponse{}
	for rows.Next() {
		user := protos.UserResponse{StatisticsType: req.Type}
		if err := rows.Scan(&user.Username, &user.StatisticsCount); err != nil {
			return nil, err
		}
		resp.Users = append(resp.Users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &resp, nil
}
