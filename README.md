# Codeflix Video Encoder microsservice

## How to run it

1. Run `make up` command to run the containers. It will create the containers.
2. Run `make execute_rabbit` command to enter in the RabbitMQ container and run `cd /usr/src && chmod +x setup.sh && ./setup.sh` to set up the resources  (`videos-result` and `videos-failed` queues) used by the Go code.

> [!NOTE] The setup.sh file:
> - Creates the exchange `dlx` of type `fanout`
> - Creates the dead-letter queue (`videos-failed`) binded with `dlx` 
> - Creates the `videos-result` queue binded with `amq.direct` and routing key `jobs` 
  
3. On another terminal session, run `make execute_app` command to enter in the app container (if you want to visualize the output) and run the `server.go` file with `go run framework/cmd/server/server.go`
4. Send a message in the `videos` queue:
```json
{
    "resource_id": "id-client-1",
    "file_path": "test.mp4"
}
```