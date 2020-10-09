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
## [2][If interviewing at a JS-only place, would you say you like Typescript because it makes your code more type-safe?](https://www.reddit.com/r/typescript/comments/j7x355/if_interviewing_at_a_jsonly_place_would_you_say/)
- url: https://www.reddit.com/r/typescript/comments/j7x355/if_interviewing_at_a_jsonly_place_would_you_say/
---
Something like 75% of all errors are TypeErrors and all my portfolio items showcase mid level Typescript even though by experience I'm a junior. It's the one technology I feel strongly about because in the JS days I spent 30% of my time building the project, 70% of the time figuring out where the interpreter's dynamic type inference is throwing an error.

Do JS only companies take offense if you start raving about how much TS has reinforced your code and sense of clarity on what is being returned where? Should I downplay it if I know they're not big on TS? 

Intuition says yes especially given my lack of corporate experience (3 years learning/practicing at home though). PS: I know the best answer is to find a company that has TS but in these times I can't be picky.
## [3][They changed the compiler and now the code no longer compiles](https://www.reddit.com/r/typescript/comments/j7tkwz/they_changed_the_compiler_and_now_the_code_no/)
- url: https://www.reddit.com/r/typescript/comments/j7tkwz/they_changed_the_compiler_and_now_the_code_no/
---
I know how to fix the code but I don't understand the logic behind the new behaviour, I think this is a bug.

I really don't understand why the following code compiles under TS 3.7.5 but under TS 3.9.7 is an error

&amp;#x200B;

    interface Data {
      nullableArray?: number[]
      notNullable: number[];
    }
    
    function dummyFunction(data: Data) {
      data.nullableArray?.flat().map(x =&gt; x * 10); // &lt;-- no longer compile "object is possibly undefined
      data.notNullable.flat().map(x =&gt; x * 10);
    }
    

The error occurs on `x =&gt; x * 10` but the [Optional Chaining](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-7.html) ensures `flat()` is not called if `nullableArray` is `null` or `undefined`, so why now this code is incorrect?
## [4][Ts is garbage](https://www.reddit.com/r/typescript/comments/j7ywc6/ts_is_garbage/)
- url: https://www.reddit.com/r/typescript/comments/j7ywc6/ts_is_garbage/
---
ts is garbage , adding headache to js and produce a lot of bugs ... wasting the time to add types and dealing with it which reduce productivity.
## [5][How to have correct "this" type in callback function?](https://www.reddit.com/r/typescript/comments/j7eq51/how_to_have_correct_this_type_in_callback_function/)
- url: https://www.reddit.com/r/typescript/comments/j7eq51/how_to_have_correct_this_type_in_callback_function/
---
    function wrapFunction&lt;F extends (...fnArgs: unknown[]) =&gt; void&gt;(
      fn: F
    ): (...args: Parameters&lt;F&gt;) =&gt; void {
      return function (this: unknown, ...args: unknown[]): void {
        fn.apply(this, args)
      }
    }
    
    class SomeClass {
      myProperty = 'hey'
      myFunction: wrapFunction(function(){
        this.myProperty // "this" is not correct and "myProperty" is not known
      })
    }

wrapFunction could also be used in any class, so scope is quite dynamic.
## [6][How do I use `keyof typeof` on a Map in a class?](https://www.reddit.com/r/typescript/comments/j7h1nr/how_do_i_use_keyof_typeof_on_a_map_in_a_class/)
- url: https://www.reddit.com/r/typescript/comments/j7h1nr/how_do_i_use_keyof_typeof_on_a_map_in_a_class/
---
I have the following but don't want to type `checkEntity(entity: string)`s parameter with a string. I want it to only accept strings that are keys in `this.entities`.

I don't want to resort to hard coding `type entityKeys = "hero" | "zombie"`, there are a lot of keys in the actual program. I am missing the syntax to get what I need... `keyof typeof this.entities` is not a thing.

```typescript
export class World {
  entities;

  constructor() {
    this.entities = new Map&lt;string, [][]&gt;([
      ["hero", []],
      ["zombies", [[], [], []]],
    ]);
  }

  checkEntity(entity: string): void {
    if (this.entities.has(entity)) {
      console.log(this.entities.get(entity));
    }
  }
}
const world = new World();
world.checkEntity("doesnotexist"); // type error wanted
```
## [7][Best TypeScript Node.js API Postgres Starter Kit?](https://www.reddit.com/r/typescript/comments/j771br/best_typescript_nodejs_api_postgres_starter_kit/)
- url: https://www.reddit.com/r/typescript/comments/j771br/best_typescript_nodejs_api_postgres_starter_kit/
---
I see Microsoft has one for MongoDB. Not sure about it. Has too much for my test -- especially with Express.

I am looking for a starter kit that has TypeScript with best practices ready to go however having express (or not) is not a deal breaker.

Has anyone found anything similiar?
## [8][TS inferring object type as null](https://www.reddit.com/r/typescript/comments/j7b0e5/ts_inferring_object_type_as_null/)
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
    
## [9][Convincing typescript resources - talks, blog articles, studies, ...](https://www.reddit.com/r/typescript/comments/j7cw03/convincing_typescript_resources_talks_blog/)
- url: https://www.reddit.com/r/typescript/comments/j7cw03/convincing_typescript_resources_talks_blog/
---
I'm looking for any kind of resources that advocate typescript. Could be an interview with some CTO talking about all the benefits they had when switching form JS to TS. Could be a tech talk ("Airbnb could have prevented 38% of bugs"). Could be a blog article.

Thanks!
## [10][The most useless API you will ever see](https://www.reddit.com/r/typescript/comments/j6wtlj/the_most_useless_api_you_will_ever_see/)
- url: https://rand-api-docs.vercel.app
---

## [11][denogent - A build system built on Deno and TypeScript](https://www.reddit.com/r/typescript/comments/j72sjo/denogent_a_build_system_built_on_deno_and/)
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
