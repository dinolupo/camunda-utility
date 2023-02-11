package client

// ProcessDefinition a client for ProcessDefinition
type ProcessInstance struct {
	Client *Client
}

// ResProcessInstance a response object for process instance
type ResProcessInstance struct {
	// The id of the process instance.
	Id string `json:"id"`
	// The id of the process definition that this process instance belongs to.
	DefinitionId string `json:"definitionId"`
	// The business key of the process instance.
	BusinessKey string `json:"businessKey"`
	// The id of the case instance associated with the process instance.
	CaseInstanceId string `json:"caseInstanceId"`
	// A flag indicating whether the process instance is suspended or not.
	Suspended bool `json:"suspended"`
	// The tenant id of the process instance.
	TenantId string `json:"tenantId"`
	// A JSON array containing links to interact with the instance
	//Links []ResLink `json:"links"`
}

// Delete deletes a running process instance by id.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/delete/#query-parameters
func (p *ProcessInstance) Delete(id string, query map[string]string) error {
	_, err := p.Client.doDelete("/process-instance/"+id, query)
	return err
}

// GetList queries for process instances that fulfill given parameters.
// Parameters may be static as well as dynamic runtime properties of process instances.
// The size of the result set can be retrieved by using the GetCount method.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/get-query/#query-parameters
func (p *ProcessInstance) GetList(query map[string]string) (processInstances []*ResProcessInstance, err error) {
	res, err := p.Client.doGet("/process-instance", query)
	if err != nil {
		return
	}

	err = p.Client.readJsonResponse(res, &processInstances)
	return
}