# go-test
A simple demonstration of how coverage works and how not. Note that both
examples use the exact same code and test, except their package definition.
Note that the different package definition requires the failing example to
import the package the test actually tests.

### working
The working example shows that coverage works as expected when the test package
is within the package it tests. Looking at the files `working/foo.go` and
`working/foo_test.go` you see that both define `package foo`.

### failing
The failing example shows that coverage does not work as expected when the test package
is not within the package it tests. Looking at the files `working/foo.go` and
`working/foo_test.go` you see that one defines `package foo` and the other `package foo_test`.

### test
Run the tests.
```
$ make
ok    _/home/vagrant/projects/private/go-test/failing 0.005s  coverage: 0.0% of statements
ok    _/home/vagrant/projects/private/go-test/working 0.003s  coverage: 100.0% of statements
```
