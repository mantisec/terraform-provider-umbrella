I need your help to determine the best way forward to complete the implemention of a terraform-provider.
Im trying to determine if the best approach will be to simply collaborate with you to write the code or whether I should have you conduct research and formulate a detailed implementation plan and directives which I can then use to develop the code with an IDE agent such as Cline or Roo Code.
I have a partially completed terraform-provider which reads openapi spec yaml files and then infers the resources, schemas and endpoints required for the terraform provider. To be more precise the build of this terraform provider runs some transformation logic on the openapi yaml files to generate the terraform resource and data_source code files before they are then compiled into the terraform provider binary.

i am using libopenapi which ensures that we can interpret the openapi document correctly as it uses some complex nested object structures and $ref statements, the libopenapi integration is working well.
I need to help define the rules which will ensure the transformation generates a complete set of fully operational resources and data_sources.

When we think about translating rest api's into a terraform provider, there are some clear synergies that will help us. The CRUD operations of many REST API's align somewhat with the lifecycle operations of terraform resource, however to build a fully feature terraform provider:
- we need to know how to completely define resource schema.
- we need to know the endpoints to perform CRUD operations.
- we need to intelligently consolidate the different crud endpoints and operations when they relate to a single resource type.
- we need to know which attributes are mandatory for each CRUD operation
- we need to know which attributes can be updated and which attributes should force the creation of a new resource.
- in summary, we need to know how to construct and use the appropriate API request (endpoint and payload) for a each given resource and operation.

I have observed some patterns which provide a basis for some of our transformation logic. (I will use a combination of yaml code and inline #comments to highlight important information)
IMPORTANT !!! - the raw yaml file may contain yaml $ref statements, so it is critically important that the yaml is properly parsed before transformation logic is applied.


We will use excerpts from the Cisco Umbrella REST API OpenAPI spec as our example.

standard "top-level" resource example:
```yaml
paths:
  '/tunnels':
    get:
      #list ALL tunnel resources
    post:
      #create a new tunnel resource
      requestBody:
        description: The tunnel to create.
        required: true
        content:
            application/json:
              schema:
                #create tunnel resource schema
```
terraform names based on the example above:
```text
resource_tunnels
data_source_tunnels
```
in the example above the path "/tunnels" is the endpoint for listing and creating Cisco Umbrella Tunnel resources.
IMPORTANT - top-level paths eg. '/tunnels' will normally provide only a get and post endpoint to list ALL resource of that type or create a new resource.
IMPORTANT!!  paths['/tunnels'].post.requestBody.content['application/json'].schema provides the full schema required to create the resource, descriptions here are very useful for terraform provider docs !


The next example will show how we view and manage individual items.
```yaml
paths:
  '/tunnels/{id}':
    get:
      summary: Get Tunnel
      description: # useful for documentation
      parameters:
        - in: path
          name: id
          description: The ID of the tunnel.
          required: true
          schema:
            type: integer
      responses:
        '200':
          content:
            application/json:
              schema:
                #schema definition>
              example:
                #detailed example for documentation>
    put:
      # see next example
    delete:
      # see next example
```

this example above shows resource specific paths, here we can see the prefix '/tunnels/' from our already discovered tunnel resource, however this path includes the {id} parameter which allows us to view or manage a specific tunnel instance. this endpoint contains get, put/patch and delete http verbs.
IMPORTANT!! note here the 'parameters' attribute inside the get object describes how to construct the request to read the target resource.
IMPORTANT!!  get.responses['200'].content['application/json'].schema provides the full schema of the resource, descriptions here are very useful for terraform provider docs !
IMPORTANT!!  get.responses['200'].content['application/json'].example provides excellent additional information for the terraform provider docs !


The next example focuses on the put and delete https verb for the same /tunnel/{id} path as the last example resource
```yaml
paths:
  '/tunnels/{id}':
    get:
      # get a specific tunnel resource
    put:
      summary: Update Tunnel
      description: |-
        Update the `name`, `siteOriginId`, `networkCIDRs`, and client `deviceType` properties for a tunnel.
        Updates to read-only attributes are ignored.
      parameters:
        - in: path
          name: id
          description: The ID of the tunnel.
          required: true
          schema:
            type: integer
      requestBody:
        description: Provide a tunnel to update.
        required: true
        content:
            application/json:
              schema:
                # resource update schema
    delete:
      summary: Delete Tunnel
      description: Delete a tunnel in the organization.
      parameters:
        - in: path
          name: id
          description: The ID of the tunnel.
          required: true
          schema:
            type: integer
```
IMPORTANT!! note here the 'parameters' attribute inside the put and delete objects describe how to construct the request to target the correct resource.
IMPORTANT!!  paths['/tunnels/{id}'].put.requestBody.content['application/json'].schema provides the full schema required to update the resource, descriptions here are very useful for terraform provider docs !


The next example describes how to manage resource "sub-types", In practical terms this is where a given resource may contain nested resources or configuration items.
we can provide a concrete example using the Cisco Umbrella example. A Cisco Umbrella Tunnel can have policies attached.
```yaml
paths:
  /tunnels/{id}/policies:
```
In the specific case of tunnel_policies, there is only a get method, meaning this could only work as a data-source.

IMPORTANT NOTE: The api specs include a number of endpoints which are only used for debugging, These should NOT be included in the terraform provider. All debugging endpoints include a 'tags' list attribute that must contain the value 'Debugging' for example:
```yaml
paths:
  '/tunnelsState':
    get:
      tags:
        - Debugging
```

I will also provide you with links to the complete OpenApi specs and REST API documentation for Cisco Umbrella, please review all the content carefully.

I need your advanced reasoning capabilities to identify the patterns which we can use to translate the OpenAPi specs into a fully functional terraform provider with complient,idempotent operations and detailed documentation.

here is the link to my repo with partially implemented code:
https://github.com/mantisec/terraform-provider-umbrella/tree/feature/0.4.0

here is the link to the root of the Cisco Umbrella API docs:
https://developer.cisco.com/docs/cloud-security/umbrella-api-reference-overview/

here are the links to the full OpenAPI specs for the Umbrella API:
```
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/auth/cisco_umbrella_token_authorization_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_networks_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_internal_networks_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_internal_domains_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_sites_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_virtual_appliances_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_roaming_computers_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_network_tunnels_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_network_devices_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_deployments_policies_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_tagging_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/deployments/cisco_umbrella_secure_web_gateway_device_settings_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/admin/cisco_umbrella_key_admin_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/admin/cisco_umbrella_service_providers_console_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/admin/cisco_umbrella_providers_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/admin/cicso_umbrella_managed_providers_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/admin/cisco_umbrella_users_and_roles_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/admin/cisco_umbrella_s_3_bucket_key_rotation_api_1_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/investigate/cisco_umbrella_investigate_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/policies/cisco_umbrella_destination_lists_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/reports/cisco_umbrella_reporting_api_2_0_0.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/reports/cisco_umbrella_app_discovery_api_2_0_1.yaml
https://pubhub.devnetcloud.com/media/cloud-security-apis-in-eft/docs/reference/reports/cisco_umbrella_api_usage_reports_2_0_0.yaml
```
