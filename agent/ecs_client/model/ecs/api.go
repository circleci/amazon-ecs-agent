// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package ecs

import (
	"sync"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

var oprw sync.Mutex

// CreateClusterRequest generates a request for the CreateCluster operation.
func (c *ECS) CreateClusterRequest(input *CreateClusterInput) (req *aws.Request, output *CreateClusterOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateCluster == nil {
		opCreateCluster = &aws.Operation{
			Name:       "CreateCluster",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opCreateCluster, input, output)
	output = &CreateClusterOutput{}
	req.Data = output
	return
}

// Creates a new Amazon ECS cluster. By default, your account will receive a
// default cluster when you launch your first container instance. However, you
// can create your own cluster with a unique name with the CreateCluster action.
func (c *ECS) CreateCluster(input *CreateClusterInput) (output *CreateClusterOutput, err error) {
	req, out := c.CreateClusterRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateCluster *aws.Operation

// CreateServiceRequest generates a request for the CreateService operation.
func (c *ECS) CreateServiceRequest(input *CreateServiceInput) (req *aws.Request, output *CreateServiceOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateService == nil {
		opCreateService = &aws.Operation{
			Name:       "CreateService",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opCreateService, input, output)
	output = &CreateServiceOutput{}
	req.Data = output
	return
}

// Runs and maintains a desired number of tasks from a specified task definition.
// If the number of tasks running in a service drops below desiredCount, Amazon
// ECS will spawn another instantiation of the task in the specified cluster.
func (c *ECS) CreateService(input *CreateServiceInput) (output *CreateServiceOutput, err error) {
	req, out := c.CreateServiceRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateService *aws.Operation

// DeleteClusterRequest generates a request for the DeleteCluster operation.
func (c *ECS) DeleteClusterRequest(input *DeleteClusterInput) (req *aws.Request, output *DeleteClusterOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteCluster == nil {
		opDeleteCluster = &aws.Operation{
			Name:       "DeleteCluster",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDeleteCluster, input, output)
	output = &DeleteClusterOutput{}
	req.Data = output
	return
}

// Deletes the specified cluster. You must deregister all container instances
// from this cluster before you may delete it. You can list the container instances
// in a cluster with ListContainerInstances and deregister them with DeregisterContainerInstance.
func (c *ECS) DeleteCluster(input *DeleteClusterInput) (output *DeleteClusterOutput, err error) {
	req, out := c.DeleteClusterRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteCluster *aws.Operation

// DeleteServiceRequest generates a request for the DeleteService operation.
func (c *ECS) DeleteServiceRequest(input *DeleteServiceInput) (req *aws.Request, output *DeleteServiceOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteService == nil {
		opDeleteService = &aws.Operation{
			Name:       "DeleteService",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDeleteService, input, output)
	output = &DeleteServiceOutput{}
	req.Data = output
	return
}

// Deletes a specified service within a cluster.
func (c *ECS) DeleteService(input *DeleteServiceInput) (output *DeleteServiceOutput, err error) {
	req, out := c.DeleteServiceRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteService *aws.Operation

// DeregisterContainerInstanceRequest generates a request for the DeregisterContainerInstance operation.
func (c *ECS) DeregisterContainerInstanceRequest(input *DeregisterContainerInstanceInput) (req *aws.Request, output *DeregisterContainerInstanceOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeregisterContainerInstance == nil {
		opDeregisterContainerInstance = &aws.Operation{
			Name:       "DeregisterContainerInstance",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDeregisterContainerInstance, input, output)
	output = &DeregisterContainerInstanceOutput{}
	req.Data = output
	return
}

// Deregisters an Amazon ECS container instance from the specified cluster.
// This instance will no longer be available to run tasks.
func (c *ECS) DeregisterContainerInstance(input *DeregisterContainerInstanceInput) (output *DeregisterContainerInstanceOutput, err error) {
	req, out := c.DeregisterContainerInstanceRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeregisterContainerInstance *aws.Operation

// DeregisterTaskDefinitionRequest generates a request for the DeregisterTaskDefinition operation.
func (c *ECS) DeregisterTaskDefinitionRequest(input *DeregisterTaskDefinitionInput) (req *aws.Request, output *DeregisterTaskDefinitionOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeregisterTaskDefinition == nil {
		opDeregisterTaskDefinition = &aws.Operation{
			Name:       "DeregisterTaskDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDeregisterTaskDefinition, input, output)
	output = &DeregisterTaskDefinitionOutput{}
	req.Data = output
	return
}

// NOT YET IMPLEMENTED.
//
// Deregisters the specified task definition. You will no longer be able to
// run tasks from this definition after deregistration.
func (c *ECS) DeregisterTaskDefinition(input *DeregisterTaskDefinitionInput) (output *DeregisterTaskDefinitionOutput, err error) {
	req, out := c.DeregisterTaskDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeregisterTaskDefinition *aws.Operation

// DescribeClustersRequest generates a request for the DescribeClusters operation.
func (c *ECS) DescribeClustersRequest(input *DescribeClustersInput) (req *aws.Request, output *DescribeClustersOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeClusters == nil {
		opDescribeClusters = &aws.Operation{
			Name:       "DescribeClusters",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDescribeClusters, input, output)
	output = &DescribeClustersOutput{}
	req.Data = output
	return
}

// Describes one or more of your clusters.
func (c *ECS) DescribeClusters(input *DescribeClustersInput) (output *DescribeClustersOutput, err error) {
	req, out := c.DescribeClustersRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeClusters *aws.Operation

// DescribeContainerInstancesRequest generates a request for the DescribeContainerInstances operation.
func (c *ECS) DescribeContainerInstancesRequest(input *DescribeContainerInstancesInput) (req *aws.Request, output *DescribeContainerInstancesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeContainerInstances == nil {
		opDescribeContainerInstances = &aws.Operation{
			Name:       "DescribeContainerInstances",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDescribeContainerInstances, input, output)
	output = &DescribeContainerInstancesOutput{}
	req.Data = output
	return
}

// Describes Amazon EC2 Container Service container instances. Returns metadata
// about registered and remaining resources on each container instance requested.
func (c *ECS) DescribeContainerInstances(input *DescribeContainerInstancesInput) (output *DescribeContainerInstancesOutput, err error) {
	req, out := c.DescribeContainerInstancesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeContainerInstances *aws.Operation

// DescribeServicesRequest generates a request for the DescribeServices operation.
func (c *ECS) DescribeServicesRequest(input *DescribeServicesInput) (req *aws.Request, output *DescribeServicesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeServices == nil {
		opDescribeServices = &aws.Operation{
			Name:       "DescribeServices",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDescribeServices, input, output)
	output = &DescribeServicesOutput{}
	req.Data = output
	return
}

// Describes the specified services running in your cluster.
func (c *ECS) DescribeServices(input *DescribeServicesInput) (output *DescribeServicesOutput, err error) {
	req, out := c.DescribeServicesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeServices *aws.Operation

// DescribeTaskDefinitionRequest generates a request for the DescribeTaskDefinition operation.
func (c *ECS) DescribeTaskDefinitionRequest(input *DescribeTaskDefinitionInput) (req *aws.Request, output *DescribeTaskDefinitionOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeTaskDefinition == nil {
		opDescribeTaskDefinition = &aws.Operation{
			Name:       "DescribeTaskDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDescribeTaskDefinition, input, output)
	output = &DescribeTaskDefinitionOutput{}
	req.Data = output
	return
}

// Describes a task definition. You can specify a family and revision to find
// information on a specific task definition, or you can simply specify the
// family to find the latest revision in that family.
func (c *ECS) DescribeTaskDefinition(input *DescribeTaskDefinitionInput) (output *DescribeTaskDefinitionOutput, err error) {
	req, out := c.DescribeTaskDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTaskDefinition *aws.Operation

// DescribeTasksRequest generates a request for the DescribeTasks operation.
func (c *ECS) DescribeTasksRequest(input *DescribeTasksInput) (req *aws.Request, output *DescribeTasksOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeTasks == nil {
		opDescribeTasks = &aws.Operation{
			Name:       "DescribeTasks",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDescribeTasks, input, output)
	output = &DescribeTasksOutput{}
	req.Data = output
	return
}

// Describes a specified task or tasks.
func (c *ECS) DescribeTasks(input *DescribeTasksInput) (output *DescribeTasksOutput, err error) {
	req, out := c.DescribeTasksRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTasks *aws.Operation

// DiscoverPollEndpointRequest generates a request for the DiscoverPollEndpoint operation.
func (c *ECS) DiscoverPollEndpointRequest(input *DiscoverPollEndpointInput) (req *aws.Request, output *DiscoverPollEndpointOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDiscoverPollEndpoint == nil {
		opDiscoverPollEndpoint = &aws.Operation{
			Name:       "DiscoverPollEndpoint",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opDiscoverPollEndpoint, input, output)
	output = &DiscoverPollEndpointOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Returns an endpoint for the Amazon EC2 Container Service agent to poll for
// updates.
func (c *ECS) DiscoverPollEndpoint(input *DiscoverPollEndpointInput) (output *DiscoverPollEndpointOutput, err error) {
	req, out := c.DiscoverPollEndpointRequest(input)
	output = out
	err = req.Send()
	return
}

var opDiscoverPollEndpoint *aws.Operation

// ListClustersRequest generates a request for the ListClusters operation.
func (c *ECS) ListClustersRequest(input *ListClustersInput) (req *aws.Request, output *ListClustersOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListClusters == nil {
		opListClusters = &aws.Operation{
			Name:       "ListClusters",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opListClusters, input, output)
	output = &ListClustersOutput{}
	req.Data = output
	return
}

// Returns a list of existing clusters.
func (c *ECS) ListClusters(input *ListClustersInput) (output *ListClustersOutput, err error) {
	req, out := c.ListClustersRequest(input)
	output = out
	err = req.Send()
	return
}

var opListClusters *aws.Operation

// ListContainerInstancesRequest generates a request for the ListContainerInstances operation.
func (c *ECS) ListContainerInstancesRequest(input *ListContainerInstancesInput) (req *aws.Request, output *ListContainerInstancesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListContainerInstances == nil {
		opListContainerInstances = &aws.Operation{
			Name:       "ListContainerInstances",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opListContainerInstances, input, output)
	output = &ListContainerInstancesOutput{}
	req.Data = output
	return
}

// Returns a list of container instances in a specified cluster.
func (c *ECS) ListContainerInstances(input *ListContainerInstancesInput) (output *ListContainerInstancesOutput, err error) {
	req, out := c.ListContainerInstancesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListContainerInstances *aws.Operation

// ListServicesRequest generates a request for the ListServices operation.
func (c *ECS) ListServicesRequest(input *ListServicesInput) (req *aws.Request, output *ListServicesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListServices == nil {
		opListServices = &aws.Operation{
			Name:       "ListServices",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opListServices, input, output)
	output = &ListServicesOutput{}
	req.Data = output
	return
}

// Lists the services that are running in a specified cluster.
func (c *ECS) ListServices(input *ListServicesInput) (output *ListServicesOutput, err error) {
	req, out := c.ListServicesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListServices *aws.Operation

// ListTaskDefinitionFamiliesRequest generates a request for the ListTaskDefinitionFamilies operation.
func (c *ECS) ListTaskDefinitionFamiliesRequest(input *ListTaskDefinitionFamiliesInput) (req *aws.Request, output *ListTaskDefinitionFamiliesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListTaskDefinitionFamilies == nil {
		opListTaskDefinitionFamilies = &aws.Operation{
			Name:       "ListTaskDefinitionFamilies",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opListTaskDefinitionFamilies, input, output)
	output = &ListTaskDefinitionFamiliesOutput{}
	req.Data = output
	return
}

// Returns a list of task definition families that are registered to your account.
// You can filter the results with the familyPrefix parameter.
func (c *ECS) ListTaskDefinitionFamilies(input *ListTaskDefinitionFamiliesInput) (output *ListTaskDefinitionFamiliesOutput, err error) {
	req, out := c.ListTaskDefinitionFamiliesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTaskDefinitionFamilies *aws.Operation

// ListTaskDefinitionsRequest generates a request for the ListTaskDefinitions operation.
func (c *ECS) ListTaskDefinitionsRequest(input *ListTaskDefinitionsInput) (req *aws.Request, output *ListTaskDefinitionsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListTaskDefinitions == nil {
		opListTaskDefinitions = &aws.Operation{
			Name:       "ListTaskDefinitions",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opListTaskDefinitions, input, output)
	output = &ListTaskDefinitionsOutput{}
	req.Data = output
	return
}

// Returns a list of task definitions that are registered to your account. You
// can filter the results by family name with the familyPrefix parameter.
func (c *ECS) ListTaskDefinitions(input *ListTaskDefinitionsInput) (output *ListTaskDefinitionsOutput, err error) {
	req, out := c.ListTaskDefinitionsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTaskDefinitions *aws.Operation

// ListTasksRequest generates a request for the ListTasks operation.
func (c *ECS) ListTasksRequest(input *ListTasksInput) (req *aws.Request, output *ListTasksOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListTasks == nil {
		opListTasks = &aws.Operation{
			Name:       "ListTasks",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opListTasks, input, output)
	output = &ListTasksOutput{}
	req.Data = output
	return
}

// Returns a list of tasks for a specified cluster. You can filter the results
// by family name or by a particular container instance with the family and
// containerInstance parameters.
func (c *ECS) ListTasks(input *ListTasksInput) (output *ListTasksOutput, err error) {
	req, out := c.ListTasksRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTasks *aws.Operation

// RegisterContainerInstanceRequest generates a request for the RegisterContainerInstance operation.
func (c *ECS) RegisterContainerInstanceRequest(input *RegisterContainerInstanceInput) (req *aws.Request, output *RegisterContainerInstanceOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opRegisterContainerInstance == nil {
		opRegisterContainerInstance = &aws.Operation{
			Name:       "RegisterContainerInstance",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opRegisterContainerInstance, input, output)
	output = &RegisterContainerInstanceOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Registers an Amazon EC2 instance into the specified cluster. This instance
// will become available to place containers on.
func (c *ECS) RegisterContainerInstance(input *RegisterContainerInstanceInput) (output *RegisterContainerInstanceOutput, err error) {
	req, out := c.RegisterContainerInstanceRequest(input)
	output = out
	err = req.Send()
	return
}

var opRegisterContainerInstance *aws.Operation

// RegisterTaskDefinitionRequest generates a request for the RegisterTaskDefinition operation.
func (c *ECS) RegisterTaskDefinitionRequest(input *RegisterTaskDefinitionInput) (req *aws.Request, output *RegisterTaskDefinitionOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opRegisterTaskDefinition == nil {
		opRegisterTaskDefinition = &aws.Operation{
			Name:       "RegisterTaskDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opRegisterTaskDefinition, input, output)
	output = &RegisterTaskDefinitionOutput{}
	req.Data = output
	return
}

// Registers a new task definition from the supplied family and containerDefinitions.
// Optionally, you can add data volumes to your containers with the volumes
// parameter. For more information on task definition parameters and defaults,
// see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
// in the Amazon EC2 Container Service Developer Guide.
func (c *ECS) RegisterTaskDefinition(input *RegisterTaskDefinitionInput) (output *RegisterTaskDefinitionOutput, err error) {
	req, out := c.RegisterTaskDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opRegisterTaskDefinition *aws.Operation

// RunTaskRequest generates a request for the RunTask operation.
func (c *ECS) RunTaskRequest(input *RunTaskInput) (req *aws.Request, output *RunTaskOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opRunTask == nil {
		opRunTask = &aws.Operation{
			Name:       "RunTask",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opRunTask, input, output)
	output = &RunTaskOutput{}
	req.Data = output
	return
}

// Start a task using random placement and the default Amazon ECS scheduler.
// If you want to use your own scheduler or place a task on a specific container
// instance, use StartTask instead.
//
//  The count parameter is limited to 10 tasks per call.
func (c *ECS) RunTask(input *RunTaskInput) (output *RunTaskOutput, err error) {
	req, out := c.RunTaskRequest(input)
	output = out
	err = req.Send()
	return
}

var opRunTask *aws.Operation

// StartTaskRequest generates a request for the StartTask operation.
func (c *ECS) StartTaskRequest(input *StartTaskInput) (req *aws.Request, output *StartTaskOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opStartTask == nil {
		opStartTask = &aws.Operation{
			Name:       "StartTask",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opStartTask, input, output)
	output = &StartTaskOutput{}
	req.Data = output
	return
}

// Starts a new task from the specified task definition on the specified container
// instance or instances. If you want to use the default Amazon ECS scheduler
// to place your task, use RunTask instead.
//
//  The list of container instances to start tasks on is limited to 10.
func (c *ECS) StartTask(input *StartTaskInput) (output *StartTaskOutput, err error) {
	req, out := c.StartTaskRequest(input)
	output = out
	err = req.Send()
	return
}

var opStartTask *aws.Operation

// StopTaskRequest generates a request for the StopTask operation.
func (c *ECS) StopTaskRequest(input *StopTaskInput) (req *aws.Request, output *StopTaskOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opStopTask == nil {
		opStopTask = &aws.Operation{
			Name:       "StopTask",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opStopTask, input, output)
	output = &StopTaskOutput{}
	req.Data = output
	return
}

// Stops a running task.
func (c *ECS) StopTask(input *StopTaskInput) (output *StopTaskOutput, err error) {
	req, out := c.StopTaskRequest(input)
	output = out
	err = req.Send()
	return
}

var opStopTask *aws.Operation

// SubmitContainerStateChangeRequest generates a request for the SubmitContainerStateChange operation.
func (c *ECS) SubmitContainerStateChangeRequest(input *SubmitContainerStateChangeInput) (req *aws.Request, output *SubmitContainerStateChangeOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opSubmitContainerStateChange == nil {
		opSubmitContainerStateChange = &aws.Operation{
			Name:       "SubmitContainerStateChange",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opSubmitContainerStateChange, input, output)
	output = &SubmitContainerStateChangeOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Sent to acknowledge that a container changed states.
func (c *ECS) SubmitContainerStateChange(input *SubmitContainerStateChangeInput) (output *SubmitContainerStateChangeOutput, err error) {
	req, out := c.SubmitContainerStateChangeRequest(input)
	output = out
	err = req.Send()
	return
}

var opSubmitContainerStateChange *aws.Operation

// SubmitTaskStateChangeRequest generates a request for the SubmitTaskStateChange operation.
func (c *ECS) SubmitTaskStateChangeRequest(input *SubmitTaskStateChangeInput) (req *aws.Request, output *SubmitTaskStateChangeOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opSubmitTaskStateChange == nil {
		opSubmitTaskStateChange = &aws.Operation{
			Name:       "SubmitTaskStateChange",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opSubmitTaskStateChange, input, output)
	output = &SubmitTaskStateChangeOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Sent to acknowledge that a task changed states.
func (c *ECS) SubmitTaskStateChange(input *SubmitTaskStateChangeInput) (output *SubmitTaskStateChangeOutput, err error) {
	req, out := c.SubmitTaskStateChangeRequest(input)
	output = out
	err = req.Send()
	return
}

var opSubmitTaskStateChange *aws.Operation

// UpdateContainerAgentRequest generates a request for the UpdateContainerAgent operation.
func (c *ECS) UpdateContainerAgentRequest(input *UpdateContainerAgentInput) (req *aws.Request, output *UpdateContainerAgentOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opUpdateContainerAgent == nil {
		opUpdateContainerAgent = &aws.Operation{
			Name:       "UpdateContainerAgent",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opUpdateContainerAgent, input, output)
	output = &UpdateContainerAgentOutput{}
	req.Data = output
	return
}

// Updates the Amazon ECS container agent on a specified container instance.
func (c *ECS) UpdateContainerAgent(input *UpdateContainerAgentInput) (output *UpdateContainerAgentOutput, err error) {
	req, out := c.UpdateContainerAgentRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateContainerAgent *aws.Operation

// UpdateServiceRequest generates a request for the UpdateService operation.
func (c *ECS) UpdateServiceRequest(input *UpdateServiceInput) (req *aws.Request, output *UpdateServiceOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opUpdateService == nil {
		opUpdateService = &aws.Operation{
			Name:       "UpdateService",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = c.newRequest(opUpdateService, input, output)
	output = &UpdateServiceOutput{}
	req.Data = output
	return
}

// Modify the desired count or task definition used in a service.
//
// You can add to or subtract from the number of instantiations of a task definition
// in a service by specifying the cluster that the service is running in and
// a new desiredCount parameter.
//
// You can use UpdateService to modify your task definition and deploy a new
// version of your service, one task at a time. If you modify the task definition
// with UpdateService, Amazon ECS spawns a task with the new version of the
// task definition and then stops an old task after the new version is running.
// Because UpdateService starts a new version of the task before stopping an
// old version, your cluster must have capacity to support one more instantiation
// of the task when UpdateService is run. If your cluster cannot support another
// instantiation of the task used in your service, you can reduce the desired
// count of your service by one before modifying the task definition.
func (c *ECS) UpdateService(input *UpdateServiceInput) (output *UpdateServiceOutput, err error) {
	req, out := c.UpdateServiceRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateService *aws.Operation

// A regional grouping of one or more container instances on which you can run
// task requests. Each account receives a default cluster the first time you
// use the Amazon ECS service, but you may also create other clusters. Clusters
// may contain more than one instance type simultaneously.
//
//  During the preview, each account is limited to two clusters.
type Cluster struct {
	// The number of services that are running on the cluster in an ACTIVE state.
	// You can view these services with ListServices.
	ActiveServicesCount *int64 `locationName:"activeServicesCount" type:"integer"`

	// The Amazon Resource Name (ARN) that identifies the cluster. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the cluster, the AWS
	// account ID of the cluster owner, the cluster namespace, and then the cluster
	// name. For example, arn:aws:ecs:region:012345678910:cluster/test.
	ClusterARN *string `locationName:"clusterArn" type:"string"`

	// A user-generated string that you can use to identify your cluster.
	ClusterName *string `locationName:"clusterName" type:"string"`

	// The number of tasks in the cluster that are in the PENDING state.
	PendingTasksCount *int64 `locationName:"pendingTasksCount" type:"integer"`

	// The number of container instances registered into the cluster.
	RegisteredContainerInstancesCount *int64 `locationName:"registeredContainerInstancesCount" type:"integer"`

	// The number of tasks in the cluster that are in the RUNNING state.
	RunningTasksCount *int64 `locationName:"runningTasksCount" type:"integer"`

	// The status of the cluster. The valid values are ACTIVE or INACTIVE. ACTIVE
	// indicates that you can register container instances with the cluster and
	// the associated instances can accept tasks.
	Status *string `locationName:"status" type:"string"`

	metadataCluster `json:"-", xml:"-"`
}

type metadataCluster struct {
	SDKShapeTraits bool `type:"structure"`
}

type Container struct {
	// The Amazon Resource Name (ARN) of the container.
	ContainerARN *string `locationName:"containerArn" type:"string"`

	// The exit code returned from the container.
	ExitCode *int64 `locationName:"exitCode" type:"integer"`

	// The last known status of the container.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// The name of the container.
	Name *string `locationName:"name" type:"string"`

	NetworkBindings []*NetworkBinding `locationName:"networkBindings" type:"list"`

	// A short (255 max characters) human-readable string to provide additional
	// detail about a running or stopped container.
	Reason *string `locationName:"reason" type:"string"`

	// The Amazon Resource Name (ARN) of the task.
	TaskARN *string `locationName:"taskArn" type:"string"`

	metadataContainer `json:"-", xml:"-"`
}

type metadataContainer struct {
	SDKShapeTraits bool `type:"structure"`
}

// Container definitions are used in task definitions to describe the different
// containers that are launched as part of a task.
type ContainerDefinition struct {
	// The number of cpu units reserved for the container. A container instance
	// has 1,024 cpu units for every CPU core. This parameter specifies the minimum
	// amount of CPU to reserve for a container, and containers share unallocated
	// CPU units with other containers on the instance with the same ratio as their
	// allocated amount.
	//
	// For example, if you run a single-container task on a single-core instance
	// type with 512 CPU units specified for that container, and that is the only
	// task running on the container instance, that container could use the full
	// 1,024 CPU unit share at any given time. However, if you launched another
	// copy of the same task on that container instance, each task would be guaranteed
	// a minimum of 512 CPU units when needed, and each container could float to
	// higher CPU usage if the other container was not using it, but if both tasks
	// were 100% active all of the time, they would be limited to 512 CPU units.
	//
	// If this parameter is omitted, 0 CPU units are reserved for the container,
	// and it will only receive CPU time when other containers are not using it.
	CPU *int64 `locationName:"cpu" type:"integer"`

	// The CMD that is passed to the container. For more information on the Docker
	// CMD parameter, see https://docs.docker.com/reference/builder/#cmd (https://docs.docker.com/reference/builder/#cmd).
	Command []*string `locationName:"command" type:"list"`

	// Early versions of the Amazon ECS container agent do not properly handle entryPoint
	// parameters. If you have problems using entryPoint, update your container
	// agent or enter your commands and arguments as command array items instead.
	//
	//  The ENTRYPOINT that is passed to the container. For more information on
	// the Docker ENTRYPOINT parameter, see https://docs.docker.com/reference/builder/#entrypoint
	// (https://docs.docker.com/reference/builder/#entrypoint).
	EntryPoint []*string `locationName:"entryPoint" type:"list"`

	// The environment variables to pass to a container.
	Environment []*KeyValuePair `locationName:"environment" type:"list"`

	// If the essential parameter of a container is marked as true, the failure
	// of that container will stop the task. If the essential parameter of a container
	// is marked as false, then its failure will not affect the rest of the containers
	// in a task. If this parameter is omitted, a container is assumed to be essential.
	//
	//  All tasks must have at least one essential container.
	Essential *bool `locationName:"essential" type:"boolean"`

	// The image used to start a container. This string is passed directly to the
	// Docker daemon. Images in the Docker Hub registry are available by default.
	// Other repositories are specified with repository-url/image:tag.
	Image *string `locationName:"image" type:"string"`

	// The link parameter allows containers to communicate with each other without
	// the need for port mappings, using the name parameter. The name:internalName
	// construct is analogous to name:alias in Docker links. For more information
	// on linking Docker containers, see https://docs.docker.com/userguide/dockerlinks/
	// (https://docs.docker.com/userguide/dockerlinks/).
	Links []*string `locationName:"links" type:"list"`

	// The number of MiB of memory reserved for the container. If your container
	// attempts to exceed the memory allocated here, the container is killed.
	Memory *int64 `locationName:"memory" type:"integer"`

	// The mount points for data volumes in your container.
	MountPoints []*MountPoint `locationName:"mountPoints" type:"list"`

	// The name of a container. If you are linking multiple containers together
	// in a task definition, the name of one container can be entered in the links
	// of another container to connect the containers.
	Name *string `locationName:"name" type:"string"`

	// The list of port mappings for the container.
	PortMappings []*PortMapping `locationName:"portMappings" type:"list"`

	// Data volumes to mount from another container.
	VolumesFrom []*VolumeFrom `locationName:"volumesFrom" type:"list"`

	metadataContainerDefinition `json:"-", xml:"-"`
}

type metadataContainerDefinition struct {
	SDKShapeTraits bool `type:"structure"`
}

// An Amazon EC2 instance that is running the Amazon ECS agent and has been
// registered with a cluster.
type ContainerInstance struct {
	// This parameter returns true if the agent is actually connected to Amazon
	// ECS. Registered instances with an agent that may be unhealthy or stopped
	// will return false, and instances without a connected agent cannot accept
	// placement request.
	AgentConnected *bool `locationName:"agentConnected" type:"boolean"`

	// The status of the most recent agent update. If an update has never been requested,
	// this value is NULL.
	AgentUpdateStatus *string `locationName:"agentUpdateStatus" type:"string"`

	// The Amazon Resource Name (ARN) of the container instance. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the container instance,
	// the AWS account ID of the container instance owner, the container-instance
	// namespace, and then the container instance UUID. For example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_UUID.
	ContainerInstanceARN *string `locationName:"containerInstanceArn" type:"string"`

	// The Amazon EC2 instance ID of the container instance.
	EC2InstanceID *string `locationName:"ec2InstanceId" type:"string"`

	// The number of tasks on the container instance that are in the PENDING status.
	PendingTasksCount *int64 `locationName:"pendingTasksCount" type:"integer"`

	// The registered resources on the container instance that are in use by current
	// tasks.
	RegisteredResources []*Resource `locationName:"registeredResources" type:"list"`

	// The remaining resources of the container instance that are available for
	// new tasks.
	RemainingResources []*Resource `locationName:"remainingResources" type:"list"`

	// The number of tasks on the container instance that are in the RUNNING status.
	RunningTasksCount *int64 `locationName:"runningTasksCount" type:"integer"`

	// The status of the container instance. The valid values are ACTIVE or INACTIVE.
	// ACTIVE indicates that the container instance can accept tasks.
	Status *string `locationName:"status" type:"string"`

	VersionInfo *VersionInfo `locationName:"versionInfo" type:"structure"`

	metadataContainerInstance `json:"-", xml:"-"`
}

type metadataContainerInstance struct {
	SDKShapeTraits bool `type:"structure"`
}

// The name of a container in a task definition and the command it should run
// instead of its default.
type ContainerOverride struct {
	// The command to send to the container that overrides the default command from
	// the Docker image or the task definition.
	Command []*string `locationName:"command" type:"list"`

	// The name of the container that receives the override.
	Name *string `locationName:"name" type:"string"`

	metadataContainerOverride `json:"-", xml:"-"`
}

type metadataContainerOverride struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateClusterInput struct {
	// The name of your cluster. If you do not specify a name for your cluster,
	// you will create a cluster named default. Up to 255 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	ClusterName *string `locationName:"clusterName" type:"string"`

	metadataCreateClusterInput `json:"-", xml:"-"`
}

type metadataCreateClusterInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateClusterOutput struct {
	// The full description of your new cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`

	metadataCreateClusterOutput `json:"-", xml:"-"`
}

type metadataCreateClusterOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateServiceInput struct {
	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. Up to 32 ASCII characters are allowed.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to run your service on. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the specified task definition that you would
	// like to place and keep running on your cluster.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer" required:"true"`

	// A list of load balancer objects, containing the load balancer name, the container
	// name (as it appears in a container definition), and the container port to
	// access from the load balancer.
	LoadBalancers []*LoadBalancer `locationName:"loadBalancers" type:"list"`

	// The name or full Amazon Resource Name (ARN) of the IAM role that allows your
	// Amazon ECS container agent to make calls to your load balancer on your behalf.
	// This parameter is only required if you are using a load balancer with your
	// service.
	Role *string `locationName:"role" type:"string"`

	// The name of your service. Up to 255 letters (uppercase and lowercase), numbers,
	// hyphens, and underscores are allowed.
	ServiceName *string `locationName:"serviceName" type:"string" required:"true"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run in your service.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataCreateServiceInput `json:"-", xml:"-"`
}

type metadataCreateServiceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateServiceOutput struct {
	// The full description of your service following the create call.
	Service *Service `locationName:"service" type:"structure"`

	metadataCreateServiceOutput `json:"-", xml:"-"`
}

type metadataCreateServiceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteClusterInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to delete.
	Cluster *string `locationName:"cluster" type:"string" required:"true"`

	metadataDeleteClusterInput `json:"-", xml:"-"`
}

type metadataDeleteClusterInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteClusterOutput struct {
	// The full description of the deleted cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`

	metadataDeleteClusterOutput `json:"-", xml:"-"`
}

type metadataDeleteClusterOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteServiceInput struct {
	// The name of the cluster that hosts the service you want to delete.
	Cluster *string `locationName:"cluster" type:"string"`

	// The name of the service you want to delete.
	Service *string `locationName:"service" type:"string" required:"true"`

	metadataDeleteServiceInput `json:"-", xml:"-"`
}

type metadataDeleteServiceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteServiceOutput struct {
	Service *Service `locationName:"service" type:"structure"`

	metadataDeleteServiceOutput `json:"-", xml:"-"`
}

type metadataDeleteServiceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Deployment struct {
	// The Unix time in seconds and milliseconds when the service was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The most recent desired count of tasks that was specified for the service
	// to deploy and/or maintain.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The ID of the deployment.
	ID *string `locationName:"id" type:"string"`

	// The number of tasks in the deployment that are in the PENDING status.
	PendingCount *int64 `locationName:"pendingCount" type:"integer"`

	// The number of tasks in the deployment that are in the RUNNING status.
	RunningCount *int64 `locationName:"runningCount" type:"integer"`

	// The status of the deployment. Valid values are PRIMARY (for the most recent
	// deployment), ACTIVE (for previous deployments that still have tasks running,
	// but are being replaced with the PRIMARY deployment), and INACTIVE (for deployments
	// that have been completely replaced).
	Status *string `locationName:"status" type:"string"`

	// The most recent task definition that was specified for the service to use.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`

	// The Unix time in seconds and milliseconds when the service was last updated.
	UpdatedAt *time.Time `locationName:"updatedAt" type:"timestamp" timestampFormat:"unix"`

	metadataDeployment `json:"-", xml:"-"`
}

type metadataDeployment struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeregisterContainerInstanceInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instance you want to deregister. If you do not specify a cluster,
	// the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) of the container
	// instance you want to deregister. The ARN contains the arn:aws:ecs namespace,
	// followed by the region of the container instance, the AWS account ID of the
	// container instance owner, the container-instance namespace, and then the
	// container instance UUID. For example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_UUID.
	ContainerInstance *string `locationName:"containerInstance" type:"string" required:"true"`

	// Force the deregistration of the container instance. You can use the force
	// parameter if you have several tasks running on a container instance and you
	// don't want to run StopTask for each task before deregistering the container
	// instance.
	Force *bool `locationName:"force" type:"boolean"`

	metadataDeregisterContainerInstanceInput `json:"-", xml:"-"`
}

type metadataDeregisterContainerInstanceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeregisterContainerInstanceOutput struct {
	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`

	metadataDeregisterContainerInstanceOutput `json:"-", xml:"-"`
}

type metadataDeregisterContainerInstanceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeregisterTaskDefinitionInput struct {
	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to deregister.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataDeregisterTaskDefinitionInput `json:"-", xml:"-"`
}

type metadataDeregisterTaskDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeregisterTaskDefinitionOutput struct {
	// The full description of the deregistered task.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`

	metadataDeregisterTaskDefinitionOutput `json:"-", xml:"-"`
}

type metadataDeregisterTaskDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeClustersInput struct {
	// A space-separated list of cluster names or full cluster Amazon Resource Name
	// (ARN) entries. If you do not specify a cluster, the default cluster is assumed.
	Clusters []*string `locationName:"clusters" type:"list"`

	metadataDescribeClustersInput `json:"-", xml:"-"`
}

type metadataDescribeClustersInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeClustersOutput struct {
	// The list of clusters.
	Clusters []*Cluster `locationName:"clusters" type:"list"`

	Failures []*Failure `locationName:"failures" type:"list"`

	metadataDescribeClustersOutput `json:"-", xml:"-"`
}

type metadataDescribeClustersOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeContainerInstancesInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instances you want to describe. If you do not specify a cluster,
	// the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A space-separated list of container instance UUIDs or full Amazon Resource
	// Name (ARN) entries.
	ContainerInstances []*string `locationName:"containerInstances" type:"list" required:"true"`

	metadataDescribeContainerInstancesInput `json:"-", xml:"-"`
}

type metadataDescribeContainerInstancesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeContainerInstancesOutput struct {
	// The list of container instances.
	ContainerInstances []*ContainerInstance `locationName:"containerInstances" type:"list"`

	Failures []*Failure `locationName:"failures" type:"list"`

	metadataDescribeContainerInstancesOutput `json:"-", xml:"-"`
}

type metadataDescribeContainerInstancesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeServicesInput struct {
	// The name of the cluster that hosts the service you want to describe.
	Cluster *string `locationName:"cluster" type:"string"`

	// A list of services you want to describe.
	Services []*string `locationName:"services" type:"list" required:"true"`

	metadataDescribeServicesInput `json:"-", xml:"-"`
}

type metadataDescribeServicesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeServicesOutput struct {
	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`

	// The list of services described.
	Services []*Service `locationName:"services" type:"list"`

	metadataDescribeServicesOutput `json:"-", xml:"-"`
}

type metadataDescribeServicesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeTaskDefinitionInput struct {
	// The family for the latest revision, family and revision (family:revision)
	// for a specific revision in the family, or full Amazon Resource Name (ARN)
	// of the task definition that you want to describe.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataDescribeTaskDefinitionInput `json:"-", xml:"-"`
}

type metadataDescribeTaskDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeTaskDefinitionOutput struct {
	// The full task definition description.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`

	metadataDescribeTaskDefinitionOutput `json:"-", xml:"-"`
}

type metadataDescribeTaskDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeTasksInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task you want to describe. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A space-separated list of task UUIDs or full Amazon Resource Name (ARN) entries.
	Tasks []*string `locationName:"tasks" type:"list" required:"true"`

	metadataDescribeTasksInput `json:"-", xml:"-"`
}

type metadataDescribeTasksInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeTasksOutput struct {
	Failures []*Failure `locationName:"failures" type:"list"`

	// The list of tasks.
	Tasks []*Task `locationName:"tasks" type:"list"`

	metadataDescribeTasksOutput `json:"-", xml:"-"`
}

type metadataDescribeTasksOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DiscoverPollEndpointInput struct {
	// The cluster that the container instance belongs to.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) of the container
	// instance. The ARN contains the arn:aws:ecs namespace, followed by the region
	// of the container instance, the AWS account ID of the container instance owner,
	// the container-instance namespace, and then the container instance UUID. For
	// example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_UUID.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	metadataDiscoverPollEndpointInput `json:"-", xml:"-"`
}

type metadataDiscoverPollEndpointInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DiscoverPollEndpointOutput struct {
	// The endpoint for the Amazon ECS agent to poll.
	Endpoint *string `locationName:"endpoint" type:"string"`

	// The telemetry endpoint for the Amazon ECS agent.
	TelemetryEndpoint *string `locationName:"telemetryEndpoint" type:"string"`

	metadataDiscoverPollEndpointOutput `json:"-", xml:"-"`
}

type metadataDiscoverPollEndpointOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Failure struct {
	// The Amazon Resource Name (ARN) of the failed resource.
	ARN *string `locationName:"arn" type:"string"`

	// The reason for the failure.
	Reason *string `locationName:"reason" type:"string"`

	metadataFailure `json:"-", xml:"-"`
}

type metadataFailure struct {
	SDKShapeTraits bool `type:"structure"`
}

type HostVolumeProperties struct {
	// The path on the host container instance that is presented to the container.
	// If this parameter is empty, then the Docker daemon has assigned a host path
	// for you.
	SourcePath *string `locationName:"sourcePath" type:"string"`

	metadataHostVolumeProperties `json:"-", xml:"-"`
}

type metadataHostVolumeProperties struct {
	SDKShapeTraits bool `type:"structure"`
}

type KeyValuePair struct {
	// The name of the key value pair.
	Name *string `locationName:"name" type:"string"`

	// The value of the key value pair.
	Value *string `locationName:"value" type:"string"`

	metadataKeyValuePair `json:"-", xml:"-"`
}

type metadataKeyValuePair struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListClustersInput struct {
	// The maximum number of cluster results returned by ListClusters in paginated
	// output. When this parameter is used, ListClusters only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListClusters
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListClusters returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListClusters request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListClustersInput `json:"-", xml:"-"`
}

type metadataListClustersInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListClustersOutput struct {
	// The list of full Amazon Resource Name (ARN) entries for each cluster associated
	// with your account.
	ClusterARNs []*string `locationName:"clusterArns" type:"list"`

	// The nextToken value to include in a future ListClusters request. When the
	// results of a ListClusters request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListClustersOutput `json:"-", xml:"-"`
}

type metadataListClustersOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListContainerInstancesInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instances you want to list. If you do not specify a cluster,
	// the default cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The maximum number of container instance results returned by ListContainerInstances
	// in paginated output. When this parameter is used, ListContainerInstances
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListContainerInstances request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListContainerInstances returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListContainerInstances
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListContainerInstancesInput `json:"-", xml:"-"`
}

type metadataListContainerInstancesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListContainerInstancesOutput struct {
	// The list of container instance full Amazon Resource Name (ARN) entries for
	// each container instance associated with the specified cluster.
	ContainerInstanceARNs []*string `locationName:"containerInstanceArns" type:"list"`

	// The nextToken value to include in a future ListContainerInstances request.
	// When the results of a ListContainerInstances request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListContainerInstancesOutput `json:"-", xml:"-"`
}

type metadataListContainerInstancesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListServicesInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the services you want to list. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The maximum number of container instance results returned by ListServices
	// in paginated output. When this parameter is used, ListServices only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// ListServices request with the returned nextToken value. This value can be
	// between 1 and 100. If this parameter is not used, then ListServices returns
	// up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListServices request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListServicesInput `json:"-", xml:"-"`
}

type metadataListServicesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListServicesOutput struct {
	// The nextToken value to include in a future ListServices request. When the
	// results of a ListServices request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of full Amazon Resource Name (ARN) entries for each service associated
	// with the specified cluster.
	ServiceARNs []*string `locationName:"serviceArns" type:"list"`

	metadataListServicesOutput `json:"-", xml:"-"`
}

type metadataListServicesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTaskDefinitionFamiliesInput struct {
	// The familyPrefix is a string that is used to filter the results of ListTaskDefinitionFamilies.
	// If you specify a familyPrefix, only task definition family names that begin
	// with the familyPrefix string are returned.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition family results returned by ListTaskDefinitionFamilies
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitionFamilies request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListTaskDefinitionFamilies returns up to 100 results and a nextToken value
	// if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitionFamilies
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListTaskDefinitionFamiliesInput `json:"-", xml:"-"`
}

type metadataListTaskDefinitionFamiliesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTaskDefinitionFamiliesOutput struct {
	// The list of task definition family names that match the ListTaskDefinitionFamilies
	// request.
	Families []*string `locationName:"families" type:"list"`

	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults,
	// this value can be used to retrieve the next page of results. This value is
	// null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListTaskDefinitionFamiliesOutput `json:"-", xml:"-"`
}

type metadataListTaskDefinitionFamiliesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTaskDefinitionsInput struct {
	// The full family name that you want to filter the ListTaskDefinitions results
	// with. Specifying a familyPrefix will limit the listed task definitions to
	// task definition revisions that belong to that family.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition results returned by ListTaskDefinitions
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitions request with the returned nextToken value. This
	// value can be between 1 and 100. If this parameter is not used, then ListTaskDefinitions
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitions
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListTaskDefinitionsInput `json:"-", xml:"-"`
}

type metadataListTaskDefinitionsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTaskDefinitionsOutput struct {
	// The nextToken value to include in a future ListTaskDefinitions request. When
	// the results of a ListTaskDefinitions request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task definition Amazon Resource Name (ARN) entries for the ListTaskDefintions
	// request.
	TaskDefinitionARNs []*string `locationName:"taskDefinitionArns" type:"list"`

	metadataListTaskDefinitionsOutput `json:"-", xml:"-"`
}

type metadataListTaskDefinitionsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTasksInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the tasks you want to list. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) of the container
	// instance that you want to filter the ListTasks results with. Specifying a
	// containerInstance will limit the results to tasks that belong to that container
	// instance.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	// The task status that you want to filter the ListTasks results with. Specifying
	// a desiredStatus of STOPPED will limit the results to tasks that are in the
	// STOPPED status, which can be useful for debugging tasks that are not starting
	// properly or have died or finished. The default status filter is RUNNING.
	DesiredStatus *string `locationName:"desiredStatus" type:"string"`

	// The name of the family that you want to filter the ListTasks results with.
	// Specifying a family will limit the results to tasks that belong to that family.
	Family *string `locationName:"family" type:"string"`

	// The maximum number of task results returned by ListTasks in paginated output.
	// When this parameter is used, ListTasks only returns maxResults results in
	// a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListTasks request with
	// the returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, then ListTasks returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTasks request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the service that you want to filter the ListTasks results with.
	// Specifying a serviceName will limit the results to tasks that belong to that
	// service.
	ServiceName *string `locationName:"serviceName" type:"string"`

	// The startedBy value that you want to filter the task results with. Specifying
	// a startedBy value will limit the results to tasks that were started with
	// that value.
	StartedBy *string `locationName:"startedBy" type:"string"`

	metadataListTasksInput `json:"-", xml:"-"`
}

type metadataListTasksInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTasksOutput struct {
	// The nextToken value to include in a future ListTasks request. When the results
	// of a ListTasks request exceed maxResults, this value can be used to retrieve
	// the next page of results. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task Amazon Resource Name (ARN) entries for the ListTasks request.
	TaskARNs []*string `locationName:"taskArns" type:"list"`

	metadataListTasksOutput `json:"-", xml:"-"`
}

type metadataListTasksOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type LoadBalancer struct {
	// The name of the container to associate with the load balancer.
	ContainerName *string `locationName:"containerName" type:"string"`

	// The port on the container to associate with the load balancer. This port
	// must correspond to a containerPort in the service's task definition. Your
	// container instances must allow ingress traffic on the hostPort of the port
	// mapping.
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The name of the load balancer.
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string"`

	metadataLoadBalancer `json:"-", xml:"-"`
}

type metadataLoadBalancer struct {
	SDKShapeTraits bool `type:"structure"`
}

type MountPoint struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `locationName:"containerPath" type:"string"`

	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume. The default
	// value is false.
	ReadOnly *bool `locationName:"readOnly" type:"boolean"`

	// The name of the volume to mount.
	SourceVolume *string `locationName:"sourceVolume" type:"string"`

	metadataMountPoint `json:"-", xml:"-"`
}

type metadataMountPoint struct {
	SDKShapeTraits bool `type:"structure"`
}

type NetworkBinding struct {
	// The IP address that the container is bound to on the container instance.
	BindIP *string `locationName:"bindIP" type:"string"`

	// The port number on the container that is be used with the network binding.
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The port number on the host that is used with the network binding.
	HostPort *int64 `locationName:"hostPort" type:"integer"`

	// The protocol used for the network binding. The default is TCP.
	Protocol *string `locationName:"protocol" type:"string"`

	metadataNetworkBinding `json:"-", xml:"-"`
}

type metadataNetworkBinding struct {
	SDKShapeTraits bool `type:"structure"`
}

// Port mappings allow containers to access ports on the host container instance
// to send or receive traffic. Port mappings are specified as part of the container
// definition.
type PortMapping struct {
	// The port number on the container that is bound to the user-specified or automatically
	// assigned host port. If you specify a container port and not a host port,
	// your container will automatically receive a host port in the 49153 to 65535
	// port range.
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The port number on the container instance to reserve for your container.
	// You can specify a non-reserved host port for your container port mapping,
	// or you can omit the hostPort (or set it to 0) while specifying a containerPort
	// and your container will automatically receive a port in the 49153 to 65535
	// port range. You should not attempt to specify a host port in the 49153 to
	// 65535 port range, since these are reserved for automatic assignment.
	//
	// The default reserved ports are 22 for SSH, the Docker ports 2375 and 2376,
	// and the Amazon ECS Container Agent port 51678. Any host port that was previously
	// specified in a running task is also reserved while the task is running (once
	// a task stops, the host port is released).The current reserved ports are displayed
	// in the remainingResources of DescribeContainerInstances output, and a container
	// instance may have up to 50 reserved ports at a time, including the default
	// reserved ports (automatically assigned ports do not count toward this limit).
	HostPort *int64 `locationName:"hostPort" type:"integer"`

	// The protocol used for the port mapping. The default is TCP.
	Protocol *string `locationName:"protocol" type:"string"`

	metadataPortMapping `json:"-", xml:"-"`
}

type metadataPortMapping struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterContainerInstanceInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to register your container instance with. If you do not specify a cluster,
	// the default cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	InstanceIdentityDocument *string `locationName:"instanceIdentityDocument" type:"string"`

	InstanceIdentityDocumentSignature *string `locationName:"instanceIdentityDocumentSignature" type:"string"`

	TotalResources []*Resource `locationName:"totalResources" type:"list"`

	VersionInfo *VersionInfo `locationName:"versionInfo" type:"structure"`

	metadataRegisterContainerInstanceInput `json:"-", xml:"-"`
}

type metadataRegisterContainerInstanceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterContainerInstanceOutput struct {
	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`

	metadataRegisterContainerInstanceOutput `json:"-", xml:"-"`
}

type metadataRegisterContainerInstanceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterTaskDefinitionInput struct {
	// A list of container definitions in JSON format that describe the different
	// containers that make up your task.
	ContainerDefinitions []*ContainerDefinition `locationName:"containerDefinitions" type:"list" required:"true"`

	// You must specify a family for a task definition, which allows you to track
	// multiple versions of the same task definition. You can think of the family
	// as a name for your task definition. Up to 255 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	Family *string `locationName:"family" type:"string" required:"true"`

	// A list of volume definitions in JSON format that containers in your task
	// may use.
	Volumes []*Volume `locationName:"volumes" type:"list"`

	metadataRegisterTaskDefinitionInput `json:"-", xml:"-"`
}

type metadataRegisterTaskDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterTaskDefinitionOutput struct {
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`

	metadataRegisterTaskDefinitionOutput `json:"-", xml:"-"`
}

type metadataRegisterTaskDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes the resources available for a container instance.
type Resource struct {
	// When the doubleValue type is set, the value of the resource must be a double
	// precision floating-point type.
	DoubleValue *float64 `locationName:"doubleValue" type:"double"`

	// When the integerValue type is set, the value of the resource must be an integer.
	IntegerValue *int64 `locationName:"integerValue" type:"integer"`

	// When the longValue type is set, the value of the resource must be an extended
	// precision floating-point type.
	LongValue *int64 `locationName:"longValue" type:"long"`

	// The name of the resource, such as CPU, MEMORY, PORTS, or a user-defined resource.
	Name *string `locationName:"name" type:"string"`

	// When the stringSetValue type is set, the value of the resource must be a
	// string type.
	StringSetValue []*string `locationName:"stringSetValue" type:"list"`

	// The type of the resource, such as INTEGER, DOUBLE, LONG, or STRINGSET.
	Type *string `locationName:"type" type:"string"`

	metadataResource `json:"-", xml:"-"`
}

type metadataResource struct {
	SDKShapeTraits bool `type:"structure"`
}

type RunTaskInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to run your task on. If you do not specify a cluster, the default cluster
	// is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the specified task that you would like to
	// place on your cluster.
	//
	//  The count parameter is limited to 10 tasks per call.
	Count *int64 `locationName:"count" type:"integer"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the command it should run instead of
	// its default. A total of 8192 characters are allowed for overrides. This limit
	// includes the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// An optional tag specified when a task is started. For example if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataRunTaskInput `json:"-", xml:"-"`
}

type metadataRunTaskInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RunTaskOutput struct {
	// Any failed tasks from your RunTask action are listed here.
	Failures []*Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were run. Each task that was successfully
	// placed on your cluster will be described here.
	Tasks []*Task `locationName:"tasks" type:"list"`

	metadataRunTaskOutput `json:"-", xml:"-"`
}

type metadataRunTaskOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Service struct {
	// The Amazon Resource Name (ARN) of the of the cluster that hosts the service.
	ClusterARN *string `locationName:"clusterArn" type:"string"`

	// The current state of deployments for the service.
	Deployments []*Deployment `locationName:"deployments" type:"list"`

	// The desired number of instantiations of the task definition to keep running
	// on the service. This value is specified when the service is created with
	// CreateService, and it can be modified with UpdateService.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The event stream for your service. A maximum of 100 of the latest events
	// are displayed.
	Events []*ServiceEvent `locationName:"events" type:"list"`

	// A list of load balancer objects, containing the load balancer name, the container
	// name (as it appears in a container definition), and the container port to
	// access from the load balancer.
	LoadBalancers []*LoadBalancer `locationName:"loadBalancers" type:"list"`

	// The number of tasks in the cluster that are in the PENDING state.
	PendingCount *int64 `locationName:"pendingCount" type:"integer"`

	// The Amazon Resource Name (ARN) of the IAM role associated with the service
	// that allows the Amazon ECS container agent to register container instances
	// with a load balancer.
	RoleARN *string `locationName:"roleArn" type:"string"`

	// The number of tasks in the cluster that are in the RUNNING state.
	RunningCount *int64 `locationName:"runningCount" type:"integer"`

	// The Amazon Resource Name (ARN) that identifies the service. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the service, the AWS
	// account ID of the service owner, the service namespace, and then the service
	// name. For example, arn:aws:ecs:region:012345678910:service/my-service.
	ServiceARN *string `locationName:"serviceArn" type:"string"`

	// A user-generated string that you can use to identify your service.
	ServiceName *string `locationName:"serviceName" type:"string"`

	// The status of the service. The valid values are ACTIVE, DRAINING, or INACTIVE.
	Status *string `locationName:"status" type:"string"`

	// The task definition to use for tasks in the service. This value is specified
	// when the service is created with CreateService, and it can be modified with
	// UpdateService.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`

	metadataService `json:"-", xml:"-"`
}

type metadataService struct {
	SDKShapeTraits bool `type:"structure"`
}

type ServiceEvent struct {
	// The Unix time in seconds and milliseconds when the event was triggered.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The ID string of the event.
	ID *string `locationName:"id" type:"string"`

	// The event message.
	Message *string `locationName:"message" type:"string"`

	metadataServiceEvent `json:"-", xml:"-"`
}

type metadataServiceEvent struct {
	SDKShapeTraits bool `type:"structure"`
}

type StartTaskInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to start your task on. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUIDs or full Amazon Resource Name (ARN) entries for
	// the container instances on which you would like to place your task.
	//
	//  The list of container instances to start tasks on is limited to 10.
	ContainerInstances []*string `locationName:"containerInstances" type:"list" required:"true"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the command it should run instead of
	// its default. A total of 8192 characters are allowed for overrides. This limit
	// includes the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// An optional tag specified when a task is started. For example if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to start.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataStartTaskInput `json:"-", xml:"-"`
}

type metadataStartTaskInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type StartTaskOutput struct {
	// Any failed tasks from your StartTask action are listed here.
	Failures []*Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were started. Each task that was successfully
	// placed on your container instances will be described here.
	Tasks []*Task `locationName:"tasks" type:"list"`

	metadataStartTaskOutput `json:"-", xml:"-"`
}

type metadataStartTaskOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type StopTaskInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task you want to stop. If you do not specify a cluster, the default cluster
	// is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The task UUIDs or full Amazon Resource Name (ARN) entry of the task you would
	// like to stop.
	Task *string `locationName:"task" type:"string" required:"true"`

	metadataStopTaskInput `json:"-", xml:"-"`
}

type metadataStopTaskInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type StopTaskOutput struct {
	Task *Task `locationName:"task" type:"structure"`

	metadataStopTaskOutput `json:"-", xml:"-"`
}

type metadataStopTaskOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SubmitContainerStateChangeInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container.
	Cluster *string `locationName:"cluster" type:"string"`

	// The name of the container.
	ContainerName *string `locationName:"containerName" type:"string"`

	// The exit code returned for the state change request.
	ExitCode *int64 `locationName:"exitCode" type:"integer"`

	// The network bindings of the container.
	NetworkBindings []*NetworkBinding `locationName:"networkBindings" type:"list"`

	// The reason for the state change request.
	Reason *string `locationName:"reason" type:"string"`

	// The status of the state change request.
	Status *string `locationName:"status" type:"string"`

	// The task UUID or full Amazon Resource Name (ARN) of the task that hosts the
	// container.
	Task *string `locationName:"task" type:"string"`

	metadataSubmitContainerStateChangeInput `json:"-", xml:"-"`
}

type metadataSubmitContainerStateChangeInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SubmitContainerStateChangeOutput struct {
	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`

	metadataSubmitContainerStateChangeOutput `json:"-", xml:"-"`
}

type metadataSubmitContainerStateChangeOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SubmitTaskStateChangeInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task.
	Cluster *string `locationName:"cluster" type:"string"`

	// The reason for the state change request.
	Reason *string `locationName:"reason" type:"string"`

	// The status of the state change request.
	Status *string `locationName:"status" type:"string"`

	// The task UUID or full Amazon Resource Name (ARN) of the task in the state
	// change request.
	Task *string `locationName:"task" type:"string"`

	metadataSubmitTaskStateChangeInput `json:"-", xml:"-"`
}

type metadataSubmitTaskStateChangeInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SubmitTaskStateChangeOutput struct {
	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`

	metadataSubmitTaskStateChangeOutput `json:"-", xml:"-"`
}

type metadataSubmitTaskStateChangeOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Task struct {
	// The Amazon Resource Name (ARN) of the of the cluster that hosts the task.
	ClusterARN *string `locationName:"clusterArn" type:"string"`

	// The Amazon Resource Name (ARN) of the container instances that host the task.
	ContainerInstanceARN *string `locationName:"containerInstanceArn" type:"string"`

	// The containers associated with the task.
	Containers []*Container `locationName:"containers" type:"list"`

	// The desired status of the task.
	DesiredStatus *string `locationName:"desiredStatus" type:"string"`

	// The last known status of the task.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// One or more container overrides.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// The tag specified when a task is started. If the task is started by an Amazon
	// ECS service, then the startedBy parameter contains the deployment ID of the
	// service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The Amazon Resource Name (ARN) of the task.
	TaskARN *string `locationName:"taskArn" type:"string"`

	// The Amazon Resource Name (ARN) of the of the task definition that creates
	// the task.
	TaskDefinitionARN *string `locationName:"taskDefinitionArn" type:"string"`

	metadataTask `json:"-", xml:"-"`
}

type metadataTask struct {
	SDKShapeTraits bool `type:"structure"`
}

type TaskDefinition struct {
	// A list of container definitions in JSON format that describe the different
	// containers that make up your task. For more information on container definition
	// parameters and defaults, see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	ContainerDefinitions []*ContainerDefinition `locationName:"containerDefinitions" type:"list"`

	// The family of your task definition. You can think of the family as the name
	// of your task definition.
	Family *string `locationName:"family" type:"string"`

	// The revision of the task in a particular family. You can think of the revision
	// as a version number of a task definition in a family. When you register a
	// task definition for the first time, the revision is 1, and each time you
	// register a task definition in the same family, the revision value increases
	// by one.
	Revision *int64 `locationName:"revision" type:"integer"`

	// The full Amazon Resource Name (ARN) of the of the task definition.
	TaskDefinitionARN *string `locationName:"taskDefinitionArn" type:"string"`

	// The list of volumes in a task. For more information on volume definition
	// parameters and defaults, see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	Volumes []*Volume `locationName:"volumes" type:"list"`

	metadataTaskDefinition `json:"-", xml:"-"`
}

type metadataTaskDefinition struct {
	SDKShapeTraits bool `type:"structure"`
}

// A list of container overrides in JSON format that specify the name of a container
// in a task definition and the command it should run instead of its default.
type TaskOverride struct {
	// One or more container overrides sent to a task.
	ContainerOverrides []*ContainerOverride `locationName:"containerOverrides" type:"list"`

	metadataTaskOverride `json:"-", xml:"-"`
}

type metadataTaskOverride struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateContainerAgentInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// container instance is running on. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) entries for
	// the container instance on which you would like to update the Amazon ECS container
	// agent.
	ContainerInstance *string `locationName:"containerInstance" type:"string" required:"true"`

	metadataUpdateContainerAgentInput `json:"-", xml:"-"`
}

type metadataUpdateContainerAgentInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateContainerAgentOutput struct {
	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`

	metadataUpdateContainerAgentOutput `json:"-", xml:"-"`
}

type metadataUpdateContainerAgentOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateServiceInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// service is running on. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the task that you would like to place and
	// keep running in your service.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The name of the service that you want to update.
	Service *string `locationName:"service" type:"string" required:"true"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run in your service. If you modify
	// the task definition with UpdateService, Amazon ECS spawns a task with the
	// new version of the task definition and then stops an old task after the new
	// version is running.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`

	metadataUpdateServiceInput `json:"-", xml:"-"`
}

type metadataUpdateServiceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateServiceOutput struct {
	// The full description of your service following the update call.
	Service *Service `locationName:"service" type:"structure"`

	metadataUpdateServiceOutput `json:"-", xml:"-"`
}

type metadataUpdateServiceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VersionInfo struct {
	// The Git commit hash for the Amazon ECS container agent build on the amazon-ecs-agent
	//  (https://github.com/aws/amazon-ecs-agent/commits/master) GitHub repository.
	AgentHash *string `locationName:"agentHash" type:"string"`

	// The version number of the Amazon ECS container agent.
	AgentVersion *string `locationName:"agentVersion" type:"string"`

	// The Docker version running on the container instance.
	DockerVersion *string `locationName:"dockerVersion" type:"string"`

	metadataVersionInfo `json:"-", xml:"-"`
}

type metadataVersionInfo struct {
	SDKShapeTraits bool `type:"structure"`
}

type Volume struct {
	// The path on the host container instance that is presented to the containers
	// which access the volume. If this parameter is empty, then the Docker daemon
	// assigns a host path for you.
	Host *HostVolumeProperties `locationName:"host" type:"structure"`

	// The name of the volume. This name is referenced in the sourceVolume parameter
	// of container definition mountPoints.
	Name *string `locationName:"name" type:"string"`

	metadataVolume `json:"-", xml:"-"`
}

type metadataVolume struct {
	SDKShapeTraits bool `type:"structure"`
}

type VolumeFrom struct {
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume. The default
	// value is false.
	ReadOnly *bool `locationName:"readOnly" type:"boolean"`

	// The name of the container to mount volumes from.
	SourceContainer *string `locationName:"sourceContainer" type:"string"`

	metadataVolumeFrom `json:"-", xml:"-"`
}

type metadataVolumeFrom struct {
	SDKShapeTraits bool `type:"structure"`
}
