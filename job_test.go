package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2026"
	"github.com/rocktavious/autopilot/v2023"
)

func TestRunnerRegister(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerRegister($subscribedTemplateIds:[Int!]){runnerRegister(subscribedTemplateIds: $subscribedTemplateIds){runner{id,status},errors{message,path}}}`,
		`{"subscribedTemplateIds": null}`,
		`{"data": {"runnerRegister": { "runner": { "id": "1234", "status": "registered" }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "job/register", testRequest)
	// Act
	result, err := client.RunnerRegister()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("1234"), result.Id)
}

func TestRunnerRegisterWithTemplateIds(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RunnerRegister($subscribedTemplateIds:[Int!]){runnerRegister(subscribedTemplateIds: $subscribedTemplateIds){runner{id,status},errors{message,path}}}`,
		`{"subscribedTemplateIds": [1, 2, 3]}`,
		`{"data": {"runnerRegister": { "runner": { "id": "1234", "status": "registered" }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "job/register_with_templates", testRequest)
	// Act
	result, err := client.RunnerRegister(1, 2, 3)
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
		`mutation RunnerGetPendingJob($id:ID!$token:ID){runnerGetPendingJob(runnerId: $id lastUpdateToken: $token){runnerJob{commands,id,image,outcome,status,variables{key,sensitive,value},files{name,contents}},lastUpdateToken,errors{message,path}}}`,
		`{"id":"1234567890", "token":  "1234"}`,
		`{"data": {
      "runnerGetPendingJob": {
        "runnerJob": {
          {{ template "id1" }},
          "image": "public.ecr.aws/opslevel/cli:v2022.02.25",
          "outcome": "unstarted",
          "status": "running",
          "commands": [
            "pwd",
            "ls -al",
            "env | grep AWS"
          ],
          "variables": [
            {
              "key": "AWS_ACCESS_KEY",
              "value": "XXXXXXX"
            }
          ]
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
	autopilot.Equals(t, "public.ecr.aws/opslevel/cli:v2022.02.25", result.Image)
	autopilot.Equals(t, "ls -al", result.Commands[1])
	autopilot.Equals(t, ol.ID("12344321"), token)
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
