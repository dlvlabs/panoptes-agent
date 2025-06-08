package telegram

import (
  "context"
  "github.com/go-telegram/bot"
)

type TelegramClient struct {
  ctx      context.Context
  botToken string
  chatID   string
  client   *bot.Bot
}
