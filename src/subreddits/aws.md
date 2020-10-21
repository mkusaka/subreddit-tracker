# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][Week of Oct 19th - What have you learned recently about AWS?](https://www.reddit.com/r/aws/comments/je2wmx/week_of_oct_19th_what_have_you_learned_recently/)
- url: https://www.reddit.com/r/aws/comments/je2wmx/week_of_oct_19th_what_have_you_learned_recently/
---
Weekly Discussion post   
Sharing is caring
## [3][A Kubernetes operator to sync secrets from AWS Secrets Manager](https://www.reddit.com/r/aws/comments/jfa4lz/a_kubernetes_operator_to_sync_secrets_from_aws/)
- url: https://www.contentful.com/blog/2020/10/20/open-source-kube-secret-syncer/
---

## [4][New – Use AWS PrivateLink to Access AWS Lambda Over Private AWS Network](https://www.reddit.com/r/aws/comments/jexzcu/new_use_aws_privatelink_to_access_aws_lambda_over/)
- url: https://aws.amazon.com/blogs/aws/new-use-aws-privatelink-to-access-aws-lambda-over-private-aws-network/
---

## [5][A Cost-Benefit Analysis of VPC Interface Endpoints](https://www.reddit.com/r/aws/comments/jfclgq/a_costbenefit_analysis_of_vpc_interface_endpoints/)
- url: https://www.sentiatechblog.com/a-cost-benefit-analysis-of-vpc-interface-endpoints?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=cost_benefit_privatelink
---

## [6][how to enable cors in AWS api gateway and only allow access from specific websites/origins for GET Requests?](https://www.reddit.com/r/aws/comments/jfa9g3/how_to_enable_cors_in_aws_api_gateway_and_only/)
- url: https://www.reddit.com/r/aws/comments/jfa9g3/how_to_enable_cors_in_aws_api_gateway_and_only/
---
Hey guys I enabled CORS for get request and I only want the system to allow requests that come in from the following domain:  "toi.lightning.force.com"

&amp;#x200B;

As seen in the screenshot, during the response, I adjust the access-control-allow-origin.

&amp;#x200B;

What I am finding, is that the GET request still accepts from any place I access the link. Pretty much I was hoping that CORS would block people from accessing the API if they weren't coming from this specific domain.

&amp;#x200B;

Let me know if I am missing something!

https://preview.redd.it/wq51kq4ybfu51.png?width=1964&amp;format=png&amp;auto=webp&amp;s=71c482521579a223a16c6e64a9868f080c0073bb

&amp;#x200B;

&amp;#x200B;

&amp;#x200B;

Thank you in advance.

&amp;#x200B;

&amp;#x200B;

EDIT:

**CORS does not do server-side blocking of request**
## [7][Need Ideas For Survey Website](https://www.reddit.com/r/aws/comments/jf7qar/need_ideas_for_survey_website/)
- url: https://www.reddit.com/r/aws/comments/jf7qar/need_ideas_for_survey_website/
---
Hello, My fiancé needs a website with the ability to store answers to 3 questions, It's for a class and will only be up for a month. Students would simply click on a link, see 3 questions, each with a text box that would allow them to type a 100 char response. 

Since I already have the domain on Route 53 + S3, I figured I might as well implement a full AWS solution. Any one have an idea on how I could do this? 

My idea was to host a Windows server, but then I saw that price tag...

PS. I do realize Survey monkey is a solution but I've been interested in the possibilities of AWS. Figured this is the best way to get my feet wet.
## [8][Does the AWSCLI S3 SYNC or COPY command copy data to the local machine?](https://www.reddit.com/r/aws/comments/jfctg0/does_the_awscli_s3_sync_or_copy_command_copy_data/)
- url: https://www.reddit.com/r/aws/comments/jfctg0/does_the_awscli_s3_sync_or_copy_command_copy_data/
---
I am trying to copy objects from an S3 bucket in one region to a bucket in another region and I came across the aws s3 sync and cp commands. Do these commands copy the data to my local machine or is it completely done on aws? I would prefer that this is done without copying the data to my local machine and I was trying to find more details about this in the documentation, but did not find this mentioned anywhere.
## [9][How to change the handles in SES template data dynamically?](https://www.reddit.com/r/aws/comments/jfcrai/how_to_change_the_handles_in_ses_template_data/)
- url: https://www.reddit.com/r/aws/comments/jfcrai/how_to_change_the_handles_in_ses_template_data/
---
I was trying to send mail via Lambda using SES, and I created a Template, with two handles

&amp;#x200B;

[{{name}}, name of the owner of an instance and the instance status {{status}}](https://preview.redd.it/sjnpcujf8gu51.png?width=943&amp;format=png&amp;auto=webp&amp;s=ad6b5eaabded6c0c1beed4fa1c3d7c8dc2a96a5d)

I'm getting these two values via a Boto call and I need to pass those values to the template data.

&amp;#x200B;

[This is the block of code where I'm sending the template](https://preview.redd.it/5gfem4sx8gu51.png?width=850&amp;format=png&amp;auto=webp&amp;s=aed5b60e55e9b9c963128a361ab75b2dd199e9e3)

I'm getting the below error - 

"Unable to parse template data (invalid JSON)."

&amp;#x200B;

Any idea how to solve this?
## [10][How to switch role to an organization member account from root account?](https://www.reddit.com/r/aws/comments/jfcman/how_to_switch_role_to_an_organization_member/)
- url: https://www.reddit.com/r/aws/comments/jfcman/how_to_switch_role_to_an_organization_member/
---
Any help is highly appreciated, 
One of the employees left and he had a member account in our organization, 
- He created some resources before and now we can't access his account to delete it and 
- He didn't complete all fields required in his account so we can't delete his account either.


I tried the following from AWS docs with no luck.
- Create IAM user with full administration policy
- Create a role of type user and enter the employee ID in it and give it a specific name
- Create a policy with service STS (assume) and give it the employee ID and role created from previous step
- Create a group and attach that policy to it, then add the IAM user to this group

Expected behavior: 
- Choose switch role from the upright account menu
- enter employee account ID and role name, I should be granted access to his account
But in this step it fails while stating one of the two fields are not correct
## [11][Aws training](https://www.reddit.com/r/aws/comments/jfcaun/aws_training/)
- url: https://www.reddit.com/r/aws/comments/jfcaun/aws_training/
---
I wrote whole code locally, data is 5giga, I want to train it online since i got no nvidia, what is proper way of doining, i want to run scripts notebook
## [12][Server Side Blocking using API Gateway and Authrizer](https://www.reddit.com/r/aws/comments/jfax5i/server_side_blocking_using_api_gateway_and/)
- url: https://www.reddit.com/r/aws/comments/jfax5i/server_side_blocking_using_api_gateway_and/
---
I am looking to block requests on the server-side using API Gateway + Authorizer (lambda function). 

I pretty much want to grab the HTTP Origin and say "Yup, this is allowed" or "No, I do not know this origin"

&amp;#x200B;

Origin grabs something like "example.com"

&amp;#x200B;

I know you can put this in your mapping template and it will grab the associated header information:

    { 
    "origin" : "$input.params('origin')", 
    "referer" : "$input.params('referer')" 
    }
    

Unfortunately, I have no way of passing in body mapping templates or accessing the origin of the request from the Authorizer.

&amp;#x200B;

Below is a screenshot of the authorizer, but the Origin variable isn't taking it from the actual origin of the HTTP request. It sees origin as a custom header. I want to be able to find the actual Origin from the HTTP Header.

&amp;#x200B;

https://preview.redd.it/v0dqqhb0mfu51.png?width=940&amp;format=png&amp;auto=webp&amp;s=8e8a3ec026e353b2a5680068681a4fb52ff922d5

How can I pass in the actual HTTP Header to the authorizer?
