# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][Where is the best place to learn typescript?](https://www.reddit.com/r/typescript/comments/hym86c/where_is_the_best_place_to_learn_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hym86c/where_is_the_best_place_to_learn_typescript/
---

## [3][How to get intellisense to work with union types?](https://www.reddit.com/r/typescript/comments/hyook5/how_to_get_intellisense_to_work_with_union_types/)
- url: https://www.reddit.com/r/typescript/comments/hyook5/how_to_get_intellisense_to_work_with_union_types/
---
Consider the following example:

    const paintColor = {
      /** This color is good fo X */
      Waltz: "#d9e5d7",
      /** This color is good fo Y */
      "Indiana Clay": "#ed7c4b",
      /** This color is good fo Z */
      Odyssey: "#575b6a"
    };
    
    const locations = {
      /** There are no good school at this location*/
      City: "123 city center road",
      /** There are lots of parking at this location but it is very far*/
      West: "245 some other road"
    };
    
    interface HouseProp {
      color: keyof typeof paintColor;
      location: keyof typeof locations;
    }
    
    const House = ({ color, location }: HouseProp) =&gt; {
      ...
    };

Where House is a react component that renders a house based on the color and location props. 

And this House component is used everywhere throughout the project.

With the current setup, I could use House like this:  


    &lt;House color="Indiana Clay" location="City" /&gt;

The problem is, the intellisense can't pick up on the docs I've written as part of the code:  


https://preview.redd.it/w6zdhj275dd51.png?width=623&amp;format=png&amp;auto=webp&amp;s=ce856f9d33147b03017c06af8261b855d6ea1879

What do I need to do to get this:  


https://preview.redd.it/aqzfb5vh5dd51.png?width=623&amp;format=png&amp;auto=webp&amp;s=3bcd362c1126f8ee5c1ef9786a48cb2d5315649c

P.S. I know that I could turn paintColor and locations into enums, and use things like this:  


    import House, {HouseColor, HouseLocation} from './House';
    &lt;House color={HouseColor["Indiana Clay"]} location={HouseLocation.City} /&gt;

But that component interface just isn't as nice as my original proposal.
## [4][Is every JSX Element to be annotated as a JSX.Element?](https://www.reddit.com/r/typescript/comments/hynhj3/is_every_jsx_element_to_be_annotated_as_a/)
- url: https://www.reddit.com/r/typescript/comments/hynhj3/is_every_jsx_element_to_be_annotated_as_a/
---
Below I wanted to annotate `linkComponents` to be an array of [`JSX.Link`](https://JSX.Link) elements but tsc suggested there was nothing like that.

Curious if you guys just annotate `JSX.Element` for every type of JSX element or if you have a way of using tighter annotations.

    export function makeAllLinksIntoListItems &lt;genericKeyName extends string&gt; (jsxData: Record&lt;genericKeyName, ComponentData&gt;) {
      const allKeys: Array&lt;string&gt; = Object.keys(jsxData);
    
      const linkComponents: JSX.Element[] = allKeys.map((keyName: string) =&gt; {
        const linkComponent: JSX.Element = makeLink(jsxData[keyName as genericKeyName]);
        return linkComponent;
      })
    
      return linkComponents;
    }
## [5][Looking for tutorials on backend monorepo](https://www.reddit.com/r/typescript/comments/hydam2/looking_for_tutorials_on_backend_monorepo/)
- url: https://www.reddit.com/r/typescript/comments/hydam2/looking_for_tutorials_on_backend_monorepo/
---
Hello, I'm interested in building a monorepo using typescript for learning purposes.

I'd like to have two distinct services communicating through a Queue, and I'd like to share the messages types between these two services.

Is there any good resource or tutorial on this? Most of the ones I see involve React and/or browser related stuff. So I'm not sure what's the best way of developing for the backend in a microservice architecture.

Also, what is the best workflow in your opinion? As in build, run watching, and generating deployable artifact?
## [6][Types from imported modules](https://www.reddit.com/r/typescript/comments/hygl3p/types_from_imported_modules/)
- url: https://www.reddit.com/r/typescript/comments/hygl3p/types_from_imported_modules/
---
Hi, I'm using a somewhat complicated project structure, where I have a client and server written in typescript, and a shared library that they both use. In order to include the shared code, I am using npm link and this seems to be working well.

The problem I have is that I would like to define interfaces in the shared repo, and implement them in the server. Something like this:

Shared code:

    export interface IDataSource { get(): Promise&lt;object&gt;; set(): Promise&lt;object&gt;; }

Server code:

    import { IDataSource } from 'shared'; // no relative path, interface import
    export default class DataSource implements IDataSource {} // concrete implementation in server

This is causing me tons of problems with exporting and importing interfaces however, with "refers to a value, but is being used as a type here" and having to import and re-export so many mundane things in my shared repo.

If I don't use npm link and use relative paths it works just fine, but it breaks my desire to have a reference to shared code within the server and client.

I feel like I have a fundamental misunderstanding of how typescript is meant to be used. Is is intended that a library can contain interfaces implemented in other code that includes the library?
## [7][Animal Tribes: How to create a full-stack Typescript GraphQL Application](https://www.reddit.com/r/typescript/comments/hxufeg/animal_tribes_how_to_create_a_fullstack/)
- url: https://www.reddit.com/r/typescript/comments/hxufeg/animal_tribes_how_to_create_a_fullstack/
---
This tutorial teaches how to create an application, a game called Animal Tribes, from scratch using  Typescript, Graphql, NodeJs, ReactJS, and MongoDB.

* [Part 1—Project Overview]( https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-76786e5520ed?source=friends_link&amp;sk=1430f7c8bc45d0192f8052bb0569fd3e)

* [Part 2— Backend](https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-pt-2-backend-cae1735f13dd?source=friends_link&amp;sk=8fbd56e780dedf980ecbcb23358e9148)

* [Part 3— Frontend](https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-pt-3-frontend-dc69f71e1d62?source=friends_link&amp;sk=dbd77c7eef65c081f0c1053bbb1335af)

* [Part 4— Deploy to Heroku](https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-e7891ec4963a?source=friends_link&amp;sk=d3a77a7a3d0e4ab45c0b8750250d8595)
## [8][Experienced front-end dev yet really basically no experience packing and publishing npm packages. I've got a question about what's necessary to make a library tree shakeable.](https://www.reddit.com/r/typescript/comments/hxwwog/experienced_frontend_dev_yet_really_basically_no/)
- url: https://www.reddit.com/r/typescript/comments/hxwwog/experienced_frontend_dev_yet_really_basically_no/
---
Just like in react-bootstrap, I'm wondering how I would need to structure a component library so that I can include a single component without loading the entire thing.
What does the project structure have to look like? What does the build HAVE to look like? What about the package.json and webpack config?


EG. @my/library has ComponentA through componentZ.

import {ComponentA, ComponentB} from @my/library
I want this to just include these components, not the whole library.

Follow up - what if there is shared code among componentA andB that they both pull in? Will that be included twice in the bundle?

Thanks!
## [9][Using PassportJS within a decorator](https://www.reddit.com/r/typescript/comments/hxzyqd/using_passportjs_within_a_decorator/)
- url: https://www.reddit.com/r/typescript/comments/hxzyqd/using_passportjs_within_a_decorator/
---
Hey guys! I've being testing and trying to find out a way to do this. So basically, I've used PassportJS (with JWT) in some different projects, but always as a middleware in express. I wanted to step up that, and instead build my own custom method decorator, that would be put in some GraphQL resolver queries and mutations (using TypeGraphQL).

The idea is that if a Query / Mutation / Subcription is decorated with that, it would do some "passport.logIn" type of operation, retrieving the request of that flow, parsing headers, etc.

I'm pretty stuck on that, because as much as I can see, PassportJS is only meant to be used as a middleware, not as an "standalone function" you can call wherever. But the question that raises is: if I only have a "/graphql" endpoint for my whole app, and I use the PassportJS middleware, there is no point in putting in decorators, as PassportJS is already filtering and "middleware"-ing the request. But then how could I expose a logIn mutation? I would need that to be exposed without requiring authentication.

A bit lost here, happy to see if anybody has any idea about it. Cheers!
## [10][Assigning a number to a var of type string using another variable?](https://www.reddit.com/r/typescript/comments/hy39ws/assigning_a_number_to_a_var_of_type_string_using/)
- url: https://www.reddit.com/r/typescript/comments/hy39ws/assigning_a_number_to_a_var_of_type_string_using/
---
    let userName: string;
    let userInputAny: any;
    userInputAny = 5;
    
    userName = userInputAny
    userName += 5
    console.log(userName) // 10

New to TS. Why are operations such as these allowed?
## [11][Gamedev Patterns and Algorithms with TypeScript. Game Loop (Part 1/2)](https://www.reddit.com/r/typescript/comments/hxovmt/gamedev_patterns_and_algorithms_with_typescript/)
- url: https://medium.com/@gregsolo/gamedev-patterns-and-algorithms-with-typescript-game-loop-part-1-2-699919bb9b71
---

