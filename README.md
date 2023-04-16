# plugin-example

A) To reproduce the error, first do the following steps on linux machine

1. Clone the code
2. go to cutils directory
3. make build
4. It will generate myplugin.so, put this so under /lib on a linux box

B) Next, do the following on a Mac:

1. Clone the code
2. Under the top directory, do
   make GOOS=linux GOARCH=amd64 build
3. It will generate go binary named mytest
4. Put this binary on a linux box and run it in a window

Open another session on the same linux box and do:
curl -i -X POST http://localhost:65008/api/test/testfunction

The curl will return a 500 error. Any error returned will also be printed on the window running mytest


To run this code correctly, do the step B on the same linux box where step A was done. The curl will then not return 500, it will return 200
