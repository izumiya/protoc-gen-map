// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testdata/gentest/unary_type_test.proto

package gentest

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"

	//protoc-gen-map packages
	bytes "bytes"
	context "context"
	sql "database/sql"
	mapper "github.com/jackskj/protoc-gen-map/mapper"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	log "log"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Code generated by protoc-gen-map. DO NOT EDIT.
// To Use:
// 1. Instantiate MapperServers with sql.DB connection
// 2. Register MapperServer as the gRPC service server
// 3. Begin serving

type ExecTypeServiceMapServer struct {
	DB              *sql.DB
	CreateMapper    *mapper.Mapper
	DeleteMapper    *mapper.Mapper
	ExecFiveMapper  *mapper.Mapper
	ExecFourMapper  *mapper.Mapper
	ExecOneMapper   *mapper.Mapper
	ExecThreeMapper *mapper.Mapper
	ExecTwoMapper   *mapper.Mapper
	InSeRtMapper    *mapper.Mapper
	UpdateMapper    *mapper.Mapper

	mapperGenMux sync.Mutex
}

func (m *ExecTypeServiceMapServer) ExecOne(ctx context.Context, r *EmptyRequest) (*EmptyResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "ExecOne", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := EmptyResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) ExecTwo(ctx context.Context, r *EmptyRequest) (*EMpTyResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "ExecTwo", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := EMpTyResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) ExecThree(ctx context.Context, r *EmptyRequest) (*NilResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "ExecThree", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := NilResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) ExecFour(ctx context.Context, r *EmptyRequest) (*NullResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "ExecFour", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := NullResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) ExecFive(ctx context.Context, r *EmptyRequest) (*NuLlResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "ExecFive", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := NuLlResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) InSeRt(ctx context.Context, r *EmptyRequest) (*EmptyResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "InSeRt", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := EmptyResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) Delete(ctx context.Context, r *EmptyRequest) (*SampleResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "Delete", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := SampleResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) Update(ctx context.Context, r *EmptyRequest) (*SampleResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "Update", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := SampleResponse{}
	return &resp, nil

}

func (m *ExecTypeServiceMapServer) Create(ctx context.Context, r *EmptyRequest) (*SampleResponse, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "Create", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n EmptyRequest request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := SampleResponse{}
	return &resp, nil

}