package session

import (
	"github.com/gofiber/fiber/v3"
)

// Get fetches a value from the session and casts it to T, returning the value and a boolean
// indicating whether the value was present in the session.
func Get[T any](c fiber.Ctx, key string) (T, bool) {
	if value, found := get(c, key).(T); found {
		return value, true
	}
	var value T
	return value, false
}

// GetString fetches a string from the session, returning the empty string if it isn't found. If you
// need to be able to differentiate between absence and the empty string, use Get.
func GetString(c fiber.Ctx, key string) string {
	return GetOrDefault[string](c, key)
}

// GetInt fetches an int64 from the session, returning 0 if it isn't found. If you need to be able to
// differentiate between absence and 0, use Get.
func GetInt(c fiber.Ctx, key string) int64 {
	return GetOrDefault[int64](c, key)
}

// GetUint fetches a uint64 from the session, returning 0 if it isn't found. If you need to be able to
// differentiate between absence and 0, use Get.
func GetUint(c fiber.Ctx, key string) uint64 {
	return GetOrDefault[uint64](c, key)
}

// GetOrDefault fetches a value from the session and casts it to T, returning the value or the
// default value for the type if it was not present in the session.
func GetOrDefault[T any](c fiber.Ctx, key string) T {
	value, _ := Get[T](c, key)
	return value
}

// GetOnce fetches a value from the session and casts it to T, then deletes it from the session.
func GetOnce[T any](c fiber.Ctx, key string) (T, bool) {
	if value, found := getOnce(c, key).(T); found {
		return value, true
	}
	var value T
	return value, false
}

// GetStringOnce fetches a string from the session and then deletes it from the session. If you need
// to be able to differentiate between absence and the empty string, use GetOnce.
func GetStringOnce(c fiber.Ctx, key string) string {
	return GetOnceOrDefault[string](c, key)
}

// GetIntOnce fetches an int64 from the session and then deletes it from the session. If you
// need to be able to differentiate between absence and 0, use GetOnce.
func GetIntOnce(c fiber.Ctx, key string) int64 {
	return GetOnceOrDefault[int64](c, key)
}

// GetUintOnce fetches an int64 from the session and then deletes it from the session. If you
// need to be able to differentiate between absence and 0, use GetOnce.
func GetUintOnce(c fiber.Ctx, key string) uint64 {
	return GetOnceOrDefault[uint64](c, key)
}

// GetOnceOrDefault fetches a value from the session and casts it to T, returning the value or the
// default value for the type if it was not present in the session.
func GetOnceOrDefault[T any](c fiber.Ctx, key string) T {
	value, _ := GetOnce[T](c, key)
	return value
}

// Set stores a value in the session. If T is a custom type then it may need to be registered with
// gob.Register first.
func Set[T any](c fiber.Ctx, key string, value T) {
	session := must(s.Get(c))
	session.Set(key, value)
	must0(session.Save())
}

// Delete deletes a value from the session.
func Delete(c fiber.Ctx, key string) {
	session := must(s.Get(c))
	session.Delete(key)
}

// Destroy destroys the session.
func Destroy(c fiber.Ctx) {
	session := must(s.Get(c))
	must0(session.Destroy())
}

// GetID returns the session ID.
func GetID(c fiber.Ctx) string {
	return must(s.Get(c)).ID()
}

func get(c fiber.Ctx, key string) any {
	session := must(s.Get(c))
	return session.Get(key)
}

func getOnce(c fiber.Ctx, key string) any {
	session := must(s.Get(c))
	value := session.Get(key)
	session.Delete(key)
	must0(session.Save())
	return value
}
