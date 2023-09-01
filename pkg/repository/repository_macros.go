package repository

import "golang.org/x/exp/slog"

func (r *Repository) CreateUserIfNotExists(id int64, username string) {
	exists, err := r.ExistsUser(id)
	if err != nil {
		slog.Error("Check user exists error: " + err.Error())
	} else {
		if exists {
			slog.Info("Connected old user: " + username)
		} else {
			slog.Info("Connected new user: " + username)
			if err = r.CreateUser(id, username); err != nil {
				slog.Info("creating new user error: " + err.Error())
			}
		}
	}
}
