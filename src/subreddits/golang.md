# golang
## [1][Go server for Apple/Google contact tracing framework](https://www.reddit.com/r/golang/comments/g6xsxu/go_server_for_applegoogle_contact_tracing/)
- url: https://github.com/dstotijn/ct-diag-server
---

## [2]['For Loops' To 'WorkerPools' in Four Steps (Go Serverless)](https://www.reddit.com/r/golang/comments/g77zoj/for_loops_to_workerpools_in_four_steps_go/)
- url: https://www.youtube.com/watch?v=z_ZGS0C-qto
---

## [3][Reducing debugging time while programming. Use static and stack-trace analysis to determine which func call causes that error.](https://www.reddit.com/r/golang/comments/g75nos/reducing_debugging_time_while_programming_use/)
- url: https://github.com/snwfdhmp/errlog
---

## [4][Service API version vs build version](https://www.reddit.com/r/golang/comments/g769tu/service_api_version_vs_build_version/)
- url: https://www.reddit.com/r/golang/comments/g769tu/service_api_version_vs_build_version/
---
When you build services that have an API, do you ever separate your API version from the service build version?

I think that an API version is very important for any consumer because that version is part of the API contract. I usually manually set the API version to a whole number and increment the value when I make backward incompatible changes.

The service build number/version doesn't have anything to do with the API contract and instead deals with the implementation of the service itself. For example, say you use version 3 of my API and find a bug. When the bug is fixed, the API version is still 3, but the service build version/number changes. This number is typically set automatically at build time to either the git commit hash or to the git tag name (if one exists).

That's my current approach, but I wanted to see what other people think of this and if there are any documented patterns for doing this or for doing something better?
## [5][A comparison of three programming languages for a full-fledged next-generation sequencing tool | BMC Bioinformatics](https://www.reddit.com/r/golang/comments/g71mzs/a_comparison_of_three_programming_languages_for_a/)
- url: https://bmcbioinformatics.biomedcentral.com/articles/10.1186/s12859-019-2903-5
---

## [6][How to use grpc with vue](https://www.reddit.com/r/golang/comments/g784au/how_to_use_grpc_with_vue/)
- url: https://www.reddit.com/r/golang/comments/g784au/how_to_use_grpc_with_vue/
---
I want to use grpc with vue how can i use it. I am using go lang as the backend launguage. Stuck in the use of envoy or nginx dilemma is there is any easy way to do such thing or should i learn envoy or nginx from beginning. If so how can i learn those things. If there is any other way to create grpc microservices ? The main idea is to use http2 as http endpoint. For eg localhost:\\something\
## [7][go2dotnet - A converter from Go to .NET (C#)](https://www.reddit.com/r/golang/comments/g77n0y/go2dotnet_a_converter_from_go_to_net_c/)
- url: https://github.com/hajimehoshi/go2dotnet
---

## [8][A programming language based on IKEA product's names](https://www.reddit.com/r/golang/comments/g77e6f/a_programming_language_based_on_ikea_products/)
- url: https://github.com/hugolgst/ikea-sharp
---

## [9][is a compiled Go binary truly free from external dependencies?](https://www.reddit.com/r/golang/comments/g6y43r/is_a_compiled_go_binary_truly_free_from_external/)
- url: https://www.reddit.com/r/golang/comments/g6y43r/is_a_compiled_go_binary_truly_free_from_external/
---
I am coming from a Python &amp; R background, and I am sick of having to tote around practically an entire operating system just to ensure that my scripts and programs' dependencies are available. I am very interested in alternatives that provide static binaries that are free from external dependencies. So that instead of going through the process of

- git clone repo

- run install script to setup virtualenv / conda env with dependencies

- (manually) activate the env every. single. time. I want to use said program

I can just 

- download compiled program binary

- put it in PATH

- now its always available for use

There are plenty of other things that I hate about Python, such as the duck typing, but today the environment management is the one that is bugging me the most. I work with on-premises HPC servers building and managing the infrastructure that orchestrates data analysis workflows, among other things. We do a lot of Django and Flask API's, along with CLI tools to access them easier. If you aren't familiar with HPC, its basically a giant bare-metal server owned by the company with hundreds of employees using it at once; so environment management is a constant nightmare that I would really like to put to bed, and Python is not helping. 

I am trying to build a case for, and explore, alternatives that might ease this burden, but I want to be sure that if I compile my Go program targeted for our CentOS server, then stick the binary on GitHub for coworkers to download, its really truly gonna work without a hitch. Is that the case? It seems too good to be true.
## [10][Lockgate is a cross-platform locking library with distributed locks using Kubernetes and OS file locks support](https://www.reddit.com/r/golang/comments/g6lypt/lockgate_is_a_crossplatform_locking_library_with/)
- url: https://github.com/flant/lockgate
---

