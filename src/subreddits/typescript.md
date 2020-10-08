# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
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
## [2][Best TypeScript Node.js API Postgres Starter Kit?](https://www.reddit.com/r/typescript/comments/j771br/best_typescript_nodejs_api_postgres_starter_kit/)
- url: https://www.reddit.com/r/typescript/comments/j771br/best_typescript_nodejs_api_postgres_starter_kit/
---
I see Microsoft has one for MongoDB. Not sure about it. Has too much for my test -- especially with Express.

I am looking for a starter kit that has TypeScript with best practices ready to go however having express (or not) is not a deal breaker.

Has anyone found anything similiar?
## [3][TS inferring object type as null](https://www.reddit.com/r/typescript/comments/j7b0e5/ts_inferring_object_type_as_null/)
- url: https://www.reddit.com/r/typescript/comments/j7b0e5/ts_inferring_object_type_as_null/
---
This is React code. Can anyone tell me why I'm facing issues with type assertion below, and how to fix it? i don't understand why `flashMessage` is inferred as a boolean type.

    // outside return ()
    const [flashMessage, setFlashMessage] = useState({});
    
    // inside return()
    &lt;SnackbarContent  
      message={ 'errors' in flashMessage as object ? 
        flashMessage.errors : flashMessage.message }
     /&gt;
    
    /* Conversion of type 'boolean' to type 'object' may 
    be a mistake because neither type sufficiently overlaps 
    with the other. If this was intentional, convert the expression 
    to 'unknown' first.ts(2352) */
    
    
    // removing 'as object' 
    // Property 'message' does not exist on type '{}'.ts(2339)
    
    // this also fails
      message={ 'errors' in flashMessage as Partial&lt;{ errors: string, message: string }&gt; ? flashMessage.errors : flashMessage.message }
    
## [4][The most useless API you will ever see](https://www.reddit.com/r/typescript/comments/j6wtlj/the_most_useless_api_you_will_ever_see/)
- url: https://rand-api-docs.vercel.app
---

## [5][denogent - A build system built on Deno and TypeScript](https://www.reddit.com/r/typescript/comments/j72sjo/denogent_a_build_system_built_on_deno_and/)
- url: https://www.reddit.com/r/typescript/comments/j72sjo/denogent_a_build_system_built_on_deno_and/
---
I want to introduce a project I've been working on for the past week\~two.

[https://github.com/areller/denogent](https://github.com/areller/denogent)

I was inspired by [https://nuke.build/](https://nuke.build/) which is a build system for .NET

But my goal is to create a more universal build system that can be used for Deno, NodeJS, .NET, Python or any other project that needs to have a build pipeline.

denogent comes as a CLI tool.

You can use `denogent create` to create a `build.ts` file, where you specify your build pipeline

    const compile = task('compile')
        .does(async ctx =&gt; {
            // compile program
        });
    
    const test = task('test')
        .dependsOn(compile)
        .does(async ctx =&gt; {
            // run tests
        });
    
    const buildImage = task('build image')
        .dependsOn(compile)
        .does(async ctx =&gt; {
            // build image
        });
    
    const publishImage = task('publish image')
        .dependsOn([test, buildImage])
        .does(async ctx =&gt; {
            // publish image
        });

You can then use `denogent run` to run the build pipeline locally, or use `denogent generate` to generate manifests for various CI systems. Right now, only GitHub Actions is supported but I plan to add integration with GitLab CI, Travis, Jenkins, etc...

Here is what currently supported

## Commands

* `denogent run` \- runs the pipeline that's defined in `build.ts` locally
* `denogent create` \- creates a `build.ts` file
* `denogent tasks` \- gets a list of tasks (`compile`, `test`, etc...)
* `denogent generate` \- generates manifest for a CI system (GH Actions, etc...)

## CI Systems Integration

* GitHub Actions

## API

* Runtime - run CLI commands, access secrets/arguments, etc...
* Git - run git commands
* FS - abstractions over the file system
* Docker - run docker commands (build, push) + run docker services (e.g. running a database containers along side the integration tests task)

## Build Kits (Integration with languages/platforms)

* Deno - running tests
* NodeJS - installing NodeJS

Right now, I plan to

1. Add more documentation
2. Continue refactoring the code and increasing coverage
3. Add support for more CI systems - Jenkins, GitLab CI, etc...
4. Expand existing build kits (e.g. add more APIs for NodeJS)
5. Add new build kits (e.g. dotnet, python, etc...)

## Sample Projects

1. [https://github.com/areller/denogent-samples-deno](https://github.com/areller/denogent-samples-deno)
2. [https://github.com/areller/denogent-samples-nodejs](https://github.com/areller/denogent-samples-nodejs)

Please share your thoughts about the project. Thank you :)
## [6][Convincing typescript resources - talks, blog articles, studies, ...](https://www.reddit.com/r/typescript/comments/j7cw03/convincing_typescript_resources_talks_blog/)
- url: https://www.reddit.com/r/typescript/comments/j7cw03/convincing_typescript_resources_talks_blog/
---
I'm looking for any kind of resources that advocate typescript. Could be an interview with some CTO talking about all the benefits they had when switching form JS to TS. Could be a tech talk ("Airbnb could have prevented 38% of bugs"). Could be a blog article.

Thanks!
## [7][How does typescript help?](https://www.reddit.com/r/typescript/comments/j7c2qi/how_does_typescript_help/)
- url: https://www.reddit.com/r/typescript/comments/j7c2qi/how_does_typescript_help/
---
I've never used typescript before but recently a lot of people said me to do it and that is better for grand scale projects but they didnt really told me why. So I'm asking you now, how does typescript help you?
## [8][I need help understanding how to share types between my front-end and back-end while using TypeORM](https://www.reddit.com/r/typescript/comments/j70br2/i_need_help_understanding_how_to_share_types/)
- url: https://www.reddit.com/r/typescript/comments/j70br2/i_need_help_understanding_how_to_share_types/
---
I'm using Express w/ Typescript for my backend and React w/ Typescript for my frontend. I want to share types between them, so I don't have to duplicate code and so I can easily pass things through my Express API. 

I know how to share interfaces between my Express app and React app, which I've done successfully  in the past when I was defining my types as interfaces and persisting them to Firestore.   


I'm now using Postgres for a new project, using TypeORM, and struggling to figure out how to share types.   


TypeORM has you use entity classes that are decorated with TypeORM decorators. This is working great for me on the backend, but using those types with my API and the client is causing me troubles.  


Issues I'm running into:  
1) Using the TypeORM entities in my client pulls in TypeORM and unnecessarily bloats the front end build.  
2) References in TypeORM entities are causing circular references and causing errors when trying to convert objects to JSON when returning data in my API.   


I'm kinda new to Typescript, so I feel like I'm missing something basic in  terms of best practices.
## [9][Typescript or runtime responsibility?](https://www.reddit.com/r/typescript/comments/j6un7q/typescript_or_runtime_responsibility/)
- url: https://www.reddit.com/r/typescript/comments/j6un7q/typescript_or_runtime_responsibility/
---
Hi everyone 👋I'm writing a React component for render collections.

Through properties

* you can provide a custom itemExtractor to adapt item type
* you can provide a custom itemRender to change item component

some use cases:

    // Suppose a component(Collection) that renders a set of items(CollectionItem[])
    // We use the default renderItem and the default extractItem (noop)... OK
    const items: CollectionItem[] = [];
    return &lt;Collection items={items} /&gt;;
    
    // Suppose that items type differs from CollectionItem[]
    // We implements extractItem to get collectionItem type. Use default renderItem... OK
    const items2: OtherType[] = [];
    return &lt;Collection items={items} extractItem={(other: OtherType) =&gt; ({ title: other.name })} /&gt;;
    
    // Suppose that we need to render a different type
    // Implements renderItem to render special fields. Use default extractItem (noop)... OK
    const items3: OtherType[] = [];
    return &lt;Collection items={items3} renderItem={(other: OtherType) =&gt; &lt;p&gt;{other.name}&lt;/p&gt;} /&gt;;

**Thats works fine, but...  It is possible to expect an implementation for extractItem or renderItem if items type differs from CollectionItem?**for example&lt;Collection items={otherTypeItems} **\*expected\_extractItem / expected\_renderItem type error?\***/&gt;;

    Implementation:
    
    export interface CollectionProps&lt;In, Out&gt; { 
        items: In[]; 
        extractItem: (data: In) =&gt; Out; 
        renderItem: (data: Out, index: number) =&gt; React.ReactNode; 
    }
    
    export const Collection = function &lt;In = CollectionItem, Out = CollectionItem&gt;(
        props: CollectionProps&lt;In, Out&gt;
    ) {
        if (!props.items.length) return null;
        return props.items.map((item, index) =&gt; 
       props.renderItem(props.extractItem(item), index));
    };
    
    Collection.defaultProps = {
        extractItem: (i: CollectionItem) =&gt; i,
        renderItem: (data: CollectionItem, index: number) =&gt; (
            &lt;DefaultCollectionItem key={index} data={data} /&gt;)
    };

Is this Typescript responsability or should I check this on runtime? 🤔

Any help regarding this will be appreciated  🙏
## [10][I wrote my first Typescript React App. Would love some feedback if anyone is interested in giving it. I will post the link in the comments.](https://www.reddit.com/r/typescript/comments/j6s02p/i_wrote_my_first_typescript_react_app_would_love/)
- url: https://www.reddit.com/r/typescript/comments/j6s02p/i_wrote_my_first_typescript_react_app_would_love/
---

## [11][How to use jQuery Plugins like owl-carousel in typeScript?](https://www.reddit.com/r/typescript/comments/j6npge/how_to_use_jquery_plugins_like_owlcarousel_in/)
- url: https://www.reddit.com/r/typescript/comments/j6npge/how_to_use_jquery_plugins_like_owlcarousel_in/
---
Hi!  
I am new to typeScript. I can't figure out how to install the owl-carousel in my typeScript project.

Any help regarding this will be appreciated! Thanks in advance.
