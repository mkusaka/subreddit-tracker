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
## [2][From flow to monorepo in Typescript](https://www.reddit.com/r/typescript/comments/eqg5w3/from_flow_to_monorepo_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/eqg5w3/from_flow_to_monorepo_in_typescript/
---
Hello, I have recently moved my multiple repositories javascript packages to a single monorepo in typescript. It's my first-time complete Typescript project and would love it if you had some time to give me some feedback. 
I had a lot of trouble finding some types of external libraries. Does anyone have experience with typing webpack and gulp plugin?

[Repository](https://github.com/FullHuman/purgecss)
## [3][Is it possible to Compile and Update the Browser at the same time?](https://www.reddit.com/r/typescript/comments/eqcv0e/is_it_possible_to_compile_and_update_the_browser/)
- url: https://www.reddit.com/r/typescript/comments/eqcv0e/is_it_possible_to_compile_and_update_the_browser/
---
Hello:  
I am using **lite server** to watch and run my code.  
However, each time I make a code change to a typescript file nothing happens, the only time the browser refreshes is when I make changes to the index.html file.

To see any updates I have to stop **lite server,**  compiled the code and then run the server again.

I attempted watching the code with **tsc --watch** but I have to run the code on a server.  


Are there solutions where you can run a server which compile on any code change and update the browser?
## [4][TSC compiling .ts file dependencies below the folder holding the active package.json](https://www.reddit.com/r/typescript/comments/eq86sy/tsc_compiling_ts_file_dependencies_below_the/)
- url: https://www.reddit.com/r/typescript/comments/eq86sy/tsc_compiling_ts_file_dependencies_below_the/
---
My 1st firebase project is kinda messy as inside `/functions` it installed another package.json file and tsconfig.json. I initially used it to output JS to `/functions/outDir`.

Trouble is now I want to setup database.ts inside `/shared`. This is outside of /functions`,` in terms of heirarchy it's a sibling folder. `/shared/database.ts` will be imported into `/functions/src/index.ts`, and this file does compile to .js.

In this scenario, how will a TS file outside the watch of an active TSC instance be handled? 

If the current setup results in `/shared/database.ts` not being read, is there a way to patch it with tsconfig.json or do I need to move everything (output JS and the package.json being used) down to the project root?
## [5][How to ensure a value implements an interface without casting easily?](https://www.reddit.com/r/typescript/comments/eqcg6e/how_to_ensure_a_value_implements_an_interface/)
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
## [6][Function that reads the first line from a given file](https://www.reddit.com/r/typescript/comments/epzgr8/function_that_reads_the_first_line_from_a_given/)
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
## [7][Recommended TypeScript game engines?](https://www.reddit.com/r/typescript/comments/epxkpk/recommended_typescript_game_engines/)
- url: https://www.reddit.com/r/typescript/comments/epxkpk/recommended_typescript_game_engines/
---
I've been wanting to dabble into video game development so I'd figure it'd be easier for me to learn using something I'm the most familiar with which is TypeScript and JavaScript.

Any recommendations from the sub for one?
## [8][Non-undefined type guards by calling a function](https://www.reddit.com/r/typescript/comments/epzo1g/nonundefined_type_guards_by_calling_a_function/)
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
## [9][Angular + Web Components: a complete guide](https://www.reddit.com/r/typescript/comments/epx6k2/angular_web_components_a_complete_guide/)
- url: https://medium.com/@Armandotrue/angular-web-components-a-complete-guide-5270e5b07e93
---

## [10][Iterating over enums without waste writing duplicate code or](https://www.reddit.com/r/typescript/comments/epyjz6/iterating_over_enums_without_waste_writing/)
- url: https://www.reddit.com/r/typescript/comments/epyjz6/iterating_over_enums_without_waste_writing/
---
Hi,  


I'm new to TS and I'm working on a project with a lot of communication through binary data.  


Oftentimes I get an integer value which is actually a bit field, and I have a list of it's possible values and I have to check which flags are stored in that field. This requires iterating over the possible flags.  


So, in JS I have these objects with the lists of possible flags for a given field and their numeric value.  
In TS I thought about migrating those objects to enums so that I can reference them in interfaces. However, I realized it's impossible to easily iterate over enums as TS automatically compiles those enums (them being numeric enums) to objects with a both a regular mapping of key to value and a reverse mapping of value to key.  
This basically means I have to filter every object I iterate through and discard non-numeric values, which sounds very unidiomatic and inefficient, or I can write duplicate code where I basically write an interface for every possible object and it's implementation, so I use the object to iterate and I reference the interface for stuff like function parameters.  


Is there anything obvious that I'm missing that I could use?  


Thanks.
## [11][Can't get ahold of type definitions VS Code sees](https://www.reddit.com/r/typescript/comments/epp4z3/cant_get_ahold_of_type_definitions_vs_code_sees/)
- url: https://www.reddit.com/r/typescript/comments/epp4z3/cant_get_ahold_of_type_definitions_vs_code_sees/
---
I have a newb question. I am using a third party npm package that has type definitions (if it matters, it's [aws-amplify](https://github.com/aws-amplify/amplify-js)). I know there are type definitions because VS Code complains when I assign a wrong type to a variable returned and expects a certain type. For instance:

    import { Auth } from 'aws-amplify'
    const session = Auth.getCurrentSession()
    const idToken: string = session.getIdToken()  // Type 'CognitoIdToken' is not assignable to type 'string' ts(2322)


However, how do I actually use those type definitions VS Code is seeing? Unfortunately for the above example, `CognitoIdToken` is not exported from 'aws-amplify', so I can't reference it in an `import` statement. Is there some manual way to get ahold the type definitions if they are not exported?

Thank you.

EDIT: formatting
