package opslevel

import (
	"encoding/base64"
	"github.com/relvacode/iso8601"
	"github.com/shurcooL/graphql"
	"strings"
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
	RunnerJobOutcomeEnumPodTimeout       RunnerJobOutcomeEnum = "pod_timeout"       // A pod could not be scheduled for the job in time.
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
		string(RunnerJobOutcomeEnumPodTimeout),
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

// RunnerStatusTypeEnum represents The status of an OpsLevel runner.
type RunnerStatusTypeEnum string

const (
	RunnerStatusTypeEnumInactive   RunnerJobStatusEnum = "inactive"   // The runner will not actively take jobs.
	RunnerStatusTypeEnumRegistered RunnerJobStatusEnum = "registered" // The runner will process jobs.
)

// All RunnerStatusTypeEnum as []string
func AllRunnerStatusTypeEnum() []string {
	return []string{
		string(RunnerStatusTypeEnumInactive),
		string(RunnerStatusTypeEnumRegistered),
	}
}

type Runner struct {
	Id     graphql.ID           `json:"id"`
	Status RunnerStatusTypeEnum `json:"status"`
}

type RunnerJobVariable struct {
	Key       string `json:"key"`
	Sensitive bool   `json:"sensitive"`
	Value     string `json:"value"`
}

type RunnerJobFile struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
}

type RunnerJob struct {
	Commands  []string             `json:"commands"`
	Id        graphql.ID           `json:"id"`
	Image     string               `json:"image"`
	Outcome   RunnerJobOutcomeEnum `json:"outcome"`
	Status    RunnerJobStatusEnum  `json:"status"`
	Variables []RunnerJobVariable  `json:"variables"`
	Files     []RunnerJobFile      `json:"files"`
}

func (j *RunnerJob) Number() string {
	id := j.Id.(string)
	decoded, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return id
	}
	return strings.ReplaceAll(string(decoded), "gid://opslevel/Runners::JobRun/", "")
}

type RunnerAppendJobLogInput struct {
	RunnerId    graphql.ID   `json:"runnerId"`
	RunnerJobId graphql.ID   `json:"runnerJobId"`
	SentAt      iso8601.Time `json:"sentAt"`
	Logs        []string     `json:"logChunk"`
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

type RunnerScale struct {
	RecommendedReplicaCount int `json:"recommendedReplicaCount"`
}

func (c *Client) RunnerRegister() (*Runner, error) {
	var m struct {
		Payload struct {
			Runner Runner
			Errors []OpsLevelErrors
		} `graphql:"runnerRegister"`
	}
	v := PayloadVariables{}
	if err := c.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Runner, FormatErrors(m.Payload.Errors)
}

func (c *Client) RunnerGetPendingJob(runnerId graphql.ID, lastUpdateToken graphql.ID) (*RunnerJob, *graphql.ID, error) {
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
	if err := c.Mutate(&m, v); err != nil {
		return nil, nil, err
	}
	return &m.Payload.RunnerJob, &m.Payload.LastUpdateToken, FormatErrors(m.Payload.Errors)
}

func (c *Client) GetRunnerScale(runnerId graphql.ID, currentReplicaCount int32, jobConcurrency int) (*RunnerScale, error) {
	var q struct {
		Account struct {
			RunnerScale RunnerScale `graphql:"runnerScale(runnerId: $runnerId, currentReplicaCount: $currentReplicaCount, jobConcurrency: $jobConcurrency)"`
		}
	}
	v := PayloadVariables{
		"runnerId":            runnerId,
		"currentReplicaCount": currentReplicaCount,
		"jobConcurrency":      jobConcurrency,
	}
	if err := c.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.RunnerScale, nil
}

func (c *Client) RunnerAppendJobLog(input RunnerAppendJobLogInput) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"runnerAppendJobLog(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := c.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

func (c *Client) RunnerReportJobOutcome(input RunnerReportJobOutcomeInput) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"runnerReportJobOutcome(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := c.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

func (c *Client) RunnerUnregister(runnerId *graphql.ID) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"runnerUnregister(runnerId: $runnerId)"`
	}
	v := PayloadVariables{
		"runnerId": *runnerId,
	}
	if err := c.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}
