package main

import (
	"context"
	"encounters-service/dto"
	"encounters-service/model"
	"encounters-service/proto/encounter"
	repository "encounters-service/repositories"
	"encounters-service/service"
	"fmt"
	"log"
	"net"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer host=explorer_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println(database)
	database.AutoMigrate(&model.Encounter{},
		&model.EncounterExecution{},
		&model.EncounterRequest{})

	err = database.AutoMigrate(&model.Encounter{}, &model.EncounterExecution{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}
	return database
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8087")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database := initDB()
	if database == nil {
		fmt.Println("FAILED TO CONNECT TO DB")
		return
	}

	encounterRequestRepo := &repository.EncounterRequestRepository{DatabaseConnection: database}
	encounterRequestService := &service.EncounterRequestService{Repo: encounterRequestRepo}

	encounterRepo := &repository.EncountersRepository{DatabaseConnection: database}
	encounterService := &service.EncounterService{Repo: encounterRepo}

	encounterExecutionRepo := &repository.EncountersExecutionRepository{DatabaseConnection: database}
	encounterExecutionService := &service.EncounterExecutionService{Repo: encounterExecutionRepo}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	encounter.RegisterEncounterServer(grpcServer, Server{encounterRequestService: encounterRequestService, encounterService: encounterService, encounterExecutionService: encounterExecutionService})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

type Server struct {
	encounter.UnimplementedEncounterServer
	encounterRequestService   *service.EncounterRequestService
	encounterService          *service.EncounterService
	encounterExecutionService *service.EncounterExecutionService
}

func (s Server) Create(ctx context.Context, request *encounter.CreateRequest) (*encounter.EncounterDto, error) {
	fmt.Println("HEY CREATE")
	enc := dtoFromRequest(*request.Encounter)
	p, ok := s.encounterService.Create(enc)
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error creating encounter.")
	}
	response := responsefromDto(p)
	return &response, nil
}
func (s Server) Update(ctx context.Context, request *encounter.UpdateRequest) (*encounter.EncounterDto, error) {
	fmt.Println("HEY UPDATE")
	enc := dtoFromRequest(*request.Encounter)
	p, ok := s.encounterService.Update(enc)
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error updating encounter.")
	}
	response := responsefromDto(p)
	return &response, nil
}
func (s Server) Delete(ctx context.Context, request *encounter.EncounterId) (*emptypb.Empty, error) {
	fmt.Println("HEY DELETE")
	s.encounterService.Delete(int64(request.Id))
	return &emptypb.Empty{}, nil
}
func (s Server) GetById(ctx context.Context, request *encounter.EncounterId) (*encounter.EncounterDto, error) {
	fmt.Println("HEY GET")
	p, ok := s.encounterService.Get(int64(request.Id))
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error getting encounter.")
	}
	response := responsefromDto(p)
	return &response, nil
}
func (s Server) CreateEncounterExecution(ctx context.Context, request *encounter.EncounterExecutionDto) (*encounter.EncounterExecutionDto, error) {
	fmt.Println("HEY CREATE EXECUTION")
	p, ok := s.encounterExecutionService.Create(executionDtoFromRequest(*request), request.TouristId)
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error getting encounter.")
	}
	response := executionResponsefromDto(p)
	return &response, nil
}
func (s Server) GetEncounterById(ctx context.Context, request *encounter.EncounterId) (*encounter.EncounterDto, error) {
	fmt.Println("HEY GET")
	p, ok := s.encounterService.Get(int64(request.Id))
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error getting encounter.")
	}
	response := responsefromDto(p)
	return &response, nil
}
func (s Server) UpdateEncounterExecution(ctx context.Context, request *encounter.EncounterExecutionDto) (*encounter.EncounterExecutionDto, error) {
	fmt.Println("HEY UPDATE EXECUTION")
	p, ok := s.encounterExecutionService.Update(executionDtoFromRequest(*request))
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error updating execution.")
	}
	response := executionResponsefromDto(p)
	return &response, nil
}
func (s Server) ActivateEncounterExecution(ctx context.Context, request *encounter.ActivateRequest) (*encounter.EncounterExecutionDto, error) {
	fmt.Println("HEY ACTIVATE EXECUTION")
	p, ok := s.encounterExecutionService.Activate(1, request.TouristLatitude, request.TouristLongitude, int64(request.Id))
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error activating execution.")
	}
	response := executionResponsefromDto(p)
	return &response, nil
}
func (s Server) CompleteExecution(ctx context.Context, request *encounter.ActivateRequest) (*encounter.EncounterExecutionDto, error) {
	fmt.Println("HEY COMPLETE EXECUTION")
	p, ok := s.encounterExecutionService.CompleteExecution(int64(request.Id), 1, request.TouristLatitude, request.TouristLongitude)
	if ok != nil {
		return nil, status.Error(codes.NotFound, "Error completing execution.")
	}
	response := executionResponsefromDto(p)
	return &response, nil
}
func (s Server) DeleteExecution(ctx context.Context, request *encounter.EncounterId) (*emptypb.Empty, error) {
	fmt.Println("HEY DELETE EXECUTION")
	s.encounterExecutionService.Delete(int64(request.Id), 1)
	return &emptypb.Empty{}, nil
}
func (s Server) GetAllExecutionsByTourist(ctx context.Context, request *encounter.EncounterId) (*encounter.PagedExecutions, error) {
	fmt.Println("HEY GET ALL EXECUTIONS BY TOURIST")
	p := s.encounterExecutionService.GetAllByTourist(int64(request.Id))
	result := make([]*encounter.EncounterExecutionDto, 0, len(p))
	for _, r := range p {
		dto := executionResponsefromDto(r)
		result = append(result, &dto)
	}
	count := len(result)
	response := encounter.PagedExecutions{
		Results:    result,
		TotalCount: int32(count),
	}
	return &response, nil
}
func (s Server) GetAllCompletedExecutionsByTourist(ctx context.Context, request *encounter.PagedRequestWithId) (*encounter.PagedExecutions, error) {
	fmt.Println("HEY GET ALL EXECUTIONS BY TOURIST")
	p := s.encounterExecutionService.GetAllByTourist(int64(request.Id))
	result := make([]*encounter.EncounterExecutionDto, 0, len(p))
	for _, r := range p {
		dto := executionResponsefromDto(r)
		result = append(result, &dto)
	}
	count := len(result)
	response := encounter.PagedExecutions{
		Results:    result,
		TotalCount: int32(count),
	}
	return &response, nil
}

func dtoFromRequest(req encounter.EncounterDto) dto.EncounterDto {
	return dto.EncounterDto{
		Id:                int64(req.Id),
		AuthorId:          req.AuthorId,
		Name:              req.Name,
		Description:       req.Description,
		Xp:                int(req.XP),
		Status:            req.Status,
		Type:              req.Type,
		Latitude:          req.Latitude,
		Longitude:         req.Longitude,
		LocationLongitude: req.LocationLongitude,
		LocationLatitude:  req.LocationLatitude,
		Image:             req.Image,
		Range:             req.Range,
		RequiredPeople:    int(req.RequiredPeople),
		ActiveTouristsIds: req.ActiveTouristsIds,
	}
}
func responsefromDto(p dto.EncounterDto) encounter.EncounterDto {
	return encounter.EncounterDto{ //here is the error
		Id:                int32(p.Id),
		AuthorId:          p.AuthorId,
		Name:              p.Name,
		Description:       p.Description,
		XP:                int32(p.Xp),
		Status:            p.Status,
		Type:              p.Type,
		Latitude:          p.Latitude,
		Longitude:         p.Longitude,
		LocationLongitude: p.LocationLongitude,
		LocationLatitude:  p.LocationLatitude,
		Image:             p.Image,
		Range:             p.Range,
		RequiredPeople:    int32(p.RequiredPeople),
		ActiveTouristsIds: p.ActiveTouristsIds,
	}
}

func executionDtoFromRequest(req encounter.EncounterExecutionDto) dto.EncounterExecutionDto {
	return dto.EncounterExecutionDto{
		Id:           int64(req.Id),
		TouristId:    req.TouristId,
		StartTime:    req.StartTime.AsTime(),
		EncounterId:  req.EncounterId,
		EncounterDto: dtoFromRequest(*req.EncounterDto),
		Status:       req.Status,
	}
}
func executionResponsefromDto(p dto.EncounterExecutionDto) encounter.EncounterExecutionDto {
	encounterDto := responsefromDto(p.EncounterDto)
	return encounter.EncounterExecutionDto{
		Id:           int64(p.Id),
		TouristId:    p.TouristId,
		StartTime:    timestamppb.New(p.StartTime),
		EncounterId:  p.EncounterId,
		EncounterDto: &encounterDto,
		Status:       p.Status,
	}
}

// func startServer(requestHandler *handler.EncounterRequestHandler, encounterHandler *handler.EncounterHandler, executionHandler *handler.EncounterExecutionHandler, touristEncounterHandler *handler.TouristEncounterHandler) {
// 	router := mux.NewRouter().StrictSlash(true)

// 	//*requests
// 	router.HandleFunc("/encounterRequests/getAll", requestHandler.GetAll).Methods("GET")
// 	router.HandleFunc("/encounterRequests/accept/{id}", requestHandler.AcceptRequest).Methods("PUT")
// 	router.HandleFunc("/encounterRequests/reject/{id}", requestHandler.RejectRequest).Methods("PUT")

// 	//*encounters
// 	router.HandleFunc("/encounter/getAll", encounterHandler.GetAll).Methods("GET")
// 	router.HandleFunc("/encounter/get/{id}", encounterHandler.GetById).Methods("GET")
// 	//! update checkpoint
// 	router.HandleFunc("/encounter/create/{checkpointId}/{isSecretPrerequisite}", encounterHandler.Create).Methods("POST")
// 	router.HandleFunc("/encounter/update", encounterHandler.Update).Methods("PUT")
// 	//! update checkpoint
// 	router.HandleFunc("/encounter/delete/{id}", encounterHandler.Delete).Methods("DELETE")

// 	//*executions
// 	router.HandleFunc("/execution/get/{id}", executionHandler.GetById).Methods("GET")
// 	router.HandleFunc("/execution/getAllByTourist/{id}", executionHandler.GetAllByTourist).Methods("GET")
// 	router.HandleFunc("/execution/getAllCompletedByTourist/{id}", executionHandler.GetAllCompletedByTourist).Methods("GET")
// 	//! need body (encounterIds)
// 	router.HandleFunc("/execution/getByTour/{touristLatitude}/{touristLongitude}/{touristId}", executionHandler.GetByTour).Methods("PUT")
// 	router.HandleFunc("/execution/checkPosition/{id}/{touristLatitude}/{touristLongitude}/{touristId}", executionHandler.CheckPosition).Methods("PUT")
// 	router.HandleFunc("/execution/checkPositionLocationEncounter/{id}/{touristLatitude}/{touristLongitude}/{touristId}", executionHandler.CheckPositionLocationEncounter).Methods("PUT")
// 	router.HandleFunc("/execution/getActiveByTour/{touristId}", executionHandler.GetActiveByTour).Methods("PUT")
// 	//! end of body required methods
// 	router.HandleFunc("/execution/activate/{id}/{touristId}/{touristLatitude}/{touristLongitude}", executionHandler.Activate).Methods("PUT")
// 	//! update tourists xp points
// 	router.HandleFunc("/execution/complete/{id}/{touristId}/{touristLatitude}/{touristLongitude}", executionHandler.CompleteExecution).Methods("PUT")
// 	router.HandleFunc("/execution/update", executionHandler.Update).Methods("PUT")
// 	router.HandleFunc("/execution/delete/{id}/{touristId}", executionHandler.Update).Methods("DELETE")

// 	//*tourist encounter
// 	router.HandleFunc("/touristEncounter/getAll", touristEncounterHandler.GetAll).Methods("GET")
// 	router.HandleFunc("/touristEncounter/get/{id}", touristEncounterHandler.GetById).Methods("GET")
// 	//! update checkpoint
// 	router.HandleFunc("/touristEncounter/create/{checkpointId}/{isSecretPrerequisite}", touristEncounterHandler.Create).Methods("POST")
// 	router.HandleFunc("/touristEncounter/update", touristEncounterHandler.Update).Methods("PUT")
// 	//! update checkpoint
// 	router.HandleFunc("/touristEncounter/delete/{id}", touristEncounterHandler.Delete).Methods("DELETE")

// 	println("Server listening on port 8090")
// 	log.Fatal(http.ListenAndServe(":8090", router))
// }
