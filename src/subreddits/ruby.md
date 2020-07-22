# ruby
## [1][Ruby Method Overloading](https://www.reddit.com/r/ruby/comments/hvpwll/ruby_method_overloading/)
- url: https://lucaguidi.com/2020/07/22/ruby-method-overloading/
---

## [2][How to build Ruby in Windows natively without WLS, MSYS2 or Cygwin](https://www.reddit.com/r/ruby/comments/hvpgqu/how_to_build_ruby_in_windows_natively_without_wls/)
- url: http://gromnitsky.blogspot.com/2020/07/how-to-build-ruby-in-windows-natively.html
---

## [3][A Deeper Look at Ruby's Enumerable](https://www.reddit.com/r/ruby/comments/hvr9xe/a_deeper_look_at_rubys_enumerable/)
- url: https://tutswiki.com/ruby-enumerable/
---

## [4][E-book: Mastering Roda](https://www.reddit.com/r/ruby/comments/hvjj6j/ebook_mastering_roda/)
- url: https://fiachetti.gitlab.io/mastering-roda/
---

## [5][Ractor - Ruby's Actor-like concurrent abstraction](https://www.reddit.com/r/ruby/comments/hvpwe0/ractor_rubys_actorlike_concurrent_abstraction/)
- url: https://github.com/ko1/ruby/blob/ractor_parallel/doc/ractor.md
---

## [6][url redirection vulnerability explained](https://www.reddit.com/r/ruby/comments/hvrjzg/url_redirection_vulnerability_explained/)
- url: https://www.reddit.com/r/ruby/comments/hvrjzg/url_redirection_vulnerability_explained/
---
Today's  learning url redirection issue,  High pay cyber security 100 unique web  security issues learning easily in just 20 hours for cyber security  jobs.

Search in YouTube for "uday datrak 100 bug bounty lessons" [https://www.youtube.com/watch?v=rFKoHlJ\_SG0](https://www.youtube.com/watch?v=rFKoHlJ_SG0)
## [7][Rollout UI - Admin Dashboard for Feature Flags](https://www.reddit.com/r/ruby/comments/hvsip1/rollout_ui_admin_dashboard_for_feature_flags/)
- url: https://github.com/fetlife/rollout-ui
---

## [8][Need some help translating Python if possible.](https://www.reddit.com/r/ruby/comments/hvosx1/need_some_help_translating_python_if_possible/)
- url: https://www.reddit.com/r/ruby/comments/hvosx1/need_some_help_translating_python_if_possible/
---
It's for a leetcode problem I'm studying up on. I'm having a hard time coming up with Ruby for this (my syntax is God awful as a self-taught rookie), but I found some Python code that may just be exactly what I need. Assuming there's such a thing as "stack" in Ruby that is. Otherwise, any fix you may have would be very helpful. Here's the challenge:

&gt;The match method takes in a string, and the index of an opening parentheses. The method needs to return the index of the matching parentehses. If the given index is not to an open parentheses, or there is no matching parentheses, return -1. See examples below.

Example:

ParenthesesMatcher.match("(())", 1) =&gt; 2   
The given index is for the second parentheses.   
The matching is the closing parentheses at index 2.

`def match_string(s, index):`  
`stack = []`  
`marker = -1`  
`for i in range(len(s)):`  
`if i != index:`  
`if s[i] == '(':`  
`stack.append(i)`  
`elsif s[i] == ')':`  
`if len(stack) == 0:`  
`continue`  
`peek = stack[-1]`  
`if marker &gt;=0 and peek == marker:`  
`return i`  
`else:`   
`stack.pop()`  
`else:`  
`if s[i] == ')':`  
`return stack[-1] if len(stack) &gt; 0 else -1`  
`elsif s[i] == '(':`  
`marker = i`  
`stack.append(i)`  
`else:`   
`return -1`  
`return -1`
## [9][Solution for those who like Action Cable but hate writing JavaScript](https://www.reddit.com/r/ruby/comments/hv6uk2/solution_for_those_who_like_action_cable_but_hate/)
- url: https://pdabrowski.com/articles/cable-ready-with-action-cable
---

## [10][Ractor updates and English documentation](https://www.reddit.com/r/ruby/comments/hva15c/ractor_updates_and_english_documentation/)
- url: https://www.reddit.com/r/ruby/comments/hva15c/ractor_updates_and_english_documentation/
---
[https://github.com/ko1/ruby/commit/c543e3dc230c82eca9bac40a7d43cba4e4ec206d](https://github.com/ko1/ruby/commit/c543e3dc230c82eca9bac40a7d43cba4e4ec206d)
