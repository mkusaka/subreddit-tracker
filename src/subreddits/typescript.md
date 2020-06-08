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
## [2][How should a variable initialized to empty object be assigned an interface before it is returned?](https://www.reddit.com/r/typescript/comments/gywrtm/how_should_a_variable_initialized_to_empty_object/)
- url: https://www.reddit.com/r/typescript/comments/gywrtm/how_should_a_variable_initialized_to_empty_object/
---
    export default interface ConfigData {
       languageCode : string
       , numberOfRepeats : number
    }

Probably a basic thing for you guys to sort out. I'm thinking on initialization, allow TS to infer variable type, then enforce the interface as return type:

       run() {
          // data to be returned at end
          const configData = {};
    
          this.steps.forEach(async (configStep: WizardSteps): ConfigData =&gt; {
             // capture answer
             const rawUserResponse: string|void = await configStep.explain();
    
             if (configStep.needsValidation) {
                // loops until validated
                const validatedInput: string = configStep.validatedInput(rawUserResponse);
                
                if (configStep.hasSaveableData) { 
                   // convert to k/v pair
                   const formattedForConfiguration = configStep.format(rawUserResponse);
    
                   // save to local object
                   Object.assign(configData, formattedForConfiguration);
                }
             }
          })
    
          return configData;
       }

Is this the proper pattern? If not, how should an interface enforce types on something that is initially set to one value and then mutated?
## [3][Creating Your First React Typescript Project from Scratch](https://www.reddit.com/r/typescript/comments/gyauln/creating_your_first_react_typescript_project_from/)
- url: https://www.youtube.com/watch?v=ODvirqIC09A&amp;list=PLITOO2g_PUHTtMcBFiDk3ITbOh6YR0nhV&amp;index=2&amp;t=0s
---

## [4][How do you type server responses and FE models? (+ examples of large real world apps?)](https://www.reddit.com/r/typescript/comments/gy6pxr/how_do_you_type_server_responses_and_fe_models/)
- url: https://www.reddit.com/r/typescript/comments/gy6pxr/how_do_you_type_server_responses_and_fe_models/
---
Hi everyone!  
I'm starting to implement TS in my company and i'm not sure how to go about typing server responses and front-end models.   


For example, I have a server response for a Block resource:  
`export interface IBlockResponse {`  
 `_id: string;`  
 `page_id: string;`  
 `type: TBlockType;`  
 `styles: IStyles;`  
`}`  


Inside my api service i do a transformation that adds a parent\_id field to each block. So now I also have a:  
`export interface ITransformedBlockResponse`  
 `extends IBlockResponse {`  
 `parent_id: string;`  
`}`  


I then create JS class instances out of these responses, and these classes have some computed fields and might not include all fields from the server response.  
So now I have a third representation of basically the same data. I'm wondering if there's a more canonical way of typing these interfaces between server and client representations of the same data.  
Also wondering if creating a new interface for the transformed data is necessary?  


**P.S** \- If anyone has some examples of some open-source real world apps that have this type of behaviour I'd be very happy to  see them!  


Thanks!  
Uri
## [5][typecheck.macro - A library/macro for automatically generating validation functions from Typescript types now supports intersection types, circular references/recursive types, detailed error messages, type analysis, maps, sets, and more!](https://www.reddit.com/r/typescript/comments/gxsret/typecheckmacro_a_librarymacro_for_automatically/)
- url: https://www.reddit.com/r/typescript/comments/gxsret/typecheckmacro_a_librarymacro_for_automatically/
---
A while back, I made[ this post](https://www.reddit.com/r/typescript/comments/ges5r9/auto_generate_typechecker_from_typescript_types/) asking whether it was automatically possible to generate validation functions for Typescript types. It seems like this is not the case. Libraries like ajv, io-ts, zod, and runtypes exist, but none of them are truly automatic (and there are other issues as well).

About 3 weeks ago, I[ released](https://www.reddit.com/r/typescript/comments/glyq1m/typecheckmacro_automatically_generate_validation/)[ typecheck.macro](https://github.com/vedantroy/typecheck.macro), a compile macro/library for automatically generating validation functions for Typescript types. I've been working on it non-stop since then, and I've added a bunch more features.

The macro supports a fairly large portion of the Typescript type system, so you can automatically generate a validator for most Typescript types. (Look at the README/the bottom of this post for a full list of supported features).

So why use this library/macro? A few reasons:

\- It's seamless. You don't need to define your types in a DSL. Just use... Typescript! The macro automatically parses your type declarations and generates validation functions. This removes a lot of mental overhead and makes the overall process way easier since you don't need to worry about issues like,[ io-ts](https://github.com/gcanti/io-ts/issues/477) or[ runtypes](https://github.com/pelotom/runtypes/issues/147) not supporting rest types. (Note: These libraries are also super cool, all comparisons are friendly!)

\- It's really really really fast. I did some benchmarking in the original post and some more benchmarking for this post. It turns out when you generate "boolean validators" (validators that only return true/false) this library is consistently 2x - 3x faster than ajv, which is the fastest JSON schema validation library. When you enable error messages, this library is still generally faster than ajv, **and** it gives waayyy better error messages.

The reason this macro is fast is because it generates optimized code at compile time whereas other libraries don't generate code. (ajv does generate code, but it 1. only does this at runtime 2. can't handle circular references and 3. this library tends to generate less/faster code).

Here's an example of an error message for the following type with this library:[ ](https://pastebin.com/SiibGU2w)

`| number`

`| {`

`a?: [`

`number | { a: [number, Array&lt;number | string&gt;] },`

`number,`

`...string[]`

`];`

`b: "bar" | false | 42;`

`c: Array&lt;Array&lt;number | boolean | "bar" | "zar"&gt;&gt;;`

`}`

Error message (first entry is path, 2nd is actual value, 3rd is expected value): [https://pastebin.com/SiibGU2w](https://pastebin.com/SiibGU2w)

Here's ajv's equivalent of the same error message: "data should be number, data should be object, data should match some schema in anyOf".

Supported Typescript constructs:

\- interfaces/object literals (no extends, but all the other goodies are there, such as index signatures and optional properties)

\- intersection types

\- union types

\- generics

\- literals

\- arrays/maps/sets

\- tuples

\- type aliases

\- circular types/circular references
## [6][TypeError: Cannot assign to read only property 'isNew' of object '#&lt;Session&gt;](https://www.reddit.com/r/typescript/comments/gxucpp/typeerror_cannot_assign_to_read_only_property/)
- url: https://www.reddit.com/r/typescript/comments/gxucpp/typeerror_cannot_assign_to_read_only_property/
---

	static createSession(admin:AdminDocument) : CookieSessionInterfaces.CookieSessionObject {
			const adminToken = jwt.sign({
				id: admin.id,
				email: admin.email
			}, process.env.JWT_KEY!);

			const sessionObject = {
				isChanged: false,
				isNew: true,
				isPopulated: true,
				jwt: adminToken
			};
			return sessionObject;
	}

Ok, so I tried to create a sessionObject with just jwt, but it wouldn't let me, because it said I needed isNew, isChanged and isPopulated as properties of the session object, but now I get another error and it seems to tell me to not add isNew, so what do you need to do? Typescript is telling me to do two opposite things.

	 TypeError: Cannot assign to read only property 'isNew' of object '#&lt;Session&gt;

How do we avoid the above error?
## [7][Getting error for types that I have installed but dont use when I tsc .](https://www.reddit.com/r/typescript/comments/gxsafu/getting_error_for_types_that_i_have_installed_but/)
- url: https://www.reddit.com/r/typescript/comments/gxsafu/getting_error_for_types_that_i_have_installed_but/
---
Steps to reproduce :

create a folder and in it execute :

    npm init -y;
    npm install --save-dev typescript;
    npm install @types/jest-environment-puppeteer;

create file `./tsconfig.json` that has content :

    {
    	"compilerOptions": {
    		"module": "commonjs",
    		"outDir": "./dist",
    		"rootDir": "./src",
    		"target": "ESNext",
    		"declaration": true
    	},
    	"include": [
    		"./src"
    	],
    	"exclude": [
    		"node_modules"
    	]
    }

create file `./src/index.ts` that has content :

    const a = 3;

Run `npx tsc` and look at the errors you get :

    node_modules/@jest/environment/build/index.d.ts:8:23 - error TS2688: Cannot find type definition file for 'jest'.
    
    8 /// &lt;reference types="jest" /&gt;
                            ~~~~
    
    node_modules/@jest/environment/build/index.d.ts:11:8 - error TS1259: Module '"~/Desktop/test_jest_types_bug_with_tsc/node_modules/jest-mock/build/index"' can only be default-imported using the 'esModuleInterop' flag
    
    11 import jestMock, { ModuleMocker } from 'jest-mock';
              ~~~~~~~~
    
      node_modules/jest-mock/build/index.d.ts:133:1
        133 export = JestMock;
            ~~~~~~~~~~~~~~~~~~
        This module is declared with using 'export =', and can only be used with a default import when using the 'esModuleInterop' flag.
    
    node_modules/@jest/source-map/build/getCallsite.d.ts:8:100 - error TS2503: Cannot find namespace 'callsites'.
    
    8 declare const _default: (level: number, sourceMaps?: Record&lt;string, string&gt; | null | undefined) =&gt; callsites.CallSite;
                                                                                                         ~~~~~~~~~
    
    
    Found 3 errors.

Why am I getting errors for types that I have installed and dont use ?
## [8][[CheckPoint] VSCode extension for easy file state tracking. No more hammering the undo button to recover snippets. More info in the comments.](https://www.reddit.com/r/typescript/comments/gxw9e9/checkpoint_vscode_extension_for_easy_file_state/)
- url: https://www.github.com/BurntBanana/CheckPoint
---

## [9][My validation library written with type inference in mind](https://www.reddit.com/r/typescript/comments/gx6c8l/my_validation_library_written_with_type_inference/)
- url: https://www.reddit.com/r/typescript/comments/gx6c8l/my_validation_library_written_with_type_inference/
---
Hi all!

I just finished work on README for my first user input validation library.

This is still in alpha but I already use it in my production projects and I hope it would be useful for you too.

This is my first step to open source, so feel free to comment, open issues and create pull requests.

Any critics is welcome.

I know there are many user input validating libraries, but I did not found one with type inference. And type inference is a one of the greatest features of TypeScript.

NPM: [https://www.npmjs.com/package/treat-like](https://www.npmjs.com/package/treat-like)

GitHub: [https://github.com/atomAltera/treat-like](https://github.com/atomAltera/treat-like)
## [10][Looking for something more advanced? Learn Data Structures and Algorithms in Typescript](https://www.reddit.com/r/typescript/comments/gx67lc/looking_for_something_more_advanced_learn_data/)
- url: https://www.reddit.com/r/typescript/comments/gx67lc/looking_for_something_more_advanced_learn_data/
---
https://github.com/jeffzh4ng/dsa-ts

MIT/Harvard/Stanford's online DS/algos courses are great for theory and understanding. Real life application and real code matters too though. You learn best when you actually implement these DSA for yourself.

This repository is a collection of DS/algorithms implemented in Typescript. If you also need video lectures, they're there too.

I finished the first section of data structures. sequences.

- static/dynamic arrays
- linked lists
- stacks
- queues
- double-ended queues
- circular buffers

Next week we'll cover the Priority Queue ADT and Heaps.

- binary heaps
- indexed binary heaps
- binomial heaps
- fibonnaci heaps
- pairing heap
- soft heap

The repository and series is just getting started! The plan is to go over all classic data structures. Hashing, self balancing trees, tries, graphs.

And algorithms: sorting, searching, backtracking, dynamic, greedy, graph theory, minimum spanning trees, and more.

Videos and code here: https://github.com/jeffzh4ng/dsa-ts
## [11][Use types or interfaces to constrain generics?](https://www.reddit.com/r/typescript/comments/gxertm/use_types_or_interfaces_to_constrain_generics/)
- url: https://www.reddit.com/r/typescript/comments/gxertm/use_types_or_interfaces_to_constrain_generics/
---
I thought interfaces should be limited to class implementations. However:

    // https://www.typescriptlang.org/docs/handbook/generics.html
    
    interface Lengthwise {
        length: number;
    }
    
    function loggingIdentity&lt;T extends Lengthwise&gt;(arg: T): T {
        console.log(arg.length);  // Now we know it has a .length property, so no more error
        return arg;
    }

I checked this code and it also compiles.

    type Lengthwise = {
        length: number;
    }
    
    function loggingIdentity&lt;T extends Lengthwise&gt;(arg: T): T {
        console.log(arg.length); 
        return arg;
    }

Can you guys tell me which you would prefer? I would think the type definition should almost always be preferred unless the generic is being used to create a class factory.
