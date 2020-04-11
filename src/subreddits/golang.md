# golang
## [1][Error handling within a goroutine](https://www.reddit.com/r/golang/comments/fz27nl/error_handling_within_a_goroutine/)
- url: https://www.reddit.com/r/golang/comments/fz27nl/error_handling_within_a_goroutine/
---
I have a simple goroutine that acts as a poller and executes a function every x minutes based on an interval passed in.

I currently create my error channel in the factory method of the poller and then if there's an error pass the error into the channel and then return the error channel from the function if the goroutine is stopped.

    func (p *Poller) Start() &lt;-chan error {
    	go func() {
    		logrus.Info("polling initialized. Status: Running")
    		defer close(p.Errs)
    		for {
    			select {
    			case &lt;-p.interval.C:
    				err := p.HandlerFunc()
    				if err != nil {
    					p.Errs &lt;- err
    					return
    				}
    			case &lt;-p.done:
    				logrus.Info("polling shutting down. Status: stopped.")
    				return
    			}
    		}
    	}()
    
    	return p.Errs
    }

I then just check for errors when I run the poller

    errChan := poller.Start()
for err := range errChan {
 if err != nil {
 logrus.Fatalf("fatal error: %s", err)
   }
}

Is there a more 'go like' way to handle the error within a goroutine? I'm not 100% sure this is the correct approach. 

Any advice would be greatly appreciated!
## [2][Cuelang v0.1.0 is out.](https://www.reddit.com/r/golang/comments/fykhox/cuelang_v010_is_out/)
- url: https://www.reddit.com/r/golang/comments/fykhox/cuelang_v010_is_out/
---
[https://github.com/cuelang/cue/releases/tag/v0.1.0](https://github.com/cuelang/cue/releases/tag/v0.1.0)

Did you know Cue is the younger sibling of Go? Lots of go sources in Cue and being built by a member of the Go team.  


https://cuelang.org

&amp;#x200B;
## [3][iso9660 v0.1.2 - generate iso images on the fly and stream over http or write to file](https://www.reddit.com/r/golang/comments/fz1b4i/iso9660_v012_generate_iso_images_on_the_fly_and/)
- url: https://www.reddit.com/r/golang/comments/fz1b4i/iso9660_v012_generate_iso_images_on_the_fly_and/
---
https://github.com/KarpelesLab/iso9660

This was forked from https://github.com/kdomanski/iso9660 and modified as to allow serving dynamic ISO images over http.

Changes:

* do not use a staging directory, but instead store state in memory and do not copy data from files until the time comes
* Expose primary descriptor structure in writer.Primary for easy modification (volume name, date, etc)
* write to io.Writer so output can be streamed (to a http.ResponseWriter for example)
* support for El Torito bootable images (EFI not tested. I'll implement isohybrid bootable images eventually if there's some interest)
* Removed bugs, probably added some others. Bug reports are welcome.
## [4][Configuro - an opinionated configuration loading and validation framework focused towards Containerized and 12-Factor compliant applications.](https://www.reddit.com/r/golang/comments/fynphn/configuro_an_opinionated_configuration_loading/)
- url: https://www.reddit.com/r/golang/comments/fynphn/configuro_an_opinionated_configuration_loading/
---
I always have a similar structure to loading configurations and validating them on all my projects, I exported them inside a package and created `configuro`.

It provides the set of features you would want to implement 12-Factor compliant config and be container ready. **It has one defined strategy to load configurations** while only exposing a few options for simplicity.

Tell me what do you think.  
[https://github.com/sherifabdlnaby/configuro](https://github.com/sherifabdlnaby/configuro)
## [5][How to properly communicate between goroutines to start, restart or stop them?](https://www.reddit.com/r/golang/comments/fz3vvw/how_to_properly_communicate_between_goroutines_to/)
- url: https://www.reddit.com/r/golang/comments/fz3vvw/how_to_properly_communicate_between_goroutines_to/
---
So I want to write an API for a sort of worker pool which would look something like this, if I user tries to use my API in their go code

    wp := NewWorkerPool()
    tasks := []task{...}
    // start the background task in goroutines
    wp.StartWorkers()
    // StartWorkers is non blocking, you can go about your business after starting it
    // you can also stop all the workers before they finish
    wp.StopWorkers()

The worker pool starts a bunch of unix processes like sleep or ls or some curl request, so the tasks are actually names of commands and arguments.

I want the StartWorkers to be a non-blocking call itself. I also want the following features.

    All the processes are started as background goroutines
    In case of a crash or failure they are restarted
    They can also be stopped midway before they complete

Now I want my workers to communicate their state back to the worker pool so that it can take necessary actions based on that.

This is what I wrote pool.go


    type Pool struct {
      // chan for a successfull completion signal
      completed chan struct{}
      // chan for a crash signal
      crashed chan *worker
       // wait group to keep a number of running processes
       procCounter *sync.WaitGroup
       // a state of all child processes
       children []*worker
   }

       func New(processes []*Process) *Pool {
        var childProcesses []*worker
          for _, proc := range processes {
           childProcesses = append(childProcesses, newWorker(proc))
        }

      return &amp;Pool{
        completed:   make(chan struct{}),
        crashed:     make(chan *worker),
        procCounter: &amp;sync.WaitGroup{},
        children:    childProcesses,
      }
   }

    func (p *Pool) StartWorkers() {
      for _, child := range p.children {
        p.procCounter.Add(1)
        childWorker := child
        childWorker.start(p.crashed, p.completed)
      }

    // listen for crashes/completions of all running processes
      go func() {
        for {
            select {
            case &lt;-p.completed:
                p.procCounter.Done()
            case worker := &lt;-p.crashed:
             // in case we get a crashed signal we start the processes again
                p.restartProcess(worker)
            }
        }
      }()

      // putting the wait call in a goroutine since I want `StopProcess` to be non-blocking at its call site
      go func() {
        p.procCounter.Wait()
        close(p.completed)
        close(p.crashed)
       }()
    }

     // StopWorkers can be used to stop all the child processes.
    func (p *Supervisor) StopWorker() {
      for _, child := range p.children {
        p.StopProcess(child)
      }
    }

    // restart worker in case of a crash
    func (p *Supervisor) restartProcess(proc *worker) {
      proc.start(p.crashed, p.completed)
    }

This is worker.go

    type worker struct {
       status  string
      command *exec.Cmd
    }

    func newWorker(proc *Process) *worker {
        cmd := exec.Command(proc.Executable, proc.Args...)
        cmd.Env = proc.Env
        return &amp;worker{
            status:  "ready",
            command: cmd,
        }
    }

    func (w *worker) start(crashed chan *worker, completed chan struct{}) chan *worker {
        w.status = "running"
        err := w.command.Run()
        if err != nil {
            fmt.Println(err)
            // signal a crash
            w.status = "crashed"
            crashed &lt;- w
            return
        }
        // signal successful completion
        completed &lt;- struct{}{}
       return restart
    }

    func (w *worker) stop(crashed chan *worker, completed chan struct{}) {
     // stop the process
     w.status = "stopped"
     err := w.command.Process.Kill()
     if err != nil {
        fmt.Println("couldn't kill process")
     }
     completed &lt;- struct{}{}
    }

But this code deadlocks on running it, someone also mentioned to me that it is an anti pattern to have multiple writers to the same channel.

Is there a good way to do this? I want my worker pool to act as a monitor or a supervisor to start all the processes in the background and restart them if they crash or stop them if the user wants to kill them before they finish

Another thing I am struggling with is that I can't figure out a way to exit from this infinite for when all my processes have completed and none are running, not sure what case statement to use here

    go func() {
        for {
            select {
            case &lt;-p.completed:
                p.procCounter.Done()
            case worker := &lt;-p.crashed:
             // in case we get a crashed signal we start the processes again
                p.restartProcess(worker)
            }
        }
        }()
## [6][Package API recommendation for almost-duplicate methods.](https://www.reddit.com/r/golang/comments/fyzkhj/package_api_recommendation_for_almostduplicate/)
- url: https://www.reddit.com/r/golang/comments/fyzkhj/package_api_recommendation_for_almostduplicate/
---
Title probably doesn't make much sense, but I'll try to explains as best as I can. I am working on an experiment with a golang package for ssh based IT automation like Ansible, but no yaml ([https://github.com/krilor/gossh](https://github.com/krilor/gossh)). I'm hoping to get some input on the API.

A core part of the problem I'm working on is the ability to do "stuff" on remote (or local) targets, either as the logged in user or as another user - aka sudo. A target interface would maybe have 30 methods. I've been thinking about three ways to go about the interface for the target.

What do you think about these? Do you see any problems or benefits with either? Do you know any good examples where any of these patterns have been used (or articles discussing it)? Do I have other options? What would you do?

I've outlined the patterns below, with just a subset of methods.

**1 - Add user as a param to all methods**

```go
type Target interface {
  Create(user, path string) (io.WriteCloser, error)
  Run(user, cmd string) Response, error
  Mkdir(user, path string) error
  ...
}
```

**2 - Add duplicate methods with the user param to all methods**

The non-*As methods will just be calls to *As, with the current user specified.

```go
type Target interface {
  Create(path string) (io.WriteCloser, error)
  CreateAs(user, path string) (io.WriteCloser, error)
  Run(cmd string) Response, error
  RunAs(user, cmd string) Response, error
  Mkdir(path string) error
  MkdirAs(user, path string) error
  ...
}
```
**3 - Add a separate method for getting a wrapped target with the new user.

Usage would be like `f, err := t.As("anotheruser").Create("/some/path").

```go
type Target interface {
  As(user string) Target

  Create(path string) (io.WriteCloser, error)
  Run(cmd string) Response, error
  Mkdir(path string) error
  ...
}
```
## [7][CLI for ordering Dominos pizza](https://www.reddit.com/r/golang/comments/fy2qya/cli_for_ordering_dominos_pizza/)
- url: https://www.reddit.com/r/golang/comments/fy2qya/cli_for_ordering_dominos_pizza/
---
[https://github.com/harrybrwn/apizza](https://github.com/harrybrwn/apizza)

My hobby project is finally at a place where I'm comfortable sharing it with the world. Tell me what you think.
## [8][Is CGO Inevitable?](https://www.reddit.com/r/golang/comments/fynqlk/is_cgo_inevitable/)
- url: https://www.reddit.com/r/golang/comments/fynqlk/is_cgo_inevitable/
---
The Go programming language is awesome for many reasons including GC, fast compile, readability, etc.  I've wanted to start using Go for everything which eventually means graphics. However, any amount of research will expose the lack of 100% Go libraries. All there is, are bindings for GLFW, OpenGL, SMFL, Qt, etc all using CGO. Even libraries that tout a "100% Go" experience still contain CGO deep inside to interface with the os. It is well known that writing CGO is a downgrade from Go code as expressed [here](https://dave.cheney.net/2016/01/18/cgo-is-not-go). I know CGO has it's purpose: it can save tons of time from not having to write ports or to quickly interface with c libraries. But I want to take the position of a "pure" Go coder and only use Go.

My question is: Can a purely 100% Go graphics API be written?

What answers I'm looking for:

* Yes/No
* By What Means
* Further readings

What I'm NOT looking for:

* If it's hard or not
* Benefits of CGO 
* Why I shouldn't care about the use of CGO
## [9][Protect CRUD API against third-party app creation](https://www.reddit.com/r/golang/comments/fyf4sr/protect_crud_api_against_thirdparty_app_creation/)
- url: https://www.reddit.com/r/golang/comments/fyf4sr/protect_crud_api_against_thirdparty_app_creation/
---
I'm building a GO CRUD API that I would like to protect as much as possible.  I'm already issuing JWTs and providing access to certain handlers based on that user's level of access.  The problem is that the API is exposed and anyone can sniff out the calls I'm making, reverse engineer how JWTs are created and create their own front end, or automated tool to perform unauthenticated/authenticated calls to my API.  I'd like to avoid third-party apps from using the API, even if they properly authenticate a user and gain a proper JWT.

My first thought is an API key for the web front ends and mobile apps (when that part comes). But even if we're sending the key via HTTPS with the other payload, that can still be sniffed with a MITM attack, so it's pointless. How can I appropriately verify that API calls are coming from a legitimate source before processing the request?
## [10][gRPC Go resources](https://www.reddit.com/r/golang/comments/fyj6ai/grpc_go_resources/)
- url: https://www.reddit.com/r/golang/comments/fyj6ai/grpc_go_resources/
---
What's the best way to learn best practices for gRPC Go and in depth. Always find myself guessing on how to do things. I know the basics like generating proto, calling rpc methods, streaming and implementing servers but I need a little bit more of a detailed guide. Do I just read everything in the gRPC go repo?

&amp;#x200B;

If you know an excellent repo that uses all the features of gRPC please link it.
