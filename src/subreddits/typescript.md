# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][GitHub - twitchtv/twirp: A simple RPC framework with protobuf service definitions](https://www.reddit.com/r/typescript/comments/f6rqis/github_twitchtvtwirp_a_simple_rpc_framework_with/)
- url: https://github.com/twitchtv/twirp
---

## [3][Building an application with Vue and TypeScript. Best practices, thoughts and recommendations.](https://www.reddit.com/r/typescript/comments/f6ghfx/building_an_application_with_vue_and_typescript/)
- url: https://stefan-bauer.online/building-an-application-with-vue-and-type-script-best-practices-thoughts-and-recommendations/
---

## [4][I programmed in TypeScript like in Haskell (Lazy Evaluation)](https://www.reddit.com/r/typescript/comments/f6t3hj/i_programmed_in_typescript_like_in_haskell_lazy/)
- url: https://www.youtube.com/watch?v=E5yAoMaVCp0&amp;feature=share
---

## [5][AdminBro v2 beta is out!!!](https://www.reddit.com/r/typescript/comments/f6szkd/adminbro_v2_beta_is_out/)
- url: https://www.reddit.com/r/typescript/comments/f6szkd/adminbro_v2_beta_is_out/
---
AdminBro is an automated admin interface for Node.js apps based on React, written in TypeScript = Django admin for node on steroids !!!

Now, we almost finish the entire scope for v2.0 and we released it along with the demo version: 

Demo page can be found here:

[https://admin-bro-example-app-staging.herokuapp.com/admin](https://admin-bro-example-app-staging.herokuapp.com/admin)

Github repo is here: [https://github.com/SoftwareBrothers/admin-bro](https://github.com/SoftwareBrothers/admin-bro)

and you can install it using: 

```
yarn add admin-bro@2.0.0-beta.26
```

In this version: 

* there is an entire Design System written specially for AdminBro
* i18n support
* hooks support
* the entire UI change to Drawer'like exploring resources

Now we are working on the documentation which (WiP) can be fond here:

[https://softwarebrothers.github.io/admin-bro-dev/v2/](https://softwarebrothers.github.io/admin-bro-dev/v2/)

Let me know what do you think - and star the repo if you like the idea :)
## [6][Typesafe replacement for node's EventEmitter embracing functional programming](https://www.reddit.com/r/typescript/comments/f6d6r7/typesafe_replacement_for_nodes_eventemitter/)
- url: https://github.com/garronej/ts-evt
---

## [7][How do I disable ESLint typescript for JS files?](https://www.reddit.com/r/typescript/comments/f6p73o/how_do_i_disable_eslint_typescript_for_js_files/)
- url: https://www.reddit.com/r/typescript/comments/f6p73o/how_do_i_disable_eslint_typescript_for_js_files/
---
Hello all,

I am starting with Typescript and I would like to disable ESLint for JS files, my .eslintrc.json contains the following:

```
    "parser":  "@typescript-eslint/parser",
    "extends": [
        "airbnb",
        "prettier",
        "prettier/react",
        "plugin:jsx-a11y/recommended",
        "plugin:jest/recommended"
        // "plugin:@typescript-eslint/recommended"
    ],
    "overrides": [
      {
        "files": ["*.js", "*.jsx"],
        "rules": {
            "@typescript-eslint/...": "off"
        }
      }
    ],

```

But I continue getting eslint messages due to "plugin:@typescript-eslint/recommended". Any help with this please?

Thank you in advance and regards.
## [8][Trying to set type for single, dynamic property of indexed type](https://www.reddit.com/r/typescript/comments/f6gudt/trying_to_set_type_for_single_dynamic_property_of/)
- url: https://www.reddit.com/r/typescript/comments/f6gudt/trying_to_set_type_for_single_dynamic_property_of/
---
I have a variety of records that look like this:
```typescript
{
   id: string;
   fields: number[];
   field1: Record&lt;string, unknown&gt;;
}

{
   other_id: string;
   more_fields: number[];
   field2: number
}
```

I'm trying to set the typing on a function where you pass the property name of the `number[]` field to the function so it knows where to grab the numbers from.

Something like:
```typescript
function sum_fields&lt;T, K extends keyof T &amp; number[]&gt;(obj: T, field: K): number {
  let sum = 0;
  obj[field].forEach((i: number) =&gt; {
      sum += i;
  });
  return sum;
}
```

How can I tell typescript that the property name I'm passing in has the type number[], and/or can I get typescript complain if the type of field is *not* number[]?
## [9][Module federation and code sharing between bundles. Huge changes coming to frontend with webpack@5](https://www.reddit.com/r/typescript/comments/f68284/module_federation_and_code_sharing_between/)
- url: https://github.com/webpack/webpack/issues/10352
---

## [10][TypeScript Groovy Console Logger](https://www.reddit.com/r/typescript/comments/f65ist/typescript_groovy_console_logger/)
- url: https://www.reddit.com/r/typescript/comments/f65ist/typescript_groovy_console_logger/
---
I wanted to create a class that allowed me to pass lots of information for debugging purposes. Based on the Type of Debug message, different sections of the resulting Console message will have different colors.

I would like to get feedback from you guys.
Comments, questions, suggestions all welcome

[Github Groovy Logger!](https://github.com/RecursiveCTE/Visual_Studio_Code_Samples)

I'm new to github as well so please let me know if that link doesn't work.

Thanks!
## [11][Implementing an opaque type in typescript](https://www.reddit.com/r/typescript/comments/f5wny3/implementing_an_opaque_type_in_typescript/)
- url: https://evertpot.com/opaque-ts-types/
---

