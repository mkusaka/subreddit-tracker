# golang
## [1][Olric v0.3.0-beta.6 is out: Distributed cache and in-memory key/value data store. It can be used both as an embedded Go library and as a language-independent service.](https://www.reddit.com/r/golang/comments/jc7gfg/olric_v030beta6_is_out_distributed_cache_and/)
- url: https://github.com/buraksezer/olric/releases/tag/v0.3.0-beta.6
---

## [2][How to customize Go's HTTP client](https://www.reddit.com/r/golang/comments/jc5a1i/how_to_customize_gos_http_client/)
- url: https://rafallorenz.com/go/customize-http-client/
---

## [3][How do you log?](https://www.reddit.com/r/golang/comments/jc7mkx/how_do_you_log/)
- url: https://www.reddit.com/r/golang/comments/jc7mkx/how_do_you_log/
---
How do you people manage logger throughout whole application?

Been using Uber's zap for past 4 years, initially via dependency injection:

    type myImpl struct {
        logger *zap.Logger
    }
    
    func (impl myImpl) myFunct() {
        impl.logger.Info("Hello world!")
    }

But for past year [via context](https://github.com/cantor-systems/zapctx):

    func myFunc(ctx context.Context) {
        zapctx.Info(ctx, "Hello world!)
    }

The latter approach works, however it implies global context and whenever you need to access logger, you need to pass context to that function.

However, I have grown a bit doubtful of that and I am not sure if this is the way to go in the long term as for example I would like to add additional field "service", that contains the struct name, for all logger method calls in that struct implementation, however that is not doable with context logger, only with dependency inversion.

So I was wondering, how do you do it?. What are your approaches to logging?
## [4][Is it a bad practice to assign variable to other and then reassign it back?](https://www.reddit.com/r/golang/comments/jc7att/is_it_a_bad_practice_to_assign_variable_to_other/)
- url: https://www.reddit.com/r/golang/comments/jc7att/is_it_a_bad_practice_to_assign_variable_to_other/
---
I've got a quick question more in terms of code quality, but it's a go code, so I decided to post it here.

I have such piece of code:

    shortVar := some.Nested.Struct.Field
    ... // do some processing (around 40 lines if that shortVar)
    ...
    some.Nested.Struct.Field = shortVar // is it a bad practice?

I'm wondering if it's a bad practice to do so? Should I do it like in the above example or should I use `some.Nested.Struct.Field` everywhere?
## [5][Looking for project to contribute bugfixes](https://www.reddit.com/r/golang/comments/jc5b2o/looking_for_project_to_contribute_bugfixes/)
- url: https://www.reddit.com/r/golang/comments/jc5b2o/looking_for_project_to_contribute_bugfixes/
---
I've been teaching my son programming and go and he has progressed to the point that I'm trying to find some actual real world problems for him to swing at. Can anyone suggest some good open source projects that have a decemt sized backlog of issues I could unleash him at?
## [6][I wrote a tool to dump SSL Informations and test TLS Versions like SSLShopper e SSL Labs](https://www.reddit.com/r/golang/comments/jc96nk/i_wrote_a_tool_to_dump_ssl_informations_and_test/)
- url: https://github.com/msfidelis/cassler
---

## [7][Returning gRPC error status WithDetails, but not showing up in body?](https://www.reddit.com/r/golang/comments/jc0521/returning_grpc_error_status_withdetails_but_not/)
- url: https://www.reddit.com/r/golang/comments/jc0521/returning_grpc_error_status_withdetails_but_not/
---
I'm currently following this:

[https://medium.com/utility-warehouse-technology/advanced-grpc-error-usage-1b37398f0ff4](https://medium.com/utility-warehouse-technology/advanced-grpc-error-usage-1b37398f0ff4)

I'm able to correctly add the details to the Error, but when I make a request to my endpoint via Postman. The body is just:  
{  
 "error": "This is an error!"  
}

What steps should I follow to get the details to show up in the body as well?
## [8][Need help with GORM!](https://www.reddit.com/r/golang/comments/jc3bj6/need_help_with_gorm/)
- url: https://www.reddit.com/r/golang/comments/jc3bj6/need_help_with_gorm/
---
So, I am very new to golang and I am working on a personal project of mine. So I am using gorm for db stuffs and whenever I do any operation gorm automatically  populate created, modified and deleted column depending on what operation I do. Now I want to also populate some custom columns of mine whenever any operation is done. Does anyone knows best and easy solution for this?
## [9][[Academic] Software Development Survey (Software Developers)](https://www.reddit.com/r/golang/comments/jc61ak/academic_software_development_survey_software/)
- url: /r/SampleSize/comments/jbsju0/academic_software_development_survey_software/
---

## [10][Free offline solution to convert PDFs into audiobooks -](https://www.reddit.com/r/golang/comments/jbmmga/free_offline_solution_to_convert_pdfs_into/)
- url: https://github.com/Harry-027/go-audio
---

