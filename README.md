# Competition Corner App

![Wide window displaying the table title and top scores including user avatars](https://github.com/Billiam/competition-corner-desktop/assets/242008/2a6b0574-166d-4143-8800-76149ad32b7b)
![Skinnier window with more compact view, hinding the user avatar](https://github.com/Billiam/competition-corner-desktop/assets/242008/0c08cc86-db3f-4895-99fe-1880dae132a1)

This app displays and periodically refreshes the Virtual Pinball Chat weekly competition leaderboards:

https://virtualpinballchat.com/#/weekly-rankings/competition-corner

## Configuration

Configuration is stored in json, located at:

* Linux: `~/.config/competition-corner`
* Windows `%localappdata%/competition-corner/`

### Value dictionary:

* `width`: Window width
* `height`: Window height
* `x`: Window X position (not used)
* `y`: Window Y position (not used)
* `me`: Discord username. Matching user score will be visually highlighted

## Help

> How do I limit the number of scores displayed?

Resize the window so that you don't see all of the scores

> How do I see more scores at once

Resize the window so that you can see more of the scores

## Development

Install golang and wails: https://wails.io/docs/gettingstarted/installation/

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

