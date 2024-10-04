package database

import (
	"log"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/config"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	SQLite3   *gorm.DB
	AppConfig *config.AppConfig
}

var dbConn = &DB{}

// db tables
var users *models.User
var messages *models.Message

func ConnectDB() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("GO_live_chat.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	dbConn.SQLite3 = db
	testDB, _ := db.DB()
	err = testDB.Ping()
	if err != nil {
		return nil, err
	}
	db.Exec("DROP TABLE users")
	db.Exec("DROP TABLE messages")
	err = execMigrations(db)
	if err != nil {
		return dbConn, err
	}
	err = execSeeder(db)
	if err != nil {
		return dbConn, err
	}
	return dbConn, nil
}

// execMigrations execute all the migrations
func execMigrations(db *gorm.DB) error {
	// add all models's structs to AutoMigrate
	err := db.AutoMigrate(&users, &messages)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func execSeeder(db *gorm.DB) error {

	// all passwords are: password
	users := []*models.User{
		{
			UserName:    "admin",
			Email:       "admin@admin.com",
			Password:    "$2a$12$39JTEON1eLjhQ4uHq/L8SuQNn9kUgqCuCA3LmSZ3A9iJK6Ay82VvC",
			AccessLevel: "3",
		},
		{
			UserName:    "moderator",
			Email:       "moderator@moderator.com",
			Password:    "$2a$12$yMUW6GklJCw3ehtbs9kDQ.AtlTYPCLnimGNgWN6BH9bjvAOlXge1G",
			AccessLevel: "2",
		},
		{
			UserName: "user",
			Email:    "user@user.com",
			Password: "$2a$12$JCdNB2z/3YwQhUjd1TVlDeaf4ULeNmNoKcj1V6qWUUFKjkC7b.q2q",
		},
	}

	// dummy chat
	messages := []*models.Message{
		{
			UserID:  1,
			Content: "I'm just replaying The Legend of Zelda: A Link to the Past for the umpteenth time. It's timeless! The pixel art, the music, the dungeons... it's all so perfect.",
		},
		{
			UserID:  2,
			Content: "A Link to the Past is definitely a classic, but I'm more of a fan of the 3D era. Ocarina of Time changed the course of gaming for me. The sense of scale, the exploration, the characters... it was groundbreaking.",
		},
		{
			UserID:  3,
			Content: "I have to agree. Ocarina of Time is a masterpiece. But for me, Majora's Mask takes the cake. The dark tone, the time-traveling mechanic, the emotional impact... it's a unique and unforgettable experience.",
		},
		{
			UserID:  1,
			Content: "Majora's Mask is definitely a bold departure from the series. I love the way it tackles heavier themes. But I still prefer the classic top-down style.",
		},
		{
			UserID:  2,
			Content: "I think Breath of the Wild is the best Zelda game ever made. The open world, the freedom to explore, the challenging combat... it's a game-changer.",
		},
		{
			UserID:  3,
			Content: "Breath of the Wild is amazing, but I still have a soft spot for the older games. Link's Awakening is a hidden gem. The small island setting, the charming characters, the puzzle-solving... it's a perfect blend of adventure and nostalgia.",
		},
		{
			UserID:  1,
			Content: "I haven't played Link's Awakening in years. Maybe I should give it another try. I remember loving the game as a kid.",
		},
		{
			UserID:  2,
			Content: "I've never played Link's Awakening. I should check it out. I've heard great things about it.",
		},
		{
			UserID:  3,
			Content: "You won't regret it. It's a timeless classic. But we can't forget about The Wind Waker. The cel-shaded art style, the sailing mechanic, the whimsical atmosphere... it's a magical experience.",
		},
		{
			UserID:  1,
			Content: "The Wind Waker is so underrated. I love the art style. It's so unique and charming.",
		},
		{
			UserID:  2,
			Content: "I'm not a huge fan of the cel-shaded style, but I do enjoy The Wind Waker. The sailing mechanic is a lot of fun.",
		},
		{
			UserID:  3,
			Content: "I think the cel-shaded style adds a lot of charm to the game. It's like a cartoon adventure.",
		},
		{
			UserID:  1,
			Content: "I'm curious to hear your thoughts on Twilight Princess. It's one of my favorite Zelda games.",
		},
		{
			UserID:  2,
			Content: "Twilight Princess is a great game. I love the dark atmosphere and the mature tone.",
		},
		{
			UserID:  3,
			Content: "Twilight Princess is a bit too serious for me. I prefer the more lighthearted and whimsical games.",
		},
		{
			UserID:  1,
			Content: "I think Twilight Princess strikes a good balance between light and dark. It's a well-rounded game.",
		},
		{
			UserID:  2,
			Content: "What about Skyward Sword? It's a divisive game, but I enjoyed it. The motion controls were a bit clunky, but the story and gameplay were solid.",
		},
		{
			UserID:  3,
			Content: "I'm not a fan of Skyward Sword. The motion controls were a major turn-off for me.",
		},
		{
			UserID:  1,
			Content: "I think Skyward Sword is overrated. The story is weak, and the gameplay is repetitive.",
		},
		{
			UserID:  2,
			Content: "I disagree. I thought the story was interesting, and the gameplay was challenging.",
		},
		{
			UserID:  3,
			Content: "I think A Link Between Worlds is a great game. The ability to switch between 2D and 3D is a cool mechanic.",
		},
		{
			UserID:  1,
			Content: "A Link Between Worlds is a lot of fun. I love the puzzles and the exploration.",
		},
		{
			UserID:  2,
			Content: "I haven't played A Link Between Worlds yet. I should give it a try.",
		},
		{
			UserID:  3,
			Content: "You won't regret it. It's a great game.",
		},
		{
			UserID:  1,
			Content: "So, what's your favorite Zelda game of all time?",
		},
		{
			UserID:  2,
			Content: "It's a tough choice, but I'd have to say Ocarina of Time. It's a classic that has stood the test of time.",
		},
		{
			UserID:  3,
			Content: "I'd have to go with Majora's Mask. It's a unique and unforgettable experience.",
		},
		{
			UserID:  1,
			Content: "I'm torn between A Link to the Past and Twilight Princess. They're both amazing games.",
		},
		{
			UserID:  2,
			Content: "I think we can all agree that the Zelda series is one of the greatest franchises of all time.",
		},
		{
			UserID:  3,
			Content: "Absolutely. There's something for everyone in this series.",
		},
	}

	result := db.Create(users)
	if err := result.Error; err != nil {
		log.Println(err)
		return err
	}
	result = db.Create(messages)
	if err := result.Error; err != nil {
		log.Println(err)
		return err
	}
	log.Println("seeded users and chat")
	return nil
}
