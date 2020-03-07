# aws
## [1][EC2 has detected degradation of the underlying hardware hosting your Amazon EC2 instance...](https://www.reddit.com/r/aws/comments/feft7z/ec2_has_detected_degradation_of_the_underlying/)
- url: https://www.reddit.com/r/aws/comments/feft7z/ec2_has_detected_degradation_of_the_underlying/
---
Dear Jeff,

Is there any way we could get this email to include the \`Name\` tag??

Thanks,

Me
## [2][Experimenting AWS RedShift Column Level ACL](https://www.reddit.com/r/aws/comments/fev4e5/experimenting_aws_redshift_column_level_acl/)
- url: https://thedataguy.in/experimenting-aws-redshift-column-level-acl/
---

## [3][Has anyone interviewed for a Cloud Infrastructure Architect role at AWS?](https://www.reddit.com/r/aws/comments/feobee/has_anyone_interviewed_for_a_cloud_infrastructure/)
- url: https://www.reddit.com/r/aws/comments/feobee/has_anyone_interviewed_for_a_cloud_infrastructure/
---
Does anyone have any experience with the interview process for this role? It seems like I can find a variety of information on SD2 but not on this role specifically.
## [4][Cognito: Unable to signUp with amplify:](https://www.reddit.com/r/aws/comments/fes9az/cognito_unable_to_signup_with_amplify/)
- url: https://www.reddit.com/r/aws/comments/fes9az/cognito_unable_to_signup_with_amplify/
---
    {
    code: "InvalidParameterException", name: "InvalidParameterException", message: "Attributes did not conform to the schema: custom:f…astname: Attribute does not exist in the schema.↵"}
    code: "InvalidParameterException"
    name: "InvalidParameterException"
    message: "Attributes did not conform to the schema: custom:firstname: Attribute does not exist in the schema.↵custom:lastname: Attribute does not exist in the schema.↵"
    }

The custom attributes are present in the userpool

[https://imgur.com/Hc3uQIa](https://imgur.com/Hc3uQIa)

code:

     Auth.signUp(
          {
            'username' : email,
            'password' : password,
            attributes: {
              'email' : email,
              'custom:firstname' : firstname,
              'custom:lastname' : lastname
            }
          }
        ).then(data =&gt; console.log(data))
         .catch(err =&gt; console.log(err));

&amp;#x200B;
## [5][CloudWatch now offers composite alarms. Great for reducing alarm fatigue and triggering scale down actions](https://www.reddit.com/r/aws/comments/fe6s3y/cloudwatch_now_offers_composite_alarms_great_for/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-cloudwatch-now-allows-you-to-combine-multiple-alarms/
---

## [6][None of our SQS to lambda triggers are working](https://www.reddit.com/r/aws/comments/fel75q/none_of_our_sqs_to_lambda_triggers_are_working/)
- url: https://www.reddit.com/r/aws/comments/fel75q/none_of_our_sqs_to_lambda_triggers_are_working/
---
For some bizarre reason, all of our sqs to lambda triggers stopped working at a certain time. However the API gateway integrated with lambdas are. Anyone have a similar circumstance occurred to them?
## [7][Api call fails due to CORS policy.](https://www.reddit.com/r/aws/comments/feplc0/api_call_fails_due_to_cors_policy/)
- url: https://www.reddit.com/r/aws/comments/feplc0/api_call_fails_due_to_cors_policy/
---
The api works fine when the integration response has only the default response 200. But doesn't work when after I added 503 in integration response. I have enabled CORS after adding 503 and redeployed. Still not working.
## [8][Looking for a replacement to Coinbase's assume-role CLI tool? Check out aws_longer.](https://www.reddit.com/r/aws/comments/feirqp/looking_for_a_replacement_to_coinbases_assumerole/)
- url: https://www.reddit.com/r/aws/comments/feirqp/looking_for_a_replacement_to_coinbases_assumerole/
---
Do you use Coinbase's now deprecated [assume-role](https://github.com/coinbase/assume-role)? Are you annoyed by having to
enter your MFA token into multiple terminal sessions, or simply having
to type your MFA token in more than once a day? If so, check out
`aws_longer`!

`aws_longer` uses your computer's keyring to store a 36-hour temporary
credential that is used to subsequently assume-roles. The assumed
credentials last an hour each, necessitating re-assuming, however,
these subsequent re-assumes do not require MFA unless 36 hours hours
have passed since your last MFA.

To get started check out https://github.com/appfolio/aws_longer for installation and
testing instructions.
## [9][Proof My Kinesis Firehose Calculations](https://www.reddit.com/r/aws/comments/fehzsx/proof_my_kinesis_firehose_calculations/)
- url: https://www.reddit.com/r/aws/comments/fehzsx/proof_my_kinesis_firehose_calculations/
---
Heyo! I often think I'm good at calculating the cost of my individual AWS services but sometimes find I miss some small/obvious things, I was hoping yall could proof my costing train of thought.

[Firehose Pricing](https://aws.amazon.com/kinesis/data-firehose/pricing/)

I've enabled logging on my WAF using Firehose, and a Lambda script that removes all valid requests and outputs only blocked requests (as per [this guide](https://aws.amazon.com/blogs/security/trimming-aws-waf-logs-with-amazon-kinesis-firehose-transformations/)). The WAF is looking at my ELB, which receives a max (recorded so far) number of requests of 288,000 a day.  Going with that as the "daily amount of requests" I went with the following calculations.

Costing works on 0.032/GB (In Canada) where each WAF request is about 1.5KB, but rounds up to 5KB per req for costing, at our max of 288,000 requests per day gives us:

288,000 requests per day \* 5KB/req \* 30 days in a month = 43,200,000KB worth of requests per month, or 43.2GB \* .032 per GB = $1.38 a month.

Is that right? Am I really only going to pay up to $1.38/month at my current load?
## [10][Backuping of AWS](https://www.reddit.com/r/aws/comments/feoe96/backuping_of_aws/)
- url: https://www.reddit.com/r/aws/comments/feoe96/backuping_of_aws/
---
Good day, usually I have small projects on AWS. I export for Git policies, API Gateway swagger. But looks like I cannot export IAM with users, roles. What you can recommend? What is the best practices about backuping entire AWS account? Also this is about Infrastructure as a Code. CloudFormation? But looks like I cannot export current AWS account.
