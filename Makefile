generate:
	protoc --go_out=./generated/gameRoomManagerService --go_opt=paths=source_relative \
    --go-grpc_out=./generated/gameRoomManagerService --go-grpc_opt=paths=source_relative \
    ./services/game_room_manager_service.proto

