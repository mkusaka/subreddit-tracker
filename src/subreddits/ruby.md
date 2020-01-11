# ruby
## [1][Building an Asynchronous Scheduler for Ruby](https://www.reddit.com/r/ruby/comments/en4pyh/building_an_asynchronous_scheduler_for_ruby/)
- url: https://www.youtube.com/watch?v=OlmySf03GUo&amp;feature=youtu.be
---

## [2][puma disable multi-threading](https://www.reddit.com/r/ruby/comments/en2jls/puma_disable_multithreading/)
- url: https://www.reddit.com/r/ruby/comments/en2jls/puma_disable_multithreading/
---
Hi, if I don't want to worry about my code being thread-safe, i can just set the max thread count to 1, and crank up the worker count to like 10. Is this the general way to configure puma to be multi-processing only without multi-threading?
## [3][A few questions about Gherkin syntax](https://www.reddit.com/r/ruby/comments/emwqrn/a_few_questions_about_gherkin_syntax/)
- url: https://www.reddit.com/r/ruby/comments/emwqrn/a_few_questions_about_gherkin_syntax/
---
Hi,   
I've seen a tutorial or two, read a few blogpost/random sources, but I have two questions I couldn't find an answer for.   
1. If I have a variable outcome  
showing errors or showing success messages  
can I do it as   


&gt;when \[fail\]  
then \[show error\]  
but if \[pass\]  
then \[show pass\]

  
2. In Feature, Background, and Scenario (btw, it's Scenario, or Scenario outline, or both?) I'm supposed to use extremely short sentences, or it can be normal writing?
## [4][How to raise exception, send it as a parameter to another method, and rescue the exception in the other method?](https://www.reddit.com/r/ruby/comments/emr4l0/how_to_raise_exception_send_it_as_a_parameter_to/)
- url: https://www.reddit.com/r/ruby/comments/emr4l0/how_to_raise_exception_send_it_as_a_parameter_to/
---
I have two methods as shown below, one tries to make a connection, if it fails

it raises an exception. That exception should then be rescued by another method,

who has been called by the method who raised the exception.

&amp;#x200B;

Currently the `conn` method raises the error, breaks the application and it's

never rescued. How do I rescue the exception, and is it possible to send the

exception as the response parameter?

&amp;#x200B;

    def conn
      begin 
        reponse = parse_response(
          conn.post(
            ..
          ).tap do |res|
            unless res.status == 200
              raise Error::AuthenticationError.new
            end
          end
        )
        ..
      end
    end
    
    def parse_response(response)
      body = 
        begin
          JSON.parse(response.body)
        rescue JSON::ParseError
        end
      rescue Error::AuthenticationError =&gt; e
        LOGGER.error e
      end
    end
## [5][Rubygem release heat maps on the Ruby Toolbox](https://www.reddit.com/r/ruby/comments/emh8tr/rubygem_release_heat_maps_on_the_ruby_toolbox/)
- url: https://www.ruby-toolbox.com/blog/2020-01-09/rubygem-release-heatmaps
---

## [6][new app with no javascript and webpack](https://www.reddit.com/r/ruby/comments/emfip7/new_app_with_no_javascript_and_webpack/)
- url: https://www.reddit.com/r/ruby/comments/emfip7/new_app_with_no_javascript_and_webpack/
---
How does using the --skip-javascript and --skip-webpack-install to create a new app affect anything?
## [7][is there a tty: true equivalent for docker run?](https://www.reddit.com/r/ruby/comments/emo2o2/is_there_a_tty_true_equivalent_for_docker_run/)
- url: https://www.reddit.com/r/ruby/comments/emo2o2/is_there_a_tty_true_equivalent_for_docker_run/
---
I want to start a container in detached mode, then I can exec -it into it in my own pace. I'm currently able to do that via docker-compose with **tty: true** in the yml.

Is there a version without docker-compose, that is, I can run with the **docker run** command? And without a running service in the container. (I'm able to do that without a running service with the docker-compose method)
## [8][Deep dive into Did You Mean](https://www.reddit.com/r/ruby/comments/em9fdc/deep_dive_into_did_you_mean/)
- url: https://shime.sh/deep-dive-into-did-you-mean
---

## [9][Ruby low memory and cpu on file upload](https://www.reddit.com/r/ruby/comments/eme8e4/ruby_low_memory_and_cpu_on_file_upload/)
- url: https://www.reddit.com/r/ruby/comments/eme8e4/ruby_low_memory_and_cpu_on_file_upload/
---
I'm uploading pdf files to the cloud storage using fog-backblaze gem. I dig little in to the source code. And now I'm thinking of, Instead of ruby handle IO objects is there any better way to totally delegate this on to a better library like curl, by just passing file path. So MRI threads get little weight. Do you know any libraries to do that?
## [10][Acts_as_taggable_on and JSONAPI::Resources for react-rails app, how should I collect information to read for my front-end?](https://www.reddit.com/r/ruby/comments/emcls0/acts_as_taggable_on_and_jsonapiresources_for/)
- url: https://www.reddit.com/r/ruby/comments/emcls0/acts_as_taggable_on_and_jsonapiresources_for/
---
There is Rails code for example (from the github page) which is used to find information

    ActsAsTaggableOn::Tag.most_used
    User.tagged_with("awesome").by_join_date
    @tom.find_related_skills # =&gt; [&lt;User name="Bobby"&gt;, &lt;User name="Frankie"&gt;]

But I cannot for the life of me try and get them into JSONAPI data using the JSONAPI::Resources gem.

Should I be coding in reactJS to find similar information? I could simply just display the raw data of the tags. It is doable (through tags and taggings), but I don't want to repeat myself if any methods are available.

Is there a way to run rails code from a react app?
