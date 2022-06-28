package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

func TestRunnerRegister(t *testing.T) {
	// Arrange
	client := ATestClient(t, "job/register")
	// Act
	result, err := client.RunnerRegister()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, *ol.NewID("1234"), result.Id)
}

func TestRunnerGetPendingJobs(t *testing.T) {
	// Arrange
	client := ATestClient(t, "job/get_pending")
	token := ol.NewID("1234")
	// Act
	result, token, err := client.RunnerGetPendingJob("1234567890", &token)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "public.ecr.aws/opslevel/cli:v2022.02.25", result.Image)
	autopilot.Equals(t, "ls -al", result.Commands[1])
	autopilot.Equals(t, ol.NewID("12344321"), token)
}

func TestRunnerAppendJobLog(t *testing.T) {
	// Arrange
	client := ATestClient(t, "job/append_log")
	// Act
	err := client.RunnerAppendJobLog(ol.RunnerAppendJobLogInput{
		RunnerId:    ol.NewID("1234"),
		RunnerJobId: ol.NewID("5678"),
		SentAt:      ol.NewISO8601Date("2022-07-01T01:00:00.000Z"),
		Logs:        []string{"Log1", "Log2"},
	})
	// Assert
	autopilot.Ok(t, err)
}

func TestRunnerReportJobOutcome(t *testing.T) {
	// Arrange
	client := ATestClient(t, "job/report_outcome")
	// Act
	err := client.RunnerReportJobOutcome(ol.RunnerReportJobOutcomeInput{
		RunnerId:    ol.NewID("1234567890"),
		RunnerJobId: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvUnVubmVyczo6Sm9iUnVuLzE"),
		Outcome:     ol.RunnerJobOutcomeEnumExecutionTimeout,
	})
	// Assert
	autopilot.Ok(t, err)
}

func TestRunnerUnregister(t *testing.T) {
	// Arrange
	client := ATestClient(t, "job/unregister")
	//id := ol.NewID("1234")
	// Act
	err := client.RunnerUnregister(ol.NewID("1234"))
	// Assert
	autopilot.Ok(t, err)
}
