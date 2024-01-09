# Codeflix Video Encoder microsservice

## How to run it

1. Run `make up` command to run the containers
2. Run `make execute` command to enter in the app container (if you want to visualize the output)
3. Create the dead-letter queue (`videos-failed`) and the `videos-result` queues on RabbitMQ  
4. Send a message in the `videos` queue:
```json
{
    "resource_id": "id-client-1",
    "file_path": "test.mp4"
}
```