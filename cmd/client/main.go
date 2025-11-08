package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	pb "movies-golang/gen/go/csv_data"
)

func main() {
	fmt.Println("Reading ratings")
	file, err := os.Open(".data/ratings_small.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error at closing file: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't read the record: %v", err)
		}
		userId, err := strconv.ParseInt(record[0], 10, 32)
		if err != nil {
			log.Printf("Error parsing userId '%s': %v", record[0], err)
			continue
		}

		movieId, err := strconv.ParseInt(record[1], 10, 32)
		if err != nil {
			log.Printf("Error parsing movieId '%s': %v", record[1], err)
			continue
		}

		ratingVal, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			log.Printf("Error parsing rating '%s': %v", record[2], err)
			continue
		}

		timestampVal, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			log.Printf("Error parsing timestamp '%s': %v", record[3], err)
			continue
		}
		goTime := time.Unix(timestampVal, 0)
		protoTimestamp := timestamppb.New(goTime)

		rating := &pb.Rating{
			UserId:    int32(userId),
			MovieId:   int32(movieId),
			Rating:    float32(ratingVal),
			Timestamp: protoTimestamp,
		}

		fmt.Printf("Leído y convertido: %s\n", rating.String())

	}
}
