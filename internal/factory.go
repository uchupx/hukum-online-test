package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/uchupx/geek-garden-test/config"
	"github.com/uchupx/geek-garden-test/internal/handler"
	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/internal/repo/cache"
	"github.com/uchupx/geek-garden-test/internal/repo/store"
	"github.com/uchupx/geek-garden-test/internal/service"
	"github.com/uchupx/kajian-api/pkg/db"
	"github.com/uchupx/kajian-api/pkg/mysql"
	redisConn "github.com/uchupx/kajian-api/pkg/redis"
)

type Factory struct {
	mysqlConn   *db.DB
	redisClient *redis.Client

	// Member
	memberHandler handler.Handler
	memberRepo    repo.MemberRepo
	memberService *service.MemberService

	// Attende
	attendeeHandler handler.Handler
	attendeeRepo    repo.AttendeeRepo
	attendeeService *service.AttendeeService

	// Gathering
	gatheringHandler handler.Handler
	gatheringRepo    repo.GatheringRepo
	gatheringService *service.GatheringService

	// Invitation
	invitationHandler handler.Handler
	invitationRepo    repo.InvitationRepo
	invitationService *service.InvitationService
}

func (f *Factory) MySQL() *db.DB {
	if f.mysqlConn == nil {
		conn, err := mysql.NewConnection(mysql.DBPayload{
			Host:     config.GetConfig().Database.Host,
			Port:     config.GetConfig().Database.Port,
			Database: config.GetConfig().Database.Name,
			Username: config.GetConfig().Database.User,
			Password: config.GetConfig().Database.Password,
		})
		if err != nil {
			panic(err)
		}

		f.mysqlConn = conn
	}

	return f.mysqlConn
}

func (f *Factory) Redis() *redis.Client {
	if f.redisClient == nil {
		conn, err := redisConn.NewRedisConn(redisConn.RedisConfig{
			Host:     config.GetConfig().Redis.Host,
			Database: int(config.GetConfig().Redis.Database),
			Password: config.GetConfig().Redis.Password,
		})
		if err != nil {
			panic(err)
		}

		f.redisClient = conn
	}

	return f.redisClient
}

func (f *Factory) MemberHandler() handler.Handler {
	if f.memberHandler == nil {
		f.memberHandler = handler.MemberHandler{
			Service: f.MemberService(),
		}
	}

	return f.memberHandler
}

func (f *Factory) MemberRepo() repo.MemberRepo {
	if f.memberRepo == nil {
		f.memberRepo = cache.NewMemberCacheRepo(store.NewMemberStoreRepo(f.MySQL()))
	}

	return f.memberRepo
}

func (f *Factory) MemberService() *service.MemberService {
	if f.memberService == nil {
		f.memberService = &service.MemberService{
			MemberRepo: f.MemberRepo(),
		}
	}

	return f.memberService
}

func (f *Factory) AttendeeHandler() handler.Handler {
	if f.attendeeHandler == nil {
		f.attendeeHandler = handler.AttendeeHandler{
			Service: f.AttendeeService(),
		}
	}

	return f.attendeeHandler
}

func (f *Factory) AttendeeRepo() repo.AttendeeRepo {
	if f.attendeeRepo == nil {
		f.attendeeRepo = cache.NewAttendeecacheRepo(store.NewAttendeeStoreRepo(f.MySQL()))
	}

	return f.attendeeRepo
}

func (f *Factory) AttendeeService() *service.AttendeeService {
	if f.attendeeService == nil {
		f.attendeeService = &service.AttendeeService{
			AttendeeRepo: f.AttendeeRepo(),
		}
	}

	return f.attendeeService
}

func (f *Factory) GatheringHandler() handler.Handler {
	if f.gatheringHandler == nil {
		f.gatheringHandler = handler.GatheringHandler{
			Service: f.GatheringService(),
		}
	}

	return f.gatheringHandler
}

func (f *Factory) GatheringRepo() repo.GatheringRepo {
	if f.gatheringRepo == nil {
		f.gatheringRepo = cache.NewGatheringcacheRepo(store.NewGatheringStore(f.MySQL()))
	}

	return f.gatheringRepo
}

func (f *Factory) GatheringService() *service.GatheringService {
	if f.gatheringService == nil {
		f.gatheringService = &service.GatheringService{
			GatheringRepo: f.GatheringRepo(),
		}
	}

	return f.gatheringService
}

func (f *Factory) InvitationHandler() handler.Handler {
	if f.invitationHandler == nil {
		f.invitationHandler = handler.InvitationHandler{
			Service: f.InvitationService(),
		}
	}

	return f.invitationHandler
}

func (f *Factory) InvitationRepo() repo.InvitationRepo {
	if f.invitationRepo == nil {
		f.invitationRepo = cache.NewInvitationcacheRepo(store.NewInvitationStoreRepo(f.MySQL()))
	}

	return f.invitationRepo
}

func (f *Factory) InvitationService() *service.InvitationService {
	if f.invitationService == nil {
		f.invitationService = &service.InvitationService{
			Repo:          f.InvitationRepo(),
			AttendRepo:    f.AttendeeRepo(),
			MemberRepo:    f.MemberRepo(),
			GatheringRepo: f.GatheringRepo(),
			DB:            f.MySQL(),
		}
	}

	return f.invitationService
}

func (f Factory) InitRoute(e *gin.RouterGroup) {
	routes := []handler.Handler{
		f.MemberHandler(),
		f.GatheringHandler(),
		f.AttendeeHandler(),
		f.InvitationHandler(),
	}

	for _, route := range routes {
		route.Routes(e)
	}
}
