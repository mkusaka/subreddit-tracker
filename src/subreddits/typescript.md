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
## [2][Possible to use TypeScript to check JS only? Trying to figure out incremental conversion path for huge project](https://www.reddit.com/r/typescript/comments/grj31m/possible_to_use_typescript_to_check_js_only/)
- url: https://www.reddit.com/r/typescript/comments/grj31m/possible_to_use_typescript_to_check_js_only/
---
tl;dr trying to just use TypeScript as analysis tool is that possible?

Hey we have a large project some 100,000 lines of code with open contribution within the company. The code is abused and contributing on by hundreds of devs and people don't care about standards, they are just trying to get their code in to meet their deadline.

I am one of the core maintainers and our main defense is very strict testing requirements and static analysis on our CI which will block pull requests.

In the past when we've introduced new standards, we did it incrementally by making sure the code base did not get worse (e.g. we introduced coverage requirements by only checking that only new code was covered).

We are trying to find a similar path forward for TypeScript but have so far failed. 

However, I keep hearing the benefits of using TypeScript even without converting everything and fixing errors. I noticed the "checkJS" flag and found this pretty useful and was curious if we could run tsc with no output and just use it to check code and find certain errors TypeScript is good at like unresolved variables?
## [3][TypeScript's Language Features Documentation?](https://www.reddit.com/r/typescript/comments/gr8jt4/typescripts_language_features_documentation/)
- url: https://www.reddit.com/r/typescript/comments/gr8jt4/typescripts_language_features_documentation/
---
I'm trying to get on-board with TypeScript coming from a JavaScript background. All of the documentation only talks about the Types and how to use them. But I'm finding hidden secrets like TypeScript supporting [Optional Chaining](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-7.html) or automatically setting instance variables based on private class constructor arguments.

```
class Test {
    constructor(private name) {}

    printName() {
        console.log(this.name); // &lt;- Somehow this is magically defined by TypeScript?
    }
}
```

I totally get the Type features that TypeScript adds. That's easy to follow and well documented. But there seems to be actual language features beyond just types that it adds to and I'd like to know what all of those are. Where can I find documentation for all of the changes to the JavaScript language TypeScript adds?
## [4][Exhaustive check in TypeScript](https://www.reddit.com/r/typescript/comments/grjnrm/exhaustive_check_in_typescript/)
- url: https://www.instagram.com/p/CAsUYVEghoO/
---

## [5][I recently spoke w/a recruiter about a role requiring substantial TS experience. TS is a superset of JS that applies strict and custom types to JS code. If two developers both have very strong JS but one has two years and the other six months of TS, is there really THAT much of a difference?](https://www.reddit.com/r/typescript/comments/gr1ybe/i_recently_spoke_wa_recruiter_about_a_role/)
- url: https://www.reddit.com/r/typescript/comments/gr1ybe/i_recently_spoke_wa_recruiter_about_a_role/
---
I know that TypeScript has its idioms, its ins and outs, etc., but what I'm questioning is whether there's that substantial a difference in two strong JavaScript engineers who have different amounts of experience specifically with TypeScript. TypeScript itself doesn't seem to be a "years of experience" sort of technology; it's very easy to pick up and, with some practice and study, to implement properly. It's more like Lodash (eh, maybe that's a bit crude a comparison, but you take my meaning--it's just a tool that serves a specific purpose in JavaScript, but not a fundamentally different paradigm. It's not like going from JavaScript to, say, Java.)

Edit: A bit more information.

I picked up TypeScript about three months ago and have been using it since. I like type safety, but TypeScript itself doesn't seem to be so radical a departure from JavaScript that I'd eliminate a strong engineer from the running if they have only a few months of experience with the technology. It's really not that tough to pick up and use properly.

Let me know if I'm off-base here.
## [6][End-to-end Type Safety in Clean Architecture: a possible solution with TypeScript, GraphQL, MongoDB, React.](https://www.reddit.com/r/typescript/comments/gqx14r/endtoend_type_safety_in_clean_architecture_a/)
- url: https://charlesagile.com/end-to-end-type-safety
---

## [7][Why is it so hard to iterate over objects in TypeScript?](https://www.reddit.com/r/typescript/comments/gqxh2h/why_is_it_so_hard_to_iterate_over_objects_in/)
- url: https://effectivetypescript.com/2020/05/26/iterate-objects/
---

## [8][Is there a way to have TypeScript recompile my project when I save a file without the use of an IDE? (I'm using Vim.)](https://www.reddit.com/r/typescript/comments/gr1o6h/is_there_a_way_to_have_typescript_recompile_my/)
- url: https://www.reddit.com/r/typescript/comments/gr1o6h/is_there_a_way_to_have_typescript_recompile_my/
---

## [9][Mutex-Server, a new npm module I've developed (+ need your advise)](https://www.reddit.com/r/typescript/comments/gqso87/mutexserver_a_new_npm_module_ive_developed_need/)
- url: https://www.npmjs.com/package/mutex-server
---

## [10][convert d.ts tree as a single d.ts file .](https://www.reddit.com/r/typescript/comments/gquyk2/convert_dts_tree_as_a_single_dts_file/)
- url: https://www.reddit.com/r/typescript/comments/gquyk2/convert_dts_tree_as_a_single_dts_file/
---
Lets suppose that I have a `d.ts` file that it imports types from other `d.ts` files which import types from other `d.ts` and it goes on .

Is there any tool that will allow me to make a single `d.ts` file out of that tree that will have no imported types from other `d.ts` files because it has inside of it all its dependencies ?

Edit : [Solution](https://www.reddit.com/r/typescript/comments/go5iy7/make_typescript_understand_that_the_esmodule_i/frz19a5?utm_source=share&amp;utm_medium=web2x) .
## [11][Build/output issues in TypeScript](https://www.reddit.com/r/typescript/comments/gr0l5q/buildoutput_issues_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gr0l5q/buildoutput_issues_in_typescript/
---
When I run \`tsc\` from the directory in which my \`.tsconfig\` file is located, no errors appear yet no \`build\` directory is generated. Any idea what the problem might be? Here's my \`.tsconfig\`:

&amp;#x200B;

`{`

`"compilerOptions": {`

`"target": "es5",`

`"outDir": "./build",`

`"allowJs": true,`

`"checkJs": true,`

`"resolveJsonModule": true,`

`"listEmittedFiles": true,`

`"listFiles": true,`

`"esModuleInterop": true,`

`"skipLibCheck": true,`

`"allowSyntheticDefaultImports": true,`

`"strict": true,`

`"alwaysStrict": true,`

`"forceConsistentCasingInFileNames": true,`

`"module": "esnext",`

`"moduleResolution": "node",`

`"resolveJsonModule": true,`

`"noEmit": true`

`},`

`"include": ["./**/*"],`

`"exclude": ["./**/node_modules/*"],`

`"typeroots": ["src/types/index.ts", "node_modules/@types"]`

`}`
