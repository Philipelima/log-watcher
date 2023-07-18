# Log-watcher 
<p align="center">
<img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/7e2140c3dbcd227ed71182dfbf5ea6ccca920969.png" height="300">
</p>
<br>

<p align="center"> 
    <b>Just a simple program to watch every single change in a log file.</b>
</p>

<br>

It's requires: **Go 1.20**


<br>

### How to Install it:

<br>

First of all, clone this repository to your machine:

~~~bash
  git clone https://github.com/Philipelima/log-watcher.git
~~~

After run the command:

~~~bash
  GOOS={os} GOARCH={arch} go build -o log-watcher
~~~

<br>

### How to Use it:

<br>

In development mode, on the log-watcher dir: 

~~~bash
  go run ./cmd {filename}
~~~
<br>
After compile, on the application dir:

~~~bash
   ./log-watcher {filename}
~~~