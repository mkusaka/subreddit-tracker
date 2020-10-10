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
## [2][How to view JSDoc (or at least signature) for untyped JS dependencies? Tried both WebStrorm and VisualCode](https://www.reddit.com/r/typescript/comments/j8hj0y/how_to_view_jsdoc_or_at_least_signature_for/)
- url: https://i.redd.it/2q0surn3j8s51.png
---

## [3][Typescript project starter: npm-ready, plus testing and CLI templates](https://www.reddit.com/r/typescript/comments/j88pwf/typescript_project_starter_npmready_plus_testing/)
- url: https://github.com/bscotch/typescript-template
---

## [4][Jeff and Dom Make A Game - a coworker and his friend have been making a multiplayer game in TypeScript the past 5 months and streaming the whole thing](https://www.reddit.com/r/typescript/comments/j8b8nv/jeff_and_dom_make_a_game_a_coworker_and_his/)
- url: https://youtu.be/z8e_cUsJyxM
---

## [5][If interviewing at a JS-only place, would you say you like Typescript because it makes your code more type-safe?](https://www.reddit.com/r/typescript/comments/j7x355/if_interviewing_at_a_jsonly_place_would_you_say/)
- url: https://www.reddit.com/r/typescript/comments/j7x355/if_interviewing_at_a_jsonly_place_would_you_say/
---
Something like 75% of all errors are TypeErrors and all my portfolio items showcase mid level Typescript even though by experience I'm a junior. It's the one technology I feel strongly about because in the JS days I spent 30% of my time building the project, 70% of the time figuring out where the interpreter's dynamic type inference is throwing an error.

Do JS only companies take offense if you start raving about how much TS has reinforced your code and sense of clarity on what is being returned where? Should I downplay it if I know they're not big on TS? 

Intuition says yes especially given my lack of corporate experience (3 years learning/practicing at home though). PS: I know the best answer is to find a company that has TS but in these times I can't be picky.
## [6][Help getting declaration files to register with typeRoots setting](https://www.reddit.com/r/typescript/comments/j86fer/help_getting_declaration_files_to_register_with/)
- url: https://www.reddit.com/r/typescript/comments/j86fer/help_getting_declaration_files_to_register_with/
---
    // tsconfig.json inside "compilerOptions"
        "typeRoots": ["./@types", "./node_modules/@types"],
    
    // {projectRoot}/@types/express/index.d.ts
    declare global {
      namespace Express {
        interface Request {
          verifiedAccessToken: string
        }
      }
    }

In another file with req annotated to Request inside the function parameter, I am still not able to set verifiedAccessToken without a lint error in VsCode:

    req.verifiedAccessToken = accessToken;
    // Property 'verifiedAccessToken' does not exist on 
    // type 'Request&lt;ParamsDictionary, any, any, ParsedQs&gt;'.ts(2339)

Anything else that needs to be done? Or syntax errors in index.d.ts above?
## [7][They changed the compiler and now the code no longer compiles](https://www.reddit.com/r/typescript/comments/j7tkwz/they_changed_the_compiler_and_now_the_code_no/)
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
## [8][How to have correct "this" type in callback function?](https://www.reddit.com/r/typescript/comments/j7eq51/how_to_have_correct_this_type_in_callback_function/)
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
## [9][How do I use `keyof typeof` on a Map in a class?](https://www.reddit.com/r/typescript/comments/j7h1nr/how_do_i_use_keyof_typeof_on_a_map_in_a_class/)
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
## [10][Best TypeScript Node.js API Postgres Starter Kit?](https://www.reddit.com/r/typescript/comments/j771br/best_typescript_nodejs_api_postgres_starter_kit/)
- url: https://www.reddit.com/r/typescript/comments/j771br/best_typescript_nodejs_api_postgres_starter_kit/
---
I see Microsoft has one for MongoDB. Not sure about it. Has too much for my test -- especially with Express.

I am looking for a starter kit that has TypeScript with best practices ready to go however having express (or not) is not a deal breaker.

Has anyone found anything similiar?
## [11][TS inferring object type as null](https://www.reddit.com/r/typescript/comments/j7b0e5/ts_inferring_object_type_as_null/)
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
    
