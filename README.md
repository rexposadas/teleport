# teleport
Teleport Challenge


# Running the app

In one terminal, run the server

```
$ go run cmd/teleportd/main.go
```


In another terminal, make a call to sleep for 50 seconds. The 50 seconds is so you have 
enough time to get the status of the running job. 
```
$ cd cmd/teleport
$ go build ./... && ./teleport start /bin/sleep 50  
job id: 32bc2e95-c420-4ea9-8e2f-c9056f7f0ed4 
```

Get the status of the running job.  

```bash
$ go build ./... && ./teleport status 32bc2e95-c420-4ea9-8e2f-c9056f7f0ed4
job status: STATUS_RUNNING     
```

Once the job has finished, call status again and see the results. 

```bash
$ go build ./... && ./teleport status 32bc2e95-c420-4ea9-8e2f-c9056f7f0ed4
job status: STATUS_EXITED    
```

