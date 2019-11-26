# beatrice
A simple static blog generator written in Go.

https://beatrice-example.surge.sh/

### Installation
```go get github.com/pprobst/beatrice```

Or just download this repository and run ```go build``` or ```go install```.

### How do I even...?
After "getting" the package/repo, you can initially configure your blog by editing
the file ```config.yml```. It comes with some pre-defined configs just as an
example.

There're mainly two directories you'll work with: ```about``` and ```posts```. 

```about``` will hold a markdown file regarding yourself, while ```posts```
will, of course, hold your posts' markdown files. You can get started by
looking at the pre-defined files; get used at how the header in each file must
be written, and then... have some fun writing your own posts!

### Running

When you're ready, just run ```beatrice```. If everything went right, 
the message ```Fin.``` will be displayed. The generated files will be
inside the ```static``` directory. 

---

### Inspirations
* [YASBE](https://github.com/underr/yasbe/)
* [zupzup's blog generator](https://github.com/zupzup/blog-generator)

#### Todo
* Better error handling, maybe.
* Dedicated directory for images.
