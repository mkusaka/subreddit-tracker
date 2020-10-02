# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
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
## [2][I could use some help with properly typing an event emitter function](https://www.reddit.com/r/typescript/comments/j3s87i/i_could_use_some_help_with_properly_typing_an/)
- url: https://www.reddit.com/r/typescript/comments/j3s87i/i_could_use_some_help_with_properly_typing_an/
---
I have a use case for an event emitter like API, but I can't seem to properly type one of the functions.

I have this interface

```
interface EventChannel&lt;API extends {}&gt; {
  emit: (eventName: keyof API, payload: API[typeof eventName]) =&gt; void
}
```

However, that doesn't feel right, because we get into this sort of situation

```
interface ExampleInterface {
  onNumber: number;
  onString: string;
}

type MyInterface = EventChannel&lt;ExampleInterface&gt;;
```

And now when you look at `MyInterface.emit`, the `eventName` is correct, but the payload is `number | string`, which isn't entirely correct.

Is there a way to make it so that the type of the payload matches the type expected by `ExampleInterface[eventName]`?
## [3][Cannot get my head around this typing issue.](https://www.reddit.com/r/typescript/comments/j3ghqt/cannot_get_my_head_around_this_typing_issue/)
- url: https://www.reddit.com/r/typescript/comments/j3ghqt/cannot_get_my_head_around_this_typing_issue/
---
Here's my challenge (I've simplified it a little). I have a class call `TestClient` which I use for writing integration tests. There are two kinds of users who can access my application, web users (standard humans who log in) and service users (machine-to-machine users who access with an API key). These users have different structures


    interface WebUserIdentity {
       userId: string
       emails: string[]
       type: 'Web'
    }

and 


    interface ServiceUserIdentity {
       authorizedOrganization: string
       createdByUserId: string
       type: 'Service'
    }


What I want to be able to do is have code in my `TestClient` that can do something depending on what type of user I have.



	class TestClient&lt;IdentityType&gt; {
		private identity: IdentityType;

		public addEmailForUser(emailAddress: string, optionationUserId?: string){
			 let userId;

			 if(this.identity.type === Web){
				 userId = this.identity.userId;
			 }else{
				 // assume there's error checking in here.
				  userId = optionationUserId;
			 }


			 const newEmail = addTheEmailToTheSystem(emailAddress, userId);
			 if(this.identity.type === Web){
				 this.identity.emails.push(newEmail);
			 }
		}
	}

This way when I'm writing a test I can have something like this


    const webTestClient: TestClient&lt;WebUserIdentity&gt;
    webTestClient.addEmail('a@b.com');


My problem is in the generic for `TestClient&lt;IdentityType&gt;`. I'm pretty sure I'll need conditional types but I can't get it to work so that once I confirm the type I can access the appropriate properties.
## [4][Need ideas for a 30 minutes workshop on Typescript](https://www.reddit.com/r/typescript/comments/j349k2/need_ideas_for_a_30_minutes_workshop_on_typescript/)
- url: https://www.reddit.com/r/typescript/comments/j349k2/need_ideas_for_a_30_minutes_workshop_on_typescript/
---
So, I started convincing my team (all JS devs) to make the switch to Typescript, and the CTO of my company told me to make a 30 minutes workshop to teach them about typescript and SHOW how it could benefit our company. 

We have a front end in React and React-Native and a back-end in Node, so to me the advantage is obvious, but I feel that some people still needs a bit of convincing.

So this workshop needs both to be an introduction to TS as well as a showcase of how powerful it can be, any ideas of how I could/should do it? 

P.S.: I don't have any experience making workshops either, so any tips could help, really.
## [5][Working Static Type Inference Without Functions as Type Proxies](https://www.reddit.com/r/typescript/comments/j360y4/working_static_type_inference_without_functions/)
- url: https://www.reddit.com/r/typescript/comments/j360y4/working_static_type_inference_without_functions/
---
Holy Smokes!

This took me days to work out, and it needed the nominal types hack to work, but [boy does it work](https://www.typescriptlang.org/play?#code/C4TwDgpgBAKgcgHgMpQgD2BAdgEwM5R7ABOAllgOYB8UAvFCgGRQDeA+gFxQCuuEAZuQg4AvgG4AUBNCQoAJRgB5AEYArCAGNgdVsC4ByAPZrNwfQBooxLgpjgIAWQCGYS4Zsw7kZ2HHT78jBIJOQUOix6UPpEZJT6fjLQCnDcALbKEMThkfpYaRnE8ZKJgQBChoYANhBOWNkGyhXVtUVSJbYBtBJQPVAAPoEq6lrdvQMKwbEUoz3j8PmZM-1lTTVYEpL+sh3eLuEA2gDWHDGhALoeXhB+W9AAklj8mQgwNF29sKgY2PiBk6FQAD8hBClCgHCWMC+mFwBGSCyywLy6Uy4Mh0J+cJg5SqayBUEauNqaI+UPQMN+CiGpnxLCWH32xAAjlByFBDhAQIZ+LB9vpCmcLvSPr0Hk9iC8+QLGUyzlQliIoMw6SLevtDGBWXUOVyeTA+YZ9ILARDVaqxc99UYjeqwHKFSTelgIAA3Rabfi8LSkQx1PCkVJgapXF4Y2EgqZUAAUkRgAEouBEuPAXlRFSxiBBgNxiHUk1BtE4CCnXiIbp6sN7fRGowndMnENFQRR9GnWJns7nCAGgxArlGm1N9HGxCIJBWq3VkbXE7HG8iCq30x2c36e8H7AOF5lh6OpBPgD6-SEtAg5GHKZ57D5o5mmdxSJmcDYE0t8ymjCYtK3LNZ5G5EzLCQVV6FcuxA1UcmMYYzCgItYEbaDTB-YUej-O8HyfcxUKgdxWDHD4-DHccvUPatM0qJw0GEBBFAvLErhvKMNTIrAnEqLhFFfD530Qr8zCoX9AIAqBFDHCCrCzVdWBwqD+P0ODiz4mCULNKxhJwvCWKPdjsMIiRiIPI8oFSUhqJwM96MCRiXEsOjyUxazrxcW8IHvR9hBsNwwFY9jOO43peIQT8VME9T-1wzjxKWMC81kgwkO-RSEOCxKBL0s10LczDhAy1UtJ8nTKjyqAx2IjRfSICMAXoGJa0kCqsCq7csnoacRwkRqqsJZo6lq9c+03fQerWXcpC67QkEMDRSHY0hMB0U4tCjCS2NSCAuFOSgSqcCgNqgFqSv4WoNBALgRtqbCRDjW4GGm2bKnm6B6AtCVEm5O6ZrmzAqE2CbYGIJwcAgPBUmJehTPMlaljW-atooHa9q4Q6DPMCTlAfSpgb-C6sBKgB3CAscyc7VkupYidMMgvqwUmiTxgybpKGBAeB0HwagV6EHevVWZBsGsF+qQAHphagF05pwXD+LwTrKu0PAAEYuCmr7HoW+gWFhgwAAtuAKMxLF2-aABYAAZLGOyszoLYhuGuBr5YLZWAaB-mOa1px1oMSpuA5Cw4KRqAAGYACZLEprRqdm2nbftvx-uAUPkz59m+tYbWomIEAWiNoPQ4ATksDHSGJv8SHtyxCbLrhjsqPAHZFsXyAlx6pbS2X-rwZPPoep7wkzk3w8D-b9BwXgsENS2TptvJKkqBOneAYOU7dtOB690fVCcDRDgD42uFDgBWYvMexgwQBB+IJCAA)!

I've been looking at the [various runtime type &amp; validation libs](https://github.com/moltar/typescript-runtime-type-benchmarks) and I'm generally impressed by what they do, but the dependence on function objects to represent types was driving me crazy, so I set out to create the same effect without them.

It was surprisingly challenging to coax the type inference engine to do what I wanted, and IMHO the results are amazing and I'm super stoked about it.

Unfortunately I simply could not avoid using functions to build structures, but for singular types the proxies are just the name of the type!

If there's any interest I will try to roll it into a library.
## [6][[help] Restricting a mapped type to a list of properties to be used in another type?](https://www.reddit.com/r/typescript/comments/j33067/help_restricting_a_mapped_type_to_a_list_of/)
- url: https://www.reddit.com/r/typescript/comments/j33067/help_restricting_a_mapped_type_to_a_list_of/
---
so hard to explain here (or to google) exactly what I'm looking for. I want to 

1. have a single source of possible events (either as a union of strings, or keys in an interface)
2. then use that list to build new types where I get hints from that list of possible events

Let's imagine some fake DOM events:


    interface DomEvents {
        'click': DomClickEvent;
        'keydown': DomKeyDownEvent;
    }

then I want to use these events (click, keydown) in another type while getting autocompletions and I don't know how to do that.

I want to do something like this (type a type), but it's obviously impossible:

    type DomEventMap: {[K in keyof DomEvents]: () =&gt; void} = {
         // get autocomplete here for the props (click, keydown)
    }

I'm sure I'm spacing out on something I've seen before, but I can't find a good example of this anywhere
## [7][I tried to write clean architecture based back-end application in TypeScript](https://www.reddit.com/r/typescript/comments/j27rqe/i_tried_to_write_clean_architecture_based_backend/)
- url: https://github.com/pvarentsov/typescript-clean-architecture
---

## [8][Excluding the keys of a base type](https://www.reddit.com/r/typescript/comments/j2bggl/excluding_the_keys_of_a_base_type/)
- url: https://www.reddit.com/r/typescript/comments/j2bggl/excluding_the_keys_of_a_base_type/
---
I have a base type (ex: Record&lt;string, any&gt;) that I want to ensure is true for all sub-type instances. For each sub-type, I have a set of specific keys which will be valid which return a specific set of values.

For example: type MyType extends Record&lt;string, any&gt; = { a: string, b: number }

I am looking for a way to get the type "a" | "b" from MyType. However, using "keyof MyType" == "string | number".

What is the right way to extract only ("a" | "b") into a type of its own, given MyType as the input?

Playground:
https://www.typescriptlang.org/play?#code/C4TwDgpgBAksEFsDOUC8UBKEDGB7ATgCYA8Sw+AlgHYDmANFAIZUgB8A3FAPRdQCyEZsCjBcUMPlwA3CoWiMoAI0ZJooSFDL4ArtmDb8EBlVzCkjEFABEzS1uo0oFFLgDWFqwChq8fADNGbGg+EDhEFAgAD3gqQhQw5CgAb25eASERMQlpWXkqKCiwABsKbAphErIoXD8oKUYSwihXCBAUZibgAAsICnxNSDK-UqhDfXx89QgkTyg5pgAuTXIHOln5xSWqbQRFCHxPAF9PTymoAGlWpABBPW0GtGbWmv5Q+GROHjR0e1ooAB8oNtdvtTuBoJc2gARaZ9CBNdA2KwA6yKKyfXhyJBwwgnPzaKh6Ci4fI0CDABLEAAqBWiEFiKEhNzuDVYAApUlAAMJdZhkkRdZxQEpUNRiJkw7GGTpiQKCiBSaBYnFKCC8mQEAB06zm5XCSxCCSQDB1TkISypAEpPAa3uEANpUgC6yVNYwM+T1yHtsid7COJzwVCqXqQtqNjySpsYSyshEYwEYVjW8yUSwArAHPEGqvUitoINdHmSKe82aGGEjLRjlpRaJqoAB1AiuBzanPCPMFgBCxfJCXL72NqKs1c5wL2+Abzfwrfr2ZJuYaBa5fdLiEH4Ur2FHNdsUAAhEemxlROJJIRdPIROCCvhJP0eoYoH4CExPVQ87JrDuniBNUAA
## [9][What am I missing here?](https://www.reddit.com/r/typescript/comments/j27k9e/what_am_i_missing_here/)
- url: https://www.reddit.com/r/typescript/comments/j27k9e/what_am_i_missing_here/
---
[Playground](https://www.typescriptlang.org/play?#code/C4TwDgpgBAKgjFAvFA2gcgM7AE4EsB2A5mgDRQDeAtgQPwBcU+ArpQEYTZmUCGAHvYxbtsAXwC6AbgBQoSLABMSVGm7Zs3EKVgxJUmeGgwYSKVDNQAPrDinzVmPKnSpEXmAD22YFABmTfADGwLju+FAAFtwYUAA8MAB8UAAU3Azc+CBkANYMWRAg7j6wAJRpULjR5CgA+lnlYXkFRToMMDVZYiJS5LZm2BDATNgN9VDcUABkE2MoHVAAhIjI-gAmED4EECvSXVIEwBw+3AHQAFLuBDEAysDcwQEwZABK-sGUEAkUvWMBJ2DASWADBe+DeHzIADcGDc7rgHqUoKx3O4ADYQdI7PQBUJYKBPeAMc6XLB4IhkeCJZA9cw-P4A4CQ4pfGk0-qDYbJWQQQpQCFIJZQTA4AjEJlTb4s8xJeaRDCAlBwMRkNDUfBoJkWKzABViAB0qtiyAhurRRGA4TFEwlkuSMqi8sVyp4vHVli1Ov1fCg8SNJogZot0hpXV22PwuPx8kJF3wMX8WXw7gA7vgUEqFJTmeZjnTASQIUzqTaw7iMExWFylNrFUGbWyhmEkgBBNQaXUVFvqEBJAuW62S40QCEcbtMJCJfFtMsVgwoAAMYj1OYg-yS065ZCYxWKtbMIaxOO8k6URbMQtJxGB8BI3xUrc0V-kN5EQA)

I want to create a simple type validator for some data I get coming in.

This is a sub-example but it already doesn't work. What am I doing wrong?
## [10][I created a FULLY TYPED Twitter API Client for Node.js](https://www.reddit.com/r/typescript/comments/j1bidy/i_created_a_fully_typed_twitter_api_client_for/)
- url: https://www.reddit.com/r/typescript/comments/j1bidy/i_created_a_fully_typed_twitter_api_client_for/
---
I created a FULLY TYPED Twitter API Client for Node.js 🔧  
This client is taking full advantage of modern TypeScript capabilites!

✅ **Fully typed!** Both for query parameters and responses.  
✅ Includes 90% of the **official Twitter API** endpoints.  
✅ **Promise-based!** No ugly callbacks.  
✅ Inbuilt in-memory **cache** for rate-limit friendly usage.

Check it out 👇

[https://github.com/Silind/twitter-api-client](https://github.com/Silind/twitter-api-client)
## [11][Is it standard to use the not null operator everywhere Env variables are used?](https://www.reddit.com/r/typescript/comments/j1iuc4/is_it_standard_to_use_the_not_null_operator/)
- url: https://www.reddit.com/r/typescript/comments/j1iuc4/is_it_standard_to_use_the_not_null_operator/
---
It's by far my most common use of !. In strict mode all Env variables have an implicit type of `string | undefined`.

Just curious if you guys also use ! with them. Or if you do something else such as asserting `process.env` as `any`
