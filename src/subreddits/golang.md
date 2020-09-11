# golang
## [1][We are the Go Time podcast. Ask us anything. (AMA)](https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/)
- url: https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/
---
Hi everyone! I'm Jon Calhoun, one of the panelists on the Go Time podcast. For those of you unfamiliar - it is a Go podcast that we record live every Tuesday at 3pm ET. We usually have a guest or two on each episode, and discuss a specific topic (defer, testing, databases, infra, etc). You can check it out here: &lt;https://changelog.com/gotime&gt;

This coming episode we want to try something a little different - we want to make a Q&amp;A episode. There are two reasons for doing this:

1. We are hoping this inspires topics for future episodes.
2. We want a venue to discuss questions that don't really fit into an entire episode on their own.

To make this happen we would like everyone here to post any Go-related questions that you would like us to discuss on air, and we will try to get to as many as possible. I'll also try to type up answers here while we discuss them on the episode.

We will be answering questions live tomorrow, Tuesday, Sep 8. We will repeat questions on air, and since we record live you can join in on the Gophers Slack to ask follow-up questions or to elaborate on questions.
## [2][Project introduces breaking changes to library - bumps minor versioning instead of major version because go module's v2+ handling is so quirky](https://www.reddit.com/r/golang/comments/iqoiok/project_introduces_breaking_changes_to_library/)
- url: https://www.reddit.com/r/golang/comments/iqoiok/project_introduces_breaking_changes_to_library/
---
I'm very unhappy with go modules and this project shying away from bumping the major version in favor of bumping the minor version while introducing huge breaking changes just speaks for itself.

[https://github.com/gofiber/fiber/issues/736#issuecomment-690750255](https://github.com/gofiber/fiber/issues/736#issuecomment-690750255)

What do you think?

I just can't stop thinking that in other ecosystems people are not scared to stick to semantic versioning, so this might be a "smell" that Go has, that is not present in other languages.

Oh and please please *don't you all jump into that Github issue and tell the maintainers how they should run their project*. I don't want to cause them any trouble by posting here!
## [3][Goyave v3: a huge release for the REST API framework](https://www.reddit.com/r/golang/comments/iq4vw0/goyave_v3_a_huge_release_for_the_rest_api/)
- url: https://www.reddit.com/r/golang/comments/iq4vw0/goyave_v3_a_huge_release_for_the_rest_api/
---
Today, I am releasing Goyave v3, which is a huge step forward for the framework and represents months of work. Goyave is an opinionated REST API framework with tons of features to help you focus on your business logic.  


**Repo link:** [https://github.com/System-Glitch/goyave](https://github.com/System-Glitch/goyave)

**Docs link:** [https://system-glitch.github.io/goyave/guide/](https://system-glitch.github.io/goyave/guide/)

## Discord

Come and learn about the framework, stay updated on the progress, be notified when there is a new release, get help, suggest new features or changes, contribute to the project, and more!

&amp;#x200B;

https://preview.redd.it/9g22qm0f3cm51.png?width=320&amp;format=png&amp;auto=webp&amp;s=f292843a29bfa2e9b0b1eb01ab683837f3fa3967

Click here to join: [https://discord.gg/mfemDMc](https://discord.gg/mfemDMc)

## Key changes

🌞 Quality of life improvements, such as cleaner route definition  
⚙ Revamped configuration system  
📘 GORM v2  
💨 Performance improvements

## Full Changelog

* Changed conventions:
   * `validation.go` and `placeholders.go` moved to a new `http/validation` package.
   * Validation rule sets are now located in a `request.go` file in the same package as the controller.

**Motivation**: *Separating the requests in another package added unnecessary complexity to the directory structure and was not convenient to use. Package naming was far from ideal with the "request" suffix. Moving requests to the same package as the controller is more intuitive and requires less imports and makes route definition cleaner and easier.*  


* Validation system overhaul, allowing rule sets to be parsed only once instead of every time a request is received, giving better overall performance. This new system also allows a more verbose syntax for validation, solving the comma rule parameter value and a much easier use in your handlers.
   * Rule functions don't check required parameters anymore. This is now done when the rules are parsed at startup time. The amount of required parameters is given when registering a new rule.
   * Optimized regex-based validation rules by compiling expressions once.
   * A significant amount of untested cases are now tested.
   * The following rules now pass if the validated data type is not supported: `greater_than`, `greater_than_equal`, `lower_than`, `lower_than_equal`, `size`.
   * Type-dependent rules now try to determine what is the expected type by looking up in the rule set for a type rule. If no type rule is present, falls back to the inputted type. This change makes it so the validation message is correct even if the client didn't input the expected type.
   * Fixed a bug triggering a panic if the client inputted a non-array value in an array-validated field.

**Motivation**: *The validation system had a lot of room for improvement when it comes to performance, as* `RuleSet` *were parsed every time a request was received. Moving this process out of the request life-cycle to execute it only once saves a good amount of execution time. Moreover, any handler who would want to read the rules applied to the current request needed to parse them too, which was inconvenient and not effective. With a structure containing everything you need, making middleware interacting with the request's rules is much easier.*  


* Routing has been improved by changing how validation and route-specific middleware are registered. The signature of the router functions have been simplified by removing the validation and middleware parameters from `Route()`, `Get()`, `Post()`, etc. This is now done through two new chainable methods on the `Route`: `route.Validate()` and  `route.Middleware()`.

**Motivation**: *In the original design, the validation parameter was included in the main route definition function because most routes were expected to be validated, which turned out not to be the case. In a typical CRUD, only the create and update actions are validated, which made the route definition dirty and filled with* `nil` *parameters. Separating the rules and middleware definition is more in line with their optional nature and makes routes definition cleaner and more readable, although sometimes slightly longer.*  


* Log `Formatter` now receive the length of the response (in bytes) instead of the full body.

**Motivation:** *Keeping in memory the full response has an important impact on memory when sending files or large responses. Using the response content in a log formatter is also a marginal use-case which doesn't justify the performance loss described previously. It is still possible to retrieve the content of the response by writing your own chained writer.*  


* Configuration system has been revamped.
   * Added support for tree-like configurations, allowing for better categorization. Nested values can be accessed using dot-separated path.
   * Improved validation: nested entries can now be validated too and all entries can have authorized values. Optional entries can now be validated too.
   * Improved support for slices. The validation system is also able to check slices.
   * Entries that are validated with the `int` type are now automatically converted from `float64` if they don't have decimal places. It is no longer necessary to manually cast `float64` that are supposed to be integers.
   * More openness: entries can be registered with a default value, their type and authorized values from any package. This allows config entries required by a specific package to be loaded only if the latter is imported.
   * Core configuration has been sorted in categories. This is a breaking change that will require you to update your configuration files.
   * Entries having a `nil` value are now considered unset.
   * Added accessors `GetInt()` and `GetFloat()`.
   * Added slice accessors: `GetStringSlice()`, `GetBoolSlice()`, `GetIntSlice()`, `GetFloatSlice()`
   * Added `LoadFrom()`, letting you load a configuration file from a custom path.
   * Added the ability to use environment variables in configuration files.
   * Bug fix: `config.IsLoaded()` returned `true` even if config failed to load.
   * `maxUploadSize` config entry now supports decimal places.

**Motivation:** *Configuration was without a doubt one of the weakest and inflexible feature of the framework. It was possible to use objects in custom entries, but not for core config, but it was inconvenient because it required a lot of type assertions. Moreover, core config entries were not handled the same as custom ones, which was a lack of openness. Hopefully, this revamped system will cover more potential use-cases, ease plugin development and allow you to produce cleaner code and configuration files.*  


* Database improvements
   * Goyave has moved to [GORM v2](https://gorm.io/). Read the [release note](https://gorm.io/docs/v2_release_note.html) to learn more about what changed.
   * Protect the database instance with mutex.
   * `database.Close()` can now return errors.
   * Added [database connection initializers](https://system-glitch.github.io/goyave/guide/basics/database.html#connection-initializers).
   * Added the ability to regsiter new SQL dialects to use with GORM.
   * Use `utf8mb4` by default in database options.
   * Added a short alias for `database.GetConnection()`: `database.Conn()`.
   * Factories now use batch insert.
   * Factories now return `interface{}` instead of `[]interface{}`. The actual type of the returned value is a slice of the the type of what is returned by your generator, so you can type-assert safely.
* Status handlers improvements
   * Export panic and error status handlers so they can be expanded easily.
   * Added `goyave.ValidationStatusHandler()`, a status handler for validation errors. Therefore, the format in which validation errors are sent to the client can be customized by using your own status handler for the HTTP status 400 and 422.
* `goyave.Response` improvements
   * `response.Render` and `response.RenderHTML` now execute and write the template to a `bytes.Buffer` instead of directly to the `goyave.Response`. This allows to catch and handle errors before the response header has been written, in order to return an error 500 if the template doesn't execute properly for example.
   * Added `response.GetStacktrace()`, `response.IsEmpty()` and `response.IsHeaderWritten()`.
   * Re-organised the `goyave.Response` structure fields to save some memory.
   * Removed deprecated method `goyave.CreateTestResponse()`. Use `goyave.TestSuite.CreateTestResponse()` instead.
* Recovery middleware now correctly handles panics with a `nil` value.
* Test can now be run without the `-p 1` flag thanks to a lock added to the `goyave.RunTest` method. Therefore, `goyave.TestSuite` still **don't run in parallel** but are safe to use with the typical test command.
* Cache the regex used by `helper.ParseMultiValuesHeader()` to improve performance. This also improves the performance of the language middleware.
* Bug fix: data under validation wasn't considered from JSON payload if the content type included the charset.
* The Gzip middleware will now skip requests that have the `Upgrade` HTTP header set to any value.
* `response.String()` and `response.JSON()` don't write header before calling `Write` anymore. This behavior prevented middleware and chained writers to alter the response headers.
* Added `goyave.PreWriter` interface for chained writers needing to alter headers or status before they are written.
   * Even if this change is not breaking, it is recommended to update all your chained writers to call `PreWrite()` on their child writer if they implement the interface.
   * Thanks to this change, a bug with the gzip middleware has been fixed: header `Content-Length` wasn't removed, resulting in false information sent to the clients, which in turn failed to decompress the response.

## Upgrade guide

You can find the upgrade guide [here](https://system-glitch.github.io/goyave/guide/upgrade-guide.html).
## [4][Http proxy for caching file on disk](https://www.reddit.com/r/golang/comments/iqqf9g/http_proxy_for_caching_file_on_disk/)
- url: https://www.reddit.com/r/golang/comments/iqqf9g/http_proxy_for_caching_file_on_disk/
---
Hi, for a project, we need a way to cache files ( \~240mb) with a web proxy. I was planning to use go for that, the web proxy part is fine. I'm just wondering if you know a caching library/server/database which will be able to handle large file like that? One solution would be to create chunk in bigcache or badger, what do you think?
## [5][Build fault tolerant applications with Cassandra API for Azure Cosmos DB](https://www.reddit.com/r/golang/comments/iqq5cj/build_fault_tolerant_applications_with_cassandra/)
- url: https://medium.com/abhishek1987/build-fault-tolerant-applications-with-cassandra-api-for-azure-cosmos-db-d83f5f1fffb7?source=friends_link&amp;sk=c74d33ade6d407d756abba390b2cb4a3
---

## [6][chromedp fails to start a new context](https://www.reddit.com/r/golang/comments/iqporw/chromedp_fails_to_start_a_new_context/)
- url: https://www.reddit.com/r/golang/comments/iqporw/chromedp_fails_to_start_a_new_context/
---
Hey guys I have been trying to automate a task on my browser's machine using [chromedp](https://github.com/chromedp/chromedp).

I started with the examples on the repo however it always fails:

&gt;unexpected fault address 0x7f7a36461000 fatal error: fault \[signal SIGBUS: bus error code=0x2 addr=0x7f7a36461000 pc=0x53a9d5\]

As I understand chromedp comes with headless pre-compiled version of chrome. I have browsed the API documentation to see if it possible to pass a port on which the dev tools is listening, no success.

However when I tried [mafredri/cdp](https://github.com/mafredri/cdp) it worked. I came to understanding that chromedp is failing to start the chrome it comes with. Why this is happening?

My env:

&gt;Google Chrome 85.0.4183.102  
&gt;  
&gt;Ubuntu 18.04.4 LTS  
&gt;  
&gt;go version go1.13.5 linux/amd64

Any help would be truly appreciated.
## [7][Caddy Server Acquired By Apilayer](https://www.reddit.com/r/golang/comments/iq861b/caddy_server_acquired_by_apilayer/)
- url: https://www.ardanlabs.com/news/2020/08/caddy-server-is-acquired/
---

## [8][Sugolver - A Sudoku Solver written in Go](https://www.reddit.com/r/golang/comments/iq8kvo/sugolver_a_sudoku_solver_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/iq8kvo/sugolver_a_sudoku_solver_written_in_go/
---
**Sugolver** is a Sudoku solver written in Go.

This [article](http://thomas-joly.com/index.php/2020/08/31/sugolver/) is describing the data structures and the algorithms used in this project.

The project is hosted on Github [sugolver](https://github.com/lunatikub/sugolver).

**Have a good reading !**
## [9][converting excel to struct](https://www.reddit.com/r/golang/comments/iqmry6/converting_excel_to_struct/)
- url: https://www.reddit.com/r/golang/comments/iqmry6/converting_excel_to_struct/
---
Hi,

I have a use case where I need to read an excel and process it , which will be uploaded by the user.

I am using  github.com/360EntSecGroup-Skylar/excelize this package to read the excel , the problem here is that to reach value of each cell I need to read the row as array and cell with the index , this will hardly bind the column sequence and will fail if the sequence is changed, I was thinking is there a way to read excel directly to a struct.

I found this [https://github.com/szyhf/go-excel](https://github.com/szyhf/go-excel) , but the issue here is it requires the file path to open it,

I receive the file in request which is  \*multipart.FileHeader  format , here for this I need to first create a file and save it somewhere and then process it .

is there any other way where I can convert  excel into struct .
## [10][Sqlize: generate migration MySQL Compatible SQL](https://www.reddit.com/r/golang/comments/iqgms3/sqlize_generate_migration_mysql_compatible_sql/)
- url: https://github.com/sunary/sqlize
---

## [11][Extension interfaces and proxies](https://www.reddit.com/r/golang/comments/iqhzrb/extension_interfaces_and_proxies/)
- url: https://www.reddit.com/r/golang/comments/iqhzrb/extension_interfaces_and_proxies/
---
I'm working through the design of a package for CAS.  There is a main interface, but also 2 extension interfaces.  The design question that is bothering me at the moment is that the package needs to support proxies.  For example, there is an HTTP server and client to provide remote access to a CAS.  The HTTP client and server also support the extension interfaces, even though they may simply return an error if the methods are unsupported by the underlying implementation.

I'm kind of reaching for a QueryInterface approach for proxies where I can provide the extension interface, but only if supported.  Has anyone hit a similar problem with extension interfaces?  How did you solve it?
