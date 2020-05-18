# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][typecheck.macro - Automatically generate validation functions for Typescript types.](https://www.reddit.com/r/typescript/comments/glyq1m/typecheckmacro_automatically_generate_validation/)
- url: https://www.reddit.com/r/typescript/comments/glyq1m/typecheckmacro_automatically_generate_validation/
---
10 days ago, I made [this post](https://www.reddit.com/r/typescript/comments/ges5r9/auto_generate_typechecker_from_typescript_types/) asking whether it was possible to automatically generate validation functions for Typescript types. Unfortunately, it seemed like that was not the case.

So I built a library (well, compile time macro) to do it: [https://github.com/vedantroy/typecheck.macro](https://github.com/vedantroy/typecheck.macro)

 Once you configure babel macros, you can seamlessly generate validation functions for your Typescript types. No need to write your types again in a DSL (like runtypes/zod/io-ts). Compile time macros automatically parse your type declarations and generate validation functions.

The library is pretty fresh, but it already handles unions/tuple types/optional types/index signatures/interfaces (no extending yet)/type aliases/generics/object patterns/arrays, so it's pretty functional/usable.

The resulting validation functions are also really fast because most other libraries (except ajv) don't do compile time code generation. After some bench marking (not that I'm a benchmarking expert), I found it was up to 3 times as fast as ajv.

This is my first open source library, so I'm excited to hear what you guys think!

Above all, I feel like this can be a community effort. I'm open to suggestions/any form of feedback/contribution/etc.
## [3][Help with @HostListener("window:beforeunload", ["$event"])](https://www.reddit.com/r/typescript/comments/glwwmi/help_with_hostlistenerwindowbeforeunload_event/)
- url: https://www.reddit.com/r/typescript/comments/glwwmi/help_with_hostlistenerwindowbeforeunload_event/
---
 

Hi, my requirement is upon refreshing or closing the browser tab, there should be my custom alert

instead of browser alert. I have to get this work done in typescript so I am returning the

$event.returnValue = false on beforeunload event which will stop my browser from closing or

refreshing but at the same time I want my custom pop-up to be shown instead of browser pop-up.

Here is my Code -

      @HostListener("window:beforeunload", ["$event"])
      beforeUnload($event: Event) {
    //loading my custom pop up
          $event.returnValue =false;
    }

 My Custom Pop-up loads but loads after I cancel the browser popup- 

&amp;#x200B;

[Browser Popup](https://preview.redd.it/93ls5ftj7hz41.png?width=557&amp;format=png&amp;auto=webp&amp;s=989a3847460893e33f634256c31fea49dab7824f)

 How can I achieve this? what am I doing wrong? Please help.
## [4][Is this the correct way to handle an input with a mask?](https://www.reddit.com/r/typescript/comments/gllnww/is_this_the_correct_way_to_handle_an_input_with_a/)
- url: https://www.reddit.com/r/typescript/comments/gllnww/is_this_the_correct_way_to_handle_an_input_with_a/
---
I'm creating a reusable text input (functional component). It's a simple text input. The input will have a phone prop that when it's true, it will add a phone mask to the value it returns to it's parent.

    return (
    &lt;input id={props.id} type={props.type} placeholder={props.placeholder} ref={ref} onChange={handleChange} /&gt;
    );

And the change is handled with this function:

    function handleChange(event: React.ChangeEvent&lt;HTMLInputElement&gt;) {

    // format phone mask if phone prop is true
    if(props.phone) {
      event.target.value = "phone mask will be added here";
    }

    if(props.onChange) { 
        props.onChange(event); 
    }
   }

What I'm trying to achieve is that the state of all the form values is handled by the parent of all inputs.

The input achieves that, but I feel it's kind of hacky to assign a value like that to event.target.value. It's about the second time I'm using typescript so I'm pretty this is not the correct way.

Any help?
## [5][Use io-ts to validate data without isLeft/isRight stuff?](https://www.reddit.com/r/typescript/comments/gl9x8z/use_iots_to_validate_data_without_isleftisright/)
- url: https://www.reddit.com/r/typescript/comments/gl9x8z/use_iots_to_validate_data_without_isleftisright/
---
I swear I something in io-ts that allows you to check whether an arbitrary object matches a io-ts type without the entire isLeft/isRight ceremony. Just a simple boolean value. Does anyone remember what this method in the library is?
## [6][How would I add types to dynamic modeling library?](https://www.reddit.com/r/typescript/comments/glgrnw/how_would_i_add_types_to_dynamic_modeling_library/)
- url: https://www.reddit.com/r/typescript/comments/glgrnw/how_would_i_add_types_to_dynamic_modeling_library/
---
I recently wrote a dynamodb library (https://github.com/tywalch/electrodb) that helps folks model their access patterns by supplying a schema and dynamically implementing an API for that schema. 

As an example, the user would  supply:

    const {Entity} = require("electrodb");
    const tasks = new Entity({
        attributes: {
            id: {
                type: "string",
                default: () =&gt; uuid()
            },
            user: {
                type: "string",
                required: true
            },
            project: {
                type: "string",
                required: true
            },
            name: {
                type: "string",
                required: true
            },
            description: {
                type: "string",
                required: false
            },
            isActive: {
                type: "boolean",
                required: true
            }
        },
        indexes: {
            task: {
                pk: {
                    field: "pk",
                    facets: ["id"]
                },
                sk: {
                    field: "sk",
                    facets: [""]
                }
            },
            assigned: {
                index: "gsi1pk-gsi1sk-index"
                pk: {
                    field: "gsi1pk",
                    facets: ["employee"]
                },
                sk: {
                    field: "gsi1sk",
                    facets: ["project", "id"]
                }
            }
        }
    })

The above example would yield an dynamic api implementing that model, for example:

    // "task" method comes from index "task." 
    // "id" comes from facets value "id" under pk.
    tasks.query.task({id: "1"}).go()

    // "assigned" method comes from index "assigned".
    // "user" comes from facets value "user" under pk.
    // "project" comes from facets value "project" under sk.
    tasks.query.assigned({user: "johnsmith"})
        .between(
            {project: "project-a"},
            {project: "project-z"})
        .go()
    
    // "description" property in required so no error expected.
    // "isActive" property is required but not included so error expected.
    tasks.put({
        id: "23", 
        project: "project",
        user: "johnsmith",
        name: "task1"
    }).go()

It feels like this sort of API isnt really something that could be implemented in typescript, but I'm interested in how close I could get. It feels like the only way I'd accomplish this would be to write some sort of interface generation tool to parse and build the interfaces for the user.
## [7][While assigning value to a property, how to enforce the value of that property to belong to only values defined in string literal type of the property key?](https://www.reddit.com/r/typescript/comments/glcsop/while_assigning_value_to_a_property_how_to/)
- url: https://www.reddit.com/r/typescript/comments/glcsop/while_assigning_value_to_a_property_how_to/
---
Very new to typescript. Been scratching my head on how to do in the right way for sometime now.

# Context

Im using hapijs, and was trying to setup the routes. The `server.route` method takes in a array of route configuration objects `&lt;ServerRoute[]&gt;`. In each of the `ServerRoute`, there is a property called `method` which has the following type

    /**
     * (required) the HTTP method. Typically one of 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', or 'OPTIONS'. Any HTTP method is allowed, except for 'HEAD'. Use '*' to match against any HTTP method.
     */
    method: Util.HTTP_METHODS_PARTIAL | Util.HTTP_METHODS_PARTIAL[] | string | string[];

where the `Util.HTTP_METHODS_PARTIAL` is:

    type HTTP_METHODS_PARTIAL =
    'GET'
    | 'POST'
    | 'PUT'
    | 'PATCH'
    | 'DELETE'
    | 'OPTIONS'
    | HTTP_METHODS_PARTIAL_LOWERCASE;

# What I have tried

This is how i create the routes:-

    server.route(&lt;ServerRoute[]&gt;[
        {
            "method": "GETzz", // here "GETzz" is not a valid UTIL.HTTP_METHODS_PARTIAL
            "path": "/",
            "handler": rootHandler().get,
        }
    ]);

I tried asserting by `"method": &lt;Util.HTTP_METHODS_PARTIAL&gt;"GETzz";`, which gives me auto-completions for that type inside quotes, but the value is still not checked to be part of `UTIL.HTTP_METHODS_PARTIAL` or not.

(Edit)But take this example:-

    type Methods = "valid1" | "valid2";
    const obj:{method:Methods} = {
            method: "inValidValue",
    }

The above would result in the error that i was hoping for, which's `Type '"inValidValue"' is not assignable to type 'Methods'.`

# Question

The question is how to enforce the property `method` to only be assigned values that exists in type `UTIL.HTTP_METHODS_PARTIAL`? Currently i can give any value like `GETzz` etc but still no error is thrown, until I run the server and a request is made to the corresponding `path`. The information from that error is unreadable and in no way helps in pointing out that, giving the value of `method` as `GETzz` is why the server failed.

(Edit)Is it a problem with the type definition files of hapijs in DefinitelyTyped?

(Edit)The issue can be solved by removing the \`string | string\[\]\` from the type definition for hapi from DefinitelyTyped as \`Util.HTTP\_METHODS\_PARTIAL\` already has all the necessary header methods. I'm not sure how good of a solution this is. But this temporary hack works. 
## [8][linting types](https://www.reddit.com/r/typescript/comments/glcpew/linting_types/)
- url: https://www.reddit.com/r/typescript/comments/glcpew/linting_types/
---
How to solve that (I am using vscode):

1. eslint shows that I do not use **imported** types , despite the fact that I  use them.
2. eslint does shows error for types that I have **defined** but do not use, but when I hover over the unused type I see a ts error `"type" is declared but never used.ts(6196)`. However there is no red highlighting . Just the color of syntax highlighter is deem and is difficult to spot .

What I have done so far :

    npm install --save-dev eslint typescript
    npx eslint --init

I have choosen that I use typescript and esmodules and installed the extra packages suggested by eslint.
## [9][mapping errors from js code to ts code](https://www.reddit.com/r/typescript/comments/gld4qy/mapping_errors_from_js_code_to_ts_code/)
- url: https://www.reddit.com/r/typescript/comments/gld4qy/mapping_errors_from_js_code_to_ts_code/
---
So when I compile a `.ts` file to a `.js` file and then run the `.js` file and get errors at line `x` column `y` I can not look at my `.ts` file at line `x` column `y` because the `.js` code has code at different places than the corresponding code in the `.ts` file . Is there any way to fix that?
## [10][Safe decorator implementation with "^" type operator.](https://www.reddit.com/r/typescript/comments/glfkra/safe_decorator_implementation_with_type_operator/)
- url: https://github.com/microsoft/TypeScript/issues/38620
---

## [11][Why does this happen?](https://www.reddit.com/r/typescript/comments/gl3mxz/why_does_this_happen/)
- url: https://www.reddit.com/r/typescript/comments/gl3mxz/why_does_this_happen/
---
    function foo(p: {x:number,y:number}) {
    	return p;
    }
    
    let obj = {x : 1,y :2};
    let parameterObj = Object.assign({z:3},obj);
    foo(parameterObj);//does not lint error
    foo({x:1,y:2,z:3});//lints error
