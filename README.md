# beatrice
A simple static blog generator written in Go.

### Installation
```go get github.com/pprobst/beatrice```

### How do I even...?
After "getting" the package, you can initially configure your blog by editing
the file ```config.yml```. It comes with some pre-defined configs just as an
example.

There're mainly two directories you'll work with: ```about``` and ```posts```. 

```about``` will hold a markdown file regarding yourself, while ```posts```
will, of course, hold your posts' markdown files. You can get started by
looking at the pre-defined files; get used at how the header in each file must
be written, and then... have some fun writing your own posts!

### Running

When you're ready, just run ```beatrice```. The generated files will be
inside the ```static``` directory. 

If just running ```beatrice``` didn't work for some reason, try ```go install``` 
inside the directory where ```main.go``` is located and then try running
```beatrice``` again.

---

#### Todo
* Better error handling, maybe.
* Umm...
