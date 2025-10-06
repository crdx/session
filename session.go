package session

import (
	"github.com/gofiber/fiber/v3"
)

// Get fetches a value from the session as T. If the session doesn't contain a T then the zero value
// of T is returned.
func Get[T any](c fiber.Ctx, key string) T {
	value, _ := TryGet[T](c, key)
	return value
}

// GetOnce fetches a value from the session as T, then deletes it. If the session doesn't contain a
// T then the zero value of T is returned.
func GetOnce[T any](c fiber.Ctx, key string) T {
	value, _ := TryGetOnce[T](c, key)
	return value
}

// TryGet fetches a value from the session as T, returning the value and a boolean indicating
// whether the value was present in the session.
func TryGet[T any](c fiber.Ctx, key string) (T, bool) {
	if value, found := get(c, key).(T); found {
		return value, true
	}
	var value T
	return value, false
}

// TryGetOnce fetches a value from the session as T, then deletes it from the session.
func TryGetOnce[T any](c fiber.Ctx, key string) (T, bool) {
	if value, found := getOnce(c, key).(T); found {
		return value, true
	}
	var value T
	return value, false
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
