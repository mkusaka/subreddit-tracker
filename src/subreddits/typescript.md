# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][Optional chaining](https://www.reddit.com/r/typescript/comments/fe4r99/optional_chaining/)
- url: https://www.reddit.com/r/typescript/comments/fe4r99/optional_chaining/
---
For some reason I can't use optional chaining in my new project, other TS features like interfaces, types, ... are working fine. Also my IDE is not throwing out any errors, but upon compiling the code, I get this error for the following line of code:

Line of code: `country: location?.country`

`country: location?.country,`

`SyntaxError: Unexpected token '.'`

`at Module._compile (internal/modules/cjs/loader.js:895:18)`

What's going on?
## [3][How to replace all non-const enum value with literals? And should I do it?](https://www.reddit.com/r/typescript/comments/fe9e7a/how_to_replace_all_nonconst_enum_value_with/)
- url: https://www.reddit.com/r/typescript/comments/fe9e7a/how_to_replace_all_nonconst_enum_value_with/
---
## Problem

I want to optimize my TypeScript by replacing all literal enums with their values.

## Why not `const enum`?

* Third-party libraries may not provide them.
* Babel does not support them.
## [4][Best way to type an object with string keys?](https://www.reddit.com/r/typescript/comments/fe40sr/best_way_to_type_an_object_with_string_keys/)
- url: /r/LearnTypescript/comments/fe3zip/best_way_to_type_an_object_with_string_keys/
---

## [5][Different function return types?](https://www.reddit.com/r/typescript/comments/fe2mgz/different_function_return_types/)
- url: https://www.reddit.com/r/typescript/comments/fe2mgz/different_function_return_types/
---
Is it possible to have a function return different return types under different circumstances?

Edit: I mean without overloading.

example:

    const test = (name: string) : int | string =&gt; {
    
        if(name === 'beatdook04') return 200;
    
        return "Who are you?";    
    
    }
    
    test('beatdook04');
## [6][TypeBox - Mapping TypeScript as JSONSchema](https://www.reddit.com/r/typescript/comments/fdoolj/typebox_mapping_typescript_as_jsonschema/)
- url: https://github.com/sinclairzx81/typebox
---

## [7][Trying to store typescript code in string to use in runtime. [Guidance Appreciated]](https://www.reddit.com/r/typescript/comments/fdwf0e/trying_to_store_typescript_code_in_string_to_use/)
- url: https://www.reddit.com/r/typescript/comments/fdwf0e/trying_to_store_typescript_code_in_string_to_use/
---
As the description implies i'm trying to store actual code from typescript file in a variable so that when I console.log it i'll see the typescript implementation instead of the compiled code.  
Have any of you tried anything similar?   


What I've tried:  
1. Using decorators and Reflection,  but struggle getting the actual code (as a string) as it is in the file.

2. looking at the typescript npm package. There's just so many functions to read up on. 

3.  Tried to retrofit this code [this code](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API#user-content-re-printing-sections-of-a-typescript-file) from the typescript compiler API docs for my purposes.

So far I haven't had any luck. All pointers welcome ;)
## [8][How to conditionally determine return value in a Promise](https://www.reddit.com/r/typescript/comments/fdtpov/how_to_conditionally_determine_return_value_in_a/)
- url: https://www.reddit.com/r/typescript/comments/fdtpov/how_to_conditionally_determine_return_value_in_a/
---
I'm trying to write a wrapper around fs.readFile that will return JSON if the option is set to true.

Having a hard time to let TS figure out when it should return JSON and a string.

Trying to avoid having to do this (since I already know it is going to be JSON):

\`\`\`

const f = await readFile(…, { json: true });

if (typeof f !== 'string') {

  return f.a;

}

\`\`\`

https://preview.redd.it/peb06d2q6uk41.png?width=1300&amp;format=png&amp;auto=webp&amp;s=aec55267486ff42d3f9720d2e6f926c32b18ab6c
## [9][Static Site with JS from TS using exports?](https://www.reddit.com/r/typescript/comments/fdtoqt/static_site_with_js_from_ts_using_exports/)
- url: https://www.reddit.com/r/typescript/comments/fdtoqt/static_site_with_js_from_ts_using_exports/
---
Could anyone please explain to me (Java/Kotlin lad) how to get modules working? I am very comfortable with vanilla JS but was busy playing with the JVM when the whole Node / ES module thing went down.

I'm not creating a Node app here... just transpiling multiple TS files to JS, for use in a very basic, boring site.

I have satisfied both my ATE (VSC, fwiw) and the TSC (transpiler irl?) by using the `export` and `import` keywords (at least the notion of *import* is familiar) ... but when I load the app, on good ol' localhost, in browser, i'm understandably seeing the *exports undefined* exception in the console (RIP Firebug)... is it as simple as including a `common.js` script in the client, to satisfy this unfamiliar concept of *exports* or am I going about this wrong?

I just want to visit my static site with multiple, minifed JS files that were generated by TSC, understanding that multiple TS files rely on definitions from multiple files without too many complications in the build / compilation process..?

Cheers lads.
## [10][TypeScript Migration Strategy (from JavaScript)](https://www.reddit.com/r/typescript/comments/fdhx6f/typescript_migration_strategy_from_javascript/)
- url: https://www.youtube.com/watch?v=v3lI29trIN8&amp;feature=share
---

## [11][Is this even possible in TS?](https://www.reddit.com/r/typescript/comments/fdkxcp/is_this_even_possible_in_ts/)
- url: https://www.reddit.com/r/typescript/comments/fdkxcp/is_this_even_possible_in_ts/
---
Hi,

If I have a union type like:

    type KeyType = 'one' | 'two' | 'three';

And I want to declare an interface that has all those keys (of a given type, say number) plus additional properties. E.g.:

    interface MyInterface {
        [key in KeyType]: number;
        anotherProperty: string;
    }

Basically I want an interface that ensures that keys won't clash, and specify that all the keys in a given union will be there, plus any number of additional properties I like.

The above doesn't work.

Thanks!
