# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Where to see basic docs for types](https://www.reddit.com/r/typescript/comments/eqs63d/where_to_see_basic_docs_for_types/)
- url: https://www.reddit.com/r/typescript/comments/eqs63d/where_to_see_basic_docs_for_types/
---
Where are the docs for stuff like @types/express or @types/pg
I know the JavaScript docs for the most part but not the names of the specific types.
Is looking in the specific files in node_modules the only way to know?
## [3][Trouble finding this return type](https://www.reddit.com/r/typescript/comments/eqtiha/trouble_finding_this_return_type/)
- url: https://www.reddit.com/r/typescript/comments/eqtiha/trouble_finding_this_return_type/
---
There are 3 modules I'm working with: firebase, firebase-admin, and firebase-testing. admin handles authentication and has access to some classes and types. The core module does too.

Testing is needed for testing but is making things extra complicated at the moment.

Take this example:

    import * as firebase from "firebase"; // also tried without * aliasing
    import {firestore} from "firebase-admin";
    import firebaseTesting from "@firebase/testing";
    
    // https://firebase.google.com/docs/reference/js/firebase.firestore.Firestore
    // TS2740: Type '(credentials: any) =&gt; Firestore' is missing the following properties from type 'Firestore': settings, collection, doc, collectionGroup, and 4 more.
    const createTestDatabase: firestore.Firestore = credentials =&gt; {
      return firebaseTesting
        .initializeTestApp({
          projectId: 'testProject',
          auth: credentials
        })
        .firestore();
    };
    
    // Another try...
    // https://firebase.google.com/docs/reference/js/firebase.html#initializeapp
    // TS2740: Type '(credentials: any) =&gt; Firestore' is missing the following properties from type 'App': auth, database, delete, installations, and 8 more.
    const createTestDatabase: firebase.app.App = credentials =&gt; {
      return firebaseTesting
        .initializeTestApp({
          projectId: 'testProject',
          auth: credentials
        })
        .firestore();
    };

When faced with tough typings do you guys just do the easy thing and mark it as "any" or is there a process you follow to find the right typing? I tried following the docs to find the right return type but the two above haven't worked.

&amp;#x200B;

PS: What I would really like is if the IDE could help me deduce the correct type. I hear VSCode outright tells you. I'm using Phpstorm. It seems to know the shape but 1) doesn't give the full member list 2) doesn't make any type suggestions even though Phpstorm scans all libraries deeply.
## [4][Become a Better Software Architect](https://www.reddit.com/r/typescript/comments/eqinar/become_a_better_software_architect/)
- url: https://github.com/justinamiller/SoftwareArchitect
---

## [5][From flow to monorepo in Typescript](https://www.reddit.com/r/typescript/comments/eqg5w3/from_flow_to_monorepo_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/eqg5w3/from_flow_to_monorepo_in_typescript/
---
Hello, I have recently moved my multiple repositories javascript packages to a single monorepo in typescript. It's my first-time complete Typescript project and would love it if you had some time to give me some feedback. 
I had a lot of trouble finding some types of external libraries. Does anyone have experience with typing webpack and gulp plugin?

[Repository](https://github.com/FullHuman/purgecss)
## [6][Is it possible to Compile and Update the Browser at the same time?](https://www.reddit.com/r/typescript/comments/eqcv0e/is_it_possible_to_compile_and_update_the_browser/)
- url: https://www.reddit.com/r/typescript/comments/eqcv0e/is_it_possible_to_compile_and_update_the_browser/
---
Hello:  
I am using **lite server** to watch and run my code.  
However, each time I make a code change to a typescript file nothing happens, the only time the browser refreshes is when I make changes to the index.html file.

To see any updates I have to stop **lite server,**  compiled the code and then run the server again.

I attempted watching the code with **tsc --watch** but I have to run the code on a server.  


Are there solutions where you can run a server which compile on any code change and update the browser?
## [7][How to ensure a value implements an interface without casting easily?](https://www.reddit.com/r/typescript/comments/eqcg6e/how_to_ensure_a_value_implements_an_interface/)
- url: https://www.reddit.com/r/typescript/comments/eqcg6e/how_to_ensure_a_value_implements_an_interface/
---
For example, we have interface A:

```typescript
interface A { a: string }
```

and create an object implements A.

```typescript
const a = {
  a: 'foo',
  b: 'bar',
};
```

Can I ensure the a implementing A?

I want to keep the type of a as literal type, `{ a: string; b: string }`, but I want ensure that a is implementing A easily.

Anyone have idea?
## [8][TSC compiling .ts file dependencies below the folder holding the active package.json](https://www.reddit.com/r/typescript/comments/eq86sy/tsc_compiling_ts_file_dependencies_below_the/)
- url: https://www.reddit.com/r/typescript/comments/eq86sy/tsc_compiling_ts_file_dependencies_below_the/
---
My 1st firebase project is kinda messy as inside `/functions` it installed another package.json file and tsconfig.json. I initially used it to output JS to `/functions/outDir`.

Trouble is now I want to setup database.ts inside `/shared`. This is outside of /functions`,` in terms of heirarchy it's a sibling folder. `/shared/database.ts` will be imported into `/functions/src/index.ts`, and this file does compile to .js.

In this scenario, how will a TS file outside the watch of an active TSC instance be handled? 

If the current setup results in `/shared/database.ts` not being read, is there a way to patch it with tsconfig.json or do I need to move everything (output JS and the package.json being used) down to the project root?
## [9][Function that reads the first line from a given file](https://www.reddit.com/r/typescript/comments/epzgr8/function_that_reads_the_first_line_from_a_given/)
- url: https://www.reddit.com/r/typescript/comments/epzgr8/function_that_reads_the_first_line_from_a_given/
---
Hey everyone,

In my TS project I needed a function that reads the first line from a file, so I took the JS example from [this SO post](https://stackoverflow.com/questions/28747719/what-is-the-most-efficient-way-to-read-only-the-first-line-of-a-file-in-node-js) and turned into a TS function, also adding support for handling multiple line endings. Please review my code and offer suggestions.

    import fs from "fs";
    
    export const readFirstLine = (path: string) =&gt; new Promise&lt;string&gt;((resolve, reject) =&gt; {
        const rs = fs.createReadStream(path, {encoding: "utf8"});
    
        let acc = "";
        let pos = 0;
    
        rs
            .on("data", chunk =&gt; {
                acc += chunk;
    
                const indexCR = chunk.indexOf("\r");
                const indexLF = chunk.indexOf("\n");
    
                if (indexCR === -1 &amp;&amp; indexLF === -1) {
                    pos += chunk.length;
                } else {
                    pos += (indexCR !== -1 &amp;&amp; indexLF !== -1)
                        ? Math.min(indexCR, indexLF) // [CRLF]
                        : Math.max(indexCR, indexLF); // [CR] | [LF]
    
                    rs.close()
                }
            })
            .on("close", () =&gt; resolve(acc.slice(0, pos)))
            .on("error", e =&gt; reject(e))
    });

(It's a named arrow function out of personal preference (fewer blocks), but I'm open to suggestions if you prefer the classical `function` notation).

**Update**: minor changes.
## [10][Recommended TypeScript game engines?](https://www.reddit.com/r/typescript/comments/epxkpk/recommended_typescript_game_engines/)
- url: https://www.reddit.com/r/typescript/comments/epxkpk/recommended_typescript_game_engines/
---
I've been wanting to dabble into video game development so I'd figure it'd be easier for me to learn using something I'm the most familiar with which is TypeScript and JavaScript.

Any recommendations from the sub for one?
## [11][Non-undefined type guards by calling a function](https://www.reddit.com/r/typescript/comments/epzo1g/nonundefined_type_guards_by_calling_a_function/)
- url: https://www.reddit.com/r/typescript/comments/epzo1g/nonundefined_type_guards_by_calling_a_function/
---
This non-undefined type guard works:

    function f(x?: number) {
      let y: number
      if (x === undefined) throw 'undefined detected'
      y = x // No compile error here
    }

Is there any way to get a type guard from calling a function?

    function check(x?: number) {
      if (x === undefined) throw 'undefined detected'
    }
    
    function f(x?: number) {
      let y: number
      check(x)
      y = x // Error: Type 'number | undefined' is not assignable to type 'number'.
    }

Is there any way to make this work? At the moment I'm using the non-null assertion operator `!` but I'm trying for a more elegant solution. Thanks.
