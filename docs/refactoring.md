# Refactoring documentation

During this microsservice implementation, some architectural or desing issues were encountered. 
This documentation was created to list those issues for further revision and adjustment.

## Backlog

### Repositories
- [ ] Remove database configuration from entity definition on `Video` and `Job`
- [ ] Fix ORM error when handling with column referencing objects. Ex.: Object `Job` has a attribute `Video` that is returned by the ORM (Maybe it is a `sqlite3.orm` limitation)

### Services

#### Video service
- [ ] Create an abstraction (cloud repository) to download the video and upload its encoded parts
- [ ] Refactor the tests to use a fake repository rather than the real cloud repository
- [ ] replace the object names to a more expressive one (services, functions, channels, etc.)
- [ ] Handle error when there is no file to be processed by the upload_manager
- [ ] Change the manager to pass the destination path to the UploadObject function. Its responsibility is only upload an object and not know about how to convert a local source file path into a "some place in the cloud" file path
- [ ] Remove the `getClientUpload` from the service. It must be abstracted by a singleton pattern in the cloud repository
- [ ] Create a single function that abstract all steps related to the upload process and hide the other functions from external packages
- [ ] Create an Enum of result status, so the goroutines don't have to rely on simple string values
- [ ] Manipulate paths using built-in functions rather than simple strings


#### Job service
- [ ] Remove upload process from the job service. This is not its responsibility. The job service will only instanciate the cloud repository
- [ ] The message body from the outside world is coupled with the domain where the `resource_id` and `file_path` of the incomming message are declared in the Video struct. Change this by adding a DTO-ish or tadutor to address this issue.

#### Job Manager
- [ ] Rename the `checkParseErrors` function to a more expressive name. Investigate with it does not receive a channel to notify about errors rather than the ` notifySuccess`  does.
- [ ] Investigate the need to add more mutexes in the `Marchal/Unmarsal` calls

### Infrastructure

- [ ] Create an abstraction layer for how the microservice is integrated with the external services. It cannot refer the RabbitMQ directy
- [ ] Hide database secrets: remove them from the `.env` file

#### RabbitMQ

- [ ] Defining a way to create the necessary infrastructure for the RabbitMQ. Namely, the `dlx` exchange to receive the dead-letter queue and the `videos-result` (binded with the `amq.direct` exchange and route key `jobs`), `videos-failed` (binded with the `dlx` exchange)

## Done

- [x] Create a `resources` folder to be used by the unit tests
- [x] In the Finish function, replace the `os.Getenv` calls by a variable
- [x] Find a standard way to comment on the `queue.go` giving the credits to the author (wesley willians)