package client

// ProcessDefinition a client for ProcessDefinition
type ProcessDefinition struct {
	Client *Client
}

// ResProcessDefinition a JSON object corresponding to the ProcessDefinition interface in the engine
type ResProcessDefinition struct {
	// The id of the process definition
	Id string `json:"id"`
	// The key of the process definition, i.e., the id of the BPMN 2.0 XML process definition
	Key string `json:"key"`
	// The category of the process definition
	Category string `json:"category"`
	// The description of the process definition
	Description string `json:"description"`
	// The name of the process definition
	Name string `json:"name"`
	// The version of the process definition that the engine assigned to it
	Version int `json:"Version"`
	// The file name of the process definition
	Resource string `json:"resource"`
	// The deployment id of the process definition
	DeploymentId string `json:"deploymentId"`
	// The file name of the process definition diagram, if it exists
	Diagram string `json:"diagram"`
	// A flag indicating whether the definition is suspended or not
	Suspended bool `json:"suspended"`
	// The tenant id of the process definition
	TenantId string `json:"tenantId"`
	// The version tag of the process definition
	VersionTag string `json:"versionTag"`
	// History time to live value of the process definition. Is used within History cleanup
	HistoryTimeToLive int `json:"historyTimeToLive"`
	// A flag indicating whether the process definition is startable in Tasklist or not
	StartableInTasklist bool `json:"startableInTasklist"`
}

// GetList queries for process definitions that fulfill given parameters.
// Parameters may be the properties of process definitions, such as the name, key or version.
// The size of the result set can be retrieved by using the Get Definition Count method
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/get-query/#query-parameters
func (p *ProcessDefinition) GetList(query map[string]string) (processDefinitions []*ResProcessDefinition, err error) {
	res, err := p.Client.doGet("/process-definition", query)
	if err != nil {
		return
	}

	err = p.Client.readJsonResponse(res, &processDefinitions)
	return
}

// QueryProcessDefinitionBy path builder
type QueryProcessDefinitionBy struct {
	Id       *string
	Key      *string
	TenantId *string
}

// String a build path part
func (q *QueryProcessDefinitionBy) String() string {
	if q.Key != nil && q.TenantId != nil {
		return "key/" + *q.Key + "/tenant-id/" + *q.TenantId
	} else if q.Key != nil {
		return "key/" + *q.Key
	}

	return *q.Id
}

// Delete deletes a process definition from a deployment by id
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/delete-process-definition/#query-parameters
func (p *ProcessDefinition) Delete(by QueryProcessDefinitionBy, query map[string]string) error {
	_, err := p.Client.doDelete("/process-definition/"+by.String(), query)
	return err
}
