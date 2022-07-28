# A simple lambda to demonstrate the error `Error: fork/exec /var/task/main: no such file or directory","errorType` when the lambda is invoked in Localstack

## Requirements

- go (go1.16.3)
- docker (v20.10.14)
- localstack (0.14.3.2)
- awscli (1.18.217)

# Steps to reproduce

1. Clone this repo
2. `cd sample-lambda/src`
3. Start localstack: `docker compose up`
4. Deploy the lambda: `make deploy`
5. Test the lambda: `make test`. An error like this should show: 
  ```
  {"errorMessage":"RequestId: d46ef559-68f4-1719-76e0-25b0cd4d72a6 Error: fork/exec /var/task/main: no such file or directory","errorType":"exitError"}
  {
      "StatusCode": 200,
      "LogResult": "",
      "ExecutedVersion": "$LATEST"
  }
   ```
