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
## [2][importing with "@"](https://www.reddit.com/r/typescript/comments/gef6u2/importing_with/)
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
## [3][How Our Stack Evolved in 10 Years](https://www.reddit.com/r/typescript/comments/gek2mm/how_our_stack_evolved_in_10_years/)
- url: https://typescript-open-source-resources.s3.us-east-2.amazonaws.com/the-stack-evolution.htm
---

## [4][ts-transformer-properties-rename - a custom transfomer for TypeScript which renames all private and internal properties, so you can mangle them with uglify-js/terser and reduce size of your bundles](https://www.reddit.com/r/typescript/comments/ge56dn/tstransformerpropertiesrename_a_custom_transfomer/)
- url: https://github.com/timocov/ts-transformer-properties-rename
---

## [5][Foal TS - May release (version 1.8) - Dependency injection with interfaces, permission management, service initialization](https://www.reddit.com/r/typescript/comments/geg8up/foal_ts_may_release_version_18_dependency/)
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
## [6][How do you declare a variable of type T that extends type U without "losing" type T?](https://www.reddit.com/r/typescript/comments/ge96nf/how_do_you_declare_a_variable_of_type_t_that/)
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
## [7][Ultimate Vim TypeScript Setup](https://www.reddit.com/r/typescript/comments/gdv3al/ultimate_vim_typescript_setup/)
- url: https://pragmaticpineapple.com/ultimate-vim-typescript-setup/
---

## [8][Incorrect module resolution causing type errors](https://www.reddit.com/r/typescript/comments/ge8bqd/incorrect_module_resolution_causing_type_errors/)
- url: https://www.reddit.com/r/typescript/comments/ge8bqd/incorrect_module_resolution_causing_type_errors/
---
Hi all, I'm trying to import a package ([filestack-js](https://github.com/filestack/filestack-js)) into my typescript code which has multiple entry points defined in its `package.json` file.

When trying to run the typescript compiler i get type errors relating to not being able to find  '`@types/node'` (which I've excluded from the project as it causes other type problems with our browser typings).

&amp;#x200B;

I did some digging and saw that the typescript compiler is using the '`main`' entry point as described in the `package.json` file. However, the package authors only intend this entry point to be used in a node js environment. I see that they have defined a '`browser`' entry point that doesn't have these problems but the compiler is favoring main instead. Anyway I can force this resolution?

&amp;#x200B;

I imagine this isn't a unique problem, how do others resolve these conflicts?

&amp;#x200B;

If I set `skipLibCheck` to `true` the errors go away but that feels like overkill to me. I'm also using webpack to bundle my compiled js files and that happily resolves to the correct source code.
## [9][I have just released Jackson-js: Powerful JavaScript decorators to serialize/deserialize objects into JSON and vice versa](https://www.reddit.com/r/typescript/comments/ge1jth/i_have_just_released_jacksonjs_powerful/)
- url: https://medium.com/@pichillilorenzo/df952454cf?source=friends_link&amp;sk=a65bd247eca2f95fdfddda34447a6db6
---

## [10][VIM Users: ale vs tsuquyomi?](https://www.reddit.com/r/typescript/comments/ge7y2d/vim_users_ale_vs_tsuquyomi/)
- url: https://vi.stackexchange.com/questions/25111/typescript-ale-vs-tsuquyomi
---

## [11][Difference between property and method in Typescript](https://www.reddit.com/r/typescript/comments/ge7evn/difference_between_property_and_method_in/)
- url: https://www.reddit.com/r/typescript/comments/ge7evn/difference_between_property_and_method_in/
---
What is the difference between a property and method in Typescript? I am trying to add a property to each of my functions in a class. It works fine with anonymous functions but with methods tsserver doesn't display the correct type. Any help would be appreciated here is a [stackblitz](https://stackblitz.com/edit/typescript-1jgdjg)

    type Func = (...args: any[]) =&gt; any;
    
    /** Add a property to the function */
    interface Spy&lt;Fn extends Func&gt; {
      (...params: Parameters&lt;Fn&gt;): ReturnType&lt;Fn&gt;;
      extraProperty: string;
    }
    /** Loop through all properties of T, if the property is a function, make it a spy */
    type MakeSpys&lt;T&gt; = { [x in keyof T]: T[x] extends Func ? Spy&lt;T[x]&gt; : false };
    
    class Class {
      method(): null {
        return null;
      }
      property = (): null =&gt; null;
    }
    
    let _: Class['method']; // type: () =&gt; null (expected)
    let __: Class['property']; //type : () =&gt; null (expected)
    let spyObj: MakeSpys&lt;Class&gt;;
    spyObj.property; // type: Spy&lt;() =&gt; null&gt; (expected)
    spyObj.method; // type: method(): null (why?????)
    const method = spyObj.method; // type: Spy&lt;() =&gt; null&gt; (expected)

spyObj.method should be of type Spy&lt;() =&gt; null&gt; so that the reader is aware that there was a property added to the function
