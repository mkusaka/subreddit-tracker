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
## [2][Rome: an experimental JavaScript toolchain from Facebook. It includes a compiler, linter, formatter, bundler, testing framework and more...](https://www.reddit.com/r/typescript/comments/fa83wh/rome_an_experimental_javascript_toolchain_from/)
- url: https://github.com/facebookexperimental/rome
---

## [3][You'd be surprised how much you can do without JS (or TS) :D](https://www.reddit.com/r/typescript/comments/fa7xlt/youd_be_surprised_how_much_you_can_do_without_js/)
- url: https://github.com/you-dont-need/You-Dont-Need-JavaScript
---

## [4][[Angular] How do I properly use requestAnimationFrame in an angular component ?](https://www.reddit.com/r/typescript/comments/fac1ba/angular_how_do_i_properly_use/)
- url: https://www.reddit.com/r/typescript/comments/fac1ba/angular_how_do_i_properly_use/
---
This is my attempt :  


      ngAfterViewInit(){
        // get the context
        const canvasEl: HTMLCanvasElement = this.canvas.nativeElement;
        this.ctx = canvasEl.getContext('2d');
    
        // set the width and height
        canvasEl.width = this.width;
        canvasEl.height = this.height;
    
    
    
        this.animate();
     }
     
     
      animate() {  
        window.requestAnimationFrame(this.animate);
      }

But I am getting this error in console : 

*Processing img 6vj6524kwgj41...*

&amp;#x200B;

  
So how do I create an animation loop in an angular component ?  
I know that setTimeout/setInterval works, however I've heard they're inefficient.
## [5][Drash's addendum to Deno's deps.ts file solution](https://www.reddit.com/r/typescript/comments/fabgik/drashs_addendum_to_denos_depsts_file_solution/)
- url: https://www.reddit.com/r/typescript/comments/fabgik/drashs_addendum_to_denos_depsts_file_solution/
---
One question I see come up multiple times in the Deno gitter revolves around dependency management. When I first read about using a deps.ts file in my Deno projects, I was still confused as to why I had to use one. I needed more of an explanation other than the statements on Deno's website.

Drash has an addendum to Deno's deps.ts file solution and it gives me the "why" that I needed when I first started. Hopefully it helps someone here understand the deps.ts file better as well.

[https://drash.land/docs/#/dependency-management](https://drash.land/docs/#/dependency-management)
## [6][ImTool: TypeScript canvas-based image manipulation library that will save you money - no need to waste your server CPU time generating thumbnails anymore](https://www.reddit.com/r/typescript/comments/f9tsse/imtool_typescript_canvasbased_image_manipulation/)
- url: https://github.com/mat-sz/imtool
---

## [7][Type-based formal verification](https://www.reddit.com/r/typescript/comments/fa2n0x/typebased_formal_verification/)
- url: https://youtu.be/JboZel47XU0
---

## [8][Forget NodeJS! Build native TypeScript applications with Deno 🦖](https://www.reddit.com/r/typescript/comments/f9fayv/forget_nodejs_build_native_typescript/)
- url: https://deepu.tech/deno-runtime-for-typescript/
---

## [9][tsParticles - Resurrecting particles.js](https://www.reddit.com/r/typescript/comments/f9p0gf/tsparticles_resurrecting_particlesjs/)
- url: https://www.reddit.com/r/typescript/comments/f9p0gf/tsparticles_resurrecting_particlesjs/
---
tsparticles - Resurrecting particles.js

Hi all,

I'm resurrecting the old particles.js, updated and converted to typescript. I've implemented some new features and fixes requested by users on the old project.

[https://www.npmjs.com/package/tsparticles](https://www.npmjs.com/package/tsparticles)
## [10][Modern Typescript with Examples Cheat Sheet](https://www.reddit.com/r/typescript/comments/f9es9d/modern_typescript_with_examples_cheat_sheet/)
- url: https://www.reddit.com/r/typescript/comments/f9es9d/modern_typescript_with_examples_cheat_sheet/
---
I was looking for something like this but could find nothing that fitted. The goal is to be a reference when looking things up with great examples that only demonstrate the syntax and problem/solution, not a load of other stuff which tends to be the case in blog examples.

Once I am happy with it I will create an attractive formatted PDF to print and have that as a link too.

I have taken examples from mainly the official docs and various blog posts by some great authors including quite a bit of Marius Schulz.

The other goal is to try and include only what most people would need, and omit any old ways of doing things, so it is quite opinionated like that. 

In my personal TS journey this is all the stuff I have needed and wanted a reference to come back to, especially with easy to understand examples. Obvious things are left out:

[https://github.com/David-Else/modern-typescript-with-examples-cheat-sheet](https://github.com/David-Else/modern-typescript-with-examples-cheat-sheet)

Feedback and pull requests very welcome, especially factual corrections!
## [11][Types for PlantUML Parser: In this post I demonstrate how to parse PlantUML with TypeScript in a few simple steps.](https://www.reddit.com/r/typescript/comments/f9c4hy/types_for_plantuml_parser_in_this_post_i/)
- url: https://duckpond.ch/plantuml-parser/javascript/2020/02/25/types-for-plantuml-parser.html
---

