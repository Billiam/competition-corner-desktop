# Competition Corner App
![Wide window displaying the table title and top scores including user avatars](https://github.com/Billiam/competition-corner-desktop/assets/242008/e927b8e8-ff73-4123-8c58-aff2662da296)
![Skinnier window with more compact view, hinding the user avatar](https://github.com/Billiam/competition-corner-desktop/assets/242008/0ba26684-7ab3-407b-8c7f-4c1460b1ae2d)


This app displays and periodically refreshes the Virtual Pinball Chat weekly competition leaderboards:

https://virtualpinballchat.com/#/weekly-rankings/competition-corner

## Development

Install golang and wails: https://wails.io/docs/gettingstarted/installation/

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

