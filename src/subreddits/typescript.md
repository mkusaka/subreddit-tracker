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
## [2][A deployed version of my ECS circles demo!](https://www.reddit.com/r/typescript/comments/hv3y5h/a_deployed_version_of_my_ecs_circles_demo/)
- url: https://avokadoen.github.io/ts-ecs-circles-deployed/
---

## [3][Using an AWS sdk call in class contructor](https://www.reddit.com/r/typescript/comments/hv5qbr/using_an_aws_sdk_call_in_class_contructor/)
- url: https://www.reddit.com/r/typescript/comments/hv5qbr/using_an_aws_sdk_call_in_class_contructor/
---
Hey guys, I am trying to use a call to AWS Secret Manager in the constructor of an API class. The instance will be used to send calls to that API, and the call to AWS Secret Manager is needed to fill the auth.

However, Typescript warns me to not use await in the constructor. Since the AWS Secret Manager calls seem to be async, how would I use this service in my API class? Any advice would be welcome.


```js
export class SOMEServiceAPI {
    private readonly options: any;

    constructor() {
        this.options = {
            timeout: 60000,
            method: 'PUT'
        };
        await this.makeOptions();
    }


    makeRequest(body: string) {
        return Object.assign({}, this.options, {
            body
        });
    }


    async put(event: IncomingEvent): Promise&lt;Response&gt; {
        const request = this.makeRequest(JSON.stringify(event.event_payload));
        const date: Date = new Date();
   &lt;....&gt;
    }

    private makeOptions = async() =&gt; {
        const Auth = await awssdkcall; 
&lt;..fills in the authHeader and Api key..&gt;



        return this.options.headers = {
            Authorization: authHeader,
            'x-api-key': apiHeader
        };
    }
}

```
## [4][Trying to follow an article for implementing Context with Typescript but getting an error. Anyone know how to resolve this error?](https://www.reddit.com/r/typescript/comments/hv0sa6/trying_to_follow_an_article_for_implementing/)
- url: https://www.reddit.com/r/typescript/comments/hv0sa6/trying_to_follow_an_article_for_implementing/
---
**Article for reference:** https://kentcdodds.com/blog/how-to-use-react-context-effectively

**My code for reference:** https://pastebin.com/xNPH1pRf

I am getting the following error:

&gt;No overload matches this call.
  Overload 1 of 5, '(reducer: ReducerWithoutAction&lt;any&gt;, initializerArg: any, initializer?: undefined): [any, DispatchWithoutAction]', gave the following error.
    Argument of type '(state: State, action: Action) =&gt; {}' is not assignable to parameter of type 'ReducerWithoutAction&lt;any&gt;'.
  Overload 2 of 5, '(reducer: (state: State, action: Action) =&gt; {}, initialState: never, initializer?: undefined): [never, Dispatch&lt;Action&gt;]', gave the following error.
    Type 'number' is not assignable to type 'never'.ts(2769)

The line of code (line 25) the error appears on:

`const [state, dispatch] = React.useReducer(orderReducer, { count: 0 })`

Any help would be much appreciated!
## [5][How do you run a firebase cloud function snapshot that tou coded in swift?](https://www.reddit.com/r/typescript/comments/hv5ydj/how_do_you_run_a_firebase_cloud_function_snapshot/)
- url: https://www.reddit.com/r/typescript/comments/hv5ydj/how_do_you_run_a_firebase_cloud_function_snapshot/
---
How do you run an elaborate firebase snapshot via cloud functions instead of on the client?

//Here is the structure of the RT DB

    "people" : {
     "1ZWT7FAE2qThNQfBj7tbMO7BnMo1" : {
      "Coordinates" : {
         "latitude" : 50.054738,
         "longitude" : 8.226809826085624
       "peopleWhoLike2" : {
         "1vLVFwrXrHUoakmDrnQKwbv08Yj1" : 1581548952597,
         "F9NX0UCG4fVHCKFk2VZ1NZKsLro2" : 1586210112155,
         "IrrBgFY9C1ekMmHUkQRzc5LhbDu1" : 1581547417432,

///Here is the elaborate snapshot that I want to run via cloud functions. I already have the JS,TS files setup, but have never coded in JS.

    refArtists.observe(DataEventType.value,  with: {snapshot in
            
            if snapshot.childrenCount&gt;0{
                
                self.people.removeAll()
                
                for people in snapshot.children.allObjects as! [DataSnapshot] {
                    
                        if people.key != thisUsersUid {
                            print("peoplekey",people.key)
                            
    
                   
                        let peopleObject = people.value as? [String: AnyObject]
                        let peopleEducation = peopleObject?["Education"] as? String
                        let peopleWhatIamConsideringBuying = peopleObject?["WhatIamConsideringBuying"] as? String
                        let peoplePhotoPosts = peopleObject?["PhotoPosts"]  as? String
                        let peopleimageDownloadURL = peopleObject?["imageDownloadURL"]  as? String
                      ......
         let peopl = Userx(Education: peopleEducation, .........)
                     
                        self.people.append(peopl)
    
                            } else {
                                print ("w")
                            }
                                        } else {
                                print ("alphaaa")
                                                                                                                                                              
                         }
    }}
            }   
        })
## [6][VS Code and TS: Type error parser for these errors?](https://www.reddit.com/r/typescript/comments/hup5uj/vs_code_and_ts_type_error_parser_for_these_errors/)
- url: https://i.redd.it/ts5fqfqho1c51.png
---

## [7][Best resources to learn Typescript](https://www.reddit.com/r/typescript/comments/hv2nhf/best_resources_to_learn_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hv2nhf/best_resources_to_learn_typescript/
---
Hello everyone, i was wondering what is the best way to learn Typescript since it seems like a solid option for web development.
## [8][Can I make multiple assertions with a single function?](https://www.reddit.com/r/typescript/comments/hultk9/can_i_make_multiple_assertions_with_a_single/)
- url: https://www.reddit.com/r/typescript/comments/hultk9/can_i_make_multiple_assertions_with_a_single/
---
I have a case where a condition being true describes the state of two entities.  Here's an example:

    const compare(x: UnknownEntity, y: UnknownEntity): x is EntityA =&gt; {
        return x.property === y.property;
    }

The above example works, but it only asserts that `x is EntityA`.  In my case this condition also proves that `y is EntityB` but I don't know how (or if it's possible) to make the return type something like:

    x is EntityA AND y is EntityB

Is it possible to make two assertions with a single function?

Appreciate any help.

Edit: More context here: https://www.reddit.com/r/typescript/comments/hultk9/can_i_make_multiple_assertions_with_a_single/fynv0mr/
## [9][A directory structure for React projects](https://www.reddit.com/r/typescript/comments/hui3yt/a_directory_structure_for_react_projects/)
- url: https://medium.com//a-directory-structure-for-react-projects-99fb084c61f1?source=friends_link&amp;sk=e5c0c00d0e03e6661ec7e710f84eebfb
---

## [10][How to Infer Types from Schema](https://www.reddit.com/r/typescript/comments/hub8rd/how_to_infer_types_from_schema/)
- url: https://www.reddit.com/r/typescript/comments/hub8rd/how_to_infer_types_from_schema/
---
Hi all. I've recently started using typescript more heavily, and  I'm working a generic form-validation library. One of the things I'd like to be able to do is pass in a data model, and be able to get type inference/intellisense from the keys from that object. I've not learned all the details of how to use generics yet, which I know will be key for this question, so any pointers towards what I am missing would be greatly appreciated.

A simplified example:

```
interface SpecificSchemaInterface {
    fields: Record&lt;String, any&gt;
}

interface SpecificSchemaInterface {
    fields: {
      firstName: string;
      lastName: string;
      age: number;
    }
}

const schema = {
  fields: {
    firstName: 'Bob,
    lastName: 'Bobson',
    age: 25,
  },
};

function createValidatorFromSchema(s: SpecificSchemaInterface) {
  // ???
}

const validator: Validator&lt;SpecificSchemaInterface&gt; = createValidatorFromSchema(schema)

// Then validator would have an interface like:
interface Validator {
  rules: {
    firstName(firstName: string): boolean,
    lastName(lastName: string): boolean,
    age(age: number): boolean, // note that the type is inferred here from the schema
  }
}
```

My questions are that then
1) I probably want something like `Validator&lt;SpecificSchemaInterface&gt;`?
2) I need to be able to both extract the object keys, and the types from the schema in a reusable way inside my new `Validator`

I hope this isn't too vague, but if there is some terminology I should be looking into (beyond just "look at generics") or guides that go into this sort of thing, that would be very helpful.
## [11][Am I generating `index.d.ts` file for an `npm` module `@camwiegert/typical` correctly?](https://www.reddit.com/r/typescript/comments/hugbmu/am_i_generating_indexdts_file_for_an_npm_module/)
- url: https://www.reddit.com/r/typescript/comments/hugbmu/am_i_generating_indexdts_file_for_an_npm_module/
---
I am a noob at TS &amp; certainly don't know how to write `index.d.ts`. Guided by some auto-completion inside VSCode &amp; some of my existing knowledge, I managed to cobble this together.

I am writing types for `@camwiegert/typical` npm module. Source code could be found at https://github.com/camwiegert/typical

Here is the main file `typical.js` for brevity:

```js
export async function type(node, ...args) {
    for (const arg of args) {
        switch (typeof arg) {
            case 'string':
                await edit(node, arg);
                break;
            case 'number':
                await wait(arg);
                break;
            case 'function':
                await arg(node, ...args);
                break;
            default:
                await arg;
        }
    }
}

async function edit(node, text) {
    const overlap = getOverlap(node.textContent, text);
    await perform(node, [...deleter(node.textContent, overlap), ...writer(text, overlap)]);
}

async function wait(ms) {
    await new Promise(resolve =&gt; setTimeout(resolve, ms));
}

async function perform(node, edits, speed = 60) {
    for (const op of editor(edits)) {
        op(node);
        await wait(speed + speed * (Math.random() - 0.5));
    }
}

export function* editor(edits) {
    for (const edit of edits) {
        yield (node) =&gt; requestAnimationFrame(() =&gt; node.textContent = edit);
    }
}

export function* writer([...text], startIndex = 0, endIndex = text.length) {
    while (startIndex &lt; endIndex) {
        yield text.slice(0, ++startIndex).join('');
    }
}

export function* deleter([...text], startIndex = 0, endIndex = text.length) {
    while (endIndex &gt; startIndex) {
        yield text.slice(0, --endIndex).join('');
    }
}

export function getOverlap(start, [...end]) {
    return [...start, NaN].findIndex((char, i) =&gt; end[i] !== char);
}
```

I want to write types for this file. So I made `src/@types/@camwiegert/typical/index.d.ts` like:

```ts
declare module '@camwiegert/typical' {
	export function type(node: HTMLElement, ...args: any[]): Promise&lt;void&gt;

	declare function edit(node: HTMLElement, text: string): Promise&lt;void&gt;

	declare function wait(ms: number): Promise&lt;void&gt;

	declare function perform(node: HTMLElement, edits: any, speed?: number): Promise&lt;void&gt;

	export function editor(edits: any): Generator&lt;(node: any) =&gt; number, void, unknown&gt;

	export function writer(
		[...text]: Iterable&lt;any&gt;,
		startIndex?: number,
		endIndex?: number
	): Generator&lt;string, void, unknown&gt;

	export function deleter(
		[...text]: Iterable&lt;any&gt;,
		startIndex?: number,
		endIndex?: number
	): Generator&lt;string, void, unknown&gt;

	export function getOverlap(start: any, [...end]: Iterable&lt;any&gt;): number
}
```

Idk if I converted it correctly or not. Can anyone check if it's correct?
