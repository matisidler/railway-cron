// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package railway

import (
	"github.com/Khan/genqlient/graphql"
)

type DeploymentListInput struct {
	EnvironmentId  string                 `json:"environmentId"`
	IncludeDeleted bool                   `json:"includeDeleted"`
	ProjectId      string                 `json:"projectId"`
	ServiceId      string                 `json:"serviceId"`
	Status         *DeploymentStatusInput `json:"status,omitempty"`
}

// GetEnvironmentId returns DeploymentListInput.EnvironmentId, and is useful for accessing the field via an interface.
func (v *DeploymentListInput) GetEnvironmentId() string { return v.EnvironmentId }

// GetIncludeDeleted returns DeploymentListInput.IncludeDeleted, and is useful for accessing the field via an interface.
func (v *DeploymentListInput) GetIncludeDeleted() bool { return v.IncludeDeleted }

// GetProjectId returns DeploymentListInput.ProjectId, and is useful for accessing the field via an interface.
func (v *DeploymentListInput) GetProjectId() string { return v.ProjectId }

// GetServiceId returns DeploymentListInput.ServiceId, and is useful for accessing the field via an interface.
func (v *DeploymentListInput) GetServiceId() string { return v.ServiceId }

// GetStatus returns DeploymentListInput.Status, and is useful for accessing the field via an interface.
func (v *DeploymentListInput) GetStatus() *DeploymentStatusInput { return v.Status }

// DeploymentRedeployDeploymentRedeployDeployment includes the requested fields of the GraphQL type Deployment.
type DeploymentRedeployDeploymentRedeployDeployment struct {
	Id string `json:"id"`
}

// GetId returns DeploymentRedeployDeploymentRedeployDeployment.Id, and is useful for accessing the field via an interface.
func (v *DeploymentRedeployDeploymentRedeployDeployment) GetId() string { return v.Id }

// DeploymentRedeployResponse is returned by DeploymentRedeploy on success.
type DeploymentRedeployResponse struct {
	// Redeploys a deployment.
	DeploymentRedeploy *DeploymentRedeployDeploymentRedeployDeployment `json:"deploymentRedeploy"`
}

// GetDeploymentRedeploy returns DeploymentRedeployResponse.DeploymentRedeploy, and is useful for accessing the field via an interface.
func (v *DeploymentRedeployResponse) GetDeploymentRedeploy() *DeploymentRedeployDeploymentRedeployDeployment {
	return v.DeploymentRedeploy
}

// DeploymentRestartResponse is returned by DeploymentRestart on success.
type DeploymentRestartResponse struct {
	// Restarts a deployment.
	DeploymentRestart bool `json:"deploymentRestart"`
}

// GetDeploymentRestart returns DeploymentRestartResponse.DeploymentRestart, and is useful for accessing the field via an interface.
func (v *DeploymentRestartResponse) GetDeploymentRestart() bool { return v.DeploymentRestart }

type DeploymentStatus string

const (
	DeploymentStatusBuilding     DeploymentStatus = "BUILDING"
	DeploymentStatusCrashed      DeploymentStatus = "CRASHED"
	DeploymentStatusDeploying    DeploymentStatus = "DEPLOYING"
	DeploymentStatusFailed       DeploymentStatus = "FAILED"
	DeploymentStatusInitializing DeploymentStatus = "INITIALIZING"
	DeploymentStatusQueued       DeploymentStatus = "QUEUED"
	DeploymentStatusRemoved      DeploymentStatus = "REMOVED"
	DeploymentStatusRemoving     DeploymentStatus = "REMOVING"
	DeploymentStatusSkipped      DeploymentStatus = "SKIPPED"
	DeploymentStatusSuccess      DeploymentStatus = "SUCCESS"
	DeploymentStatusWaiting      DeploymentStatus = "WAITING"
)

type DeploymentStatusInput struct {
	In    []DeploymentStatus `json:"in"`
	NotIn []DeploymentStatus `json:"notIn"`
}

// GetIn returns DeploymentStatusInput.In, and is useful for accessing the field via an interface.
func (v *DeploymentStatusInput) GetIn() []DeploymentStatus { return v.In }

// GetNotIn returns DeploymentStatusInput.NotIn, and is useful for accessing the field via an interface.
func (v *DeploymentStatusInput) GetNotIn() []DeploymentStatus { return v.NotIn }

// DeploymentsDeploymentsQueryDeploymentsConnection includes the requested fields of the GraphQL type QueryDeploymentsConnection.
type DeploymentsDeploymentsQueryDeploymentsConnection struct {
	Edges []*DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdge `json:"edges"`
}

// GetEdges returns DeploymentsDeploymentsQueryDeploymentsConnection.Edges, and is useful for accessing the field via an interface.
func (v *DeploymentsDeploymentsQueryDeploymentsConnection) GetEdges() []*DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdge {
	return v.Edges
}

// DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdge includes the requested fields of the GraphQL type QueryDeploymentsConnectionEdge.
type DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdge struct {
	Node *DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment `json:"node"`
}

// GetNode returns DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdge.Node, and is useful for accessing the field via an interface.
func (v *DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdge) GetNode() *DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment {
	return v.Node
}

// DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment includes the requested fields of the GraphQL type Deployment.
type DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment struct {
	Id     string           `json:"id"`
	Status DeploymentStatus `json:"status"`
}

// GetId returns DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment.Id, and is useful for accessing the field via an interface.
func (v *DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment) GetId() string {
	return v.Id
}

// GetStatus returns DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment.Status, and is useful for accessing the field via an interface.
func (v *DeploymentsDeploymentsQueryDeploymentsConnectionEdgesQueryDeploymentsConnectionEdgeNodeDeployment) GetStatus() DeploymentStatus {
	return v.Status
}

// DeploymentsResponse is returned by Deployments on success.
type DeploymentsResponse struct {
	// Get all deployments
	Deployments *DeploymentsDeploymentsQueryDeploymentsConnection `json:"deployments"`
}

// GetDeployments returns DeploymentsResponse.Deployments, and is useful for accessing the field via an interface.
func (v *DeploymentsResponse) GetDeployments() *DeploymentsDeploymentsQueryDeploymentsConnection {
	return v.Deployments
}

// MeMeUser includes the requested fields of the GraphQL type User.
type MeMeUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetName returns MeMeUser.Name, and is useful for accessing the field via an interface.
func (v *MeMeUser) GetName() string { return v.Name }

// GetEmail returns MeMeUser.Email, and is useful for accessing the field via an interface.
func (v *MeMeUser) GetEmail() string { return v.Email }

// MeResponse is returned by Me on success.
type MeResponse struct {
	// Gets the authenticated user.
	Me *MeMeUser `json:"me"`
}

// GetMe returns MeResponse.Me, and is useful for accessing the field via an interface.
func (v *MeResponse) GetMe() *MeMeUser { return v.Me }

// ProjectProject includes the requested fields of the GraphQL type Project.
type ProjectProject struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

// GetName returns ProjectProject.Name, and is useful for accessing the field via an interface.
func (v *ProjectProject) GetName() string { return v.Name }

// GetId returns ProjectProject.Id, and is useful for accessing the field via an interface.
func (v *ProjectProject) GetId() string { return v.Id }

// ProjectResponse is returned by Project on success.
type ProjectResponse struct {
	// Get a project by ID
	Project *ProjectProject `json:"project"`
}

// GetProject returns ProjectResponse.Project, and is useful for accessing the field via an interface.
func (v *ProjectResponse) GetProject() *ProjectProject { return v.Project }

// ServiceResponse is returned by Service on success.
type ServiceResponse struct {
	// Get a service by ID
	Service *ServiceService `json:"service"`
}

// GetService returns ServiceResponse.Service, and is useful for accessing the field via an interface.
func (v *ServiceResponse) GetService() *ServiceService { return v.Service }

// ServiceService includes the requested fields of the GraphQL type Service.
type ServiceService struct {
	Name string `json:"name"`
}

// GetName returns ServiceService.Name, and is useful for accessing the field via an interface.
func (v *ServiceService) GetName() string { return v.Name }

// __DeploymentRedeployInput is used internally by genqlient
type __DeploymentRedeployInput struct {
	Id string `json:"id"`
}

// GetId returns __DeploymentRedeployInput.Id, and is useful for accessing the field via an interface.
func (v *__DeploymentRedeployInput) GetId() string { return v.Id }

// __DeploymentRestartInput is used internally by genqlient
type __DeploymentRestartInput struct {
	Id string `json:"id"`
}

// GetId returns __DeploymentRestartInput.Id, and is useful for accessing the field via an interface.
func (v *__DeploymentRestartInput) GetId() string { return v.Id }

// __DeploymentsInput is used internally by genqlient
type __DeploymentsInput struct {
	First int                  `json:"first"`
	Input *DeploymentListInput `json:"input,omitempty"`
}

// GetFirst returns __DeploymentsInput.First, and is useful for accessing the field via an interface.
func (v *__DeploymentsInput) GetFirst() int { return v.First }

// GetInput returns __DeploymentsInput.Input, and is useful for accessing the field via an interface.
func (v *__DeploymentsInput) GetInput() *DeploymentListInput { return v.Input }

// __ProjectInput is used internally by genqlient
type __ProjectInput struct {
	Id string `json:"id"`
}

// GetId returns __ProjectInput.Id, and is useful for accessing the field via an interface.
func (v *__ProjectInput) GetId() string { return v.Id }

// __ServiceInput is used internally by genqlient
type __ServiceInput struct {
	Id string `json:"id"`
}

// GetId returns __ServiceInput.Id, and is useful for accessing the field via an interface.
func (v *__ServiceInput) GetId() string { return v.Id }

// The query or mutation executed by DeploymentRedeploy.
const DeploymentRedeploy_Operation = `
mutation DeploymentRedeploy ($id: String!) {
	deploymentRedeploy(id: $id) {
		id
	}
}
`

func DeploymentRedeploy(
	client graphql.Client,
	id string,
) (*DeploymentRedeployResponse, error) {
	req := &graphql.Request{
		OpName: "DeploymentRedeploy",
		Query:  DeploymentRedeploy_Operation,
		Variables: &__DeploymentRedeployInput{
			Id: id,
		},
	}
	var err error

	var data DeploymentRedeployResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by DeploymentRestart.
const DeploymentRestart_Operation = `
mutation DeploymentRestart ($id: String!) {
	deploymentRestart(id: $id)
}
`

func DeploymentRestart(
	client graphql.Client,
	id string,
) (*DeploymentRestartResponse, error) {
	req := &graphql.Request{
		OpName: "DeploymentRestart",
		Query:  DeploymentRestart_Operation,
		Variables: &__DeploymentRestartInput{
			Id: id,
		},
	}
	var err error

	var data DeploymentRestartResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by Deployments.
const Deployments_Operation = `
query Deployments ($first: Int, $input: DeploymentListInput!) {
	deployments(input: $input, first: $first) {
		edges {
			node {
				id
				status
			}
		}
	}
}
`

func Deployments(
	client graphql.Client,
	first int,
	input *DeploymentListInput,
) (*DeploymentsResponse, error) {
	req := &graphql.Request{
		OpName: "Deployments",
		Query:  Deployments_Operation,
		Variables: &__DeploymentsInput{
			First: first,
			Input: input,
		},
	}
	var err error

	var data DeploymentsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by Me.
const Me_Operation = `
query Me {
	me {
		name
		email
	}
}
`

func Me(
	client graphql.Client,
) (*MeResponse, error) {
	req := &graphql.Request{
		OpName: "Me",
		Query:  Me_Operation,
	}
	var err error

	var data MeResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by Project.
const Project_Operation = `
query Project ($id: String!) {
	project(id: $id) {
		name
		id
	}
}
`

func Project(
	client graphql.Client,
	id string,
) (*ProjectResponse, error) {
	req := &graphql.Request{
		OpName: "Project",
		Query:  Project_Operation,
		Variables: &__ProjectInput{
			Id: id,
		},
	}
	var err error

	var data ProjectResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by Service.
const Service_Operation = `
query Service ($id: String!) {
	service(id: $id) {
		name
	}
}
`

func Service(
	client graphql.Client,
	id string,
) (*ServiceResponse, error) {
	req := &graphql.Request{
		OpName: "Service",
		Query:  Service_Operation,
		Variables: &__ServiceInput{
			Id: id,
		},
	}
	var err error

	var data ServiceResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}
