# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][Doorbot.ts – Let me in the building, with Typescript](https://www.reddit.com/r/typescript/comments/f3hggg/doorbotts_let_me_in_the_building_with_typescript/)
- url: http://www.wumpus-cave.net/2020/02/13/doorbot-ts-let-me-in-the-building-with-typescript/
---

## [3][extract discriminants into type](https://www.reddit.com/r/typescript/comments/f3awxt/extract_discriminants_into_type/)
- url: https://www.reddit.com/r/typescript/comments/f3awxt/extract_discriminants_into_type/
---
Hi, i need a new type that is a union of all the disciminants in a discriminated uniton.

    interface Square {
        kind: "square";
        size: number;
    }
    
    interface Rectangle {
        kind: "rectangle";
        width: number;
        height: number;
    }
    
    interface Circle {
        kind: "circle";
        radius: number;
    }
    
    type Shape = Square | Rectangle | Circle;
    
    type KindList = SomeMagic&lt;Shape&gt; // "square" | "rectangle" | "circle"

can someone help me?
## [4][Just blogged: TypeScript types cheat sheet](https://www.reddit.com/r/typescript/comments/f33on2/just_blogged_typescript_types_cheat_sheet/)
- url: https://medium.com/@maximzhukov_dev/typescript-types-cheat-sheet-5da7c6ee083d
---

## [5][Unit tests for d.ts files in JS project](https://www.reddit.com/r/typescript/comments/f2wbaf/unit_tests_for_dts_files_in_js_project/)
- url: https://github.com/ai/check-dts
---

## [6][Angular 9: What’s New?](https://www.reddit.com/r/typescript/comments/f2tsre/angular_9_whats_new/)
- url: https://auth0.com/blog/angular-9-whats-new/?utm_source=twitter&amp;utm_medium=sc&amp;utm_campaign=angular_9
---

## [7][an attribute type which knows the type of another attribute within the same interface](https://www.reddit.com/r/typescript/comments/f2xt9o/an_attribute_type_which_knows_the_type_of_another/)
- url: https://www.reddit.com/r/typescript/comments/f2xt9o/an_attribute_type_which_knows_the_type_of_another/
---
I have a type:

    type Schema &lt;M extends Record&lt;string, {}&gt;&gt; = {
      [E in keyof M]: {
        [L in keyof Partial&lt;M[E]&gt;]: {
          entity: keyof M,
          reciprocal: string,
        }
      }
    }

... whose usage would look like:

    interface M extends Record&lt;string, object&gt; {
      user: { articleIds: string[] },
      article: { authorId: string, editorId: string },
    }
    
    const mySchema: Schema&lt;M&gt; = {
      user: {
        articleIds: {
          entity: 'article',
          reciprocal: 'authorId'
        }
      },
      article: {
        authorId: {
          entity: 'user',
          reciprocal: 'articleIds'
        }
      }
    };

The schema's purpose is to define data relationships through their attributes.

An `entity` must be any key of type `M`. I am trying to constrain `reciprocal` to only keys within the value of that given key of type `M`. To illustrate this problem:

    user: {
      articleIds: {
        entity: 'article',
        reciprocal: 'this value can as any string' // how to constrain to 'authorId'|'editorId' ?  
      }
    },

Looking back at type `Schema`, the type of `reciprocal` would have to know concretely what `keyof M` is:

    type Schema &lt;M extends Record&lt;string, {}&gt;&gt; = {
      [E in keyof M]: {
        [L in keyof Partial&lt;M[E]&gt;]: {
          entity: keyof M,
          reciprocal: string, // how to reference the concrete value of `entity`?
        }
      }
    }

Is there a way to do this?
## [8][Typeclasses in Typescript with fp-ts](https://www.reddit.com/r/typescript/comments/f2kokc/typeclasses_in_typescript_with_fpts/)
- url: https://paulgray.net/typeclasses-in-typescript/
---

## [9][Axios requests to custom API in TS](https://www.reddit.com/r/typescript/comments/f2tbcb/axios_requests_to_custom_api_in_ts/)
- url: https://www.reddit.com/r/typescript/comments/f2tbcb/axios_requests_to_custom_api_in_ts/
---
Hey guys,

I've created my own API in TS (fairly new into this ngl)  and I'd like to know how to pass data over requests with axios. Assuming I have a form which takes 1 email and 1 password, when I press a button, it sends this request :

```
const login = async (user: any) =&gt; {
  const data = {
    email: user.email,
    password: user.password
  }
  const res = await axios.get('/users/login', { data })

  console.log(res)

  return res;
}
```

User is basically an object containing two fields, email and password

Everything is working well with cURL requests but I have some issues with requests from my app. My API gets the data from request body.
My API login function is :

```
public async standardLogin(req: Request, res: Response): Promise&lt;Response&gt; {
    try {
      const { email, password } = req.body

      if (!email || !password)
        return super.sendMissingFieldsError(res)

      const user = await auth.signInWithEmailAndPassword(email, password)

      return super.sendSuccessMessage(res, 'Success')
    } catch (e) {
      return super.sendErrorMessage(res, e)
    }
  }
```


Anyone can help me through this? Thanks :D
## [10][Foal TS - February release (version 1.6) - Local and Cloud file storage (AWS S3), enhanced dependency injection, cleaner builds](https://www.reddit.com/r/typescript/comments/f2ubou/foal_ts_february_release_version_16_local_and/)
- url: https://www.reddit.com/r/typescript/comments/f2ubou/foal_ts_february_release_version_16_local_and/
---
Foal TS framework version 1.6 is officially released!

Local and Cloud file storage (AWS S3), enhanced dependency injection, cleaner builds, bug fixes, you can discover the features of this version here: [https://github.com/FoalTS/foal/releases/tag/v1.6.0](https://github.com/FoalTS/foal/releases/tag/v1.6.0)

The big of feature of this release is FoalTS new file system which allows you to easily choose different storage for each of your environments. For example, when coding the application locally, you can use the file system of your operating system. Then, when deploying the application to staging or production, you can change the configuration to use AWS S3. Regardless of the storage chosen in the background, the code remains the same. Only the configuration changes.

More documentation here: [https://foalts.gitbook.io/docs/topic-guides/file-system/local-and-cloud-storage](https://foalts.gitbook.io/docs/topic-guides/file-system/local-and-cloud-storage)

&amp;#x200B;

[Foal TS - February release \(version 1.6\)](https://preview.redd.it/8lp1wb3dyig41.png?width=1122&amp;format=png&amp;auto=webp&amp;s=60c15ac855b175b99e7bc0307a7f4a77ad72850b)

Foal, in a few words, it's a Node.js framework:

* written in TypeScript
* provided with batteries included (Auth, OpenAPI, GraphQL, ORM, CLI, scripts, file storage)
* and with a simple and intuitive architecture (no magic, no over-engineering).

And the must: it has more than 11,000 lines of documentation.

[https://foalts.org](https://foalts.org/)

[https://github.com/FoalTS/foal](https://github.com/FoalTS/foal)

[https://foalts.gitbook.io/docs/](https://foalts.gitbook.io/docs/)

# TypeScript #JavaScript #NodeJS #FoalTS #AWS #S3
## [11][PeerTube v2.1 released! (PeerTube is an Open Source &amp; Decentralized YouTube Alternative)](https://www.reddit.com/r/typescript/comments/f2a539/peertube_v21_released_peertube_is_an_open_source/)
- url: https://github.com/Chocobozzz/PeerTube/releases/tag/v2.1.0
---

