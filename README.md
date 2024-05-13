# Codeflix Video Encoder microsservice

## How to run it

1. Run `make up` command to run the containers
2. Run `make execute` command to enter in the app container (if you want to visualize the output)
3. Prepare the infrastructure for RabbitMQ:
   1. Create the exchange `dlx` of type `fanout`
   2. Create the dead-letter queue (`videos-failed`) binded with `dlx` 
   3. Create the `videos-result` queue binded with `amq.direct` and routing key `jobs` 
4. Run the server.go file `go run framework/cmd/server/server.go`
5. Send a message in the `videos` queue:
```json
{
    "resource_id": "id-client-1",
    "file_path": "test.mp4"
}
```