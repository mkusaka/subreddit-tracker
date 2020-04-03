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
## [2][dtsearch: CLI to find npm packages with types, either bundled or on DefinitelyTyped](https://www.reddit.com/r/typescript/comments/ftvegk/dtsearch_cli_to_find_npm_packages_with_types/)
- url: https://github.com/danvk/dtsearch
---

## [3][How to infer callback function return type?](https://www.reddit.com/r/typescript/comments/fu5kkg/how_to_infer_callback_function_return_type/)
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
## [4][Humpf: Damped Spring motion as a function of time](https://www.reddit.com/r/typescript/comments/fu74v8/humpf_damped_spring_motion_as_a_function_of_time/)
- url: https://github.com/etienne-dldc/humpf
---

## [5][Building Vue Enterprise Application: Part 3. The Store](https://www.reddit.com/r/typescript/comments/ftzy3l/building_vue_enterprise_application_part_3_the/)
- url: https://medium.com/@gregsolo/building-vue-enterprise-application-part-3-the-store-dbda0e4bb117
---

## [6][How to optimize a serverless typescript eslint webpack setup for performance](https://www.reddit.com/r/typescript/comments/ftobyk/how_to_optimize_a_serverless_typescript_eslint/)
- url: https://medium.com//how-to-optimise-your-serverless-typescript-webpack-eslint-setup-for-performance-86d052284505?source=friends_link&amp;sk=79e3c3629452ac39d2eb8ea8c041c609
---

## [7][Junior Dev Help](https://www.reddit.com/r/typescript/comments/fu0i38/junior_dev_help/)
- url: https://www.reddit.com/r/typescript/comments/fu0i38/junior_dev_help/
---
I have a project right now for a job I just started that is stumping me and making me sweat. I need to do checks on a dataset to ensure that there is data in a sub-array of a sub-array, and if there is no data, it should return. I have the code to check if the initial array is empty, but I can't really figure out how to check for each array as shown in the picture. For example, the device\_name has something in the sub\_bucket, but dhcp\_fingerprint does not. Do I have put in the code for every el.sub\_bucket\[#\].sub\_bucket\[0\] is empty? Or is there a better way to run this check than writing it 6 times over?

`if (!Array.isArray(result) || result.length &lt; 1 ) {`

`return [];`

https://preview.redd.it/9x4mf5suriq41.png?width=1010&amp;format=png&amp;auto=webp&amp;s=c6dc9f6eb127ac03e63417dfe2de959ffd1afd82
## [8][TypeScript (WebStorm) suggest imports outside of rootDir](https://www.reddit.com/r/typescript/comments/ftmqyl/typescript_webstorm_suggest_imports_outside_of/)
- url: https://www.reddit.com/r/typescript/comments/ftmqyl/typescript_webstorm_suggest_imports_outside_of/
---
I'm the only one who gets caught by this bug???? I hate it. 

[imgconv ist the module-root \(where tsconfig is\)](https://preview.redd.it/p4u0q3hwteq41.png?width=909&amp;format=png&amp;auto=webp&amp;s=548c3b0b9a7f5573cbd1ae8a85a94bc190159949)

[https://youtrack.jetbrains.com/issue/WEB-38051](https://youtrack.jetbrains.com/issue/WEB-38051) \- this bug is now open for ONE year!
## [9][Assign nested Object as copy, not reference](https://www.reddit.com/r/typescript/comments/ftl7pi/assign_nested_object_as_copy_not_reference/)
- url: https://www.reddit.com/r/typescript/comments/ftl7pi/assign_nested_object_as_copy_not_reference/
---
Hi guys, i seem to not understand the assignment operator completly, when it comes to nested Objects.

I have initialization objects for my classes data, which is defined by an interface.

    export const MY_DATA_INIT {
        device: {
            id: 0,
            ip: "unknown"
        }
        anotherProperty: "unknown"
    }
    
    export interface MyData {
        device: Device;
        anotherProperty: string;
    }
    
    interface Device {
        id: number;
        ip: string;
    }

When i make an Object of type MyData and use MY\_DATA\_INIT to initialize it, i am only pointing to  MY\_DATA\_INIT

    let myDataObject: MyData = MY_DATA_INIT;
    
    myDataObject.device.ip = "192.168.0.2";    // This updates on MY_DATA_INIT
    myDataObject.anotherProperty = "value";    // This does'nt
    
    console.log(JSON.stringify(MY_DATA_INIT)).
    /*    {
     *        device: {
     *            id: 0,
     *            ip: "192.168.0.2",
     *        }
     *        anotherProperty: "unknown"
     *     }
     */

Using the spread operator {...MY\_DATA\_INIT} does not make a difference.

Same using the spread operator in the MY\_DATA\_INIT itself.

    export const MY_DATA_INIT {
        device: {...MY_DEVICE_INIT}
        anotherProperty: "unknown"
    }
    
    export const MY_DEVICE_INIT {
        id: 0,
        ip: "unknown"
    }

I don't get why this is failing. The direct Object properties of MyData seems to work as expected, but the nested device object is passed as reference.

Any ideas?

Edit: you're not seeing the class itself cause i stumbled upon this while writing Mocha unit tests for my DB access &amp; i only need to pass the data interface MyData.
## [10][Setting up VS project](https://www.reddit.com/r/typescript/comments/ftmgu1/setting_up_vs_project/)
- url: https://www.reddit.com/r/typescript/comments/ftmgu1/setting_up_vs_project/
---
Greetings,

sorry, i'm simply getting crazy. How do i create a TypeScript project in visual studio? I've installed the modules: "support for JavaScript and TypeScript" and "TypeScript 3.8 SDK", but what now? I don't want an angular project, i don't want a node project, i simply need to create a webpage that runs my script, and i don't want to write it in javascript.

Can someone direct me pls? I've only used Visual Studio for C++ until now.

(note Visual Studio the IDE, not Visual Studio Code the baby text editor)

Thanks!
## [11][Best way of figuring out the proper type to use?](https://www.reddit.com/r/typescript/comments/ftfqcf/best_way_of_figuring_out_the_proper_type_to_use/)
- url: https://www.reddit.com/r/typescript/comments/ftfqcf/best_way_of_figuring_out_the_proper_type_to_use/
---
TypeScript can be occasionally fussy about types, which is what I want, but at the same time it can be annoying, some things I've basically given up on figuring out and just put "any." Which kind of defeats the purpose of using it.   I'm learning React with TypeScript at the moment. Mostly it's okay, there are a few times that I've been defeated by React/TypesScript typings.  So where do you go for figuring out the proper types for variables, methods etc?
