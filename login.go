package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

type Pengguna struct {
	Nama     string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func hashPassword(teks string) string {
	hash := sha1.New()
	hash.Write([]byte(teks))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	app := fiber.New()

	redisDB := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var dataLogin struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&dataLogin); err != nil {
			return c.Status(400).SendString("Format data salah")
		}

		kunci := "login_" + dataLogin.Username
		dataPengguna, err := redisDB.Get(ctx, kunci).Result()
		if err != nil {
			return c.Status(401).SendString("Pengguna tidak ditemukan")
		}

		var pengguna Pengguna
		if err := json.Unmarshal([]byte(dataPengguna), &pengguna); err != nil {
			return c.Status(500).SendString("Gagal membaca data pengguna")
		}

		if pengguna.Password != hashPassword(dataLogin.Password) {
			return c.Status(401).SendString("Password salah")
		}

		return c.SendString("Login berhasil. Halo, " + pengguna.Nama)
	})

	log.Fatal(app.Listen(":3000"))
}
