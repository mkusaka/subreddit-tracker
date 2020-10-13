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
## [2][Make tsconfig reference all exported variables (for testing mock/spy/stub)](https://www.reddit.com/r/typescript/comments/ja4zqn/make_tsconfig_reference_all_exported_variables/)
- url: https://www.reddit.com/r/typescript/comments/ja4zqn/make_tsconfig_reference_all_exported_variables/
---
I have a file like this:

&amp;#x200B;

Foo.tsx:

`export default function Foo() {`

`return 'hi';`

`}`

GroupItem.tsx:

`import Foo from 'Foo.tsx';`

`export const useGroupIds = () =&gt; [];`

`export default function GroupItem() {`

`const groupIds = useGroupIds();`

`return &lt;Foo /&gt;`

`}`

We see in \`GroupItem\` I reference \`useGroupIds\`. This is captured by reference. Is it possible to make TypeScript refer to it as \`module.useGroupIds\`?

I ask this because I am trying to write a test and put a spy on \`useGroupIds\` but when I do that \`spy(module, 'useGroupIds')\` the \`GroupItem\` is not using this spied one as it caught it by reference.

My hope is to get ts to compile the above like this:

`import Foo from 'Foo.tsx';`

`export const useGroupIds = () =&gt; [];`

`export default function GroupItem() {`

`const groupIds = GroupItemModule.useGroupIds();`

`return &lt;`[`FooModule.Foo`](https://FooModule.Foo) `/&gt;`

`}`

&amp;#x200B;

So now when I mock like this:

`import * as FooModule from './Foo';`

`import * as GroupItemModule from './GroupItem';`

`jest.spyOn(FooModule, 'default')`

`jest.spyOn(GroupItemModule, 'useGroupIds')`

Then the spied functions will be properly called.
## [3][Require a type for certain components and not others](https://www.reddit.com/r/typescript/comments/ja7gwc/require_a_type_for_certain_components_and_not/)
- url: https://www.reddit.com/r/typescript/comments/ja7gwc/require_a_type_for_certain_components_and_not/
---
I have a type prop:

`type prop = { width?: number; height?: number; id?: number /* etc... */ [key: string]: any }`

I use `prop` (as shown above) for several components. However, I want to require `id: number` for other components. So these components will have the `prop` from above looking more like this: `type prop = { width?: number; height?: number; id: number; /* etc... */ [key: string]: any }`

Is it possible to require an `id` for certain components, but not others, using the same `type prop = {}`?
## [4][A Complete Guide To TypeScript Decorator](https://www.reddit.com/r/typescript/comments/j9sbaj/a_complete_guide_to_typescript_decorator/)
- url: https://saul-mirone.github.io/a-complete-guide-to-typescript-decorator/
---

## [5][I made a tool that makes writing Minecraft datapack recipes easier](https://www.reddit.com/r/typescript/comments/j9r7uk/i_made_a_tool_that_makes_writing_minecraft/)
- url: https://github.com/Asha20/simple-recipe
---

## [6][I couldn't seem to find help for the following question.](https://www.reddit.com/r/typescript/comments/ja2uz7/i_couldnt_seem_to_find_help_for_the_following/)
- url: https://www.reddit.com/r/typescript/comments/ja2uz7/i_couldnt_seem_to_find_help_for_the_following/
---
I'm sorry to cross post but having working at this for several days and cannot find an answer.

https://preview.redd.it/kw1i46ci9rs51.png?width=1554&amp;format=png&amp;auto=webp&amp;s=3ba46d4d44e5f96d0b963f61b60ad7030e548182
## [7][The performance of index.ts files](https://www.reddit.com/r/typescript/comments/j9si8i/the_performance_of_indexts_files/)
- url: https://www.reddit.com/r/typescript/comments/j9si8i/the_performance_of_indexts_files/
---
I'm working on a CLI that uses index.ts files to slice the code base into unidirectional layers:

    - core   
      /index.ts 
    - util   
      /index.ts 
    - etc 

I'm really happy with the architecture, but recently I noticed a huge performance boost when not using barrel imports from index.ts files. When instead importing from individual files, the CLI ran multitudes faster.

It seems that when a CLI-feature imports from the core-barrel, the import would include all internal and external modules in that barrel, even though it only uses a fraction of them. Take this as an example:

    - core   
      /expensive.ts   
      /quick.ts   
      /index.ts 

If the said CLI-feature uses only quick.ts, it would still import everything associated to expensive.ts through the barrel-import.

Is there any way to get around this? I'd guess that 99.9% of developers using index.ts files do so to improve readability and architecture. The almost accidental import of everything inside a barrel just seems like an unwanted side-effect.

Right now I'm thinking of rewiring all barrel-imports in a prebuild step, but that feels so ugly!
## [8][TS can't infer the returns from a callback function](https://www.reddit.com/r/typescript/comments/j9zknb/ts_cant_infer_the_returns_from_a_callback_function/)
- url: https://www.reddit.com/r/typescript/comments/j9zknb/ts_cant_infer_the_returns_from_a_callback_function/
---
I'm not clear on how to handle async callbacks with TS. In the following method TS can't infer the return paths from the async callback:

        async send(): Promise&lt;object&gt; {
          &lt;config removed&gt;
    
          mailgun.messages().send(messageDetails, (error, message) =&gt; {
            if (error) return { error };
            return { message };
          });
        }
    
    // A function whose declared type is neither 'void' nor 'any' 
    // must return a value.ts(2355)

What is the proper way to assert that the callback will handle the returns here?
## [9][Isomorphic TypeScript client for PostgreSQL](https://www.reddit.com/r/typescript/comments/j9nlal/isomorphic_typescript_client_for_postgresql/)
- url: https://www.reddit.com/r/typescript/comments/j9nlal/isomorphic_typescript_client_for_postgresql/
---
I've developed this open source library which allows safely exposing PostgreSQL APIs with minimal amount of code. Features:

\- CRUD operations

\- Subscriptions

\- Optimistic Client-Server data replication

\- Security Rules

\- Generated TS Types for database schema

WIP: Auth + Android library

&amp;#x200B;

I've recently rewritten it in TypeScript and I am now open to any feedback or comments you might have.

Thanks for reading!

https://prostgles.com/
## [10][JSDoc comments from typescript file to markdown .](https://www.reddit.com/r/typescript/comments/j9u7bh/jsdoc_comments_from_typescript_file_to_markdown/)
- url: https://www.reddit.com/r/typescript/comments/j9u7bh/jsdoc_comments_from_typescript_file_to_markdown/
---
So I have a ts file from which I export a class and an object . The class has JSDoc comments on its methods. There are no JSDoc comments somewhere else in the file . Is there any way to convert the JSDoc comments for the exported class to single markdown file?
## [11][Asserting type of value of keyof some object](https://www.reddit.com/r/typescript/comments/j9sl7u/asserting_type_of_value_of_keyof_some_object/)
- url: https://www.reddit.com/r/typescript/comments/j9sl7u/asserting_type_of_value_of_keyof_some_object/
---
Time and time again I encounter this problem:
Say I have a generic object T.
In typescript I can easily describe a string key of T:
key: keyof T &amp; string.

What I dont understand how to do is type the value of that lookup using key. For example, how do I express that a key leads to a Date value? As far as I can understand, indexers dont solve this problem as they index the whole object, not a specific key.

I keep trying to create functions that recieve an object and a key of said object, and return the value from that lookup - again while asserting a specific return type, without a general indexer.
