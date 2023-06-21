package main

const (
	webPort  = "8081"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

type Config struct{}

func main() {

	app := Config{}

	// start web server
	//log.Println("Starting service on port", webPort)
	app.gRPCListen()
	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf(":%s", webPort),
	// 	Handler: app.routes(),
	// }

	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Panic()
	// }

}
