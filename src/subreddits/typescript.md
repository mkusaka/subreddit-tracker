# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][Is strict mode worth it?](https://www.reddit.com/r/typescript/comments/guu3a8/is_strict_mode_worth_it/)
- url: https://www.reddit.com/r/typescript/comments/guu3a8/is_strict_mode_worth_it/
---
I just got caught out so I've turned on strictNullChecks, but apart from that is strict mode worth it?
## [3][Can I improve my knowledge and more deeply understand OOP with TypeScript?](https://www.reddit.com/r/typescript/comments/gux7fw/can_i_improve_my_knowledge_and_more_deeply/)
- url: https://www.reddit.com/r/typescript/comments/gux7fw/can_i_improve_my_knowledge_and_more_deeply/
---
A couple years a go I switched careers and became a front end dev. I love it but I have a lot to learn obviously. One of the things I really want to get a better grasp of is OOP. I have a pretty solid JavaScript knowledge but I have read a lot that it isn't truly an OOP language. I would like to learn TS to start using with React and Redux anyway so I was wondering if TS would help me to learn OOP properly, and if so, can you suggest any resources.   


Thank you very much in advance.
## [4][What's the purpose of having a type variable extend any?](https://www.reddit.com/r/typescript/comments/gupooo/whats_the_purpose_of_having_a_type_variable/)
- url: https://www.reddit.com/r/typescript/comments/gupooo/whats_the_purpose_of_having_a_type_variable/
---
I've randomly run into this looking at random TS snippets online or library source codes. Here's an example of it inside [`ts-toolbelt`](https://github.com/pirix-gh/ts-toolbelt/blob/0d5f604/src/Any/Compute.ts#L13).

I thought the following two examples were equivalent; are they not?

    type WithAny&lt;T extends any&gt; = T;
    type WithoutAny&lt;T&gt; = T;
## [5][Generating TypeScript definitions for open source project with over 300k LOC | cesium.com](https://www.reddit.com/r/typescript/comments/gutl6m/generating_typescript_definitions_for_open_source/)
- url: https://cesium.com/blog/2020/06/01/cesiumjs-tsd/
---

## [6][Do you guys use eslint on your typescript projects?](https://www.reddit.com/r/typescript/comments/guiovl/do_you_guys_use_eslint_on_your_typescript_projects/)
- url: https://www.reddit.com/r/typescript/comments/guiovl/do_you_guys_use_eslint_on_your_typescript_projects/
---
I used to use eslint with the plugins it suggested to me to install after executing :

    npx eslint --init;

and chooseing the option : I use typescript  , but I have found it to produce some akward linting errors recently like not allowing me to put type `any` on a function parameter. 

I have stopped using eslint since then. 

If you use eslint then what you use it for?

I was using it just because it was linting errors while I was writing js and so I believed it will also be useful in typescript.
## [7][Create type for sync function wrapper](https://www.reddit.com/r/typescript/comments/guue8c/create_type_for_sync_function_wrapper/)
- url: https://www.reddit.com/r/typescript/comments/guue8c/create_type_for_sync_function_wrapper/
---
I would like to create a wrapper for any \_synchronous\_ function (the function shouldn't return a promise), is there a way to do that?

I did the following, but it still allows promise returning functions:

    export function wrapSyncFn&lt;T extends any[]&gt;(
        handler: (...args: T) =&gt; any,
    ): typeof handler {
      return (...args) =&gt; {
        return handler(...args)
      }
    }
    
    async function promiseFn() {
      return 0
    }
    
    wrapSyncFn(promiseFn)
    
## [8][Pass on specific props, based on interface [REACT]](https://www.reddit.com/r/typescript/comments/guruw2/pass_on_specific_props_based_on_interface_react/)
- url: https://www.reddit.com/r/typescript/comments/guruw2/pass_on_specific_props_based_on_interface_react/
---
So, let's say I have a TopComponent.tsx which gets Props1 of interface Styles, Helpers and some other stuff. TopComponent ist a FunctionComponent. 

Inside this Component I use a different Component, lets call it HolderComponent, which is expecting Props2 of interface Styles and some other different stuff. And they share some props (Styles)

So, I am aware that I can pass on props to other components by using {...props}, or by restructuring. So I get my props from Props1 and pass it on to the HolderComponent - but I don't new all of these props, I only need a subset of props - that of interface Styles. I would expect something like (pseudocode) {...props | Styles}, which pass all props from the TopComponent which are part of the interface Styles, vomiting the rest.

Is that possible, am I missing something? 

The reason for this is: performance hit (if there is any), and code readability.
## [9][Having trouble using imports and exports](https://www.reddit.com/r/typescript/comments/gusict/having_trouble_using_imports_and_exports/)
- url: https://www.reddit.com/r/typescript/comments/gusict/having_trouble_using_imports_and_exports/
---
Hello,

I am currently working on my first TS project and am having trouble with exports and imports.

I originally wrote my code in one file, while ran great; but splitting the project up into multiple files is finding to be challenging.

&amp;#x200B;

Upon splitting up into multiple files I get *"Uncaught ReferenceError: exports is not defined"* in the browser console.

&amp;#x200B;

What I have tried:

* npm install commonjs then updating my tsconfig.json to {  
 "compilerOptions": {  
 "target": "es5",  
 "module": "commonjs",  
 "strict": true,  
 "outDir": "dist",  
 "sourceMap": true  
  }  
}

I did that based off of  [https://stackoverflow.com/questions/43042889/typescript-referenceerror-exports-is-not-defined](https://stackoverflow.com/questions/43042889/typescript-referenceerror-exports-is-not-defined) 

Unfortunately this did not help.

&amp;#x200B;

Any help or advice would be greatly appreciated.

If looking at the code (3 files less than 90 lines) would be helpful please let me know and Ill paste it. I just did not want to flood this thread.
## [10][[help] Compiler doesn't seem to detect enum variants correctly](https://www.reddit.com/r/typescript/comments/gunlh8/help_compiler_doesnt_seem_to_detect_enum_variants/)
- url: https://www.reddit.com/r/typescript/comments/gunlh8/help_compiler_doesnt_seem_to_detect_enum_variants/
---
at this [playground link](https://www.typescriptlang.org/play/#code/JYOwLgpgTgZghgYwgAgKJSgHlQPmQbwChkTkw4BzALmQHJopaAaY0hgeyhtUIF9DCoSLEQoA8gGtMAFTxFSZSjVrsJzViQAmccjWl8BMAK4gEYYOxDJJMnAAoIegJQ0bsghuRQIYI1Cv4itR0qszI2rrIKLwA3AaExqbmlmgY2PaOaC6pWLgeCt6+-gRBygxhHFxRyLHxYACeAA4oAEoQAM5GADZgMkxoeAC81lLuAD456XGEECBGALbIAIJQwGAAFvM+wAjoUJz5pCYQAG6zyMO0x2cg6goglgCS4BAU0Bd0DyDPkG+M8YkzBYrD9XtAACLAE7ATQQOxwGhzeYAI2g-WRiIWqKg2TanR6mCR2P6KzWm22uwwnDknmAMGQdgAhHYAHJY6AAOmA7VBf3hTmQADJBcg2SjOdzedA7MinHLDgoSIU-FY9nZSRstuZKfsoByvlKcXEFPwFAhLO0wF4Oh84MgAPTIZHG0h0hmMsXYrk8l587zteXyRVKnwqnLq1aail7Tgc66zJwukim0jK4qSOz+xPxc0gdrsLoQDlddgUOyGyHQ2F2AAs-QArHK4rn84Xi6Xy76IVCYXCAJz9ADMTcILYLRZLZYrPerNY59f6ACYR6OLVarMNp1W4fWAGz9ADs2faAHc1gh1gyQBzyBQBUGSAg4O0UCo1FRPApC1aAI6Y8VQB814RHASbIE+L50OUH7BiQ37IOwNAauS2oxoBwzXpUcS8EAA) on line 46, it thinks `o` is of type `Result` instead of just the `Err`variant. Not sure why this is the case.
## [11][Can enums be treated as numbers for incrementation purposes?](https://www.reddit.com/r/typescript/comments/gugq66/can_enums_be_treated_as_numbers_for/)
- url: https://www.reddit.com/r/typescript/comments/gugq66/can_enums_be_treated_as_numbers_for/
---
By default I believe enums are assigned numbers starting from 0. I was thinking of using an enum to define a series of steps for a CLI script. that way I could name each step, for example, "intro", "gatherName", "configureGoogle" etc. There are no branching paths so in this setup, reaching the final enum would signal that the next step should perform an exit procedure.

Does it ever make sense to increment enums? Are there side effects to consider?

If this should not be done, I suppose additional logic will be needed to define step pathing. I'm open to any ideas you guys have. The first idea I have is to iterate over the enumb definition, push each enum into an array to create a sequence, and then allow the indexes to become the sequence values.
