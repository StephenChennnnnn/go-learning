package server

import (
	"crypto/tls"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-learning/grpc/grpc-hello-world/pkg/ui/data/swagger"
	"go-learning/grpc/grpc-hello-world/pkg/util"
	pb "go-learning/grpc/grpc-hello-world/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
	"path"
	"strings"
)

var (
	ServerPort  string
	CertName    string
	CertPemPath string
	CertKeyPath string
	SwaggerDir  string
	EndPoint    string
)

func Serve() (err error) {
	log.Println(ServerPort)
	log.Println(CertName)
	log.Println(CertPemPath)
	log.Println(CertKeyPath)

	EndPoint = ":" + ServerPort
	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Println("TCP Listen err: %v\n", err)
	}

	tlsConfig := util.GetTLSConfig(CertPemPath, CertKeyPath)
	srv := createInternalServer(conn, tlsConfig)

	log.Println("gPRC and https listen on: %s\n", ServerPort)

	if err = srv.Serve(tls.NewListener(conn, tlsConfig)); err != nil {
		log.Println("ListenAndServe: %v\n", err)
	}

	return err
}

func createInternalServer(conn net.Listener, tlsConfig *tls.Config) *http.Server {
	var opts []grpc.ServerOption

	// grpc server
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Println("Failed to create server TLS credentials %v", err)
	}

	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)

	// register grpc pb
	pb.RegisterHelloWorldServer(grpcServer, NewHelloService())

	// gw server
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Println("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		log.Println("Failed to register gw server: %v\n", err)
	}

	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}

}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(SwaggerDir, p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
