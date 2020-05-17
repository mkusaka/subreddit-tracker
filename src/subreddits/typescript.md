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
## [2][Use io-ts to validate data without isLeft/isRight stuff?](https://www.reddit.com/r/typescript/comments/gl9x8z/use_iots_to_validate_data_without_isleftisright/)
- url: https://www.reddit.com/r/typescript/comments/gl9x8z/use_iots_to_validate_data_without_isleftisright/
---
I swear I something in io-ts that allows you to check whether an arbitrary object matches a io-ts type without the entire isLeft/isRight ceremony. Just a simple boolean value. Does anyone remember what this method in the library is?
## [3][While assigning value to a property, how to enforce the value of that property to belong to only values defined in string literal type of the property key?](https://www.reddit.com/r/typescript/comments/glcsop/while_assigning_value_to_a_property_how_to/)
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
## [4][[Suggest] Safe decorator implementation with ^ type operator.](https://www.reddit.com/r/typescript/comments/glec40/suggest_safe_decorator_implementation_with_type/)
- url: https://github.com/microsoft/TypeScript/issues/38620
---

## [5][Why does this happen?](https://www.reddit.com/r/typescript/comments/gl3mxz/why_does_this_happen/)
- url: https://www.reddit.com/r/typescript/comments/gl3mxz/why_does_this_happen/
---
    function foo(p: {x:number,y:number}) {
    	return p;
    }
    
    let obj = {x : 1,y :2};
    let parameterObj = Object.assign({z:3},obj);
    foo(parameterObj);//does not lint error
    foo({x:1,y:2,z:3});//lints error
## [6][mapping errors from js code to ts code](https://www.reddit.com/r/typescript/comments/gld4qy/mapping_errors_from_js_code_to_ts_code/)
- url: https://www.reddit.com/r/typescript/comments/gld4qy/mapping_errors_from_js_code_to_ts_code/
---
So when I compile a `.ts` file to a `.js` file and then run the `.js` file and get errors at line `x` column `y` I can not look at my `.ts` file at line `x` column `y` because the `.js` code has code at different places than the corresponding code in the `.ts` file . Is there any way to fix that?
## [7][linting types](https://www.reddit.com/r/typescript/comments/glcpew/linting_types/)
- url: https://www.reddit.com/r/typescript/comments/glcpew/linting_types/
---
How to solve that (I am using vscode):

1. eslint shows that I do not use **imported** types , despite the fact that I  use them.
2. eslint does shows error for types that I have **defined** but do not use, but when I hover over the unused type I see a ts error `"type" is declared but never used.ts(6196)`. However there is no red highlighting . Just the color of syntax highlighter is deem and is difficult to spot .

What I have done so far :

    npm install --save-dev eslint typescript
    npx eslint --init

I have choosen that I use typescript and esmodules and installed the extra packages suggested by eslint.
## [8][Build a Cryptocurrency Trading Exchange App in React Native](https://www.reddit.com/r/typescript/comments/gl4g5n/build_a_cryptocurrency_trading_exchange_app_in/)
- url: https://www.reddit.com/r/typescript/comments/gl4g5n/build_a_cryptocurrency_trading_exchange_app_in/
---
The tutorial will go through how to get started with the Web3 library to handle transactions and interact with a Blockchain in a React Native app. 

[https://www.education-ecosystem.com/glenettn/2gzEB-how-to-build-a-cryptocurrency-trading-exchange-app-in-react-native/voGo5-01-introduction/](https://www.education-ecosystem.com/glenettn/2gzEB-how-to-build-a-cryptocurrency-trading-exchange-app-in-react-native/voGo5-01-introduction/)
## [9][I created a zero-dependncy promise-based job queue with concurrency limiting](https://www.reddit.com/r/typescript/comments/gksmub/i_created_a_zerodependncy_promisebased_job_queue/)
- url: https://github.com/Tanuel/async-queue
---

## [10][What does this code in the express source code mean?](https://www.reddit.com/r/typescript/comments/glahd9/what_does_this_code_in_the_express_source_code/)
- url: https://www.reddit.com/r/typescript/comments/glahd9/what_does_this_code_in_the_express_source_code/
---
This might be a simple question but I search everywhere and couldn't find the answer.

`export interface IRouterMatcher&lt;T, Method extends 'all' | 'get' | 'post' | 'put' | 'delete' | 'patch' | 'options' | 'head' = any&gt; {`  
 `// tslint:disable-next-line no-unnecessary-generics (This generic is meant to be passed explicitly.)`  
`&lt;P extends Params = ParamsDictionary, ResBody = any, ReqBody = any, ReqQuery = Query&gt;(path: PathParams, ...handlers: Array&lt;RequestHandler&lt;P, ResBody, ReqBody, ReqQuery&gt;&gt;): T;`  
 `// tslint:disable-next-line no-unnecessary-generics (This generic is meant to be passed explicitly.)`  
`&lt;P extends Params = ParamsDictionary, ResBody = any, ReqBody = any, ReqQuery = Query&gt;(path: PathParams, ...handlers: Array&lt;RequestHandlerParams&lt;P, ResBody, ReqBody, ReqQuery&gt;&gt;): T;`  
`(path: PathParams, subApplication: Application): T;`  
`}`

 It seems like this is the interface for 'get, post put, delete...' but what does this line mean?

`Method extends 'all' | 'get' | 'post' | 'put' | 'delete' | 'patch' | 'options' | 'head' = any`

Is `method` also a generic in here? If it is then why it isn't used inside the interface? Also what is this whole thing 

`'all' | 'get' | 'post' | 'put' | 'delete' | 'patch' | 'options' | 'head'`

Are they also some kinds of interface? 

&amp;#x200B;

One more question, how come express is written in typescript but I can require this module in pure JS ? Sorry for asking so many questions :(
## [11][I created a simple, dependency free web server with live reload for Deno called Denoliver (sorry..)](https://www.reddit.com/r/typescript/comments/gl3g5e/i_created_a_simple_dependency_free_web_server/)
- url: https://www.reddit.com/r/typescript/comments/gl3g5e/i_created_a_simple_dependency_free_web_server/
---
Hey everyone,

I wanted to get into Deno and see what it had to offer during the weekend so I decided to write a small file server that:

* Is dependency free
* Has live reload
* Supports client routing

All written in Typescript of course :)

It was heavily inspired by Luke Jackssons [Servor](https://github.com/lukejacksonn/servor)

Project link: [https://github.com/joakimunge/denoliver](https://github.com/joakimunge/denoliver)

Would love if you wanted to try it out, or contribute!
