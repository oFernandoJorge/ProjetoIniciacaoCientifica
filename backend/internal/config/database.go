package config
import (
"fmt"
"log"
"os"
"github.com/joho/godotenv"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)
// DB é instância global do banco
var DB *gorm.DB
// ConnectDatabase realiza conexão com PostgreSQL
func ConnectDatabase() {
// Carrega variáveis do .env
if err := godotenv.Load(); err != nil {
log.Println("[WARN] .env not found, using system env vars")
}
host := os.Getenv("DB_HOST")
port := os.Getenv("DB_PORT")
user := os.Getenv("DB_USER")
password := os.Getenv("DB_PASSWORD")
dbname := os.Getenv("DB_NAME")
sslmode := os.Getenv("DB_SSLMODE")
// DSN PostgreSQL
dsn := fmt.Sprintf(
"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
host,
user,
password,
dbname,
port,
sslmode,
)
database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
log.Fatal("[ERROR] failed to connect database:", err)
}
DB = database
log.Println("Database connection established")
}