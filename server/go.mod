module github.com/conradludgate/playlist/server

go 1.12

replace github.com/conradludgate/playlist/server/exchanges v0.0.0 => ./exchanges

replace github.com/conradludgate/playlist/server/exchanges/spotify v0.0.0 => ./exchanges/spotify

require (
	github.com/conradludgate/playlist/server/exchanges v0.0.0
	github.com/conradludgate/playlist/server/exchanges/spotify v0.0.0
	github.com/gorilla/mux v1.7.3
	github.com/joho/godotenv v1.3.0
)
