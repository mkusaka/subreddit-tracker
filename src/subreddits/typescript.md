# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
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
## [2][Announcing TypeScript 3.9 RC](https://www.reddit.com/r/typescript/comments/gaecct/announcing_typescript_39_rc/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-9-rc/
---

## [3][A collection of challenging TypeScript exercises](https://www.reddit.com/r/typescript/comments/gabb46/a_collection_of_challenging_typescript_exercises/)
- url: https://github.com/mdevils/typescript-exercises
---

## [4][Restore mock of imported class function in ts-jest while testing](https://www.reddit.com/r/typescript/comments/gav2w0/restore_mock_of_imported_class_function_in_tsjest/)
- url: https://www.reddit.com/r/typescript/comments/gav2w0/restore_mock_of_imported_class_function_in_tsjest/
---
**Issue**

I am mocking a function (only one) which is imported from another class which is being called inside the main class. I am able to mock the function and return the values which I provide. But I am not able to restore the mocked function back to normal for the subsequent tests.

Any help would be appreciated!

Framework Using: jest + ts-jest

**Code**

\~main.ts

\`\`\`

import {SubClass} from './subclass.ts' 

export class MainClass { 

let sub: SubClass = new SubClass() 

public async Car(){ 

let start = await sub.key(); 

return start } }

\`\`\`

\~sub.ts

\`\`\`

export class SubClass{ 

public async key(){ return "you can start the car" } }

\`\`\`

\`\`\`  
import {SubClass} from './subclass.ts' 

import {MainClass} from './mainclass.ts' 

import {mocked} from 'ts-jest/utils'  

jest.mock(./subclass.ts) 

let main = new MainClass() 

let sub = new SubClass() 

let mockedFn = mocked(sub,true)  

mockedFn.key = jest.fn().mockImplementation(() =&gt; console.log('mocking succesfull'))  

afterEach(()=&gt;{ mockedFn.key.mockRestore() // tried mockClear(), mockReset() })

  it('test for main func mocked result',async ()=&gt;{ 

const result = await main.car()  

expect(result).toEqual("mocking succesfull") } 

it('test for main func result',async ()=&gt;{ 

const result = await main.car()  

expect(result).toEqual("you can start the car") // getting result as undefined 

}

\`\`\`
## [5][How can I have intellisense for a dictionary object but also type definition.](https://www.reddit.com/r/typescript/comments/gakqj5/how_can_i_have_intellisense_for_a_dictionary/)
- url: https://www.reddit.com/r/typescript/comments/gakqj5/how_can_i_have_intellisense_for_a_dictionary/
---
    /**
     * @type {{[x:string] : (type1) =&gt; type2}}
    */
    const obj;

This does not give me intellisense for the properties of the `obj`.
## [6][Converting Elements from Script Element?](https://www.reddit.com/r/typescript/comments/gafg9l/converting_elements_from_script_element/)
- url: https://www.reddit.com/r/typescript/comments/gafg9l/converting_elements_from_script_element/
---
I'm starting a new React project in my company and the lead has decided to use TypeScript.  However, there are a TON of standardized company-wide HTML tags (generated using the browser.createElement API) that are imported via the index.html file through a script tag.  This worked well when we had just plain React, but now adding TypeScript is forcing us to add each of these elements individually to the JSX namespace (we have to add to an interface called IntrinsicElements with a type of 'any', though I think we can extend this out a bit with more specific types) and this seems like a kind of hacky fix.  I think I understand is happening (in that TypeScript doesn't like undefined elements), but I am wondering if there's any kind of library or configuration that can help here.

I've read the TS docs for converting from  JavaScript, but the suggestion seems to be mostly rewriting existing code.  Are there any solutions for converting existing elements from a repo and getting them to work with TS?
## [7][Is there a VSCode extension that displays all properties (inherited or not) of a given type?](https://www.reddit.com/r/typescript/comments/gabieg/is_there_a_vscode_extension_that_displays_all/)
- url: https://www.reddit.com/r/typescript/comments/gabieg/is_there_a_vscode_extension_that_displays_all/
---
The title says it all.

OutlinedTextFieldProps extends BaseTextFieldProps, which extends StandardProps, which is a generic receiving FormControlProps, and so on and on. 

It would be nice to have a list of all properties available without hacks (like creating an empty element to explore its props).

https://preview.redd.it/150l682fyrv41.png?width=608&amp;format=png&amp;auto=webp&amp;s=f2bc88c452f3a7ae8ae014a270f70fd6a4c02b62
## [8][RxJs for 5 minutes or The Little Engine that Could…](https://www.reddit.com/r/typescript/comments/gad6kh/rxjs_for_5_minutes_or_the_little_engine_that_could/)
- url: https://www.reddit.com/r/typescript/comments/gad6kh/rxjs_for_5_minutes_or_the_little_engine_that_could/
---
RxJS is everywhere in Angular. So whether you are a junior developer or not, let's consider data transformation, combining the data from two streams and discuss Hot and Cold Observables.

[RxJs for 5 minutes or The Little Engine that Could…](https://2muchcoffee.com/blog/rxjs-for-5-minutes-or-the-little-engine-that-could/)
## [9][Defining a function that has as a parameter a single object that has initial values.](https://www.reddit.com/r/typescript/comments/ga97kq/defining_a_function_that_has_as_a_parameter_a/)
- url: https://www.reddit.com/r/typescript/comments/ga97kq/defining_a_function_that_has_as_a_parameter_a/
---
Lets suppose I have this function :

    function foo(obj/*has property a and b*/) {
    	/*some code*/
    }

I want property `a` to be optional.When no value is provided for `a` then `a` will be equal to `1`, regardless of whether a value was provided to `b`.Correct me if I am wrong but as far as I understand,for that to happen ,and also have all the parameters gathered in `obj` , I have to manually code it inside the function body.

Can I at least write something in typescript (or JSDoc comment) so my IDE(vscode) understands what I want to do?More specifically I want when I type `foo` and hover over it,my IDE will show me that `a` has default value `1` and is supposed to be a number if I give a different value.

Is it possible to add also some comments about the functionality of the parameter `a` that will also be visible when I hover over the function (and not just when I hover over the parameter)?

This is what I have done so far :

    /**
     * @param {object} obj
     * @param {number} [obj.a=1] - some comment about the functionality of the parameter
     * @param {number} obj.b
    */
    function foo(obj) {
    	/*some code*/
    }

In vscode it does not show me initial value. For the comment about each parameter I have to hover over each individual parameter.
## [10][TypeScript VS Flow: Type Checking Front End JavaScript](https://www.reddit.com/r/typescript/comments/gaguon/typescript_vs_flow_type_checking_front_end/)
- url: https://blog.bitsrc.io/should-you-use-typescript-or-flow-abb2716b68e5
---

## [11][I made a simple TypeScript REST API boilerplate with integration tests and deployment setup](https://www.reddit.com/r/typescript/comments/gaf5wt/i_made_a_simple_typescript_rest_api_boilerplate/)
- url: https://github.com/nya1/rest-api-boilerplate
---

