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
## [2][Animal Tribes: How to create a full-stack Typescript GraphQL Application](https://www.reddit.com/r/typescript/comments/hxufeg/animal_tribes_how_to_create_a_fullstack/)
- url: https://www.reddit.com/r/typescript/comments/hxufeg/animal_tribes_how_to_create_a_fullstack/
---
This tutorial teaches how to create an application, a game called Animal Tribes, from scratch using  Typescript, Graphql, NodeJs, ReactJS, and MongoDB.

* [Part 1—Project Overview]( https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-76786e5520ed?source=friends_link&amp;sk=1430f7c8bc45d0192f8052bb0569fd3e)

* [Part 2— Backend](https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-pt-2-backend-cae1735f13dd?source=friends_link&amp;sk=8fbd56e780dedf980ecbcb23358e9148)

* [Part 3— Frontend](https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-pt-3-frontend-dc69f71e1d62?source=friends_link&amp;sk=dbd77c7eef65c081f0c1053bbb1335af)

* [Part 4— Deploy to Heroku](https://medium.com/@samarony.barros/animal-tribes-how-to-create-your-first-full-stack-typescript-graphql-application-e7891ec4963a?source=friends_link&amp;sk=d3a77a7a3d0e4ab45c0b8750250d8595)
## [3][Experienced front-end dev yet really basically no experience packing and publishing npm packages. I've got a question about what's necessary to make a library tree shakeable.](https://www.reddit.com/r/typescript/comments/hxwwog/experienced_frontend_dev_yet_really_basically_no/)
- url: https://www.reddit.com/r/typescript/comments/hxwwog/experienced_frontend_dev_yet_really_basically_no/
---
Just like in react-bootstrap, I'm wondering how I would need to structure a component library so that I can include a single component without loading the entire thing.
What does the project structure have to look like? What does the build HAVE to look like? What about the package.json and webpack config?


EG. @my/library has ComponentA through componentZ.

import {ComponentA, ComponentB} from @my/library
I want this to just include these components, not the whole library.

Follow up - what if there is shared code among componentA andB that they both pull in? Will that be included twice in the bundle?

Thanks!
## [4][Using PassportJS within a decorator](https://www.reddit.com/r/typescript/comments/hxzyqd/using_passportjs_within_a_decorator/)
- url: https://www.reddit.com/r/typescript/comments/hxzyqd/using_passportjs_within_a_decorator/
---
Hey guys! I've being testing and trying to find out a way to do this. So basically, I've used PassportJS (with JWT) in some different projects, but always as a middleware in express. I wanted to step up that, and instead build my own custom method decorator, that would be put in some GraphQL resolver queries and mutations (using TypeGraphQL).

The idea is that if a Query / Mutation / Subcription is decorated with that, it would do some "passport.logIn" type of operation, retrieving the request of that flow, parsing headers, etc.

I'm pretty stuck on that, because as much as I can see, PassportJS is only meant to be used as a middleware, not as an "standalone function" you can call wherever. But the question that raises is: if I only have a "/graphql" endpoint for my whole app, and I use the PassportJS middleware, there is no point in putting in decorators, as PassportJS is already filtering and "middleware"-ing the request. But then how could I expose a logIn mutation? I would need that to be exposed without requiring authentication.

A bit lost here, happy to see if anybody has any idea about it. Cheers!
## [5][Gamedev Patterns and Algorithms with TypeScript. Game Loop (Part 1/2)](https://www.reddit.com/r/typescript/comments/hxovmt/gamedev_patterns_and_algorithms_with_typescript/)
- url: https://medium.com/@gregsolo/gamedev-patterns-and-algorithms-with-typescript-game-loop-part-1-2-699919bb9b71
---

## [6][Assigning a number to a var of type string using another variable?](https://www.reddit.com/r/typescript/comments/hy39ws/assigning_a_number_to_a_var_of_type_string_using/)
- url: https://www.reddit.com/r/typescript/comments/hy39ws/assigning_a_number_to_a_var_of_type_string_using/
---
    let userName: string;
    let userInputAny: any;
    userInputAny = 5;
    
    userName = userInputAny
    userName += 5
    console.log(userName) // 10

New to TS. Why are operations such as these allowed?
## [7][I have two modules called List.ts and ListNode.ts which both export a class . List class depends on ListNode class , but ListNode depends on the type of class List. How do I import the type of class List in ListNode ?](https://www.reddit.com/r/typescript/comments/hxw1pp/i_have_two_modules_called_listts_and_listnodets/)
- url: https://www.reddit.com/r/typescript/comments/hxw1pp/i_have_two_modules_called_listts_and_listnodets/
---
Here is an example :

    //List.ts
    import {ListNode} from "./ListNode";
    export class List {
    	//some stuff that use ListNode
    }
    //ListNode.ts
    export class ListNode {
    	ownningLists : List[]
    }

Should I create a class in a separate module just to use it as a type for `List` and `ListNode`? Is this called a non concrete class or class used as interface?
## [8][Followup: Setting up browser-native websockets with exclusive chat rooms. Does this approach make sense?](https://www.reddit.com/r/typescript/comments/hy07dz/followup_setting_up_browsernative_websockets_with/)
- url: https://www.reddit.com/r/typescript/comments/hy07dz/followup_setting_up_browsernative_websockets_with/
---
This is a followup to an earlier question: it basically seems that [Socket.IO](https://Socket.IO), while solving the problem, isn't the best solution to what I'm trying to do - that browsers have come a long way since 2014 when I last used it. 

I'm just having trouble wrapping my head around how to do what I want to do with native websockets, though. 

Specifically, what I'd like to do is make it possible to set up different chat rooms.  Specifically, I'd like to be able to have a user login to a chat room, and make sure that messages sent TO that chat room are broadcast ONLY to those sockets that have joined that chat room.  [Socket.IO](https://Socket.IO) had this as a feature.  

Now, my current idea to approach this would be to set up an in-memory store, (possibly as a singleton class?) to determine which channels which sockets were subscribed to, then do something like:   

```
interface UserSocket {
  userName: string;
  socket: WebSocket;
}
class ChatroomManager {
  private chatrooms: Record&lt;string, UserSocket[]&gt; = {};
  public subscribeToChatroom = (roomName: string, userName: string, socket: WebSocket) =&gt; {
    if(!this.chatrooms[roomName]){
      this.chatrooms[roomName] = [];
    }
    this.chatrooms[roomName].push({userName, socket}); 
  }
  public sendToUsersOfChatroom = (roomName:string, messageBody: any) =&gt; {
    const payload = { type: 'message', roomName, body: messageBody };
    for(let userSocket of this.chatrooms[roomName]){
      userSocket.socket.send(payload); 
    }
  }
}
```

But I'm wondering if there's a simpler way to do this.  Basically, this chatroom manager class would have to have methods for joining a chatroom, leaving a chatroom, etc.  Let's not worry about things like IRC ops, kicks/bans, etc.  

The one thing that I'm glad I'm doing is working in a monorepo, as I'll be able to create a library that can be used by both client and server, with the same types for messages, so I don't have to really create a whole lot of different addEventListener types, what I can do is just create a few basic ones, ('message', 'leave', etc.) and have it run through a reducer-like switching function to determine how it should be handled based upon the data of the message.
## [9][Regex groups say 'property does not exist on type', how can I fix it?](https://www.reddit.com/r/typescript/comments/hxqzy8/regex_groups_say_property_does_not_exist_on_type/)
- url: https://www.reddit.com/r/typescript/comments/hxqzy8/regex_groups_say_property_does_not_exist_on_type/
---
It seems there is an error on the playground to confound things further, but I get no error on `result.groups` in VS Code and it runs with type checking off. The errors are:
```
Property 'time' does not exist on type '{ [key: string]: string; } | undefined'.ts(2339)
Property 'chapterTitle' does not exist on type '{ [key: string]: string; } | undefined'.ts(2339)
```
How can I fix this and make it compile properly in strict mode?

```
#!/usr/bin/env -S deno run --allow-read --allow-write --no-check

const example = `Introduction
---------------------
⌨️ (00:00:00) Introduction`;

const SEARCH_PATTERN = /(?&lt;time&gt;\d{2}:\d{2}:\d{2})\)\u0020(?&lt;chapterTitle&gt;.*)/;

const lines = example.split("\n");

lines.forEach((line) =&gt; {
  const result = line.match(SEARCH_PATTERN)!;

  if (result !== null) {
    let { time, chapterTitle } = result.groups;
    console.log(time, chapterTitle);
  }
});

```

[Playground Link](https://www.typescriptlang.org/play?#code/MYewdgzgLgBApgDwIYFsAOAbOMC8MAGAkmFAE4gAmArsFAJbgBQAtK2+x544BTEg8H8wAKAAxCAXCPFCAlDGJlKNeuHyNGoSLADKAUQCCAJQDCACQD6ABV0AVK9v0A5XDAD0AgPwAeeijgA+ADoUAN4ATAC+ooGhEVHhUv7xVCIhQu4ewAAWSGhQcKRWdFBYvgB0AFRSzgDcqurQMBh0YHAQToiomHAlEJiFAgBE-mD9UjWMjc0QJQBmIKTaSJkCAhNwMji+MEGMMLswdbCkLVQYsHirJShIUEs6BiYW1rYOUgCEY3swdNOCRxAnsFeODwYBOGBk20+nywsCCMG8cAANPssjk8gUitgwk4-gCSgBzchUNAQGpQvZ1EBYEoYED4gQI5GZbK5fKFLBSVSfMKMMKjIA)
## [10][Handling real-time communications - is Socket.IO still a good choice?](https://www.reddit.com/r/typescript/comments/hxhr1c/handling_realtime_communications_is_socketio/)
- url: https://www.reddit.com/r/typescript/comments/hxhr1c/handling_realtime_communications_is_socketio/
---
Back in 2014, I built my first React app using [Socket.IO](https://Socket.IO) because it contained a real-time communication component.  (It was basically Kahoot).  I've been thinking about revisiting real time communication. Chatrooms, basically, and then once I get that working, maybe build some party-game clones (like Cards Against Humanity or a clone of some of the Jackbox games).  At any rate, I'm just wondering if I'm overlooking some information on real time data. 

I know Firebase is often considered for real-time communication, but I'd kind of like to keep this independent of requiring 3rd party services - at least off the bat.
## [11][Type nested object](https://www.reddit.com/r/typescript/comments/hxo7c9/type_nested_object/)
- url: https://www.reddit.com/r/typescript/comments/hxo7c9/type_nested_object/
---
Hi can anyone show me how to type this in the best possible way.

`const initialData = {`  
 `tasks: {`  
 `"task-1": { id: "task-1", content: "Take out the garbage" },`  
 `"task-2": { id: "task-2", content: "Watch my favorite show" },`  
 `"task-3": { id: "task-3", content: "Charge my phone" },`  
 `"task-4": { id: "task-4", content: "Cook dinner" },`  
 `},`  
 `columns: {`  
 `"column-1": {`  
 `id: "column-1",`  
 `title: "To do",`  
 `taskIds: ["task-1", "task-2", "task-3", "task-4"],`  
 `},`  
 `},`  
 `columnOrder: ["column-1"],`  
`};`  
`export default initialData;`
