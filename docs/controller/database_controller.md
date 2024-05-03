# Example controller with basic SELECT SQL Query

1. Create `model/get_users.go` structure in model folder

```
package model

type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}
```

2. Create a SQL Query in `repository/db/get_users.go`, which makes an SQL Query against database

```
package sql

import (
	"web/model"
)

func GetUsers() ([]model.User, error) {
	rows, err := DB.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
```

3. Create API Controller in `controller/api/get_users.go`, which uses this struct to fill the data

```
package api

import (
	"net/http"
	sql "web/repository/db"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := sql.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

```

4. Add into middleware your API Path

```
	// API Handling
	apiGroup := router.Group(os.Getenv("API_PATH"))
	{
		apiGroup.GET("/users", api.GetUsers)
	}
```

5. Setup environment and run migrations

**Note!** Check that you have PostgreSQL enabled and running

`make`

`make migrate-up`

6. Make an request against `/api/v1/users`

```
curl http://localhost:8000/api/v1/users
[{"user_id":"1","name":"Alice"}]⏎   
```
