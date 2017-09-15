package client

const (
	PUBLISH_TYPE = "publish"
)

type Publish struct {
	Resource

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	PreviousId string `json:"previousId,omitempty" yaml:"previous_id,omitempty"`

	PreviousIds []string `json:"previousIds,omitempty" yaml:"previous_ids,omitempty"`

	ResourceId string `json:"resourceId,omitempty" yaml:"resource_id,omitempty"`

	ResourceType string `json:"resourceType,omitempty" yaml:"resource_type,omitempty"`

	Time int64 `json:"time,omitempty" yaml:"time,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
}

type PublishCollection struct {
	Collection
	Data   []Publish `json:"data,omitempty"`
	client *PublishClient
}

type PublishClient struct {
	rancherClient *RancherClient
}

type PublishOperations interface {
	List(opts *ListOpts) (*PublishCollection, error)
	Create(opts *Publish) (*Publish, error)
	Update(existing *Publish, updates interface{}) (*Publish, error)
	ById(id string) (*Publish, error)
	Delete(container *Publish) error
}

func newPublishClient(rancherClient *RancherClient) *PublishClient {
	return &PublishClient{
		rancherClient: rancherClient,
	}
}

func (c *PublishClient) Create(container *Publish) (*Publish, error) {
	resp := &Publish{}
	err := c.rancherClient.doCreate(PUBLISH_TYPE, container, resp)
	return resp, err
}

func (c *PublishClient) Update(existing *Publish, updates interface{}) (*Publish, error) {
	resp := &Publish{}
	err := c.rancherClient.doUpdate(PUBLISH_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PublishClient) List(opts *ListOpts) (*PublishCollection, error) {
	resp := &PublishCollection{}
	err := c.rancherClient.doList(PUBLISH_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PublishCollection) Next() (*PublishCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PublishCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PublishClient) ById(id string) (*Publish, error) {
	resp := &Publish{}
	err := c.rancherClient.doById(PUBLISH_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *PublishClient) Delete(container *Publish) error {
	return c.rancherClient.doResourceDelete(PUBLISH_TYPE, &container.Resource)
}
