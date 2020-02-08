# aws
## [1][Worst AWS consultant experiences?](https://www.reddit.com/r/aws/comments/f0f2yx/worst_aws_consultant_experiences/)
- url: https://www.reddit.com/r/aws/comments/f0f2yx/worst_aws_consultant_experiences/
---
People that have engaged third party AWS consultants, what did they do badly? What surprised you most?
## [2][Is there a way to avoid duplicates in Dynamodb?](https://www.reddit.com/r/aws/comments/f0q036/is_there_a_way_to_avoid_duplicates_in_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/f0q036/is_there_a_way_to_avoid_duplicates_in_dynamodb/
---
I'm adding records with auto-generated primary keys. So, is there a way to avoid duplicate records in Dynamodb?
## [3][Why would I use Node.js in Lambda? Node main feature is handling concurrent many requests. If each request to lambda will spawn a new Node instance, whats the point?](https://www.reddit.com/r/aws/comments/f0dsq3/why_would_i_use_nodejs_in_lambda_node_main/)
- url: https://www.reddit.com/r/aws/comments/f0dsq3/why_would_i_use_nodejs_in_lambda_node_main/
---
Maybe I'm missing something here, from an architectural point of view, I can't wrap my head on using node inside a lambda. Let's say I receive 3 requests, a single node instance would be able to handle this with ease, but if I use lambda, 3 lambdas with Node inside would be spawned, each would be idle while waiting for the callback.

Edit: Many very good answers. I will for sure discuss this with the team next week. Very happy with this community. Thanks and please keep them coming!
## [4][What tools are useful for testing api response speeds.?](https://www.reddit.com/r/aws/comments/f0rfkr/what_tools_are_useful_for_testing_api_response/)
- url: https://www.reddit.com/r/aws/comments/f0rfkr/what_tools_are_useful_for_testing_api_response/
---
I've setup few apis in api gateway and connected to lambda and Dynamodb. After deploying these apis. I want to be able to test them in rapid succession to measure the response time and latency and data transfered.
## [5][Database in RDS keeps disappearing. Please Help](https://www.reddit.com/r/aws/comments/f0n5ua/database_in_rds_keeps_disappearing_please_help/)
- url: https://www.reddit.com/r/aws/comments/f0n5ua/database_in_rds_keeps_disappearing_please_help/
---
I have an amazon student AWS account and for my project I am using the RDS with MSQL (basic config). I keep finding my database missing every 2-3 days and I have to keep building up. The instance still remains active, but it is the database that keeps disappearing. Any idea what might be causing it and how to solve it .
## [6][Question about the most simple and secure way to deploy a React Website and a Django Framework](https://www.reddit.com/r/aws/comments/f0r04p/question_about_the_most_simple_and_secure_way_to/)
- url: https://www.reddit.com/r/aws/comments/f0r04p/question_about_the_most_simple_and_secure_way_to/
---
Hello everyone,

I've developed a React JS where the users can retrieve data from my Django REST backend to query data etc.  Optinally: I want to create a login functionality.My initial thought was to create a Kubernetes cluster and deploy there the services but I now think that it might be an overkill.

I did some research and found the AWS API Gateway.

My new idea is to create an API Gateway and connect this to my Website that I host on S3. Then I route the GET request from the website through the API Gateway to my Django REST Service and return there the data. I want to use the Django Backend because I think about to implement a Sign up and Login functionality to my website. Otherwise I also could use Lambdas.

The architecture might look like this:

Client &lt;----&gt; Website (public) &lt;----&gt; GET &lt;----&gt; API Gateway (public) &lt;------&gt; Django REST service (inside private vpc) &lt;-----&gt; AWS RDS (also in the private vpc, contains my MySQL db).

The REST services is hosted on an EC2 instance and retrieve data from a MySQL db that are both in a private vpc. So the only public visible address would be the API Gateway.

What do you think about this approach?

Edit: For me it seems to be way too much overhead to have the API Gateway and also a Django REST service.
## [7][ELI5: What is capacity units?](https://www.reddit.com/r/aws/comments/f0qblg/eli5_what_is_capacity_units/)
- url: https://www.reddit.com/r/aws/comments/f0qblg/eli5_what_is_capacity_units/
---
What is read and write capacity units? I've been using DynamoDB and this concept has eluded me the most. 

Is this related to concurrency? Like, if Read and Write is set to 1 unit, does this mean only one lambda/ one request can be reading and writing to the DB?
## [8][Automatic Capacity Reservations](https://www.reddit.com/r/aws/comments/f099cu/automatic_capacity_reservations/)
- url: https://www.reddit.com/r/aws/comments/f099cu/automatic_capacity_reservations/
---
Hi all 👋,

Today, AWS released a great new [update](https://aws.amazon.com/about-aws/whats-new/2020/02/amazon-ec2-adds-ability-to-easily-query-billing-information-of-amazon-machine-images/) which now allows us to retrieve the full platform type for running instances (indirectly via DescribeImages). With this, I was able to complete a small tool which automatically provisions capacity reservations to meet the accounts instance count.

This tool would be helpful for those that maintain a fleet of instances that are regularly stopped/started and need guaranteed capacity to be able to start on demand.

It's available as a deployable CloudFormation stack at: [https://github.com/iann0036/auto-capacity-reservations](https://github.com/iann0036/auto-capacity-reservations)

Comments, suggestions and general discussion welcomed!
## [9][Dockerization of NodeJS Applications on Amazon Elastic Containers](https://www.reddit.com/r/aws/comments/f0ie1a/dockerization_of_nodejs_applications_on_amazon/)
- url: https://blog.soshace.com/dockerization-of-node-js-applications-on-amazon-elastic-containers/
---

## [10][What to look out for when moving from EC2 to RDS for Postgres?](https://www.reddit.com/r/aws/comments/f0iz0f/what_to_look_out_for_when_moving_from_ec2_to_rds/)
- url: https://www.reddit.com/r/aws/comments/f0iz0f/what_to_look_out_for_when_moving_from_ec2_to_rds/
---
Hi Everyone,

I'm seeking out some first hand knowledge from anyone that has migrated from running your own Postgres instances on EC2 to RDS. We're excited by a number of features that will allow us to free up DBA time to tuning vs. configuration management, as well as the new RDS Proxy functionality coming out in the near future that will support Postgres. We're currently running a fairly large instance type to support the increasing number of connections hitting our instance, but I'm hoping we can scale that back when we eventually implement RDS Proxy, which should hopefully offset the cost of moving to RDS. I know we could use something like pgpool for that area specifically, but I'm a big fan of managed services where available.

In years past, we were looking to make the move from our EC2 MS SQL Server instances to RDS, but there were too many specific use cases that prevented us from doing that properly. Is anyone aware of limitations or "gotchas" when migrating from EC2 to Postgres RDS specifically? They have a great section on their MS SQL Server docs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html) in the SQL Server Security section, but I can't seem to find anything comparable for postgres.

I'm the farthest thing from a Postgres DBA, so all of this information will be gathered and reviewed collectively with our actual DBAs, but they're not too familiar with the AWS side of things (yet).

Thanks!
