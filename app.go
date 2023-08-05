package main

import (
  "context"
  "github.com/spf13/viper"
  "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
  ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
  return &App{}
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
  cx, cy := runtime.WindowGetPosition(a.ctx)
  cw, ch := runtime.WindowGetSize(a.ctx)

  viper.Set("x", cx)
  viper.Set("y", cy)
  viper.Set("width", cw)
  viper.Set("height", ch)

  err := viper.WriteConfig()
  if err != nil {
    runtime.LogFatal(a.ctx, err.Error())
  }
  return false
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
  a.ctx = ctx
}

func (a *App) ExitApp() {
  runtime.Quit(a.ctx)
}
func (a *App) domReady(ctx context.Context) {
// unreliable on multimonitor setups
//   cx := viper.GetInt("x")
//   cy := viper.GetInt("y")
//
  runtime.WindowShow(a.ctx)
//   runtime.WindowSetPosition(a.ctx, cx, cy)
}

func (a *App) GetUser() (user string) {
  return viper.GetString("me")
}
