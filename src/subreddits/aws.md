# aws
## [1][AWS workflows to set up services automatically via GitHub CI/CD: just fork &amp; reuse!](https://www.reddit.com/r/aws/comments/j3ssqy/aws_workflows_to_set_up_services_automatically/)
- url: https://github.com/didier-durand/aws-workflows-on-github/
---

## [2][AWS SSO with Just-in-time Privileges](https://www.reddit.com/r/aws/comments/j3nue4/aws_sso_with_justintime_privileges/)
- url: https://www.reddit.com/r/aws/comments/j3nue4/aws_sso_with_justintime_privileges/
---
Azure has a great concept called PIM or Privileged Identity Management

[https://docs.microsoft.com/en-us/azure/active-directory/privileged-identity-management/pim-configure](https://docs.microsoft.com/en-us/azure/active-directory/privileged-identity-management/pim-configure)

It allows specific users to elevate privileges, with customizable expiration, authorization (currently unfortuntely only email no mobile push), notifications etc.

It's a great way for admins to have only basic permissions for day-to-day admin, but consciously (with authorization, explicit reason, for limited duration) elevate privs when required.

I'm looking at moving our company's AWS account auth to using a single identity account (with AWS SSO &amp;. Azure as IdP) and was wondering if anyone had built similar? Just-in-time privesc.

In terms of the assumable roles - I guess you could update AssumeRole principles, or probably better use ABAC [https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction\_attribute-based-access-control.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction_attribute-based-access-control.html)

While we could try to leverage Azure PIM, it's limited to Azure Roles ... so you're need e.g. one Azure Role per Assumable AWS Role ... then based on that role membership, add SAML Assertion .... reauth ... kind of yuck.

Probably best would be to deploy a system in the Identity Account to manage it, dynamically add/remove tags on users' Principal roles (the ones they sign in to, from which they'll assume into others), and the target roles could add a Condition based on the principle's tags to the Assume role policy (and require the assuming principal be in the Identity Account(!))

I did find this excellent writeup on the area, by our previous Solutions Architect Louay (legend btw), it's related, but not 100% what I'm after [https://aws.amazon.com/blogs/security/attribute-based-access-control-ad-fs-simplify-iam-permissions-management/](https://aws.amazon.com/blogs/security/attribute-based-access-control-ad-fs-simplify-iam-permissions-management/)

(also now I notice this isn't necessarily about AWS SSO but the more general concept of "just-in-time approved/audited privesc in AWS")

Just spitballing ideas and thought it might be interesting to discuss here, cheers
## [3][IAM having timeout issues?](https://www.reddit.com/r/aws/comments/j3gync/iam_having_timeout_issues/)
- url: https://www.reddit.com/r/aws/comments/j3gync/iam_having_timeout_issues/
---
Errors from Terraform complaining about the connection being reset, and getting this from the CLI: 

    Connection was closed before we received a valid response from endpoint URL: "https://iam.amazonaws.com/".

Console shows this error:
&gt; Http request timed out enforced after 999ms

Not happening with all my accounts, strangely enough.

EDIT: Just resolved?
## [4][AWS Aurora - What issues have you had?](https://www.reddit.com/r/aws/comments/j3opqz/aws_aurora_what_issues_have_you_had/)
- url: https://www.reddit.com/r/aws/comments/j3opqz/aws_aurora_what_issues_have_you_had/
---
We're considering moving from RDS Mysql to AWS Aurora. The promise is great - but in reality we've found that the devil is in the details with a lot of AWS services. 

Has anyone migrated from RDS Mysql to AWS Aurora? Are there any gotchas or things to watch out for?

If its relevant, our application is using Ruby on Rails v5.0.7 - soon to upgrade to v6 where we can leverage different database connections (ie utilizing read only nodes)
## [5][We're building 7777, a tool to simplify connecting to RDS databases in VPC, feedback is welcome](https://www.reddit.com/r/aws/comments/j3blej/were_building_7777_a_tool_to_simplify_connecting/)
- url: https://port7777.com/
---

## [6][S3 NextContinuationToken question](https://www.reddit.com/r/aws/comments/j3r7o7/s3_nextcontinuationtoken_question/)
- url: https://www.reddit.com/r/aws/comments/j3r7o7/s3_nextcontinuationtoken_question/
---
Hey, guys, i was wondering about how will nextContinuationToken behave in the scenario where i listObjects, say a 1000, get a nextContinuationToken and delete those 1000 objects i got through listObjects call from s3, will the token be invalidated and fail next call to listObjects using the token? Or will it work correctly and list other objects in the bucket?
## [7][AWS::CertificateManager::Certificate Automatic DNS validation?](https://www.reddit.com/r/aws/comments/j3tonj/awscertificatemanagercertificate_automatic_dns/)
- url: https://www.reddit.com/r/aws/comments/j3tonj/awscertificatemanagercertificate_automatic_dns/
---
So, I was thinking about migrating away from the lambda to generate my certificates - and to jump straight in the new Cloudformation feature for the AWS Certificate manager, that looks like this:
```
  MyCertificate:
    Type: AWS::CertificateManager::Certificate
    Properties:
      DomainName: "stuff.mydomain.com"
      SubjectAlternativeNames:
        - "morestuff.mydomain.com"
      ValidationMethod: DNS
      DomainValidationOptions:
        - HostedZoneId: XXXXXXXXXXXXXXXXXXXXX
          DomainName: mydomain.com
```
I thought that this would put the required entries straight in the Route53 Zone config, but I ended up waiting for about 2hrs.

Is the above example somehow wrong, or is the AWS::CertificateManager::Certificate resource *not* supposed to actually input the records to the supplied Route53 zone?
## [8][How are you providing access to S3 data in EMR](https://www.reddit.com/r/aws/comments/j3tdvg/how_are_you_providing_access_to_s3_data_in_emr/)
- url: https://www.reddit.com/r/aws/comments/j3tdvg/how_are_you_providing_access_to_s3_data_in_emr/
---
Hey everyone I have an AWS account that stores tons of large datasets using parquet on S3.

I have a separate account that will leverage EMR and Spark to process this data.

I am wondering how people are providing access to their data. S3 bucket policies? Lake Formation (not supported on spark submit or cross account), etc.

Thanks in advance!
## [9][How to securely share resources in bulk with other accounts?](https://www.reddit.com/r/aws/comments/j3q3nr/how_to_securely_share_resources_in_bulk_with/)
- url: https://www.reddit.com/r/aws/comments/j3q3nr/how_to_securely_share_resources_in_bulk_with/
---
Our AWS resources are organized by tags for different projects. Looking for a secure way to share all resources tagged “ProjectName=Alpha” with another AWS account.
## [10][Can anyone tell me or send me documentation on how AWS has enough IPv4 addresses for all of the publicly routed instances?](https://www.reddit.com/r/aws/comments/j3luvy/can_anyone_tell_me_or_send_me_documentation_on/)
- url: https://www.reddit.com/r/aws/comments/j3luvy/can_anyone_tell_me_or_send_me_documentation_on/
---
I'm assuming they have tens of thousands of instances running at all times(or more) which are publicly routed with IPv4 addresses. 

I understand that they recycle IPs when not in use, but it seems that they'd run out at some point at the rate they're growing.
