# Forking gRPC!

#### Getting started.

 1. `script/bootstrap` - This builds a docker image with everything we need and starts an interactive shell in the container.

 2. `make test` - This will clone the grpc repo, build and install grpc with the php extension, build a grpc server, and start the text fixture. Subsequent runs will only invoke the test fixture.

Hopefully that worked. Note that the first time you run this, it's going to take long af. That's because the commands try to build, download and will into place all the things you need. Subsequent runs should be considerably faster.

