# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Is there a easy way for us to extend the language of typescript](https://www.reddit.com/r/typescript/comments/es8hdh/is_there_a_easy_way_for_us_to_extend_the_language/)
- url: https://www.reddit.com/r/typescript/comments/es8hdh/is_there_a_easy_way_for_us_to_extend_the_language/
---
Hi All

Is there a easy way for us to extend the language of typescript?  adding my own syntax?

thanks

Peter
## [3][Type vs Interface (React/React-native)](https://www.reddit.com/r/typescript/comments/es2e4c/type_vs_interface_reactreactnative/)
- url: https://www.reddit.com/r/typescript/comments/es2e4c/type_vs_interface_reactreactnative/
---
Hey beautiful people!

I just ported my project (react native) over to typescript but dont get the hang when to use **type and when to use interface**

I looked at some examples, but just added to the confusion, because apparently both have their place but i cannot figure it out

Could you tell me on a high level the difference between these two and if you have expereicne with React/React-Native where to use which?

&amp;#x200B;

Thanks a lot!
## [4][Dev Diary: 5 Tips for Building Beautiful CLIs with Node.js, Yargs, &amp; Ink](https://www.reddit.com/r/typescript/comments/erzcbn/dev_diary_5_tips_for_building_beautiful_clis_with/)
- url: https://medium.com/eximchain/dev-diary-5-tips-for-building-beautiful-clis-with-node-js-yargs-ink-16d184ea0d14
---

## [5][Thoughts on Map and Set in TypeScript](https://www.reddit.com/r/typescript/comments/es2fav/thoughts_on_map_and_set_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/thoughts-on-map-and-set-in-typescript-efd43c3bf2ad?source=friends_link&amp;sk=5863d3228b0d37f3d493bb1d001a2c13
---

## [6][on restricting Generic Arguments with `extends`](https://www.reddit.com/r/typescript/comments/es2zht/on_restricting_generic_arguments_with_extends/)
- url: https://twitter.com/phry/status/1219701490530562049
---

## [7][JavaScript design patterns #5. The Observer pattern with TypeScript](https://www.reddit.com/r/typescript/comments/erel2u/javascript_design_patterns_5_the_observer_pattern/)
- url: https://wanago.io/2020/01/20/javascript-design-patterns-observer-typescript/
---

## [8][Lack of Private Access Interfaces](https://www.reddit.com/r/typescript/comments/erj669/lack_of_private_access_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/erj669/lack_of_private_access_interfaces/
---
I understand that the focus of TypeScript interfaces is to create public contracts, so that objects within a running application inter-operate appropriately with one another. However, I think some would still find benefits if interfaces also specify the private side of a class. Private contracts ensure that the objects within a running application *intra-operate* appropriately with themselves. 

I understand this behavior is similar in other programming languages, namely C#. I've recently run into this issue when structuring my application with composition as opposed to inheritance. If my classes instead are defined in terms of what they do, TypeScript interfaces with their duck-typing become a powerful tool. Even if some of the characteristics of what a class does is private, I would like to specify this behavior with the usage of interfaces.

What do you all think? Have there ever been cases where you wish that interfaces could provide type assurance for the private side of your classes? Am I trying to inappropriately overload the purpose of interfaces?
## [9][typescript in a nodejs project starter](https://www.reddit.com/r/typescript/comments/erd5j9/typescript_in_a_nodejs_project_starter/)
- url: https://www.reddit.com/r/typescript/comments/erd5j9/typescript_in_a_nodejs_project_starter/
---
Does anyone have a reference/instructions/template on how to setup a nodejs project initially with typescript. I have googled and followed links like below multiple times but get errors like "Cannot use import statement outside a module" or "Error: Cannot find module index.ts". All I am simply doing in my file to test is just a VERY basic file write and I cannot even do that.

import \* as fs from "fs"

fs.writeFileSync('notes.txt','Hello')

to run the project I am just running "node index.ts"

I am running node v13.6.0

[https://levelup.gitconnected.com/how-to-set-up-a-typescript-node-js-app-in-5-minutes-93ffee3b1768](https://levelup.gitconnected.com/how-to-set-up-a-typescript-node-js-app-in-5-minutes-93ffee3b1768)

[https://dev.to/theghostyced/setting-up-a-node-typescript-project-in-under-4-minutes-4gk2](https://dev.to/theghostyced/setting-up-a-node-typescript-project-in-under-4-minutes-4gk2)
## [10][Question, considering learning Typescript](https://www.reddit.com/r/typescript/comments/er8fjv/question_considering_learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/er8fjv/question_considering_learning_typescript/
---
Hi everyone,

I'm considering learning Typescript as I'm hearing lots of great things about it however my dilemma is that I'm half way through a React Native project, are there any downsides to using Typescript for only part of you code base instead of the entire project?

I should mention that if I like typescript I would consider rewriting the code base overtime

Thanks!
## [11][Getter and Setter on String Array](https://www.reddit.com/r/typescript/comments/er9opl/getter_and_setter_on_string_array/)
- url: https://www.reddit.com/r/typescript/comments/er9opl/getter_and_setter_on_string_array/
---
Hey guys!

Okay, not sure about this:

I am using a simple **getter** and **setter** as shown below:  
If I use different names for the getting, then it works but if I use the same name, then it fails:

`private _person: string[] = [];`  
 `get person(){`  
 `return this._person;`  
  `}`  
 `set person(value: string){`  
 `this._person.push(name);`  
  `}`  


I get the following error:  
 ***Type 'string\[\]' is not assignable to type 'string'.ts(2322)***
