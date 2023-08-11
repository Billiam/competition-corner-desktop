package main

import (
  "embed"

  "github.com/wailsapp/wails/v2"
  "github.com/wailsapp/wails/v2/pkg/runtime"
  "github.com/wailsapp/wails/v2/pkg/options"
  "github.com/wailsapp/wails/v2/pkg/options/assetserver"
  "github.com/wailsapp/wails/v2/pkg/options/windows"
  "github.com/wailsapp/wails/v2/pkg/options/linux"
  "github.com/spf13/viper"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
  // Create an instance of the app structure
  app := NewApp()

  configErr := LoadConfig()
  if configErr != nil {
    runtime.LogError(app.ctx, configErr.Error())
  }

  // Create application with options
  err := wails.Run(&options.App{
    Title:  "competition-corner-wails",
    Width: viper.GetInt("width"),
    Height: viper.GetInt("height"),
    MinWidth: 200,
    MinHeight: 200,
    BackgroundColour: &options.RGBA{R: 0, G: 0, B: 5, A: 0},
    Frameless: true,
    OnBeforeClose: app.beforeClose,
    OnDomReady: app.domReady,
    StartHidden: true,
    AssetServer: &assetserver.Options{
      Assets: assets,
    },
    OnStartup: app.startup,
    Bind: []interface{}{
      app,
    },
    Windows: &windows.Options{
      WebviewIsTransparent: true,
      WindowIsTranslucent: true,
      BackdropType: windows.Auto,
      DisableFramelessWindowDecorations: false,
      Theme: windows.Dark,
    },
    Linux: &linux.Options{
      WindowIsTranslucent: true,
      Icon: icon,
    },
  })

  if err != nil {
    println("Error:", err.Error())
  }
}
