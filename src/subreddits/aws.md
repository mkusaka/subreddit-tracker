# aws
## [1][Welcome to new mod u/_abhayshah!](https://www.reddit.com/r/aws/comments/j5trs4/welcome_to_new_mod_u_abhayshah/)
- url: https://www.reddit.com/r/aws/comments/j5trs4/welcome_to_new_mod_u_abhayshah/
---
Thrilled to expand the mod team in order better serve the community. As /r/aws grows, u/_abhayshah will be able to assist with AMAs, post flair, mod queue/mail, building out the Wiki, and more!   


Please give him a nice /r/aws welcome and let us know how we can improve things together going forward.
## [2][GraphQL Tools &amp; libraries pt.3](https://www.reddit.com/r/aws/comments/j7bltp/graphql_tools_libraries_pt3/)
- url: https://blog.graphqleditor.com/graphql-tools-part3/
---

## [3][The Complete AWS Lambda Handbook for Beginners (Part 1)](https://www.reddit.com/r/aws/comments/j6tlnf/the_complete_aws_lambda_handbook_for_beginners/)
- url: https://dashbird.io/blog/complete-aws-lambda-handbook-beginners-part-1/
---

## [4][New – Redis 6 Compatibility for Amazon ElastiCache](https://www.reddit.com/r/aws/comments/j70d5r/new_redis_6_compatibility_for_amazon_elasticache/)
- url: https://aws.amazon.com/blogs/aws/new-redis-6-compatibility-for-amazon-elasticache/
---

## [5][reuse cloudfront for dev, test and prod?](https://www.reddit.com/r/aws/comments/j7a4qc/reuse_cloudfront_for_dev_test_and_prod/)
- url: https://www.reddit.com/r/aws/comments/j7a4qc/reuse_cloudfront_for_dev_test_and_prod/
---
Due to company policies, our global security team runs an uninformed Qualys scan on every CloudFront we have. This results in roughly 50k requests every other day (it scans for everything you can imagine, even WordPress and Joomla even though we don't use that).

To avoid doubling or tripling our costs, can I reuse a cloudfront that points to [dev.example.com](https://dev.example.com) and [test.example.com](https://test.example.com) ? Do I need to reuse my bucket for that as well or can those be separate?
## [6][Testing tomorrow](https://www.reddit.com/r/aws/comments/j7cwx0/testing_tomorrow/)
- url: https://www.reddit.com/r/aws/comments/j7cwx0/testing_tomorrow/
---
I’m testing for the Practitioner tomorrow. I’ve been studying for the past 3 weeks for this and every practice test I find I fail. Admittedly, I would be able to recite everything I’ve studied and I expect to get some answers wrong but failing these practice tests is starting to give me doubts. Anyone have any tips or advice I can use in my last 24 hours before the test?
## [7][What DB for Lambda to work with GEO SPATIAL data?](https://www.reddit.com/r/aws/comments/j78pyb/what_db_for_lambda_to_work_with_geo_spatial_data/)
- url: https://www.reddit.com/r/aws/comments/j78pyb/what_db_for_lambda_to_work_with_geo_spatial_data/
---
Hi Guys,

What would you recommend to use to be able to store and query (efficiently) spatial data with a serverless app on AWS?  
I am working with Node and there is a lib for geo hash for DynamoDB but its support is pretty uncertain, so I am not sure about it and unfortunately, DynamoDB does not support spatial data out of the box.  
I thought about using Mongo Atlas hosted in same DC, but network peering is available from M10 which is $70/month which does not go well with cost scaling and without it, first of all, will be slower and it will incur fees for data transfer out of AWS (pretty stupid...).  
Maybe someone has some experience with similar serverless app on AWS?
## [8][Any lamda sample/tutorial to serve a static website from S3?](https://www.reddit.com/r/aws/comments/j79j8e/any_lamda_sampletutorial_to_serve_a_static/)
- url: https://www.reddit.com/r/aws/comments/j79j8e/any_lamda_sampletutorial_to_serve_a_static/
---
This is what I post earlier:

[https://www.reddit.com/r/aws/comments/j6imbx/tried\_to\_build\_my\_personal\_blog\_using\_aws\_static/](https://www.reddit.com/r/aws/comments/j6imbx/tried_to_build_my_personal_blog_using_aws_static/)

I want to serve folder with HTML files on S3, Couldn't find on google

Please help, thank you
## [9][403 error when uploading file to S3 from browser](https://www.reddit.com/r/aws/comments/j7bypr/403_error_when_uploading_file_to_s3_from_browser/)
- url: https://www.reddit.com/r/aws/comments/j7bypr/403_error_when_uploading_file_to_s3_from_browser/
---
I'm trying to upload file from browser direct to S3 presigned url.  And uploading always gets stuck at \~80% and throws 403 Forbidden error (without response body).

I use PHP (Laravel) backend to generate presigned url:

    $s3 = Storage::disk('s3');
    $client = $s3-&gt;getDriver()-&gt;getAdapter()-&gt;getClient();
    $expiry = "+1 hour";
    
    $command = $client-&gt;getCommand('GetObject', [
        'Bucket' =&gt; config('filesystems.disks.s3.bucket'),
        'Key'    =&gt; Str::random(10),
        'ContentType'    =&gt; 'video/mp4',
    ]);
    
    $request = $client-&gt;createPresignedRequest($command, $expiry);
    
    return response()-&gt;json([
        'uri' =&gt; (string) $request-&gt;getUri(),
    ]);

And then execute PUT request with axios from browser:

    const options = {
        headers: {
            'Content-Type': file.type,
            'x-amz-acl': 'public-read',
        },
    };
    
    axios.put(signedUrl, file, options);
## [10][Has anyone used ChatQL?](https://www.reddit.com/r/aws/comments/j7b25h/has_anyone_used_chatql/)
- url: https://www.reddit.com/r/aws/comments/j7b25h/has_anyone_used_chatql/
---
I want to use AWS AppSync. I wanted to have Chat functionality in my app so my friend suggested me to use ChatQL. The feature I am looking most in ChatQL is to notify the sender that msg has been delivered or not, whether it is read by the receiver or not, along with caching functionality. I am using React Native on client side. Does anyone has experience with ChatQL? Is there any other service from aws which provides real time msg sending along with informing sender whether msg delivered to receiver or not, caching etc. It could be GraphQL or REST architecture. I was using Nodejs socket io lib for real time chat but I needed some trigger for whether msg was send correctly or not. Of course I can create those triggers on my own but I want to stay away from it as there is a high chance I might make some mistake on my part
## [11][Why isn't my s3 bucket secure?](https://www.reddit.com/r/aws/comments/j7azky/why_isnt_my_s3_bucket_secure/)
- url: https://www.wolfe.id.au/2020/10/08/why-isnt-my-s3-bucket-secure/
---

