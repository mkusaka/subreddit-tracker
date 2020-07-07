# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][Best way to setup reloading ?](https://www.reddit.com/r/typescript/comments/hmpxnn/best_way_to_setup_reloading/)
- url: https://www.reddit.com/r/typescript/comments/hmpxnn/best_way_to_setup_reloading/
---
What is your preferred choice:

1. tsc + node
2. nodemon + ts-node
3. ts-node-dev
4. ?
## [3][How can i achieve object copy ??](https://www.reddit.com/r/typescript/comments/hmp6zr/how_can_i_achieve_object_copy/)
- url: https://www.reddit.com/r/typescript/comments/hmp6zr/how_can_i_achieve_object_copy/
---
const target = { a: 1, b: { c: 2} };

const source = { a: 4, b: { d: 3} };

const res = Object.assign(target, source);

console.log(res)

//Output :   { a: 4, b: { d: 3 } };

//Expected : { a: 4, b: { c: 2, d: 3 } };

Any help ???
## [4][Priority Queue in TypeScript](https://www.reddit.com/r/typescript/comments/hmsedt/priority_queue_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hmsedt/priority_queue_in_typescript/
---
 Queues and stacks are used in programming to process a collection of items in a certain order. Those data structures work either with a *first-in-first-out* or *last-in-first-out* system in which the order of processing is always directly coupled to the order of insertion. However, sometimes you need to have more control over the processing order then offered by the insertion order. 

[https://medium.com/@wim\_jongeneel/priority-queue-in-typescript-6ef23116901?source=friends\_link&amp;sk=a35b45b2a26f4cd5daafb84feafb6c88](https://medium.com/@wim_jongeneel/priority-queue-in-typescript-6ef23116901?source=friends_link&amp;sk=a35b45b2a26f4cd5daafb84feafb6c88)
## [5][open-source and free resume builder](https://www.reddit.com/r/typescript/comments/hm9e1s/opensource_and_free_resume_builder/)
- url: https://www.reddit.com/r/typescript/comments/hm9e1s/opensource_and_free_resume_builder/
---
I am Javascript developer, I created this [open-source resume builder](https://github.com/sramezani/resume-builder) with React js (Next js)

The question I faced from many companies (I looking for a job): did you have any experience with Typescript, then I started to refactor my codes from Javascript to Typescript.

I will be happy if You look at the codes and tell me which part needs to improve.

Demo: [**WTFresume.com**](https://wtfresume.com/)

Github repo: [**resume-builder**](https://github.com/sramezani/resume-builder)

Thanks
## [6][Can you require constrained generic parameters?](https://www.reddit.com/r/typescript/comments/hmi15i/can_you_require_constrained_generic_parameters/)
- url: https://www.reddit.com/r/typescript/comments/hmi15i/can_you_require_constrained_generic_parameters/
---
You can essentially require a generic to be passed to a function using the following:

    function contractType&lt;T = void, U extends T = T&gt;(value: unknown): U { return value as U }
    
    const example1: string = contractType(17) // error
    const example2: string = contractType("value") // error
    const example3: string = contractType&lt;string&gt;("value") // ok

However I have a situation where I want to require a generic in this way, but with the generic being constrained to another type:

    function contractType&lt;T extends Type = void, U extends T = T&gt;(value: unknown): U { return value as U }

This of course fails as void is not assignable to type 'Type'. Is there a way I can achieve a similar thing?
## [7][Best place to start learning TS](https://www.reddit.com/r/typescript/comments/hmcp0m/best_place_to_start_learning_ts/)
- url: https://www.reddit.com/r/typescript/comments/hmcp0m/best_place_to_start_learning_ts/
---
Hi,
Please suggest me good resources to start learning TypeScript apart from official docs. I have knowledge of JS and react already but I have not have experience working in company.

Thank you for resource suggestion in advanced.
## [8][typescript is ignoring lib specified in tsconfig.json](https://www.reddit.com/r/typescript/comments/hmae3z/typescript_is_ignoring_lib_specified_in/)
- url: https://www.reddit.com/r/typescript/comments/hmae3z/typescript_is_ignoring_lib_specified_in/
---
typescript is ignoring the libs specified in the tsconfig array

the issue I am running into is Object.values is working in my typescript files, and ctrl+click takes me to lib.es2017.object.d.ts where I would expect this to error

this is the tsconfig for a larger react project, I copy-pasted this tsconfig to a new project and it worked as expected, so I suspect it is not anything in the tsconfig

my tsconfig.json looks like this

    {  
      "compilerOptions": {  
        "target": "es5",  
        "allowJs": true,  
        "checkJs": true,  
        "skipLibCheck": true,  
        "esModuleInterop": true,  
        "allowSyntheticDefaultImports": true,  
        "strict": true,  
        "forceConsistentCasingInFileNames": true,  
        "noUnusedLocals": true,  
        "noUnusedParameters": true,  
        "module": "commonjs",  
        "moduleResolution": "node",  
        "resolveJsonModule": true,  
        "isolatedModules": true,  
        "noEmit": true,  
        "jsx": "preserve",  
        "lib": \["dom", "es5"\],  
        "baseUrl": "./",  
        "paths": {  
            "src/\*": \["src/\*"\]  
        },  
        "typeRoots": \["./src/@types", "./node\_modules/@types"\]  
      },  
      "include": \["src", "test/unit-tests"\]  
    }
## [9][Can some explain? Weird behaviour with enums and union types as type arguments](https://www.reddit.com/r/typescript/comments/hmfsvx/can_some_explain_weird_behaviour_with_enums_and/)
- url: https://www.reddit.com/r/typescript/comments/hmfsvx/can_some_explain_weird_behaviour_with_enums_and/
---
So I came across this behaviour and I can't really grasp why the compiler processes things this way. I will paste an example snippet and the resulting error messages:  


    enum Enum {
      Uno,
      Dos,
    }
    type Union = 'Uno' | 'Dos';
    function testEnum11(...arg: [Enum]) {
      return { arg: arg[0] };
    }
    function testEnum12(arg: Enum) {
      return testEnum11(arg); // good
    }
    function testUnion11(...arg: [Union]) {
      return { arg: arg[0] };
    }
    function testUnion12(arg: Union) {
      return testUnion11(arg); // good
    }
    function testEnum21&lt;T&gt;(...args: T extends undefined ? [] : [T]) {
      return { arg: args[0] };
    }
    function testEnum22(arg: Enum) {
      return testEnum21(arg);  // bad
    }
    function testUnion21&lt;T&gt;(...args: T extends undefined ? [] : [T]) {
      return { arg: args[0] };
    }
    function testUnion22(arg: Union) {
      return testUnion21(arg);  // bad
    }
    function testEnum31&lt;T&gt;(...args: [T]) {
      return { arg: args[0] };
    }
    function testEnum32(arg: Enum) {
      return testEnum31(arg); // good
    }
    function testUnion31&lt;T&gt;(...args: [T]) {
      return { arg: args[0] };
    }
    function testUnion32(arg: Union) {
      return testUnion31(arg); // good
    }

Where I've added the comment 'bad', the arg to the callee function has a squiggly line with the following errors:

* Enum case: Argument of type '\[Enum\]' is not assignable to parameter of type '\[Enum.Uno\] | \[Enum.Dos\]'. Type '\[Enum\]' is not assignable to type '\[Enum.Uno\]'. Type 'Enum' is not assignable to type '[Enum.Uno](https://Enum.Uno)'.ts(2345)
* Union type case: Argument of type '\[Union\]' is not assignable to parameter of type '\["Uno"\] | \["Dos"\]'. Type '\[Union\]' is not assignable to type '\["Uno"\]'. Type 'Union' is not assignable to type '"Uno"'. Type '"Dos"' is not assignable to type '"Uno"'.ts(2345)

Can you explain why in tests 22 it kind of spreads the Enum/union types?

Thanks
## [10][Recommendations for cli interfaces](https://www.reddit.com/r/typescript/comments/hmaccb/recommendations_for_cli_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/hmaccb/recommendations_for_cli_interfaces/
---
I am looking to build an interactive CLI (an interactive, shell like application). IOW something that is like readline with history and autocompletion (e.g. async, REST-roundtrip required type autocompletion), but offers an API similar to commander.js so that I can use a fluent interface to build the CLI syntax and add actions and help easily. Any recommendations?
## [11][Complimentary technologies for someone wanting to specialise in TypeScript](https://www.reddit.com/r/typescript/comments/hmeyc9/complimentary_technologies_for_someone_wanting_to/)
- url: https://www.reddit.com/r/typescript/comments/hmeyc9/complimentary_technologies_for_someone_wanting_to/
---
I'm currently a full stack dev in the .Net space hoping to move into a more TypeScript central role (I use it within Angular at the moment but it only accounts for ~5-10% of my work). What languages and/or frameworks would you all recommend learning to compliment TypeScript and provide the best base for getting a job?

From looking at some job specs (which is the most obvious place to start researching this), Node and React seem to be the most common. From my last few development jobs though, I've found that the languages listed in a job spec haven't matched what I've actually ended up spending most my time on.

I've done some searching, but I'd like to get some insights from real TypeScript developers!
