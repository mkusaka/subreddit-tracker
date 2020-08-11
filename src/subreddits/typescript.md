# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Type error - Missing 200+ properties](https://www.reddit.com/r/typescript/comments/i7evqe/type_error_missing_200_properties/)
- url: https://www.reddit.com/r/typescript/comments/i7evqe/type_error_missing_200_properties/
---
    Type error: Type '{ children: Element; className: string; closeButton: true; }' is missing the following properties from type 'Pick&lt;ModalHeaderProps, "style" | "title" | "slot" | "children" | "bsPrefix" | "className" | "color" | "id" | "lang" | "role" | "tabIndex" | "dangerouslySetInnerHTML" |

     ... 246 more ... | "closeButton"&gt;': translate, onAuxClick, onAuxClickCapture  TS2739

https://imgur.com/a/XpLvrpo

Do I need to actually pass in 246+ properties to this Modal.Header component (using bootstrap)? I am pretty bad so please treat me like a huge noob! What should I do here?
## [3][Add Typescript to your Next.js project](https://www.reddit.com/r/typescript/comments/i7m2kv/add_typescript_to_your_nextjs_project/)
- url: https://blog.bhanuteja.dev/add-typescript-to-your-nextjs-project-ckdpgej1h0174kbs1h9fg86yq
---

## [4][The best Rollup config for TypeScript libraries](https://www.reddit.com/r/typescript/comments/i7635n/the_best_rollup_config_for_typescript_libraries/)
- url: https://www.reddit.com/r/typescript/comments/i7635n/the_best_rollup_config_for_typescript_libraries/
---
https://gist.github.com/aleclarson/9900ed2a9a3119d865286b218e14d226

### Features
🔥 Blazing fast builds  
😇 CommonJS bundle  
🌲 .mjs bundle  
✨ .d.ts bundle + type-checking  
🧐 Source maps  

**Please retweet this if it helps you:**  
https://twitter.com/alecdotbiz/status/1291742610374131712
## [5][Why didn't my typings match this overload](https://www.reddit.com/r/typescript/comments/i7e6pl/why_didnt_my_typings_match_this_overload/)
- url: https://www.reddit.com/r/typescript/comments/i7e6pl/why_didnt_my_typings_match_this_overload/
---
This is purely practice with working through the complex method overloads I see. In the case of passport.serializeUser

    (method) passport.Authenticator&lt;e.Handler, any, any,
     passport.AuthenticateOptions&gt;.serializeUser&lt;unknown, unknown&gt;(
    fn: (user: unknown, done: (err: any, id?: unknown) =&gt; void) =&gt; void)
    : void (+1 overload)

I thought this would work:

    passport.serializeUser(&lt;UserShape extends { id: string }&gt;
    (user: UserShape, callback: any) =&gt; {
      callback(null, user.id as string);
    });

The error lint:

    No overload matches this call.
      Overload 1 of 2, '(fn: (user: unknown, done: (err: any, 
    id?: unknown) =&gt; void) =&gt; void): void', gave the following error.
        Argument of type '&lt;UserShape extends { id: string; }&gt;
    (user: UserShape, callback: any) =&gt; void' is not assignable to 
    parameter of type '(user: unknown, done: (err: any, 
    id?: unknown) =&gt; void) =&gt; void'.
          Types of parameters 'user' and 'user' are incompatible.
    
            Type 'unknown' is not assignable to type '{ id: string; }'.
      Overload 2 of 2, '(fn: (req: Request&lt;ParamsDictionary, any, any, 
    ParsedQs&gt;, user: unknown, done: (err: any, id?: unknown) =&gt; 
    void) =&gt; void): void', gave the following error.
    
    Argument of type '&lt;UserShape extends { id: string; }&gt;(user: UserShape, 
    callback: any) =&gt; void' is not assignable to parameter of type '(req: Request&lt;ParamsDictionary, any, any, ParsedQs&gt;, user: unknown, 
    done: (err: any, id?: unknown) =&gt; void) =&gt; void'.
          Types of parameters 'user' and 'req' are incompatible.
            Property 'id' is missing in type 'Request&lt;ParamsDictionary, 
    any, any, ParsedQs&gt;' but required in type '{ id: string; }'.ts(2769)

Seems like the first overload gets close. What was my mistake?
## [6][Healthcare Software Development Team](https://www.reddit.com/r/typescript/comments/i7qfwj/healthcare_software_development_team/)
- url: http://software-development.icu
---

## [7][Type definition for function that could return undefined](https://www.reddit.com/r/typescript/comments/i7d62e/type_definition_for_function_that_could_return/)
- url: https://www.reddit.com/r/typescript/comments/i7d62e/type_definition_for_function_that_could_return/
---
I need help writing a type definition for a function that could return `undefined` or anything else, but the return type has to include `undefined`. I'm wanting to avoid a generic parameter in the type as well.

```
// this doesn't work, but is kind of what I'm going for
type CouldReturnUndefined = &lt;T&gt;(...args: any[]) =&gt; T | undefined

// this should be allowed
const a: CouldReturnUndefined = () =&gt; { return undefined }

// this should cause an error
const b: CouldReturnUndefined = () =&gt; { return 10 }

// this should be allowed
const c: CouldReturnUndefined = (): number | undefined =&gt; { return 10 }
```

Any ideas?
## [8][I built YJs bindings for slate - need feedback on my code](https://www.reddit.com/r/typescript/comments/i77bqu/i_built_yjs_bindings_for_slate_need_feedback_on/)
- url: https://github.com/BitPhinix/slate-yjs
---

## [9][Custom types not recognized](https://www.reddit.com/r/typescript/comments/i76lwt/custom_types_not_recognized/)
- url: https://www.reddit.com/r/typescript/comments/i76lwt/custom_types_not_recognized/
---
Hello, I have an issue using custom definition files inside a project. My folder looks like : 

    tsconfig.json
        |- src/    
            |- common/       
                |- types/          
                    |- index.d.ts          
                    |- arandomfile.d.ts          
                    |- anotherrandomfile.d.ts

my declarations files looks like : 

    interface myFirstCustomType {    
        customStringProperty: string,
        customArrayProperty: Array&lt;string&gt;
        // and so on, whatever
    }
    
    interface mySecondCustomType {    
        customStringProperty: string,
        customArrayProperty: Array&lt;string&gt;
        // and so on, whatever
    }

and In my tsconfig.json, I have specified the path of my types folder with 2 different ways : 

    "typeRoots": ["src/common/types/*"]

or  

    "paths": {   
        "*" : ["src/common/types/*"]
    }

but both doesn't seems to work since my types aren't recognized, any solutions? thanks
## [10][TypeScript/Express/React/Webpack starter project?](https://www.reddit.com/r/typescript/comments/i74q3u/typescriptexpressreactwebpack_starter_project/)
- url: https://www.reddit.com/r/typescript/comments/i74q3u/typescriptexpressreactwebpack_starter_project/
---
I'm trying to make a TypeScript/Express/React/Webpack project and there's a few bits that I can't get plumbed right. Specifically, I want to have a "good" dev experience without a ton of recompiling. 

As such, I want to run the project when developing using both:

1. Nodemon+TS-Node
2. Webpack Hot Module Replacement

I can't seem to find a combination of configurations that will allow me to get both of those working at the same time. Does anyone have any examples of this setup working well?

Thanks!
## [11][Game of Life in TypeScript - text/video tutorial](https://www.reddit.com/r/typescript/comments/i6mtcf/game_of_life_in_typescript_textvideo_tutorial/)
- url: https://medium.com/hypersphere-codes/conways-game-of-life-in-typescript-a955aec3bd49
---

