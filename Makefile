OUT_DIR := ./generated/gameRoomManagerService

generate:
	protoc --go_out=${OUT_DIR} --go_opt=paths=source_relative \
    --go-grpc_out=${OUT_DIR} --go-grpc_opt=paths=source_relative \
    ./services/game_room_manager_service.proto

