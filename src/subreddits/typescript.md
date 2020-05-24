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
## [2][Best Way To Use markAllAsTouched in Angular Forms | Geekstrick](https://www.reddit.com/r/typescript/comments/gpomm6/best_way_to_use_markallastouched_in_angular_forms/)
- url: https://www.geekstrick.com/markallastouched-in-angular-forms/
---

## [3][first steps to take to learn typescript? (most common question asked over time probably)](https://www.reddit.com/r/typescript/comments/gpnwf9/first_steps_to_take_to_learn_typescript_most/)
- url: https://www.reddit.com/r/typescript/comments/gpnwf9/first_steps_to_take_to_learn_typescript_most/
---
as this is liekly the most common question asked over time

im guessing some of the very knowledgable ppl here had sometimes during the past created a script in which other ppl would just copy / paste to new ppl that ask this question (since you're coders and would want to automate things so they can be made efficient &amp; effective)

here's everything i know about coding so far:

[https://www.reddit.com/r/learnpython/comments/goy5xa/where\_to\_continue\_learning\_coding/](https://www.reddit.com/r/learnpython/comments/goy5xa/where_to_continue_learning_coding/)

&amp;#x200B;

it seems like typescript is mainly for 'enterprise use' and for 'angular'  [https://www.quora.com/Should-I-learn-JavaScript-before-learning-TypeScript](https://www.quora.com/Should-I-learn-JavaScript-before-learning-TypeScript) so maybe should learn something simpler but it doesnt seem like there are any simpler options out there...
## [4][Exporting a Jest mock of a module with type casting](https://www.reddit.com/r/typescript/comments/gpgg8d/exporting_a_jest_mock_of_a_module_with_type/)
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
## [5][Verify cookie with JWT payload inside.](https://www.reddit.com/r/typescript/comments/gppio8/verify_cookie_with_jwt_payload_inside/)
- url: https://www.reddit.com/r/typescript/comments/gppio8/verify_cookie_with_jwt_payload_inside/
---
https://preview.redd.it/nm5u38x6qp051.png?width=596&amp;format=png&amp;auto=webp&amp;s=b1c75a2fd39e0e077840f4ee40188c3c4e1607ed

Hi everyone, I'm learning TypeScript and I'm stuck in that problem.

The return of function verify is the code commented above the interface Decoded

Thank you for your help!

&amp;#x200B;
## [6][ORM Challenges](https://www.reddit.com/r/typescript/comments/gp8fx8/orm_challenges/)
- url: https://www.reddit.com/r/typescript/comments/gp8fx8/orm_challenges/
---
Hey All,

I have some fairly extensive professional experience using python with sql alchemy, Django, etc. I am currently working on building a side project with typescript being my language of choice. 

I’ve noticed there is much more scaffolding necessary compared to the python ecosystem. Currently I’m using typeorm. It feels also though everything is extremely manual from setting up testing classes for end to end tests, connection management, simple things like get_or_create statements. 

I love typescript as a language but I feel far less productive building out a backend with the tools I’ve tried so far. 

I’m not sure if I’m simply approaching this with the wrong perspective or using poor tools. 

Does anyone have any recommendations or thoughts on the above? I may just have a poor perspective coming from a more mature python ecosystem.
## [7][Typescript is not checking function arguments.](https://www.reddit.com/r/typescript/comments/gp7sl4/typescript_is_not_checking_function_arguments/)
- url: https://www.reddit.com/r/typescript/comments/gp7sl4/typescript_is_not_checking_function_arguments/
---
Why TS is not checking function arguments?

    interface Functions {
      add(x: number, y: number): number;
      isPositive(n: number): boolean;
    }
    const functions: Functions = {
      add: () =&gt; 5,
      isPositive: () =&gt; true,
    };

TS doesn't complain although `x` and `y` arguments are missing in `add` function. But does his job if the return value (number in add) doesn't match.
## [8][Change property with decorator in TypeScript?](https://www.reddit.com/r/typescript/comments/gpc10u/change_property_with_decorator_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gpc10u/change_property_with_decorator_in_typescript/
---
I'm trying to get a simple decorator example in TS to work. I'm simply trying to change the value of an instance member. Since property decorators receive the prototype object as a target argument, defining a property will actually add it to the prototype and not the instance. So, I can't access it directly. I'm not sure how I could actually change the value here?

    function valueDecorator(target: any, propertyKey: any) {
        Object.defineProperty(target, propertyKey, {
            value: 5,
            writable: true,
        })
    }
    
    
    class Example {
        @valueDecorator
        public value: number;
    
        constructor() {
            this.value = 0
        }
    }
    
    const ex = new Example()
    console.log(ex.value)

**The output is still 0. I expected it to be 5.**
## [9][query in Optional parameter: typescript](https://www.reddit.com/r/typescript/comments/gp3t7g/query_in_optional_parameter_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gp3t7g/query_in_optional_parameter_typescript/
---
Hello All,

I  have a query

topic: optional parameter

I want to call a specific optional parameter? is there a way to do that, without sending all  the optional parameter. For eg:- 

in the below example I want to send value only for z variable without sending x and y value ,

something like

add( z: 5);

function add (x ?:  number,  y ?: number, z ?: number){

}

is there a way to do it?

thanks in advance
## [10][TypeScript taught me something today](https://www.reddit.com/r/typescript/comments/goovef/typescript_taught_me_something_today/)
- url: https://www.reddit.com/r/typescript/comments/goovef/typescript_taught_me_something_today/
---
So I'm working on a form today, and when I was coding one particular select input, I noticed I needed two different pieces of information from whatever I was selecting. I tried a few different things but nothing was working. As I was noodling around, I hovered over the value attribute in the option tag, and TS told me that the value can be a string\[\] as well as a string or a number. I guess I would have eventually solved the problem, but TypeScript taught me how to get multiple values and from there it was just a matter of of hovering instead of reading a lot of documentation. Thank you Typescript!
## [11][Intellisense not working on returned parameter (ide : vscode) .](https://www.reddit.com/r/typescript/comments/gp6ifp/intellisense_not_working_on_returned_parameter/)
- url: https://www.reddit.com/r/typescript/comments/gp6ifp/intellisense_not_working_on_returned_parameter/
---
    function component&lt;T&gt;(p: {
    	actions : () =&gt; T,
    	htmlTemplate: &lt;S extends T&gt;(actions : S) =&gt; any
    }) {
    	return p;
    }
    
    component({
    	actions : () =&gt; ({
    		a : (x,y) =&gt; x+y
    	}),
    	htmlTemplate : (actions) =&gt; { 
    		/*try to dot into the action and you will get intellisense*/ 
    		actions;
    		/*but if you dot here you will get no intellisense*/ 
    		return actions;
    	}
    });

here is it with syntax highlighting just for it to be more readable :

https://preview.redd.it/1ogu0pcy5j051.png?width=658&amp;format=png&amp;auto=webp&amp;s=093845151b5d35b74f073d872443eb95634a1473

do you know why intellisense is not working and how to fix that ?
