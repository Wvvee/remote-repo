package router

import (
	"gin-ranking/config"
	"gin-ranking/controllers"
	"gin-ranking/pkg/logger"
	"log"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	store, err := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "Ssh123", []byte("secret"))
	if err != nil {
		log.Fatalf("session_redis.NewStore:%v", err)
	}
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
	}

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)

	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("/ranking", controllers.PlayerController{}.GetRanking)

	return r
}
