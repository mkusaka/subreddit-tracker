# aws
## [1][7 Common Mistakes in Architecture Diagrams](https://www.reddit.com/r/aws/comments/eokico/7_common_mistakes_in_architecture_diagrams/)
- url: https://blog.ilograph.com/posts/diagram-mistakes/
---

## [2][Part 5 Building an Imgur clone with Vue.js and Serverless - Authentication between our Backend and Frontend - Now Live!](https://www.reddit.com/r/aws/comments/eoaxu0/part_5_building_an_imgur_clone_with_vuejs_and/)
- url: https://tutorialedge.net/projects/building-imgur-clone-vuejs-nodejs/part-5-authentication/
---

## [3][Introducing Pubby - DAZN's custom, high-scale WebSockets solution](https://www.reddit.com/r/aws/comments/eojr9r/introducing_pubby_dazns_custom_highscale/)
- url: https://medium.com/dazn-tech/introducing-pubby-our-custom-websockets-solution-c5764e3a7dcb?source=friends_link&amp;sk=a206d4ba54317ed6eae3fce0521e03b2
---

## [4][Preferred architecture diagram software?](https://www.reddit.com/r/aws/comments/eodjob/preferred_architecture_diagram_software/)
- url: https://www.reddit.com/r/aws/comments/eodjob/preferred_architecture_diagram_software/
---

Hi all,

Wanted to hear what folks think is the 'best' architecture / flow diagram software.

I personally was looking for something free, easy to use, could sync across devices, and natively had AWS icons. I ended up going to draw.io. 

I made a video highlighting my rationale and why I chose draw.io here - https://youtu.be/i0pMpnxHN6g

Curious to hear everyone's thoughts and what you use daily.
## [5][Bundesliga Goes All-In on AWS](https://www.reddit.com/r/aws/comments/eo78u1/bundesliga_goes_allin_on_aws/)
- url: https://www.reddit.com/r/aws/comments/eo78u1/bundesliga_goes_allin_on_aws/
---
As an AWS user, German and football fan, I find this in multiple ways interesting.  

[Bundesliga Goes All-In on AWS](https://press.aboutamazon.com/news-releases/news-release-details/bundesliga-goes-all-aws-revolutionize-football-viewing)

Mainz 05 ftw!
## [6][Getting software to run on two AWS accounts.](https://www.reddit.com/r/aws/comments/eokylt/getting_software_to_run_on_two_aws_accounts/)
- url: https://www.reddit.com/r/aws/comments/eokylt/getting_software_to_run_on_two_aws_accounts/
---
I am trying to run software on two separate EC2 windows server and keep getting an error that the software is already running on another computer.

I created a new AWS account with new account ID and fresh intsalled the software.

How would the software know that I'm using it when it's a new instance, on a new account, with a new ID?

Potentially it picks up on where I'm remote viewing from?

Any ideas on how I can fix this?

Thanks.
## [7][Can AWS credit be applied towards support plans?](https://www.reddit.com/r/aws/comments/eoeeoi/can_aws_credit_be_applied_towards_support_plans/)
- url: https://www.reddit.com/r/aws/comments/eoeeoi/can_aws_credit_be_applied_towards_support_plans/
---
We have some startup credit. If I select a support plan, will it be paid for with our unused credit?
## [8][I am an idiot who needs help using the Glacier Vaults](https://www.reddit.com/r/aws/comments/eofcfe/i_am_an_idiot_who_needs_help_using_the_glacier/)
- url: https://www.reddit.com/r/aws/comments/eofcfe/i_am_an_idiot_who_needs_help_using_the_glacier/
---
I want to use the Glacier Vault to store a few TB of video.  I set up an IAM use account. I created a vault.  Now, I cant figure out what I need to do to upload my video for long term storage. 

I downloaded the AWS Command Line Interface Setup Wizard. Once it is installed it is not giving me any nice prompts on what to do next.  

Oh, so very confused.
## [9][CodeBuild+GitHub - How can I build a branch upon PULL_REQUEST_MERGED?](https://www.reddit.com/r/aws/comments/eok0hn/codebuildgithub_how_can_i_build_a_branch_upon/)
- url: https://www.reddit.com/r/aws/comments/eok0hn/codebuildgithub_how_can_i_build_a_branch_upon/
---
**The need** \- when merging a pull-request to a branch, I want CodeBuild to build the latest branch's commit, **not** the pull-request.I'm using CloudFormation, here's the triggers snippet:

     Triggers:
       Webhook: true
       FilterGroups:
         - - Type: EVENT
             Pattern: PULL_REQUEST_CREATED, PULL_REQUEST_UPDATED, PULL_REQUEST_REOPENED
           - Type: BASE_REF
             Pattern: !Sub "refs/heads/${GithubBranchName}$"
             ExcludeMatchedPattern: false

I've tried adding PULL\_REQUEST\_MERGED in the same CodeBuild project, but it builds the PR.

I've also tried creating a new CodeBuild project with PULL\_REQUEST\_MERGED only, and I tweaked the BASE\_REF and HEAD\_REF, but still no luck, the pull-request is built, instead of the branch.

Even though I'm using CloudFormation, feel free to reply with screenshots referring to AWS Console.

Is it even possible?
## [10][Maintanance Page when Target Group is down](https://www.reddit.com/r/aws/comments/eoi2x4/maintanance_page_when_target_group_is_down/)
- url: https://www.reddit.com/r/aws/comments/eoi2x4/maintanance_page_when_target_group_is_down/
---
Hello everybody,  
we have several development and staging system who are shut down every night with the ec2-scheduler. In the front we have a ALB and using target groups. This works very well.   
But, what is the best practice to get a custom maintnance Page for this case? I want a Page with "This Enviroment is recently shut down" or something.  


Any tipps?  


Thanks for your help!
