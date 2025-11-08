package proto

//go:generate protoc --proto_path=. --go_out=../gen/go --go_opt=paths=source_relative csv_data/ratings.proto
