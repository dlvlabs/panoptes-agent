package telegram

import (
  "context"
  "fmt"

  "github.com/go-telegram/bot"
)

func NewTelegramClient(ctx context.Context, botToken string, chatID string) (*TelegramClient, error) {
  telegramBot, err := bot.New(botToken)
  if err != nil {
    return nil, fmt.Errorf("failed to create telegram client: %w", err)
  }

  return &TelegramClient{
    ctx:      ctx,
    botToken: botToken,
    chatID:   chatID,
    client:   telegramBot,
  }, nil
}

func (tc *TelegramClient) Close() error {
  tc.client.Close(tc.ctx)
  return nil
}

func (tc *TelegramClient) SendMessage(message string) error {
  _, err := tc.client.SendMessage(tc.ctx, &bot.SendMessageParams{
    ChatID: tc.chatID,
    Text:   message,
  })
  return err
}
