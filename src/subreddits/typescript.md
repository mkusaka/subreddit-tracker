# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
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
## [2][Idea: exporting "function bundles"](https://www.reddit.com/r/typescript/comments/fumqcz/idea_exporting_function_bundles/)
- url: https://www.reddit.com/r/typescript/comments/fumqcz/idea_exporting_function_bundles/
---
It would be cool if there was a CLI of some sort that accepted a file and converted each of its exports into a bundled, portable module.

I was wishing something like this existed recently when I was migrating a single API endpoint (written in TypeScript) from an AWS instance (I'll call this Repo #1) to Google Cloud Functions (Repo #2). Even though all the code was written and working in Repo #1, it was a huge chore to move it to Repo #2 because the endpoint code used a few classes and utility functions that were scattered throughout my codebase. All said and done, I ended up copying and pasting a dozen files (mostly containing simple util functions) over to Repo #2 and having to do a lot of manual rewriting of \`import\` statements to account for the changes in project structure. It was very annoying and dumb.

So I was wishing I could have just exported the endpoint from Repo #1 to a bundled, single-file module that I could drop into Repo #2. Ideally it would do tree shaking so it only imports the code from Repo #1 that it actually needs to run. Everything was implemented in strict TypeScript so it was all types and there were no weird global variables or what have you.

Does anyone have a sense of how difficult this would be to do? I don't have a good grasp of the inner workings of the TS engine or parser.
## [3][Variant: More advanced discriminated unions in TypeScript](https://www.reddit.com/r/typescript/comments/fu9wkk/variant_more_advanced_discriminated_unions_in/)
- url: https://github.com/paarthenon/variant
---

## [4][Humpf: Damped Spring motion as a function of time](https://www.reddit.com/r/typescript/comments/fu74v8/humpf_damped_spring_motion_as_a_function_of_time/)
- url: https://github.com/etienne-dldc/humpf
---

## [5][Function with type as parameter](https://www.reddit.com/r/typescript/comments/fuawqh/function_with_type_as_parameter/)
- url: https://www.reddit.com/r/typescript/comments/fuawqh/function_with_type_as_parameter/
---
Hi there,

I'm having an hard time to create a function receiving a type as a parameter.

The main idea is the to make the following code work. What kind of magic do I have to put on the "feed" function?

    class Cat {
      name: string;
      age: number;
    } 
    
    feed(Pick&lt;Cat, 'name') // Error: 'Pick' only refers to a type, but is being used as a value here.ts(2693)
## [6][dtsearch: CLI to find npm packages with types, either bundled or on DefinitelyTyped](https://www.reddit.com/r/typescript/comments/ftvegk/dtsearch_cli_to_find_npm_packages_with_types/)
- url: https://github.com/danvk/dtsearch
---

## [7][How to infer callback function return type?](https://www.reddit.com/r/typescript/comments/fu5kkg/how_to_infer_callback_function_return_type/)
- url: https://www.reddit.com/r/typescript/comments/fu5kkg/how_to_infer_callback_function_return_type/
---
Hi there! I'm fairly new to typescript, so maybe the question is wrong itself but, I need to do get the possible outputs of a function as the keys for an object. That is:

    interface Tagger {
        factory: (x: number) =&gt; T
        tags: { [key of T]: () =&gt; void }
    }
    
    const tagger: Tagger = {
        factory: x =&gt; x % 2 === 0 ? "even" : "odd",
        tags: {
            even: () =&gt; console.log('is even'),
            odd: () =&gt; console.log('is odd'),
        }
    }

The property `tags` should have possible keys `"even" | "odd"`.

Is this even possible? Thanks!
## [8][Building Vue Enterprise Application: Part 3. The Store](https://www.reddit.com/r/typescript/comments/ftzy3l/building_vue_enterprise_application_part_3_the/)
- url: https://medium.com/@gregsolo/building-vue-enterprise-application-part-3-the-store-dbda0e4bb117
---

## [9][How to optimize a serverless typescript eslint webpack setup for performance](https://www.reddit.com/r/typescript/comments/ftobyk/how_to_optimize_a_serverless_typescript_eslint/)
- url: https://medium.com//how-to-optimise-your-serverless-typescript-webpack-eslint-setup-for-performance-86d052284505?source=friends_link&amp;sk=79e3c3629452ac39d2eb8ea8c041c609
---

## [10][Junior Dev Help](https://www.reddit.com/r/typescript/comments/fu0i38/junior_dev_help/)
- url: https://www.reddit.com/r/typescript/comments/fu0i38/junior_dev_help/
---
I have a project right now for a job I just started that is stumping me and making me sweat. I need to do checks on a dataset to ensure that there is data in a sub-array of a sub-array, and if there is no data, it should return. I have the code to check if the initial array is empty, but I can't really figure out how to check for each array as shown in the picture. For example, the device\_name has something in the sub\_bucket, but dhcp\_fingerprint does not. Do I have put in the code for every el.sub\_bucket\[#\].sub\_bucket\[0\] is empty? Or is there a better way to run this check than writing it 6 times over?

`if (!Array.isArray(result) || result.length &lt; 1 ) {`

`return [];`

https://preview.redd.it/9x4mf5suriq41.png?width=1010&amp;format=png&amp;auto=webp&amp;s=c6dc9f6eb127ac03e63417dfe2de959ffd1afd82
## [11][TypeScript (WebStorm) suggest imports outside of rootDir](https://www.reddit.com/r/typescript/comments/ftmqyl/typescript_webstorm_suggest_imports_outside_of/)
- url: https://www.reddit.com/r/typescript/comments/ftmqyl/typescript_webstorm_suggest_imports_outside_of/
---
I'm the only one who gets caught by this bug???? I hate it. 

[imgconv ist the module-root \(where tsconfig is\)](https://preview.redd.it/p4u0q3hwteq41.png?width=909&amp;format=png&amp;auto=webp&amp;s=548c3b0b9a7f5573cbd1ae8a85a94bc190159949)

[https://youtrack.jetbrains.com/issue/WEB-38051](https://youtrack.jetbrains.com/issue/WEB-38051) \- this bug is now open for ONE year!
