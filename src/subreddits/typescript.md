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
## [2][Auto generate type-checker from Typescript types?](https://www.reddit.com/r/typescript/comments/ges5r9/auto_generate_typechecker_from_typescript_types/)
- url: https://www.reddit.com/r/typescript/comments/ges5r9/auto_generate_typechecker_from_typescript_types/
---
Does a library exist that automatically generates code to validate your Typescript types. I'm pretty sure the answer is no, but just wanna confirm before I build a babel-macro that does this.

Would people find such a tool useful?

Example:

`interface Foo {`

`bar: string`

`}`

`const validator = createValidator("Foo") // compile time babel macro`

`// validator is now a function that can be called at runtime to check whether an object matches the Foo interface`
## [3][Documenting API responses](https://www.reddit.com/r/typescript/comments/gewoym/documenting_api_responses/)
- url: https://www.reddit.com/r/typescript/comments/gewoym/documenting_api_responses/
---
Curious how other people handle API responses. The apps I develop the front end for are built with mostly API integrations. 

Sometimes I think it's easier to just do 

```
{ [key: string]: any; }
```

I have built out the response types for the convenience of intellicense.
## [4][How to handle async restful response return ?](https://www.reddit.com/r/typescript/comments/gf0jnr/how_to_handle_async_restful_response_return/)
- url: https://www.reddit.com/r/typescript/comments/gf0jnr/how_to_handle_async_restful_response_return/
---
I have a case where request comes in trigger some processing, but the processing is async, and I do want to response return the result(success or not) from the processing. How can I handle this case?

  
Here I cannot simply use await to stop for processing, because the piece of code in the req &amp; res level is just sending a message to message queue and actually it needs to wait the consumer to fully process it.  


My thought is to add eventemitter and listen on event in the req &amp; res level, and emit the event after the consumer processed it. (My service and consumer are in same node process). But How can I handle different requests? There could be racing problems, any better ideas?

Thanks
## [5][importing with "@"](https://www.reddit.com/r/typescript/comments/gef6u2/importing_with/)
- url: https://www.reddit.com/r/typescript/comments/gef6u2/importing_with/
---
I'm trying to find a best practice here. What's best?
Importing like this:  
`import { collectionRef} from '@mycoolpackage/prelude/control/function/collection';`  
When I run the `tsc` transpiler and importing like above, it doesn't find the import.

or like this:  
`import { collectionRef } from '../../prelude/control/function/collection';`  
This works when I run the `tsc` transpiler.

Are there any advantages?
## [6][Which is the preferred way to write a function? And can I mark one as pure somehow?](https://www.reddit.com/r/typescript/comments/gekwdq/which_is_the_preferred_way_to_write_a_function/)
- url: https://www.reddit.com/r/typescript/comments/gekwdq/which_is_the_preferred_way_to_write_a_function/
---
The TypeScript docs present functions in the style

    const add: (x: number, y: number) =&gt; number = function (
      x: number,
      y: number
    ): number {
      return x + y;
    };

or as an arrow function

    const add: (x: number, y: number) =&gt; number = (x: number, y: number): number =&gt;
      x + y;

`add` is declared with a type signature (if that's the correct term, I'm new to this), then a function is assigned with its own type signature.

But this doesn't seem to add any information over the much more readable options of either letting `add`'s type be inferred from a function literal that's already got fully specified types

    const add = (x: number, y: number): number =&gt; x + y;

Or declaring a type on `add` and then giving it a function literal without specified types

    const add: (x: number, y: number) =&gt; number = (x, y) =&gt; x + y;

My questions:

- Is there any reason to use the more verbose example #1 over my preference #3?
- Is there any way I can have the TypeScript compiler guarantee that my function is pure and doesn't refer to or modify any outside data? I couldn't see this in the docs but I also don't fully understand everything I can see, so I thought I might be missing features or not understanding the power of it.

Thanks to anyone who can offer advice!
## [7][Foal TS - May release (version 1.8) - Dependency injection with interfaces, permission management, service initialization](https://www.reddit.com/r/typescript/comments/geg8up/foal_ts_may_release_version_18_dependency/)
- url: https://www.reddit.com/r/typescript/comments/geg8up/foal_ts_may_release_version_18_dependency/
---
Foal TS framework version 1.8 is officially released!

[Foal TS - May release \(version 1.8\)](https://preview.redd.it/xm3t9m4ev3x41.png?width=1606&amp;format=png&amp;auto=webp&amp;s=d9e0d848c980f799c5e0d25ffa1370a6c4b79a9b)

This version brings mainly **three new features**.

1. **Interfaces and generic classes** (such as TypeORM repositories) **can** now **be injected** into services and controllers using FoalTS dependency injection system.

[Example of service](https://preview.redd.it/umiskrw6w3x41.png?width=746&amp;format=png&amp;auto=webp&amp;s=b1d61cab00658afdfcf41891acaf9725358edd0b)

[Injection of the concrete instances.](https://preview.redd.it/jevwe7q8w3x41.png?width=747&amp;format=png&amp;auto=webp&amp;s=0f01dbc77d59ab54155884e308d9d90bdbf5a2ce)

Documentation : [https://foalts.gitbook.io/docs/topic-guides/architecture/services-and-dependency-injection#usage-with-interfaces-and-generic-classes](https://foalts.gitbook.io/docs/topic-guides/architecture/services-and-dependency-injection#usage-with-interfaces-and-generic-classes)

2. **Services** accept an optional `boot` **method for initialization**

[Example of \\"boot\\" method](https://preview.redd.it/68sm4vmqw3x41.png?width=745&amp;format=png&amp;auto=webp&amp;s=c0fa86820f28abe83d37158d3adcc203f515137e)

Documentation: [https://foalts.gitbook.io/docs/topic-guides/architecture/initialization#the-services-boot-method](https://foalts.gitbook.io/docs/topic-guides/architecture/initialization#the-services-boot-method)

3. `Group`, `Permission` and `UserWithPermissions` extend `BaseEntity` (TypeORM) so as to add new methods (find, save, etc) and `UserWithPermissions` has a new method `withPerm` which allows to retrieve all users with a given permission.

[Permission extends BaseEntity](https://preview.redd.it/8uzmpoq9x3x41.png?width=745&amp;format=png&amp;auto=webp&amp;s=5bf024e17155b16aeb4a6162b27d19787650a54c)

[Retrieve all users with the permission \\"perm1\\"](https://preview.redd.it/u9kqau8bx3x41.png?width=748&amp;format=png&amp;auto=webp&amp;s=6e8b612834a836fdf3e86d71d71d18587ad17955)

Documentation: [https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/groups-and-permissions#baseentity-inheritance](https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/groups-and-permissions#baseentity-inheritance)

&amp;#x200B;

Foal, in a few words, it's a Node.js framework:

* written in TypeScript
* provided with batteries included (Auth, OpenAPI, GraphQL, ORM, CLI, scripts, file storage)
* and with a simple and intuitive architecture (no magic, no over-engineering).

And the must: it has more than 11,000 lines of documentation.

[https://foalts.org](https://foalts.org/)

[https://github.com/FoalTS/foal](https://github.com/FoalTS/foal)

[https://foalts.gitbook.io/docs/](https://foalts.gitbook.io/docs/)

\#TypeScript #JavaScript #NodeJS #FoalTS #DI #JWT #permissions
## [8][Convert string to enum values](https://www.reddit.com/r/typescript/comments/gekbfj/convert_string_to_enum_values/)
- url: https://www.reddit.com/r/typescript/comments/gekbfj/convert_string_to_enum_values/
---
In my database, I have a string saved like this: `"[Value 1, Value 2, Value 3]"`    

In my frontend I have an enum that looks like this:    

    export enum Values {
        Value1 = 'Value 1',
        Value2 = 'Value 2',
        Value3 = 'Value 3'
    }

My question is when I get this string from my backend, how can I convert it into my enum? Is this possible?
## [9][Has anyone written a Typescript to Documentation generator?](https://www.reddit.com/r/typescript/comments/geo0ih/has_anyone_written_a_typescript_to_documentation/)
- url: https://www.reddit.com/r/typescript/comments/geo0ih/has_anyone_written_a_typescript_to_documentation/
---
This has been a big point of frustration for me. Typescript has all of the parameters, types &amp; 95% of what you need to automatically generate documentation. But... I can't find a tool that actually reads the types &amp; generates documentation from them.

Does something like this exist?
## [10][ts-transformer-properties-rename - a custom transfomer for TypeScript which renames all private and internal properties, so you can mangle them with uglify-js/terser and reduce size of your bundles](https://www.reddit.com/r/typescript/comments/ge56dn/tstransformerpropertiesrename_a_custom_transfomer/)
- url: https://github.com/timocov/ts-transformer-properties-rename
---

## [11][How do you declare a variable of type T that extends type U without "losing" type T?](https://www.reddit.com/r/typescript/comments/ge96nf/how_do_you_declare_a_variable_of_type_t_that/)
- url: https://www.reddit.com/r/typescript/comments/ge96nf/how_do_you_declare_a_variable_of_type_t_that/
---
e.g. I want to declare some record `foo`, where the values are all of a certain type `Bar`, so `foo` must extend `Record&lt;string, Bar&gt;` (`{ [ key: string ]: Bar}`).

I can simply write `const foo: Record&lt;string, Bar&gt; = { fooKeyA: bar1, fooKeyB: bar2 };` but now all that I can extract from the declared `foo` is that it is of type `Record&lt;string, Bar&gt;`, which means I only know that its keys are `string`, i.e. I have lost the key names `fooKeyA` and `fooKeyB`.

Here is a workaround:

&gt;Declare a function that does nothing other than return its argument, and ensure that its argument extends a certain type. Anytime we declare something like foo, wrap the actual value we want to use in this function. i.e   
`function iAmTotallyUselessAtRuntime&lt;T extends Record&lt;string, Bar&gt;&gt;(foo: Foo): Foo { return foo; }`  
`const foo = iamTotallyUselessAtRuntime({ fooKeyA: bar1, fooKeyB: bar2 });`  
Now foo has maintained its original type with its unique keys, and also we can be sure that it also extends the type it needs to extend.  
However, (as I made clear in the function name) we are declaring a totally useless function that does nothing meaningful at run time. 

I feel that there must be a better way.

This seems like it should be so simple, but I can't really think of an efficient way to do it.
