# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][Race conditions in React and beyond. A race condition guard with TypeScript](https://www.reddit.com/r/typescript/comments/fcafpl/race_conditions_in_react_and_beyond_a_race/)
- url: https://wanago.io/2020/03/02/race-conditions-in-react-and-beyond-a-race-condition-guard-with-typescript/
---

## [3][Just wanted to say that I can't even quantify the number of time Typescript saved my ass, but it's a lot.](https://www.reddit.com/r/typescript/comments/fbw9x9/just_wanted_to_say_that_i_cant_even_quantify_the/)
- url: /r/LearnTypescript/comments/fbv23u/just_wanted_to_say_that_i_cant_even_quantify_the/
---

## [4][[@known-as-bmf/store] Lightweight synchronous shared state library.](https://www.reddit.com/r/typescript/comments/fc0eo6/knownasbmfstore_lightweight_synchronous_shared/)
- url: https://github.com/known-as-bmf/store
---

## [5][Set Up a Typescript React Redux Project](https://www.reddit.com/r/typescript/comments/fb56ac/set_up_a_typescript_react_redux_project/)
- url: https://typeofnan.dev/setup-a-typescript-react-redux-project/
---

## [6][Rosebox a functional typescript library for styling](https://www.reddit.com/r/typescript/comments/fb0zse/rosebox_a_functional_typescript_library_for/)
- url: https://www.reddit.com/r/typescript/comments/fb0zse/rosebox_a_functional_typescript_library_for/
---
Hi guys!

The library is still very young and is under heavy development, but getting feedback during this phase is crucial in my opinion so I've decided to share it.

Website and docs: [https://www.rosebox.dev/](https://www.rosebox.dev/0.1.4)

Github: [https://github.com/hugonteifeh/rosebox](https://github.com/hugonteifeh/rosebox)

Progress  of next release: [https://github.com/hugonteifeh/rosebox/projects/1](https://github.com/hugonteifeh/rosebox/projects/1)
## [7][What editor/IDE do you use, and why?](https://www.reddit.com/r/typescript/comments/fbakwg/what_editoride_do_you_use_and_why/)
- url: https://www.reddit.com/r/typescript/comments/fbakwg/what_editoride_do_you_use_and_why/
---
For me, a mix between vscode &amp; phpstorm (but TS is better supported in vscode)
## [8][Fun practice using Generics to restrict function arguments based on other function arguments! Note the linter immediately telling me when I make missteps. One of my favorite parts of Typescript is that it helps you fail faster, enabling you to fix mistakes sooner.](https://www.reddit.com/r/typescript/comments/falwk7/fun_practice_using_generics_to_restrict_function/)
- url: https://i.redd.it/uiu34cki2kj41.png
---

## [9][Curveball — A typescript microframework with first class AWS Lambda and HTTP/2 Push support](https://www.reddit.com/r/typescript/comments/far19m/curveball_a_typescript_microframework_with_first/)
- url: https://medium.com/@evertp/curveball-a-typescript-microframework-with-first-class-aws-lambda-and-http-2-push-support-6efeba66ca45?source=friends_link&amp;sk=beff2d5d3e89693ecf3cc43a55a33e3f
---

## [10][Writing a types declaration file for a library where one of the exported functions is a function imported from another library. How should I write this?](https://www.reddit.com/r/typescript/comments/fawz1m/writing_a_types_declaration_file_for_a_library/)
- url: https://www.reddit.com/r/typescript/comments/fawz1m/writing_a_types_declaration_file_for_a_library/
---
Every approach I take, my IDE complains about it.

    // index.d.ts
    import foo from 'foo-library';
    
    declare module 'bar' {
      ...
      export = foo; // error ts(2666) "Exports and export assignments are not permitted in module augmentations"
      export function foo; // error ts(1005) (brackets expected)
      export const foo; // type info lost on foo
    }

Any ideas?
## [11][How to get started with gulp/browserify and npm](https://www.reddit.com/r/typescript/comments/fb2za9/how_to_get_started_with_gulpbrowserify_and_npm/)
- url: https://www.reddit.com/r/typescript/comments/fb2za9/how_to_get_started_with_gulpbrowserify_and_npm/
---
Hi,

I'm struggling with some very basic things like importing packages I installed with NPM. I honestly tried googling it a lot, and tried out different tutorials, but I can't get it to work. Here is my current setup, trying with gulp, I guess I'm missing something quite basic..

&amp;#x200B;

&amp;#x200B;

I installed (e.g. webmidi via npm) and my folder structure looks like this:

    proj
    +--node_modules
        +--module1
        +--module2
        +--webmidi
    +--dist
    +--src
        +--main.ts
        +--index.html

&amp;#x200B;

My Gulpfile looks like this:

    var gulp = require('gulp');
    var webmidi = require('webmidi');
    var browserify = require('browserify');
    var source = require('vinyl-source-stream');
    var tsify = require('tsify');
    var paths = {
        pages: ['src/*.html']
    };
    
    gulp.task('copy-html', function () {
        return gulp.src(paths.pages)
            .pipe(gulp.dest('dist'));
    });
    
    gulp.task('default', gulp.series(gulp.parallel('copy-html'), function () {
        return browserify({
            basedir: '.',
            debug: true,
            entries: ['src/main.ts'],
            cache: {},
            packageCache: {}
        })
        .plugin(tsify)
        .bundle()
        .pipe(source('bundle.js'))
        .pipe(gulp.dest('dist'));
    }));

&amp;#x200B;

main.ts:

    import * as webmidi from "webmidi";
    
    console.log(WebMidi);

package.json:

    {
      "name": "ts4",
      "version": "1.0.0",
      "description": "",
      "main": "./dist/main.js",
      "scripts": {
        "test": "echo \"Error: no test specified\" &amp;&amp; exit 1"
      },
      "author": "",
      "license": "ISC",
      "devDependencies": {
        "browserify": "^16.5.0",
        "gulp": "^4.0.0",
        "gulp-require-modules": "^1.1.4",
        "gulp-typescript": "^6.0.0-alpha.1",
        "tsify": "^4.0.1",
        "typescript": "^3.8.2",
        "vinyl-source-stream": "^2.0.0"
      },
      "dependencies": {
        "webmidi": "^2.5.1"
      }
    }

and tsconfig.json:

    {
        "files": [
            "src/main.ts"
        ],
        "compilerOptions": {
            "noImplicitAny": true,
            "target": "es5"
        }
    }

&amp;#x200B;

The error I get when i do

    &gt;gulp
    [23:36:55] Using gulpfile ~\Documents\Projekte\code\ts4\gulpfile.js
    [23:36:55] Starting 'default'...
    [23:36:55] Starting 'copy-html'...
    [23:36:55] Finished 'copy-html' after 39 ms
    [23:36:55] Starting '&lt;anonymous&gt;'...
    [23:36:57] '&lt;anonymous&gt;' errored after 2.21 s
    [23:36:57] Error: TypeScript error: c:/users/username/documents/projekte/code/ts4/src/main.ts(4,13): Error TS2304: Cannot find name 'WebMidi'.
        at formatError (C:\Users\username\AppData\Roaming\npm\node_modules\gulp-cli\lib\versioned\^4.0.0\format-error.js:21:10)
        at Gulp.&lt;anonymous&gt; (C:\Users\username\AppData\Roaming\npm\node_modules\gulp-cli\lib\versioned\^4.0.0\log\events.js:33:15)
        at emitOne (events.js:121:20)
        at Gulp.emit (events.js:211:7)
        at Object.error (C:\Users\username\Documents\Projekte\code\ts4\node_modules\undertaker\lib\helpers\createExtensions.js:61:10)
        at handler (C:\Users\username\Documents\Projekte\code\ts4\node_modules\now-and-later\lib\mapSeries.js:47:14)
        at f (C:\Users\username\Documents\Projekte\code\ts4\node_modules\once\once.js:25:25)
        at f (C:\Users\username\Documents\Projekte\code\ts4\node_modules\once\once.js:25:25)
        at tryCatch (C:\Users\username\Documents\Projekte\code\ts4\node_modules\async-done\index.js:24:15)
        at done (C:\Users\username\Documents\Projekte\code\ts4\node_modules\async-done\index.js:40:12)
    [23:36:57] 'default' errored after 2.26 s

According to  [https://github.com/djipco/webmidi/issues/82](https://github.com/djipco/webmidi/issues/82)  this is the right way to import WebMidi, and it should work - i guess I'm doing something wrong, can anyone help me?
