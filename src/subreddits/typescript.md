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
## [2][Mutex-Server, a new npm module I've developed (+ need your advise)](https://www.reddit.com/r/typescript/comments/gqso87/mutexserver_a_new_npm_module_ive_developed_need/)
- url: https://www.npmjs.com/package/mutex-server
---

## [3][Earl - Modern chai replacement tailored for TypeScript](https://www.reddit.com/r/typescript/comments/gqc5tq/earl_modern_chai_replacement_tailored_for/)
- url: https://github.com/krzkaczor/earl
---

## [4][convert d.ts tree as a single d.ts file .](https://www.reddit.com/r/typescript/comments/gquyk2/convert_dts_tree_as_a_single_dts_file/)
- url: https://www.reddit.com/r/typescript/comments/gquyk2/convert_dts_tree_as_a_single_dts_file/
---
Lets suppose that I have a `d.ts` file that it imports types from other `d.ts` files which import types from other `d.ts` and it goes on . 

Is there any tool that will allow me to make a single `d.ts` file out of that tree that will have no imported types from other `d.ts` files because it has inside of it all its dependencies ?
## [5][End-to-end Type Safety in Clean Architecture: a possible solution with TypeScript, GraphQL, MongoDB, React.](https://www.reddit.com/r/typescript/comments/gqx14r/endtoend_type_safety_in_clean_architecture_a/)
- url: https://charlesagile.com/end-to-end-type-safety
---

## [6][How can I use path aliases in a library in a way that aliases are resolved when used by library consumers?](https://www.reddit.com/r/typescript/comments/gqupni/how_can_i_use_path_aliases_in_a_library_in_a_way/)
- url: https://www.reddit.com/r/typescript/comments/gqupni/how_can_i_use_path_aliases_in_a_library_in_a_way/
---
Hi. I'm writing a library called `cool-lib` in this library I am using TypeScript path aliases to avoid the relative import path hell.

I am using an alias like so: `"@": "./src"`.

When I try to consume `cool-lib` in an actual project I get import path errors. I'm aware this is expected because the path aliases aren't resolved in the emitted JS.

In order to try get around this I added the `module-alias` npm package to `cool-lib` but the same import errors occur in the project consuming `cool-lib`. I'm not sure if I've misconfigured this or it's not meant to be used for TypeScript libraries. Edit: I just noticed it's failing in declaration files so maybe for the source JS files it's actually working? If this is the case how do I resolve paths in generated declaration files?

Another option I thought of was adding a bundler to the library and releasing the library as a bundle? Not really keen on adding something like Webpack just for this though.

What can I do in `cool-lib` so I can use path aliases but not have the library break when consumed by actual projects due to import paths?
## [7][Do you have a method for declaring or creating fixed length arrays/tuples that are too long to write?](https://www.reddit.com/r/typescript/comments/gqe334/do_you_have_a_method_for_declaring_or_creating/)
- url: https://www.reddit.com/r/typescript/comments/gqe334/do_you_have_a_method_for_declaring_or_creating/
---
My use case is specifying arguments for an API that accepts only 100 "ids" per request.

I have looked through [this](https://stackoverflow.com/questions/41139763/how-to-declare-a-fixed-length-array-in-typescript) and [that](https://stackoverflow.com/questions/52489261/typescript-can-i-define-an-n-length-tuple-type) stackoverflow thread and [this github issue](https://github.com/Microsoft/TypeScript/issues/26223).
## [8][exporting imported interfaces](https://www.reddit.com/r/typescript/comments/gqb0kg/exporting_imported_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/gqb0kg/exporting_imported_interfaces/
---
I am creating a package where I want to export the interfaces from the main index.ts file

foo.ts

`export interface Foo {   foo: string }`

interfaces.ts

`export { Foo } from './foo';`

and then the package export

`export * as interfaces from './interfaces';`

if I try to access the interfaces

`import { interfaces } from './index';`

interface has nothing in it

I can, however, import Foo from the interfaces file, which is what I want to avoid  
`import { Foo } from './interfaces';`  
`const foo: Foo = {`  
  `foo: '',`  
`}`
## [9][Are you willing to get a mentor?](https://www.reddit.com/r/typescript/comments/gq9nl8/are_you_willing_to_get_a_mentor/)
- url: https://www.reddit.com/r/typescript/comments/gq9nl8/are_you_willing_to_get_a_mentor/
---
Hey guys, I want to take programming seriously and I am looking for a mentor and I’m not sure who I should pick. I would like to know your opinion on this.  
As a software engineer who knows the basics, are you willing to pay a mentor to help you?   
If yes:

* How much are you willing to pay if he guarantees you the best results?
* What problems do you want the mentor to help you with?
## [10][Need help importing a .gql file into my vue class component](https://www.reddit.com/r/typescript/comments/gpwi83/need_help_importing_a_gql_file_into_my_vue_class/)
- url: https://www.reddit.com/r/typescript/comments/gpwi83/need_help_importing_a_gql_file_into_my_vue_class/
---
Hello everyone. I started with typescript last week and wanted to redo one of my projects with typescript.

I'm trying to import a .gql file into my vue class component file but it gives me the following error.

" **Cannot find module** "

does anyone maybe know how to import .gql files?

I tried googling but couldn't get anything that is related to .gql files.
## [11][Object Typing Question](https://www.reddit.com/r/typescript/comments/gpvl62/object_typing_question/)
- url: https://www.reddit.com/r/typescript/comments/gpvl62/object_typing_question/
---
Hey all,

Typescript beginner here. This is probably crazy basic but I'm having a hard time putting this into words I can google. Can anyone help me wrap my brain around what's going on here?   Mostly just curious about line 4, where I'm setting the type for the errors object:

`const isRequiredError = "Missing required";`  
`const dateReg = /0?[1-9]|[12][0-9]|3[01]/-[/-]\d{4}$/;`  
`const malFormattedDateError = "Date field should be formatted DD-MM-YYYY or DD/MM/YYYY";`  
`const errors: { [field: string]: string } = {};`

`if (!value.description) errors.description = isRequiredError;`  
`if (!value.date) errors.date = isRequiredError;`  
`if (!value.specialist) errors.specialist = isRequiredError;`  
`if (!dateReg.test(value.date)) errors.date = malFormattedDateError;`  
`if (!value.discharge.date) errors.discharge.date = isRequiredError;`

`return errors;`

Everything else I've typed has been stuff like:

`interface Foo = {  bar: string;}`

Whereas this is more like:

`interface Foo = {  [bar: string]: string;}`

What are the brackets doing here?

Also this seems to accommodate doing something like

`errors.field = "You effed up";`

How would I type this to accommodate something like this as well?

`errors.field.subfield = "You effed up here too"`

Thanks for the help!

&amp;#x200B;

Edit: Formatting
