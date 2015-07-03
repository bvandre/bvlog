# bvlog
Logging library for go

http://godoc.org/github.com/bvandre/bvlog

Currently the library will set up a logger first trying to
setup the logger with journald.  It then tries to setup the
logger to syslog if journald isn't found.

# Future

Will add additional logging backends, including stdout, file,
and possibly windows logging services.
