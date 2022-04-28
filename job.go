package opslevel

import (
	"github.com/shurcooL/graphql"
)

// RunnerJobOutcomeEnum represents the runner job outcome.
type RunnerJobOutcomeEnum string

const (
	RunnerJobOutcomeEnumUnstarted        RunnerJobOutcomeEnum = "unstarted"         // translation missing: en.graphql.types.runner_job_outcome_enum.unstarted.
	RunnerJobOutcomeEnumCanceled         RunnerJobOutcomeEnum = "canceled"          // Job was canceled.
	RunnerJobOutcomeEnumFailed           RunnerJobOutcomeEnum = "failed"            // Job failed during execution.
	RunnerJobOutcomeEnumSuccess          RunnerJobOutcomeEnum = "success"           // Job succeded the execution.
	RunnerJobOutcomeEnumQueueTimeout     RunnerJobOutcomeEnum = "queue_timeout"     // Job was not assigned to a runner for too long.
	RunnerJobOutcomeEnumExecutionTimeout RunnerJobOutcomeEnum = "execution_timeout" // Job run took too long to complete, and was marked as failed.
)

// All RunnerJobOutcomeEnum as []string
func AllRunnerJobOutcomeEnum() []string {
	return []string{
		string(RunnerJobOutcomeEnumUnstarted),
		string(RunnerJobOutcomeEnumCanceled),
		string(RunnerJobOutcomeEnumFailed),
		string(RunnerJobOutcomeEnumSuccess),
		string(RunnerJobOutcomeEnumQueueTimeout),
		string(RunnerJobOutcomeEnumExecutionTimeout),
	}
}

// RunnerJobStatusEnum represents the runner job status.
type RunnerJobStatusEnum string

const (
	RunnerJobStatusEnumCreated  RunnerJobStatusEnum = "created"  // A created runner job, but not yet ready to be run.
	RunnerJobStatusEnumPending  RunnerJobStatusEnum = "pending"  // A runner job ready to be run.
	RunnerJobStatusEnumRunning  RunnerJobStatusEnum = "running"  // A runner job being run by a runner.
	RunnerJobStatusEnumComplete RunnerJobStatusEnum = "complete" // A finished runner job.
)

// All RunnerJobStatusEnum as []string
func AllRunnerJobStatusEnum() []string {
	return []string{
		string(RunnerJobStatusEnumCreated),
		string(RunnerJobStatusEnumPending),
		string(RunnerJobStatusEnumRunning),
		string(RunnerJobStatusEnumComplete),
	}
}

type RunnerJobVariable struct {
	Key       string `json:"key"`
	Sensitive bool   `json:"sensitive"`
	Value     string `json:"value"`
}

type RunnerJob struct {
	Commands  []string             `json:"commands"`
	Id        graphql.ID           `json:"id"`
	Image     string               `json:"image"`
	Outcome   RunnerJobOutcomeEnum `json:"outcome"`
	Status    RunnerJobStatusEnum  `json:"status"`
	Variables []RunnerJobVariable  `json:"variables"`
}

type RunnerJobOutcomeVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RunnerReportJobOutcomeInput struct {
	RunnerId         graphql.ID                 `json:"runnerId"`
	RunnerJobId      graphql.ID                 `json:"runnerJobId"`
	Outcome          RunnerJobOutcomeEnum       `json:"outcome"`
	OutcomeVariables []RunnerJobOutcomeVariable `json:"outcomeVariables,omitempty"`
}

func (s *Client) ReportJobOutcome(input RunnerReportJobOutcomeInput) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"runnerReportJobOutcome(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := s.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

func (s *Client) GetPendingJob(runnerId graphql.ID, lastUpdateToken graphql.ID) (*RunnerJob, *graphql.ID, error) {
	var m struct {
		Payload struct {
			RunnerJob       RunnerJob
			LastUpdateToken graphql.ID
			Errors          []OpsLevelErrors
		} `graphql:"runnerGetPendingJob(runnerId: $id lastUpdateToken: $token)"`
	}
	v := PayloadVariables{
		"id":    runnerId,
		"token": lastUpdateToken,
	}
	if err := s.Mutate(&m, v); err != nil {
		return nil, nil, err
	}
	return &m.Payload.RunnerJob, &m.Payload.LastUpdateToken, FormatErrors(m.Payload.Errors)
}
