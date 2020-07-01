# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][TIL What TypeScript Is](https://www.reddit.com/r/typescript/comments/hj8azp/til_what_typescript_is/)
- url: https://www.reddit.com/r/typescript/comments/hj8azp/til_what_typescript_is/
---
So last night, I finally watched a short intro on what exactly typescript is, and... Wow.

I wish I had known about this years ago; as someone who loves languages like C++, and never really was a fan of JS, TS makes the web world so much more inviting.

Anyway, happy to be here, sad it took me so long to learn about it.
## [3][How to check if a whole number (1.0) is a double in TypeScript](https://www.reddit.com/r/typescript/comments/hj9qtk/how_to_check_if_a_whole_number_10_is_a_double_in/)
- url: https://www.reddit.com/r/typescript/comments/hj9qtk/how_to_check_if_a_whole_number_10_is_a_double_in/
---
I have a library for developers with a method like this:

        function add (value: number | string) {
          if (typeof value === 'number') {
            if (value % 1 == 0) {
              // post to database as an int
            } else {
              // post to database as a double
            }
          }
          ...
        }

The problem here arises when *value* is something like 2.0 or 10.0, this gets treated as an int. However, I need those values to be passed as doubles. How can I check for these 'whole numbers' (10.0 or 2.0) to be doubles?
## [4][Typescript Infer types from a tuple in React render props](https://www.reddit.com/r/typescript/comments/hj9b5l/typescript_infer_types_from_a_tuple_in_react/)
- url: https://www.reddit.com/r/typescript/comments/hj9b5l/typescript_infer_types_from_a_tuple_in_react/
---
I want to infer the types from my fetchData array so that the parameters I pass to my children would be properly typed.

    import * as React from 'react';
    
    interface DataSource {
      dataSource: () =&gt; Promise&lt;unknown&gt;,
    }
    
    interface Data {
      data: unknown,
    }
    
    interface Props{
      fetchData: DataSource[],
      children: (...data: any[]) =&gt; React.ReactNode;
    }
    
    interface State {
      data: Data[];
    }
        
    const Placeholder : React.FC&lt;Props&gt; = (props) =&gt; {
      const [state, setState] = React.useState&lt;State&gt;({
        data: [],
      });
    
      const getAllData = async () =&gt; {
        const result = await Promise.all(props.fetchData.map((item) =&gt; item.dataSource()));
        return result as Data[];
      };
    
      const fetchData = async () =&gt; {
          const data = await getAllData();
          setState({
              data,
            });
      };
      React.useEffect(() =&gt; {
        fetchData();
      }, []);
    
      const {
        data
      } = state;
    
      const { children } = props;
    
      return (
          {children(...data)}
      );
    };

In the below example:

         // const fetchCategories: () =&gt; Promise&lt;Category[]&gt;
         // const fetchRoles: () =&gt; Promise&lt;Role[]&gt;
    
          &lt;PlaceHolder fetchData={[{dataSource: fetchCategories }, {dataSource: fetchRoles}]}&gt;
            {(categories /* right now 'categories' type is 'any' but I want to infer type 
             Category[] from 'fetchCategories'  */,
             roles /* same as categories */) =&gt; (
              &lt;CreateCategory categories={categories} roles={roles} /&gt;
            )}
          &lt;/PlaceHolder&gt;
## [5][Help digging a return type out of a big library (Google Cloud)](https://www.reddit.com/r/typescript/comments/hj6evi/help_digging_a_return_type_out_of_a_big_library/)
- url: https://www.reddit.com/r/typescript/comments/hj6evi/help_digging_a_return_type_out_of_a_big_library/
---
[https://googleapis.dev/nodejs/text-to-speech/latest/v1.TextToSpeechClient.html](https://googleapis.dev/nodejs/text-to-speech/latest/v1.TextToSpeechClient.html)

&gt;synthesizeSpeech(request, optionsopt) → {Promise}  
&gt;  
&gt;Synthesizes speech synchronously: receive results after all text input has been processed.

The library comment on this method:

    * @returns {Promise} - The promise which resolves to an array.
         *   The first element of the array is an object representing [SynthesizeSpeechResponse]{@link google.cloud.texttospeech.v1.SynthesizeSpeechResponse}.
         *   The promise has a method named "cancel" which cancels the ongoing API call.
         */

And on that page:

[https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#google.cloud.texttospeech.v1.SynthesizeSpeechResponse](https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#google.cloud.texttospeech.v1.SynthesizeSpeechResponse)

&gt;SynthesizeSpeechResponse  
&gt;  
&gt;The message returned to the client by the SynthesizeSpeech method.

&amp;#x200B;

My code so far:

    const TextToSpeechClient = require("@google-cloud/text-to-speech").TextToSpeechClient;
    
    // Cannot find namespace 'TextToSpeechClient'.ts(2503)
          : Promise&lt;TextToSpeechClient.SynthesizeSpeechResponse&gt; 

Full method below. I got close but I'm still rough with digging around an API, trying to extract these obscure types.

I want to learn how to do it though. Marking `any` is easy, specifying the correct type on a fetch especially is a lot more insightful and easy to backtrace later.

Thanks for any pointers on getting it right.

&amp;#x200B;

Full method (PS: I plan to split the fetch and save next. Maybe also implement a named `options` interface)

       protected async fetchAndWriteAudio(options: {
          input : { text : string }
          , voice : { languageCode : string, ssmlGender: voiceGender }
          , audioConfig : { audioEncoding : string }
       }, fileNameAndPath: string)
          : Promise&lt;TextToSpeechClient.SynthesizeSpeechResponse&gt; {
    
          const textToSpeech = new TextToSpeechClient();
          const writeFileAsync = util.promisify(writeFile);
    
          try {
             const [audioResponse] = await textToSpeech.synthesizeSpeech(options);
             await writeFileAsync(fileNameAndPath, audioResponse.audioContent);
          }
          catch(error) { console.log(error); }
       }
## [6][Compiler gives no warning with wrong extended class fed into a constructor](https://www.reddit.com/r/typescript/comments/hj67q0/compiler_gives_no_warning_with_wrong_extended/)
- url: https://www.reddit.com/r/typescript/comments/hj67q0/compiler_gives_no_warning_with_wrong_extended/
---
Is there a way for the compiler to give warning/error because I am giving Y class instead of Z class on P constructor?

    class X { }
    
    class Y extends X { } 
    class Z extends X { }
    
    const y = new Y()
    
    class P { 
        constructor(private z: Z) { }
    
        get get(): Z { 
            return this.z  
        } 
    }
    const p = new P(y) 
    console.log(p.get)

&amp;#x200B;
## [7][Prints a dependency graph in dot format for your typescript project.](https://www.reddit.com/r/typescript/comments/hj7hct/prints_a_dependency_graph_in_dot_format_for_your/)
- url: https://github.com/PSeitz/ts-dependency-graph
---

## [8][A module "manager" for Deno](https://www.reddit.com/r/typescript/comments/hixcao/a_module_manager_for_deno/)
- url: https://www.reddit.com/r/typescript/comments/hixcao/a_module_manager_for_deno/
---
Hello, I'm Rivier

My team and I we've been working in a way to manage the packages from deno. Ok we know that deno doesn't need a package manager 'cuz the way to bring them to a project if from an url but that's the point Trex is a way to point and manage more easy the philosophy of the import\_maps in deno: [Trex](https://github.com/crewdevio/Trex) 

any collaborations and feedbacks are welcome :)
## [9][Import from package with dependencies](https://www.reddit.com/r/typescript/comments/hijylv/import_from_package_with_dependencies/)
- url: https://www.reddit.com/r/typescript/comments/hijylv/import_from_package_with_dependencies/
---
Hello,

I am working an an application that is deployed as some microservices that are mostly written with Typescript and run with Node. We have split some functionality that is used in many backends to small NPM packages that those backends consume. Some of these packages provide abstractions for others and all is well so far.

But we encountered a problem on some packages that is keeping us busy for a while now. [I'll asked on stack overflow](https://stackoverflow.com/questions/61363752/typescript-import-type-from-package-without-dependencies-from-other-types-in-th) about this some months ago with no avail. So I hope somebody here can help us or provide us with some ideas.

&amp;#x200B;

To outline the problem a little bit:

Consider a package that provides a class A and class B. Class A is using an external dependency that also some other packages use. Class B is not using this dependency. To avoid having the dependency installed multiple times and to not having to install the dependency if you don't need it, it is set as peerDependency.

Now when I consume class A in a project and also make use of the dependency I install it and everything works fine. But if in another project I only need class B one could argue, that the dependency won't bring anything of value here, thus should not be installed. However if we try to use class B in a project, Typescript would always complain about the missing dependency.

&gt;error TS2307: Cannot find module 'x' or its corresponding type declarations.

&amp;#x200B;

Note: we use an index.ts for all our packages, where all classes, types and interfaces that should be used outside of the package itself are exported.

&amp;#x200B;

Now, is there any way, apart from splitting class B in it's own package or having to set "skipLibCheck": true in every tsconfig file, to avoid the typescript error?

&amp;#x200B;

Edit:

To clarify a little bit and to give some real world examples:

The original SO post was about a logger abstractions. Since we have microservices we want our logging to be centralized, While the deployment environment allows us to scrape all those logs and we have a tool to view them centralized it was our intention to unify the way the backends or microservices log. We also wanted to implement a correlation ID to follow requests through multiple backends and packages in the logs without the need to handle the request object (that has the ID of the current request stored) down to the deepest function of some sub package.

For this I created a logger abstraction that provides a factory where each class (we use mostly classes but it would work for functions too) can retrieve an instance of the logger to use. This logger would have the correlationID already set, so all those backends, packages and whatever don't need to mind about it. The abstractions also abstract what logger is even used (since we replaced one logger with another, this was also a goal of the whole action) so those packages don't really care what logger is even installed.

So each package would install the logger abstractions and can then rely on it to retrieve a logger instance and write logs to it. For example we have a database service that wraps all the connection handling. Every backend that needs to do some database operations has this service installed. Now the idea is, the database service installs the logger abstractions, gets the instance of the logger and write logs (for example if you want to log queries or just the init/shutdown logs). The backend would then install the database service and also a concrete implementation of the abstract logger factory.

Initially this concrete implementation was also part of the logger package. This implementation needs a real logger (in this instance it was pino) and thus has a dependency on it. And that was were the problems began, since the database service won't transpile and complain that pino was not installed, even though it used only the abstract logger without any dependency.

We "solved" this problem, by extracting the concrete logger implementation into its own package. But that was not really a solution but more a workaround.

&amp;#x200B;

Right now I'm trying to build a centralized metrics service that backends can use to unify metrics we expose for our Prometheus server. For example the database service has now a method that returns stats about the connection pool (how many connections are idle/waiting/active). Now every backend can call this function and then create Gauges in the Prometheus client. Again much duplicate code so I build an intermediate service that wraps the Prometheus stuff. It also exposes a ScrapableService interface, that other services like the db service can implement and provide unified methods to be scraped later.

This is where again the problem begins. As the metrics are optional, I would rather not install the metrics service (and it's prom-client dependency) everywhere the database service is used. The db service itself only needs the interface which doesn't depend on anything. This interface only adds 2 public methods to the service that can but does not have to be used.

But I can't figure out how to make this optional. A backend that needs the db service but doesn't expose any metrics and therefore has no metrics service would fail to build, because the dependency is missing. If we add metrics-service as dependency of the database service it would be duplicated on all backends that install and use the metrics service ( I know NPM dedupes modules if they have the same version).

Again we could solve this by extracting the interface, but since it's only purpose is to glue together the metrics-service with re-usable services in the backend, it seems a little bit too overkill to have a package with only an interface in it.

I hope this (somewhat long) explanation provides some insight and explains why we have these problems.
## [10][Is it worth setting up an enum for a single use?](https://www.reddit.com/r/typescript/comments/hilkz9/is_it_worth_setting_up_an_enum_for_a_single_use/)
- url: https://www.reddit.com/r/typescript/comments/hilkz9/is_it_worth_setting_up_an_enum_for_a_single_use/
---
I have a method that should process text/audio differently depending on whether the input was a sentence or word. I could setup a boolean like `isWord` but that is the least clear or extensible (phrase support may be added). 

Probably the JS option would be to use string option values.

In Typescript I have a file `minorTypes.ts` that I have used to import little enums. But I'm not sure if this is a good pattern or not, if the enum is only used in one method. So I wanted to ask about it. What do you guys think?
## [11][How do I disable specific ESLint rules?](https://www.reddit.com/r/typescript/comments/hiyww7/how_do_i_disable_specific_eslint_rules/)
- url: https://www.reddit.com/r/typescript/comments/hiyww7/how_do_i_disable_specific_eslint_rules/
---
I set up a new project and I'm already getting things like `img elements must have an alt prop`. This is a project for work and accessibility is not a concern.

I don't have an `.esconfig` file or anything
