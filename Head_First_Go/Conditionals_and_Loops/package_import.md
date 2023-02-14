When we say math/rand we’re referring to the package’s import path, not its name. An import path is just a unique string that identifies a package and that you use in an import statement. Once you’ve imported the package, you can refer to it by its package name.

Import path 	Package name
"fmt" 		fmt
"log" 		log
"strings" 	strings

mport path and package name don’t have to be identical. Many Go packages fall into similar categories, like compression or complex math. So they’re grouped together under similar import path prefixes, such as "archive/" or "math/". (Think of them as being similar to the paths of directories on your hard drive.)


Import path 	Package name
"archive" 	archive
"archive/tar" 	tar
"archive/zip" 	zip
"math" 		math
"math/cmplx" 	cmplx
"math/rand" 	rand

Go language doesn’t require that a package name have anything to do with its import path. But by convention, the last (or only) segment of the import path is also used as the package name. So if the import path is "archive", the package name will be archive, and if the import path is "archive/zip", the package name will be zip.


Import path 	Package name
"archive" 	archive
"archive/tar" 	tar
"archive/zip" 	zip
"math" 		math
"math/cmplx" 	cmplx
"math/rand" 	rand


