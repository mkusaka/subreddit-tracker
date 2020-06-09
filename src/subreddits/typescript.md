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
## [2][Design Doc: Use JavaScript instead of TypeScript for internal Deno Code](https://www.reddit.com/r/typescript/comments/gzkodc/design_doc_use_javascript_instead_of_typescript/)
- url: https://docs.google.com/document/d/1_WvwHl7BXUPmoiSeD8G83JmS8ypsTPqed4Btkqkn_-4/preview?pru=AAABcry4Yhs*kxYogrqCiIU7LhIaH3-5XA
---

## [3][Should this be throwing an error (optional chaining comparison with 0)? (see comments)](https://www.reddit.com/r/typescript/comments/gze2cx/should_this_be_throwing_an_error_optional/)
- url: https://i.redd.it/roruv41tps351.png
---

## [4][Is there an ESLint rule to stop you from coercing numbers to booleans?](https://www.reddit.com/r/typescript/comments/gzeeer/is_there_an_eslint_rule_to_stop_you_from_coercing/)
- url: https://www.reddit.com/r/typescript/comments/gzeeer/is_there_an_eslint_rule_to_stop_you_from_coercing/
---
Suppose I have the following check to see if str &amp; number is defined:

    str = undefined
    num = undefined

    if (str)
      console.log("not undefined")
    if (num)
      console.log("not undefined")

However this will fail if num == 0.

    num = 0

    if (num)
      console.log("not undefined")
    else
      console.log("undefined")

And console.log will print that it's undefined instead, even though it was defined as 0, because 0 is evaluated to false.

Is there an ESLint rule that checks if you attempt to evaluate a number as true / false and warn you? I've made way too many bugs because I keep forgetting that numbers get coerced to 0.
## [5][How to add more declaration for an existing library](https://www.reddit.com/r/typescript/comments/gzjth7/how_to_add_more_declaration_for_an_existing/)
- url: https://www.reddit.com/r/typescript/comments/gzjth7/how_to_add_more_declaration_for_an_existing/
---
hi, i was using typescript recently(using my limited ES knowledges on server side)

and there were a 3rd party library, which used to be a pure ES module, but now they had wrote some \`.d.ts\` files, unfortunately they havnt wrote for all the members inside the module, while only about 1/3 had. this caused my project throw error when compiling(i guess since the new version had types file, compiler might treate it as a pure TS module)

so my question is how to add more declaration for this 3rd party library? without sending PR to the upstream?  

if that's impossible, is there anyway to indicate the compiler(tsc) to fallback to treat the library as a ES module?

also i noticed that in the existing declarations, there were some types which can not be accept by my compiler, should i explicit import them from that library?

sorry for the poor english and knowledges
## [6][Why is here no type error when calling super?](https://www.reddit.com/r/typescript/comments/gzgatm/why_is_here_no_type_error_when_calling_super/)
- url: https://www.reddit.com/r/typescript/comments/gzgatm/why_is_here_no_type_error_when_calling_super/
---
```
class Person {
    name!: string
    constructor(input: Person) {
       Object.assign(this, input)
    }
}

class Employee extends Person {
    age!: number
    constructor(input: Employee) {
       super(input)
    }
}
```
I thought that `super(input)` calls `class Person`'s constructor and that requires `input` to be of type `Person` but the `input` passed to super has one more field than `type Person` (`age`).
## [7][How should a variable initialized to empty object be assigned an interface before it is returned?](https://www.reddit.com/r/typescript/comments/gywrtm/how_should_a_variable_initialized_to_empty_object/)
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
## [8][Creating Your First React Typescript Project from Scratch](https://www.reddit.com/r/typescript/comments/gyauln/creating_your_first_react_typescript_project_from/)
- url: https://www.youtube.com/watch?v=ODvirqIC09A&amp;list=PLITOO2g_PUHTtMcBFiDk3ITbOh6YR0nhV&amp;index=2&amp;t=0s
---

## [9][How do you type server responses and FE models? (+ examples of large real world apps?)](https://www.reddit.com/r/typescript/comments/gy6pxr/how_do_you_type_server_responses_and_fe_models/)
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
## [10][typecheck.macro - A library/macro for automatically generating validation functions from Typescript types now supports intersection types, circular references/recursive types, detailed error messages, type analysis, maps, sets, and more!](https://www.reddit.com/r/typescript/comments/gxsret/typecheckmacro_a_librarymacro_for_automatically/)
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
## [11][TypeError: Cannot assign to read only property 'isNew' of object '#&lt;Session&gt;](https://www.reddit.com/r/typescript/comments/gxucpp/typeerror_cannot_assign_to_read_only_property/)
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
