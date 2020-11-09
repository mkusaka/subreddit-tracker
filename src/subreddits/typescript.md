# typescript
## [1][Who's hiring Typescript developers - November](https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/)
- url: https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/
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
## [2][Stator: A full-stack boilerplate – releases, deployments, enforced conventions](https://www.reddit.com/r/typescript/comments/jqpizp/stator_a_fullstack_boilerplate_releases/)
- url: https://www.reddit.com/r/typescript/comments/jqpizp/stator_a_fullstack_boilerplate_releases/
---
Have you ever started a new project by yourself?

If so, you probably know that it is tedious to set up all the necessary tools. 

Just like you, the part I enjoy the most is coding, not boilerplate.

&amp;#x200B;

Say hi to [stator](https://github.com/chocolat-chaud-io/stator), a full-stack TypeScript template that enforces conventions, handles releases, deployments and many more features!
## [3][Help with typing a deep merge function.](https://www.reddit.com/r/typescript/comments/jqqtg5/help_with_typing_a_deep_merge_function/)
- url: https://www.reddit.com/r/typescript/comments/jqqtg5/help_with_typing_a_deep_merge_function/
---
Hey all, I'm wondering if there are any smart people out there looking a little bit of a challenge.
On the surface this didn't seem like it would be very hard when I started but it's been a lot more difficult than I expected.

But anyway, essentially I'm trying to type 2 functions. One that will deeply merge any 2 given objects into a new resulting object. And the other that will deeply merge an array of 2+ objects into a new resulting object.

I feel I'm close to getting the typings all correct but I'm having trouble with the second of the 2 functions. I'd really appreciate any help with solving this issue.
I'd also really love if anyone was able to peer review my current PR.

https://github.com/TehShrike/deepmerge/pull/211

Thanks for reading :)
## [4][Terraform with TypeScript](https://www.reddit.com/r/typescript/comments/jqxmoq/terraform_with_typescript/)
- url: https://medium.com/francisvitullo/terraform-with-typescript-7643defb4eb1?source=friends_link&amp;sk=975dc30cd7d48d989f2d7f0c16882c2e
---

## [5][How to use Contact Picker API in typescript?](https://www.reddit.com/r/typescript/comments/jqu6yo/how_to_use_contact_picker_api_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/jqu6yo/how_to_use_contact_picker_api_in_typescript/
---
Hi.I need to use [Contact Picker API](https://wicg.github.io/contact-api/spec/) in my angular project. but it doesn't work because contacts and ContactsManager is not declared in Navigator in typescript dom library. and the error say "property contatcs does not exist on Navigator"

here's some working examples in js:[https://whatwebcando.today/contacts.html](https://whatwebcando.today/contacts.html)[https://contact-picker.glitch.me](https://contact-picker.glitch.me/)

here's my test angular app:  [https://codesandbox.io/s/angular-contacts-85x29-85x29](https://codesandbox.io/s/angular-contacts-85x29-85x29)

I copied one of the examples and it doesn't work. I tried to add interface for contactManager and contact and merge it with Navigator in contacts.d.ts and added it in types in tsconfig, but I think it's not correct. I appreciate if you check my code or give me some hint on what to do to make it work.

thanks in advance.
## [6][I have released a web sdk to reduce video streaming costs by 90%](https://www.reddit.com/r/typescript/comments/jqazdm/i_have_released_a_web_sdk_to_reduce_video/)
- url: https://www.reddit.com/r/typescript/comments/jqazdm/i_have_released_a_web_sdk_to_reduce_video/
---
 Hi everyone,

I am working on a javascript sdk which can reduce video streaming costs of CDN by up-to 90% using a hybrid decentralized load sharing technology. I have opened it up for beta-access for developers to try it out. Looking for feedback w.r.t the technology and any features you would want to have.

A web demo is available here [https://api.peervadoo.com/test](https://api.peervadoo.com/test) . Click on **Add new peer** to see the tech in action,

**Javascript sdk link** :- [https://github.com/vadootvpeer/sdk-javascript](https://github.com/vadootvpeer/sdk-javascript)
## [7][Is there any way to make `tsc` "inline" imported modules?](https://www.reddit.com/r/typescript/comments/jqkrn3/is_there_any_way_to_make_tsc_inline_imported/)
- url: https://www.reddit.com/r/typescript/comments/jqkrn3/is_there_any_way_to_make_tsc_inline_imported/
---
I'm writing a firefox extension using TypeScript. So far, the only tool in my build chain is `tsc` and ideally I would like it to stay that way.

I basically have three source files:

    src/background.ts
    src/context.ts
    src/shared.ts

`background` and `context` both use a few small methods from `shared`. However, the resulting JS code uses `import` statements (which do not really work in an extension environent). Ideally, I would like the import statement to be replaced by expanding the original code from `shared`, such that the output consists of only `background.js` and `context.js` (each with duplicated copies of the shared code).

I'm pretty new to typescript and have found the documentation unhelpful
## [8][how to remove debug TypeScript](https://www.reddit.com/r/typescript/comments/jqm9r1/how_to_remove_debug_typescript/)
- url: https://www.reddit.com/r/typescript/comments/jqm9r1/how_to_remove_debug_typescript/
---
I'm currently remote debugging in webstorm with node and everything works great. but now I'd like to remote  debug TypeScript and wondering how is that set up with the map files and everything?

thanks
## [9][Any good ways to make a shared tsconfig.json?](https://www.reddit.com/r/typescript/comments/jqe69e/any_good_ways_to_make_a_shared_tsconfigjson/)
- url: https://www.reddit.com/r/typescript/comments/jqe69e/any_good_ways_to_make_a_shared_tsconfigjson/
---
So I am writing a general library containing JS/TS configuration utilities that I use across all my projects. Webpack, Babel, ESLint, etc. The goal is to be able to have something that I can literally just drop in and use. So I'm looking for help in a few areas:

1. Is there the equivalent of tsconfig.js? JavaScript configuration files, IMO, are much easier to use for shared configuration. I can just alternate between `__dirname` for references to code in the configuration library and `process.cwd()` for references to code in the consuming project. I have seen no references to this capability, so I assume no. If there is a way to do this, all subsequent questions are moot. If the answer is no, then please provide guidance for the subsequent questions for the JSON configuration.
2. Is there any way to have configurable settings? For example, whether to compile the code to have esmodules or commonjs modules is a setting I want configurable.
3. How do I handle relative paths? The `include` array, `compilerOptions.outDir`, `exclude` array, etc. I would want to be able to have the tsconfig.json in my configuration library while still pointing to the locations in the project that uses it.

That's it for now. Help is appreciated. Thanks.
## [10][Filtering array of Record&lt;&gt; and Objects](https://www.reddit.com/r/typescript/comments/jqd9c8/filtering_array_of_record_and_objects/)
- url: https://www.reddit.com/r/typescript/comments/jqd9c8/filtering_array_of_record_and_objects/
---
I'm new to TypeScript, and since for the past few years I only programmed in JS, I'm not really used to using types.

I want to filter an array of objects like this:`(Record&lt;string, unknown&gt; | iItem)[]`

I want it to only contain `iItem`. I know that iItem will contain properties like id and type, inherited from `BaseItem` interface. What I've tried to do is to create a filtering function

    isItem( v: Record&lt;string, unknown&gt; | iItem ):v is iItem {
 return v instanceof BaseItem
 }

And then pass it to `.filter` method, however it doesn't seem to recognize that the output type will be iItem. Is there an easier way of filtering it? Am I doing something wrong?
## [11][Check if string is AWS ARN](https://www.reddit.com/r/typescript/comments/jqj59y/check_if_string_is_aws_arn/)
- url: https://www.reddit.com/r/typescript/comments/jqj59y/check_if_string_is_aws_arn/
---
Hey guys, how can I check if a string is an AWS ARN in Typescript? I'm having difficulty implementing such a functionality.
