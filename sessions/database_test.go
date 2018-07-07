package sessions_test

import (
	"reflect"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/teamlint/iris/sessions"
	"github.com/teamlint/iris/sessions/sessiondb/badger"
	"github.com/teamlint/iris/sessions/sessiondb/boltdb"
	"github.com/teamlint/iris/sessions/sessiondb/redis"
	"github.com/teamlint/iris/sessions/sessiondb/redis/service"
)

func TestDatabase(t *testing.T) {
	var db sessions.Database
	// mem
	db = sessions.NewMemDB()
	logTitle(t, "MemDB")
	testInterfacer(t, db)
	// Badger
	logTitle(t, "BadgerDB")
	db, _ = badger.New("./testdb/badger")
	testInterfacer(t, db)
	// boltdb
	logTitle(t, "BoltDB")
	db, _ = boltdb.New("./testdb/bolt.db", 0600)
	testInterfacer(t, db)
	// Redis
	logTitle(t, "RedisDB")
	db = redis.New(service.Config{
		Network:     service.DefaultRedisNetwork,
		Addr:        service.DefaultRedisAddr,
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      "",
	})
	testInterfacer(t, db)
	//
}
func testInterfacer(t *testing.T, db sessions.Database) {
	type Model struct {
		ID         int
		Name       string
		Height     float64
		IsApproved bool
		CreatedAt  time.Time
		DeletedAt  *time.Time
	}
	uuid, _ := uuid.NewV4()
	sid := uuid.String()
	intValue := 10
	stringValue := "session read test"
	floatValue := 3.1415926
	boolValue1 := true
	boolValue2 := false
	timeValue, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-07 11:37:45", time.Local)
	structValue := Model{ID: intValue, Name: stringValue, Height: floatValue, IsApproved: boolValue1, CreatedAt: timeValue, DeletedAt: &timeValue}
	sliceIntValue := []int{1, 2, 3, 4, 5, 6}
	mapStringIntValue := map[string]int{"k1": 1, "k2": 2, "k3": 3, "k4": 4}
	sliceStructValue := make([]Model, 0)
	sliceStructValue = append(sliceStructValue, Model{ID: 1, Name: "struct1"})
	sliceStructValue = append(sliceStructValue, Model{ID: 2, Name: "struct2"})

	values := map[string]interface{}{
		"IntKey":           intValue,
		"IntPKey":          &intValue,
		"StringKey":        stringValue,
		"StringPKey":       &stringValue,
		"FloatKey":         floatValue,
		"FloatPKey":        &floatValue,
		"BoolKey1":         boolValue1,
		"BoolPKey1":        &boolValue1,
		"BoolKey2":         boolValue2,
		"BoolPKey2":        &boolValue2,
		"TimeKey":          timeValue,
		"TimePKey":         &timeValue,
		"StructKey":        structValue,
		"StructPKey":       &structValue,
		"SliceIntKey":      sliceIntValue,
		"SliceIntPKey":     &sliceIntValue,
		"SliceStructKey":   sliceStructValue,
		"SliceStructPKey":  &sliceStructValue,
		"MapStringIntKey":  mapStringIntValue,
		"MapStringIntPKey": &mapStringIntValue,
	}

	var err error
	expires := time.Hour * 24
	lifetime := db.Acquire(sid, expires)
	for k, v := range values {
		db.Set(sid, lifetime, k, v, false)
		switch k {
		case "IntKey", "IntPKey":
			var act int
			err = db.Read(sid, k, &act)
			log(t, err, k, v, act)
		case "TimeKey", "TimePKey":
			var act time.Time
			err = db.Read(sid, k, &act)
			log(t, err, k, v, act)
		case "StructKey", "StructPKey":
			// var act Model
			var act = new(Model)
			err = db.Read(sid, k, act)
			log(t, err, k, v, *act)
		case "SliceIntKey", "SliceIntPKey":
			var act []int
			err = db.Read(sid, k, &act)
			log(t, err, k, v, act)
		case "SliceStructKey", "SliceStructPKey":
			var act []Model
			err = db.Read(sid, k, &act)
			log(t, err, k, v, act)
		case "MapStringIntKey", "MapStringIntPKey":
			var act map[string]int
			err = db.Read(sid, k, &act)
			log(t, err, k, v, act)
		default:
			var act interface{}
			err = db.Read(sid, k, &act)
			log(t, err, k, v, act)
		}
		t.Log("---------------------------------------------------------------------------")
	}
}
func log(t *testing.T, err error, k string, v interface{}, act interface{}) {
	t.Logf("[key:%s]\texpected type: %T\t\tactual type: %T", k, v, act)
	assert.Equal(t, nil, err)
	curr := reflect.ValueOf(v)
	if curr.Kind() == reflect.Ptr {
		curr = curr.Elem()
	}
	currAct := reflect.ValueOf(act)
	if currAct.Kind() == reflect.Ptr {
		currAct = currAct.Elem()
	}
	t.Logf("[key:%s]\texpected value: %v\tactual value: %v", k, curr.Interface(), currAct.Interface())
	assert.Equal(t, curr.Interface(), act)
}
func logTitle(t *testing.T, title string) {
	t.Logf("=======================================[%s]=====================================", title)
}
