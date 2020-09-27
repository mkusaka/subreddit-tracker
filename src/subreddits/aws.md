# aws
## [1][S3 Path Deprecation Plan Updated](https://www.reddit.com/r/aws/comments/iyd84u/s3_path_deprecation_plan_updated/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/
---

## [2][Custom Servers for Games?](https://www.reddit.com/r/aws/comments/j0qdkf/custom_servers_for_games/)
- url: https://www.reddit.com/r/aws/comments/j0qdkf/custom_servers_for_games/
---
Hi all,

Obviously I'm a noob in AWS but I want to leverage the cloud. I'm planning to setup a custom server for this game (Left 4 Dead 2) to play with my friends and I'm using EC2. When I tried to ping it from my game to do a test, I can't seem to connect to the server I created. TCP and UDP ports have been opened under security groups and I also enabled it from the Windows Defender Firewall. Now my real question, is there anything I may have missed? Or perhaps hosting a custom server in AWS is not possible?
## [3][Cognito: Static S3 website or API gateway/lambda?](https://www.reddit.com/r/aws/comments/j0osrv/cognito_static_s3_website_or_api_gatewaylambda/)
- url: https://www.reddit.com/r/aws/comments/j0osrv/cognito_static_s3_website_or_api_gatewaylambda/
---
Hi all.

I’m working on a personal project, and looking to secure access to an “admin” site that allows authorised users access to certain functionality (sending SMS through Twilio etc). I’ve read a bit around using Cognito to restrict access but I’ve a couple of questions regarding using this with S3 and Lambda.

Currently I’ve a static website in S3 using CloudFront. My initial thought was to create my static webpages, and restrict access to the lambdas they call using Cognito, similar to the “wild rydes” tutorial. However, webpages with restricted functionality are still viewable before the redirect kicks in, although obviously these pages won’t allow unauthenticated users to use the functionality. I’d prefer to not allow these pages to be viewed.

I’m now considering a completely serverless approach. My thought is to create a registration page and login page on the static website, retrieve the JWT and redirect the user to a serverless drive sub domain, admin.website.com. This would be fronted by API gateway calling lambdas to build the pages and run the functionality for the admin section. Initially I’ll pull the webpages in from a private S3 bucket and build the pages myself, though I’d maybe look to move to React for building the customised webpages.

Is there anything I’ve overlooked in this approach? If this is an appropriate solution, do you have any thoughts on how I can improve this approach? I’d appreciate your thoughts on this!

Thanks!
## [4][How to create lambda applications with many functions? (CloudFormaton limit)](https://www.reddit.com/r/aws/comments/j0qki4/how_to_create_lambda_applications_with_many/)
- url: https://www.reddit.com/r/aws/comments/j0qki4/how_to_create_lambda_applications_with_many/
---
Hi, I'm creating a website that needs many functions. Currently, we have multiple repositories (users, processing, management, etc.), each repo with multiple functions each serving a specific logic.

We just hit the CloudFormation limit of 200 resources (tho we have only 25 functions) in one of the repos. As it's not logical to separate the same type of service (e.g. two user repos), any ideas on how to handle this? A repo for each lambda (e.g. EditImageDetails)? Separate the services further? We do need the lambdas to access a shared general code and that's a problem too.

Thanks for the help!
## [5][How are you securely tunneling to RDS instances in a private subnet?](https://www.reddit.com/r/aws/comments/j0droo/how_are_you_securely_tunneling_to_rds_instances/)
- url: https://www.reddit.com/r/aws/comments/j0droo/how_are_you_securely_tunneling_to_rds_instances/
---
I was curious what process people are using to securely access their RDS databases in a private subnet.

My current process is:

* Generate &amp; send temp ssh keys to the bastion in public subnet using ec2-instance-connect
* Bastion is firewalled off to only allow ingress from my ip
* SSH tunnel to the RDS instance through the bastion to use psql

Is there a more secure way to do it? I'm not fond on having to patch/maintain a bastion host. Perhaps a managed solution that doesn't require a bastion host to get the tunnel? I checked out Secure Session Manager but that seems to still require an EC2 server in the middle.
## [6][Caddy or Nginx ?](https://www.reddit.com/r/aws/comments/j0qmby/caddy_or_nginx/)
- url: https://www.reddit.com/r/aws/comments/j0qmby/caddy_or_nginx/
---
We need to automatically and programmatically generate domain names and certificates for customers (potentially 10-100Ks of customers) in a scalable, reliable and responsive way.

We have a serverless infrastructure (cloudfront / S3 / with dynamodb + lambda + api gate way serverless backend), so ideally we would have liked to use route 53 and AWS certificate manager and route the domains to our cloudfront distribution but there is no way to attach the customers' certificates.

Hence, we've been thinking about nginx or caddy as alternative. What are your thoughts ?  Is there a way to do this serverless ?

Or should we go for nginx or caddy proxy that generates domains and certificates on the go behind an  ELB ?
## [7][AWS API Gateway: HTTP vs REST](https://www.reddit.com/r/aws/comments/j0l1eh/aws_api_gateway_http_vs_rest/)
- url: https://www.reddit.com/r/aws/comments/j0l1eh/aws_api_gateway_http_vs_rest/
---
[https://www.learnaws.org/2020/09/12/rest-api-vs-http-api/](https://www.learnaws.org/2020/09/12/rest-api-vs-http-api/)
## [8][AWS is adopting Google deprecation policies and want your feedback :)](https://www.reddit.com/r/aws/comments/j05uk0/aws_is_adopting_google_deprecation_policies_and/)
- url: https://i.redd.it/bbe3obg5shp51.png
---

## [9][Create CloudWatch Alarm for EC2 running over 2hr?](https://www.reddit.com/r/aws/comments/j0k4p1/create_cloudwatch_alarm_for_ec2_running_over_2hr/)
- url: https://www.reddit.com/r/aws/comments/j0k4p1/create_cloudwatch_alarm_for_ec2_running_over_2hr/
---
Is it possible to create a CloudWatch alarm to send an SNS when an EC2 is running over 2hr? Currently I have this (in yaml):

    F1RunningAlarm:
        Type: AWS::CloudWatch::Alarm
        Properties:
          AlarmName: F1LongRunning
          AlarmDescription: Alarms when the F1 instance is running for a long time
          AlarmActions:
            - !Ref 'F1SNSTopic'
          MetricName: CPUUtilization
          Namespace: AWS/EC2
          ComparisonOperator: GreaterThanOrEqualToThreshold
          EvaluationPeriods: 1
          Period: 7200
          Statistic: Average
          Threshold: 0
          TreatMissingData: notBreaching
          Dimensions:
          - Name: "InstanceId"
            Value:
              !Ref EC2

Currently it Alarms immediately as the CPU utilization passes 0 in the 2 hour period. I'd like it to alarm after the CPU has been running for more than 2 hour. Is this possible?
## [10][AWS CDK - Validation Error; Codepipeline](https://www.reddit.com/r/aws/comments/j0ocmx/aws_cdk_validation_error_codepipeline/)
- url: https://www.reddit.com/r/aws/comments/j0ocmx/aws_cdk_validation_error_codepipeline/
---
Morning all, 

I'm running into some problems with AWS CDK, I'm trying to create 2x AMI's in a single build job. Here is my code:

```typescript
let jenkinsNodeTypes = ["master", "slave"]
    for (let value of jenkinsNodeTypes) {
    const jenkinsDeployment = new codebuild.PipelineProject(this, `JenkinsDeployment-${value}`, {
      buildSpec: codebuild.BuildSpec.fromSourceFilename(`jenkins-${value}/buildspec.yml`),
      projectName: 'JenkinsInfraStack',
      environment: {
        buildImage: codebuild.LinuxBuildImage.AMAZON_LINUX_2_3
      }
    })
      
    const jenkinsSourceOutput = new codepipeline.Artifact();
    const jenkinsBuildOutput = new codepipeline.Artifact(`jenkinsBuildOutput-${value}`);
    
    const jenkinsPipeline = new codepipeline.Pipeline(this, `jenkinsPipeline-${value}`, {
      pipelineName: 'jenkinsPipeline',
      restartExecutionOnUpdate: true,
    })

    jenkinsPipeline.addStage({
      stageName: 'PullFromGit',
      actions: [
        new codepipeline_actions.GitHubSourceAction({
          actionName: `PullFromGithub-Jenkins${value}`,
          owner: '&lt;OWNER&gt;',
          repo: '&lt;REPO&gt;',
          oauthToken: cdk.SecretValue.secretsManager('GitHubToken'),
          branch: 'master',
          output: jenkinsBuildOutput,
          trigger: codepipeline_actions.GitHubTrigger.WEBHOOK,
        }),
      ],
    })

    jenkinsPipeline.addStage({
      stageName: 'Build',
      actions: [
        new codepipeline_actions.CodeBuildAction({
          actionName: `buildJenkins${value}`,
          project: jenkinsDeployment,
          input: jenkinsBuildOutput,
          outputs: [jenkinsBuildOutput],
        })
      ]
    })
  }
```

but I'm getting this output :/ 

```
{path}/node_modules/@aws-cdk/core/lib/private/synthesis.ts:151
    throw new Error(`Validation failed with the following errors:\n  ${errorList}`);
          ^
Error: Validation failed with the following errors:
  [JenkinsInfraStackGBP/jenkinsPipeline-master] Both Actions 'PullFromGithub-Jenkinsmaster' and 'buildJenkinsmaster' are producting Artifact 'jenkinsBuildOutput-master'. Every artifact can only be produced once.
  [JenkinsInfraStackGBP/jenkinsPipeline-slave] Both Actions 'PullFromGithub-Jenkinsslave' and 'buildJenkinsslave' are producting Artifact 'jenkinsBuildOutput-slave'. Every artifact can only be produced once.
    at validateTree ({path}/node_modules/@aws-cdk/core/lib/private/synthesis.ts:151:11)
    at Object.synthesize ({path}/node_modules/@aws-cdk/core/lib/private/synthesis.ts:25:5)
    at App.synth ({path}/node_modules/@aws-cdk/core/lib/stage.ts:175:23)
    at process.&lt;anonymous&gt; ({path}/node_modules/@aws-cdk/core/lib/app.ts:112:45)
    at Object.onceWrapper (events.js:421:26)
    at process.emit (events.js:314:20)
    at process.EventEmitter.emit (domain.js:486:12)
    at process.emit ({path}/node_modules/source-map-support/source-map-support.js:495:21)
```

Any help would be much appreciated. 

Thank you.
## [11][Out of my depth here. Got asked to set up a Windows VM.](https://www.reddit.com/r/aws/comments/j0pij6/out_of_my_depth_here_got_asked_to_set_up_a/)
- url: https://www.reddit.com/r/aws/comments/j0pij6/out_of_my_depth_here_got_asked_to_set_up_a/
---
I'm totally out of my depth and got asked to set up a Windows VM with .net, SQL, IIS to allow someone to RDP onto it and set up a web application. 

I guess this can be done with EC2? But how do I do all the other things? I don't even know what to google and whether this could just turn into an expensive waste of time. Any help?
