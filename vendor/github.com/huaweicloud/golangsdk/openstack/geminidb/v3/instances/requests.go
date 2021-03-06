package instances

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type FlavorOpt struct {
	Num      string `json:"num" required:"true"`
	Size     int    `json:"size" required:"true"`
	Storage  string `json:"storage" required:"true"`
	SpecCode string `json:"spec_code" required:"true"`
}

type BackupStrategyOpt struct {
	StartTime string `json:"start_time,omitempty"`
	KeepDays  string `json:"keep_days,omitempty"`
}

type CreateGeminiDBOpts struct {
	Name                string             `json:"name" required:"true"`
	Region              string             `json:"region" required:"true"`
	AvailabilityZone    string             `json:"availability_zone" required:"true"`
	VpcId               string             `json:"vpc_id" required:"true"`
	SubnetId            string             `json:"subnet_id" required:"true"`
	SecurityGroupId     string             `json:"security_group_id,omitempty"`
	Password            string             `json:"password" required:"true"`
	Mode                string             `json:"mode" required:"true"`
	ConfigurationId     string             `json:"configuration_id,omitempty"`
	EnterpriseProjectId string             `json:"enterprise_project_id,omitempty"`
	Ssl                 string             `json:"ssl_option,omitempty"`
	DataStore           DataStore          `json:"datastore" required:"true"`
	Flavor              []FlavorOpt        `json:"flavor" required:"true"`
	BackupStrategy      *BackupStrategyOpt `json:"backup_strategy,omitempty"`
}

type CreateGeminiDBBuilder interface {
	ToInstancesCreateMap() (map[string]interface{}, error)
}

func (opts CreateGeminiDBOpts) ToInstancesCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateGeminiDBBuilder) (r CreateResult) {
	b, err := opts.ToInstancesCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, instanceId string) (r DeleteResult) {
	url := deleteURL(client, instanceId)

	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})

	return
}

type ListGeminiDBInstanceOpts struct {
	Id            string `q:"id"`
	Name          string `q:"name"`
	Mode          string `q:"mode"`
	DataStoreType string `q:"datastore_type"`
	VpcId         string `q:"vpc_id"`
	SubnetId      string `q:"subnet_id"`
	Offset        int    `q:"offset"`
	Limit         int    `q:"limit"`
}

type ListGeminiDBBuilder interface {
	ToGeminiDBListDetailQuery() (string, error)
}

func (opts ListGeminiDBInstanceOpts) ToGeminiDBListDetailQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListGeminiDBBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToGeminiDBListDetailQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return GeminiDBPage{pagination.SinglePageBase(r)}
	})
	// Headers supplies additional HTTP headers to populate on each paged request
	pageList.Headers = map[string]string{"Content-Type": "application/json"}

	return pageList
}

func GetInstanceByID(client *golangsdk.ServiceClient, instanceId string) (GeminiDBInstance, error) {
	var instance GeminiDBInstance

	opts := ListGeminiDBInstanceOpts{
		Id: instanceId,
	}

	pages, err := List(client, &opts).AllPages()
	if err != nil {
		return instance, err
	}

	all, err := ExtractGeminiDBInstances(pages)
	if err != nil {
		return instance, err
	}
	if all.TotalCount == 0 {
		return instance, nil
	}

	instance = all.Instances[0]
	return instance, nil
}
