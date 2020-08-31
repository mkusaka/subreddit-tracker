# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][How do I create a data access layer correctly?](https://www.reddit.com/r/typescript/comments/ijtgcv/how_do_i_create_a_data_access_layer_correctly/)
- url: https://www.reddit.com/r/typescript/comments/ijtgcv/how_do_i_create_a_data_access_layer_correctly/
---
I use TypeORM to work with the database, but I think it would be better to write my own abstraction (correct me if I'm wrong). In DAL, do I need to describe competitive methods for getting data, such as getUser (id), or develop generic methods for any data models, such as findOne&lt;User&gt;(args)?  

Can you recommend resources on this topic?
## [3][Does using await mean all the catches will just go to the try -catch block?](https://www.reddit.com/r/typescript/comments/ijl6il/does_using_await_mean_all_the_catches_will_just/)
- url: https://www.reddit.com/r/typescript/comments/ijl6il/does_using_await_mean_all_the_catches_will_just/
---
I’m using like 3 awaits. I use them because I need the data of the 3 awaits before doing anything else. I put them all in a try catch block. If any of them fail it’ll go to the catch block right? How do I know which one failed?
## [4][fp-ts equivalent of Scala Option.foreach](https://www.reddit.com/r/typescript/comments/ijlruv/fpts_equivalent_of_scala_optionforeach/)
- url: https://www.reddit.com/r/typescript/comments/ijlruv/fpts_equivalent_of_scala_optionforeach/
---
So Scala has been my big introduction to functional programming and I love it. Now in my TS projects I'm embracing the fp-ts library, which clearly adheres to a different style of functional structures than Scala does.

One big thing I've been looking for is foreach. In Scala if you want to execute code against the contents of an Option (if there's a value present), foreach is how you do it. I don't see any similarly named function in fp-ts tho. I've been going over the docs but I'm a bit unclear on what to use.

Guidance would be appreciated. Thanks.
## [5][Smart way to debug data - Online Editor and Chart Generator for JSON, XML, CSV, YAML documents https://jxcy.dev https://www.youtube.com/watch?v=N1ff9Sdq4w0](https://www.reddit.com/r/typescript/comments/ijh69d/smart_way_to_debug_data_online_editor_and_chart/)
- url: https://i.redd.it/udbudxkji6k51.png
---

## [6][TSDoc annotation or describing objects in code help?](https://www.reddit.com/r/typescript/comments/ijnerg/tsdoc_annotation_or_describing_objects_in_code/)
- url: https://www.reddit.com/r/typescript/comments/ijnerg/tsdoc_annotation_or_describing_objects_in_code/
---
I inherited a project, and they they didn't use types anywhere, even in obvious places. Is it possible in TSDoc (or an alternative) specify the types of variables and the structure of objects like you can in JS/PHPDoc?

The docs make me think that they removed that in favor of using the built-in strict typing. Am I mistaken?

If the above won't work, is there a compromise between comments and the built-in typing? I don't want to add types to code unless I'm actually working on it, where as I don't mind taking a minute and adding comments to an entire class.
## [7][stack traces not pointing back to source (ts) files](https://www.reddit.com/r/typescript/comments/ijc1g3/stack_traces_not_pointing_back_to_source_ts_files/)
- url: https://www.reddit.com/r/typescript/comments/ijc1g3/stack_traces_not_pointing_back_to_source_ts_files/
---
Other than `sourceMap: true`, do any other settings need to be configured to get stack traces to point to typescript files? My stack traces are hard to follow because they point to the compiled js files.

current settings in tsconfig.json:

    {
      "compilerOptions": {
        "target": "es5",
        "lib": [
          "dom",
          "dom.iterable",
          "esnext"
        ],
        "allowJs": true,
        "skipLibCheck": true,
        "esModuleInterop": true,
        "allowSyntheticDefaultImports": true,
        "strict": true,
        "forceConsistentCasingInFileNames": true,
        "module": "CommonJS",
        "moduleResolution": "node",
        "resolveJsonModule": true,
        "isolatedModules": true,
        "sourceMap": true,
        "rootDir": "src",
        "outDir": "compiled"
      },
      "include": [
        "src", "src/__mocks__"
      ],
      "exclude": ["node_modules", "**/*.test.ts", "compiled"]
    }
## [8][TypeScript First Test Framework](https://www.reddit.com/r/typescript/comments/ijie1y/typescript_first_test_framework/)
- url: https://www.reddit.com/r/typescript/comments/ijie1y/typescript_first_test_framework/
---
Hi. I am starting to work on a TypeScript first test framework.

All major JavaScript frameworks are JavaScript first and require transpilation before files can be tested. This includes Jest, Mocha and Ava, with different degrees of support.

The new test framework I am proposing would need to implement the following features:

1. Support TypeScript source files and TypeScript test files out of the box.
2. Still be BDD based, with \`describe\` and \`it\` keywords.
3. Support arrow functions for test definitions, passing the test context to the callbacks.
4. Support async and parallel testing.
5. MIT License.

I was wondering if anyone would be interested to be day-1 collaborators on this project?

Thanks!
## [9][Ways to ensure the properties of an interface are of certain types?](https://www.reddit.com/r/typescript/comments/iixvio/ways_to_ensure_the_properties_of_an_interface_are/)
- url: https://www.reddit.com/r/typescript/comments/iixvio/ways_to_ensure_the_properties_of_an_interface_are/
---
Currently, when extending a type that has index signatures, like this:
    
    type DataRecord = { [key in string]: string | number };
    interface LoginRecord extends DataRecord {
      account: string,
      time: number,
    }
No compilation error will occur when assigning a object with more properties to it,

    const record: LoginRecord = {
      account: 'Admin',
      time: 0,
      password: 'admin',
    }; // no error

Currently I'm using a utility type to help me:

    type ExtendProperty&lt;B, T extends Record&lt;keyof T, B&gt;&gt; = T;
    type RecordProperty = string | number;
    type LoginRecord = ExtendProperty&lt;RecordProperty, {
      account: string,
      time: number,
    }&gt;;
    const record: LoginRecord = {
      account: 'Admin',
      time: 0,
      password: 'admin',
    }; // error: ...'password' does not exist in type...
I wonder if there are other ways to do it?
## [10][Spotify shuffler](https://www.reddit.com/r/typescript/comments/ijebkz/spotify_shuffler/)
- url: https://www.reddit.com/r/typescript/comments/ijebkz/spotify_shuffler/
---
I know absolutely nothing about coding, (went to a very basic course in html, wordpress and css and that's it) but I was wondering if there is a way you could reorganise your songs in your spotify playlist? Shuffle them around
## [11][Why is any nonsense assignable to InstanceType&lt;...&gt;[] here?](https://www.reddit.com/r/typescript/comments/iiqcv9/why_is_any_nonsense_assignable_to_instancetype/)
- url: https://www.reddit.com/r/typescript/comments/iiqcv9/why_is_any_nonsense_assignable_to_instancetype/
---
I want to create a function which accepts a class and a callback, which also accepts this class and returns its instances. It's not very practical, but it is a simplified example of my problem.

My code looks like this ([playground link](https://www.typescriptlang.org/play?ts=4.0.2#code/C4TwDgpgBAwg9gOwM7AE4FcDGw6qgXigQgHcAKAOioENUBzJALimoRAG0BdASgID4WbANwAoEQDN0CbAEtEUYBBQAeACpQIAD0UIAJkliIUGbLj5ladZqoA0USQmYX613vgEBJZMFaYIq8Ag1Pi5eAG8RKCioVAhgdFQEeylnOm5RAF8xTAAbaiQDAFkQGDyCqDCskQB6aqhVAAtoGQQdXQhdKHQkajpmgzgAayoKEUUUMmLS-KQ7AH1cmf4odmISKAWypDJuOzWNxYKdnlEauqoAI3RgKA8oTFYWbHRqHJyQGLiEpNpUahABuJBB8EEYIMhoOJUHAALbJBAAQjO9Sa9jgbzgJBadCgSAacHQOU6eMxgg0qGhqDsVxuMhuujgSgQAHJgEjxsBJiUtvNDgZ3CsAIwAJgAzHZmeI4HALrRmScgA)):

    type Constructor = new(...args: any[]) =&gt; any;

    function test&lt;T extends Constructor&gt;(arg: T, fun: (arg: T) =&gt; InstanceType&lt;T&gt;[]) 
    {
        return fun(arg);
    }

    class MyClass {}

    // The intended usage is ok...
    test(MyClass, _class =&gt; [new _class(), new _class()]);

    // ..but I can actually return arrays of any nonsense from fun!
    // The following should show an error, but it doesn't!
    test(MyClass, _class =&gt; [123, 'foobar']);

The second call to `test` allows me to return anything from the callback, despite its type clearly inferred as `(arg: typeof MyClass) =&gt; MyClass[]`. It seems wrong to me that a `(string | number)[]` return type is assignable to `MyClass[]`.

Is there a way to fix this and make it report the correct error?

Thanks!
