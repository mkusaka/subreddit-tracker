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
## [2][How to handle __dirname transpilation?](https://www.reddit.com/r/typescript/comments/gwbb8s/how_to_handle_dirname_transpilation/)
- url: https://www.reddit.com/r/typescript/comments/gwbb8s/how_to_handle_dirname_transpilation/
---
TS transpiles `__dirname` to the folder `__dirname` was used in a class definition of. Do you guys just go along with this, or is there a workaround to have it evaluate to the folder of the calling file?
## [3][Naming convention for interfaces + decorators](https://www.reddit.com/r/typescript/comments/gwezqm/naming_convention_for_interfaces_decorators/)
- url: https://www.reddit.com/r/typescript/comments/gwezqm/naming_convention_for_interfaces_decorators/
---

I have this very specific issue where I cannot find a clean way of naming interfaces and decorators.

For example, I have a `Middleware` interface that I can use to create implementations like `AuthMiddleware`;
```ts
interface Middleware {
    handle():boolean;
}

class AuthMiddleware implements Middleware {
    handle(){ /* implementation here */ }
}
```

The problem arises when I need a decorator to add middleware to a class:

```ts
@Middleware(AuthMiddleware)
class UserController{

}
```

I cannot name it also `Middleware` because this is the name of the interface. 

Of course, one solution is to name the interface `MiddlwareInterface` but I feel it's tautology.
One other solution is to change the decorator to `AddMiddleware` but I also feel it's tautology and doesn't look that good in code.

What are some other beautiful ways to name interfaces that have decorators with the same name?
## [4][How would I properly write the example for React-Transition-Group's 'Transition' so that eslint stops yelling at me?](https://www.reddit.com/r/typescript/comments/gwgfzk/how_would_i_properly_write_the_example_for/)
- url: https://www.reddit.com/r/typescript/comments/gwgfzk/how_would_i_properly_write_the_example_for/
---
[ https://reactcommunity.org/react-transition-group/transition]( https://reactcommunity.org/react-transition-group/transition)

In particular this part:
    
    const defaultStyle = {
      transition: 'opacity ${duration}ms ease-in-out'
      opacity: 0,
    }
    
    const transitionStyles = {
      entering: { opacity: 1 },
      entered:  { opacity: 1 },
      exiting:  { opacity: 0 },
      exited:  { opacity: 0 },
    };
    
    const Fade = ({ in: inProp }) =&gt; (
      &lt;Transition in={inProp} timeout={duration}&gt;
        {state =&gt; (
          &lt;div style={{
            ...defaultStyle,
            ...transitionStyles[state]
          }}&gt;
            I'm a fade Transition!
          &lt;/div&gt;
        )}
      &lt;/Transition&gt;
    );                
    


I'm not really sure how to type the state in the arrow func or the transitionStyles[state] in the style
## [5][Is utility types a dev depedency or a depedency ?](https://www.reddit.com/r/typescript/comments/gw5ysy/is_utility_types_a_dev_depedency_or_a_depedency/)
- url: https://www.reddit.com/r/typescript/comments/gw5ysy/is_utility_types_a_dev_depedency_or_a_depedency/
---
`utility-types` is a typescript library of utility type functions. Since some of those functions end up in the `d.ts` file that `package.json` point with the field `types` I was wondering if it is a dependency and not a dev dependency.

Here is how I see stuff :

* dev depedency is any package of which the code does not end up in the  entry point files
* depedency is any package of which some part (or all)  of the code ends up in the entry point files
## [6][What are the actual reasons some devs don't like Typescript?](https://www.reddit.com/r/typescript/comments/gvrrsu/what_are_the_actual_reasons_some_devs_dont_like/)
- url: https://www.reddit.com/r/typescript/comments/gvrrsu/what_are_the_actual_reasons_some_devs_dont_like/
---
I usually hear "excess code writing/verbosity." Personally I think it covers for not wanting to learn another language. Or maybe the dev is in a field with fairly low complexity or routine work.

I just hooked up part of a CLI program across 5 different files with abstract base classes, interfaces enforcing consistency between a main controller class, and an array of instances of various types. Would have been impossible for it to work on the first go but TS tells you where there are issues every step of the way. I personally think one actually codes faster on the first pass too, when you consider the superior feedback loop.
## [7][What's the easiest way to deal with identical enums with different names?](https://www.reddit.com/r/typescript/comments/gvyfck/whats_the_easiest_way_to_deal_with_identical/)
- url: https://www.reddit.com/r/typescript/comments/gvyfck/whats_the_easiest_way_to_deal_with_identical/
---
I'm using  `graphql-codegen` and in the process it's automatically generating types for me to work with, however it's creating a lot of ENUMs with the same name for each different schema.  I don't have control of what these names are because the schema is generated automatically by `graphql-compose-mongoose`

Lets say I have two identical enums:

    enum A {
        This = 'THIS',
        That = 'THAT,
    }


    enum B {
        This = 'THIS',
        That = 'THAT,
    }

If I have a field on an interface that is designed to accept A, what would be the safest way for me to assert that it's okay to accept enum B because I know it's identical, but Typescript doesn't.

What I'm looking for ideally is something like:

`type AssertSameEnum&lt;T&gt; = T extends A ? A : never` 

But from what I understand you can't compare two enums with `extends` so that's where I'm stuck.

Thanks!
## [8][Watch mode for multiple child workspaces?](https://www.reddit.com/r/typescript/comments/gvv8wb/watch_mode_for_multiple_child_workspaces/)
- url: https://www.reddit.com/r/typescript/comments/gvv8wb/watch_mode_for_multiple_child_workspaces/
---
I have the following project structure:

```
├── foo                   (workspace @project/foo)
│   ├── package.json
│   └── src
│       └── index.ts
├── bar                   (workspace @project/bar)
│   ├── package.json
│   └── src
│       └── index.ts
├── package.json          (monorepo root)
└── tsconfig.json         (base tsconfig)
```

If I run `tsc --watch` in each child workspace, TypeScript watches the file changes correctly and rebuilds the package. However, If I run the same command at the monorepo root, it doesn't work - it doesn't throw an error but if I change anything in file, the package is not rebuilt.

Reading the [docs](https://www.typescriptlang.org/docs/handbook/project-references.html#build-mode-for-typescript), I understand the the `--build` flag does have an option to pass multiple child projects (or `tsconfig.json` files). Is there something equivalent for `--watch`?
## [9][Understanding "baseUrl" and "paths" in TypeScript with * glob](https://www.reddit.com/r/typescript/comments/gvq2mh/understanding_baseurl_and_paths_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gvq2mh/understanding_baseurl_and_paths_in_typescript/
---
I have a monorepo created with yarn workspaces and the following folder structure:

```
├── foo                   (workspace @project/foo)
│   ├── package.json
│   └── src
│       └── index.ts
├── bar                   (workspace @project/bar)
│   ├── package.json
│   └── src
│       └── index.ts
├── package.json          (monorepo root)
└── tsconfig.json         (base tsconfig)
```

And the following settings in `tsconfig.json`:

```ts
{
  "compilerOptions": {
    "baseUrl": ".",
    "module": "commonjs",
    "paths": {
      "@project/foo/*": "./packages/foo/src/*",
      "@project/bar/*": "./packages/bar/src/*"
    },
    ...
  }
}
```

In the `@project/bar` workspace, I want to import modules from `@project/foo`:


```ts
import foo from "@project/foo";
```

But I'm getting the following error:

&gt; Cannot find module '@project/foo' or its corresponding type declarations.ts(2307)

If I remove the `*` symbols from both the keys and the values of the "paths" object, the code compiles. Why is that? How can I keep the `*` glob pattern and make non-relative imports to my local modules?
## [10][How to see TypeScript at my company?](https://www.reddit.com/r/typescript/comments/gvuwr6/how_to_see_typescript_at_my_company/)
- url: https://www.reddit.com/r/typescript/comments/gvuwr6/how_to_see_typescript_at_my_company/
---
I love working with React + TypeScript, but for a new project I have to make a case for this and some of the devs are resistant. 

Are there any good articles or resources to do this? I'm assuming most complex JavaScript apps started today would use TypeScript rather than ES6 or Flow, but is there any data on this?
## [11][Define a type of Record where the key should be of type T](https://www.reddit.com/r/typescript/comments/gvlse5/define_a_type_of_record_where_the_key_should_be/)
- url: https://www.reddit.com/r/typescript/comments/gvlse5/define_a_type_of_record_where_the_key_should_be/
---
I have a `create` method inside a users-service like so:

```
async create(fields: Record&lt;string, any&gt;) {}
```

but the `fields` is always an object of keys where the keys are a partial of `User` keys, so how do I achieve this typing? Something like: `Record&lt;K extends keyof User, V&gt;` ?
