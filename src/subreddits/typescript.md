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
## [2][Building a Synthesizer in TypeScript](https://www.reddit.com/r/typescript/comments/gq75go/building_a_synthesizer_in_typescript/)
- url: https://medium.com//building-a-synthesizer-in-typescript-5a85ea17e2f2?source=friends_link&amp;sk=e443409e6d34185e8496f0aeee22eb5e
---

## [3][Red - Type-safe, composable, boilerplateless reducers ; https://github.com/betafcc/red](https://www.reddit.com/r/typescript/comments/gq1t2v/red_typesafe_composable_boilerplateless_reducers/)
- url: https://i.redd.it/fb93h7x0et051.gif
---

## [4][Humble Book Bundle: Definitive Guides to All Things Programming by O'Reilly (pay what you want and help charity)](https://www.reddit.com/r/typescript/comments/gpwdzu/humble_book_bundle_definitive_guides_to_all/)
- url: https://www.humblebundle.com/books/definitive-guides-to-all-things-programming-oreilly-books?partner=indiekings
---

## [5][exporting imported interfaces](https://www.reddit.com/r/typescript/comments/gqb0kg/exporting_imported_interfaces/)
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
## [6][Are you willing to get a mentor?](https://www.reddit.com/r/typescript/comments/gq9nl8/are_you_willing_to_get_a_mentor/)
- url: https://www.reddit.com/r/typescript/comments/gq9nl8/are_you_willing_to_get_a_mentor/
---
Hey guys, I want to take programming seriously and I am looking for a mentor and I’m not sure who I should pick. I would like to know your opinion on this.  
As a software engineer who knows the basics, are you willing to pay a mentor to help you?   
If yes:

* How much are you willing to pay if he guarantees you the best results?
* What problems do you want the mentor to help you with?
## [7][Need help importing a .gql file into my vue class component](https://www.reddit.com/r/typescript/comments/gpwi83/need_help_importing_a_gql_file_into_my_vue_class/)
- url: https://www.reddit.com/r/typescript/comments/gpwi83/need_help_importing_a_gql_file_into_my_vue_class/
---
Hello everyone. I started with typescript last week and wanted to redo one of my projects with typescript.

I'm trying to import a .gql file into my vue class component file but it gives me the following error.

" **Cannot find module** "

does anyone maybe know how to import .gql files?

I tried googling but couldn't get anything that is related to .gql files.
## [8][Object Typing Question](https://www.reddit.com/r/typescript/comments/gpvl62/object_typing_question/)
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
## [9][Verify cookie with JWT payload inside.](https://www.reddit.com/r/typescript/comments/gppio8/verify_cookie_with_jwt_payload_inside/)
- url: https://www.reddit.com/r/typescript/comments/gppio8/verify_cookie_with_jwt_payload_inside/
---
https://preview.redd.it/nm5u38x6qp051.png?width=596&amp;format=png&amp;auto=webp&amp;s=b1c75a2fd39e0e077840f4ee40188c3c4e1607ed

Hi everyone, I'm learning TypeScript and I'm stuck in that problem.

The return of function verify is the code commented above the interface Decoded

Thank you for your help!

&amp;#x200B;
## [10][#1 thing of coding is 'data structures' - everything learned about the first steps in learning typescript (and any other coding languages)](https://www.reddit.com/r/typescript/comments/gpz597/1_thing_of_coding_is_data_structures_everything/)
- url: https://www.reddit.com/r/typescript/comments/gpz597/1_thing_of_coding_is_data_structures_everything/
---
# #1 learning this thing call 'data structure' is the most important thing of coding

so most important thing is to find a good site that shows 'data structure' so everyone can learn it

&amp;#x200B;

"Programming is not really defined in the way you describe it. Namely, you describe functions as *actions* and arguments as *adjustments*. Sure, that’s a decent analogy. But you need to then comprehend how to actually write those functions (or actions as you describe them) to manipulate a *data structure*.

This is the key piece you are missing. You have to have data stored somewhere, and it has to have the correct *structure* for the data you want it to hold. And those actions need to change that data in a well-defined way so that the state of program can progress. You are calling functions on premade objects which is all well and good, but you need to understand how to make data structures, operate on them, and interpret them so that you can begin to write actual programs.

This is going to be very condensed and simplified, but in essence, you:

1. Start your program
2. Make your data structures to hold any data
3. Initialize them with the starting data they need (or they might not need any initialization at all).
4. Enter into the part of your program that does something with that data (actions and adjustments).
5. Exit out of the part that did something with that data
6. And either exit the program, or read the data and figure out what to do next to it
7. if you need to do more things to the data after reading it, do it

Here’s a good medium article [about different types of data structures and how to use them.](https://medium.com/swlh/introduction-to-data-structures-9134b7d064a6) But it might be a bit high level for you. Best is to take an intro course or typescript or python and learn them gradually. But don’t give up, people tend to hit a wall when “hard” stuff like this pops up in programming. But honestly, it isn’t hard once you get the hang of how it works. Best of luck."

&amp;#x200B;

&amp;#x200B;

**here's all the sites i used to try to learn coding**

[https://www.reddit.com/r/learnpython/comments/gp0zsi/heres\_all\_the\_sites\_i\_used\_to\_try\_to\_learn\_coding/](https://www.reddit.com/r/learnpython/comments/gp0zsi/heres_all_the_sites_i_used_to_try_to_learn_coding/)

**here's everything i know about coding so far:**

[https://www.reddit.com/r/learnpython/comments/goy5xa/where\_to\_continue\_learning\_coding/](https://www.reddit.com/r/learnpython/comments/goy5xa/where_to_continue_learning_coding/)

&amp;#x200B;

# #2 random links to learn ts / js - no idea which is ACTUALLY for beginniners, ill test all the js ones

First of all, TypeScript is a typed superset of JavaScript. So, you need to understand JavaScript/ECMAScript first.Few resource to learn:

* [A re-introduction to JavaScript (JS tutorial)](https://developer.mozilla.org/en-US/docs/Web/JavaScript/A_re-introduction_to_JavaScript)
* [The Modern JavaScript Tutorial](https://javascript.info/)
* [33 Concepts Every JavaScript Developer Should Know](https://github.com/leonardomso/33-js-concepts/)
* [Modern JavaScript Cheatsheet](https://github.com/mbeaudru/modern-js-cheatsheet)
* [Clean code JavaScript](https://github.com/ryanmcdermott/clean-code-javascript)

And after that, you can go with TS:

* [Official TypeScript docs](https://www.typescriptlang.org/docs/home)
* [TypeScript Cheatsheet](https://github.com/rmolinamir/typescript-cheatsheet)

&amp;#x200B;

&amp;#x200B;

**original question**

as this is liekly the most common question asked over time

im guessing some of the very knowledgable ppl here had sometimes during the past created a script in which other ppl would just copy / paste to new ppl that ask this question (since you're coders and would want to automate things so they can be made efficient &amp; effective)

&amp;#x200B;

# #3 if coding was designed well, then it would be quick + easy to learn, just lilke hopscotch or walking

'have quite a way to go ' - so is it cos coding is not at the state of progress where it could be easiely and quickly learned? kinda like how some things can be quickly and easiely learned, like hopscotch or walking (for example many or even every single thing is that designed well is easy to learn, but things that are not, are not easy to learn but hard)

i also asked if coding was not at the state of progress where it would be easy and quick  to learn (which it can if it was designed well which you didnt consider). you didnt consider the things taht are easy and quick to learn (especially when there are good sites / resources for new ppl)

&amp;#x200B;

**anything about 'enterpries' doesnt matter**

'why such and such language is an "enterprise" language will *fly* over your head at the moment, don't even read them, forget about them.'

it seems like typescript is mainly for 'enterprise use' and for 'angular'  [https://www.quora.com/Should-I-learn-JavaScript-before-learning-TypeScript](https://www.quora.com/Should-I-learn-JavaScript-before-learning-TypeScript) so maybe should learn something simpler but it doesnt seem like there are any simpler options out there...

&amp;#x200B;

# #4 good learning sites are hard to find

there are no beginning resources on r/reddit, for example the helpful thing i used to learn 'codecombat', its not even on there etc
## [11][Exporting a Jest mock of a module with type casting](https://www.reddit.com/r/typescript/comments/gpgg8d/exporting_a_jest_mock_of_a_module_with_type/)
- url: https://www.reddit.com/r/typescript/comments/gpgg8d/exporting_a_jest_mock_of_a_module_with_type/
---
Hey everyone. I have a quick question which kind of has to do with typescript, but I think may be more of a Jest question. If this is in the wrong place, I apologize.

I'm learning typescript while building a [nuxt.js](https://nuxtjs.org/) app with jest for unit testing. I am coding my API service layer right now, and have the following code for mocking the [nuxt.js axios module](https://axios.nuxtjs.org/), which as a different type than a standard axios instance:

```TypeScript
import axios from 'axios'
import { NuxtAxiosInstance } from '@nuxtjs/axios'

jest.mock('axios')

// cast mock to correct type
const mockNuxtAxios = axios as jest.Mocked&lt;NuxtAxiosInstance&gt;

// return mock data for test
mockNuxtAxios.$get = jest.fn().mockReturnValue(mockData)
```

This is quite a lot of code, and I am using this mock in a few different test files. I would like to avoid duplicating this code whenever I need to mock nuxt axios. What would be ideal would be to be able to import my `mockNuxtAxios` mock into multiple test files with something as simple as:

```TypeScript
import { mockNuxtAxios } from '~/test/utils'
```

I imagine that this is a fairly common thing that someone would want to do, but I cannot figure out how to.

Any suggestions?
