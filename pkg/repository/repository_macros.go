package repository

import "golang.org/x/exp/slog"

func (r *Repository) CreateUserIfNotExists(id int64, username string) {
	exists, err := r.ExistsUser(id)
	if err != nil {
		slog.Error("check user exists error: " + err.Error())
	} else {
		if exists {
			slog.Info("connected old user: " + username)
		} else {
			slog.Info("connected new user: " + username)
			if err = r.CreateUser(id, username); err != nil {
				slog.Info("creating new user error: " + err.Error())
			}
		}
	}
}
