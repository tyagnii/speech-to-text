package cmd

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	sttservice "github.com/shipherman/speech-to-text/gen/stt/service/v1"
	"github.com/shipherman/speech-to-text/internal/clients"
	"github.com/shipherman/speech-to-text/internal/db"
	"github.com/shipherman/speech-to-text/internal/models"
	"github.com/shipherman/speech-to-text/internal/services/auth"
	"github.com/shipherman/speech-to-text/pkg/audioconverter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TranscribeServer struct {
	sttservice.UnimplementedSttServiceServer
	DBClient db.Connector
	Store    models.Store
	auth     auth.Auth
}

// TranscribeAudio provides main functionality of service
// It adds new audio to the pipeline, sends processing statuses to client
// and provides result of transcription to client
func (t *TranscribeServer) TranscribeAudio(
	audio *sttservice.Audio,
	stream sttservice.SttService_TranscribeAudioServer,
) error {
	// TODO
	// Implement pipline for processing audio
	// 1. Check file, if it's a valid wav file
	//		send status ACCEPTED
	// 2. Put file to the worker pool for saving
	// 		send status ORDERED
	// 3. Transcribe file
	// 		send status DONE
	// *send status INVALID on invalid files
	//
	// Additional
	//	convert *.ogg to *.wav [for telegram bot]

	var response *sttservice.Status

	// Check wave header and if its valid send status ACCEPTED
	_, err := audioconverter.CheckWAVHeader(audio.Audio)
	if err != nil {
		return err
	}
	response = &sttservice.Status{Status: sttservice.EnumStatus_STATUS_ACCEPTED}
	stream.Send(response)

	// Calculate hash sum
	h := md5.New()
	_, err = h.Write(audio.Audio)
	if err != nil {
		return err
	}
	// Use hex string as filename
	audioFileHashSum := hex.EncodeToString(h.Sum(nil))

	// Execute user email frome jwt
	// get User object from db

	// Save data to DB
	// email, err := t.auth.GetEmail(context.Background())
	// t.DBClient.SaveNewAudio(audioFileHashSum, t.Store)

	// Save audio to store
	err = t.Store.Save(string(audioFileHashSum), audio.Audio)
	if err != nil {
		return err
	}

	// Mark audio as ordered
	response = &sttservice.Status{Status: sttservice.EnumStatus_STATUS_ORDERED}
	stream.Send(response)

	// Call remote STT neural network service
	text, err := clients.ReqSTT(audio.Audio)
	if err != nil {
		return err
	}

	response = &sttservice.Status{Status: sttservice.EnumStatus_STATUS_DONE}
	stream.Send(response)

	log.Println(text)

	return nil
}

// TODO
// Implement inmemory results store
//
// GetResults executing requested audio trascription from memory and
// sends it to a client
func (t *TranscribeServer) GetResult(ctx context.Context,
	in *sttservice.Text,
) (*sttservice.Text, error) {
	return nil, nil
}

// TODO
// Implement request to DB
//
// GetHistory returns array of {audiohash: text} pairs to a client
func (t *TranscribeServer) GetHistory(ctx context.Context,
	in *sttservice.User,
) (*sttservice.History, error) {
	// t.DBClient.GetHistory()
	return nil, nil
}

// Register registers new user in stt service
func (t *TranscribeServer) Register(ctx context.Context,
	in *sttservice.RegisterRequest,
) (*sttservice.RegisterResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	userID, err := t.auth.RegisterNewUser(ctx, in.Username, in.Email, in.Password)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &sttservice.RegisterResponse{UserId: userID}, nil
}

// Login authenticate users
func (t *TranscribeServer) Login(ctx context.Context, in *sttservice.LoginRequest) (*sttservice.LoginResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	fmt.Println("TServer Login")
	token, err := t.auth.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to login")
	}
	return &sttservice.LoginResponse{Token: token}, nil
}
