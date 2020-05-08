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
## [2][Convert parameters to destructured object](https://www.reddit.com/r/typescript/comments/gfre6t/convert_parameters_to_destructured_object/)
- url: https://www.reddit.com/r/typescript/comments/gfre6t/convert_parameters_to_destructured_object/
---
I just noticed this refactoring was available in one of my methods. Is this a refactor I should do? And what really does it do? Tried to read about it but I can't quite get it. What's the benefit of it?    

Here's the method:    

    private methodName(exampleString: string, array: any) {
        const json = JSON.parse(exampleString);
        json.map((value: any) =&gt; {
            array.map((v: { Type: any; Value: boolean; }) =&gt; {
                if (v.Type === value) {
                    v.Value = true;
                }
            });
        });
    }

If I use the refactoring it looks like this:    

    private methodName({ exampleString, array }: { exampleString: string; array: any; }) {
        const json = JSON.parse(exampleString);
        json.map((value: any) =&gt; {
            array.map((v: { Type: any; Value: boolean; }) =&gt; {
                if (v.Type === value) {
                    v.Value = true;
                }
            });
        });
    }

What's the benefit?
## [3][​​Setting Up Webpack and Typescript - TypeScript Webpack example](https://www.reddit.com/r/typescript/comments/gft88l/setting_up_webpack_and_typescript_typescript/)
- url: https://codesource.io/setting-up-webpack-and-typescript/
---

## [4][Front-end Software Engineer opening @Romania (Iasi)](https://www.reddit.com/r/typescript/comments/gfr66i/frontend_software_engineer_opening_romania_iasi/)
- url: https://www.reddit.com/r/typescript/comments/gfr66i/frontend_software_engineer_opening_romania_iasi/
---
Hi there! :)

I am currently looking for great Front-end Software Engineers for our new project in Iasi (Romania).

Our aim is to build Single Page Applications that drive decision making, with big care about writing high-quality code and providing excellent user experience in our product.

Key required skills: #JavaScript, #ReactJS, #TypeScript, #Redux &amp; familiarity with #CI/CD tools.

Want to discuss more? Just leave me a message! :)
## [5][Publishing TypeScript NPM Libraries using NX and Github Actions](https://www.reddit.com/r/typescript/comments/gfejco/publishing_typescript_npm_libraries_using_nx_and/)
- url: https://tane.dev/2020/05/publishing-npm-libraries-using-nx-and-github-actions/
---

## [6][Arithmetic(s) for RxJS observables](https://www.reddit.com/r/typescript/comments/gf85bi/arithmetics_for_rxjs_observables/)
- url: https://loreanvictor.github.io/rxmetics
---

## [7][Am I typing so much?](https://www.reddit.com/r/typescript/comments/gfbye8/am_i_typing_so_much/)
- url: https://www.reddit.com/r/typescript/comments/gfbye8/am_i_typing_so_much/
---
I'm using TS in a React project for the *first time*, and I have some questions about how much should I type. Let's say I want to handle an input state in a function component.

    const SomeComp = (): JSX.Element =&gt; {
      const [inputValue1, setInputValue1] = useState('');
      const [inputValue2, setInputValue2] = useState('');
    
      const handleInputs = () =&gt; {} // a function to handle inputs, see below.
    
      return (
        &lt;div&gt;
          &lt;input value={inputValue1} onChange={handleInputs(setInputValue1)} /&gt;
          &lt;input value={inputValue2} onChange={handleInputs(setInputValue2)} /&gt;
      &lt;/div&gt;
    }

**And we have this function:**

    const handleInputs = (setter: React.Dispatch&lt;React.SetStateAction&lt;string&gt;&gt;) =&gt; 
      (e: React.ChangeEvent&lt;HTMLInputElement&gt;): void =&gt; {
        setter(e.currentTarget.value);
    };

**Vs. this one:**

    const handleInputs = (setter) =&gt; (e): void =&gt; {
      setter(e.currentTarget.value);
    };

The first one is a little bit to read, but it ensures that it accepts a `setState` function and I can safely use the input event. But the second one is really easy to understand but, you know, seems typing dangerous. I've seen a few TS + React projects, they used the second one.

What do you think I should use? How do you handle this kind of typing?
## [8][Nest or FoalTS?](https://www.reddit.com/r/typescript/comments/gfd8hf/nest_or_foalts/)
- url: https://www.reddit.com/r/typescript/comments/gfd8hf/nest_or_foalts/
---
Deciding between these 2 serverside frameworks, nest is looking more mature and stable, but foal looks easier to pick up. I want to use it to create GraphQL API with Postgres.

If someone had experience with these 2, would be interesting to hear your opinions.
## [9][TS.ED Great Node.JS Typescript framework](https://www.reddit.com/r/typescript/comments/gfg7km/tsed_great_nodejs_typescript_framework/)
- url: https://tsed.io/
---

## [10][‘’ is sometimes undefined](https://www.reddit.com/r/typescript/comments/gfb4zd/is_sometimes_undefined/)
- url: https://www.reddit.com/r/typescript/comments/gfb4zd/is_sometimes_undefined/
---
Working on a TS project where a build process replaces a const string, which is otherwise ‘’. The odd part is that it complains about a .replace call later on in the app, saying that replace does not exist on never. Sure enough putting 

const test = ‘’
console.log(typeof test)

Prints ‘undefined’

It’s very inconsistent - so I’m not sure how to reproduce but it’s absolutely occurring in our production code which I cannot share. Any ideas?
## [11][[TS.ED] Autoload services](https://www.reddit.com/r/typescript/comments/gfg56p/tsed_autoload_services/)
- url: https://tsed.io/docs/services.html#configuration
---

