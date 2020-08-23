# aws
## [1][CS graduate - concerned about unexpected charges](https://www.reddit.com/r/aws/comments/iezimy/cs_graduate_concerned_about_unexpected_charges/)
- url: https://www.reddit.com/r/aws/comments/iezimy/cs_graduate_concerned_about_unexpected_charges/
---
I want to start using AWS for my projects as a lot of employers in my area are looking for AWS experience.

My concern is that it appears I would be liable for any server costs if someone were to do something like point a botnet at one of my sites. Even if I set up the usage alert, couldn't they rack up tens of thousands of dollars in debt before I could even get to a computer to shutdown my service?

Is everyone just okay with that level of exposure or is there something I'm missing?
## [2][Is it dangerous to publicly display the names of S3 buckets?](https://www.reddit.com/r/aws/comments/ieyxi8/is_it_dangerous_to_publicly_display_the_names_of/)
- url: https://www.reddit.com/r/aws/comments/ieyxi8/is_it_dangerous_to_publicly_display_the_names_of/
---
If I do a screen recording or post in youtube a video that shows my S3 bucket list, is it dangerous in anyway? Do I need to blur or cover the names of my buckets? I have buckets which work as html sites that say "objects can be public".
## [3][Building an AWS flowchart.. what do you use?](https://www.reddit.com/r/aws/comments/if1xdr/building_an_aws_flowchart_what_do_you_use/)
- url: https://www.reddit.com/r/aws/comments/if1xdr/building_an_aws_flowchart_what_do_you_use/
---
So I'm just starting my journey into the cloud realm and I'm wondering what software or website do you experienced folks use to generate the service flowcharts using the actual AWS symbology?
## [4][Tiny CLI to save AWS costs in development environments when you're sleep](https://www.reddit.com/r/aws/comments/iell1j/tiny_cli_to_save_aws_costs_in_development/)
- url: https://news.ycombinator.com/item?id=24245166
---

## [5][Aws Amplify Github Continous Deployment Issues - React App - Looking for Advice](https://www.reddit.com/r/aws/comments/iewpdq/aws_amplify_github_continous_deployment_issues/)
- url: https://www.reddit.com/r/aws/comments/iewpdq/aws_amplify_github_continous_deployment_issues/
---
I'm having issues with deploying my SPA through AWS Amplify. I have tried deploying through Github hooks? &amp; S3 and Cloudfront.

My issue stems from
1. code not being updated on a push to my Github master
2. code not being updated when using S3 &amp; Cloudfront w/ ```amplify publish```

I understand that I'm just looking to use an S3 bucket with a Cloudfront distribution so I'm looking for advice on how your continuous deployment pipeline is setup with a react app / SPA.

Any tips and or pointers on a flawless setup would be great. I can't have these issues on a production site.

#### side note
- I have 0 of these issues when using the elastic beanstalk CLI with my node/express back-end, deployment is flawless.
## [6][Better way to use AWS CLI with multiple accounts](https://www.reddit.com/r/aws/comments/ies847/better_way_to_use_aws_cli_with_multiple_accounts/)
- url: https://www.reddit.com/r/aws/comments/ies847/better_way_to_use_aws_cli_with_multiple_accounts/
---
I'm not sure if the way I'm handling profiles and credentials for AWS CLI tool set makes sense.

I have to occasionally run AWS CLI across one hundred accounts.

I have all one hundred  accounts under profile in ./aws/credentials like so (the account names really are $org1, $org2, and so on to 100).

`[default]`

`aws_access_key_id = redacted`

`aws_secret_access_key = redacted`

`[org1]`

`role_arn = arn:aws:iam::account1:role/redacted`

`source_profile = default`

`[org2]`

`role_arn = arn:aws:iam::account2:role/redacted`

`source_profile = default`

&amp;#x200B;

Use case: 'need a list of all bucket names for all accounts'. In pseudo shell

`$ aws s3 ls --profile $name`

Where $name is incremented from 1 to 100 in a scipt.  This works.

But i have to think there is a better way than maintaining a 303-line credentials file.

Your thoughts?
## [7][Using multiple authorization types with AWS AppSync GraphQL APIs](https://www.reddit.com/r/aws/comments/iey2an/using_multiple_authorization_types_with_aws/)
- url: https://www.reddit.com/r/aws/comments/iey2an/using_multiple_authorization_types_with_aws/
---
**What AWS Services am I utilizing?**

- Cognito
- Appsync GraphQL api
- Amplify CLI (4.27.3)


[This blog post][1] shows how to set multiple authorisation types to models and fields in graphql transform.

Lets say I have an `@model Blog`

```
    type Blog @model {
       id: ID!
       adminUserId: String
       name: String!
       posts: [Post] @connection(keyName: "byBlog", fields: ["id"])
    }
```

Using this shcema will autogenerate the following mutations/queries;
`createBlog`
`updateBlog`
`deleteBlog`
`getBlog`
`listBlogs`


I want `createBlog`, `updateBlog` and `deleteBlog` to have authorisation type `@aws_iam`.
I want `getBlog` and `listBlogs` to have my default authorisation type `@aws_cognito_user_pools`


How can I define this in my `schema.graphql`?
I can not set the authorisation type directly on the mutations/queries as they are not defined in my `schema.graphql` file.

I am able to set the auth types directly in the complete schema that is generated in the cloud (`AWS AppSync &gt; API Name &gt; Schema` ) because here all the queries/mutations are all defined. But this schema will be re-written every time I run `amplify push`.

There must be a better way?

  [1]: https://aws.amazon.com/blogs/mobile/using-multiple-authorization-types-with-aws-appsync-graphql-apis/
## [8][Elastic Beanstalk + SSL woes](https://www.reddit.com/r/aws/comments/ieroxf/elastic_beanstalk_ssl_woes/)
- url: https://www.reddit.com/r/aws/comments/ieroxf/elastic_beanstalk_ssl_woes/
---
I am trying to get my node project up and running with SSL(from lets encrypt) on a single instance EB, but I can't seem to configure this properly.

I have tried [Amazon's way](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/https-singleinstance-nodejs.html) by dumping the contents of the certs and keys into a file in .ebextensions. This way doesn't throw any errors when deploying. But refuses to connect when accessing on HTTPS.  
Googling told me this was a issue with DNS-- Needing to configure my domain with a CNAME record rather an A record . Which is how I had this setup from the beginning, But that advice was for EB with a load balancer, So I figured perhaps with a single instance the its configured the other way around, So I've switched it an  A record, still refuses to connect(haven't changed back to a CNAME though).  
I have port 443 open.  


I have tried the other [widely recommended way](https://www.tutcodex.com/ssl-on-single-instance-elastic-beanstalk-tutorial/)(link for example, I've looked at &amp; tried dozens of variations on it), which leads to errors after errors while deploying. 

This is the second full(12+ hours) day I have wasted trying to get this to work. I know enough sysadmin stuff to get by, but my wheelhouse is development, and I'd really like to get back to that. If someone can offer some advice on how to fix the "connection refused" problem I would very much appreciate it. Alternatively if you have a tutorial going about this another way, and you're sure it will work, I'll be happy to have that too.
## [9][How to organize infrastructure responsibilities if we build Micro-services with AWS CDK](https://www.reddit.com/r/aws/comments/ieece7/how_to_organize_infrastructure_responsibilities/)
- url: https://www.reddit.com/r/aws/comments/ieece7/how_to_organize_infrastructure_responsibilities/
---
I really like AWS CDK very much because it allows us to organize/align our team purely into a developer first way, i.e., each repo (say \`billing-service\`, or \`checking-out-service\`) directly corresponds to a vertical function team -- the repo contains the node.js service code and also contains the infrastructure setup code (i.e, how to setup Beanstalk cluster).

&amp;#x200B;

However, it seems that we still need a horizontal repo (and team) to take care the shared things across the vertical function repos (and teams) -- for example, if \`billing-service\` and  \`checking-out-service\` shares the same VPC, or even share the same ECS cluster, then the repo that is in charge of the shared VPC and ECS cluster seemingly have to be an independent repo, say 'vpc-and-ecs'

&amp;#x200B;

My question here are the following two:

1. In the above example, if we have to have the third repo of \`vpc-and-ecs\`, how can the  \`billing-service\` and  \`checking-out-service\` repo knows the output of \`vpc-and-ecs\`, such as CIDR block or ECS cluster ID and etc.? (I guess hard coding is ok at the very beginning, but I feel it's very hard to maintain across the team)
2. If we need to update the shared infrastructure code (vpc-and-ecs), say we want to change the VPC CIDR or change the subnets, it probably will have inevitable effect towards the  \`billing-service\` and  \`checking-out-service\` repo, how can we manage the repos change dependency and cross team communication?

Anyone thought of how to work with CDK in a large cross team?
## [10][Is there any eBook / website with sample aws architecture diagrams? Something similar to the aws this is my architectur series?](https://www.reddit.com/r/aws/comments/iehi9b/is_there_any_ebook_website_with_sample_aws/)
- url: https://www.reddit.com/r/aws/comments/iehi9b/is_there_any_ebook_website_with_sample_aws/
---

