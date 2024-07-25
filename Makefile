OUT_DIR := ./generated/gamesManagerService
FILE_PATH := ./services/games_manager_service.proto

clean_generated:
	@if [ -d $(OUT_DIR) ]; then \
		echo "Removing folder: $(OUT_DIR)"; \
		rm -rf $(OUT_DIR); \
	fi

create_generated: 
	mkdir -p generated

generate_grpc:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
    --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
    $(FILE_PATH)

generate: clean_generated create_generated generate_grpc