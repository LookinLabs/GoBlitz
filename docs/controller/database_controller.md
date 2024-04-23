# Example controller with basic SELECT SQL Query

1. Create `model/get_users_struct.go` structure in model folder

```
package model

type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}
```

2. Create API Controller in `controller/api/get_users.go`, which uses this struct to fill the data

```
package api

import (
	"net/http"
	model "web/model"
	"web/repository"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := repository.DB.Query("SELECT id, name FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
```

3. Add into middleware your API Path

```
	// API Handling
	apiGroup := router.Group(os.Getenv("API_PATH"))
	{
		apiGroup.GET("/users", api.GetUsers)
	}
```

4. Setup environment and run migrations

**Note!** Check that you have PostgreSQL enabled and running

`make`

`make migrate-up`

5. Make an request against `/api/v1/users`

```
curl http://localhost:8000/api/v1/users
[{"user_id":"1","name":"Alice"}]⏎   
```
