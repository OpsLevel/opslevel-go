package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2026"
	"github.com/rocktavious/autopilot/v2023"
)

func TestRunnerRegister(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerRegister($queue:String){runnerRegister(queue: $queue){runner{id,status},errors{message,path}}}`,
		`{"queue": null}`,
		`{"data": {"runnerRegister": { "runner": { "id": "1234", "status": "registered" }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "job/register", testRequest)
	// Act
	result, err := client.RunnerRegister()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("1234"), result.Id)
}

func TestRunnerRegisterWithQueue(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerRegister($queue:String){runnerRegister(queue: $queue){runner{id,status},errors{message,path}}}`,
		`{"queue": "my-queue"}`,
		`{"data": {"runnerRegister": { "runner": { "id": "1234", "status": "registered" }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "job/register_with_queue", testRequest)
	// Act
	result, err := client.RunnerRegister("my-queue")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("1234"), result.Id)
}

func TestRunnerGetScale(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query RunnerScale($currentReplicaCount:Int!$jobConcurrency:Int!$runnerId:ID!){account{runnerScale(runnerId: $runnerId, currentReplicaCount: $currentReplicaCount, jobConcurrency: $jobConcurrency){recommendedReplicaCount}}}`,
		`{"currentReplicaCount":2, "jobConcurrency":3, "runnerId":"1234567890" }`,
		`{"data": { "account": { "runnerScale": { "recommendedReplicaCount": 6 }}}}`,
	)

	client := BestTestClient(t, "job/scale", testRequest)
	// Act
	result, err := client.RunnerScale("1234567890", 2, 3)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 6, result.RecommendedReplicaCount)
}

func TestRunnerGetPendingJobs(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerGetPendingJob($id:ID!$token:ID){runnerGetPendingJob(runnerId: $id lastUpdateToken: $token){runnerJob{commands,id,image,outcome,resources{requests{cpu,memory,ephemeralStorage},limits{cpu,memory,ephemeralStorage}},status,variables{key,sensitive,value,scope},files{name,contents},initCommands,initImage},lastUpdateToken,errors{message,path}}}`,
		`{"id":"1234567890", "token":  "1234"}`,
		`{"data": {
      "runnerGetPendingJob": {
        "runnerJob": {
          {{ template "id1" }},
          "image": "public.ecr.aws/opslevel/cli:v2026.7.8",
          "outcome": "unstarted",
          "resources": null,
          "status": "running",
          "commands": [
            "pwd",
            "ls -al",
            "env | grep AWS"
          ],
          "variables": [
            {
              "key": "AWS_ACCESS_KEY",
              "value": "XXXXXXX",
              "scope": "main"
            },
            {
              "key": "REPO_CLONE_URL",
              "value": "https://token@example.com/repo.git",
              "sensitive": true,
              "scope": "init"
            }
          ],
          "initCommands": [
            "/opslevel/clone-repo ."
          ],
          "initImage": "public.ecr.aws/opslevel/cli:v2026.7.8"
        },
        "lastUpdateToken": "12344321",
        "errors": []
      }}}`,
	)

	client := BestTestClient(t, "job/get_pending", testRequest)
	// Act
	result, token, err := client.RunnerGetPendingJob("1234567890", "1234")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "public.ecr.aws/opslevel/cli:v2026.7.8", result.Image)
	autopilot.Equals(t, "ls -al", result.Commands[1])
	autopilot.Equals(t, ol.ID("12344321"), token)
	autopilot.Equals(t, []string{"/opslevel/clone-repo ."}, result.InitCommands)
	autopilot.Equals(t, "public.ecr.aws/opslevel/cli:v2026.7.8", result.InitImage)
	autopilot.Equals(t, ol.RunnerJobVariableScopeMain, result.Variables[0].Scope)
	autopilot.Equals(t, ol.RunnerJobVariableScopeInit, result.Variables[1].Scope)
	autopilot.Assert(t, result.Resources == nil, "expected Resources to be nil when server omits it")
}

func TestRunnerGetPendingJobsWithResourcesSpecs(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerGetPendingJob($id:ID!$token:ID){runnerGetPendingJob(runnerId: $id lastUpdateToken: $token){runnerJob{commands,id,image,outcome,resources{requests{cpu,memory,ephemeralStorage},limits{cpu,memory,ephemeralStorage}},status,variables{key,sensitive,value,scope},files{name,contents},initCommands,initImage},lastUpdateToken,errors{message,path}}}`,
		`{"id":"1234567890", "token":  "1234"}`,
		`{"data": {
      "runnerGetPendingJob": {
        "runnerJob": {
          {{ template "id1" }},
          "image": "public.ecr.aws/opslevel/cli:v2026.7.8",
          "outcome": "unstarted",
          "resources": {
            "requests": {
              "cpu": "500m",
              "memory": "1Gi",
              "ephemeralStorage": "64Mi"
            },
            "limits": {
              "cpu": "4000m",
              "memory": "4Gi",
              "ephemeralStorage": "26Gi"
            }
          },
          "status": "running",
          "commands": ["pwd"],
          "variables": [],
          "files": [],
          "initCommands": [],
          "initImage": ""
        },
        "lastUpdateToken": "12344321",
        "errors": []
      }}}`,
	)

	client := BestTestClient(t, "job/get_pending_with_resources", testRequest)
	// Act
	result, _, err := client.RunnerGetPendingJob("1234567890", "1234")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Assert(t, result.Resources != nil, "expected Resources to be non-nil")
	autopilot.Assert(t, result.Resources.Requests != nil, "expected Resources.Requests to be non-nil")
	autopilot.Equals(t, "500m", result.Resources.Requests.CPU)
	autopilot.Equals(t, "1Gi", result.Resources.Requests.Memory)
	autopilot.Equals(t, "64Mi", result.Resources.Requests.EphemeralStorage)
	autopilot.Assert(t, result.Resources.Limits != nil, "expected Resources.Limits to be non-nil")
	autopilot.Equals(t, "4000m", result.Resources.Limits.CPU)
	autopilot.Equals(t, "4Gi", result.Resources.Limits.Memory)
	autopilot.Equals(t, "26Gi", result.Resources.Limits.EphemeralStorage)
}

func TestRunnerAppendJobLog(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerAppendJobLog($input:RunnerAppendJobLogInput!){runnerAppendJobLog(input: $input){errors{message,path}}}`,
		`{"input":{ "logChunk":["Log1", "Log2"], "runnerId":"1234", "runnerJobId":"5678", "sentAt":"2022-07-01T01:00:00Z" }}`,
		`{"data": { "runnerAppendJobLog": { "errors": [] }}}`,
	)

	client := BestTestClient(t, "job/append_log", testRequest)
	// Act
	err := client.RunnerAppendJobLog(ol.RunnerAppendJobLogInput{
		RunnerId:    "1234",
		RunnerJobId: "5678",
		SentAt:      ol.NewISO8601Date("2022-07-01T01:00:00.000Z"),
		Logs:        []string{"Log1", "Log2"},
	})
	// Assert
	autopilot.Ok(t, err)
}

func TestRunnerReportJobOutcome(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerReportJobOutcome($input:RunnerReportJobOutcomeInput!){runnerReportJobOutcome(input: $input){errors{message,path}}}`,
		`{"input": { "runnerId":"1234567890", "runnerJobId": "{{ template "id1_string" }}", "outcome":"execution_timeout" }}`,
		`{"data": { "runnerReportJobOutcome": { "errors": [] }}}`,
	)

	client := BestTestClient(t, "job/report_outcome", testRequest)
	// Act
	err := client.RunnerReportJobOutcome(ol.RunnerReportJobOutcomeInput{
		RunnerId:    "1234567890",
		RunnerJobId: id1,
		Outcome:     ol.RunnerJobOutcomeEnumExecutionTimeout,
	})
	// Assert
	autopilot.Ok(t, err)
}

func TestRunnerUnregister(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerUnregister($runnerId:ID!){runnerUnregister(runnerId: $runnerId){errors{message,path}}}`,
		`{"runnerId": "1234" }`,
		`{"data": { "runnerUnregister": { "errors": [] }}}`,
	)

	client := BestTestClient(t, "job/unregister", testRequest)
	// Act
	err := client.RunnerUnregister("1234")
	// Assert
	autopilot.Ok(t, err)
}

func TestRunnerJobNumber(t *testing.T) {
	// Arrange
	job := ol.RunnerJob{
		Id: "Z2lkOi8vb3BzbGV2ZWwvUnVubmVyczo6Sm9iUnVuLzIyNQ",
	}
	// Act
	jobNumber := job.Number()
	// Assert
	autopilot.Equals(t, "225", jobNumber)
}

func TestRunnerJobNumberFailure(t *testing.T) {
	// Arrange
	job := ol.RunnerJob{
		Id: "Z2lkOi8vb3BzbGV2ZWwvUnVubmVyczo6Sm9iU",
	}
	// Act
	jobNumber := job.Number()
	// Assert
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvUnVubmVyczo6Sm9iU", jobNumber)
}
